package symbolic

/*
monomial.go
Description:
	This file defines the function associated with the Monomial object.
*/

/*
Type Definition
*/
type Monomial struct {
	Coefficient     float64
	Degrees         []int
	VariableFactors []Variable
}
