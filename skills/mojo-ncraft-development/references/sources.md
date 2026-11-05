# 源码索引与版本边界

本 skill 根据 2026-10-15 可用源码编写：Mojo `6f649d1`、NCraft `6b95135`，并交叉阅读 Armory / Tilestream 的真实服务。以下是**仓库内相对路径**，不是安装 skill 后文件所在目录；不要求开发者具有作者本机的目录。

通用组件配置模板于 2026-11-05 按上述源码版本重新核对；复制到其他版本的工程时，仍需验证其 Config 字段和初始化行为。

先根据目标模块 `go.mod` / `go.work` / `go list -m -json` 找到实际源代码。工程可能同时存在 CLI 新版本、API 旧版本、service/client 不同 NCraft 版本；新特性使用前核对编译及生成装配。

## Mojo 仓库 `github.com/mojo-lang/mojo`

| 路径 | 核对内容 |
| --- | --- |
| `docs/ncraft.md`、`docs/syntax.md`、`docs/entity-models.md` | 项目流程、IDL 语法、实体概念；具体行为以实现为准 |
| `go/cmd/mojo/cmd/create.go`、`build.go` | CLI 参数 |
| `go/pkg/compiler/cmd/commander/builder.go` | targets 与默认生成行为 |
| `go/pkg/compiler/ncraft/gokit/generator/options.go`、`handler_sources.go` | 生成路径及 handlers 包保留策略 |
| `go/pkg/compiler/ncraft/gokit/generator/handlers/` | 方法补全、middleware/hooks 保留、handler 模板 |
| `go/pkg/compiler/ncraft/gokit/generator/templates/service-go/internal/NAME-server/run.go.tmpl` | 生命周期、路由、实例共享、装配 |
| `go/pkg/compiler/ncraft/gokit/generator/templates/service-go/internal/server/run.go.tmpl` | 统一 host |
| `go/pkg/compiler/ncraft/gokit/generator/httptransport/templates/server.go` | 自定义 request/response、绑定 writer、HTTP 错误编码 |
| `go/pkg/compiler/ncraft/gokit/generator/model/templates/entity_model.go`、`db.go` | CRUD 真实语义、模型构造、单例和迁移 |
| `go/pkg/compiler/ncraft/gokit/generator/templates/client-go/` | 客户端构造、SD、options、配置 |
| `go/pkg/compiler/util/generated_file.go` | 文件覆盖条件 |
| `go/pkg/mojo/db/db.go`、`json.go`、`query/` | GORM 包装、JSON DB、查询与分页 |
| `mojo/db/config.mojo`、`go/pkg/mojo/db/config.pb.go` | DB 配置定义及字段 |
| `mojo/db/attributes.mojo`、`go/pkg/mojo/db/*_attribute.go` | 数据库标注名称与定义；有定义不等于所有后端都实现 |
| `go/pkg/compiler/go/generator/generator/protoc_gen_go.go`、`go/pkg/compiler/go/generator/injection/db_json_injector.go` | GORM tag 映射、JSON 字段包装与注入 |
| `go/pkg/mojo/core/jsoniter.go`、`timestamp.json.go`、`uuid.json.go` | 注册包装器、scalar / 指针 codec |
| `go/pkg/logs/config.go`、`sugar.go` | 日志模板字段、输出目标及文件轮转单位 |

## NCraft 仓库 `github.com/ncraft-io/ncraft`

| 路径 | 核对内容 |
| --- | --- |
| `go/pkg/gokit/server/lifecycle.go`、`README.md` | Starter / Shutdowner / ErrorSource 的执行约定 |
| `go/pkg/gokit/transport/http/README.md`、`request_context.go` | ResponseWriter 绑定及生命周期 |
| `go/pkg/gokit/middleware/` | JWT、限流、logging、metrics |
| `go/pkg/gokit/client/README.md`、`config.go` | 客户端专用合并、选择键、单位 |
| `go/pkg/gokit/sd/`、`tracing/`、`metrics/`、`retry/` | discovery、观测、重试 |
| `go/pkg/gokit/server/config.go`、`go/pkg/ncraft/logs/config.go`、`logs.go` | server / logs 模板的字段及读取路径 |
| `go/pkg/ncraft/redis/config.go`、`go/pkg/ncraft/redis/goredis/redis.go`、`go/pkg/ncraft/auth/jwt/config.go` | Redis 超时单位、实现注册与 JWT 配置 |
| `go/pkg/ncraft/config/config.go`、`utils.go`、`config_helper.go` | ScanFrom、NcraftGet、文件发现及环境变量边界 |
| `go/pkg/ncraft/config/tag/tag.go`、`reader/json/values.go` | default tag、Null 和 jsoniter Scan |
| `go/pkg/ncraft/config/reader/json/json.go`、`reader/preprocessor.go` | source merge 与变量替换 |
| `go/pkg/ncraft/config/source/env/`、`source/file/` | 来源与键名映射 |
| `go/pkg/ncraft/storage/`、`redis/`、`messaging/`、`kvstore/` | 注册实现、构造返回类型及资源关闭 |

## 可参考的业务实现

- Armory `service-go/pkg/file-service/handlers/handlers.go`：固定 protobuf 响应绑定 HTTP writer；`server.go` 接入 Start。
- Armory `service-go/pkg/file-service/util/config.go`：业务配置组合与 `type S3ProviderConfig = storage.Config`。
- Armory `service-go/pkg/unitable-service/handlers/config.go`：运行配置 wrapper 与公共 API 实体分离。
- Tilestream `service-go/`：独立业务服务与构建脚本，可借鉴其结构；存在历史 patch 脚本不代表应将修改生成代码作为新项目惯例。

示例用于理解结构与契约，不复制其中的真实连接串、业务数据或历史缺陷。若参考文档与实现冲突，记录目标版本的实际行为并对关键路径验证。
