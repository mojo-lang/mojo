```
db.users.insertOne(
   {
      name: "sue",
      age: 19,
      status: "P"
   }
)
```

```
users | insert({
    name: ""
    age:
    status: "A"
})
```


```
db.users.find( { status: "A" } )
```

```
users | status == 'A'
```

```
db.users.find( { status: { $in: [ "P", "D" ] } } )
```
```
users | status in ["P", "D"]
```

```
db.users.find( { status: "A", age: { $lt: 30 } } )
```
```
users | status == "A" && age < 30
```

```
db.users.find(
   {
     $or: [ { status: "A" }, { age: { $lt: 30 } } ]
   }
)
```

```
users | status == "A" || age < 30
```

```
db.users.find( { favorites: { artist: "Picasso", food: "pizza" } } )
```

```???
users | contains({favorites: {artist : "Picasso", food: "pizza"}})

users | (favorites | artist == "Picasso" && food == "pizza")
users | favorites.artist == "Picasso" && favorites.food == "pizza"

```

```
db.users.find( { badges: [ "blue", "black" ] } )
```

```
users | badges == ["blue", "black"]
```

```
db.users.find( { badges: "black" } )
```
```
users | badges.contains("black")
```

```
db.users.find( { "badges.0": "black" } )
```
```
users | badges[0] == "black"
```

```
db.users.find( { finished: { $elemMatch: { $gt: 15, $lt: 20 } } } )
```
```
users | (finished | $ in 15<..<20)
```

```
db.users.find( { 'points.0.points': { $lte: 55 } } )
```
```
users | points[0].points <= 55
```

```
db.users.find( { 'points.points': { $lte: 55 } } )
```
```
users | (points | $.points <= 55)
```

```
db.users.find( { points: { $elemMatch: { points: { $lte: 70 }, bonus: 20 } } } )
```

```
users |  ($.points | points <= 70 && bonus == 20 | reduce)
```
