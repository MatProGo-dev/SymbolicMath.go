package symbolic

import (
	"fmt"
)

// ScalarExpression represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant. This is a base interface that is implemented
// by single variables, constants, and general linear expressions.
type ScalarExpression interface {
	// Variables returns the variables included in the scalar expression
	Variables() []Variable

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	// Coeffs returns a slice of the coefficients in the expression
	Coeffs() []float64

	// Constant returns the constant additive value in the expression
	Constant() float64

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(rightIn interface{}, errors ...error) (Expression, error)

	// Mult multiplies the current expression to another and returns the
	// resulting expression
	//Mult(c float64, errors ...error) (Expression, error)

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhsIn interface{}, errors ...error) (Constraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhsIn interface{}, errors ...error) (Constraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhsIn interface{}, errors ...error) (Constraint, error)

	//Comparison
	// Compares the receiver expression rhs with the expression rhs in the sense of sense.
	Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error)

	//Multiply
	// Multiplies the given scalar expression with another expression
	Multiply(rightIn interface{}, errors ...error) (Expression, error)

	//Dims
	// Returns the dimensions of the scalar expression (should always be 1,1)
	Dims() []int

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression
}

// NewExpr returns a new expression with a single additive constant value, c,
// and no variables. Creating an expression like sum := NewExpr(0) is useful
// for creating new empty expressions that you can perform operatotions on
// later
//func NewScalarExpression(c float64) ScalarExpression {
//	return ScalarLinearExpr{C: c}
//}

/*
IsScalarExpression
Description:

	Determines whether or not an input object is a
	valid "ScalarExpression" according to MatProInterface.
*/
func IsScalarExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case float64:
		return true
	case K:
		return true
	case Variable:
		return true
	default:
		return false

	}
}

/*
ToScalarExpression
Description:

	Converts the input expression to a valid type that
	implements "ScalarExpression".
*/
func ToScalarExpression(e interface{}) (ScalarExpression, error) {
	// Input Processing
	if !IsScalarExpression(e) {
		return K(1.0), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a ScalarExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case float64:
		return K(e2), nil
	case K:
		return e2, nil
	case Variable:
		return e2, nil
	default:
		return K(1.0), fmt.Errorf(
			"unexpected scalar expression conversion requested for type %T!",
			e,
		)
	}
}
