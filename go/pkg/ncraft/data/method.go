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
	// the fields of the Methods.
	Bindings []*HTTPBinding

	Subscriptions []*MessagingSubscription

	Extensions map[string]interface{}
}

func (m *Method) GetRequest() *Message {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Method) GetResponse() *Message {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Method) GetEntity() *Message {
	if m != nil {
		return m.Entity
	}
	return nil
}
