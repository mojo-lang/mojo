package document

import "github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"

func NewAttribute() *Attribute {
	return &Attribute{
		Properties: make(map[string]*core.Value),
	}
}

func (x *Attribute) HasProperty(key string) bool {
	return x != nil && x.Properties != nil && x.Properties[key] != nil
}

func (x *Attribute) GetStringProperty(key string) (string, error) {
	if x != nil && x.Properties != nil {
		if v, ok := x.Properties[key]; ok {
			return v.GetString(), nil
		}
	}
	return "", core.NewNotFoundError("get property %s", key)
}

func (x *Attribute) GetIntProperty(key string) (int64, error) {
	if x != nil && x.Properties != nil {
		if v, ok := x.Properties[key]; ok {
			return v.GetInt64(), nil
		}
	}
	return 0, core.NewNotFoundError("get property %s", key)
}

func (x *Attribute) GetBoolProperty(key string) (bool, error) {
	if x != nil && x.Properties != nil {
		if v, ok := x.Properties[key]; ok {
			return v.GetBool(), nil
		}
	}
	return false, core.NewNotFoundError("get property %s", key)
}

func (x *Attribute) SetStringProperty(key string, value string) *Attribute {
	if x != nil {
		if x.Properties == nil {
			x.Properties = make(map[string]*core.Value)
		}
		x.Properties[key] = core.NewStringValue(value)
	}
	return x
}

func (x *Attribute) SetIntProperty(key string, value int64) *Attribute {
	if x != nil {
		if x.Properties == nil {
			x.Properties = make(map[string]*core.Value)
		}
		x.Properties[key] = core.NewInt64Value(value)
	}
	return x
}

func (x *Attribute) SetBoolProperty(key string, value bool) *Attribute {
	if x != nil {
		if x.Properties == nil {
			x.Properties = make(map[string]*core.Value)
		}
		x.Properties[key] = core.NewBoolValue(value)
	}
	return x
}
