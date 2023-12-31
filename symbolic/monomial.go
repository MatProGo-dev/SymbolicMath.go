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
	err := m.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		// Convert to expression
		rightAsE, _ := ToExpression(e)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}
		// Check Dimensions
		err := CheckDimensionsInAddition(m, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return m.Plus(K(right))
	case K:
		if m.IsConstant() {
			mCopy := m
			mCopy.Coefficient += float64(right)
			return mCopy
		} else {
			return Polynomial{
				Monomials: []Monomial{m, right.ToMonomial()},
			}
		}
	case Variable:
		if m.IsVariable(right) {
			mCopy := m
			mCopy.Coefficient += 1.0
			return mCopy
		} else {
			return Polynomial{
				Monomials: []Monomial{m, right.ToMonomial()},
			}
		}
	case Polynomial:
		return right.Plus(m)
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
		rightAsFloat64 := float64(right)
		monomialOut := m
		monomialOut.Coefficient *= rightAsFloat64
		return monomialOut
	case Variable:
		monomialOut := m

		if foundIndex, _ := FindInSlice(right, m.VariableFactors); foundIndex == -1 {
			monomialOut.VariableFactors = append(monomialOut.VariableFactors, right)
			monomialOut.Degrees = append(monomialOut.Degrees, 1)
		} else {
			monomialOut.Degrees[foundIndex] += 1
		}
		return monomialOut
	case Monomial:
		var monomialOut Monomial

		// Collect all variables in both monomials
		variables := append(m.VariableFactors, right.VariableFactors...)
		variables = UniqueVars(variables)

		multiDegree := make([]int, len(variables))

		// Iterate through each variable in monomial one
		for ii, variable := range m.VariableFactors {
			// Find the index of the variable in the slice variables
			foundIndex, _ := FindInSlice(variable, variables)
			// And modify the multidegree
			multiDegree[foundIndex] += m.Degrees[ii]
		}

		// Iterate through each variable in monomial two
		for ii, variable := range right.VariableFactors {
			// Find the index of the variable in the slice variables
			foundIndex, _ := FindInSlice(variable, variables)
			// And modify the multidegree
			multiDegree[foundIndex] += right.Degrees[ii]
		}

		// Create coefficient
		monomialOut.Coefficient = m.Coefficient * right.Coefficient

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
func (m Monomial) LessEq(rightIn interface{}) Constraint {
	return m.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Returns a constraint between a monomial being greater than an
	expression.
*/
func (m Monomial) GreaterEq(rightIn interface{}) Constraint {
	return m.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Returns a constraint between a monomial being equal to an
	expression.
*/
func (m Monomial) Eq(rightIn interface{}) Constraint {
	return m.Comparison(rightIn, SenseEqual)
}

/*
Comparison
Description:

	Base method for creating constraints as comparisons between
	two different expressions according to a sense.
*/
func (m Monomial) Comparison(rhsIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return ScalarConstraint{m, rhs, sense}
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

/*
IsConstant
Description:

	Returns true if the monomial defines a constant.
*/
func (m Monomial) IsConstant() bool {
	// Input Checking
	err := m.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return len(m.VariableFactors) == 0
}

/*
IsVariable
Description:

	Returns true if the monomial defines an expression containing only the
	variable v.
*/
func (m Monomial) IsVariable(v Variable) bool {
	// Input Checking
	err := m.Check()
	if err != nil {
		panic(err)
	}

	err = v.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	containsOnlyOneFactor := len(m.VariableFactors) == 1
	firstFactorMatchesV := m.VariableFactors[0].ID == v.ID
	firstFactorHasDegreeOne := m.Degrees[0] == 1
	if containsOnlyOneFactor && firstFactorMatchesV && firstFactorHasDegreeOne {
		return true
	} else {
		return false
	}
}

/*
MatchesFormOf
Description:

	Returns true if the monomial matches the form of the input monomial.
	(in other words if the input monomial has the same variables and degrees as the input monomial.)
*/
func (m Monomial) MatchesFormOf(mIn Monomial) bool {
	// Input Checking
	err := m.Check()
	if err != nil {
		panic(err)
	}

	err = mIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	if len(m.VariableFactors) != len(mIn.VariableFactors) {
		return false
	}

	for ii, v := range m.VariableFactors {
		foundIndex, _ := FindInSlice(v, mIn.VariableFactors)
		if foundIndex == -1 {
			// If v was not in mIn, then these two monomials are not the same
			return false
		} else {
			// If v was in mIn, but not of the right degree, then these two are not the same
			if m.Degrees[ii] != mIn.Degrees[ii] {
				return false
			}
		}

	}

	// If all checks pass, then return true!
	return true
}

/*
DerivativeWrt
Description:

	This function returns the derivative of the monomial with respect to the input
	variable vIn. If the monomial does not contain the variable vIn, then the
	derivative is zero.
	If the monomial does contain the variable vIn, then the derivative is the monomial
	with a decreased degree of vIn and a coefficient equal to the original coefficient.
*/
func (m Monomial) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	foundIndex, _ := FindInSlice(vIn, m.VariableFactors)
	if foundIndex == -1 {
		// If vIn is not in the monomial, then the derivative is zero
		return K(0.0)
	} else {
		// If vIn is in the monomial, then decrease that element's degree in the
		// monomial.
		var monomialOut Monomial
		if monomialOut.Degrees[foundIndex] == 1 {
			// If the degree of vIn is 1, then remove it from the monomial
			monomialOut.Coefficient = m.Coefficient
			for ii, variable := range m.VariableFactors {
				if ii != foundIndex {
					monomialOut.VariableFactors = append(monomialOut.VariableFactors, variable)
					monomialOut.Degrees = append(monomialOut.Degrees, m.Degrees[ii])
				}
			}
		} else {
			monomialOut = m
			monomialOut.Degrees[foundIndex] -= 1
		}

		// Return monomial
		return monomialOut
	}
}

/*
IsLinear
Description:

	This function returns true only if the sum of all degrees in the monomial is
	less than or equal to 1.
*/
func (m Monomial) IsLinear() bool {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	sum := 0
	for _, degree := range m.Degrees {
		sum += degree
	}

	return sum <= 1
}

/*
ToPolynomial
Description:

	Creates a copy of the monomial m as a polynomial.
*/
func (m Monomial) ToPolynomial() Polynomial {
	return Polynomial{
		Monomials: []Monomial{m},
	}
}
