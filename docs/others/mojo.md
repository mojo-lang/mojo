# Mojo

是Feature_db的数据查询语言。

## 词汇表
1. Object
2. Graph
3. No_sql: Not Only SQL
4. FQL: Feature Query Language
5. FPL: Feature Presentation Language
6. DSL: Domain Specification Language
7. RPC: Remote Procedure Call

## 系统组件

### 语言层
语言层是对数据的操作封装层，对外提供统一的语言接口。具体的语言采用FQL以及FPL。
语言层包含三个子组件：
1. Mojo Parser

    负责对输入的字符串进行语法解析，生成AST语法树。

2. Mojo AST接口层

    对输入的AST语法树进行语义分析，进行初步优化。
    AST语法树可以由Mojo语言解析器提供，或者由外部客户端通过RPC提供；客户端的AST是由各个语言的客户端通过语言内DSL生成。
    AST的RPC接口采用Google Protobuf协议。

3. Mojo Planer

    对于输入的AST语法树进行优化后，调用分布式计算框架计算结果。


## Mojo(Feature Query Language):

针对GIS应用领域，基于Feature_d_b的，一门动态类型语言，是一种偏向于函数式描述性语言。

### 特点
1. 简单的语法，比SQL简单；
2. 更强的表述能力；
3. 方便自定义函数扩展；
4. Pipeline处理流；
5. 使用基于分布式计算框架的执行器。

### 语言说明

### 语言meta类型

1. Meta Object

   ```
   type Object {
   }
   ```

   

2. Meta Function

   ```
   func foo () {
   }

   func foo(a:Int, b:Int) -> a+b;

   foo = (a:Int, b:Int) => a+b;
   foo = a:Int, b:Int -> a+b;

   foo = (a:Int, b:Int) {
   };
   ```

   

3. Enum类型

   	enum Week_day {
   	  	Monday = 1
   		Tuesday,
   		Wednesday,
   		Thursday,
   		Friday,
   		Saturday,
   		Sunday
   	}
 	
   	enum {
   	}
   ​

### 语言基本类型

1. Null
   null

2. Bool
    `true`
   `false`

3. Number（数值）
   1. Integer 整数
   2. 16进制
      1. "#" Hex_digits
   3. Real 浮点数
      inf
   4. 

4. String（字符串） "" ''

5. String Template  t"{}"

6. Bytes (数据)

7. Regex (正则表达式) r'^sss$'

8. 数量   p t m k  PB TB MB 34KB ?

9. Duration （时间间隔）
   1. Year: 3Y
   2. Month: 1M
   3. Weak: 1W
   4. Day: 1D
   5. Hour: 2h
   6. Minute: 1m
   7. Second: 2s
   8. Million_second: 3ms
   9. Micro_second: us
   10. Nano_second: 5ns

10. 
11. Tuple（元组）
   省略field name的对象，field的次序应该与对象的layout一致，不存在赋值为null

12. Sequence/Array （数组，相同Type的元素列表）[]

13. Sequence为bound或者unbound的数组

14. Array为定长数组

   //Collection ＝ Sequence<Any>

17. Range
    [4..]
    4..8
    4..<8
    
18. Object（对象）｛｝

   Generic Object
   {}

   Typed Object
   Type_symbol {}

15.    Express（表达式）
16.    一元表达式
                  ```
                  not Express
                 	! Express
                  + Express
                 	- Express
                  ```
17.    二元表达式
                  算术符
               ```
               		Express + Express
               		Express - Express
               		Express * Express
               		Express / Express
               		Express % Express
               		Express ^ Express
               ```
          ​     比较符
               ```
               		Express > Express
               		Express < Express
               		Express == Express
               		Express >= Express
               		Express <= Express
               		Express != Express
               ```
          ​     逻辑符
               ```
               		Express && Express
               		Express || Express
               ```

18.    括号表达式

19.    字面量

       5.    变量
             1. 变量名
             2. 无参数的函数名

       6.    函数调用表达式

             7.   lamda表达式

             8.   lamda arrow表达式

                        arg1,arg2->return express
                        t,v->t+v

             9.   pipeline表达式

20.    自定义类型(标准库)
21.    Geometry
               2. Feature
               3. Features = Sequence<Feature>
               4. Track : Feature
               5. Collection
               6. Db

### 语句：
	1. 赋值语句
	2. 表达式语句
	3. 类型定义语句
	​```
	type Geometry {
	}
	enum Geometry_type {
	}
	​```
	3. 函数定义语句
	func function() {
	}
	
	4. if
	
	5. match
		match(obj) {
	      1: print("one");
	      _: Error{};
		}
	
	6. for
		for (i in Sequence<Object>) {
		}

### Pipeline处理流：

```
function_1(args...) | function_2(args...)
等效于：
function_2(function_1(args...), args...)
```



### 类型处理函数
```
type key_space.User {
}

create_type(key_space.User, Meta_object);
```

```
update_type(key_space.User, Meta_object);
```
