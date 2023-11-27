package symbolic

import (
	"fmt"
)

/*
Integer constants representing commonly used numbers. Makes for better
readability
*/
const (
	Zero     = K(0)
	One      = K(1)
	Infinity = K(1e100)
)

// K is a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
type K float64

/*
Variables
Description:

	Shares all variables included in the expression that is K.
	It is a constant, so there are none.
*/
func (c K) Variables() []Variable {
	return []Variable{}
}

// Vars returns a slice of the Var ids in the expression. For constants,
// this is always nil
func (c K) IDs() []uint64 {
	return nil
}

// Coeffs returns a slice of the coefficients in the expression. For constants,
// this is always nil
func (c K) Coeffs() []float64 {
	return nil
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

// Plus adds the current expression to another and returns the resulting
// expression
func (c K) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(c, rightAsE)
		if err != nil {
			return c, err
		}
	}

	// Switching based on input type
	switch right := rightIn.(type) {
	case K:
		return K(c.Constant() + right.Constant()), nil
	case Variable:
		return right.Plus(c)
	default:
		return c, fmt.Errorf("Unexpected type in K.Plus() for constant %v: %T", right, right)
	}
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.
*/
func (c K) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// InputProcessing
	err := CheckErrors(errors)
	if err != nil {
		return ScalarConstraint{}, err
	}

	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		return ScalarConstraint{}, err
	}

	// Constants

	// Algorithm
	return ScalarConstraint{c, rhs, sense}, nil
}

/*
Multiply
Description:

	This method multiplies the input constant by another expression.
*/
func (c K) Multiply(term1 interface{}, errors ...error) (Expression, error) {
	// Constants

	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	if IsExpression(term1) {
		// Check dimensions
		term1AsE, _ := ToExpression(term1)
		err = CheckDimensionsInMultiplication(c, term1AsE)
		if err != nil {
			return c, err
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return c.Multiply(K(right))
	case K:
		return c * right, nil
	default:
		return K(0), fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1)

	}
}

func (c K) Dims() []int {
	return []int{1, 1} // Signifies scalar
}

func (c K) Check() error {
	return nil
}

func (c K) Transpose() Expression {
	return c
}

/*
ToMonomial
Description:

	Converts the constant into a monomial.
*/
func (c K) ToMonomial() Monomial {
	return Monomial{
		Coefficient:     float64(c),
		VariableFactors: []Variable{},
		Degrees:         []int{},
	}
}
