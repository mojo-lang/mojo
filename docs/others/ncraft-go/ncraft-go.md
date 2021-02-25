ncraft-go 基于gokit的Cloud Native微服务框架

ncraft-java基于spring bootde 



ncraft-go

ncraft-generator

ncraft-go微服务框架代码生成工具



```
Mojo Type

Url
Timestamp
Duration
Version
```

```
Protobuf Type
```



```
golang Type
```



```
Union<T | Reference>
Union<T | String>
Union<T | Bytes>

Union<Document | String>

Union<Url | String>
Union<Date_time | String>
Union<Version | String>


Union<T | Error>
```



```
protobufer

message Foo {
	string datetime = 1; //@format(Date_time)
}

message Foo {
	Datetime datetime = 1;
}

message Foo {
	Date_time datetime = 1;
}

message Date_time {
}

message Version {
}

message Url {
}

message Any {
}
```











