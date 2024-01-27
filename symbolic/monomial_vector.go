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
		term1AsE, err := ToExpression(term1)
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = CheckDimensionsInAddition(mv, term1AsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return mv.Plus(K(right))
	case K:
		// Create a polynomial vector
		var pv PolynomialVector
		for _, monomial := range mv {
			pv = append(pv, monomial.Plus(right).(Polynomial))
		}
		return pv
	case MonomialVector:
		// Create a polynomial vector
		var pv PolynomialVector
		for ii, monomial := range mv {
			var tempPolynomial Polynomial
			tempPolynomial.Monomials = append(tempPolynomial.Monomials, monomial, mv[ii])
			pv = append(pv, tempPolynomial.Simplify())
		}
		return pv
	}

	// Unrecognized response is a panic
	panic(
		fmt.Errorf("Unexpected type of term1 in the MonomialVector.Plus() method: %T (%v)", term1, term1),
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
		term1AsE, err := ToExpression(term1)
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = CheckDimensionsInMultiplication(mv, term1AsE)
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
		rhs, err := ToExpression(rightIn)
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
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Comparison",
				Input:        rhs,
			},
		)
	}

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
	var monomialPresent bool
	for _, monomial := range mv {
		DmDv := monomial.DerivativeWrt(v)
		if _, tf := DmDv.(Monomial); !tf {
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
		return VecDenseToKVector(OnesVector(mv.Len()))
	}
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
