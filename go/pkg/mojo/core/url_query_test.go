package core

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUrl_Query_Unmarshal_String_Slice_1(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"bar", "baz"}},
	}}

	var strs []string
	err := query.Unmarshal("foo", &strs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(strs))
	assert.Equal(t, "bar", strs[0])
}

func TestUrl_Query_Unmarshal_String_Slice_2(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"bar,baz"}},
	}}

	var strs []string
	err := query.Unmarshal("foo", &strs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(strs))
	assert.Equal(t, "bar", strs[0])
}

func TestUrl_Query_Unmarshal_String_Slice_3(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`"bar","baz"`}},
	}}

	var strs []string
	err := query.Unmarshal("foo", &strs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(strs))
	assert.Equal(t, "bar", strs[0])
}

func TestUrl_Query_Unmarshal_String_Slice_4(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`"bar,baz`}},
	}}

	var strs []string
	err := query.Unmarshal("foo", &strs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(strs))
	assert.Equal(t, "\"bar", strs[0])
	assert.Equal(t, "baz", strs[1])
}

func TestUrl_Query_Unmarshal_String_Slice_5(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`["bar","baz"]`}},
	}}

	var strs []string
	err := query.Unmarshal("foo", &strs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(strs))
	assert.Equal(t, "bar", strs[0])
}

func TestUrl_Query_Unmarshal_Integer_Slice_1(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123", "234"}},
	}}

	var vals []int32
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(123), vals[0])
}

func TestUrl_Query_Unmarshal_Integer_Slice_2(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123,234"}},
	}}

	var vals []int32
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(123), vals[0])
}

func TestUrl_Query_Unmarshal_Int32Value_Slice_1(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123", "234"}},
	}}

	var vals []*Int32Value
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(123), vals[0].Val)
}

func TestUrl_Query_Unmarshal_Int32Value_Slice_2(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123,234"}},
	}}

	var vals []*Int32Value
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(123), vals[0].Val)
}

func TestUrl_Query_Unmarshal_Version_Slice_1(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"1.2.3", "2.34.0"}},
	}}

	var vals []*Version
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, uint64(1), vals[0].Major)
}

func TestUrl_Query_Unmarshal_Version_Slice_2(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"1.2.3,2.34.0"}},
	}}

	var vals []*Version
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, uint64(1), vals[0].Major)
}

func TestUrl_Query_Unmarshal_Object_Slice_1(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`{"code":"200","message":"OK"}`, `{"code":"400","message":"invalid parameter"}`}},
	}}

	var vals []*Error
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(200), vals[0].Code.Code)
}

func TestUrl_Query_Unmarshal_Object_Slice_2(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`{"code":"200","message":"OK"},{"code":"400","message":"invalid parameter"}`}},
	}}

	var vals []*Error
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(200), vals[0].Code.Code)
}

func TestUrl_Query_Unmarshal_Object_Slice_3(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`[{"code":"200","message":"OK"},{"code":"400","message":"invalid parameter"}]`}},
	}}

	var vals []*Error
	err := query.Unmarshal("foo", &vals)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, int32(200), vals[0].Code.Code)
}

func TestUrl_Query_Unmarshal_String(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"bar"}},
	}}

	var val string
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, "bar", val)
}

func TestUrl_Query_Unmarshal_Integer(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123"}},
	}}

	var val int32
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, int32(123), val)
}

func TestUrl_Query_Unmarshal_Int32Value(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"123"}},
	}}

	var val Int32Value
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, int32(123), val.Val)
}

func TestUrl_Query_Unmarshal_Version(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"1.2.3"}},
	}}

	var val Version
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val.Major)
}

func TestUrl_Query_Unmarshal_Url(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{"https://localhost:9090/path/to/foo"}},
	}}

	var val Url
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, "https", val.Scheme)
	assert.Equal(t, "/path/to/foo", val.Path)
}

func TestUrl_Query_Unmarshal_Object(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": {Vals: []string{`{"code":"200","message":"OK"}`}},
	}}

	var val Error
	err := query.Unmarshal("foo", &val)
	assert.NoError(t, err)
	assert.Equal(t, int32(200), val.Code.Code)
}

func TestUrl_Query_splitQuotedString(t *testing.T) {
	str := `x,y,z`
	vals := splitQuotedString(str)
	assert.Equal(t, 3, len(vals))
	assert.Equal(t, "x", vals[0])

	str = `"x"`
	vals = splitQuotedString(str)
	assert.Equal(t, 1, len(vals))
	assert.Equal(t, `"x"`, vals[0])

	str = `"x,z" ,  "y"`
	vals = splitQuotedString(str)
	assert.Equal(t, 2, len(vals))
	assert.Equal(t, `"x,z"`, vals[0])

	str = `"x,z" ,  "y", "z,x"`
	vals = splitQuotedString(str)
	assert.Equal(t, 3, len(vals))
	assert.Equal(t, `"x,z"`, vals[0])
	assert.Equal(t, `"y"`, vals[1])
	assert.Equal(t, `"z,x"`, vals[2])
}

func TestUrl_Query_Format(t *testing.T) {
	query := &Url_Query{Vals: map[string]*StringValues{
		"foo": NewStringValues("bar1", "bar2"),
		"zar": NewStringValues("bar"),
	}}

	str := query.Format()
	assert.True(t, "foo=bar1&foo=bar2&zar=bar" == str || "zar=bar&foo=bar1&foo=bar2" == str)
}

func TestNewUrlQuery(t *testing.T) {
	query := NewUrlQuery("foo", "bar1", "foo", "bar2", "zar", 100)
	assert.Equal(t, []string{"bar1", "bar2"}, query.GetValues("foo"))
	assert.Equal(t, "100", query.Get("zar"))
}

func TestNewUrlQuery2(t *testing.T) {
	query := NewUrlQuery("foo", true, "foo2", 13.1, "zar", 100)
	assert.Equal(t, "true", query.Get("foo"))
	assert.Equal(t, "13.1", query.Get("foo2"))
	assert.Equal(t, "100", query.Get("zar"))
}

func TestNewUrlQuery_Add(t *testing.T) {
	query := NewUrlQuery()
	query.Add("foo", true)
	assert.Equal(t, "true", query.Get("foo"))
	assert.Equal(t, true, query.GetBool("foo"))
}

func TestUrlQueryMutationAndConversion(t *testing.T) {
	var query Url_Query
	require.NotPanics(t, func() { query.Set("tag", "first").Add("tag", []string{"second", "third"}).Set("enabled", true) })
	assert.Equal(t, []string{"first", "second", "third"}, query.GetValues("tag"))
	assert.Equal(t, "first", query.Get("tag"))
	assert.True(t, query.GetBool("enabled"))
	assert.True(t, query.Has("tag"))
	assert.ElementsMatch(t, []string{"tag", "enabled"}, query.GetKeys())
	query.Set("tag", "replacement")
	assert.Equal(t, []string{"replacement"}, query.GetValues("tag"))
	query.Del("tag").Del("missing")
	assert.False(t, query.Has("tag"))
	assert.Empty(t, query.Get("tag"))
	assert.Nil(t, query.GetValues("tag"))
	source := url.Values{"q": {"a b", "c&d"}, "empty": {}}
	copied := NewUrlQueryFrom(source)
	assert.Equal(t, source, copied.ToUrlValues())
	var parsed Url_Query
	require.NoError(t, parsed.Parse("q=a+b&q=c%26d"))
	assert.Equal(t, []string{"a b", "c&d"}, parsed.GetValues("q"))
	assert.Equal(t, "q=a+b&q=c%26d", parsed.Format())
	assert.Error(t, parsed.Parse("q=%zz"))
	var added Url_Query
	require.NotPanics(t, func() { added.Add("key", 1) })
	assert.Equal(t, "1", added.Get("key"))
}

func TestUrlQueryNilAndMissingValues(t *testing.T) {
	var absent *Url_Query
	assert.False(t, absent.Has("x"))
	assert.False(t, absent.GetBool("x"))
	assert.Empty(t, absent.Get("x"))
	assert.Nil(t, absent.GetValues("x"))
	assert.Nil(t, absent.GetKeys())
	assert.Nil(t, absent.ToUrlValues())
	assert.Nil(t, absent.FromUrlValues(url.Values{"x": {"y"}}))
	assert.Nil(t, absent.Set("x", "y"))
	assert.Nil(t, absent.Add("x", "y"))
	assert.Nil(t, absent.Del("x"))
	assert.Empty(t, absent.Format())
	assert.NoError(t, absent.Parse("x=y"))
	for _, query := range []*Url_Query{nil, {}, {Vals: map[string]*StringValues{"nil": nil, "empty": {}}}} {
		for _, key := range []string{"missing", "nil", "empty"} {
			value := "unchanged"
			require.NoError(t, query.Unmarshal(key, &value))
			assert.Equal(t, "unchanged", value)
			assert.Empty(t, query.Get(key))
		}
	}
	query := &Url_Query{Vals: map[string]*StringValues{"nil": nil}}
	assert.NotPanics(t, func() { assert.Empty(t, query.GetValues("nil")); assert.Empty(t, query.ToUrlValues()["nil"]) })
}

func TestUrlQueryUnmarshalPreservesValues(t *testing.T) {
	query := NewUrlQuery("tag", []string{"alpha", "beta"})
	for range 2 {
		var decoded []string
		require.NoError(t, query.Unmarshal("tag", &decoded))
		assert.Equal(t, []string{"alpha", "beta"}, decoded)
		assert.Equal(t, []string{"alpha", "beta"}, query.GetValues("tag"), "reading must not add quotes to stored values")
	}
	var number int
	assert.Error(t, NewUrlQuery("n", "invalid").Unmarshal("n", &number))
	var numbers []int
	assert.Error(t, UnmarshalParam("1,invalid", &numbers))
	assert.NoError(t, UnmarshalParam("", &number))
	assert.NoError(t, UnmarshalParam("   ", &numbers))
	assert.NotPanics(t, func() { assert.Error(t, UnmarshalParam("x", nil)) })
	assert.NotPanics(t, func() { assert.Error(t, UnmarshalParam("x", "not a pointer")) })
}
