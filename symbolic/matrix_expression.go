package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
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

	// Mult multiplies the current expression with another and returns the
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

	// Substitute returns the expression with the variable vIn replaced with the expression eIn
	Substitute(vIn Variable, eIn Expression) Expression

	// SubstituteAccordingTo returns the expression with the variables in the map replaced with the corresponding expressions
	SubstituteAccordingTo(subMap map[Variable]Expression) Expression

	// Power
	// Raises the scalar expression to the power of the input integer
	Power(exponent int) Expression
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
ToMatrixExpression
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ToMatrixExpression(e interface{}) (MatrixExpression, error) {
	// Input Processing
	if !IsMatrixExpression(e) {
		return DenseToKMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a MatrixExpression.",
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

/*
IsSquare
Description:

	Determines whether the input matrix expression is square.
*/
func IsSquare(e MatrixExpression) bool {
	dims := e.Dims()
	return dims[0] == dims[1]
}

/*
MatrixPowerTemplate
Description:

	Template for the matrix power function.
*/
func MatrixPowerTemplate(me MatrixExpression, exponent int) MatrixExpression {
	// Input Processing
	err := me.Check()
	if err != nil {
		panic(err)
	}

	// Check if the matrix is square
	if !IsSquare(me) {
		panic(fmt.Errorf("matrix is not square; cannot raise to power"))
	}

	// Check if the power is non-negative
	if exponent < 0 {
		panic(smErrors.NegativeExponentError{
			Exponent: exponent,
		})
	}

	// Algorithm
	var out MatrixExpression = K(0).Plus(Identity(me.Dims()[0])).(MatrixExpression)
	for i := 0; i < exponent; i++ {
		out = out.Multiply(me).(MatrixExpression)
	}
	return out
}
