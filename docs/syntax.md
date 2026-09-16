# Mojo 简明语法：type 与 interface

本文依据项目的 [ANTLR 解析规则](../antlr/mojo/MojoParser.g4) 和
[词法规则](../antlr/mojo/MojoLexer.g4)，介绍数据结构与服务接口的常用写法。
关键字使用小写 `type`、`interface`；本文中的 Interface 指接口这一概念。
项目创建、目录组织和代码生成见 [NCraft 微服务开发指南](ncraft.md)。
字面量、运算符和调用等写法见 [Mojo 表达式说明](expression.md)。

## 1. 基本书写规则

Mojo 源文件使用 `.mojo` 扩展名。类型和接口名称使用大写字母开头的名称，
例如 `Product`、`Catalog`；字段、参数和方法通常使用小写下划线命名，
例如 `display_name`、`get_product`。

声明和结构体字段通常以换行分隔，不需要分号；同一行也可以用 `;` 分隔。
方法参数通常以逗号分隔，也支持以换行分隔。字段和参数建议始终写出
`名称: 类型` 中的冒号，虽然语法允许省略该冒号。

## 2. 使用 type 定义结构体

最常用的形式是 `type 类型名 { 字段列表 }`：

```mojo
/// 商品信息
type Product {
    id: String @1 //< 商品 ID
    name: String @2 //< 商品名称
    price: Float64 @3 //< 商品价格
    tags: [String] @4 //< 商品标签
    metadata: {String: String} @5 //< 扩展属性
}
```

每个字段由字段名、类型和可选标注组成。例如 `name: String @2` 中，
`name` 是字段名，`String` 是类型，`@2` 是数字标注。生成 Protobuf 时，
数字标注用作字段编号；面向服务的数据定义应显式指定稳定、不重复的编号。
同一结构内编号唯一，已发布字段的编号不应随排列顺序改变，也不要复用
已删除字段的编号。这些编号约束属于协议生成规则。

常用字段和参数类型如下：

| 写法 | 含义或用途 |
| --- | --- |
| `String`、`Bool` | 字符串、布尔值 |
| `Int32`、`Int64`、`UInt32`、`UInt64` | 常用整数类型 |
| `Float32`、`Float64` | 浮点数 |
| `Bytes` | 字节数据 |
| `Product` | 自定义类型 |
| `[Product]`、`[String]` | 数组，元素类型写在方括号内 |
| `{String: String}` | 映射，花括号内依次写键类型和值类型 |
| `String?` | 可选类型，问号放在类型后面 |
| `shop.Product` | 使用包名限定的类型引用 |

这些名称由核心库或业务包定义，ANTLR 通过类型标识符解析它们。
语法也支持 `Array<T>` 一类泛型类型引用；具体类型及其生成结果取决于
类型定义和所选生成目标。

### 类型别名

`type 名称 = 类型` 定义类型别名，右侧可以是已有类型或组合类型：

```mojo
type ProductId = String
type ProductList = [Product]
```

注意区分：`type Product { ... }` 定义结构体，`type ProductId = String`
定义别名。语法还接受 `struct Product { ... }` 和
`type Product = struct { ... }`；业务定义统一采用第一种结构体写法即可。

### Entity 标注

需要数据库实体时，将 `@entity` 放在结构体声明之前：

```mojo
/// 商品分类
@entity
type Category {
    id: String @1
    name: String @2
}
```

Entity 仍然是 `type` 定义的结构体。`@entity` 的数据库含义由编译器和
生成器处理，详细的主键识别及 CRUD 规则见 [Entity 数据库操作代码](entity-models.md)。

## 3. 使用 interface 定义接口

接口由名称和方法签名组成。方法不写 `func`，也不在接口中编写函数体：

```mojo
/// 商品服务
interface Catalog {
    /// 获取商品
    get_product(id: String @1) -> Product

    /// 查询商品列表
    list_products(keyword: String @1, limit: Int32 @2) -> [Product]

    /// 创建商品
    create_product(product: Product @1) -> Product

    /// 删除商品，不声明返回类型
    delete_product(id: String @1)
}
```

方法签名的常用形式为 `方法名(参数名: 类型 标注, ...) -> 返回类型`。
无参数时写 `()`；不声明返回类型时省略整个 `-> 返回类型` 部分。
接口只描述调用约定，NCraft 服务的业务逻辑在生成的 Go handlers 中实现。

参数的 `@1`、`@2` 等编号在每个方法的输入参数内独立；不同方法可以
分别从 `@1` 开始。需要返回多个业务字段时，通常定义一个响应结构体，
将其作为返回类型。

语法也接受 `type Catalog = interface { ... }`，本文统一使用
`interface Catalog { ... }` 的写法。

### 为服务方法添加 HTTP 标注

在方法前指定 HTTP 方法与路径，在参数类型后指定传输方式：

```mojo
/// 商品 HTTP 服务
interface Catalog {
    /// 从路径读取商品 ID
    @http.get("/shop/v1/products/{id}")
    get_product(id: String @1) -> Product

    /// 从 query 读取查询条件
    @http.get("/shop/v1/products")
    list_products(keyword: String @1, limit: Int32 @2) -> [Product]

    /// 将商品对象作为整个请求体传入
    @http.post("/shop/v1/products")
    create_product(product: Product @1 @http.body) -> Product

    /// 同时使用路径参数和请求体
    @http.put("/shop/v1/products/{id}")
    update_product(id: String @1, product: Product @2 @http.body) -> Product
}
```

NCraft 的 HTTP 生成规则将路径占位符 `{id}` 绑定到同名参数；
`@http.body` 表示该参数通过请求体传入。未绑定路径且未标记 body 的
普通参数通过 query 传入，例如 `?keyword=phone&limit=10`。
多个请求体字段应放进一个结构体，再将对应参数标记为 `@http.body`。

标注的语法统一以 `@` 开头，可以是数字、名称、限定名称或带参数的名称，
例如 `@1`、`@entity`、`@http.body`、`@http.get("/products")`。
它们的具体行为由相应的编译器处理，ANTLR 规则定义其书写形式。

## 4. 注释与文档说明

| 写法 | 用途 |
| --- | --- |
| `// 注释` | 普通单行注释 |
| `/* 注释 */` | 普通块注释，可跨行 |
| `/// 说明` | 放在类型、字段或方法之前的文档说明 |
| `//< 说明` | 放在字段或参数之后的文档说明 |

需要出现在生成文档中的描述使用 `///` 或 `//<`。多行参数也可以分别
添加尾随说明；使用逗号分隔时，逗号写在注释之前：

```mojo
interface ProductSearch {
    /// 根据关键字分页查询商品
    search(
        keyword: String @1, //< 搜索关键字
        limit: Int32 @2 //< 返回数量
    ) -> [Product]
}
```

## 5. 对照 ANTLR 规则

需要查询完整语法时，可在 `MojoParser.g4` 中查找以下规则：

| 内容 | 规则名 |
| --- | --- |
| 结构体与字段 | `structDeclaration`、`structMemberDeclaration` |
| 类型别名 | `typeAliasDeclaration` |
| 接口与方法 | `interfaceDeclaration`、`interfaceMethodDeclaration` |
| 参数与返回值 | `functionParameter`、`functionResult` |
| 类型、数组和映射 | `type_`、`arrayType`、`mapType` |
| 标注 | `attribute`、`attributes` |
| 文档注释与分隔符 | `document`、`followingDocument`、`eosWithDocument`、`eovWithDocument` |

ANTLR 还描述了泛型声明、继承、嵌套类型、联合类型等形式；本文聚焦服务
开发最常用的定义。语法可解析不代表每个生成目标都支持对应形式。
