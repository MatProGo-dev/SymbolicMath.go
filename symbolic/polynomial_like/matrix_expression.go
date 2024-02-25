package polynomial_like

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
   matrix_expression.go
   Description:

*/

type MatrixExpression interface {
	// Check returns an error if the expression is not initialized properly
	Check() error

	// Variables returns the number of variables in the expression.
	Variables() []symbolic.Variable

	//// Coeffs returns a slice of the coefficients in the expression
	//LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.Dense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}) symbolic.Expression

	// Mult multiplies the current expression with another and returns the
	// resulting expression
	Multiply(e interface{}) symbolic.Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhs interface{}) symbolic.Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhs interface{}) symbolic.Constraint

	// Comparison
	// Returns a constraint with respect to the sense (senseIn) between the
	// current expression and another.
	Comparison(rhs interface{}, sense symbolic.ConstrSense) symbolic.Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhs interface{}) symbolic.Constraint

	//AtVec returns the expression at a given index
	At(i int, j int) symbolic.ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() symbolic.Expression

	// Dims returns the dimensions of the given expression
	Dims() []int

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn symbolic.Variable) symbolic.Expression

	// String returns a string representation of the expression
	String() string
}

/*
IsMatrixExpression
Description:

	Determines whether or not an input object is a valid "VectorExpression" according to MatProInterface.
*/
func IsMatrixExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case mat.Dense:
		return true
	case symbolic.KMatrix:
		return true
	case symbolic.VariableMatrix:
		return true
	case symbolic.MonomialMatrix:
		return true
	case symbolic.PolynomialMatrix:
		return true
	default:
		return false

	}
}

/*
ToMatrixExpression
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ToMatrixExpression(e interface{}) (MatrixExpression, error) {
	// Input Processing
	if !IsMatrixExpression(e) {
		return symbolic.DenseToKMatrix(symbolic.ZerosMatrix(1, 1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a MatrixExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case mat.Dense:
		return symbolic.DenseToKMatrix(e2), nil
	case symbolic.KMatrix:
		return e2, nil
	case symbolic.VariableMatrix:
		return e2, nil
	case symbolic.MonomialMatrix:
		return e2, nil
	case symbolic.PolynomialMatrix:
		return e2, nil
	default:
		return symbolic.DenseToKMatrix(symbolic.ZerosMatrix(1, 1)), fmt.Errorf(
			"unexpected vector expression conversion requested for type %T!",
			e,
		)
	}
}
