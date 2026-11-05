# 服务实现与可定制扩展点

## 框架集成原则

微服务的 HTTP 入口优先使用 Mojo / NCraft 生成的服务框架，复用其监听、路由、endpoint、中间件、错误编码和启动退出流程。业务实现放在服务 handler 中，特殊需求通过已支持的扩展接口接入。

不要因实现方便另写一套 `http.Server` / `ListenAndServe` 或独立 Web 应用，再靠额外端口、替换框架主 router、挂载覆盖全部请求的 catch-all handler，或补丁修改生成的 `run.go` / transport 强行接入。框架版本缺少需要的扩展点时，先核对可用扩展方式和兼容升级路径。仅在项目已有明确的独立 HTTP 服务架构或任务明确要求时，按该架构处理，并明确监听、中间件与生命周期的归属；这不是普通业务接口的默认实现方式。

## 业务实现的位置

保留工厂 `NewService() pb.CatalogServer` 的签名；有状态实现改为返回 `&catalogServer{}`，通过字段持有模型、配置和客户端，方法使用指针 receiver。嵌入 `pb.UnimplementedCatalogServer` 以适配生成的 gRPC 接口。不要再定义一个相同的 NewService 或 server 类型。

Handler 校验请求和业务约束，将 context 向 DB / 客户端传递。业务可测试构造器可命名为 `NewServiceWithConfig` / `NewServiceWithDependencies`，返回实例和 error，再由工厂及生命周期接入。正常请求错误从 RPC 方法返回，不能发送到后台致命错误通道。

使用 Mojo 错误语义，例如 `core.NewInvalidArgumentError`、`NewNotFoundError`、`NewAlreadyExistsError`、`NewInternalError`。`gorm.ErrRecordNotFound` 用 `errors.Is` 判断并映射。客户端可见错误不暴露 DSN；内部日志保留诊断信息。HTTP envelope 可能改变 HTTP 状态码，不能只断言错误文本或 status，需核对运行时配置；gRPC 错误映射也以实际传输层为准。

## 扩展点表

下表是本 skill 基线版本的调用契约；先确认目标生成 host 和运行时支持。

| 扩展点 | 用法与限制 |
| --- | --- |
| `NewService() pb.XServer` | 保留签名，返回服务实例；用 Start 报告初始化错误 |
| `Start(nserver.Config) error` | 实现 `gokit/server.Starter`，监听前初始化 DB、配置、后台任务；失败时自行回收已创建资源 |
| `Shutdown(context.Context) error` | 实现 `Shutdowner`，HTTP/gRPC 排空后释放资源；已成功启动服务按逆序退出，遵守 ctx 的截止时间 |
| `Errors() <-chan error` | 实现 `ErrorSource`，报告后台任务致命失败；稳定的服务自有 channel，至少缓冲 1，后续发送避免阻塞；nil、关闭或 nil error 不代表退出 |
| `RegisterHTTPHandlers(*mux.Router)` | 在原始服务实例上向框架提供的同一 router 注册必要的额外路由，早于生成路由；限于声明及编解码扩展无法覆盖的专用行为，不另建 HTTP 服务或整体替换主路由 |
| `WrapService(pb.XServer, map[string]interface{}) pb.XServer` | 服务级 wrapper/decorator，保持完整 Server 接口 |
| `WrapEndpoints(svc.Endpoints, map[string]interface{}) svc.Endpoints` | endpoint 中间件；全量、按名称排除或单方法包装 |
| `nhttp.RequestDecoder` | `DecodeHttpRequest(*http.Request) error`；在请求类型上实现，替代整个自动解析流程，需自行处理 path/query/body |
| `nhttp.ResponseWriter` | `WriteHttpResponse(context.Context, http.ResponseWriter) error`；响应对象自行写 HTTP |
| `nhttp.BindResponseWriter(ctx, response, writer)` | 给一次请求的具体非 nil 响应指针绑定 HTTP writer，不改变 protobuf 返回类型 |

生成的 `NewEndpoints(options, services...)`、`RegisterService(cfg, router, grpcServer, services...)` 可以接收注入的业务服务实例，便于测试和复用框架装配；这里的服务实例是 `pb.XServer` 实现，不是另一套 HTTP server。若维护项目已有或任务明确要求的自定义 host，其装配需负责 Start / Shutdown / Errors，不能假设调用 RegisterService 会启动生命周期。

生命周期在原始实例上执行，先于 `WrapService` / endpoint 包装。模板的默认工厂原本返回值；如果只给指针定义 Start 而工厂仍返回值，类型断言不会命中。新增扩展后用编译期断言和集成测试确认：

```go
var _ nserver.Starter = (*catalogServer)(nil)
var _ nserver.Shutdowner = (*catalogServer)(nil)
// 使用后台任务时才实现：
// var _ nserver.ErrorSource = (*catalogServer)(nil)
```

当前生成 host 统一处理信号、transport error 和后台错误；共享 30 秒退出预算，先停止传输再回收服务。老 `handlers.InterruptHandler` 虽可能保留在磁盘，当前 host 不再调用它。把业务清理移到 Shutdown，不要新增另一个 signal goroutine。旧项目未生成这些调用时应同步升级生成器/运行时并再生成，或使用其已有、经验证的启动机制。

## HTTP 自定义

按需求选择最小的框架扩展：

1. 普通业务 HTTP 接口：在 Mojo interface 中定义方法与 `@http.*` 路由，再生成 transport 并实现 handler。
2. 业务包装、鉴权、限流和观测：使用 `WrapService` / `WrapEndpoints`，复用现有调用链。
3. 特殊请求体、响应格式、下载或流式响应：优先使用 `RequestDecoder`、`ResponseWriter`、`BindResponseWriter`，保留生成路由和 endpoint；Range 下载也不应直接成为另起一个 HTTP 服务的理由。
4. 仍有无法覆盖的专用路由：通过 `RegisterHTTPHandlers` 注册到框架提供的同一 router，限定方法及路径，避免遮蔽无关生成路由。

专用路由不会自动经过 `WrapEndpoints`；需要它的 handler 接上所需鉴权、限流和观测，并继续由框架统一管理监听与退出。

固定 protobuf 返回值要下载文件时，可使用绑定 writer：

```go
// 在业务方法里；response 是准备返回的同一个 protobuf 指针。
if _, ok := nhttp.RequestFromContext(ctx); ok {
    if err := nhttp.BindResponseWriter(ctx, response, writer); err != nil {
        return nil, err
    }
}
return response, nil
```

生成 HTTP transport 必须通过 `RequestToContext` 安装请求 attachment。编码优先级为绑定 writer → 响应自身的 ResponseWriter → 普通 JSON/envelope。绑定按指针身份匹配，替换/复制响应会失效；RPC 不携带这种 HTTP 状态。

文件/网络流资源只在 `WriteHttpResponse` 内打开并在返回前关闭，避免 middleware 丢弃响应后泄漏。发出 headers 后发生传输错误只记录，不能返回会导致 transport 再附加 JSON 的错误。endpoint metrics 结束于响应编码之前，大文件传输耗时需在 HTTP 层统计。

自定义 RequestDecoder / ResponseWriter 方法应放在**定义该类型的包**中的独立手写文件。Go 不允许在服务包为另一个模块的外来类型添加方法；本地 API 模块可添加独立文件，外部依赖用适配器或 BindResponseWriter。不要为规避这点复制 protobuf 类型或修改生成 `.pb.go`。

## 验证重点

- HTTP 需求接入了框架调用链或明确的扩展点；普通业务没有额外启动 HTTP 服务、整体替换主路由或修改生成启动代码。
- 工厂返回的实际实例实现需要的生命周期与路由接口，HTTP/gRPC 共享该实例。
- Start 中途失败能回收资源；停止接收请求后才关闭 DB / Redis。
- 新增 endpoint 实际被需要的 middleware 包装；自定义路由的权限策略有测试。
- RequestDecoder 覆盖 path/query/body 和错误；自定义 writer 验证 pointer identity、取消与资源关闭。
