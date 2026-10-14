package lang

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Check the handwritten JSON codecs against every variant declared by the AST
// schema, including zero-valued nodes whose type must survive serialization.
func TestASTJSONPreservesUnionVariants(t *testing.T) {
	for _, union := range []proto.Message{&Expression{}, &Declaration{}, &Statement{}} {
		message := union.ProtoReflect()
		fields := message.Descriptor().Oneofs().Get(0).Fields()
		for i := 0; i < fields.Len(); i++ {
			field := fields.Get(i)
			kind := string(field.Message().Name())
			if kind == "Expression" || kind == "Declaration" {
				continue
			}
			for _, positioned := range []bool{false, true} {
				name := string(message.Descriptor().Name()) + "/" + kind
				if positioned {
					name += "/positioned"
				} else {
					name += "/zero"
				}
				t.Run(name, func(t *testing.T) {
					original := message.New()
					value := original.NewField(field)
					if positioned {
						node := value.Message()
						position := node.Descriptor().Fields().ByName("start_position")
						require.NotNil(t, position)
						node.Set(position, protoreflect.ValueOfMessage((&Position{Filename: "sample.mojo", Line: 2, Column: 3}).ProtoReflect()))
					}
					original.Set(field, value)
					require.NotPanics(t, func() {
						encoded, err := jsoniter.Marshal(original.Interface())
						require.NoError(t, err)
						require.Equal(t, kind, jsoniter.Get(encoded, "@type").ToString())
						decoded := message.New().Interface()
						require.NoError(t, jsoniter.Unmarshal(encoded, decoded))
						require.True(t, proto.Equal(original.Interface(), decoded), "AST node changed after round trip: %s", encoded)
					})
				})
			}
		}
	}
}

func TestASTJSONNestedValuesAndStatementExpressions(t *testing.T) {
	expressions := []*Expression{
		NewBoolLiteralExpressionFrom(false), NewBoolLiteralExpressionFrom(true),
		NewIntegerLiteralExpressionFrom(-42), NewFloatLiteralExpressionFrom(1.25), NewStringLiteralExpressionFrom("hello"),
		NewArrayLiteralExpressionFrom(NewStringLiteralExpressionFrom("one"), NewIntegerLiteralExpressionFrom(2)),
		NewMapLiteralExpression(&MapLiteralExpr{Entries: []*MapLiteralExpr_Entry{{Key: "enabled", Value: NewBoolLiteralExpressionFrom(true)}}}),
	}
	for _, expression := range expressions {
		t.Run(string(expression.ProtoReflect().WhichOneof(expression.ProtoReflect().Descriptor().Oneofs().Get(0)).Name()), func(t *testing.T) {
			original := NewExpressionStatement(expression)
			encoded, err := jsoniter.Marshal(original)
			require.NoError(t, err)
			var restored Statement
			require.NoError(t, jsoniter.Unmarshal(encoded, &restored))
			require.True(t, proto.Equal(original, &restored), "nested payload changed: %s", encoded)
		})
	}
	for _, union := range []proto.Message{&Expression{}, &Declaration{}, &Statement{}} {
		encoded, err := jsoniter.Marshal(union)
		require.NoError(t, err)
		require.JSONEq(t, "null", string(encoded))
		for _, input := range []string{"null", `{}`, `{"@type":"UnknownFutureNode"}`} {
			decoded := union.ProtoReflect().New().Interface()
			require.NoError(t, jsoniter.Unmarshal([]byte(input), decoded))
			require.True(t, proto.Equal(union, decoded))
		}
	}
}

func TestASTJSONNestedStatementsAndDeclarations(t *testing.T) {
	for _, original := range []*Statement{
		{Statement: &Statement_BlockStmt{BlockStmt: &BlockStmt{Statements: []*Statement{NewExpressionStatement(NewBoolLiteralExpressionFrom(true)), NewExpressionStatement(NewIntegerLiteralExpressionFrom(2))}}}},
		NewStructDeclStatement(&StructDecl{Name: "Item", Attributes: []*Attribute{NewStringAttribute("", "alias", "item")}, Type: &StructType{Fields: []*ValueDecl{{Name: "id"}, {Name: "name"}}}}),
	} {
		encoded, err := jsoniter.Marshal(original)
		require.NoError(t, err)
		var restored Statement
		require.NoError(t, jsoniter.Unmarshal(encoded, &restored))
		require.True(t, proto.Equal(original, &restored), "nested AST changed: %s", encoded)
	}
	for _, original := range []*Expression{NewNullLiteralExpressionFrom(), NewArrayLiteralExpression(&ArrayLiteralExpr{Elements: []*Expression{}})} {
		encoded, err := jsoniter.Marshal(original)
		require.NoError(t, err)
		require.NotEmpty(t, jsoniter.Get(encoded, "@type").ToString())
		require.NotContains(t, string(encoded), "elements")
		var restored Expression
		require.NoError(t, jsoniter.Unmarshal(encoded, &restored))
		require.True(t, proto.Equal(original, &restored))
	}
}
