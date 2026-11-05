---
name: mojo-ncraft-development
description: "创建、实现、维护基于 Mojo IDL 与 NCraft Go kit 的 Go 微服务；用于脚手架、GORM 数据访问、接口再生成、服务扩展、中间件、jsoniter 自定义序列化和配置复用。适用于含 package.mojo、service-go 或 github.com/ncraft-io/ncraft/go 的服务项目，不用于 Modular Mojo / Python 互操作。"
---

# Mojo / NCraft 微服务开发

把用户需求落实为可再生成、可测试的服务实现。本文中的 Mojo 是 `mojo-lang` 的接口定义语言，NCraft 是基于 Go kit 的运行时。

## 先确定项目与版本

1. 找到目标项目的 `AGENTS.md`、`package.mojo`、`mojo/` 和各个 `go.mod`，读取已有构建脚本、配置与相邻业务实现。
2. 分别确认 `go/`、`service-go/`、`client-go/` 的依赖版本及 `replace` / `go.work`。在对应模块目录用 `go list -m -json github.com/mojo-lang/mojo/go` 和 `go list -m -json github.com/ncraft-io/ncraft/go` 定位实际依赖；模块未依赖某包时无需强行加入。
3. 检查 Mojo CLI 来源及 Go / protoc 插件版本。生成器与运行时可能不一致；新增扩展点前检查目标项目生成的装配代码及锁定依赖，不能把参考中的新 API 当作旧版本必然具备的能力。

## 按任务读取

| 任务 | 参考 |
| --- | --- |
| 创建项目、定义类型/接口、更新 Mojo 并重新生成 | [项目与接口迭代](references/workflow.md) |
| 实现服务、初始化/退出、定制 HTTP、扩展接口 | [服务与扩展点](references/services.md) |
| Entity 组织、数据库标注、生成与扩展 model、GORM、事务及分页 | [数据库](references/database.md) |
| 工程 JSON 规范、自定义结构体编解码 | [JSON 序列化](references/json.md) |
| 复制通用组件配置模板、新配置、默认值、配置源及基础配置复用 | [配置](references/configuration.md) |
| 鉴权、限流、日志、发现客户端、Redis、对象存储、消息、KV | [中间件与类库](references/runtime.md) |
| 核对版本差异、查找规则对应实现 | [源码索引](references/sources.md) |

仅加载与当前任务有关的参考。新建完整服务通常需要 workflow、services、database、json、configuration。

## 必须保持的约定

- **Mojo 定义是协议源头。** 修改公开类型、方法、路由时先改 `.mojo`，再生成依赖产物；已发布字段编号不能重用。不要手改 `.pb.go`、生成的 transport / endpoints / model 来代替接口迭代。
- **文本注释遵循 Mojo 约定。** `type` 和 `interface` 的文本注释放在声明之前，使用 `///`；`type` 内部字段和 `enum` 成员可在同一行末尾使用 `//<` 尾追注释。示例见[项目与接口迭代](references/workflow.md)。
- **Entity 优先放在非版本目录。** 服务相关 Entity（`type` 结构体）放在业务包的非 `v1` / `v2` 目录，版本目录一般放服务 `interface`。Entity 通常以 ID 为首字段，Mojo 源码沿用 `id` 命名。生成时显式包含 `service` 目标，将基于 GORM 的 SQL 操作代码生成到 `service-go/pkg/model/`；用户可在同一 `model` 包的独立手写文件中扩展数据库操作。
- **数据库映射通过已支持的标注表达。** 按存储语义选择 `@db.json`、`@db.ignore`、`@db.index`、主键和关联标注，核对生成字段类型及 GORM tag；具体写法和支持边界见[数据库](references/database.md)。
- **业务与扩展留在可保留的文件中。** 使用 `handlers` 包、自建业务包及独立手写扩展文件。保留 `NewService() pb.XServer` 工厂签名；有状态服务返回指针。生命周期方法是否生效以实际生成 host 为准。
- **HTTP 服务优先沿用微服务框架。** 使用生成的 HTTP transport、路由、endpoint 和生命周期，通过框架已有扩展点实现需求。不要为了方便另外手写一套 HTTP 服务，再通过额外监听端口、替换主 router、兜底转发或修改生成的启动代码强行接入。特殊 HTTP 行为优先使用编解码接口、绑定 writer 或同一 router 上的必要路由扩展，具体选择见[服务与扩展点](references/services.md)。
- **凡工程使用 Mojo 相关包，业务 JSON 序列化统一使用 `github.com/json-iterator/go`。** 用户所称 jsonitor 指此 jsoniter。HTTP、缓存、消息、数据库 JSON、自定义 MarshalJSON 内部均遵守；不能用 `encoding/json`、`protojson` 等另起一套 Mojo JSON 编解码。标准库 `json.Marshaler` / `RawMessage` 等类型声明不等同于调用标准库序列化；不要为删除 import 改坏生成文件或第三方代码。
- **复用 NCraft / Mojo 类型。** DB 使用 `mojo/db.DB` 与 `db.Config`，配置使用 NCraft config，其他能力优先使用已有类库。真实连接、订阅、watcher、后台任务由拥有它们的服务管理生命周期。
- **通用组件配置优先使用随 skill 提供的模板。** `configs/` 内的 db、redis、server、logs、sd、metrics、tracing、jwt YAML 可按需复制到工程 `service-go/configs/`；也可运行 [copy_configs.py](scripts/copy_configs.py)，不会覆盖同名配置。复制后填写项目标识、连接参数与所需开关，复用对应 Config 类型。此目录只维护通用组件模板，项目专属业务配置留在对应工程；具体步骤见[配置](references/configuration.md)。
- **再生成不会替你更新所有手写逻辑。** 保留的 middleware、新接口的权限、旧方法签名、依赖版本和配置仍需逐项核对。生成代码成功不等于服务已经实现或能够编译。

## 完成标准

按改动运行相关模块的测试与构建；协议变更同时检查 API、服务端和客户端。新增 CRUD 要覆盖零值更新、未找到和事务行为；JSON 变更要检查实际输出及往返；配置变更要检查缺失、覆盖与默认值；扩展点要验证宿主确实调用。

核对生成差异，确认手写实现保留、没有重复方法、新方法已经接入需要的 middleware。交付说明改动的 Mojo 文件、手写实现、配置项、验证结果和真实存在的版本限制。示例仓库地址和服务名须替换为目标项目的实际值，不能复制本机绝对路径或示例凭据。
