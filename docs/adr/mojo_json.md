
### TypeAlias
```
type A = String
```

```
type B = Int | String | Array<String> | DateTime | Color
```

```
type C = Array<String>
```


#### AST

```
    NonimalType {
        name: "String"
        attributes: [{name: "type_alias", arguments:[value: "A"]}]
    }
```

#### Protobuffer

```
message B {
    option (mojo.boxed) = true;
    option (mojo.type_alias) = true;

    oneof b {
        int32 int_val = 1;
        string string_val = 2;
    }
}
```

```
message C {
    option (mojo.boxed) = true;
    option (mojo.type_alias) = true;

    repeated string values = 1;
}
```

#### json

#### Xml


### Inheritance(继承)

```
type A : String
```

```
type B : Int | String | Array<String> | DateTime | Color
```

```
type C : Array<String>
```

#### Protobuffer

```
message A {
    option (mojo.boxed) = true;

    string value = 1;
}
```

```
message B {
    option (mojo.boxed) = true;

    oneof b {
        int32 int_val = 1;
        string string_val = 2;
    }
}
```

```
message C {
    option (mojo.boxed) = true;

    repeated string values = 1;
}
```

如果是直接从`[[String]]`转换而来的话，则自动生成如下的message
```
message Strings {
    option (mojo.boxed) = true;
    option (mojo.auto_generated) = true;

    repeated string values = 1;
}
```

### Union

Union类型转换成json，由于json的类型较少，如果无法区分的情况下，需要增加`discriminator`来区分不同类型。

1. 多个string类型的区分方案
2. 多个Array类型的区分方案
3. `discriminator`字段添加方案
   1. 对于基本类型的处理方案
   2. 对于Array类型的处理方案
   3. 对于Dictionary类型的处理方案
4. Union内嵌Union
5. 