```graphql
{
    user(id:4) {
        name
    }
}
```

```mojo

user | id == 4 | {
    name
}
```

```
mutation {
  likeStory(storyID: 12345) {
    story {
      likeCount
    }
  }
}
```
```
update(story) | story_id == 12345 | ++like_count | {
    story: {
        likeCount
    }
}
```

```
{
  me {
    id
    firstName
    lastName
    birthday {
      month
      day
    }
    friends {
      name
    }
  }
}
```

For example, this operation selects fields of complex data and relationships down to scalar values.
```
{
  me {
    id
    firstName
    lastName
    birthday {
      month
      day
    }
    friends {
      name
    }
  }
}
```

```
table_name | {
    me: {
        id
        firstName
        lastName
        birthday: {
            month
            day
        }
        friends: {
            name
        }
    }
}
```

```
{
  user(id: 4) {
    id
    name
    profilePic(width: 100, height: 50)
  }
}
```
```
user | id == 4 | {
    id: $.id
    name: $.name
    profilePic: $.profilePic | width == 100 && height == 50 | url
}
```

```
{
  user(id: 4) {
    id
    name
    smallPic: profilePic(size: 64)
    bigPic: profilePic(size: 1024)
  }
}
```

```mojo

```

```
query withFragments {
  user(id: 4) {
    friends(first: 10) {
      ...friendFields
    }
    mutualFriends(first: 10) {
      ...friendFields
    }
  }
}

fragment friendFields on User {
  id
  name
  profilePic(size: 50)
}
```

```
friendFields = _:User -> {
    id
    name
    profilePic: resize(50)
}
user | id == 4 | {
    friends: first(10) | friendFields
    mutualFriends: first(10) | friendFields
}

```

```
query inlineFragmentNoType($expandedInfo: Boolean) {
  user(handle: "zuck") {
    id
    name
    ... @include(if: $expandedInfo) {
      firstName
      lastName
      birthday
    }
  }
}
```

```
func query(expandedInfo: Bool) -> Object {
    return user | handle == "zuck" | {
        id
        name
        @include(expandedInfo) {
            firstName
            lastName
            birthday
        }
    }
}
```

```
union SearchResult = Photo | Person

type Person {
  name: String
  age: Int
}

type Photo {
  height: Int
  width: Int
}

type SearchQuery {
  firstSearchResult: SearchResult
}

{
  firstSearchResult {
    ... on Person {
      name
    }
    ... on Photo {
      height
    }
  }
}
```

```
{
    firstSearchResult: {
         @include(is Person) {
            name
            age
        }
        
        @include(is Photo) {
          height 
        }
    }
}

```

```
type User {
  id: String
  name: String
  birthday: Date
}

{
  __type(name: "User") {
    name
    fields {
      name
      type {
        name
      }
    }
  }
}
```
```
User.type | {
    name
    fields: {
        name
        type: $.type.name
    }
}
```