# 中间件与 NCraft 类库

下面的路径均在 `github.com/ncraft-io/ncraft/go/pkg/` 下，除特别注明。先查目标依赖版本的签名与已有消费者，再选能力。

## Endpoint middleware

扩展已有 `handlers/middlewares.go` 的 WrapEndpoints，不另建同名函数。可用 `in.WrapAllExcept(mw, "MethodName")`、`WrapAllLabeledExcept` 或单 endpoint 包装；排除名称使用生成的 Go 方法名。最后包装的最外层最先执行。

```go
// 插入现有 WrapEndpoints，假设当前服务有 CreateProduct。
in.CreateProductEndpoint = middleware.NewJWT()(in.CreateProductEndpoint)
in.CreateProductEndpoint = middleware.EveryRateLimiter(time.Second, 20)(in.CreateProductEndpoint)
return in
```

这里 time.Second 的参数表示每个令牌的补充间隔，20 是 burst，不是 20 QPS；若需要 20 QPS 可用 `NewTokenBucketLimitMiddleware(rate.NewLimiter(20, 20))`。middleware 的构造放在装配时，不能每次请求重建 limiter。

| 能力 | API 与行为 |
| --- | --- |
| 访问日志 | `middleware.Logging(log.Logger)`；options 的 logger 是 Go kit logger，业务日志用 `ncraft/logs` |
| JWT | `middleware.NewJWT()`，读取 `ncraft.jwt` / `jwt`；未启用或构造失败可能直通，要求鉴权的服务须在启动检查配置且测试无 token 会拒绝 |
| 限流 | `NewTokenBucketLimitMiddleware(*rate.Limiter)` 或 `EveryRateLimiter(interval, burst)` |
| 指标 | `middleware.Instrumenting(histogram, counter)`；生成 options 可提供 count / latency，保持标签数与注册一致 |
| tracing | 生成代码使用 Go kit OpenTracing / NCraft tracing；保持已有 tracer 的传播与生命周期 |

JWT 示例配置用 `enable: true`、`key: ${JWT_KEY}`，并按实际需求配置有效期和 exceptions；当前实现只接受 HS256，不能仅改 algorithm 字段就声称支持 RSA。只改 WrapService 不会改变生命周期所在的原始实例。

新增方法时旧 middlewares 文件保留，逐方法 tracing/metrics 代码也不会自动增加。调用 WrapAll 会涵盖新 endpoint，但需检查排除列表与业务权限，不要误把所有方法设为公开。重试可能重复执行写入，按操作幂等性启用。

## 下游客户端与服务发现

优先使用生成 `client-go/pkg/catalog-client`，不要直接依赖下游的 service-go：

```go
client, err := catalogclient.NewClientFromServiceConfig(ctx, "catalog")
if err != nil { return err }
defer client.Close() // 在 Start 中持有时改由 Shutdown 关闭。
```

示例选择 HTTP，配置同时复用全局 sd 与客户端通用部分：

```yaml
sd:
  mode: direct
  transport: http
  direct:
    shop.v1.Catalog:
      urls: [http://localhost:21101]
client:
  http:
    timeout: 5000
    headers: [Authorization, X-Request-ID]
  catalog:
    serviceName: shop.v1.Catalog
    http:
      timeout: 1000
```

注册名须读生成的 `FullServiceName`，不要按 Go 目录名猜。`NewClient(ctx, *sd.Config, options...)` 只用明确参数；`NewClientFromConfig` 用 FullServiceName 自动选择配置；`NewClientFromServiceConfig` 用别名。

发现支持 direct / etcd / Nacos、实例变化和 round-robin；默认 transport 是 gRPC，需显式提供 `WithGRPCDialOptions(grpc.WithTransportCredentials(...))`。开发环境确需明文时明确用 insecure credentials；配置文件不能代替凭据构造。`WithGRPCOptions` 是 per-call options，不是 dial options。

直接 HTTP 用 `NewHttpClient`；`NewGrpc(connection, ...)` 的连接归调用者；`NewGrpcClient` / discovery client 持有连接，结束后 Close。可用 WithHTTPClient、WithMiddleware、CtxValuesToSend、WithEnvelope。HTTP envelope 必须与服务配置一致。HTTP 只用每方法首个 binding，无 HTTP binding 的方法不能通过 HTTP 调用。

重试默认关闭；`retry.max` 是总尝试次数，timeout 为毫秒。endpoint middleware 包裹整个逻辑调用及其重试。只有可安全重试的操作才启用。

## Redis

接口 `ncraft/redis.Redis`，配置 `redis.Config`；使用实现需要注册 import：

```go
import (
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/redis"
    _ "github.com/ncraft-io/ncraft/go/pkg/ncraft/redis/goredis"
)
```

加载并校验配置后 `redis.New(cfg)`，返回 Redis（无 error 返回值），不存在 implementor 会 panic；默认 implementor 为 goredis，空 connections 可能默认 `:6379`，不能把配置缺失误当可用。初始化后用带超时 context 的 PING 做必要连通性检查，Shutdown 调用 Close。

`Do(ctx, "GET", key)` 与转换工具 `redis.String` / `Bytes` / `Int` 等使用前检查实际签名；缓存对象由 jsoniter 编解码。NewBatch / RunBatch 是批处理接口，不等于事务，多 key 聚合不保证原子性；需要 Lua/事务/cluster 特性先确认实现支持。

## 对象存储

使用 `ncraft/storage.Config`，以 `_ ".../ncraft/storage/minio"` 注册 MinIO / S3 实现。优先 `NewStorageWithContext(ctx, cfg) (Storage, error)` 显式处理初始化错误；旧 NewStorage 失败仅返回 nil，全局 GetStorage 还可能回退 DummyStorage。

复用 endpoint、bucketName、region、sessionToken、secure、prefix、createBucket；明确 URL 的 http/https 优先于 secure。`createBucket` 默认为兼容行为 true，访问已有 bucket 且无建桶权限时设 false。

`Storage` 提供 Read / Write / Upload / Download；需要可取消、可 seek 大文件时断言 `storage.StreamingStorage`，调用 `Open(ctx, key)` 并关闭返回 reader，或 `WriteContext`。不要假造 Storage.Close；各资源按实际接口回收。

## 消息

使用 `ncraft/messaging`，显式注册实现 `_ ".../ncraft/messaging/nats"`。`NewClient()` 读取 `messaging` 然后 fallback `ncraft.messaging`；`NewClientWith(cfg)` 返回 `(*Client, error)`。Publish 传 context，服务拥有的 client 在 Shutdown 调用其 Shutdown。

带订阅的 Mojo 接口可能已经生成消费装配，先检查 run.go，避免又创建一套重复消费者。手写消费者按处理结果 Ack / Nak / Term，并验证实际 provider 的重投递规则、幂等性。消息 Data 是约定的 JSON 字符串时用 jsoniter；不要把 payload 序列化两次。业务处理失败与消费者后台致命失败分开处理。

## KV、日志与辅助类库

| 能力 | 选择与注意事项 |
| --- | --- |
| 嵌入式 KV | `ncraft/kvstore.KvStore`；需要 Bolt 时可用 `kvstore/bbolt.NewStore(Options) (KvStore,error)`，确认文件路径和 Close。badgerdb 等实现按需引入；TiKV 是独立 Go 子模块，不能假定主模块已包含 |
| 业务日志 | `ncraft/logs.Infow/Errorw` 等结构化日志；`gokit/kit.Logger()` 是 Go kit adapter；不要每请求创建新 logger |
| 过滤 | `gokit/filter` 是过滤工具；数据库查询使用 Mojo `db/query`，不可互相替代 |
| 分页 | `gokit/pagination` 具有不同 token 类型/接口；按业务选一个协议，不与 DB 查询的数字页索引混用 |
| 类型与错误 | Mojo `core.Timestamp` / Duration / Uuid / FieldMask / Ordering / Options / Error，沿用已有 Parse / Format / codec，不另定义不兼容版本 |

接口签名、配置键和关闭方法以锁定版本源码为准。只引入所需后端，避免为使用接口包就引入全部外部驱动。
