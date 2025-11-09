package core

import (
	"strings"
)

func ParseOrdering(value string) (*Ordering, error) {
	o := &Ordering{}
	if err := o.Parse(value); err != nil {
		return nil, err
	}
	return o, nil
}

func (x *Ordering) Format() string {
	if len(x.Orders) > 0 {
		var values []string
		for _, v := range x.Orders {
			values = append(values, v.Format())
		}
		return strings.Join(values, ", ")
	}
	return ""
}

func (x *Ordering) ToString() string {
	return x.Format()
}

func (x *Ordering) Parse(value string) error {
	if x != nil {
		segments := strings.Split(value, ",")
		for _, segment := range segments {
			v := &Ordering_Order{}
			if err := v.Parse(segment); err != nil {
				return err
			}
			x.Orders = append(x.Orders, v)
		}
	}
	return nil
}

func (x *Ordering_Order) Format() string {
	if x.Sort != Ordering_SORT_UNSPECIFIED {
		return x.Field + " " + x.Sort.Format()
	} else {
		return x.Field
	}
}

func (x *Ordering_Order) Parse(value string) error {
	value = strings.TrimSpace(value)
	if len(value) > 0 {
		pos := strings.Index(value, " ")
		if pos < 0 {
			x.Field = value
		} else {
			x.Field = value[:pos]
			value = strings.TrimSpace(value[pos:])
			if err := x.Sort.Parse(strings.TrimSpace(value)); err != nil {
				return err
			}
		}
	}
	return nil
}
