package smErrors

import "fmt"

/*
VectorDimensionError
Description:

	This error is thrown when two matrices do not have the appropriate dimensions
	for a given operation.
*/
type VectorDimensionError struct {
	Arg1      VectorLike
	Arg2      VectorLike
	Operation string
}

func (de VectorDimensionError) Error() string {
	return fmt.Sprintf(
		"vector dimension error: Cannot perform %v between expression of dimension %v and expression of dimension %v",
		de.Operation,
		de.Arg1.Len(),
		de.Arg2.Len(),
	)
}
