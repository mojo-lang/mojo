
type Int

type A : Array<Int>


type B<T> {
    type C<U> {
        foo : [String]
    }

    a: Int32
    b: String
    c: C<T>
    d: T
}

enum Foobar : Int {
    Foo :1
    Bar :2
}

Foobar.Foo
Foobar.Bar

type Test {
    a: Foobar
}

Type_alias {
}

Data_type_decl {
    name: 'Int'
}

Data_type_decl {
    name: 'Array'
    generic_params: [{
        name: 'T'
    }]
}

Data_type_decl {
    name: 'Map'
    generic_params: [{
        name: 'Key'
    },{
        name: 'Value'
    }]
}

Type_decl {
    name: 'A'
    type: Type_name {
        name: 'Array'
        generic_arguments: [{
            Type_name{
                name: 'Int'
            }
        }]
    }
}

Struct_decl {
    name: 'B'
    generic_params: [{
        name: 'T'
    }]
    type: Struct_type {
        fields: [{
            name: 'a',
            type: ''
        }, {
        }]
    }
    struct_decls: [{
        name: 'C'
        generic_params:[]
        type: Struct {
            fields:[{
                
            }]
        }
    }],
    enum_decls: []
    type_alias_decls: []
}