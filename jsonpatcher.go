package jsonpatcher

import (
	"encoding/json"
)

var (
	PATCH_OPERATIONS = [...]string{"add", "remove", "replace", "move", "copy", "test"}
)

type Patch struct {
	Op    string       `json:"op"`
	Path  string       `json:"path"`
	Value *interface{} `json:"value,omitempty"`
}

type JsonPatch []Patch

func CheckValidOperation(op string) bool {

	for _, standard_op := range PATCH_OPERATIONS {
		if op == standard_op {
			return true
		}
	}
	return false
}

func PatchJson(jsonBytes []byte) (JsonPatch, error) {

	var patchData JsonPatch

	err := json.Unmarshal(jsonBytes, &patchData)
	if err != nil {
		return JsonPatch{}, err
	}

	for _, patch := range patchData {

		if !CheckValidOperation(patch.Op) {
			return JsonPatch{}, &InvalidOperation{patch.Op}
		}
	}

	return patchData, nil

}
