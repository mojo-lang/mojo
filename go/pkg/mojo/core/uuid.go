package core

import (
	"fmt"

	"github.com/google/uuid"
)

const UuidTypeName = "Uuid"
const UuidTypeFullName = "mojo.core.Uuid"

func NewUuid() *Uuid {
	id := uuid.New()
	return &Uuid{Val: append([]byte(nil), id[:]...)}
}

func (x *Uuid) ToUUID() uuid.UUID {
	if x != nil && len(x.Val) > 0 {
		id, err := uuid.FromBytes(x.Val)
		if err != nil {
			return uuid.New()
		}
		return id
	}
	return uuid.New()
}

// // MarshalText implements encoding.TextMarshaler.
// func (x *Uuid) MarshalText() ([]byte, error) {
//	var js [36]byte
//	encodeHex(js[:], uuid)
//	return js[:], nil
// }
//
// // UnmarshalText implements encoding.TextUnmarshaler.
// func (x *Uuid UnmarshalText(data []byte) error {
//	id, err := ParseBytes(data)
//	if err != nil {
//		return err
//	}
//	*uuid = id
//	return nil
// }

// MarshalBinary implements encoding.BinaryMarshaler.
func (x *Uuid) MarshalBinary() ([]byte, error) {
	return x.Val[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (x *Uuid) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	x.Val = append(x.Val[:0], data...)
	return nil
}
