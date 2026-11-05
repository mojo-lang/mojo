# JSON Iterator：统一序列化与自定义结构体

工程使用 Mojo 相关包时，业务 JSON 编解码统一使用 **jsoniter**，模块是 `github.com/json-iterator/go`；不是一个名为 jsonitor 的依赖。Mojo 的 Timestamp、Duration、Uuid、union、boxed 类型等在初始化阶段注册 codec，标准库反射编码不能保证这些类型的 wire format。

## 普通结构体

沿用项目已选择的 jsoniter API。生成 HTTP 和数据库 JSON 路径常用 `ConfigFastest`，配置 Scan 使用 jsoniter 默认 API。普通带 json tag 的结构体无需专门注册 codec：

```go
package wire

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

var jsonCodec = jsoniter.ConfigFastest

type Event struct {
    Name string `json:"name"`
    CreatedAt *core.Timestamp `json:"createdAt,omitempty"`
    TTL *core.Duration `json:"ttl,omitempty"`
}

func EncodeEvent(event *Event) ([]byte, error) { return jsonCodec.Marshal(event) }
func DecodeEvent(data []byte, event *Event) error { return jsonCodec.Unmarshal(data, event) }
```

导入 core 已触发其注册，无需反复注册 Timestamp 等生成类型。所有序列化入口（HTTP writer、缓存、消息、DB JSON、自定义 Marshaler 内部）保持一致。不要经过 `encoding/json.Marshal` / `Unmarshal` 中转，不要把 protobuf 的 protojson 输出直接当作 Mojo HTTP JSON。

不同 jsoniter config 在浮点精度、HTML 转义、map 排序等方面不完全相同；协议需要精确数字或稳定输出时用现有约定及测试确定 config，不能仅以“都叫 JSON”假定相同。改变 config 是可观察协议变更。

## 自定义对象形状：显式 wire DTO

例如内部字段叫 Name / CreatedAt，对外需要 eventName / at；用独立 wire 类型实现 MarshalJSON / UnmarshalJSON，内部仍调用 jsoniter。下面代码与上一段同属 `wire` 包：

```go
type EventEnvelope struct {
    Name string
    CreatedAt *core.Timestamp
}

type eventEnvelopeJSON struct {
    EventName string `json:"eventName"`
    At *core.Timestamp `json:"at,omitempty"`
}

func (e EventEnvelope) MarshalJSON() ([]byte, error) {
    return jsonCodec.Marshal(eventEnvelopeJSON{EventName: e.Name, At: e.CreatedAt})
}

func (e *EventEnvelope) UnmarshalJSON(data []byte) error {
    var v *eventEnvelopeJSON
    if err := jsonCodec.Unmarshal(data, &v); err != nil { return err }
    if v == nil {
        *e = EventEnvelope{}
        return nil
    }
    *e = EventEnvelope{Name: v.EventName, CreatedAt: v.At}
    return nil
}
```

这里选择 null 清零，并在成功解码后整体替换对象。若需要 PATCH 的“缺失保持原值”，需另设计 presence / pointer 字段。不能在 `MarshalJSON` 中再次 Marshal 同一个 EventEnvelope，否则无限递归。不要同时给同一类型写方法和注册 codec，除非确实了解优先级并测试过。

## 自定义类型的 jsoniter codec

当需要将自定义结构体表达为 scalar / union，或确实需要注册表能力时，再实现 ValEncoder / ValDecoder。下面把 `OrderCode{Value:"p-1"}` 编码为 JSON 字符串 `"p-1"`。放在独立手写文件，不修改生成 `.json.go`：

```go
package wire

import (
    "reflect"
    "unsafe"

    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

type OrderCode struct { Value string }
type orderCodeCodec struct{}

func init() {
    name := reflect.TypeOf(OrderCode{}).String() // wire.OrderCode，非 import path
    codec := &orderCodeCodec{}
    core.RegisterJSONTypeEncoder(name, codec)
    core.RegisterJSONTypeDecoder(name, codec)
}

func (*orderCodeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString((*OrderCode)(ptr).Value)
}

func (*orderCodeCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return (*OrderCode)(ptr).Value == ""
}

func (*orderCodeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := (*OrderCode)(ptr)
    if iter.WhatIsNext() == jsoniter.NilValue {
        iter.ReadNil()
        *value = OrderCode{}
        return
    }
    if iter.WhatIsNext() != jsoniter.StringValue {
        iter.ReportError("OrderCode", "expected string or null")
        return
    }
    text := iter.ReadString()
    if iter.Error != nil { return }
    value.Value = text
}
```

`core.RegisterJSONTypeEncoder/Decoder` 是 Mojo 包装器；encoder 同时记录到 Mojo 的注册表。注册发生在 init、首次使用前，不要按请求重复或并发修改全局表。名称来自类型的包名加类型名，不是模块完整 import path；不同包若同名需检查冲突。

Type codec 的 ptr 指向注册的值类型，jsoniter 负责外层指针间接层。不要把它直接转换为 `**OrderCode`。字段级定制使用 `core.RegisterJSONFieldEncoder/Decoder(typeName, goFieldName, codec)`，field 是 Go 字段名（如 `Code`）而不是 JSON tag；字段指针需要与其实际内存布局匹配，参考生成的 `IsFieldPointer` codec，不能不加区别套用上述实现。

若需要绕开自己注册的类型 codec，定义新的非 alias 类型（`type BareEvent Event`）或显式 DTO；`type BareEvent = Event` 仍是同一类型，会继续递归。不要为 Timestamp 等已有 Mojo 类型重新注册另一套全局编码规则。

## 验证实际协议

除了能编译，至少验证：

- 顶层值、指针、嵌套字段、slice/map 的具体 JSON 形状及往返；Mojo Timestamp 应是其 Format 的字符串，而非内部 seconds 对象。
- nil 指针、null、空集合、零值和 `omitempty`；指针指向零值与 nil 可能不是同一行为。
- 非法 token / 类型必须返回 error，不吞掉错误、不留下错误被当作成功的对象。
- 用生产路径的 config 测试；若配置扫描也使用此类型，补默认 jsoniter API 的测试。

可先搜索自有代码中的 `encoding/json`、`protojson` 及 Marshal / Unmarshal / Encoder / Decoder 调用。检查实际调用，不要仅凭 import 判违规：生成 transport 使用 `json.Marshaler` 接口声明是合法的，第三方及生成内部实现也不应被盲目替换。
