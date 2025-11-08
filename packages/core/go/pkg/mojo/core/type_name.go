package core

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

type TypeName struct {
	PackageName       string
	Name              string
	GenericParameters []*TypeName
}

func ParseTypeName(name string) (*TypeName, error) {
	tn := &TypeName{}
	if err := tn.Parse(name); err != nil {
		return nil, err
	}
	return tn, nil
}

func (n *TypeName) GetFullName() string {
	if n != nil {
		if len(n.PackageName) > 0 {
			return n.PackageName + "." + n.Name
		}
		return n.Name
	}
	return ""
}

func (n *TypeName) Parse(name string) error {
	if err := n.parse(bufio.NewReader(strings.NewReader(name))); err != nil {
		return err
	}

	return nil
}

func (n *TypeName) IsSame(tn *TypeName) bool {
	if n != nil {
		if tn != nil {
			if n.PackageName == tn.PackageName && n.Name == tn.Name && len(n.GenericParameters) == len(tn.GenericParameters) {
				for i := 0; i < len(n.GenericParameters); i++ {
					if !n.GenericParameters[i].IsSame(tn.GenericParameters[i]) {
						return false
					}
				}
				return true
			} else {
				return false
			}
		}
	} else if tn != nil {
		return false
	}
	return true
}

func (n *TypeName) parse(name io.ByteScanner) error {
	buff := new(bytes.Buffer)

	state := 0 // 1: package name, 2: name
	for {
		v, err := name.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if buff.Len() > 0 {
					n.Name = buff.String()
				}
				return nil
			}
			return err
		}

		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNum := v >= '0' && v <= '9'
		vIsBlank := v == ' ' || v == '\t'
		if state == 0 {
			if vIsLow {
				state = 1
			} else if vIsCap {
				state = 2
			} else if vIsBlank {
				continue
			}
		}

		if state == 1 {
			if vIsLow || vIsNum || v == '.' || v == '_' {
				if v == '.' {
					if len(n.PackageName) > 0 {
						n.PackageName = n.PackageName + "." + buff.String()
					} else {
						n.PackageName = buff.String()
					}
					buff.Reset()
					state = 0
				} else {
					buff.WriteByte(v)
				}
			} else {
				// error
			}
		} else if state == 2 {
			if v == '<' {
				n.Name = buff.String()
				buff.Reset()

				for {
					tn := &TypeName{}
					if err = tn.parse(name); err != nil {
						return err
					}
					n.GenericParameters = append(n.GenericParameters, tn)

					b, err := name.ReadByte()
					if err != nil {
						return err
					}
					if b == ',' {
						continue
					} else if b == '>' {
						return nil
					}
				}
			} else if vIsBlank {
				if buff.Len() > 0 {
					n.Name = buff.String()
					buff.Reset()
				}
				continue
			} else if v == ',' || v == '>' {
				if buff.Len() > 0 {
					n.Name = buff.String()
					buff.Reset()
					name.UnreadByte()
					return nil
				}
			} else {
				buff.WriteByte(v)
			}
		}
	}
	return nil
}
