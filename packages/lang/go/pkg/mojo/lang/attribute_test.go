package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute_GetArrayLiteralArgument(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{Value: NewStringLiteralExpressionFrom("foo1")},
			{Value: NewStringLiteralExpressionFrom("foo2")},
			{Value: NewStringLiteralExpressionFrom("foo3")},
		},
	}

	array := attribute.GetArrayLiteralArgument()
	assert.NotNil(t, array)
	assert.Equal(t, 3, len(array.Elements))
	assert.Equal(t, "foo1", array.Elements[0].GetStringLiteralExpr().Value)
}

func TestAttribute_GetArrayLiteralArgument2(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{Value: NewArrayLiteralExpression(&ArrayLiteralExpr{Elements: []*Expression{
				NewStringLiteralExpressionFrom("foo1"),
				NewStringLiteralExpressionFrom("foo2"),
				NewStringLiteralExpressionFrom("foo3"),
			}})},
		},
	}

	array := attribute.GetArrayLiteralArgument()
	assert.NotNil(t, array)
	assert.Equal(t, 3, len(array.Elements))
	assert.Equal(t, "foo1", array.Elements[0].GetStringLiteralExpr().Value)
}

func TestAttribute_GetObjectLiteralArgument(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{
				Label: "foo1",
				Value: NewStringLiteralExpressionFrom("foo1"),
			},
			{
				Label: "foo2",
				Value: NewStringLiteralExpressionFrom("foo2"),
			},
			{
				Label: "foo2",
				Value: NewStringLiteralExpressionFrom("foo3"),
			},
		},
	}

	object := attribute.GetObjectLiteralArgument()
	assert.NotNil(t, object)
	assert.Equal(t, 3, len(object.Fields))
	assert.Equal(t, "foo1", object.Fields[0].Name)
	assert.Equal(t, "foo1", object.Fields[0].Value.GetStringLiteralExpr().Value)
}

func TestAttribute_GetObjectLiteralArgument2(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{
				Value: NewObjectLiteralExpression(&ObjectLiteralExpr{
					Fields: []*ObjectLiteralExpr_Field{
						{
							Name:  "foo1",
							Value: NewStringLiteralExpressionFrom("foo1"),
						},
						{
							Name:  "foo2",
							Value: NewStringLiteralExpressionFrom("foo2"),
						},
						{
							Name:  "foo2",
							Value: NewStringLiteralExpressionFrom("foo3"),
						},
					},
				}),
			},
		},
	}

	object := attribute.GetObjectLiteralArgument()
	assert.NotNil(t, object)
	assert.Equal(t, 3, len(object.Fields))
	assert.Equal(t, "foo1", object.Fields[0].Name)
	assert.Equal(t, "foo1", object.Fields[0].Value.GetStringLiteralExpr().Value)
}

func TestAttribute_GetMapLiteralArgument(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{
				Label: "foo1",
				Value: NewStringLiteralExpressionFrom("foo1"),
			},
			{
				Label: "foo2",
				Value: NewStringLiteralExpressionFrom("foo2"),
			},
			{
				Label: "foo2",
				Value: NewStringLiteralExpressionFrom("foo3"),
			},
		},
	}

	object := attribute.GetMapLiteralArgument()
	assert.NotNil(t, object)
	assert.Equal(t, 3, len(object.Entries))
	assert.Equal(t, "foo1", object.Entries[0].Key)
	assert.Equal(t, "foo1", object.Entries[0].Value.GetStringLiteralExpr().Value)
}

func TestAttribute_GetMapLiteralArgument2(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{
				Value: NewMapLiteralExpression(&MapLiteralExpr{
					Entries: []*MapLiteralExpr_Entry{
						{
							Key:   "foo1",
							Value: NewStringLiteralExpressionFrom("foo1"),
						},
						{
							Key:   "foo2",
							Value: NewStringLiteralExpressionFrom("foo2"),
						},
						{
							Key:   "foo2",
							Value: NewStringLiteralExpressionFrom("foo3"),
						},
					},
				}),
			},
		},
	}

	object := attribute.GetMapLiteralArgument()
	assert.NotNil(t, object)
	assert.Equal(t, 3, len(object.Entries))
	assert.Equal(t, "foo1", object.Entries[0].Key)
	assert.Equal(t, "foo1", object.Entries[0].Value.GetStringLiteralExpr().Value)
}

func TestAttribute_GetMapLiteralArgument3(t *testing.T) {
	attribute := &Attribute{
		Name: "test",
		Arguments: []*Argument{
			{
				Value: NewObjectLiteralExpression(&ObjectLiteralExpr{
					Fields: []*ObjectLiteralExpr_Field{
						{
							Name:  "foo1",
							Value: NewStringLiteralExpressionFrom("foo1"),
						},
						{
							Name:  "foo2",
							Value: NewStringLiteralExpressionFrom("foo2"),
						},
						{
							Name:  "foo2",
							Value: NewStringLiteralExpressionFrom("foo3"),
						},
					},
				}),
			},
		},
	}

	object := attribute.GetMapLiteralArgument()
	assert.NotNil(t, object)
	assert.Equal(t, 3, len(object.Entries))
	assert.Equal(t, "foo1", object.Entries[0].Key)
	assert.Equal(t, "foo1", object.Entries[0].Value.GetStringLiteralExpr().Value)
}

func TestAttribute_IsSameName(t *testing.T) {
	attribute := &Attribute{Name: "foo", PackageName: "mojo.core"}
	assert.True(t, attribute.IsSameName("foo"))
	assert.True(t, attribute.IsSameName("mojo.core.foo"))

	attribute = &Attribute{Name: "foo", PackageName: ""}
	assert.True(t, attribute.IsSameName("foo"))
	assert.True(t, attribute.IsSameName("mojo.core.foo"))

	attribute = &Attribute{Name: "foo", PackageName: "mojo.http"}
	assert.True(t, attribute.IsSameName("foo"))
	assert.True(t, attribute.IsSameName("http.foo"))
	assert.True(t, attribute.IsSameName("mojo.http.foo"))
}
