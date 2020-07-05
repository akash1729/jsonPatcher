package jsonpatcher

import (
	"encoding/json"
	"strings"
)

var (
	patchOperations = [...]string{"add", "remove", "replace", "move", "copy", "test"}
)

// Patch stuct to store each patch operation
// refer to RFC 5789 and RFC 6902
// https://tools.ietf.org/html/rfc5789, https://tools.ietf.org/html/rfc6902
type Patch struct {
	Op    string       `json:"op"`
	Path  string       `json:"path"`
	Value *interface{} `json:"value,omitempty"`
}

// JSONPatch list of Patch objects
type JSONPatch []Patch

func checkValidOperation(op string) bool {

	for _, standardOp := range patchOperations {
		if op == standardOp {
			return true
		}
	}
	return false
}

// ParseJSON pass json bytes to return a JSONPatch object
func ParseJSON(jsonBytes []byte) (JSONPatch, error) {

	var patchData JSONPatch

	err := json.Unmarshal(jsonBytes, &patchData)
	if err != nil {
		return JSONPatch{}, err
	}

	for _, patch := range patchData {

		if !checkValidOperation(patch.Op) {
			return JSONPatch{}, &InvalidOperation{patch.Op}
		}

		if strings.TrimSpace(patch.Path) == "" {
			return JSONPatch{}, &EmptyPath{}
		}
	}

	return patchData, nil

}
