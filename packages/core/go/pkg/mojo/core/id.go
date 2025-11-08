package core

const IdTypeName = "Id"
const IdTypeFullName = "mojo.core.Id"

func NewId(value interface{}) *Id {
	switch v := value.(type) {
	case uint64:
		return NewIntId(v)
	case string:
		return NewStringId(v)
	case *Uuid:
		return NewUuidId(v)
	}
	return nil
}

func NewIntId(value uint64) *Id {
	return &Id{Id: &Id_Uint64Val{Uint64Val: value}}
}

func NewStringId(value string) *Id {
	return &Id{Id: &Id_StringVal{StringVal: value}}
}

func NewUuidId(value *Uuid) *Id {
	return &Id{Id: &Id_UuidVal{UuidVal: value}}
}

func UuidId() *Id {
	return NewUuidId(NewUuid())
}

func (x *Id) SetInt(id uint64) {
	x.Id = &Id_Uint64Val{Uint64Val: id}
}

func (x *Id) SetString(id string) {
	x.Id = &Id_StringVal{StringVal: id}
}

func (x *Id) SetUuid(id *Uuid) {
	x.Id = &Id_UuidVal{UuidVal: id}
}

func (x *Id) SetUuidString(id string) {
	uuid, err := ParseUuid(id)
	if err == nil {
		x.SetUuid(uuid)
	}
}

func (x *Id) RegenerateUuid() {
	x.SetUuid(NewUuid())
}
