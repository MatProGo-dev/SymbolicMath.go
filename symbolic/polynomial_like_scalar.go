package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// ScalarExpression represents a linear general expression of the form
// c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
// variables and k is a constant. This is a base interface that is implemented
// by single variables, constants, and general linear expressions.
type PolynomialLikeScalar interface {
	// Check returns an error if the expression is not valid
	Check() error

	// Variables returns the variables included in the scalar expression
	Variables() []Variable

	// Constant returns the constant additive value in the expression
	Constant() float64

	// LinearCoeff returns the coefficient of the linear terms in the expression
	LinearCoeff(wrt ...[]Variable) mat.VecDense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(rightIn interface{}) Expression

	// Minus subtracts an expression from the current one and returns the resulting
	// expression
	Minus(rightIn interface{}) Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhsIn interface{}) Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhsIn interface{}) Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhsIn interface{}) Constraint

	//Comparison
	// Compares the receiver expression rhs with the expression rhs in the sense of sense.
	Comparison(rhsIn interface{}, sense ConstrSense) Constraint

	//Multiply
	// Multiplies the given scalar expression with another expression
	Multiply(rightIn interface{}) Expression

	//Dims
	// Returns the dimensions of the scalar expression (should always be 1,1)
	Dims() []int

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn Variable) Expression

	// Degree returns the degree of the expression
	Degree() int

	// String returns a string representation of the expression
	String() string

	// Substitute replaces the variable v with the expression e
	Substitute(v Variable, se ScalarExpression) Expression

	// SubstituteAccordingTo replaces the variables in the expression with the expressions in the map
	SubstituteAccordingTo(subMap map[Variable]ScalarExpression) Expression

	// Power raises the expression to the power of the input integer
	Power(exponent int) Expression
}

/*
IsPolynomialLikeScalar
Description:

	Determines whether or not an input object is a
	valid "PolynomialLikeScalar" according to MatProInterface.
*/
func IsPolynomialLikeScalar(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case float64:
		return true
	case K:
		return true
	case Variable:
		return true
	case Monomial:
		return true
	case Polynomial:
		return true
	default:
		return false

	}
}

/*
ToPolynomialLikeScalar
Description:

	Converts the input expression to a valid type that
	implements "PolynomialLikeScalar".
*/
func ToPolynomialLikeScalar(e interface{}) (PolynomialLikeScalar, error) {
	// Input Processing
	if !IsPolynomialLikeScalar(e) {
		return K(1.0), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a PolynomialLikeScalar.",
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
	case Monomial:
		return e2, nil
	case Polynomial:
		return e2, nil
	default:
		return K(1.0), fmt.Errorf(
			"unexpected scalar expression conversion requested for type %T!",
			e,
		)
	}
}
