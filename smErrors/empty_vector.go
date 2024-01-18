package smErrors

import "fmt"

/*
empty_vector.go
Description:

	Defining the empty vector error object and all of its associated functions.
*/

// Type
type EmptyVectorError struct {
	Expression MatrixLike
}

// Error Function
func (eve EmptyVectorError) Error() string {
	errOut := fmt.Sprintf(
		"empty vector error: the vector of type %T is empty",
		eve.Expression,
	)
	return errOut
}
