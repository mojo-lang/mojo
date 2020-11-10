## Mojo语言的多态

分为语言编译时以及runtime时的两种多态形式，编译时的多态是通过template（模板）来实现。

template类型相当于类型的函数，需要参数的代入才能计算出最终的类型。mojo语言的模板相比较C++而言是比较简单的，不支持各种偏特化特性。

### 类型的组合与继承

```
type Place {
}

type Country : Place {
}

type Place : [String] {
}

type Foo = {String, [String]}
```



### 多态类型的序列化

mojo语言当作为IDL使用时，兼容protobuf以及thrift两种标准；同时也支持json、xml、yaml、toml等通用文本序列化方式。



