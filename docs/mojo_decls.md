DataDecl

类型的定义可以使用如下语法，具体的类型由系统推导

type Number = Int   类型别名
//type Number : Int

type HouseNumber : Poi { 定义struct
}

type City : Int { 定义enum
    Shanghai = 1
}

type House { 定义interface
    func xxx() -> xxx
}

也可以使用如下的语法：

struct HouseNumber : Poi {
}

enum City : Int { 定义enum
    Shanghai = 1
}

interface House { 定义interface
    func xxx() -> xxx
}