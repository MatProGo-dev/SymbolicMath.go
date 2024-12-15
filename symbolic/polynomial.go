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
Copy
Description:

	Returns a deep copy of the polynomial.
*/
func (p Polynomial) Copy() Polynomial {
	out := Polynomial{
		Monomials: make([]Monomial, len(p.Monomials)),
	}
	copy(out.Monomials, p.Monomials)
	return out
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
		err = eAsE.Check()
		if err != nil {
			panic(err)
		}

		//err := smErrors.CheckDimensionsInAddition(p, eAsE)
		//if err != nil {
		//	panic(err)
		//}
	}

	// Constants
	switch right := e.(type) {
	case float64:
		return p.Plus(K(right))
	case K:
		pCopy := p.Copy()

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
		pCopy := p.Copy()

		// Check to see if the variable is already in the polynomial
		variableIndex := pCopy.VariableMonomialIndex(right)
		if variableIndex == -1 {
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

		return pCopy

	case Monomial:
		return p.Plus(right.ToPolynomial())

	case Polynomial:
		pCopy := p.Copy()

		// Combine the list of monomials.
		pCopy.Monomials = append(pCopy.Monomials, right.Monomials...)

		// Simplify?
		return pCopy.Simplify()
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		ve, _ := ToVectorExpression(right)
		if ve.Len() == 1 {
			return p.Plus(ve.AtVec(0)) // Reuse scalar case
		} else {
			// Return a polynomial vector
			var polVecOut PolynomialVector
			for ii := 0; ii < ve.Len(); ii++ {
				polVecOut = append(polVecOut, p.Plus(ve.AtVec(ii)).(Polynomial))
			}
			return polVecOut
		}

	case KMatrix, VariableMatrix, MonomialMatrix, PolynomialMatrix:
		// Setup

		// Convert right to as ME
		rightAsME, _ := ToMatrixExpression(right)
		nResultRows, nResultCols := rightAsME.Dims()[0], rightAsME.Dims()[1]

		switch {
		case nResultRows == 1 && nResultCols == 1:
			return p.Plus(rightAsME.At(0, 0)) // Reuse scalar case
		default:
			// Return a polynomial matrix
			var polMatOut PolynomialMatrix
			for ii := 0; ii < nResultRows; ii++ {
				var polRowOut []Polynomial
				for jj := 0; jj < nResultCols; jj++ {
					polRowOut = append(polRowOut, p.Plus(rightAsME.At(ii, jj)).(Polynomial))
				}
				polMatOut = append(polMatOut, polRowOut)
			}
			return polMatOut
		}
	}

	// Unrecognized response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Polynomial.Plus",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	Defines a subtraction between the polynomial and another expression.
*/
func (p Polynomial) Minus(e interface{}) Expression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(err)
		}

		//// Check the dimensions in this subtraction
		//err := smErrors.CheckDimensionsInSubtraction(p, eAsE)
		//if err != nil {
		//	panic(err)
		//}

		// Use Expression's Minus() method
		return Minus(p, eAsE)
	}

	// If the function has reached this point, then
	// the input is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Polynomial.Minus",
			Input:        e,
		},
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
		err = eAsE.Check()
		if err != nil {
			panic(err)
		}

		//err := smErrors.CheckDimensionsInMultiplication(p, eAsE)
		//if err != nil {
		//	panic(err)
		//}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return p.Multiply(K(right))
	case K:
		pCopy := p.Copy()
		for ii, _ := range pCopy.Monomials {
			product_ii := pCopy.Monomials[ii].Multiply(right)
			pCopy.Monomials[ii] = product_ii.(Monomial) // Convert to Monomial
		}
		return pCopy
	case Variable:
		pCopy := p.Copy()
		for ii, _ := range pCopy.Monomials {
			product_ii := pCopy.Monomials[ii].Multiply(right)
			pCopy.Monomials[ii] = product_ii.(Monomial) // Convert to Monomial
		}
		return pCopy
	case Monomial:
		pCopy := p.Copy()
		var out Polynomial
		for _, m := range pCopy.Monomials {
			out.Monomials = append(
				out.Monomials,
				m.Multiply(right).(Monomial),
			)
		}
		return out
	case Polynomial:
		pCopy := p.Copy()

		// Multiply each monomial of the polynomial by the polynomial
		var productOut Expression = K(0.0)
		for ii := 0; ii < len(right.Monomials); ii++ {
			fmt.Println(fmt.Sprintf("pCopy.Multiply(right.Monomials[ii]): %v", pCopy.Multiply(right.Monomials[ii])))
			productOut = productOut.Plus(
				pCopy.Multiply(right.Monomials[ii]),
			)
		}

		return productOut.(Polynomial).Simplify()
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		// Right must be a vector of length 1
		ve, _ := ToVectorExpression(right)
		return ve.Multiply(p) // Reuse scalar case
	case KMatrix, VariableMatrix, MonomialMatrix, PolynomialMatrix:
		// Right must be a matrix of size [1,1]
		me, _ := ToMatrixExpression(right)
		return me.Multiply(p) // Reuse scalar case
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
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	return p.Copy()
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

	if IsExpression(rightIn) {
		// Check right
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Dimensions don't need to be checked here, because the polynomial is a scalar.
		//err := smErrors.CheckDimensionsInComparison(p, rightAsE, sense)
		//if err != nil {
		//	panic(err)
		//}
	}

	// Algorithm
	switch right := rightIn.(type) {
	case float64:
		return p.Comparison(K(right), sense)
	case K:
		return ScalarConstraint{p, right, sense}
	case Variable:
		return ScalarConstraint{p, right, sense}
	case Monomial:
		return ScalarConstraint{p, right, sense}
	case Polynomial:
		return ScalarConstraint{p, right, sense}
	}

	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Polynomial.Comparison (" + sense.String() + ")",
			Input:        rightIn,
		},
	)
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
	by finding the matching terms (i.e., monomials with matching Variables and Exponents)
	and combining them.
*/
func (p Polynomial) Simplify() Polynomial {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Find first element that has nonzero coefficient
	firstNonZeroIndex := -1
	for ii, monomial := range p.Monomials {
		if monomial.Coefficient != 0.0 {
			firstNonZeroIndex = ii
			break
		}
	}

	if len(p.Monomials) == 0 || firstNonZeroIndex == -1 {
		// If there is only a single, zero monomial, then return the zero polynomial
		return p
	}

	// Copy the first element of the polynomial into the new polynomial
	pCopy := Polynomial{
		Monomials: []Monomial{p.Monomials[firstNonZeroIndex]},
	}

	// Loop through the rest of the monomials
	for ii := firstNonZeroIndex + 1; ii < len(p.Monomials); ii++ {
		// Check to see if the monomials coefficient is zero
		if p.Monomials[ii].Coefficient == 0.0 {
			// Don't add it.
			continue
		}

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
		dMonomial := monomial.DerivativeWrt(vIn)
		switch component := dMonomial.(type) {
		case Monomial:
			derivative.Monomials = append(derivative.Monomials, component)
		case K:
			// Skip zero monomials
			if float64(component) == 0.0 {
				continue
			}
			derivative.Monomials = append(derivative.Monomials, component.ToMonomial())
		default:
			panic(fmt.Errorf("Unexpected type in Polynomial.Derivative: %T", component))
		}
	}

	// If the derivative is empty, then return 0.0
	if len(derivative.Monomials) == 0 {
		return K(0.0)
	}

	return derivative
}

/*
Degree
Description:

	The degree of the polynomial is the maximum degree of any of the monomials.
*/
func (p Polynomial) Degree() int {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	degree := 0
	for _, monomial := range p.Monomials {
		if monomial.Degree() > degree {
			degree = monomial.Degree()
		}
	}
	return degree
}

/*
LinearCoeff
Description:

	This function returns a vector describing the coefficients of the linear component
	of the polynomial.
	The (ii)th element of the vector is the coefficient of the (ii)th variable in the
	p.Variables() slice as it appears in the polynomial.
*/
func (p Polynomial) LinearCoeff(wrt ...[]Variable) mat.VecDense {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Check to see if the user provided a slice of variables
	var wrtVars []Variable
	switch len(wrt) {
	case 0:
		wrtVars = p.Variables()
	case 1:
		wrtVars = wrt[0]
	default:
		panic(fmt.Errorf("Too many inputs provided to LinearCoeff() method."))
	}

	// Constants

	// If there are no variables in the slice, then return a vector of length 1 containing zero.
	if len(wrtVars) == 0 {
		panic(smErrors.CanNotGetLinearCoeffOfConstantError{p})
	}

	// Algorithm
	coeffOut := ZerosVector(len(wrtVars))
	for ii := 0; ii < len(wrtVars); ii++ {
		// Try to find the variable in the polynomial
		varIndex := p.VariableMonomialIndex(wrtVars[ii])
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

/*
String
Description:

	This method returns a string representation of the polynomial.
*/
func (p Polynomial) String() string {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	// Create string
	polynomialString := ""

	// Add monomials
	for ii, monomial := range p.Monomials {
		if ii != 0 {
			polynomialString += " + "
		}
		polynomialString += monomial.String()
	}

	// Return
	return polynomialString
}

/*
Substitute
Description:

	This method substitutes the variable vIn with the expression eIn.
*/
func (p Polynomial) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	// Input Processing
	err := p.Check()
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
	var out Expression = K(0.0)
	for _, monomial := range p.Monomials {
		newMonomial := monomial.Substitute(vIn, eIn)
		out = out.Plus(newMonomial).(Polynomial).Simplify()
	}

	return out
}

/*
SubstituteAccordingTo
Description:

	This method substitutes the variables in the map with the corresponding expressions.
*/
func (p Polynomial) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	err = CheckSubstitutionMap(subMap)
	if err != nil {
		panic(err)
	}

	// Algorithm
	var out Expression = K(0.0)
	for _, monomial := range p.Monomials {
		newMonomial := monomial.SubstituteAccordingTo(subMap)
		out = out.Plus(newMonomial)
	}

	return out
}

/*
Power
Description:

	Computes the power of the constant.
*/
func (p Polynomial) Power(exponent int) Expression {
	return ScalarPowerTemplate(p, exponent)
}

/*
At
Description:

	Returns the value at the given row and column index.

Note:

	For a constant, the indices should always be 0.
*/
func (p Polynomial) At(ii, jj int) ScalarExpression {
	// Input Processing
	err := p.Check()
	if err != nil {
		panic(err)
	}

	// Check to see whether or not the index is valid.
	err = smErrors.CheckIndexOnMatrix(ii, jj, p)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return p
}
