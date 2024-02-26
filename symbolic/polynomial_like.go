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
}

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
func IsPolynomialLike(e interface{}) bool {
	return IsPolynomialLikeScalar(e) || IsPolynomialLikeVector(e) || IsPolynomialLikeMatrix(e)
}

func ToPolynomialLike(e interface{}) (Expression, error) {
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
