package longrunning

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

func (x *Operation) GetValue() *core.Any {
	return x.GetResponse()
}

func (x *Operation) SetError(err *core.Error) *Operation {
	if x != nil {
		x.Error = err
	}
	return x
}

func (x *Operation) SetValue(any interface{}) *Operation {
	if x != nil {
		x.Response = core.NewAny(any)
	}
	return x
}

func (x *Operation) SetResponse(any interface{}) *Operation {
	return x.SetValue(any)
}
