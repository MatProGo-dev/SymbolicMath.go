package polynomial_like

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
	Variables() []symbolic.Variable

	// Dims returns a slice describing the true dimensions of a given expression (scalar, vector, or matrix)
	Dims() []int

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(rightIn interface{}) symbolic.Expression

	// Multiply multiplies the current expression to another and returns the
	// resulting expression
	Multiply(rightIn interface{}) symbolic.Expression

	// Transpose transposes the given expression
	Transpose() symbolic.Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rightIn interface{}) symbolic.Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rightIn interface{}) symbolic.Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rightIn interface{}) symbolic.Constraint

	// Comparison
	Comparison(rightIn interface{}, sense symbolic.ConstrSense) symbolic.Constraint

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn symbolic.Variable) symbolic.Expression

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
	//return IsScalarExpression(e) || IsVectorExpression(e) || IsMatrixExpression(e)
	return false
}

func ToPolynomialLike(e interface{}) (symbolic.Expression, error) {
	//switch {
	//case IsScalarExpression(e):
	//	return ToScalarExpression(e)
	//case IsVectorExpression(e):
	//	return ToVectorExpression(e)
	//case IsMatrixExpression(e):
	//	return ToMatrixExpression(e)
	//}

	// If the input is not a valid expression, panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "ToExpression",
			Input:        e,
		},
	)
}
