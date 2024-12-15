package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
monomial_vector.go
Description:
	This file defines all functions associated with a vector of monomial
	objects.
*/

// ===============
// Type Definition
// ===============

type MonomialVector []Monomial

/*
Check
Description:

	This function checks that the monomial vector is valid.
	It does this by checking that each of the monomials in the vector are valid.
	And by checking that there is a non-zero number of them.
*/
func (mv MonomialVector) Check() error {
	// Check that the polynomial has at least one monomial
	if len(mv) == 0 {
		return smErrors.EmptyVectorError{mv}
	}

	// Check that each of the monomials are well formed
	for ii, monomial := range mv {
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

	Returns the variables in the monomial vector.
*/
func (mv MonomialVector) Variables() []Variable {
	// Check that the monomial vector is well formed
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	variables := []Variable{}
	for _, monomial := range mv {
		variables = append(variables, monomial.Variables()...)
	}

	// Return
	return UniqueVars(variables)
}

/*
Len
Description:

	Returns the number of monomials in the vector.
*/
func (mv MonomialVector) Len() int {
	return len(mv)
}

/*
Dims
Description:

	Returns the dimensions of the monomial vector.
*/
func (mv MonomialVector) Dims() []int {
	return []int{mv.Len(), 1}
}

/*
Constant
Description:

	Returns the constant component of the monomial vector (if one exists).
*/
func (mv MonomialVector) Constant() mat.VecDense {
	// Check that the monomial vector is well formed
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	constant := ZerosVector(len(mv))
	for ii, monomial := range mv {
		if monomial.IsConstant() {
			// If this element is a constant,
			// then copy its coefficient into the constant vector.
			constant.SetVec(ii, monomial.Coefficient)

		}
	}

	// Return
	return constant
}

/*
Plus
Description:

	This function returns the sum of the monomial vector and the input expression.
*/
func (mv MonomialVector) Plus(term1 interface{}) Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(term1) {
		// Convert term1 to expression and check it
		term1AsE, _ := ToExpression(term1)
		err = term1AsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInAddition(mv, term1AsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return mv.Plus(K(right))
	case K:
		// Convert the scalar to a scalar vector
		tempVD := OnesVector(mv.Len())
		tempVD.ScaleVec(float64(right), &tempVD)

		return mv.Plus(VecDenseToKVector(tempVD))
	case Monomial:
		// Check to see if all elements of the monomial vector,
		// are all monomials like the input monomial.
		monomialVectorMatches := true
		for _, monomial := range mv {
			if !monomial.MatchesFormOf(right) {
				monomialVectorMatches = false
			}
		}

		if monomialVectorMatches {
			// If all elements of the monomial vector are monomials like the input monomial,
			// then simply add the coefficients. and return a monomial vector.
			var mvOut MonomialVector
			for _, monomial := range mv {
				mvOut = append(mvOut, monomial.Plus(right).(Monomial))
			}
			return mvOut
		} else {
			// Otherwise, create a polynomial vector
			var pv PolynomialVector
			for _, monomial := range mv {
				pv = append(pv, monomial.Plus(right).(Polynomial))
			}
			return pv.Simplify()
		}
	case KVector:
		if mv.IsConstant() {
			// If monomial vector is really a constant vector,
			// then don't convert down but simply update the coefficients.
			var kvOut KVector = VecDenseToKVector(mv.Constant())
			return kvOut.Plus(right)
		} else {
			// Create a polynomial vector
			var pv PolynomialVector
			for ii, monomial := range mv {
				pv = append(pv, monomial.Plus(right[ii]).(Polynomial))
			}
			return pv.Simplify()
		}
	case MonomialVector:
		// Check to see if all elements of the monomial vector,
		// are all monomials like the input monomial.
		monomialVectorMatches := true
		for ii, monomial := range mv {
			if !monomial.MatchesFormOf(right[ii]) {
				monomialVectorMatches = false
			}
		}

		if monomialVectorMatches {
			// If all elements of the monomial vector are monomials like the input monomial,
			// then simply add the coefficients. and return a monomial vector.
			var mvOut MonomialVector
			for ii, monomial := range mv {
				mvOut = append(mvOut, monomial.Plus(right[ii]).(Monomial))
			}
			return mvOut
		} else {
			// Otherwise, create a polynomial vector
			var pv PolynomialVector
			for ii, monomial := range mv {
				sumII := monomial.Plus(right[ii])
				switch sumII.(type) {
				case Monomial:
					pv = append(pv, sumII.(Monomial).ToPolynomial())
				case Polynomial:
					pv = append(pv, sumII.(Polynomial))
				default:
					panic(
						fmt.Errorf(
							"Unexpected type of sumII in the Plus() method: %T (%v)",
							sumII, sumII,
						),
					)

				}
			}
			return pv.Simplify()
		}
	}

	// Unrecognized response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "MonomialVector.Plus",
			Input:        term1,
		},
	)
}

/*
Minus
Description:

	This function returns the difference of the monomial vector and the input expression.
*/
func (mv MonomialVector) Minus(term1 interface{}) Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(term1) {
		// Convert term1 to expression and check it
		term1AsE, _ := ToExpression(term1)
		err := term1AsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInSubtraction(mv, term1AsE)
		if err != nil {
			panic(err)
		}

		// Use Expression's Minus function
		return Minus(mv, term1AsE)
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return mv.Minus(K(right))
	case mat.VecDense:
		return mv.Minus(VecDenseToKVector(right))
	case *mat.VecDense:
		return mv.Minus(VecDenseToKVector(*right))
	}

	// Unrecognized response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "MonomialVector.Minus",
			Input:        term1,
		},
	)
}

/*
Multiply
Description:

	This function returns the product of the monomial vector and the input expression.
*/
func (mv MonomialVector) Multiply(term1 interface{}) Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(term1) {
		// Convert term1 to expression and check it
		term1AsE, _ := ToExpression(term1)
		err := term1AsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInMultiplication(mv, term1AsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return mv.Multiply(K(right))
	case K:
		// Create a polynomial vector
		var mv MonomialVector
		for _, monomial := range mv {
			mv = append(mv, monomial.Multiply(right).(Monomial))
		}
		return mv
	}

	// Unrecognized response is a panic
	panic(
		fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1),
	)
}

/*
Transpose
Description:

	This function returns the transpose of the monomial vector.
*/
func (mv MonomialVector) Transpose() Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var mm MonomialMatrix
	mm = append(mm, make(MonomialVector, len(mv)))
	for ii, monomial := range mv {
		mm[0][ii] = monomial
	}

	// Return
	return mm
}

/*
LessEq
Description:

	This function creates a constraint that the monomial vector is less than or equal to the input expression.
*/
func (mv MonomialVector) LessEq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This function creates a constraint that the monomial vector is greater than or equal to the input expression.
*/
func (mv MonomialVector) GreaterEq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This function creates a constraint that the monomial vector is equal to the input expression.
*/
func (mv MonomialVector) Eq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseEqual)
}

/*
Comparison
Description:

	This function compares the monomial vector to the input expression in the sense provided by sense.
*/
func (mv MonomialVector) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		// Convert rhsIn to expression and check it
		rhs, _ := ToExpression(rightIn)
		err := rhs.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = CheckDimensionsInComparison(mv, rhs, sense)
		if err != nil {
			panic(err)
		}
	}

	// Constants

	// Algorithm
	switch rhs := rightIn.(type) {
	case float64:
		return mv.Comparison(K(rhs), sense)
	case K:
		// Convert the scalar to a scalar vector
		tempVD := OnesVector(mv.Len())
		tempVD.ScaleVec(float64(rhs), &tempVD)

		return VectorConstraint{
			LeftHandSide:  mv,
			RightHandSide: VecDenseToKVector(tempVD),
			Sense:         sense,
		}
	case KVector:
		return VectorConstraint{
			LeftHandSide:  mv,
			RightHandSide: rhs,
			Sense:         sense,
		}
	case VariableVector:
		return VectorConstraint{
			LeftHandSide:  mv,
			RightHandSide: rhs,
			Sense:         sense,
		}
	case MonomialVector:
		return VectorConstraint{
			LeftHandSide:  mv,
			RightHandSide: rhs,
			Sense:         sense,
		}
	}

	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "MonomialVector.Comparison (" + sense.String() + ")",
			Input:        rightIn,
		},
	)

}

/*
DerivativeWrt
Description:

	This function returns the derivative of the monomial vector with respect to the input variable.
*/
func (mv MonomialVector) DerivativeWrt(v Variable) Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var ev []Expression
	var monomialPresent bool = false
	for _, monomial := range mv {
		DmDv := monomial.DerivativeWrt(v)
		if _, tf := DmDv.(Monomial); tf {
			monomialPresent = true
		}
		ev = append(ev, monomial.DerivativeWrt(v))
	}

	if monomialPresent {
		var mv MonomialVector
		for _, element := range ev {
			switch e := element.(type) {
			case Monomial:
				mv = append(mv, e)
			case K:
				mv = append(mv, e.ToMonomial())
			default:
				panic(
					fmt.Errorf(
						"Unexpected type of element in the DerivativeWrt() method: %T (%v)",
						element, element,
					),
				)
			}
		}
		return mv
	} else {
		// Collect all coefficients
		var out KVector
		for _, element := range ev {
			out = append(out, element.(K))
		}
		return out
	}
}

/*
At
Description:

	This function returns the value of the monomial vector at the
	(ii, jj)-th index.

Note:

	Because this is a vector, the jj input should always be 0.
*/
func (mv MonomialVector) At(ii, jj int) ScalarExpression {
	// Input Processing

	// Check that the monomial vector is well formed
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Check that the input is valid
	err = smErrors.CheckIndexOnMatrix(ii, jj, mv)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return mv[ii]
}

/*
AtVec
Description:

	This function returns the value of the monomial vector at the input vector.
*/
func (mv MonomialVector) AtVec(idx int) ScalarExpression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return mv[idx]
}

/*
String
Description:

	This function returns a string representation of the monomial vector.
*/
func (mv MonomialVector) String() string {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var output string = "MonomialVector = ["
	for ii, monomial := range mv {
		output += monomial.String()
		if ii != len(mv)-1 {
			output += ", "
		}
	}
	output += "]"

	return output
}

/*
IsConstant
Description:

	Determines whether or not an input object is a valid "MonomialVector" according to MatProInterface.
*/
func (mv MonomialVector) IsConstant() bool {
	// Input Checking
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Check each element of mv
	for _, element := range mv {
		if !element.IsConstant() {
			return false
		}
	}

	// If all checks pass, then return true!
	return true
}

/*
ToPolynomialVector
Description:

	This function converts the input monomial vector to a polynomial vector.
*/
func (mv MonomialVector) ToPolynomialVector() PolynomialVector {
	// Input Checking
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var pv PolynomialVector
	for _, monomial := range mv {
		pv = append(pv, monomial.ToPolynomial())
	}

	// Return
	return pv
}

/*
Degree
Description:

	Returns the MAXIMUM degree in the monomial vector.
*/
func (mv MonomialVector) Degree() int {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	maxDegree := 0
	for _, m := range mv {
		degree := m.Degree()
		if degree > maxDegree {
			maxDegree = degree
		}
	}

	// Return
	return maxDegree
}

/*
Substitute
Description:

	This function substitutes the input variable with the input scalar expression.
*/
func (mv MonomialVector) Substitute(vIn Variable, seIn ScalarExpression) Expression {
	return VectorSubstituteTemplate(mv, vIn, seIn)
}

/*
SubstituteAccordingTo
Description:

	This function substitutes all instances of variables in the substitutions map with their corresponding expressions.
*/
func (mv MonomialVector) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	err = CheckSubstitutionMap(subMap)
	if err != nil {
		panic(err)
	}

	// Setup
	var out VectorExpression = mv
	for varKey, expr := range subMap {
		postSub := out.Substitute(varKey, expr.(ScalarExpression))
		out = postSub.(VectorExpression)
	}
	return out
}

/*
Power
Description:

	This function raises the monomial vector to the input power.
*/
func (mv MonomialVector) Power(exponent int) Expression {
	return VectorPowerTemplate(mv, exponent)
}
