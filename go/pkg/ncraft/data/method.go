package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Method struct {
    Decl         *lang.FunctionDecl
    Name         string
    StandardName string // get list create update delete
    CustomName   string

    Request  *Message
    Response *Message
    Entity   *Message

    Inherit *Interface

    IsEntity    bool
    IsInherited bool
    IsStandard  bool
    IsBatch     bool

    // Bindings contains information for mapping http paths and parameters onto
    // the fields of this Methods.
    Bindings []*HTTPBinding
}
