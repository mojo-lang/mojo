
```
func test(x:Int, y:Int) -> Int {
  return x + y;
}
```

```
func test(x, y) {
  return x + y;
}
```

```
test = (x:Int, y:Int) -> {
  return x + y;
}
```

```
test = (x, y) -> {
  return x + y;
}
```

```
test = (x:Int, y:Int) -> x + y
```

```
test = (x, y) -> x + y
```

```
test = x -> x + 3
```

```
test = $0 + $1
```

Expression Lambdas

A lambda expression with an expression on the right side of the => operator is called an expression lambda. Expression lambdas are used extensively in the construction of Expression Trees. An expression lambda returns the result of the expression and takes the following basic form:

Statement Lambdas
A statement lambda resembles an expression lambda except that the statement(s) is enclosed in braces:
(input parameters) -> {statement;}
