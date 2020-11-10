

{
    empireHero : hero | episode == empire | { name }
    jediHero : hero | episode == jedi | name
}

pj = (c:Character) -> {
    name
    appearsIn
    friends: c.friends | @include(with_friends) { name }
}

{
    left : hero | $.episode == empire | pj
    right : hero | $.episode == jedi | {
        name : $.name
        @mixin: pj($)
        yuu: pj($)
}


{
    key1
    key2
    key3:89 @
    @mixin: {}
    @merge
    @overlay
}

{
    key1:value1
    key2:value2
}

hero | episode == ep | {
    name
    @mixin : $.match {
                 is Droid => { $.primary_function }
                 is Human => { $.height }
                 _        => {}
             }

    @skip(x > 5) {
        ty: 12
    }
}

match = (value:Any) -> {
    match value {
        is Droid => {}
        is Human => {}
        _        => {}
    }
}



hero | episode == ep | {
    name
    @mixin : (o) -> {
        match o {
                 is Droid => { primary_function }
                 is Human => { height }
                 _        => {}
        }
    }($)
}

hero | episode == ep | reviews.push(review) | {}

func length(s: Starship) -> Int {
    return s.ss + s.gg;
}

Starship ss; ss.length()

{
data: {
    hero: hero | { name }
    droid: droid | id == 2000 | { name }
}
}

persons | ~(a : name == 'Jim')-[knowns]->(b)-[knowns]->(:c) && (:a)-[:knowns]->(:c)  

User {

@link
type Rated  {
    starts: Int,
    movie: Moive
}

rated: [Rated]


}

movie | min(year)

movie | name == '' | (*)<-[r:rated]-(*) | 


category | {

    a: name
    b: ()<-[:category]-(moive) | 

}

movie | (()-[r:rated]->() | count() > 10) | {
    a: name
    b: ()->[rated: stars]->() | average
} | sort(b) | top(10)


movie | ()-[category]->(:name=='Action') | year in 1980..1990 | ()<-[b:rated]-() | b.stars == 5