package symbolic

import "fmt"

/*
monomial.go
Description:
	This file defines the function associated with the Monomial object.
*/

/*
Type Definition
*/
type Monomial struct {
	Coefficient     float64
	Degrees         []int
	VariableFactors []Variable
}

/*
Check
Description:

	This function checks that the monomial is valid.
*/
func (m Monomial) Check() error {
	// Check that the number of degrees matches the number of variables
	if len(m.Degrees) != len(m.VariableFactors) {
		return fmt.Errorf(
			"the number of degrees (%v) does not match the number of variables (%v)",
			len(m.Degrees),
			len(m.VariableFactors),
		)
	}
	// All Checks passed
	return nil
}

/*
Variables
Description:

	Returns the variables in the monomial.
*/
func (m Monomial) Variables() []Variable {
	return m.VariableFactors
}

/*
Dims
Description:

	Returns the dimensions of the monomial. (It is a scalar, so this is [1,1])
*/
func (m Monomial) Dims() []int {
	return []int{1, 1}
}

/*
Plus
Description:

	Multiplication of the monomial with another expression.
*/
func (m Monomial) Plus(e interface{}) Expression {
	// Input Processing
	if IsExpression(e) {
		// Check dimensions
		rightAsE, _ := ToExpression(e)
		err := CheckDimensionsInAddition(m, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return m.Plus(K(right))
	}

	// Unrecornized response is a panic
	panic(
		fmt.Errorf("Unexpected type of right in the Plus() method: %T (%v)", e, e),
	)
}

/*
Multiply
Description:

	Defines the multiplication operation between a monomial and another expression.
*/
func (m Monomial) Multiply(e interface{}) Expression {
	// Input Processing
	if IsExpression(e) {
		// Check dimensions
		rightAsE, _ := ToExpression(e)
		err := CheckDimensionsInMultiplication(m, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return m.Multiply(K(right))
	case K:
		return m.Multiply(Monomial{Coefficient: float64(right)})
	case Variable:
		monomialOut := m

		if foundIndex, _ := FindInSlice(right, m.VariableFactors); foundIndex == -1 {
			monomialOut.VariableFactors = append(monomialOut.VariableFactors, right)
			monomialOut.Degrees = append(monomialOut.Degrees, 1)
		} else {
			monomialOut.Degrees[foundIndex] += 1
		}
		return monomialOut
	}

	// Unrecornized response is a panic
	panic(
		fmt.Errorf("Unexpected type of right in the Multiply() method: %T (%v)", e, e),
	)
}

/*
Transpose
Description:

	Transposes the scalar monomial and returns it. (This is the same as simply copying the monomial.)
*/
func (m Monomial) Transpose() Expression {
	return m
}

/*
LessEq
Description:

	Returns a constraint between a monomial being less than an
	expression.
*/
func (m Monomial) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return m.Comparison(rightIn, SenseLessThanEqual, errors...)
}

/*
GreaterEq
Description:

	Returns a constraint between a monomial being greater than an
	expression.
*/
func (m Monomial) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return m.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

/*
Eq
Description:

	Returns a constraint between a monomial being equal to an
	expression.
*/
func (m Monomial) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return m.Comparison(rightIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	Base method for creating constraints as comparisons between
	two different expressions according to a sense.
*/
func (m Monomial) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return ScalarConstraint{}, err
	}

	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		return ScalarConstraint{}, err
	}

	// Algorithm
	return ScalarConstraint{m, rhs, sense}, nil
}

/*
Constant
Description:

	Returns the constant component in the scalar monomial.
	This should be zero unless there are no variables present. Then it will be the coefficient.
*/
func (m Monomial) Constant() float64 {
	if len(m.VariableFactors) == 0 {
		return m.Coefficient
	} else {
		return 0
	}
}
