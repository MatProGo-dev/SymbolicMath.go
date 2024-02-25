package polynomial_like

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

// ScalarExpression represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant. This is a base interface that is implemented
// by single variables, constants, and general linear expressions.
type ScalarExpression interface {
	// Check returns an error if the expression is not valid
	Check() error

	// Variables returns the variables included in the scalar expression
	Variables() []symbolic.Variable

	// Constant returns the constant additive value in the expression
	Constant() float64

	// LinearCoeff returns the coefficient of the linear terms in the expression
	LinearCoeff(wrt ...[]symbolic.Variable) mat.VecDense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(rightIn interface{}) symbolic.Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhsIn interface{}) symbolic.Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhsIn interface{}) symbolic.Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhsIn interface{}) symbolic.Constraint

	//Comparison
	// Compares the receiver expression rhs with the expression rhs in the sense of sense.
	Comparison(rhsIn interface{}, sense symbolic.ConstrSense) symbolic.Constraint

	//Multiply
	// Multiplies the given scalar expression with another expression
	Multiply(rightIn interface{}) symbolic.Expression

	//Dims
	// Returns the dimensions of the scalar expression (should always be 1,1)
	Dims() []int

	//Transpose returns the transpose of the given vector expression
	Transpose() symbolic.Expression

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn symbolic.Variable) symbolic.Expression

	// IsLinear returns true if the expression is linear
	IsLinear() bool

	// String returns a string representation of the expression
	String() string
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
	case symbolic.K:
		return true
	case symbolic.Variable:
		return true
	case symbolic.Monomial:
		return true
	case symbolic.Polynomial:
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
		return symbolic.K(1.0), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a ScalarExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case float64:
		return symbolic.K(e2), nil
	case symbolic.K:
		return e2, nil
	case symbolic.Variable:
		return e2, nil
	case symbolic.Monomial:
		return e2, nil
	case symbolic.Polynomial:
		return e2, nil
	default:
		return symbolic.K(1.0), fmt.Errorf(
			"unexpected scalar expression conversion requested for type %T!",
			e,
		)
	}
}
