package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
Expression
Description:

	This interface should be implemented by and ScalarExpression and VectorExpression
*/
type Expression interface {
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

	// Substitute returns the expression with the variable vIn replaced with the expression eIn
	Substitute(vIn Variable, eIn Expression) Expression

	// SubstituteAccordingToMap returns the expression with the variables in the map replaced with the corresponding expressions
	SubstituteAccordingTo(subMap map[Variable]Expression) Expression

	// Power
	// Raises the scalar expression to the power of the input integer
	Power(exponent int) Expression
}

/*
NumVariables
Description:

	The number of distinct variables.
*/
func NumVariables(e Expression) int {
	return len(e.Variables())
}

/*
VariableIDs
Description:

	Returns a list of ids associated with each variable.
*/
func VariableIDs(e Expression) []uint64 {
	vSlice := e.Variables()

	var idSlice []uint64
	for _, v := range vSlice {
		idSlice = append(idSlice, v.ID)
	}

	return idSlice
}

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
func IsExpression(e interface{}) bool {
	return IsScalarExpression(e) || IsVectorExpression(e) || IsMatrixExpression(e)
}

func ToExpression(e interface{}) (Expression, error) {
	switch {
	case IsScalarExpression(e):
		return ToScalarExpression(e)
	case IsVectorExpression(e):
		return ToVectorExpression(e)
	case IsMatrixExpression(e):
		return ToMatrixExpression(e)
	}

	// If the input is not a valid expression, panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "ToExpression",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	subtracts the current expression from another and returns the resulting expression
*/
func Minus(left, right Expression) Expression {
	return left.Plus(
		right.Multiply(-1.0),
	)
}

/*
IsLinear
Description:

	Determines whether an input object is a
	valid linear expression.
	In math, this means that the polynomial like expression
	has a degree less than or equal to 1.
*/
func IsLinear(e Expression) bool {
	// Input Processing
	if !IsPolynomialLike(e) {
		return false // Not a polynomial like expression, so it can't be linear
	}

	eAsPL, _ := ToPolynomialLike(e)

	return eAsPL.Degree() <= 1
}

/*
IsQuadratic
Description:

	Determines whether or not an input object is a
	valid Quadratic Expression.
	In math, this means that the polynomial like expression
	has a degree less than or equal to 2.
*/
func IsQuadratic(e Expression) bool {
	// Input Processing
	if !IsPolynomialLike(e) {
		return false // Not a polynomial like expression, so it can't be linear
	}

	eAsPL, _ := ToPolynomialLike(e)

	return eAsPL.Degree() <= 2
}
