package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akash1729/jsonpatcher"
)

func TestCorrect(t *testing.T) {

	raw := []byte(`[
		{ "op": "replace", "path": "/baz", "value": "boo" },
		{ "op": "adder", "path": "/hello", "value": ["world"] },
		{ "op": "remove", "path": "/foo" }
		]`)

	patches, err := jsonpatcher.PatchJson(raw)
	if err != nil {
		panic(err)
	}

	// first one
	first = patches[0]
	require.Equal("replace", first.Op)
	require.Equal("/baz", first.Path)
	require.Equal("boo", first.value.(string))

}
