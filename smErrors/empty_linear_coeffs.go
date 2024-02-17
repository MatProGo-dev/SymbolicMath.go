package smErrors

import "fmt"

/*
empty_linear_coeffs.go
Description:

	Defining the empty linear coefficients error object and all of its associated functions.
*/

// Type
type EmptyLinearCoeffsError struct {
	Expression interface{}
}

// Error Function
func (elce EmptyLinearCoeffsError) Error() string {
	errOut := fmt.Sprintf(
		"the expression of type %T has no variables to compute linear coefficients for",
		elce.Expression,
	)
	return errOut
}
