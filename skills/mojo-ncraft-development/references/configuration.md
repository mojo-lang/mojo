# 新配置与基础配置复用

## 可直接复制的组件模板

本 skill 的 `configs/` 只维护可复用的通用组件 YAML 模板，项目专属业务配置留在对应工程。模板可直接作为目标工程的配置文件，不需要额外渲染 Go template；`${NAME}` 由 NCraft 配置系统在运行时读取环境变量。文件根节点采用当前消费者实际读取的 `db`、`redis`、`server`、`logs` 等名称。

| 模板 | 复制后的定制项与行为 |
| --- | --- |
| [db.yaml](../configs/db.yaml) | 默认 postgres，设置 `DB_DSN`；按数据库调整 driver。默认关闭 AutoMigrate，先执行项目迁移；本地开发可明确开启。连接池参数由实际初始化代码应用 |
| [redis.yaml](../configs/redis.yaml) | 修改 connections / dbNumber，按需设置 `REDIS_PASSWORD`，注册 goredis 实现；模板的 timeout 使用整数纳秒，1 秒为 1000000000 |
| [server.yaml](../configs/server.yaml) | 调整 HTTP / gRPC / debug 地址及端口；使用生成框架的监听器，debug 默认仅绑定本机 |
| [logs.yaml](../configs/logs.yaml) | 默认 info、JSON 编码、控制台输出；设置 `SERVICE_PACKAGE`、`SERVICE_NAME`、`SERVICE_VERSION` 或直接填写项目标识。文件输出使用 `file` 或 `console,file`，调整路径及轮转参数 |
| [sd.yaml](../configs/sd.yaml) | 默认 direct、gRPC、无实例、关闭重试；按生成 FullServiceName 添加下游地址。`retry.max` 是总尝试次数，timeout 为毫秒，不能沿用 `count` |
| [metrics.yaml](../configs/metrics.yaml) | 按需启用并填写 department / project；当前 Config 没有 group。关闭业务指标不等于关闭 debug listener 的 `/metrics` |
| [tracing.yaml](../configs/tracing.yaml) | 默认关闭；启用时填写 Jaeger UDP agent 地址，不能填成 OTLP collector URL |
| [jwt.yaml](../configs/jwt.yaml) | 默认未启用；要求鉴权时必须设置非空 `JWT_KEY`、开启 enable、装配 JWT middleware 并验证无 token 请求被拒绝；有效期为 Mojo Duration 字符串 |

普通工程从这些模板选择组件，保留字段结构和原有 Config 类型。不要把项目专属业务项、其他项目的服务名、固定凭据和部署路径当作公共默认值。新增项目可直接复制；已有项目逐项合并同名配置，保留其已确认的端口、连接、权限和部署约定。

从目标项目根目录使用复制工具，例如选择服务、日志、数据库和 Redis：

```sh
skill_dir="${CODEX_HOME:-$HOME/.codex}/skills/mojo-ncraft-development"
python3 "$skill_dir/scripts/copy_configs.py" --dest service-go/configs --components server logs db redis
```

使用仓库副本或其他 agent 时，将 `skill_dir` 改为实际包含 `SKILL.md` 的目录。已在 `service-go/` 内则使用 `--dest configs`。省略选择参数默认复制 server、logs、metrics、tracing、sd；`--all` 复制八份通用模板。也可手工复制所需 YAML。

工具只复制模板，不展开环境变量或改动 Go 依赖，也不会覆盖同名文件。生成器若已写入同名配置，工具输出 `SKIP existing`：比较模板后按需合并，不要将“跳过”当成已更新。模板的原始文件继续留在 skill，工程使用自己的可定制副本。

复制后按顺序完成：

1. 根据实际部署填写服务元数据、监听端口、DB / Redis 地址，以及需要的发现实例；凭据在进程启动前通过环境提供。NCraft 不会自动加载 `.env`，也不支持 `${NAME:-default}`；不要把这些表达式直接加入模板。
2. 按实际需求启用可选组件。配置文件存在不代表对应业务已接入；DB / Redis 仍需服务初始化和资源关闭，JWT 仍需 middleware，metrics / tracing / retry 开关按业务确定。不要给没有该字段的组件臆造统一 `enable` 配置。
3. 用目标项目锁定版本的 Config 和 NCraft Scan 验证类型及关键值。数字、布尔值保留 YAML 原生类型；不要把 duration 单位混同，也不要把扫描成功当成必需项非空或已经连通。
4. 从 `service-go/` 启动或设置正确的 `CONFIG_PATH`。基础模板使用无环境后缀的文件名；若使用 `DEPLOY_ENVIRONMENT`，先按下文确认文件选择策略，不能假定 base 文件仍会被自动加载。

direct 模式没有注册中心 URL。当前通用 `sd.New` 在 URL 为空时返回 nil，所以生成 host 会跳过注册；当前生成客户端专门使用 direct 构造器读取实例表。旧客户端或手写消费者应先核对其实现，不要为了绕过检查保留一个无关的 etcd 地址。

## 路径、结构与默认值

配置管理包为 `github.com/ncraft-io/ncraft/go/pkg/ncraft/config`。配置文件名不创建命名空间；`configs/catalog.yaml` 必须自行包含 `catalog:` 根节点。Go 字段使用准确的 JSON tag，YAML 经配置系统转换后由 jsoniter Scan。

```yaml
catalog:
  enabled: false
  pageSize: 50
  readTimeout: 5s
  storage:
    vendor: s3
    endpoint: https://objects.example.com
    bucketName: catalog
    accessKey: ${CATALOG_ACCESS_KEY}
    secretKey: ${CATALOG_SECRET_KEY}
    createBucket: false
```

服务本地配置放 `service-go/pkg/catalog/config.go`（或既有配置包），不必为了运行配置生成公共 API：

```go
package catalog

import (
    "fmt"

    "github.com/mojo-lang/mojo/go/pkg/mojo/core"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
    "github.com/ncraft-io/ncraft/go/pkg/ncraft/storage"
)

type Config struct {
    Enabled bool `json:"enabled" default:"true"`
    PageSize int `json:"pageSize" default:"50"`
    ReadTimeout *core.Duration `json:"readTimeout"`
    Storage storage.Config `json:"storage"`
}

func LoadConfig() (*Config, error) {
    value := config.NcraftGet("catalog")
    if value.Null() { return nil, fmt.Errorf("catalog configuration is required") }
    timeout, err := core.ParseDuration("5s")
    if err != nil { return nil, err }
    cfg := &Config{ReadTimeout: timeout}
    if err := value.Scan(cfg); err != nil { return nil, err }
    if cfg.PageSize < 1 || cfg.PageSize > 500 {
        return nil, fmt.Errorf("catalog.pageSize must be in [1, 500]")
    }
    if cfg.ReadTimeout == nil || cfg.ReadTimeout.Seconds <= 0 {
        return nil, fmt.Errorf("catalog.readTimeout must be at least 1s")
    }
    return cfg, nil
}
```

在 Start 中加载一次并注入业务；需要对象存储时另检查 endpoint、bucket、凭据等所需项并调用构造函数。`default` tag 是 NCraft 扫描子节点时的能力，直接 `jsoniter.Unmarshal` 不会应用它。默认 tag 先赋值再 Scan，所以 YAML 显式 false 能覆盖 true；不要再用 `if !Enabled { Enabled = true }` 抹掉用户输入。

`ScanFrom(cfg, "catalog", "legacyCatalog")` 是选择首个非 Null 节点的 fallback，不是合并。缺失节点 Scan 可成功返回零值，因此必需配置要显式检查 Null 和字段。`NcraftGet("catalog")` 优先 `ncraft.catalog` 再根 `catalog`；传入显式路径参数后不走这个 fallback。实例 `conf.Get` 接收路径 segments；模块级 `config.Get` 才额外将点分路径展开。

`*core.Duration` 能读取 Mojo 的 `5s` 字符串。不要假定任意 `time.Duration` 字段 JSON 都能读 `5s`：普通 jsoniter 通常按整数纳秒，NCraft 部分客户端配置又明确按毫秒。尊重原类型约定，不统一猜测单位。pointer、slice、map 的复杂默认值可在加载函数设置，不能假定 default loader 递归覆盖所有类型。

## 复用已有 Config

| 能力 | 复用类型 |
| --- | --- |
| 数据库 | `mojo/db.Config` |
| Redis | `ncraft/redis.Config` |
| 对象存储 / S3 | `ncraft/storage.Config` |
| JWT | `ncraft/auth/jwt.Config` |
| 日志 | `ncraft/logs.Config` |
| 服务监听 | `gokit/server.Config` |
| 发现 / 客户端 / 重试 | `gokit/sd.Config`、`gokit/client.Config`、`gokit/retry.Config` |
| metrics / tracing | `gokit/metrics.Config`、`gokit/tracing.Config` |
| 消息 | `ncraft/messaging.Config` |

只要语义相同就组合已有类型，不复制 endpoint/secretKey 等字段。兼容旧名称可用 `type S3ProviderConfig = storage.Config`，有额外字段时用包含 `Storage storage.Config` 的 wrapper；不要给别名添加不能合法定义的方法。

复用可以是**类型复用**、**路径 fallback**、**配置源覆盖**或**客户端字段合并**，需要明确是哪一种。生成 DB 单例当前只读根 `db`；不要以为所有消费者都自动读取 `ncraft.db`。同理生成 server 某些调用会传显式 `"metrics"` / `"sd"`，核对调用方后再安排命名空间。

## 默认文件发现的实际行为

基线 `config_helper.go`：

1. 导入 config 时初始化默认配置；按进程工作目录找 `conf/`、`configs/`，`CONFIG_PATH` 非空时优先它。命中第一个含可用配置文件的目录后停止，不合并所有目录。
2. 支持默认发现 `.yaml` / `.yml` / `.json` / `.toml`；文件内容按来源合并，随后加入 env source。
3. `DEPLOY_ENVIRONMENT=dev` 时选择文件名主干以 dev 结尾的文件；**不自动先加载所有基础文件再加 dev 文件**。未设置时反而可能把各环境后缀文件都读入。
4. 当前 `getEnv` 对 `NCRAFT_CONFIG_PATH` / `NCRAFT_DEPLOY_ENVIRONMENT` 的前缀分支有返回空串的问题；基线不要宣称这些前缀变量可用，优先无前缀变量并验证实际文件路径。目录递归也不应未经验证当作可靠功能。

这些是版本相关行为，升级后重新读源码。不要把它们变成其他配置系统的一般规范。

## 明确的基础配置叠加

需要可控的 base → environment → env 覆盖时，使用显式配置实例并检查加载错误：

```go
conf, err := config.NewConfig()
if err != nil { return err }
// 在启动阶段创建，由拥有它的服务在 Shutdown 关闭；测试中可 defer conf.Close()。
err = conf.Load(
    file.NewSource(file.WithPath("configs/base.yaml")),
    file.NewSource(file.WithPath("configs/catalog.dev.yaml")),
    env.NewSource(env.WithStrippedPrefix("SHOP")),
)
if err != nil { _ = conf.Close(); return err }
var cfg Config
if err := conf.Get("catalog").Scan(&cfg); err != nil {
    _ = conf.Close()
    return err
}
```

`file` / `env` import 来自 `ncraft/config/source/file`、`source/env`。此实例用于注入自身配置；不会改变生成 host 读取的全局 DefaultConfig。若要让生成 host 使用这些来源，在业务启动之前明确加载到全局 config，并在依赖已读取配置前完成；不要在并发请求中替换 DefaultConfig，也不要关闭仍被其他服务使用的全局实例。

一般 source merge 的边界以本版本 reader / mergo 为准；测试 false、0、空列表、null，不要套用客户端专门 merge 的完整语义。env source 将名称小写并以 `_` 分层：`SHOP_CATALOG_ENABLED=false` 在 strip SHOP 后是 `catalog.enabled`；`PAGE_SIZE` 会成为两层 `page.size`，不是 `pageSize`。复杂 camelCase 键宜在 YAML 中用 `${ENV_VAR}`。

`${NAME}` 是字符串替换，没有 `${NAME:-default}` 语法，未设置会变成空串。包含引号、反斜杠的值需验证转义；关键配置应在启动时做环境变量/字段校验。示例中仅放变量名，不写真实凭据。

## 客户端基础配置复用

`gokit/client.LoadConfig("catalog")` 的专用合并顺序为 sd → client 通用 → client.catalog。每一层经 NcraftGet 优先取 ncraft 节点。嵌套对象按字段合并，数组替换，显式 false/0/空字符串/null 覆盖，空对象保留继承字段，null 清空子配置。此语义来自专用客户端 loader，不代表所有 ScanFrom 都能深度合并。

生成 `NewClientFromConfig` 默认用 FullServiceName 选择客户端配置；`NewClientFromServiceConfig(ctx, "catalog", ...)` 用别名。配置选择键与注册中心 `serviceName` 是两回事；HTTP 超时和 gRPC dialTimeout 是毫秒。详见 [运行时类库](runtime.md)。

## 动态更新

配置源变化不等于已构造的 DB、Redis 或业务 struct 自动变化。需要热更新时显式 Watch → 解码新快照 → 校验 → 原子替换，并停止 watcher / goroutine。Start 失败和 Shutdown 都处理资源归属；不要仅启用文件 watch 就声称热更新完成。
