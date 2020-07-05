package jsonpatcher

import "fmt"

// InvalidOperation returned when the operation is not a standard patch operation
type InvalidOperation struct {
	Op string
}

func (in *InvalidOperation) Error() string {
	return fmt.Sprintf("found invalid operation '%s' in the parch request", in.Op)
}

// EmptyPath returned when the path in a patch operation is black
type EmptyPath struct{}

func (ep *EmptyPath) Error() string {
	return "path cannot be empty"
}
