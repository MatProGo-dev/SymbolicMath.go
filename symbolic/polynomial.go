package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
polynomial.go
Description:
	This file defines the function associated with the Polynomial object.
*/

/*
Type Definition
*/
type Polynomial struct {
	Monomials []Monomial
}

// Member Methods

/*
Check
Description:

	Verifies that all elements of the polynomial are defined correctly.
*/
func (p Polynomial) Check() error {
	// Check that the polynomial has at least one monomial
	if len(p.Monomials) == 0 {
		return fmt.Errorf("polynomial has no monomials")
	}

	// Check that each of the monomials are well formed
	for ii, monomial := range p.Monomials {
		err := monomial.Check()
		if err != nil {
			return fmt.Errorf("error in monomial %v: %v", ii, err)
		}
	}

	// All checks passed
	return nil
}

/*
Variables
Description:

	The unique variables used to define the polynomial.
*/
func (p Polynomial) Variables() []Variable {
	var variables []Variable // The variables in the polynomial
	for _, monomial := range p.Monomials {
		variables = append(variables, monomial.Variables()...)
	}
	return UniqueVars(variables)
}

/*
Dims
Description:

	The scalar polynomial should have dimensions [1,1].
*/
func (p Polynomial) Dims() []int {
	return []int{1, 1}
}

/*
Plus
Description:

	Defines an addition between the polynomial and another expression.
*/
func (p Polynomial) Plus(e interface{}) Expression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err := CheckDimensionsInAddition(p, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := e.(type) {
	case float64:
		return p.Plus(K(right))
	case K:
		pCopy := p

		// Algorithm
		constantIndex := pCopy.ConstantMonomialIndex()
		if constantIndex == -1 {
			// Monomial does not contain a constant,
			// so add a new monomial.
			rightAsMonom := right.ToMonomial()
			pCopy.Monomials = append(pCopy.Monomials, rightAsMonom)
		} else {
			// Monomial does contain a constant, so
			// modify the monomial which represents that constant.
			newMonomial := pCopy.Monomials[constantIndex]
			newMonomial.Coefficient += float64(right)
			pCopy.Monomials[constantIndex] = newMonomial
		}
		return pCopy

	case Variable:
		pCopy := p

		// Check to see if the variable is already in the polynomial
		variableIndex := pCopy.VariableMonomialIndex(right)
		if variableIndex != -1 {
			// Monomial does not contain the variable,
			// so add a new monomial.
			rightAsMonom := right.ToMonomial()
			pCopy.Monomials = append(pCopy.Monomials, rightAsMonom)
		} else {
			// Monomial does contain the variable, so
			// modify the monomial which represents that variable.
			newMonomial := pCopy.Monomials[variableIndex]
			newMonomial.Coefficient += 1.0
			pCopy.Monomials[variableIndex] = newMonomial
		}

	case Polynomial:
		pCopy := p

		// Combine the list of monomials.
		pCopy.Monomials = append(pCopy.Monomials, right.Monomials...)

		// Simplify?
		return pCopy.Simplify()
	}

	// Unrecognized response is a panic
	panic(
		fmt.Errorf("Unexpected type of right in the Plus() method: %T (%v)", e, e),
	)
}

/*
ConstantMonomialIndex
Description:

	Returns the index of the monomial in the polynomial which is a constant.
	If none are found, then this returns -1.
*/
func (p Polynomial) ConstantMonomialIndex() int {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	for ii, monomial := range p.Monomials {
		if monomial.IsConstant() {
			return ii
		}
	}

	// No constant monomial found
	return -1
}

/*
VariableMonomialIndex
Description:

	Returns the index of the monomial which represents the variable given as vIn.
*/
func (p Polynomial) VariableMonomialIndex(vIn Variable) int {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	for ii, monomial := range p.Monomials {
		if monomial.IsVariable(vIn) {
			return ii
		}
	}

	// No variable monomial found
	return -1
}

/*
MonomialIndex
Description:

	Returns the index of the monomial which has the same
	degrees and variables as the input monomial mIn.
*/
func (p Polynomial) MonomialIndex(mIn Monomial) int {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	err = mIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	for ii, monomial := range p.Monomials {
		if monomial.MatchesFormOf(mIn) {
			return ii
		}
	}

	// No monomial found
	return -1
}

/*
Multiply
Description:

	Implements the multiplication operator between a polynomial and another expression.
*/
func (p Polynomial) Multiply(e interface{}) Expression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err := CheckDimensionsInMultiplication(p, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return p.Multiply(K(right))
	case K:
		pCopy := p
		for ii, _ := range pCopy.Monomials {
			product_ii := pCopy.Monomials[ii].Multiply(right)
			pCopy.Monomials[ii] = product_ii.(Monomial) // Convert to Monomial
		}
		return pCopy
	case Variable:
		pCopy := p
		for ii, _ := range pCopy.Monomials {
			product_ii := pCopy.Monomials[ii].Multiply(right)
			pCopy.Monomials[ii] = product_ii.(Monomial) // Convert to Monomial
		}
		return pCopy
	case Monomial:
		pCopy := p
		for ii, _ := range pCopy.Monomials {
			product_ii := pCopy.Monomials[ii].Multiply(right)
			pCopy.Monomials[ii] = product_ii.(Monomial)
		}
		return pCopy
	case Polynomial:
		pCopy := p

		// Multiply each monomial of the polynomial by the polynomial
		productOut := pCopy.Multiply(right.Monomials[0])
		for ii := 1; ii < len(right.Monomials); ii++ {
			productOut = productOut.Plus(pCopy.Multiply(right.Monomials[ii]))
		}

		return productOut
	}

	// Unrecognized response is a panic
	panic(
		fmt.Errorf("Unexpected type of right in the Multiply() method: %T (%v)", e, e),
	)
}

/*
Transpose
Description:

	The transpose operator when applied to a scalar is just the same scalar object.
*/
func (p Polynomial) Transpose() Expression {
	return p
}

/*
Comparison
Description:

	Creates a constraint between the polynomial and another expression
	of the sense provided in Sense.
*/
func (p Polynomial) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	right, err := ToScalarExpression(rightIn)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return ScalarConstraint{p, right, sense}
}

/*
LessEq
Description:

	Creates a less than equal constraint between the polynomial and another expression.
*/
func (p Polynomial) LessEq(rightIn interface{}) Constraint {
	return p.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Creates a greater than equal constraint between the polynomial and another expression.
*/
func (p Polynomial) GreaterEq(rightIn interface{}) Constraint {
	return p.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Creates an equality constraint between the polynomial and another expression.
*/
func (p Polynomial) Eq(rightIn interface{}) Constraint {
	return p.Comparison(rightIn, SenseEqual)
}

/*
Constant
Description:

	Retrieves the constant component of the polynomial if there is one.
*/
func (p Polynomial) Constant() float64 {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	constantIndex := p.ConstantMonomialIndex()
	if constantIndex == -1 {
		return 0.0
	} else {
		return p.Monomials[constantIndex].Coefficient
	}
}

/*
Simplify
Description:

	This function simplifies the number of monomials in the polynomial,
	by finding the matching terms (i.e., monomials with matching Variables and Degrees)
	and combining them.
*/
func (p Polynomial) Simplify() Polynomial {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm

	// Copy the first element of the polynomial into the new polynomial
	pCopy := Polynomial{
		Monomials: []Monomial{p.Monomials[0]},
	}

	// Loop through the rest of the monomials
	for ii := 1; ii < len(p.Monomials); ii++ {
		// Check to see if the monomial is already in the polynomial
		monomialIndex := pCopy.MonomialIndex(p.Monomials[ii])
		if monomialIndex == -1 {
			// Polynomial does not contain the monomial,
			// so add a new monomial.
			pCopy.Monomials = append(pCopy.Monomials, p.Monomials[ii])
		} else {
			// Monomial does contain the variable, so
			// modify the monomial which represents that variable.
			newMonomial := pCopy.Monomials[monomialIndex]
			newMonomial.Coefficient += p.Monomials[ii].Coefficient
			pCopy.Monomials[monomialIndex] = newMonomial
		}
	}

	// Return the simplified polynomial
	return pCopy

}

/*
DerivativeWrt
Description:

	The derivative of the polynomial with respect to the input variable.
*/
func (p Polynomial) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var derivative Polynomial
	for _, monomial := range p.Monomials {
		if monomial.IsConstant() {
			// Skip constant monomials
			continue
		}

		// Append
		components := monomial.DerivativeWrt(vIn)
		derivative.Monomials = append(derivative.Monomials, components.(Monomial))
	}

	// If the derivative is empty, then return 0.0
	if len(derivative.Monomials) == 0 {
		return K(0.0)
	}

	return derivative
}

/*
IsLinear
Description:

	This function returns true only if all of the monomials in the polynomial are linear.
*/
func (p Polynomial) IsLinear() bool {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	for _, monomial := range p.Monomials {
		if !monomial.IsLinear() {
			return false
		}
	}

	// All monomials are linear
	return true
}

/*
LinearCoeff
Description:

	This function returns a vector describing the coefficients of the linear component
	of the polynomial.
	The (ii)th element of the vector is the coefficient of the (ii)th variable in the
	p.Variables() slice as it appears in the polynomial.
*/
func (p Polynomial) LinearCoeff(vSlices ...[]Variable) mat.VecDense {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Check to see if the user provided a slice of variables
	var varSlice []Variable
	switch len(vSlices) {
	case 0:
		varSlice = p.Variables()
	case 1:
		varSlice = vSlices[0]
	default:
		panic(fmt.Errorf("Too many inputs provided to LinearCoeff() method."))
	}

	// Constants

	// If there are no variables in the slice, then return a vector of length 1 containing zero.
	if len(varSlice) == 0 {
		panic(smErrors.CanNotGetLinearCoeffOfConstantError{p})
	}

	// Algorithm
	coeffOut := ZerosVector(len(varSlice))
	for ii := 0; ii < len(varSlice); ii++ {
		// Try to find the variable in the polynomial
		varIndex := p.VariableMonomialIndex(varSlice[ii])
		if varIndex != -1 {
			coeffOut.SetVec(ii, p.Monomials[varIndex].Coefficient)
		}
	}

	return coeffOut
}

/*
IsConstant
Description:

	This method returns true if and only if the polynomial
	represented by pv is a constant (i.e., it contains no variables).
*/
func (p Polynomial) IsConstant() bool {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return len(p.Variables()) == 0
}
