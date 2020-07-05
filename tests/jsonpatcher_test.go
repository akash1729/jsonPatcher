package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akash1729/jsonpatcher"
)

func TestCorrect(t *testing.T) {

	raw := []byte(`[
		{ "op": "replace", "path": "/baz", "value": "boo" },
		{ "op": "add", "path": "/hello", "value": ["world"] },
		{ "op": "remove", "path": "/foo" }
		]`)

	patches, err := jsonpatcher.ParseJSON(raw)
	if err != nil {
		panic(err)
	}

	// first one
	first := patches[0]
	require.Equal(t, "replace", first.Op)
	require.Equal(t, "/baz", first.Path)
	require.Equal(t, "boo", (*first.Value).(string))

	// second
	second := patches[1]
	require.Equal(t, "add", second.Op)
	require.Equal(t, "/hello", second.Path)
	require.Equal(t, []interface{}{"world"}, (*second.Value).([]interface{}))

	// third
	third := patches[2]
	require.Equal(t, "remove", third.Op)
	require.Equal(t, "/foo", third.Path)
	require.Nil(t, third.Value)
}

func TestInvalidOperation(t *testing.T) {

	raw := []byte(`[
		{ "op": "replace123", "path": "/baz", "value": "boo" }]`)

	_, err := jsonpatcher.ParseJSON(raw)

	require.IsType(t, &jsonpatcher.InvalidOperation{}, err)

}

func TestEmptyPath(t *testing.T) {

	raw := []byte(`[
		{ "op": "replace", "path": "   ", "value": "boo" }]`)

	_, err := jsonpatcher.ParseJSON(raw)

	require.IsType(t, &jsonpatcher.EmptyPath{}, err)

}
