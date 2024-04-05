package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

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
	Exponents       []int
	VariableFactors []Variable
}

/*
Check
Description:

	This function checks that the monomial is valid.
*/
func (m Monomial) Check() error {
	// Check that the number of degrees matches the number of variables
	if len(m.Exponents) != len(m.VariableFactors) {
		return fmt.Errorf(
			"the number of degrees (%v) does not match the number of variables (%v)",
			len(m.Exponents),
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
		err := smErrors.CheckDimensionsInAddition(m, rightAsE)
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
			return K(m.Coefficient + float64(right))
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
	case Monomial:
		if m.MatchesFormOf(right) {
			monomialOut := m.Copy()
			monomialOut.Coefficient += right.Coefficient
			return monomialOut
		} else {
			return Polynomial{
				Monomials: []Monomial{m, right},
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
Minus
Description:

	Subtraction of the monomial with another expression.
*/
func (m Monomial) Minus(e interface{}) Expression {
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
		err := smErrors.CheckDimensionsInAddition(m, rightAsE)
		if err != nil {
			panic(err)
		}

		// Use Expression's method
		return Minus(m, rightAsE)
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return m.Minus(K(right)) // Reuse K case
	}

	// Unrecognized response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Monomial.Minus",
			Input:        e,
		},
	)
}

/*
Multiply
Description:

	Defines the multiplication operation between a monomial and another expression.
*/
func (m Monomial) Multiply(e interface{}) Expression {
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

		// Check dimensions
		err := smErrors.CheckDimensionsInMultiplication(m, rightAsE)
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
		monomialOut := m.Copy()

		if foundIndex, _ := FindInSlice(right, m.VariableFactors); foundIndex == -1 {
			monomialOut.VariableFactors = append(monomialOut.VariableFactors, right)
			monomialOut.Exponents = append(monomialOut.Exponents, 1)
		} else {
			monomialOut.Exponents[foundIndex] += 1
		}
		return monomialOut
	case Monomial:

		// Collect all variables in both monomials
		variables := append(m.VariableFactors, right.VariableFactors...)
		variables = UniqueVars(variables)

		multiDegree := make([]int, len(variables))

		// Iterate through each variable in monomial one
		for ii, variable := range m.VariableFactors {
			// Find the index of the variable in the slice variables
			foundIndex, _ := FindInSlice(variable, variables)
			// And modify the multidegree
			multiDegree[foundIndex] += m.Exponents[ii]
		}

		// Iterate through each variable in monomial two
		for ii, variable := range right.VariableFactors {
			// Find the index of the variable in the slice variables
			foundIndex, _ := FindInSlice(variable, variables)
			// And modify the multidegree
			multiDegree[foundIndex] += right.Exponents[ii]
		}

		// Create monomialOut
		return Monomial{
			Coefficient:     m.Coefficient * right.Coefficient,
			Exponents:       multiDegree,
			VariableFactors: variables,
		}
	case Polynomial:
		return right.Multiply(m) // Commutative
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
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

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
	err := m.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rhsIn) {
		// Check rhs
		rhsAsE, _ := ToExpression(rhsIn)
		err = rhsAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		// Scalar function, so no need to check dimensions
		//err = CheckDimensionsInComparison(m, rhsAsE, sense)
		//if err != nil {
		//	panic(err)
		//}
	}

	// Algorithm
	switch right := rhsIn.(type) {
	case float64:
		return m.Comparison(K(right), sense)
	case K:
		return ScalarConstraint{m, right, sense}
	case Variable:
		return ScalarConstraint{m, right, sense}
	case Monomial:
		return ScalarConstraint{m, right, sense}
	case Polynomial:
		return ScalarConstraint{m, right, sense}
	}

	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Monomial.Comparison (" + sense.String() + ")",
			Input:        rhsIn,
		},
	)
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
LinearCoeffs
Description:

	Returns the coefficients of the linear terms in the monomial.
*/
func (m Monomial) LinearCoeff(wrt ...[]Variable) mat.VecDense {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	var wrtVars []Variable
	switch len(wrt) {
	case 0:
		wrtVars = m.Variables()
	case 1:
		wrtVars = wrt[0]
	default:
		panic(
			fmt.Errorf("Too many inputs provided to Monomial.LinearCoeff() method. Expected 0 or 1 input."),
		)
	}

	if len(wrtVars) == 0 {
		panic(smErrors.CanNotGetLinearCoeffOfConstantError{m})
	}

	// Algorithm
	// Create slice
	linearCoeffs := ZerosVector(len(wrtVars))

	// Iterate through each variable
	if IsLinear(m) {
		// If the monomial is linear,
		// then find the variable that is present
		idx, _ := FindInSlice(m.VariableFactors[0], wrtVars)
		linearCoeffs.SetVec(idx, m.Coefficient)
	}

	// Return
	return linearCoeffs
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

	// If the monomial is a constant, then it is not a variable
	// return false
	if m.IsConstant() {
		return false
	}

	// Algorithm
	containsOnlyOneFactor := len(m.VariableFactors) == 1
	firstFactorMatchesV := m.VariableFactors[0].ID == v.ID
	firstFactorHasDegreeOne := m.Exponents[0] == 1
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
			if m.Exponents[ii] != mIn.Exponents[ii] {
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

		if m.Exponents[foundIndex] == 1 {
			// If the degree of vIn is 1, then remove it from the monomial
			monomialOut.Coefficient = m.Coefficient
			for ii, variable := range m.VariableFactors {
				if ii != foundIndex {
					monomialOut.VariableFactors = append(monomialOut.VariableFactors, variable)
					monomialOut.Exponents = append(monomialOut.Exponents, m.Exponents[ii])
				}
			}
		} else {
			monomialOut = m
			monomialOut.Exponents[foundIndex] -= 1
		}

		// Return monomial
		return monomialOut
	}
}

/*
Degree
Description:

	Returns the degree of the monomial.
*/
func (m Monomial) Degree() int {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	degree := 0
	for _, exp := range m.Exponents {
		degree += exp
	}

	// Return
	return degree
}

/*
ToPolynomial
Description:

	Creates a copy of the monomial m as a polynomial.
*/
func (m Monomial) ToPolynomial() Polynomial {
	// Copy Values
	mCopy := m.Copy()

	// Return polynomial with copied monomial
	return Polynomial{
		Monomials: []Monomial{mCopy},
	}
}

/*
String
Description:

	Returns a string representation of the monomial.
*/
func (m Monomial) String() string {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	// Create string
	monomialString := ""

	// Add coefficient
	if (m.Coefficient != 1) || (len(m.VariableFactors) == 0) {
		monomialString += fmt.Sprintf("%v", m.Coefficient)
		if len(m.VariableFactors) != 0 {
			monomialString += " "
		}
	}

	// Add variables
	for ii, variable := range m.VariableFactors {
		monomialString += fmt.Sprintf("%v", variable)
		if m.Exponents[ii] != 1 {
			monomialString += fmt.Sprintf("^%v", m.Exponents[ii])
		}
	}

	// Return
	return monomialString
}

/*
Copy
Description:

	Returns a copy of the monomial.
*/
func (m Monomial) Copy() Monomial {
	// Copy Values
	mCopy := Monomial{
		Coefficient:     m.Coefficient,
		Exponents:       make([]int, len(m.Exponents)),
		VariableFactors: make([]Variable, len(m.VariableFactors)),
	}
	copy(mCopy.Exponents, m.Exponents)
	copy(mCopy.VariableFactors, m.VariableFactors)

	// Return
	return mCopy
}

/*
Substitute
Description:

	Substitutes all occurrences of variable vIn with the expression eIn.
*/
func (m Monomial) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	err = eIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	// Find the index of the variable in the monomial
	foundIndex, _ := FindInSlice(vIn, m.VariableFactors)

	// If the variable is not in the monomial, then return the monomial
	if foundIndex == -1 {
		return m
	}

	// If the variable is in the monomial,
	// then compute the product of the monomial with the expression
	var prod Expression = K(m.Coefficient)
	for ii, variable := range m.VariableFactors {
		prod = prod.Multiply(variable.Substitute(vIn, eIn).Power(m.Exponents[ii]))
	}

	// Return
	return prod
}

/*
SubstituteAccordingTo
Description:

	Substitutes all occurrences of the variables in the map with the corresponding expressions.
*/
func (m Monomial) SubstituteAccordingTo(subMap map[Variable]ScalarExpression) Expression {
	// Input Processing
	err := m.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	// Create the monomial
	var out Expression = K(0.0)

	// Iterate through each variable in the monomial
	for tempVar, tempExp := range subMap {
		out = out.Substitute(tempVar, tempExp)
	}

	// Return
	return out
}

/*
Power
Description:

	Computes the power of the monomial.
*/
func (m Monomial) Power(exponent int) Expression {
	return ScalarPowerTemplate(m, exponent)
}
