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

// MonomialVector is a vector of monomials.
type MonomialVector []Monomial

// Check This function checks that the monomial vector is valid.
// It does this by checking that each of the monomials in the vector are valid.
// And by checking that there is a non-zero number of them.
func (mv MonomialVector) Check() error {
	// Check that the polynomial has at least one monomial
	if len(mv) == 0 {
		return smErrors.EmptyVectorError{Expression: mv}
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

// Variables Returns the variables in the monomial vector.
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

// Len Returns the number of monomials in the vector.
func (mv MonomialVector) Len() int {
	return len(mv)
}

// Dims Returns the dimensions of the monomial vector.
func (mv MonomialVector) Dims() []int {
	return []int{mv.Len(), 1}
}

// Constant Returns the constant component of the monomial vector (if one exists).
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

// Plus This function returns the sum of the monomial vector and the input expression.
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
	var out Expression
	switch right := term1.(type) {
	case float64:
		out = mv.Plus(K(right))
	case int:
		out = mv.Plus(K(float64(right)))
	case Expression:
		out = VectorPlusTemplate(mv, right)
	default:
		// If the right hand side is an unsupported type, then panic
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Plus",
				Input:        term1,
			},
		)
	}

	// Return
	return out.AsSimplifiedExpression()
}

// Minus This function returns the difference of the monomial vector and the input expression.
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

// Multiply This function returns the product of the monomial vector and the input expression.
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
	var out Expression
	switch right := term1.(type) {
	case float64:
		out = mv.Multiply(K(right))
	case Expression:
		out = VectorMultiplyTemplate(mv, right)
	default:
		// Unrecognized response is a panic
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Multiply",
				Input:        right,
			},
		)
	}

	return out.AsSimplifiedExpression()

}

// Transpose This function returns the transpose of the monomial vector.
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

// LessEq This function creates a constraint that the monomial vector is less than or equal to the input expression.
func (mv MonomialVector) LessEq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseLessThanEqual)
}

// GreaterEq This function creates a constraint that the monomial vector is greater than or equal to the input expression.
func (mv MonomialVector) GreaterEq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseGreaterThanEqual)
}

// Eq This function creates a constraint that the monomial vector is equal to the input expression.
func (mv MonomialVector) Eq(rightIn interface{}) Constraint {
	return mv.Comparison(rightIn, SenseEqual)
}

// Comparison This function compares the monomial vector to the input expression in the sense provided by sense.
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

// DerivativeWrt This function returns the derivative of the monomial vector with respect to the input variable.
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

// At This function returns the value of the monomial vector at the
// (ii, jj)-th index.
// Note:
//
// Because this is a vector, the jj input should always be 0.
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

// AtVec This function returns the value of the monomial vector at the input vector.
func (mv MonomialVector) AtVec(idx int) ScalarExpression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return mv[idx]
}

// String returns a string representation of the monomial vector.
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

// IsConstant Determines whether or not an input object is a valid "MonomialVector" according to MatProInterface.
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

// ToPolynomialVector This function converts the input monomial vector to a polynomial vector.
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

// Degree Returns the MAXIMUM degree in the monomial vector.
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

// Substitute substitutes the input variable with the input scalar expression.
func (mv MonomialVector) Substitute(vIn Variable, seIn ScalarExpression) Expression {
	return VectorSubstituteTemplate(mv, vIn, seIn)
}

// SubstituteAccordingTo This function substitutes all instances of variables in the substitutions map with their corresponding expressions.
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

// Power This function raises the monomial vector to the input power.
func (mv MonomialVector) Power(exponent int) Expression {
	return VectorPowerTemplate(mv, exponent)
}

// LinearCoeff This function retrieves the "linear coefficient" of the monomial vector.
// In math, this is extracting the matrix A such that:
//
// mv' = L * v
//
// where:
// - v is the vector of variables for the monomial vector.
// - mv' is the monomial vector mv with ONLY the terms that have degree 1
func (mv MonomialVector) LinearCoeff(wrt ...[]Variable) mat.Dense {
	return PolynomialLikeVector_SharedLinearCoeffCalc(mv, wrt...)
}

// AsSimplifiedExpression Returns the simplest form of the expression.
func (mv MonomialVector) AsSimplifiedExpression() Expression {
	// Input Processing
	err := mv.Check()
	if err != nil {
		panic(err)
	}

	// Simplify each monomial in the vector
	var out []ScalarExpression
	for ii, monomial := range mv {
		simplified := monomial.AsSimplifiedExpression()
		simplifiedAsSE, tf := simplified.(ScalarExpression)
		if !tf {
			panic(fmt.Errorf("error simplifying monomial vector entry %v", ii))
		}
		// Add the simplified version of the monomial to the output
		out = append(out, simplifiedAsSE)
	}

	return ConcretizeVectorExpression(out)
}

// ToScalarExpressions Converts the MonomialVector into a slice of ScalarExpression type objects.
func (mv MonomialVector) ToScalarExpressions() []ScalarExpression {
	var out []ScalarExpression
	for _, monomial := range mv {
		out = append(out, monomial)
	}
	return out
}
