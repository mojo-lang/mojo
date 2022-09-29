
type A<T1, T2, T3> {
    t1: T1 @1
    t2: T2 @2
    t3: T3 @3
}

type B<T1, T2> = A<T1, T2, String>

type R<T> = B<T, Int>

type U<T1, T2> = Union<T1, T2, String> 

type V = U<String, Int>

type IntArray = [Int64]

type Foo {
    v: V @1

    r: R<DateTime> @2

    rv: R<V> @3

    arrays: [IntArray] @4

    u: U<String, Int32> @5
}

type H<T> = [T]

type J = H<String>

type Foo2<T> {
    bar: T @1
}

type Bar = Foo2<String>

type FooBar: Foo2<Bar>



