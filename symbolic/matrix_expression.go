package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
   matrix_expression.go
   Description:

*/

type MatrixExpression interface {
	// Variables returns the number of variables in the expression.
	Variables() []Variable

	//// Coeffs returns a slice of the coefficients in the expression
	//LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.Dense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, errors ...error) (Expression, error)

	// Mult multiplies the current expression with another and returns the
	// resulting expression
	Multiply(e interface{}, errors ...error) (Expression, error)

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhs interface{}, errors ...error) (Constraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhs interface{}, errors ...error) (Constraint, error)

	// Comparison
	// Returns a constraint with respect to the sense (senseIn) between the
	// current expression and another.
	Comparison(rhs interface{}, sense ConstrSense, errors ...error) (Constraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhs interface{}, errors ...error) (Constraint, error)

	//AtVec returns the expression at a given index
	At(i int, j int) ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression

	// Dims returns the dimensions of the given expression
	Dims() []int
}

/*
IsMatrixExpression
Description:

	Determines whether or not an input object is a valid "VectorExpression" according to MatProInterface.
*/
func IsMatrixExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case mat.VecDense:
		return true
	case KMatrix:
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
	if !IsVectorExpression(e) {
		return KMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a MatrixExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case mat.Dense:
		return KMatrix(e2), nil
	default:
		return KMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"unexpected vector expression conversion requested for type %T!",
			e,
		)
	}
}
