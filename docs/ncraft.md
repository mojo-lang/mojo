# NCraft 微服务开发指南

Mojo 用于定义数据结构和服务接口，并生成 API 文档、协议代码、服务端和
客户端框架。NCraft 微服务基于 Go kit，运行时类库位于
`github.com/ncraft-io/ncraft/go`。开发时先编写 Mojo 定义，再生成代码，
最后在服务端实现业务逻辑。

`type`、`interface`、标注和注释的基本写法见 [Mojo 简明语法](syntax.md)。

以下以 `shop` 项目下的 `catalog` 商品服务为例。运行前准备好 Mojo CLI、
Go、`protoc`、`protoc-gen-go` 和 `protoc-gen-go-grpc`，并将命令加入 PATH。
Go 版本需满足生成项目及其依赖的 `go.mod` 要求。

## 1. 创建项目框架

在准备存放项目的父目录执行：

```sh
mojo create -p shop -r https://github.com/acme/shop
cd shop
```

将示例仓库地址替换为自己的地址，它用于确定生成代码的模块路径。
命令会创建 `shop/`，其中包含 `package.mojo`、`README.md`、`.gitignore`
和空的 `mojo/shop/` 源码目录。`-o` 可以指定存放项目的父目录。
显式传入包名会创建空框架；不传包名则进入 Hello World 示例模式。

`package.mojo` 是项目的包清单，记录包名、版本、作者、仓库和依赖。
本例的主包名为 `shop`，应保留创建命令生成的清单并按项目需要维护。

## 2. 组织包并定义 Entity

`mojo/` 下按 package 名称建立目录。简单项目可以直接使用主包：业务类型
放在 `mojo/shop/`，服务 Interface 放在 `v1/`、`v2/` 等版本目录下。
本指南后续示例采用以下无业务子包的结构：

```text
shop/
├── package.mojo                 # 声明主包 shop
└── mojo/
    └── shop/                   # shop
        ├── product.mojo
        └── v1/                 # shop.v1
            └── catalog.mojo
```

目录层次确定源码所属的包，不需要在每个 `.mojo` 文件中重复根清单的
package 声明，版本目录用于区分 API 版本。

业务增多时，可以在主包下组织多个业务子包，例如 `shop.catalog`、
`shop.account`，将各自的类型和版本化服务定义放在对应子包内：

```text
shop/
├── package.mojo                 # 声明主包 shop
└── mojo/
    └── shop/                   # shop
        ├── catalog/            # shop.catalog
        │   ├── product.mojo
        │   └── v1/             # shop.catalog.v1
        │       └── catalog.mojo
        └── account/            # 可选的其他业务子包
            └── v1/
```

对于本指南采用的简单结构，只需要创建 `v1` 目录：

```sh
mkdir -p mojo/shop/v1
```

在 `mojo/shop/product.mojo` 中定义商品实体：

```mojo
/// 商品信息
@entity
type Product {
    id: String @1 //< 商品 ID，由业务方分配
    name: String @2 //< 商品名称
    description: String @3 //< 商品说明
}
```

`type` 定义结构体，`@entity` 标记数据库实体；`id` 会被识别为主键。
普通请求对象、响应对象及其他关键结构也可以使用 `type` 定义。
不需要数据库模型的结构体应避免被识别为 Entity；具体的自动识别、主键
和关联规则见 [Entity 数据库操作代码](entity-models.md)。

字段后的 `@1`、`@2` 等是协议字段编号，同一结构内不能重复。已发布的
编号应保持稳定，避免调整顺序时改变编号或复用已删除字段的编号。

## 3. 在版本目录中定义服务

在 `mojo/shop/v1/catalog.mojo` 中编写 Interface：

```mojo
/// 商品服务
interface Catalog {
    /// 根据商品 ID 查询商品
    @http.get("/shop/catalog/v1/products/{id}")
    get_product(id: String @1) -> Product

    /// 创建商品，请求体为商品对象
    @http.post("/shop/catalog/v1/products")
    create_product(product: Product @1 @http.body) -> Product

    /// 更新指定商品
    @http.put("/shop/catalog/v1/products/{id}")
    update_product(id: String @1, product: Product @2 @http.body) -> Product
}
```

服务可以引用父包中的类型，本例中的 `Product` 定义在父包
`shop`。Interface 中声明函数名、输入参数和返回类型，生成器据此
生成请求／响应类型及服务方法。

服务定义中的标注含义如下：

| 写法 | 含义 |
| --- | --- |
| `/// ...` | 类型或函数的说明文本，用于生成文档 |
| `//< ...` | 字段或参数后的说明文本 |
| `@http.get("...")` 等 | HTTP 方法与路由，支持相应的 POST、PUT、DELETE 等标注 |
| 参数后的 `@1`、`@2` | 生成请求消息时的字段编号，编号在每个方法的输入参数内独立 |
| 路由中的 `{id}` | 与同名输入参数绑定，通过 URL 路径传入 |
| `@http.body` | 该参数通过 HTTP 请求体传输 |

没有绑定到路径且没有标记为请求体的普通参数通过 query 传入。
例如增加 `keyword: String @2` 后，可通过 `?keyword=phone` 传值。
一个请求体应对应一个 body 参数，需要多个字段时将它们放进同一个结构体。

本例创建接口的请求体直接对应 `Product`，形如
`{"id":"p-1","name":"Phone","description":"Demo"}`。
更新接口同时包含路径参数和请求体；路径 ID 与实体 ID 的一致性由业务逻辑校验。

## 4. 生成文档、API、服务端和客户端

在包含 `package.mojo` 的项目根目录执行：

```sh
mojo build -t api,service,client .
```

各目标及对应的输出为：

| 目标 | 输出目录 | 内容 |
| --- | --- | --- |
| `api` | `document/` | Markdown 类型和服务文档 |
| `api` | `openapi/` | OpenAPI 服务说明及结构体 schema |
| `api` | `protobuf/` | Protobuf 定义 |
| `api` | `go/` | Go API 类型、序列化代码及 gRPC stub |
| `service` | `service-go/` | Go kit 服务框架、handlers、启动入口和 Entity 模型 |
| `client` | `client-go/` | HTTP、gRPC 等客户端封装 |

`service-go`、`client-go` 是输出目录名，命令中的目标名称是 `service`、
`client`。只需要服务端时使用 `mojo build -t api,service .`；已有 Go API
时可用 `mojo build -t service .` 或 `mojo build -t client .` 更新对应部分。
修改接口后应同时更新 API 和依赖它的服务端／客户端。

未指定 `-t` 时，首次构建默认只生成 `api`；项目根目录已有 `service-go/`
时默认生成 `api,service`。需要客户端时应显式指定 `client`。

## 5. 实现服务业务逻辑

本例主要使用以下目录：

```text
service-go/
├── cmd/catalog-server/          # 服务启动入口
├── configs/                     # 运行配置
├── internal/catalog-server/     # 生成的服务装配代码
└── pkg/
    ├── catalog-service/
    │   ├── handlers/handlers.go # 业务接口实现
    │   └── svc/                # 生成的 endpoint 和传输层
    ├── model/                  # Entity 数据库 CRUD
    └── catalog/                # 按需手工建立的业务包
```

一般在 `service-go/pkg/{name}-service/handlers/handlers.go` 中实现业务。
目录中的 `{name}` 来自 Interface 名称（去除 `Service`／`Server` 后缀并
转换为短横线格式）；本例使用 `Catalog`，因此生成 `catalog-service` 和
`catalog-server`。按子包划分服务时，可以让 Interface 与子包使用对应的
名称；仅创建子包目录不会自动定义服务。

生成的方法接收 `context.Context` 和请求对象，并返回响应对象及 `error`。
在这些方法中完成参数校验、数据库访问和业务调用，按生成的响应类型填写结果。
初始 handlers 只提供方法框架，需要实现业务后才会产生相应的业务行为。

数据库 Entity 的 CRUD 代码生成在 `service-go/pkg/model/`，底层依赖
GORM。模型构造函数接收 `*db.DB`，可复用服务管理的数据库连接。业务需要
自行初始化、注入连接并管理其生命周期；模型是否执行 AutoMigrate 由
`DisableAutoMigrate` 控制。具体构造函数、CRUD 方法和使用示例见
[Entity 数据库操作代码](entity-models.md)。

可复用的业务代码放在 `service-go/pkg/{package-name}/`，例如本例的
`pkg/catalog/`。handler 负责接入请求，业务包负责商品规则及业务流程。

重新生成时，生成器会保留 handlers 包中已有的方法实现，并补充缺失方法；
已有方法的签名与新接口不兼容时，需要手动调整。带有 `Code generated ...
DO NOT EDIT.` 标记的纯生成文件会被更新，业务扩展应放在手写文件中。

## 6. 配置、启动和本地验证

常用配置位于 `service-go/configs/`，包括 `server.yaml`、`logs.yaml`、
`metrics.yaml`、`tracing.yaml` 和 `sd.yaml`。根据实际运行环境配置监听地址、
日志、监控、链路追踪和服务发现；业务使用的数据库配置也应由服务加载。

从 `service-go/` 目录启动服务，便于加载当前工作目录下的 `configs/`：

```sh
cd service-go
go run ./cmd/catalog-server
```

NCraft 配置加载器会查找当前工作目录下的 `conf/`、`configs/`，也支持
通过 `CONFIG_PATH` 指定配置目录。配置路径以进程工作目录为基准。
服务器部署时，可将 `configs/` 放在启动工作目录下，并保留实际需要的配置。

可以在 `service-go/` 下测试并编译服务：

```sh
go test ./...
go build -o bin/catalog-server ./cmd/catalog-server
./bin/catalog-server
```

服务启动后，根据实际监听地址调用接口。例如使用生成的 `server.yaml`
中的 HTTP 端口 `21101`：

```sh
curl -X POST http://localhost:21101/shop/catalog/v1/products \
  -H 'Content-Type: application/json' \
  -d '{"id":"p-1","name":"Phone","description":"Demo"}'
curl http://localhost:21101/shop/catalog/v1/products/p-1
```

以上创建和查询操作需要先实现对应的 handler 及存储逻辑。

## 7. JSON 与运行时依赖

服务定义的结构体及 HTTP 传输层的 JSON 编解码依赖 jsoniter
（JSON Iterator，Go 模块为 `github.com/json-iterator/go`）。生成的 JSON
代码包含 Mojo 类型所需的编码规则，手写 JSON 处理时应与生成代码保持一致。

NCraft 的 Go kit 服务、客户端、配置、日志等运行时依赖位于
`github.com/ncraft-io/ncraft/go`。日常迭代流程是：修改 Mojo 定义，重新执行
`mojo build -t api,service,client .`，补充或调整业务实现，再测试并启动服务。
