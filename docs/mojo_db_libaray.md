# Mojo Database Libaray

[TOC]

## Types

### Db

### Collection(Table)



## Functions

### sort

```
ascend() = (l:Object, r:Object) {l.key < r.key}
ascend(field:Simbol) = (l:Object, r:Object) {l.field < r.field}

decend() = (l:Object, r:Object) {l.key < r.key}
descend(field:String) = (l:Object, r:Object) {l.field < r.field}

sort(array: Sequence<Object>[, expression: Expression]) : Sequence<Object>
```

### limit
```
limit(n: Integer)
```

### skip
```
skip(n: Integer)
```

### update
```
update(src: Collection, updated: Feature) -> Feature
update(src: Features, updated: Feature) -> Features
update(src: Feature, updated: Feature) -> Feature
```

示例：

```
db = Db()
db.my_teams.1234.teams += 'xxx'

db.my_teams.update({})


db.my_teams | | update({}, {partial: true})

update(my_teams) | id==1234 | {teams: ['xx', 'xxx']}
update(my_teams) | id==1234 | {teams: teams + 'xxxxx'}
update(my_teams) | id==1234 | {teams: teams - 'xxxxx'}
update(my_teams) | id==1234 | {teams: []}

update(students) | id == 5 | {
	scores: (scores + []) | sort(score | decend) | slice(-3)
}
```

### insert
```
insert(src: Collection, item: Object) -> Collection
```
### upsert
```
upsert()
```
### replace
```
replace()
```
### remove
```
remove()
```

### aggregate

```
group(Squence<Any>[, field_name]) -> Group<Squence<Any>>
```

Group<Object> : Object
Group<Integer>
Group<Squence<Object>>



#### 创建collection
```
create(collection, type:Simbol)
create(collection, type:Simbol, meta_obj:MetaObject) : Collection<Simbol>
```

## 示例：
```
简单的定义：

详细的对象定义：

type User : Table {
	  geom: Point
	  id: String @prime_key
    name: String @full_text
};

create<User>(name:String);

create<User>(users);
create<User>(options:Object);
```

#### 删除collection
```
remove(collection);
```

#### 变更collection的Object


### Sql vs FQL
```

select * from users;
select id, user_id, status from users;
select * from users where status = 'A';

users
users | {id, user_id, status}
users | status == "A"
users | status != "A"
users | status == "A" and age == 50
users | status == "A" or age == 50
users | age > 25
users | age in 25 .. 50
users | id != /bc/
users | id == /^bc/

pois | search("北京天安门")

users | status == "A" | sort(user_id, asc)
users | count
users | count(id)

users | age > 30 | count
users | distinct(status)
users | limit(1)
users | limit(5) | skip(10)

users | age > 25 | update({status: "C"})
users | status == "A" | update({age: age + 3})
```

根据test表单的address字段，自动添加经纬度
```
test | update({geom: geocode(address)})
```

根据test表单的geom字段，自动添加位置信息
```
test | update({
    address: geocode(geom)
    type: switches(type, (1, 45), (3, 67)) 
})
```
道路的合并

```
roads | group(nfc, (r1,r2){ r1.connected(r2) }) | merge | thinning(road_nodes)
```
检测道路shape point密度，并可视化
```
for z in 3..15 {
    roads | grid_clip(z) | {grid: $, point_count: ($ | point_size | sum) } | attach_style({
	    fill-color: steps(point_size, [100, #ffffff], [300, #ffffdd], [1000, #ffffaa], [2000, #ffff00])
    })
}

```
