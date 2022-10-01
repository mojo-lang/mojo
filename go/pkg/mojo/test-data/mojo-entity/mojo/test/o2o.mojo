
type User {
    name: String @1 @key
    age:  Int @2
    card: Card @3 @reference
}

type Card {
    number: String @1 @db.primary_key
    expired: Timestamp @2
    owner: User @3
}

type User2 {
    name: String @1 @key
    age:  Int @2
    card: Card2 @3
}

type Card2 {
    number: String @1 @db.primary_key
    expired: Timestamp @2
    owner_name: String @3 @back_reference('User2')
}

type User3 {
    name: String @1 @key
    age:  Int @2
    card: Card3 @3
}

type Card3 {
    number: String @1 @db.primary_key
    expired: Timestamp @2
    owner: User3 @3
}

type Node {
    id: String @1
    value: Int @2

    next: Node @5 @reference
    prev: Node @6
}

type Node2 {
    value: Int @1

    next: Node2 @2
    prev: Node2 @3
}

type User5 {
    id:   String @1
    name: String @2
    age:  Int @3

    spouse: User5 @5
}
