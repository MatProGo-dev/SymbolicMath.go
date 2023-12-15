package smErrors

import "fmt"

/*
out_of_bounds_error.go
Description:
	Defining the out of bounds error object and all of its associated functions.
*/

type OutOfBoundsError struct {
	Index      int
	Expression MatrixLike
}

func (obe OutOfBoundsError) Error() string {
	errOut := fmt.Sprintf(
		"out of bounds error: index %v is out of bounds for %T object of dimension %v",
		obe.Index,
		obe.Expression,
		obe.Expression.Dims(),
	)
	return errOut
}

// Explicit Functions
func CheckIndexOnVector(index int, vector MatrixLike) error {
	if index < 0 || index >= vector.Dims()[0] {
		return OutOfBoundsError{
			Index:      index,
			Expression: vector,
		}
	}
	return nil
}
