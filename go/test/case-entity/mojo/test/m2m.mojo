

type MMUser {
    id:   String @1
    name: String @2
    age:  Int    @3

    groups: [MMGroup] @5
}

type MMGroup {
    id:    String @1
    users: [MMUser] @5 @reference
}

type MMUser2 {
    id:   String @1
    name: String @2
    age:  Int    @3

    groups: [MMGroup2] @5
}

type MMGroup2 {
    id:    String @1
    users: [MMUser2] @5
}

type MMUser3 {
    id:   String @1
    name: String @2
    age:  Int    @3

    followings: [MMUser3] @5 @reference
    followers: [MMUser3] @6
}

type MMUser4 {
    id:   String @1
    name: String @2
    age:  Int    @3

    followings: [MMUser4] @5
    followers: [MMUser4] @6
}


type MMUser5 {
    id:   String @1
    name: String @2
    age:  Int    @3

    friends: [MMUser5] @5
}
