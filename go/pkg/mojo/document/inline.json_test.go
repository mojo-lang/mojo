package document

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInlineJSONTextRoundTrip(t *testing.T) {
	for _, text := range []string{"", "hello", "你好\n\"world\"\\", "<b>text & more</b>"} {
		t.Run(text, func(t *testing.T) {
			encoded, err := jsoniter.MarshalToString(NewTextInline(text))
			require.NoError(t, err)
			want, err := jsoniter.MarshalToString(text)
			require.NoError(t, err)
			assert.Equal(t, want, encoded, "inline text must retain the configured string escaping")
			var plain string
			require.NoError(t, jsoniter.UnmarshalFromString(encoded, &plain))
			assert.Equal(t, text, plain)
			decoded := NewLinkInline(&Link{})
			require.NoError(t, jsoniter.UnmarshalFromString(encoded, decoded))
			require.NotNil(t, decoded.GetText())
			assert.Equal(t, text, decoded.GetText().Val)
			assert.Nil(t, decoded.GetLink())
		})
	}
}

func TestInlineJSONLinkRoundTrip(t *testing.T) {
	for _, typ := range []string{"Link", "mojo.document.Link"} {
		t.Run(typ, func(t *testing.T) {
			input := `{"@type":"` + typ + `","attribute":{"identifier":"docs"},"description":["read"," more"],"target":{"title":"Guide","url":"https://example.com/docs?q=go"}}`
			var decoded Inline
			require.NoError(t, jsoniter.UnmarshalFromString(input, &decoded))
			link := decoded.GetLink()
			require.NotNil(t, link)
			assert.Equal(t, "docs", link.GetAttribute().GetIdentifier())
			require.Len(t, link.Description, 2)
			assert.Equal(t, "read", link.Description[0].GetText().GetVal())
			assert.Equal(t, " more", link.Description[1].GetText().GetVal())
			assert.Equal(t, "Guide", link.GetTarget().GetTitle())
			assert.Equal(t, "https://example.com/docs?q=go", link.GetTarget().GetUrl().Format())
			encoded, err := jsoniter.MarshalToString(&decoded)
			require.NoError(t, err)
			assert.JSONEq(t, `{"@type":"Link","attribute":{"identifier":"docs"},"description":["read"," more"],"target":{"title":"Guide","url":"https://example.com/docs?q=go"}}`, encoded)
			var roundTrip Inline
			require.NoError(t, jsoniter.UnmarshalFromString(encoded, &roundTrip))
			require.NotNil(t, roundTrip.GetLink())
			assert.Equal(t, "read", roundTrip.GetLink().Description[0].GetText().GetVal())
		})
	}
}

func TestInlineJSONInvalidInputs(t *testing.T) {
	for _, input := range []string{
		``, `"unterminated`, `true`, `false`, `123`, `-1.5`, `[]`, `{}`, `{"@type":123}`, `{"@type":"unknown"}`,
		`{"@type":"Link","target":42}`, `{"@type":"Link","description":[true]}`, `{"@type":"Link"`,
	} {
		t.Run(input, func(t *testing.T) {
			inline := NewTextInline("preserved")
			assert.Error(t, jsoniter.UnmarshalFromString(input, inline))
			assert.Equal(t, "preserved", inline.GetText().GetVal(), "failed decoding must preserve the receiver")
		})
	}
}

func TestInlineJSONEmptyValues(t *testing.T) {
	codec := &InlineCodec{}
	assert.True(t, codec.IsEmpty(nil))
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&Inline{})))
	assert.False(t, codec.IsEmpty(unsafe.Pointer(NewTextInline(""))))
	for _, tt := range []struct {
		name   string
		inline *Inline
	}{
		{"nil receiver", nil},
		{"empty union", &Inline{}},
		{"nil text payload", &Inline{Inline: &Inline_Text{}}},
		{"nil link payload", &Inline{Inline: &Inline_Link{}}},
		{"nil text wrapper", &Inline{Inline: (*Inline_Text)(nil)}},
		{"nil link wrapper", &Inline{Inline: (*Inline_Link)(nil)}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var encoded string
			var err error
			require.NotPanics(t, func() { encoded, err = jsoniter.MarshalToString(tt.inline) })
			require.NoError(t, err)
			assert.Equal(t, "null", encoded)
		})
	}
	inline := NewTextInline("old")
	require.NoError(t, jsoniter.UnmarshalFromString(`null`, inline))
	assert.Nil(t, inline.Inline)
	type payload struct {
		Inline Inline `json:"inline,omitempty"`
	}
	encoded, err := jsoniter.MarshalToString(&payload{})
	require.NoError(t, err)
	assert.Equal(t, `{}`, encoded)
}

func TestInlineJSONNestedValues(t *testing.T) {
	const input = `["first",null,{"@type":"Link","description":["second",{"@type":"Link","description":["nested"]}]}]`
	var inlines []*Inline
	require.NoError(t, jsoniter.UnmarshalFromString(input, &inlines))
	require.Len(t, inlines, 3)
	assert.Equal(t, "first", inlines[0].GetText().GetVal())
	assert.Nil(t, inlines[1])
	description := inlines[2].GetLink().GetDescription()
	require.Len(t, description, 2)
	assert.Equal(t, "second", description[0].GetText().GetVal())
	nested := description[1].GetLink().GetDescription()
	require.Len(t, nested, 1)
	assert.Equal(t, "nested", nested[0].GetText().GetVal())
	encoded, err := jsoniter.MarshalToString(inlines)
	require.NoError(t, err)
	assert.JSONEq(t, input, encoded)
}

func TestInlineJSONUnsupportedEncodingReturnsError(t *testing.T) {
	// The handwritten codec currently supports text and link payloads.
	_, err := jsoniter.MarshalToString(NewCodeInlineFrom("code"))
	assert.Error(t, err, "unsupported variants must not silently produce invalid JSON")
}
