package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
polynomial_vector.go
Description:
	Defines a vector of polynomials.
*/

type PolynomialVector []Polynomial

// =========
// Functions
// =========

/*
Check
Description:

	Verifies that each of the polynomials in the vector are valid.
*/
func (pv PolynomialVector) Check() error {
	// Check that the polynomial has at least one monomial
	if len(pv) == 0 {
		return fmt.Errorf("polynomial vector has no polynomials")
	}

	// Check that each of the monomials are well formed
	for ii, polynomial := range pv {
		err := polynomial.Check()
		if err != nil {
			return fmt.Errorf("error in polynomial %v: %v", ii, err)
		}
	}

	// All checks passed
	return nil
}

/*
Length
Description:

	The number of elements in the Polynomial vector.
*/
func (pv PolynomialVector) Length() int {
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	return len(pv)
}

/*
Len
Description:

	Mirrors the gonum api for vectors. This extracts the element of the variable vector at the index x.
*/
func (pv PolynomialVector) Len() int {
	return pv.Length()
}

/*
At
Description:

	Returns the polynomial at the (ii,jj) index.

Note:

  - The jj index should always be 0.
*/
func (pv PolynomialVector) At(ii, jj int) ScalarExpression {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	err = smErrors.CheckIndexOnMatrix(ii, jj, pv)
	if err != nil {
		panic(err)
	}

	// Return the polynomial at the index
	return pv[ii]
}

/*
AtVec
Description:

	Retrieves the polynomial at the index idx.
*/
func (pv PolynomialVector) AtVec(idx int) ScalarExpression {
	// Input Checking
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	err = smErrors.CheckIndexOnVector(idx, pv)
	if err != nil {
		panic(err)
	}

	//
	return pv[idx]
}

/*
Variables
Description:

	Retrieves the set of all unique variables in the polynomial vector.
*/
func (pv PolynomialVector) Variables() []Variable {
	var variables []Variable // The variables in the polynomial
	for _, polynomial := range pv {
		variables = append(variables, polynomial.Variables()...)
	}
	return UniqueVars(variables)
}

/*
Constant
Description:

	Returns all of the constant components of the polynomial vector.
*/
func (pv PolynomialVector) Constant() mat.VecDense {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var constant mat.VecDense = ZerosVector(pv.Len())

	// Algorithm
	for ii, polynomial := range pv {
		constant.SetVec(ii, polynomial.Constant())
	}
	return constant
}

/*
LinearCoeff
Description:

	Retrieves the coefficients of the linear terms in the polynomial vector.
	The output is a matrix where element (ii,jj) of the matrix describes the coefficient
	of variable jj (from pv.Variables()) in the polynomial at index ii.
*/
func (pv PolynomialVector) LinearCoeff(wrt ...[]Variable) mat.Dense {
	return PolynomialLikeVector_SharedLinearCoeffCalc(pv, wrt...)
}

/*
Plus
Description:

	Defines an addition between the polynomial vector and another expression.
*/
func (pv PolynomialVector) Plus(e interface{}) Expression {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Plus: %v", err))
		}

		err := smErrors.CheckDimensionsInAddition(pv, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := e.(type) {
	case float64:
		return pv.Plus(K(right))
	case K:
		pvCopy := pv

		for ii, polynomial := range pv {
			sum := polynomial.Plus(right)
			pvCopy[ii] = sum.(Polynomial)
		}
		return pvCopy
	case Variable:
		pvCopy := pv
		for ii, polynomial := range pv {
			sum := polynomial.Plus(right)
			pvCopy[ii] = sum.(Polynomial)
		}
		return pvCopy

	case Polynomial:
		pvCopy := pv

		// Algorithm
		for ii, polynomial := range pv {
			sum := polynomial.Plus(right)
			pvCopy[ii] = sum.(Polynomial)
		}
		return pvCopy
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		pvCopy := pv

		// Cast right
		rightAsVector, _ := ToVectorExpression(right)

		// Algorithm
		for ii, polynomial := range pv {
			sum := polynomial.Plus(rightAsVector.AtVec(ii))
			pvCopy[ii] = sum.(Polynomial)
		}
		return pvCopy.Simplify()
	}

	// Default response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "PolynomialVector.Plus",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	Defines a subtraction between the polynomial vector and another expression.
*/
func (pv PolynomialVector) Minus(e interface{}) Expression {
	// Input Processing
	// - Check the polynomial vector
	// - Checks if the input e is an expression, if so:
	//	 + Checks the expression
	//	 + Checks the dimensions of the polynomial vector and the expression

	err := pv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Minus: %v", err))
		}

		err = smErrors.CheckDimensionsInSubtraction(pv, eAsE)
		if err != nil {
			panic(err)
		}

		// Use the Expression's Minus method
		return Minus(pv, eAsE)
	}

	// Default response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "PolynomialVector.Minus",
			Input:        e,
		},
	)
}

/*
Multiply
Description:

	Computes the product of a polynomial vector and another expression.
*/
func (pv PolynomialVector) Multiply(rightIn interface{}) Expression {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}
		err = smErrors.CheckDimensionsInMultiplication(pv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := rightIn.(type) {
	case float64:
		return pv.Multiply(K(right))
	case K:
		pvCopy := pv

		for ii, polynomial := range pv {
			product := polynomial.Multiply(right)
			pvCopy[ii] = product.(Polynomial)
		}
		return pvCopy
	case Polynomial:
		pvCopy := pv

		for ii, polynomial := range pv {
			product := polynomial.Multiply(right)
			pvCopy[ii] = product.(Polynomial)
		}
		return pvCopy
	case PolynomialVector:
		// This should only be true if the polynomial vector is actually a polynomial.
		// Convert it to a polynomial and do the multiplication as if it was with just the scalar.
		return pv.Multiply(right[0])

	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "PolynomialVector.Multiply",
				Input:        rightIn,
			},
		)
	}
}

/*
Transpose
Description:

	Computes the transpose of the polynomial vector.
*/
func (pv PolynomialVector) Transpose() Expression {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var pvT PolynomialMatrix = make([][]Polynomial, 1)
	pvT[0] = make([]Polynomial, pv.Len())
	for ii, polynomial := range pv {
		pvT[0][ii] = polynomial
	}
	return pvT
}

/*
Dims
Description:

	Returns the shape of the vector which should always be (pv.Len(), 1).
*/
func (pv PolynomialVector) Dims() []int {
	return []int{pv.Len(), 1}
}

/*
Comparison
Description:

	Creates the vector constraint between the polynomial vector pv and another
	expression according to the sense senseIn.
*/
func (pv PolynomialVector) Comparison(e interface{}, senseIn ConstrSense) Constraint {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Check the dimensions of the input
	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err := eAsE.Check()
		if err != nil {
			panic(err)
		}

		err = CheckDimensionsInComparison(pv, eAsE, senseIn)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := e.(type) {
	case float64:
		return pv.Comparison(K(right), senseIn)
	case K:
		// Convert the scalar to a scalar vector
		tempVD := OnesVector(pv.Len())
		tempVD.ScaleVec(float64(right), &tempVD)

		return VectorConstraint{
			LeftHandSide:  pv,
			RightHandSide: VecDenseToKVector(tempVD),
			Sense:         senseIn,
		}
	//case Polynomial:
	//	return VectorConstraint{
	//		Expression: pv.Plus(right.Multiply(K(-1))),
	//		Sense:      senseIn,
	//	}
	case *mat.VecDense:
		// Convert the vector to a KVector
		tempKVector := VecDenseToKVector(*right)
		return VectorConstraint{
			LeftHandSide:  pv,
			RightHandSide: tempKVector,
			Sense:         senseIn,
		}
	case mat.VecDense:
		// Convert the vector to a KVector
		tempKVector := VecDenseToKVector(right)
		return VectorConstraint{
			LeftHandSide:  pv,
			RightHandSide: tempKVector,
			Sense:         senseIn,
		}
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		rightAsVE, _ := ToVectorExpression(right)
		return VectorConstraint{
			LeftHandSide:  pv,
			RightHandSide: rightAsVE,
			Sense:         senseIn,
		}
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "PolynomialVector.Comparison",
				Input:        e,
			},
		)
	}
}

/*
LessEq
Description:

	Returns a vector constraint between pv and the input expression.
	Leverages the Comparison method.
*/
func (pv PolynomialVector) LessEq(e interface{}) Constraint {
	return pv.Comparison(e, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Returns a vector constraint comparing pv and the input expression according
	to the GreaterEq sense.
	Leverages the Comparison method.
*/
func (pv PolynomialVector) GreaterEq(e interface{}) Constraint {
	return pv.Comparison(e, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Returns a vector constraint comparing pv and the input expression according
	to the Eq sense.
	Leverages the Comparison method.
*/
func (pv PolynomialVector) Eq(e interface{}) Constraint {
	return pv.Comparison(e, SenseEqual)
}

/*
DerivativeWrt
Description:

	Returns the derivative of the polynomial vector with respect to the input variable.
*/
func (pv PolynomialVector) DerivativeWrt(vIn Variable) Expression {
	// Constants
	var derivative PolynomialVector = pv

	// Algorithm
	for ii, polynomial := range pv {
		derivative[ii] = polynomial.DerivativeWrt(vIn).(Polynomial)
	}

	return derivative
}

/*
IsConstantVector
Description:

	This method returns true if the polynomial vector is constant.
*/
func (pv PolynomialVector) IsConstantVector() bool {
	// Constants
	var isConstant bool = true

	// Algorithm
	for _, polynomial := range pv {
		isConstant = isConstant && polynomial.IsConstant()
	}

	return isConstant
}

/*
Simplify
Description:

	This method simplifies the polynomial vector.
*/
func (pv PolynomialVector) Simplify() VectorExpression {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Constants

	// Algorithm
	var simplified []ScalarExpression

	for ii, polynomial := range pv {
		simplifiedEntryII := polynomial.AsSimplifiedExpression()
		entryAsSE, ok := simplifiedEntryII.(ScalarExpression)
		if !ok {
			panic(
				fmt.Errorf(
					"error converting polynomial vector entry %v to a scalar expression during simplification",
					ii,
				),
			)
		}
		simplified = append(simplified, entryAsSE)
	}

	return ConcretizeVectorExpression(simplified)
}

func (pv PolynomialVector) AsSimplifiedExpression() Expression {
	return pv.Simplify()
}

/*
String
Description:

	Returns a string representation of the polynomial vector.
*/
func (pv PolynomialVector) String() string {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var output string = "PolynomialVector = ["
	for ii, polynomial := range pv {
		output += polynomial.String()
		if ii != len(pv)-1 {
			output += ", "
		}
	}
	output += "]"
	return output
}

/*
Degree
Description:

	Returns the maximum degree of any of the entries
	in the polynomial vector.
*/
func (pv PolynomialVector) Degree() int {
	// Input Processing
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var maxDegree int = 0

	// Algorithm
	for _, polynomial := range pv {
		if polynomial.Degree() > maxDegree {
			maxDegree = polynomial.Degree()
		}
	}

	return maxDegree
}

/*
Substitute
Description:

	Substitutes the variable vIn with the expression eIn in the polynomial vector.
*/
func (pv PolynomialVector) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	return VectorSubstituteTemplate(pv, vIn, eIn)
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the polynomial vector with the expressions in the map.
*/
func (pv PolynomialVector) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	// Input Checking
	err := pv.Check()
	if err != nil {
		panic(err)
	}

	err = CheckSubstitutionMap(subMap)
	if err != nil {
		panic(err)
	}

	// Algorithm
	var result VectorExpression = pv
	for tempVar, tempExpr := range subMap {
		resultSubbed := result.Substitute(tempVar, tempExpr.(ScalarExpression))
		result = resultSubbed.(VectorExpression)
	}

	return result
}

/*
Power
Description:

	Computes the power of the polynomial vector.
*/
func (pv PolynomialVector) Power(exponent int) Expression {
	return VectorPowerTemplate(pv, exponent)
}
