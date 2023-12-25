package smErrors

import "fmt"

/*
linear_coeffs.go
Description:
	Defines the LinearCoeffs custom error which is used
	when a method attempts to find the linear coefficients of a
	constant (i.e., an object that has no variables).
*/

type LinearCoeffsError struct {
	Expression interface{}
}

func (lce LinearCoeffsError) Error() string {
	return fmt.Sprintf(
		"linear coefficients error: cannot find linear coefficients of object %v (type %T)",
		lce.Expression,
		lce.Expression,
	)
}
