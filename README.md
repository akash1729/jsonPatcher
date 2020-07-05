# jsonpatcher : go package to parse http PATCH request

## Overview [![GoDoc](https://godoc.org/github.com/akash1729/jsonpatcher?status.svg)](https://godoc.org/github.com/akash1729/jsonpatcher)

package to parse http JSON PATCH request. 

Create a list of PATCH objects to further processing

## Install

```
go get github.com/akash1729/jsonpatcher
```

## Example

```
	rawBytes := []byte(`[
		{ "op": "replace", "path": "/baz", "value": "boo"},
		{ "op": "add", "path": "/hello", "value": ["world"]},
		{ "op": "remove", "path": "/foo"}
		]`)

	patches, err := jsonpatcher.ParseJSON(raw)


```

## Author

akash1729

## License

MIT.

