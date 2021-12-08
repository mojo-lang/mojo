

type U<T> = String @1 | T @2

type Foo {
    u: U<Int32> @1
    v: U<Int32> @2
}
