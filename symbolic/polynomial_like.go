package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
polynomial_like.go
Description:

	This interface should be implemented by all objects
	that are polynomial like (i.e., are linear combinations
	of variables (potentially raised to powers) and coefficients).
*/
type PolynomialLike interface {
	// NumVars returns the number of variables in the expression
	Variables() []Variable

	// Dims returns a slice describing the true dimensions of a given expression (scalar, vector, or matrix)
	Dims() []int

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(rightIn interface{}) Expression

	// Minus subtracts an expression from the current one and returns the resulting
	// expression
	Minus(rightIn interface{}) Expression

	// Multiply multiplies the current expression to another and returns the
	// resulting expression
	Multiply(rightIn interface{}) Expression

	// Transpose transposes the given expression
	Transpose() Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rightIn interface{}) Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rightIn interface{}) Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rightIn interface{}) Constraint

	// Comparison
	Comparison(rightIn interface{}, sense ConstrSense) Constraint

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn Variable) Expression

	// Check
	Check() error

	// String returns a string representation of the expression
	String() string

	// Degree returns the degree of the expression
	Degree() int

	// Substitute returns the expression with the variable vIn replaced with the expression eIn
	Substitute(vIn Variable, eIn ScalarExpression) Expression

	// SubstituteAccordingTo returns the expression with the variables in the map replaced with the corresponding expressions
	SubstituteAccordingTo(subMap map[Variable]Expression) Expression

	// Power returns the expression raised to the power of the input exponent
	Power(exponent int) Expression

	// At returns the value at the given row and column index
	At(ii, jj int) ScalarExpression
}

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
func IsPolynomialLike(e interface{}) bool {
	return IsPolynomialLikeScalar(e) || IsPolynomialLikeVector(e) || IsPolynomialLikeMatrix(e)
}

func ToPolynomialLike(e interface{}) (PolynomialLike, error) {
	switch {
	case IsPolynomialLikeScalar(e):
		return ToPolynomialLikeScalar(e)
	case IsPolynomialLikeVector(e):
		return ToPolynomialLikeVector(e)
	case IsPolynomialLikeMatrix(e):
		return ToPolynomialLikeMatrix(e)
	}

	// If the input is not a valid expression, panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "ToPolynomialLike",
			Input:        e,
		},
	)
}
