package openapi

import (
	"reflect"
	"sort"
	"strings"
)

func (x *PathItem) GenerateExample(components *Components) {
	operations := x.GetOperations()
	for _, op := range operations {
		op.Operation.GenerateExample(components)
	}
}

func (x *PathItem) SupplementExample(components *Components) {
	operations := x.GetOperations()
	for _, op := range operations {
		op.Operation.SupplementExample(components)
	}
}

func (x *PathItem) GetOperations() []struct {
	Operation *Operation
	Method    string
} {

	var operations []struct {
		Operation *Operation
		Method    string
	}
	if x != nil {
		for _, method := range []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch", "Trace"} {
			if operation, ok := reflect.ValueOf(x).Elem().FieldByName(method).Interface().(*Operation); ok && operation != nil {
				operations = append(operations, struct {
					Operation *Operation
					Method    string
				}{Operation: operation, Method: strings.ToUpper(method)})
			}
		}
	}
	return operations
}

func (x *PathItem) SetOperation(name string, operation *Operation) {
	if x != nil {
		switch strings.ToLower(name) {
		case "get":
			x.Get = operation
		case "post":
			x.Post = operation
		case "put":
			x.Put = operation
		case "patch":
			x.Patch = operation
		case "delete":
			x.Delete = operation
		case "options":
			x.Options = operation
		case "head":
			x.Head = operation
		case "trace":
			x.Trace = operation
		}
	}
}

func (x *PathItem) Merge(item *PathItem) *PathItem {
	if x == nil {
		return item
	}

	if len(item.Ref) > 0 {
		x.Ref = item.Ref
	}
	if len(item.Summary) > 0 {
		x.Summary = item.Summary
	}
	if item.Description != nil {
		x.Description = item.Description
	}
	if item.Get != nil {
		x.Get = item.Get
	}
	if item.Put != nil {
		x.Put = item.Put
	}
	if item.Post != nil {
		x.Post = item.Post
	}
	if item.Delete != nil {
		x.Delete = item.Delete
	}
	if item.Options != nil {
		x.Options = item.Options
	}
	if item.Head != nil {
		x.Head = item.Head
	}
	if item.Patch != nil {
		x.Patch = item.Patch
	}
	if item.Trace != nil {
		x.Trace = item.Trace
	}
	if len(item.Servers) > 0 {
		x.Servers = append(x.Servers, item.Servers...)
	}
	if len(item.Parameters) > 0 {
		x.Parameters = append(x.Parameters, item.Parameters...)
	}
	return x
}

func OrderedPathItemIterator(m map[string]*PathItem, onKey func(key string, value *PathItem) error) error {
	var list []struct {
		key   string
		value *PathItem
	}
	for k, v := range m {
		list = append(list, struct {
			key   string
			value *PathItem
		}{
			key:   k,
			value: v,
		})
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].key < list[j].key
	})
	for _, el := range list {
		err := onKey(el.key, el.value)
		if err != nil {
			return err
		}
	}
	return nil
}
