package main

import "fmt"

type InvalidOperation struct {
	Op string
}

func (in *InvalidOperation) Error() string {
	return fmt.Sprintf("found invalid operation '%s' in the parch request", in.Op)
}
