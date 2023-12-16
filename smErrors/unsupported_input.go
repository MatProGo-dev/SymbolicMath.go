package smErrors

import "fmt"

/*
unsupported_input.go
Description:
	This file defines an error for when a function receives an unsupported input.
*/

// Error Definition
type UnsupportedInputError struct {
	FunctionName string
	Input        interface{}
}

func (uie UnsupportedInputError) Error() string {
	return fmt.Sprintf(
		"unsupported input error: %v does not support input of type %T",
		uie.FunctionName,
		uie.Input,
	)
}
