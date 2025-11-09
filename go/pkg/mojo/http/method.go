package http

import "strings"

const MethodTypeName = "Method"
const MethodTypeFullName = "mojo.http.Method"

func (x Method) CanCarryBody() bool {
    switch x {
    case Method_METHOD_POST, Method_METHOD_PUT, Method_METHOD_PATCH:
        return true
    }
    return false
}

func CanCarryBody(method string) bool {
    m := new(Method)
    if err := m.Parse(strings.ToUpper(method)); err != nil {
        return false
    }
    return m.CanCarryBody()
}
