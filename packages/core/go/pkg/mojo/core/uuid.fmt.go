package core

import "github.com/google/uuid"

func ParseUuid(value string) (*Uuid, error) {
	id := &Uuid{}
	err := id.Parse(value)
	return id, err
}

func (x *Uuid) Format() string {
	return x.ToUUID().String()
}

func (x *Uuid) Parse(value string) error {
	id, err := uuid.Parse(value)
	if err == nil {
		x.Val = append(x.Val, id[:]...)
	}
	return err
}
