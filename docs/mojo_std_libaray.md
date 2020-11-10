# Mojo Standard Libaray

[TOC]

##  Types

### Error

### DateTime

```
1. 2016-12-12
2. 2016-12-12T12:34:45.345Z
   1. 2016-12-12T12:34:45.345+08:00
```

### Duration

1. 1. Year: 3Y
   2. Month: 1M
   3. Weak: 1W
   4. Day: 1D
   5. Hour: 2h
   6. Minute: 1min
   7. Second: 2sec
   8. MillionSecond: 3ms
   9. MicroSecond: us
   10. NanoSecond: 5ns

###

## Functions

### Cores

#### `typeof`

#### `is_type`

`typeof(Any) -> Type`
`is_type(Any, typename:Symbol) -> Bool`

#### `length`

````
length(Any) -> Size
````

### Numbers

### Strings

`substr(str:String) -> String`

### Datetimes

#### now

```
func now() -> DateTime;
func local_now() -> DateTime;
```


monday(t:DateTime) : Bool
tuesday(t:DateTime) : Bool

### Arrays

#### slice

```
slice(Sequence, Integer) -> Array
slice(Sequence, from:Integer, to:Integer) -> Array
```

#### sort

```
sort(Sequence) -> Array;
```

#### count

```
count(Sequence) -> Integer;
```

#### 示例

```
for(v:[12,12,12,34,45]) {}
[12,12,12,34,45] | (v){[v+1, v+2, v+3]} | flat()
```

### Objects

#### each

```
each(object)
each(object, path)
```

#### tree

	tree(object)
	tree(object, path)
Pair:Object

```
to_array(Object) -> Array(Pair)
to_object(array:Array(Pair)) -> Object
```

```
obj = {
		alice: [{},{}],
		bob: [{}]
};

to_array(obj)

obj | to_array

[
	{
		key: alice,
		value: [
			{},
			{}
		]
	},
	{
			key: bob,
			value: [
				{},
				{}
			]
	}
]
```
