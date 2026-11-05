# DB / GORM 使用

## 类型与实体

数据库包装器是 `github.com/mojo-lang/mojo/go/pkg/mojo/db`，不是臆造的 `ncraft/db`。`db.DB` 嵌入 `*gorm.DB` 并持有 `*db.Config`；模型在 `service-go/pkg/model`。

- 服务 Entity（`type` 结构体）优先放在非版本业务目录，例如 `mojo/shop/product.mojo`；服务接口一般放在 `mojo/shop/v1/`、`v2/`，复用这些 Entity。新 Entity 通常以 ID 为首字段，Mojo 中写作 `id`，如 `id: String @1`。这是组织约定，已有特殊主键或已发布协议按实际兼容要求处理。
- `id` 或 `@key` / `@db.key` / `@db.primary_key` 会识别实体主键；`@entity` 没有主键时补 String id，但不会自动生成 ID 值。
- 生成模型支持单 String / 整型主键；构造时核对实际 GORM 主键列。普通 DTO 中的 id 也可能触发实体识别，需检查生成结果。
- `NewProduct(database)` 返回 `(*model.Product, error)`。配置未禁用时会 AutoMigrate；生产环境通常显式 `disableAutoMigrate: true` 并使用项目迁移流程。是否改变 schema 应遵从任务和环境约束。
- `@db.json` 字段作为 JSON 列，通过生成的 Scanner / Valuer 与 jsoniter 工作，不是关联预加载。真实关联默认加载直接层级，不递归加载所有后代。

## 数据库相关标注

类型标注放在声明前；字段标注放在字段声明中，文本说明仍可用行末 `//<`。以下行为已按当前 Go 生成链核对；使用其他版本时检查生成结果。

| 标注 | 位置与用途 | 生成行为及边界 |
| --- | --- | --- |
| `@entity` | `type` 声明前，显式标记 Entity | 纳入实体模型生成；不负责分配业务 ID |
| `@key`、`@db.key`、`@db.primary_key` | 主键字段上，通常择一使用 | 生成 GORM `primaryKey`；有 `id` 时实体识别优先选 `id`，不要同时指定另一个冲突主键；当前生成模型要求单主键 |
| `@db.json` | 需要整体存入 JSON 列的对象、结构体、数组或 map 字段 | 生成相应的数据库 Scanner / Valuer 等支持，编解码使用 jsoniter；Entity/Entity 数组字段使用此标注后按 JSON 存储，不为该字段生成关联 Preload |
| `@db.ignore` | 仅用于计算、展示等不落库的字段 | 生成 GORM `-`，忽略持久化；不会自动从 HTTP / JSON 响应中删除该字段 |
| `@db.index` | 需要普通索引的字段 | 当前 Go 生成器输出普通 GORM `index`；不能将 `@db.index("unique")` 当作已实现的唯一索引或完整索引参数透传 |
| `@db.foreign_key("customer_id")` | 关联字段上，指定关联使用的外键字段 | 当前生成器转换为 GORM `foreignKey:CustomerId`；关联实体、外键字段位置与主键类型须符合实际关系，生成后核对 GORM schema |
| `@max_length(128)` | String 字段，属于 core 标注但参与数据库映射 | 当前 Go 生成器对大于 0 且小于 256 的长度生成 `varchar(n)`；不能由此推断所有数据库约束或请求校验都自动生效 |

例如在非版本目录定义商品 Entity：

```mojo
/// 商品实体
@entity
type Product {
    id: String @1 //< 商品 ID，通常为首字段
    name: String @2 @db.index @max_length(128) //< 商品名称及普通索引
    description: String @3 //< 商品说明
    metadata: Object @4 @db.json //< 作为一个 JSON 值落库的扩展信息
    labels: [String] @5 @db.json //< 作为 JSON 数组落库的标签
    display_label: String @6 @db.ignore //< 计算得到的展示字段，不落库
}
```

JSON 列的实际 SQL 类型由数据库方言决定；JSON null、SQL NULL、空数组及零值分别验证。生成器可能为 `@db.json` 字段创建包装类型或命名类型，构造对象时读取生成的字段定义和初始化方法，不要假设它仍是原始 Go 类型。JSON 列不等于关系表，选择 JSON 存储时要考虑查询、索引和关联更新需求。

源码还包含 `@db.explode` 及 reference 相关定义，但本基线的 `db.explode` 路径只可确认生成 `db` tag，不能据此承诺 GORM `embedded` / `embeddedPrefix` 行为。也不要凭 GORM 的 tag 名猜造 `@db.column`、`@db.unique` 等 Mojo 标注；自定义列名、唯一/复合索引等需要使用目标版本已验证的能力或项目现有迁移与模型扩展方案。

## 生成模型与手写数据库扩展

在项目根目录显式执行 `mojo build -t api,service .`，生成 Entity 的 Go API 和基于 GORM 的 SQL CRUD 操作代码。已有同步 API 时可用 `mojo build -t service .`；只含 Entity、没有 interface 的包也可以生成 model。这里不需要另造 `sql` 或 `model` 构建目标。

```text
service-go/pkg/model/
├── db.go               # 生成的数据库与模型初始化支持
├── product_model.go    # 生成的 Product CRUD
└── product_queries.go  # 用户手写的 Product 数据库扩展
```

扩展文件使用同一个 `package model`，可为生成的本地模型类型添加方法，复用它的 DB 与事务。不要覆盖生成文件、修改其生成标记，或重复声明已有模型/方法。例如 `product_queries.go`：

```go
package model

import (
    "context"

    entity "github.com/acme/shop/go/pkg/shop"
)

func (m *Product) FindByName(ctx context.Context, name string) ([]*entity.Product, error) {
    products := make([]*entity.Product, 0)
    err := m.DB.WithContext(ctx).
        Where("name = ?", name).
        Find(&products).Error
    return products, err
}
```

将 import 路径替换为项目实际模块。自定义查询明确需要的过滤、关联加载、分页和事务语义，不默认继承生成 `List` 的行为；查询值使用参数绑定。再生成保留这些独立手写文件，并检查新增生成方法是否与手写方法冲突。

## 初始化与所有权

`db.New(cfg)` **仅返回 `*db.DB`**，未知 driver 或打开失败返回 nil，不能写 `database, err := db.New(cfg)`。目前实现 postgres / mysql / sqlite / dm；clickhouse 虽有常量但未实现。显式设置 driver，不能依赖源码中拼错的默认标注。

简便入口为生成的 `model.GetDB()` / `model.GetProductModel()`：读取根 `db`，用 sync.Once 单例，失败可能 panic，模型首次获取可能迁移。需要可测试注入、多个 DB 或明确启动失败时，优先下面的显式方式，不要在每次请求新建连接。

开发环境可用 `service-go/configs/db.yaml`，其顶层节点仍需写 `db`：

```yaml
db:
  driver: sqlite
  dsn: file:catalog.db
  disableAutoMigrate: false
  maxIdleConnections: 2
  maxOpenConnections: 10
```

生产数据库改为实际 driver 与环境提供的 DSN，并按迁移流程设置 `disableAutoMigrate`。不要把开发示例的自动建表策略直接沿用到生产。

以下代码用于 workflow 中的 `shop/Catalog` 示例。将模块名换成实际值，在 handlers 包修改已有类型/工厂并添加方法，不要和生成占位重复定义：

```go
package handlers

import (
    "context"
    "errors"
    "fmt"
    "time"

    api "github.com/acme/shop/go/pkg/shop"
    pb "github.com/acme/shop/go/pkg/shop/v1"
    "github.com/acme/shop/service-go/pkg/model"
    "github.com/mojo-lang/mojo/go/pkg/mojo/core"
    "github.com/mojo-lang/mojo/go/pkg/mojo/db"
    nserver "github.com/ncraft-io/ncraft/go/pkg/gokit/server"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
    "gorm.io/gorm"
)

type catalogServer struct {
    pb.UnimplementedCatalogServer
    database *db.DB
    products *model.Product
}

func NewService() pb.CatalogServer { return &catalogServer{} }

func (s *catalogServer) Start(_ nserver.Config) error {
    if config.Get("db").Null() {
        return fmt.Errorf("db configuration is required")
    }
    cfg := &db.Config{}
    if err := config.ScanFrom(cfg, "db"); err != nil {
        return fmt.Errorf("load db: %w", err)
    }
    if cfg.Driver == "" || cfg.Dsn == "" {
        return fmt.Errorf("db.driver and db.dsn are required")
    }
    database := db.New(cfg)
    if database == nil || database.DB == nil {
        return fmt.Errorf("open database failed")
    }
    pool, err := database.DB.DB()
    if err != nil { return err }
    if cfg.MaxIdleConnections > 0 {
        pool.SetMaxIdleConns(int(cfg.MaxIdleConnections))
    }
    if cfg.MaxOpenConnections > 0 {
        pool.SetMaxOpenConns(int(cfg.MaxOpenConnections))
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := pool.PingContext(ctx); err != nil {
        _ = pool.Close()
        return fmt.Errorf("ping database: %w", err)
    }
    products, err := model.NewProduct(database)
    if err != nil {
        _ = pool.Close()
        return fmt.Errorf("prepare product model: %w", err)
    }
    s.database, s.products = database, products
    return nil
}

func (s *catalogServer) Shutdown(ctx context.Context) error {
    // 若有后台任务，先用 ctx 停止并等待它们，再关闭连接池。
    if s.database == nil { return nil }
    pool, err := s.database.DB.DB()
    if err != nil { return err }
    return pool.Close()
}

func (s *catalogServer) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*api.Product, error) {
    if in == nil || in.Id == "" {
        return nil, core.NewInvalidArgumentError("id is required")
    }
    result, err := s.products.Get(ctx, in.Id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, core.NewNotFoundError("product %s not found", in.Id)
    }
    if err != nil {
        logs.Errorw("get product", "id", in.Id, "error", err)
        return nil, core.NewInternalError("cannot read product")
    }
    return result, nil
}
```

DB 包装器自身没有统一配置所有连接池参数、慢日志或 tracing；配置存在不代表功能已启用。上面的代码显式设置池大小。超时、连接寿命按实际 driver 与服务需要设置；SQL 调用沿用请求 `WithContext(ctx)`。不要在多个服务之间重复关闭同一个共享池。

## CRUD 的真实语义

| 方法 | 本基线行为 |
| --- | --- |
| `Create(ctx, value)` / `BatchCreate(ctx, values...)` | 返回 `(int64, error)`，数字是 RowsAffected；使用 `OnConflict{UpdateAll:true}`，是 upsert，不是严格只插入 |
| `Get(ctx, id)` | 返回实体指针和 error，找不到为 GORM record-not-found |
| `Update(ctx, value, columns...)` | 不指定 columns 时跳过零值；指定 columns 时由 GORM Select 选择，能够写选中零值；排除主键，不自动插入 |
| `UpdateAll(ctx, value)` | 更新全部非主键字段，包括零值；不自动插入 |
| `UpdateBy` / `UpdateAllBy` | 按指定列筛选且排除该列；非唯一列可能匹配多行，主键未必被自动排除，不当作普通按主键更新的同义词 |
| `Delete` / `BatchDelete` | 返回 RowsAffected 和 error；软删除取决于实体类型/标注 |
| `List(ctx, *query.Query)` | 返回实体数组和 error，不自动提供 total / nextPageToken |
| `BatchUpdate` / `BatchUpdateAll` | 事务执行，任一步错误回滚并返回 0；空批量不访问数据库 |

严格创建需要在业务层使用带唯一约束的纯 GORM Create 并处理冲突，不能把生成 upsert 当成 duplicate error；不要修改生成文件。修改某字段为 `false`、`0`、`""` 时用已确认的 columns、UpdateAll 或手写 map update；PATCH 不应无差别清零未提供字段。校验非 nil 实体及批量元素，一些更新入口在内部 nil 检查前会解引用。

`RowsAffected == 0` 不一定表示不存在（不同 driver 对未变化记录计数不同）；仅在业务确认存在性语义后映射 NotFound。

## 事务、查询和 JSON 字段

多模型事务必须复用同一个 tx，不要在回调中继续用连接池上的原模型。可复制已初始化模型，只替换 DB，避免事务内构造模型又触发迁移：

```go
err := s.database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    products := *s.products
    products.DB = s.database.Tx(tx)
    _, err := products.UpdateAll(ctx, changedProduct)
    return err // 后续操作也必须通过同一 tx；错误向外返回触发回滚。
})
```

`query.Query` 位于 `mojo/db/query`，支持 FieldQueries、Filter、Order、FieldMask、PageSize / PageToken 等。模型 List 会 Normalize 再 Apply；此对象可能被修改，勿跨请求共享。这里 PageToken 是分页索引的十进制字符串；不要混用其他 paginator 的 token。排序、投影、Groups、UpdateBy 列名等由业务白名单映射，值用占位参数，不能假定每条 query 分支都验证了用户 SQL 片段。

JSON 列优先在 `.mojo` 标注 `@db.json` 再生成；手写 JSON 数据库类型可参考 `db.JSONValuer`、`JSONScanner`、`JSONDbDataType`。序列化仍须 jsoniter；JSON null、SQL NULL 和零值的语义分别测试。基础类型、关联模型、数据库 driver 的差异用临时数据库验证。
