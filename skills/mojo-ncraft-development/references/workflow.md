# 项目初始化与接口迭代

## 创建框架

准备项目指定的 Mojo CLI、符合依赖要求的 Go、`protoc`、`protoc-gen-go`、`protoc-gen-go-grpc`。不要以生成模板中的 `go 1.20` 推断所有依赖只需要 Go 1.20，也不要默认安装 `@latest`。

若使用本地 Mojo 源码构建 CLI，在该仓库 `go/` 执行 `go build -o /tmp/mojo-dev ./cmd/mojo`，之后用这个明确路径执行命令；不要误用同名的 Modular Mojo。

从目标父目录创建项目（仓库 URL 决定模块路径）：

```sh
mojo create -p shop -r https://github.com/acme/shop
cd shop
mkdir -p mojo/shop/v1
```

显式 `-p` 创建空项目；省略包名进入 Hello World 模式。`-o` 指定父目录。沿用生成的 `package.mojo`，维护版本、仓库及依赖；不要凭空构造另一种 manifest 格式。

## 类型、接口与文本注释

- 服务相关 Entity（`type` 结构体）优先定义在业务包的非版本目录，例如 `mojo/shop/product.mojo`；`v1` / `v2` 目录一般用于服务 `interface`，例如 `mojo/shop/v1/catalog.mojo`，由不同版本的接口复用同一业务 Entity。
- 新 Entity 通常将 ID 放在首个字段，Mojo 源码按现有约定写作 `id`，如 `id: String @1`；生成 Go 字段名以实际代码为准（当前通常为 `Id`）。已有协议的字段编号和包名保持兼容，不为调整布局擅自移动已发布类型或重新编号。
- `type` 和 `interface` 的文本说明写在声明之前，使用 `///`；多行说明的每一行都以 `///` 开头。有 `@entity` 等标注时，文本注释放在标注及其声明之前。
- `type` 内部字段、`enum` 成员的文本说明可以采用同一行末尾的 `//<` 尾追注释。

`mojo/shop/product.mojo`：

```mojo
/// 商品
@entity
type Product {
    id: String @1 //< 商品 ID
    name: String @2 //< 商品名称
    description: String @3 //< 商品说明
}
```

`mojo/shop/v1/catalog.mojo`：

```mojo
/// 商品服务
interface Catalog {
    @http.get("/shop/catalog/v1/products/{id}")
    get_product(id: String @1) -> Product

    @http.post("/shop/catalog/v1/products")
    create_product(product: Product @1 @http.body) -> Product

    @http.put("/shop/catalog/v1/products/{id}")
    update_product(id: String @1, product: Product @2 @http.body) -> Product
}
```

目录决定包层次：这里是 `shop` 与 `shop.v1`，无需每个文件重复 package 清单。业务增多可拆 `shop.catalog/v1`、`shop.account/v1`。服务目录名由 interface 名派生，去除 Service / Server 后缀并转为 kebab-case；只创建包目录不会定义服务。

路由 `{id}` 对应同名参数；`@http.body` 参数直接成为请求体，未绑定路径或 body 的普通参数通过 query 传入。多个 body 字段应合为一个类型。上面的创建请求体是 `{"id":"p-1","name":"Phone"}`，没有额外的 `product` 包裹层。

## 生成及模块边界

在 `package.mojo` 所在根目录执行：

```sh
mojo build -t api,service,client .
```

| 目标 | 主要产物 |
| --- | --- |
| `api` | `document/`、`openapi/`、`protobuf/`、`go/` 中的协议、类型和 codec |
| `service` | `service-go/cmd/`、`internal/`、`pkg/*-service/`、`pkg/model/` |
| `client` | `client-go/pkg/*-client/`，不依赖服务实现 |

命令目标是 `service` / `client`，不是目录名 `service-go` / `client-go`。无需客户端可用 `api,service`；只更新服务或客户端时先确认 API 已同步。首次不指定目标默认 `api`，存在 `service-go/` 后默认 `api,service`，因此需要客户端时显式传 `client`。

需要 Entity 对应的 SQL 操作代码时，显式执行 `mojo build -t api,service .`（需要客户端则加 `client`）。`service` 目标会为识别到的 Entity 在 `service-go/pkg/model/` 下生成基于 GORM 的 CRUD；仅生成 `api` 不会得到该后端 model。用户可在同一目录新建 `product_queries.go` 等手写文件，扩展生成模型的数据库操作；不要修改 `product_model.go` 等生成文件。数据库标注与扩展示例见[数据库](database.md)。

生成后检查各模块的 `go.mod`，保留共享 API 的相对 `replace .../go => ../go`（若项目采用此布局）。本地协作可用项目既有 workspace；不要提交机器专属绝对路径。模板可能残留 `github.com/ncraft-io/ncraft-gokit` 等历史依赖，而代码 import 已是 `github.com/ncraft-io/ncraft/go`：依实际源码选择经过验证的兼容版本，再整理依赖，不要只执行无版本的 `go get` 让它漂移。

当前 build 会在生成过程中调用各模块的 `go mod tidy`。新脚手架的历史依赖或离线缓存不足可能使它在 API / 服务端已写入后退出，后续客户端尚未生成。保留已生成的源码，修正对应模块的版本/依赖解析后重跑同一命令；不要因看到输出目录就当作全部成功，也不要为了重试删除 handlers。记录最终命令退出状态与各模块构建结果。

`service-go/pkg/catalog-service/handlers/` 实现协议；`pkg/catalog/` 放领域逻辑；`pkg/model/` 同时容纳生成的实体 CRUD 和独立手写的数据库扩展。依赖、连接和配置不应放入共享 API 包。

## 后续接口更新

1. 查看 `git status` 和已有差异，确定源码、生成产物、手写代码。保留用户改动。
2. 改 `.mojo`。保持已发布编号、字段语义、路由与已有请求体兼容；不兼容变更考虑新的版本包。实体结构变更同时规划迁移，不能只看 API 是否编译。
3. 执行明确目标的 `mojo build -t api,service,client .`（按项目实际目标缩减）。
4. 查看生成 diff，手动调整保留方法的签名、响应构造、调用方及测试。处理废弃方法，检查手写与新生成方法是否重复。
5. 逐项补充新增方法的鉴权、限流、tracing、metrics。已有 `middlewares.go` 和 `hooks.go` 会整体保留，新增方法不会自动获得旧模板中的逐方法装配。
6. 在各 Go 模块内运行相关测试与构建；有新契约时用 HTTP/gRPC 的实际响应验证 body、错误和 JSON 格式。需要验证再生成保留性时，在临时副本重复生成并对比手写文件。

生成器会扫描 handlers 包的非测试 Go 文件，保留方法并补缺失方法，包括移到 sibling 文件的实现。它不是签名迁移器：改入参/返回类型仍需手动改。`svc/`、`internal/*-server/`、生成 model 是可覆盖产物；不要删除生成标记来阻止更新。

## 本地运行与验证

生成服务框架后，按实际组件从本 skill 的 `configs/` 复制并定制配置；已有同名文件比较后合并。可用的组件模板和复制工具见[配置](configuration.md)，不要把其他业务工程的配置整包搬入。

从模块内执行，不要在只有 `package.mojo`、没有 `go.mod` 的根目录执行 `go test ./...`：

```sh
cd service-go
go test ./...
go build -o bin/catalog-server ./cmd/catalog-server
./bin/catalog-server
```

从 `service-go/` 启动可发现 `configs/`；也可明确指定 `CONFIG_PATH`。监听地址读实际 `server.yaml`。API 与 client 模块也应分别构建；集成测试需要的外部设施应明确列出。刚生成的 handler 返回空响应只是占位，必须实现后才能声称服务可用。
