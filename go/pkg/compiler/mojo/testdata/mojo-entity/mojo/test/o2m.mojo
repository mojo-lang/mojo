
type OMUser {
    id:   String @1
    name: String @2
    age:  Int @3

    pets: [Pet] @5 @reference
}

type Pet {
    id: String @1
    name: String @2
    owner: OMUser @3
}

type OMUser2 {
    id:   String @1
    name: String @2
    age:  Int @3

    pets: [Pet2] @5
}

type Pet2 {
    id: String @1
    name: String @2

    owner_id: String @3 @back_reference('OMUser2')
}

type OMUser3 {
    id:   String @1
    name: String @2
    age:  Int @3

    pets: [Pet3] @5
}

type Pet3 {
    id: String @1
    name: String @2

    owner: OMUser3 @3
}

type OMNode {
    id: String @1
    value: Int @2

    children: [OMNode]
    parent: OMNode @back_reference
}

type OMNode2 {
    id: String @1
    value: Int @2

    children: [OMNode2]
    parent: OMNode2
}

