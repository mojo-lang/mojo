# Mojo 简明语法：表达式

本文依据 [ANTLR 解析规则](../antlr/mojo/MojoParser.g4)、
[词法规则](../antlr/mojo/MojoLexer.g4) 和当前
[表达式 AST 转换实现](../go/pkg/compiler/mojo/parser/syntax/expression_visitor.go)，
介绍 Mojo 表达式的常用写法。类型与服务接口见 [type 与 interface](syntax.md)。

表达式用来描述值、访问数据或组合计算，可以出现在初始化值、标注参数、
调用参数和集合元素等位置。以下示例侧重语法；变量、函数及运算符的实际
求值能力由使用表达式的编译器或运行时决定，不表示 NCraft 会将任意表达式
自动生成为业务代码。

## 1. 字面量与名称引用

常用的简单字面量包括数字、字符串、布尔值和空值：

```mojo
42
-12
3.14
"hello"
'hello'
true
false
null
```

字符串可以使用单引号或双引号；布尔值与空值使用小写关键字。
类型名 `String`、`Bool` 与值 `"hello"`、`true` 是不同的概念。

已命名的值直接通过名称引用，例如 `price`、`quantity`。
在声明中，`=` 右侧就是初始化表达式：

```mojo
var page_size = 20
var total = price * quantity
```

这里的 `price`、`quantity` 假设已经在使用环境中定义。

## 2. 数组、映射、对象和结构体值

```mojo
[1, 2, 3]
["phone", "tablet"]
{"color": "black", "size": "large"}
{name: "Phone", price: 99.0}
Product{name: "Phone", price: 99.0}
```

这几种写法的区别是：

| 写法 | 含义 |
| --- | --- |
| `[值, 值]` | 数组字面量 |
| `{"键": 值}` | 映射字面量，键使用字符串字面量；语法也允许整数字面量键 |
| `{字段名: 值}` | 对象字面量，字段名使用标识符 |
| `Product{字段名: 值}` | 指定类型的结构体字面量，假设已定义 `Product` |

数组、映射和对象中的元素可以是其他表达式，也可以嵌套。
它们支持逗号或换行分隔，并允许末尾分隔符：

```mojo
{
    name: "Phone"
    tags: ["mobile", "new"]
    metadata: {"color": "black"}
    total: price * quantity
}
```

注意区分类型与值：`[String]` 是数组类型，`["phone"]` 是数组值；
`{String: String}` 是映射类型，`{"color": "black"}` 是映射值。

## 3. 运算符与括号

常用运算符包括：

| 类别 | 写法示例 |
| --- | --- |
| 算术 | `a + b`、`a - b`、`a * b`、`a / b`、`a % b` |
| 幂运算形式 | `a ** b` |
| 一元运算 | `-value`、`!enabled`、`not enabled` |
| 比较 | `a == b`、`a != b`、`a < b`、`a <= b`、`a > b`、`a >= b` |
| 逻辑组合 | `a and b`、`a or b`、`a && b`、`a \|\| b` |
| 成员判断 | `item in items`、`item !in items` |
| 赋值形式 | `total = value`、`total += value` |

建议二元运算符两侧留空格。使用括号表达分组，特别是混合多种运算符时：

```mojo
price * quantity + shipping
(price + shipping) * quantity
(price >= 10) and (price <= 100)
not enabled
name in ["phone", "tablet"]
```

### 当前解析器的优先级

二元运算符优先级由
[BinaryExprParser](../go/pkg/compiler/mojo/parser/syntax/binary_expr_parser.go)
决定。下表从高到低排列，同一行优先级相同：

| 优先级 | 运算符 |
| --- | --- |
| 47 | `**` |
| 45 | `*`、`/`、`%` |
| 40 | `+`、`-` |
| 30 | `<<`、`>>` |
| 25 | `..`、`..<`、`..=` |
| 22 | `in`、`!in` |
| 21 | `<`、`<=`、`>`、`>=` |
| 20 | `==`、`!=` |
| 13 | `and`、`&&` |
| 12 | `or`、`\|\|` |
| 10 | `=`、`+=`、`-=`、`**=`、`*=`、`/=`、`%=`、`<<=`、`>>=`、`&=`、`^=`、`\|=` |
| 5 | `\|` |

当前实现对同优先级的连续二元运算从左侧合并，包括 `**` 和赋值形式。
例如 `a ** b ** c` 按 `(a ** b) ** c` 构造 AST；需要其他分组时显式加括号。
表中的 `|` 是表达式运算符位置的写法；类型定义中的 `A | B` 属于联合类型语法。

## 4. 条件表达式

使用 `条件 ? 成立时的值 : 不成立时的值`：

```mojo
enabled ? "enabled" : "disabled"
(price > 100) ? (price - 10) : price
```

复杂条件、分支中的计算及嵌套条件建议加括号，明确各部分的范围。
类型后的问号（例如 `String?`）表示可选类型，与条件表达式不同。

## 5. 成员访问、下标和函数调用

```mojo
product.name
order.customer.name
products[0]
metadata["color"]
calculate_total(price, quantity)
find_products(keyword: "phone", limit: 20)
catalog.get_product("p-1").name
```

`.` 访问成员，`[...]` 表示下标访问，`(...)` 表示调用。
调用参数可以直接写表达式，也可以用 `标签: 表达式`；多个调用参数用逗号
分隔。调用参数列表与 Interface 中的参数声明不同，传入的是值，不是类型。

成员访问、下标和调用可以连续组合，例如
`catalog.get_product("p-1").name` 表示对调用结果继续访问 `name`。
实际可用的成员、下标和函数签名由对应类型及使用环境决定。

## 6. 闭包与带标记的字面量

带显式参数的闭包可以写成 `{参数 -> 表达式}`：

```mojo
{x, y -> x + y}
```

该写法声明参数 `x`、`y`，闭包体包含表达式 `x + y`；它与
`{name: "Phone"}` 这样的对象字面量不同。

ANTLR 也支持数字后缀、字符串前缀和后缀：

```mojo
12s
r"^ab$"
r"^ab$"i
```

这些例子会保留字面量及其标记，标记的含义由后续处理器定义。
不能仅凭语法就认定 `s` 一定表示秒，或 `r`、`i` 一定执行正则匹配。
标记必须紧邻字面量，不能写成 `12 s` 或 `r "^ab$"`。

## 7. 语法范围与实现边界

ANTLR 中还有范围表达式（`1 ..< 10` 等）、字符串插值、元组、
`值 if 条件 else 值`、`match`、`is`／`as` 和尾随闭包等规则。
这些规则在当前 AST 转换和后续处理中的支持程度不同，不能仅以语法文件
中存在规则判断其可用性。例如，当前整数 AST 转换只处理十进制数字，
虽然词法规则还定义了二进制、八进制和十六进制形式。

本文中的示例用于说明已核对的常用 AST 构造形式；语法解析通过不等于
完成类型检查或运行时求值。准备使用其他形式时，应同时检查对应 visitor
和目标编译器。

查阅语法时，可按以下规则定位：

| 内容 | ANTLR 规则 |
| --- | --- |
| 表达式与运算符 | `expression`、`prefixExpression`、`binaryExpression` |
| 简单字面量 | `literal`、`numericLiteral`、`stringLiteral` |
| 集合与结构体值 | `arrayLiteral`、`mapLiteral`、`objectLiteral`、`structLiteral` |
| 条件表达式 | `conditionalOperator` |
| 括号分组 | `parenthesizedExpression` |
| 成员、下标与调用 | `explicitMemberSuffix`、`subscriptSuffix`、`functionCallSuffix` |
| 闭包 | `closureExpression` |
| 带标记的字面量 | `numericOperatorLiteral`、`stringOperatorLiteral` |
