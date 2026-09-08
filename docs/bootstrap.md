# 更新 Mojo 核心组件的 Go 代码

Mojo 的 IDL 位于 `packages/<组件>/mojo/...`，语法由
`antlr/mojo/MojoLexer.g4` 和 `MojoParser.g4` 定义。命令行入口是
`go/cmd/mojo`，编译器位于 `go/pkg/compiler`，生成代码与手工编写的
运行时共同位于 `go/pkg/mojo`。

所有内置组件共享 `go/go.mod` 中的模块
`github.com/mojo-lang/mojo/go`。`package.mojo` 的 repository 仍表示
IDL 的位置，不再对应一个独立的 Go 模块。

## 完整自举

需要 Go（版本满足 `go/go.mod`）、`protoc`、`protoc-gen-go` 和
`protoc-gen-go-grpc`，并将这些命令放入 PATH。

在仓库的 `go` 目录运行：

```sh
go run ./cmd/mojo bootstrap
go build ./cmd/mojo ./pkg/mojo/...
```

`bootstrap` 也可通过已编译的 mojo 命令运行，并接受仓库根目录或其
子目录作为参数：

```sh
mojo bootstrap /path/to/mojo
```

命令依次更新 core、document、lang、db、geom、http、openapi、rpc。
每个组件都经过本地 Mojo 源码解析、语义处理、AST 到 Protobuf 转换、
protoc 生成 Go、Go 扩展代码生成。最后刷新
`go/pkg/compiler/mojo/mpm/mojo` 中内嵌的 AST 和 Protobuf 快照。
重新构建或安装 CLI 后，新快照才会进入新的可执行文件。

生成位置为：

- Protobuf：`packages/<组件>/protobuf`
- Go：`go/pkg/mojo/<组件>`
- 内嵌快照：`go/pkg/compiler/mojo/mpm/mojo`

自举不会创建 `packages/<组件>/go`，也不会改写共享的 `go.mod` 或执行
该模块的 `go mod tidy`。现有的手写 Go 文件沿用生成器的保护规则：
只有被识别为纯生成代码的文件会被替换。需要手动保留的实现不应带有
`Code generated ... DO NOT EDIT` 文件头。

## 单个组件

在仓库的 `go` 目录运行：

```sh
go run ./cmd/mojo build -t go ../packages/core
```

`go`（或 `golang`）目标仅生成 Protobuf 和 Go；`protobuf` 目标仅生成
Protobuf；`java` 目标生成 Protobuf、Java 消息类和 gRPC stub；
`api` 目标继续包含 OpenAPI、文档、Protobuf 和 Go。
C++ 不受支持，NCraft 服务生成仅支持 gokit。

核心组件构建时优先读取本仓库中的依赖源码；显式声明的依赖 path
优先于自动定位。普通仓库外的项目继续使用 CLI 内嵌的核心组件快照。
依赖组件也有变动时，运行完整 `bootstrap`，以同步依赖的 Go 代码和快照。

`build -t go -o /path/to/output ...` 可指定 Go 输出目录。
Protobuf 中间文件仍写入组件的 `protobuf` 目录，保证 protoc 读取本次生成的文件。

## Java 核心组件

Java 主源码统一位于 `java/src/main/java`，测试位于 `java/src/test`。
原 `packages/<组件>/java` 中的手写辅助类及测试已迁入该目录。

```sh
# 在 go 目录执行
go run ./cmd/mojo bootstrap -t java
go run ./cmd/mojo build -t java ../packages/core
# 同时更新 Go 和 Java
go run ./cmd/mojo bootstrap -t go,java
```

`build -t java -o <目录>` 将 Java 源码写入该目录下的 `src/main/java`。
每个 Mojo 包在 `java/.mojo-generated` 中记录自己的生成文件，单包重生成
仅删除该包已不再使用的生成文件，并保留其他包及手写文件。

ANTLR Java 语法代码也生成到 `java/src/main/java`；可以从任意工作目录执行
`antlr/mojo/generate-java.sh`。脚本在临时副本中转换语言专用谓词，失败时
也不会改写原始 `.g4` 文件。旧 `java/mojo` 解析器原型保留在原目录，
不属于新的 Maven 类库构建。详见 `java/README.md`。
