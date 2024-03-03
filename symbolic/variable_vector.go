package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
variable_vector.go
Description:
	The VariableVector type will represent a
*/

/*
VariableVector
Description:

	Represnts a variable in a optimization problem. The variable is
*/
type VariableVector []Variable

// =========
// Functions
// =========

/*
Length
Description:

	Returns the length of the vector of optimization variables.
*/
func (vv VariableVector) Length() int {
	return len(vv)
}

/*
Len
Description:

	This function is created to mirror the GoNum Vector API. Does the same thing as Length.
*/
func (vv VariableVector) Len() int {
	return vv.Length()
}

/*
At
Description:

	Mirrors the gonum api for vectors. This extracts the element of the variable vector at the index x.
*/
func (vv VariableVector) AtVec(idx int) ScalarExpression {
	// Constants

	// Algorithm
	return vv[idx]
}

/*
Variables
Description:

	Returns the slice of all variables in the vector.
*/
func (vv VariableVector) Variables() []Variable {
	return UniqueVars(vv)
}

/*
Constant
Description:

	Returns an all zeros vector as output from the method.
*/
func (vv VariableVector) Constant() mat.VecDense {
	zerosOut := ZerosVector(vv.Len())
	return zerosOut
}

/*
LinearCoeff
Description:

	Returns the matrix which is multiplied by Variables to get the current "expression".
	For a single vector, this is an identity matrix.
*/
func (vv VariableVector) LinearCoeff() mat.Dense {
	return Identity(vv.Len())
}

/*
Plus
Description:

	This member function computes the addition of the receiver vector var with the
	incoming vector expression ve.
*/
func (vv VariableVector) Plus(rightIn interface{}) Expression {
	// Constants
	// vvLen := vv.Len()

	// Processing Errors
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = smErrors.CheckDimensionsInAddition(vv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := rightIn.(type) {
	case *mat.VecDense:
		// Use KVector's method
		return vv.Plus(VecDenseToKVector(*right))
	case mat.VecDense:
		// Use KVector's method
		return vv.Plus(VecDenseToKVector(right))
	case KVector:
		// Create a polynomial vector
		var pv PolynomialVector
		for ii := 0; ii < vv.Len(); ii++ {
			var tempPolynomial Polynomial
			if right.AtVec(ii).(K) != 0 {
				rightAsK := right.AtVec(ii).(K)
				tempPolynomial.Monomials = append(tempPolynomial.Monomials, rightAsK.ToMonomial())
			}
			tempPolynomial.Monomials = append(
				tempPolynomial.Monomials,
				vv[ii].ToMonomial(),
			)
			// Create next polynomial.
			pv = append(pv, tempPolynomial)
		}
		return pv
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "VariableVector.Plus",
				Input:        right,
			},
		)
	}
}

/*
Minus
Description:

	This function subtracts an expression from the current
	variable vector and returns the resulting expression.
*/
func (vv VariableVector) Minus(rightIn interface{}) Expression {
	// Input Checking
	// - Check that the receiver is well-defined
	// - If the rightIn is an expression, then
	//	 + Check that the rightIn is well-defined
	//	 + Check that the dimensions are compatible

	err := vv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = smErrors.CheckDimensionsInSubtraction(vv, rightAsE)
		if err != nil {
			panic(err)
		}

		// Use Expression's Minus function
		return Minus(vv, rightAsE)
	}

	// Algorithm for non-expressions
	switch right := rightIn.(type) {
	case float64:
		return vv.Minus(K(right)) // Use K method
	case mat.VecDense:
		return vv.Minus(VecDenseToKVector(right)) // Use KVector method
	case *mat.VecDense:
		return vv.Minus(VecDenseToKVector(*right)) // Use KVector method
	}

	// If input isn't recognized, then panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Minus",
			Input:        rightIn,
		},
	)
}

/*
Multiply
Description:

	Multiplication of a VariableVector with another expression.
*/
func (vv VariableVector) Multiply(rightIn interface{}) Expression {
	//Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = smErrors.CheckDimensionsInMultiplication(vv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	nResultRows := vv.Dims()[0]

	// Algorithm
	switch right := rightIn.(type) {
	case float64:
		// Use K method
		return vv.Multiply(K(right))
	case K:
		// Is output a scalar
		if nResultRows == 1 {
			return vv[0].Multiply(right)
		}
		// Create a new vector of polynomials.
		var mvOut MonomialVector
		for _, v := range vv {
			mvOut = append(mvOut, v.Multiply(right).(Monomial))
		}
		return mvOut
	case Monomial:
		// Is output a scalar?
		if nResultRows == 1 {
			return vv[0].Multiply(right)
		}
		// Otherwise, create a new vector of monomials.
		var mvOut MonomialVector
		for _, v := range vv {
			mvOut = append(mvOut, v.Multiply(right).(Monomial))
		}
		return mvOut
	case Polynomial:
		// Is output a scalar?
		if nResultRows == 1 {
			return vv[0].Multiply(right)
		}
		// Create a new vector of polynomials.
		var pvOut PolynomialVector
		for _, v := range vv {
			pvOut = append(pvOut, v.Multiply(right).(Polynomial))
		}
		return pvOut
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		// Vector of polynomials must be (1x1)
		rightAsVE, _ := ToVectorExpression(right)
		return vv.Multiply(rightAsVE.AtVec(0))
	}

	// Otherwise, panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Multiply",
			Input:        rightIn,
		},
	)
}

/*
LessEq
Description:

	This method creates a less than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VariableVector) LessEq(rightIn interface{}) Constraint {
	return vv.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method creates a greater than or equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VariableVector) GreaterEq(rightIn interface{}) Constraint {
	return vv.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method creates an equal to vector constraint using the receiver as the left hand side and the
	input rhs as the right hand side if it is valid.
*/
func (vv VariableVector) Eq(rightIn interface{}) Constraint {
	return vv.Comparison(rightIn, SenseEqual)

}

/*
Comparison
Description:

	This method creates a constraint of type sense between
	the receiver (as left hand side) and rhs (as right hand side) if both are valid.
*/
func (vv VariableVector) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = CheckDimensionsInComparison(vv, rightAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	// Constants

	// Algorithm
	switch rhsConverted := rightIn.(type) {
	case mat.VecDense:
		rhsAsKVector := VecDenseToKVector(rhsConverted)

		return vv.Comparison(rhsAsKVector, sense)

	case KVector, VariableVector, MonomialVector, PolynomialVector:
		// Convert to a vector expression
		rightAsVE, _ := ToVectorExpression(rhsConverted)
		return VectorConstraint{vv, rightAsVE, sense}
	}

	// Default option is to panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Comparison",
			Input:        rightIn,
		},
	)
}

func (vv VariableVector) Copy() VariableVector {
	// Constants

	// Algorithm
	var newVarSlice VariableVector
	for varIndex := 0; varIndex < vv.Len(); varIndex++ {
		// Append to newVar Slice
		newVarSlice = append(newVarSlice, vv[varIndex])
	}

	return newVarSlice

}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (vv VariableVector) Transpose() Expression {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	// Create a matrix representing the transpose
	var vmOut VariableMatrix
	vmOut = append(vmOut, vv.Copy())
	return vmOut
}

/*
Dims
Description:

	Dimensions of the variable vector.
*/
func (vv VariableVector) Dims() []int {
	return []int{vv.Len(), 1}
}

/*
Check
Description:

	Checks whether or not the VariableVector has a sensible initialization.
*/
func (vv VariableVector) Check() error {
	// Check that each variable is properly defined
	for ii, element := range vv {
		err := element.Check()
		if err != nil {
			return fmt.Errorf(
				"element %v has an issue: %v",
				ii, err,
			)
		}
	}

	// If nothing was thrown, then return nil!
	return nil
}

/*
DerivativeWrt
Description:

	This function returns the derivative of the VariableVector with respect to the input variable
	vIn, which is a vector where each element:
		- is 0 if the variable is not the same as vIn
		- is 1 if the variable is the same as vIn
*/
func (vv VariableVector) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	vecOut := ZerosVector(vv.Len())

	// Algorithm
	for ii, element := range vv {
		if element.ID == vIn.ID {
			vecOut.SetVec(ii, 1)
		}
	}

	return VecDenseToKVector(vecOut)
}

/*
NewVariableVector
Description:

	Returns a new VariableVector object.
*/
func NewVariableVector(N int, envs ...*Environment) VariableVector {
	// Constants

	// Input Processing
	var currentEnv = &BackgroundEnvironment
	switch len(envs) {
	case 1:
		currentEnv = envs[0]
	}

	// Algorithm
	var varVectorOut VariableVector
	for ii := 0; ii < N; ii++ {
		varVectorOut = append(varVectorOut, NewVariable(currentEnv))
	}

	return varVectorOut

}

/*
String
Description:

	Returns a string representation of the VariableVector.
*/
func (vv VariableVector) String() string {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var output string = "VariableVector = ["
	for ii, variable := range vv {
		output += variable.String()
		if ii != len(vv)-1 {
			output += ", "
		}
	}
	output += "]"

	return output
}

/*
ToMonomialVector
Description:

	This function converts the input expression to a monomial vector.
*/
func (vv VariableVector) ToMonomialVector() MonomialVector {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var out MonomialVector

	// Algorithm
	for _, variable := range vv {
		out = append(out, variable.ToMonomial())
	}

	return out
}

/*
ToPolynomialVector
Description:

	This function converts the input expression to a polynomial vector.
*/
func (vv VariableVector) ToPolynomialVector() PolynomialVector {
	// Input Processing
	err := vv.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	var out PolynomialVector

	// Algorithm
	for _, variable := range vv {
		out = append(out, variable.ToPolynomial())
	}

	return out
}
