package converter

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"
	"github.com/stretchr/testify/require"
)

func TestEnumNumbersAndRegistration(t *testing.T) {
	for _, register := range []bool{false, true} {
		ctx, file := descriptorContext()
		ctx = context.WithValues(ctx, "register_enum", register)
		decl := &lang.EnumDecl{Name: "OrderStatus", Type: &lang.EnumType{Enumerators: []*lang.ValueDecl{
			{Name: "Unknown"}, {Name: "Pending"}, {Name: "Done", Attributes: lang.SetIntegerAttribute(nil, core.NumberAttributeName, 10)},
		}}}
		enum := descriptor.NewEnum(file)
		require.NoError(t, (Enum{}).ConvertTo(ctx, decl, enum))
		require.Equal(t, "OrderStatus", enum.GetName())
		require.Len(t, enum.Proto.Value, 3)
		for i, want := range []struct {
			name   string
			number int32
		}{{"ORDER_STATUS_UNKNOWN", 0}, {"ORDER_STATUS_PENDING", 1}, {"ORDER_STATUS_DONE", 10}} {
			require.Equal(t, want.name, enum.Proto.Value[i].GetName())
			require.Equal(t, want.number, enum.Proto.Value[i].GetNumber())
		}
		if register {
			require.Equal(t, []*descriptor.Enum{enum}, file.Enums)
		} else {
			require.Empty(t, file.Enums)
		}
	}
}

func TestEnumRejectsNegativeNumber(t *testing.T) {
	ctx, file := descriptorContext()
	decl := &lang.EnumDecl{Name: "Status", Type: &lang.EnumType{Enumerators: []*lang.ValueDecl{{Name: "Bad", Attributes: lang.SetIntegerAttribute(nil, core.NumberAttributeName, -1)}}}}
	err := (Enum{}).ConvertTo(ctx, decl, descriptor.NewEnum(file))
	require.ErrorContains(t, err, "must be positive")
	require.Empty(t, file.Enums)
}
