# 如何让Protobuf支持Union类型

## 何为Union类型

## 为什么需要UNION类型

## UNION和oneof的差异

protobuf的oneof只是值的oneof，而并非传统类型系统的Union，

```protobuf
message MsgWithOneof {
    oneof union {
      string title = 1;
      int64 salary = 2;
      string Country = 3;
      string home_address = 4;
    }
}
```

## 在不改变现有protobuf工具的情况下如何实现

## 我们如何实现

```protobuf
message MsgWithOneof {
   option (mojo.union) = true;
   option (gogoproto.onlyone) = true;
   
   string string_value = 1;
   int64 	 int_value = 2;
}
```

```

Array<Array<Test>>


message TestArray {
  option (mojo.array) = true;
  option (mojo.json_unwrap) = false;
  
  repeated Test tests = 1;
}

repeated Tests tests = 2;

message TestDictionary {
  option (mojo.dictionary) = true;

  map<string, Test> tests = 1;
}

Tuple类型 转换成 message，序号默认从1开始


```


```

function的request和response的转换规则


func test(int : Int@1,
          string : String @2) -> (result : Reuslt @1)

message TestRequest {
}

message TestResponse {
}


Duration

Timestamp

Int?


```