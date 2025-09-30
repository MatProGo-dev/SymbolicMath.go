package symbolic

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

/*
   matrix_expression.go
   Description:

*/

type PolynomialLikeMatrix interface {
	// Check returns an error if the expression is not initialized properly
	Check() error

	// Variables returns the number of variables in the expression.
	Variables() []Variable

	//// Coeffs returns a slice of the coefficients in the expression
	//LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.Dense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}) Expression

	// Minus subtracts an expression from the current one and returns the resulting
	// expression
	Minus(rightIn interface{}) Expression

	// Multiply multiplies the current expression with another and returns the
	// resulting expression
	Multiply(e interface{}) Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhs interface{}) Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhs interface{}) Constraint

	// Comparison
	// Returns a constraint with respect to the sense (senseIn) between the
	// current expression and another.
	Comparison(rhs interface{}, sense ConstrSense) Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhs interface{}) Constraint

	//AtVec returns the expression at a given index
	At(i int, j int) ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression

	// Dims returns the dimensions of the given expression
	Dims() []int

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn Variable) Expression

	// String returns a string representation of the expression
	String() string

	// Degree returns the degree of the expression
	Degree() int

	// Substitute returns the expression with the variable vIn replaced with the expression eIn
	Substitute(vIn Variable, eIn ScalarExpression) Expression

	// SubstitueAccordingTo returns the expression with the variables in the map replaced with the corresponding expressions
	SubstituteAccordingTo(subMap map[Variable]Expression) Expression

	// Power returns the expression raised to the power of the input exponent
	Power(exponent int) Expression

	// Simplify simplifies the expression and returns the simplified version
	AsSimplifiedExpression() Expression
}

/*
IsPolynomialLikeMatrix
Description:

	Determines whether or not an input object is a valid "VectorExpression" according to MatProInterface.
*/
func IsPolynomialLikeMatrix(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case mat.Dense:
		return true
	case KMatrix:
		return true
	case VariableMatrix:
		return true
	case MonomialMatrix:
		return true
	case PolynomialMatrix:
		return true
	default:
		return false

	}
}

/*
ToPolynomialLikeMatrix
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ToPolynomialLikeMatrix(e interface{}) (PolynomialLikeMatrix, error) {
	// Input Processing
	if !IsPolynomialLikeMatrix(e) {
		return DenseToKMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a PolynomialLikeMatrix.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case mat.Dense:
		return DenseToKMatrix(e2), nil
	case KMatrix:
		return e2, nil
	case VariableMatrix:
		return e2, nil
	case MonomialMatrix:
		return e2, nil
	case PolynomialMatrix:
		return e2, nil
	default:
		return DenseToKMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"unexpected vector expression conversion requested for type %T!",
			e,
		)
	}
}
