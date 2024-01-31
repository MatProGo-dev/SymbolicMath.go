package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
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
Check
Description:

	Checks to make sure that the constant is initialized properly.
	Constants are always initialized properly, so this should always return
	no error.
*/
func (c K) Check() error {
	return nil
}

/*
Variables
Description:

	Shares all variables included in the expression that is K.
	It is a constant, so there are none.
*/
func (c K) Variables() []Variable {
	return []Variable{}
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

/*
Plus
Description:

	adds the current expression to another and returns the resulting expression
*/
func (c K) Plus(rightIn interface{}) Expression {
	// Input Processing
	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to %v.Plus: %v", c, err))
		}

		err = CheckDimensionsInAddition(c, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Switching based on input type
	switch right := rightIn.(type) {
	case float64:
		return c.Plus(K(right))
	case K:
		return K(c.Constant() + right.Constant())
	case Variable:
		return right.Plus(c)
	case Monomial:
		return right.Plus(c)
	case Polynomial:
		return right.Plus(c)
	case mat.VecDense:
		return c.Plus(VecDenseToKVector(right))
	case *mat.VecDense:
		return c.Plus(VecDenseToKVector(*right))
	case KVector:
		return right.Plus(c)
	case PolynomialVector:
		return right.Plus(c)
	}

	// Default response is a panic
	panic(
		fmt.Errorf("Unexpected type in K.Plus() for constant %v: %T", rightIn, rightIn),
	)
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(rightIn interface{}) Constraint {
	return c.Comparison(rightIn, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(rightIn interface{}) Constraint {
	return c.Comparison(rightIn, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(rightIn interface{}) Constraint {
	return c.Comparison(rightIn, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.
*/
func (c K) Comparison(rhsIn interface{}, sense ConstrSense) Constraint {
	// InputProcessing
	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		panic(err)
	}

	// Constants

	// Algorithm
	return ScalarConstraint{c, rhs, sense}
}

/*
Multiply
Description:

	This method multiplies the input constant by another expression.
*/
func (c K) Multiply(term1 interface{}) Expression {
	// Constants

	// Input Processing
	if IsExpression(term1) {
		// Check dimensions
		term1AsE, _ := ToExpression(term1)
		err := CheckDimensionsInMultiplication(c, term1AsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return c.Multiply(K(right))
	case K:
		return c * right
	case Variable:
		return right.Multiply(c)
	}

	// Unrecornized response is a panic
	panic(
		fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1),
	)
}

func (c K) Dims() []int {
	return []int{1, 1} // Signifies scalar
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
		Exponents:       []int{},
	}
}

/*
DerivativeWrt
Description:

	Computes the derivative of a constant, which should be 0 for any constant.
*/
func (c K) DerivativeWrt(vIn Variable) Expression {
	return Zero
}

/*
IsLinear
Description:

	Returns true always. A constant expression is always linear.
*/
func (c K) IsLinear() bool {
	return true
}

/*
String
Description:

	Returns a string representation of the constant.
*/
func (c K) String() string {
	return fmt.Sprintf("%v", float64(c))
}
