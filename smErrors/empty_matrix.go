package smErrors

import "fmt"

/*
empty_matrix.go
Description:

	Defining the empty matrix error object and all of its associated functions.
*/

// Type
type EmptyMatrixError struct {
	Expression MatrixLike
}

// Error Function
func (eme EmptyMatrixError) Error() string {
	errOut := fmt.Sprintf(
		"empty matrix error: the matrix of type %T is empty",
		eme.Expression,
	)
	return errOut
}
