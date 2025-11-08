Range

1. 语法表示
   1. ..
   2. ..<
   3. ..=
   4. <..
   
2. 默认闭开，还是闭闭
3. 是否提供step
   1. 2..10 | step(2)
   2. [2..10:2]

## C#

```
public readonly struct Range
    {
        public Range(System.Index start, System.Index end);
        public static Range StartAt(System.Index start);
        public static Range EndAt(System.Index end);
        public static Range All { get; }
    }
```

The .. syntax allows for either, both, or none of its arguments to be absent. Regardless of the number of arguments, the Range constructor is always sufficient for using the Range syntax.

https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/proposals/csharp-8.0/ranges


## Closure

```
(range)(range end)(range start end)(range start end step)
```

Returns a lazy seq of nums from start (inclusive) to end
(exclusive), by step, where start defaults to 0, step to 1, and end to
infinity. When step is equal to 0, returns an infinite sequence of
start. When start is equal to end, returns empty list.

https://clojuredocs.org/clojure.core/range


## Kotlin

(..) operator
It is the simplest way to work with range. It will create a range from the start to end including both the values of start and end. It is the operator form of rangeTo() function. Using (..) operator we can create range for integers as well as characters.
Kotlin program of integer range using (..) operator –

```kotlin
fun main(args : Array<String>){

	println("Integer range:")
	// creating integer range
	for(num in 1..5){
		println(num)
	}
}
```

outpu
```
Integer range:
1
2
3
4
5
```
https://www.geeksforgeeks.org/kotlin-ranges/


## Swift


https://www.programiz.com/swift-programming/ranges

In Swift, a range is a series of values between two numeric intervals. For example,

```swift
var numbers = 1...4
```

... is a range operator
1...4 contains values 1, 2, 3, 4
1 is lower bound (first element)
4 is upper bound (last element)
Types of Range in Swift
In Swift, there are three types of range:

Closed Range
Half-Open Range
One-Sided Range


## python

https://www.runoob.com/python/python-func-range.html

range(start, stop[, step])
参数说明：

start: 计数从 start 开始。默认是从 0 开始。例如range（5）等价于range（0， 5）;
stop: 计数到 stop 结束，但不包括 stop。例如：range（0， 5） 是[0, 1, 2, 3, 4]没有5
step：步长，默认为1。例如：range（0， 5） 等价于 range(0, 5, 1)


## Rust

https://doc.rust-lang.org/std/ops/struct.Range.html

```rust
pub struct Range<Idx> {
    pub start: Idx,
    pub end: Idx,
}
```

A (half-open) range bounded inclusively below and exclusively above (start..end).

The range start..end contains all values with start <= x < end. It is empty if start >= end.


examples

```rust
assert_eq!((3..5), std::ops::Range { start: 3, end: 5 });
assert_eq!(3 + 4 + 5, (3..6).sum());
```
```rust
let arr = [0, 1, 2, 3, 4];
assert_eq!(arr[ ..  ], [0, 1, 2, 3, 4]);
assert_eq!(arr[ .. 3], [0, 1, 2      ]);
assert_eq!(arr[ ..=3], [0, 1, 2, 3   ]);
assert_eq!(arr[1..  ], [   1, 2, 3, 4]);
assert_eq!(arr[1.. 3], [   1, 2      ]); // This is a `Range`
assert_eq!(arr[1..=3], [   1, 2, 3   ]);
```