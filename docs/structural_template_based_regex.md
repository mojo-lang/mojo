# structural template language based on regex

RegTemplate

基于正则表达式的结构化模板语言

## 目的

1. capture variable （变量提取）
2. extend variable （变量扩展）
3. verify （验证）

## 使用场景
在MOJO中，struct类型与String类型（非JSON、XML、YAML等格式）的codec。

比如DateTime的格式，EMail、HostName、IPV4、IPV6

## 语法

```abnf
template = 1* element
element = regex-element / optional-block / expression

optional-block = "[" template "]"

expression = "{" [ variable ] [ ":" restrict-expression ] "}"
expression = "{" function "}"

restrict-expression = template / function
function = "@" identifer [ "<" ">" ] [ "("  ")" ]
```

## Examples

```RegTemplate
{:.{2,4}-{@string}}
{:.*}
{:@string}
{@string}

/abc/{id}[.{format:@type<Element>}]

/abc/{id}[.{format:@string}]
/abc/{id}[.{format:@array('.')}]

$

/abc/{}/xyz
/abc/{:[^/]+}/xyz
/abc/{@array('/')}/xyz[?{@map('&', '=')}]

{longitude},{latitude}
T{kilometer}[[+.]{meter}]
{labels: @array(".", 1, 127)}
{top_left.longitude},{top_left.latitude},{right_bottom.longitude},{right_bottom.latitude}

{top_left:{longitude},{latitude}},{top_right:{}} | {}

{longitude},{latitude}[,{altitude}]
{type}/{facet}[-{producer}]-{product}-{media}.{suffix}
{major}.{minor}.{path}[-{pre_release}][+{build_metadata}]
```

## Compile

RegTemplate最终都可以编译成标准的Regex；对应Url的template可以编译成Uri Template。

编译参数：
1. 对于匿名表达式（即没有variable的expression），可以选择是否进行capture；对于纯验证目的的情况下，可以优化Regex；
2. 对于非匿名表达式，也可以进行是否capture设置；