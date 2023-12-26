package smErrors

import "fmt"

/*
linear_coeffs.go
Description:
	Defines the LinearCoeffs custom error which is used
	when a method attempts to find the linear coefficients of a
	constant (i.e., an object that has no variables).
*/

type CanNotGetLinearCoeffOfConstantError struct {
	Expression interface{}
}

func (lce CanNotGetLinearCoeffOfConstantError) Error() string {
	return fmt.Sprintf(
		"linear coefficients error: cannot find linear coefficients of object %v (type %T) which represents a constant (i.e., this is a polynomial that equals a constant like 3.14)!",
		lce.Expression,
		lce.Expression,
	)
}
