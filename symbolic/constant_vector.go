package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
vector_constant_test.go
Description:
	Creates a vector extension of the constant type K from the original goop.
*/

import (
	"fmt"
)

/*
KVector

	A type which is built on top of the KVector()
	a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
*/
type KVector []K // Inherit all methods from mat.VecDense

/*
Len

	Computes the length of the KVector given.
*/
func (kv KVector) Len() int {
	return len(kv)
}

/*
Check
Description:

	This method is used to make sure that the variable is well-defined.
	For a constant vector, the vecdense should always be well-defined.
*/
func (kv KVector) Check() error {
	return nil
}

/*
AtVec
Description:

	This function returns the value at the k index.
*/
func (kv KVector) AtVec(idx int) ScalarExpression {
	// Input Processing
	if idx < 0 || idx >= kv.Len() {
		// TODO: Create new error type to handle this; maybe put it in new package?
	}

	// Algorithm
	kvAsVector := kv.ToVecDense()
	return K(kvAsVector.AtVec(idx))
}

/*
Variables
Description:

	This function returns the empty slice because no variables are in a constant vector.
*/
func (kv KVector) Variables() []Variable {
	return []Variable{}
}

/*
LinearCoeff
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
func (kv KVector) LinearCoeff() mat.Dense {
	return ZerosMatrix(kv.Len(), kv.Len())
}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (kv KVector) Constant() mat.VecDense {
	return kv.ToVecDense()
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (kv KVector) Plus(rightIn interface{}) Expression {
	// Input Processing
	err := kv.Check()
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

		// Check dimensions
		err = smErrors.CheckDimensionsInAddition(kv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	kvLen := kv.Len()

	// Management
	switch right := rightIn.(type) {
	case float64:
		// Check to see if the output is a vector or a scalar
		if kvLen == 1 {
			return K(float64(kv[0]) + right)
		}

		// Create vector
		tempOnes := OnesVector(kvLen)
		var eAsVec mat.VecDense
		eAsVec.ScaleVec(right, &tempOnes)

		// Add the values
		return kv.Plus(VecDenseToKVector(eAsVec))
	case K:
		// Return Addition
		return kv.Plus(float64(right))
	case Variable:
		// Create a new polynomial vector
		var pvOut PolynomialVector
		for _, element := range kv {
			pvOut = append(pvOut, element.Plus(right).(Polynomial))
		}
		return pvOut

	case *mat.VecDense:
		return kv.Plus(VecDenseToKVector(*right)) // Convert to KVector

	case KVector:
		// Compute Addition
		var result mat.VecDense
		kvAsVec := kv.ToVecDense()
		eAsVec := right.ToVecDense()
		result.AddVec(&kvAsVec, &eAsVec)

		return VecDenseToKVector(result)

	case VariableVector:
		return right.Plus(kv)

	case MonomialVector:
		return right.Plus(kv)

	case PolynomialVector:
		return right.Plus(kv)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVector kv.Plus(%v)!", right, right)
		panic(fmt.Errorf(errString))
	}
}

/*
Minus
Description:

	Subtracts the current expression from another and returns the resulting expression
*/
func (kv KVector) Minus(e interface{}) Expression {
	// Input Processing
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		// Check right
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInSubtraction(kv, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return kv.Minus(K(right)) // Reuse K case
	case mat.VecDense:
		return kv.Minus(VecDenseToKVector(right)) // Convert to KVector
	case *mat.VecDense:
		return kv.Minus(VecDenseToKVector(*right)) // Convert to KVector
	}

	// Default response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KVector.Minus",
			Input:        e,
		},
	)
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kv KVector) LessEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kv KVector) GreaterEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kv KVector) Eq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseEqual)
}

func (kv KVector) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Checking
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		// Check dimensions
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = CheckDimensionsInComparison(kv, rightAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	switch rhsConverted := rightIn.(type) {
	case KVector:
		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsConverted,
			Sense:         sense,
		}

	case mat.VecDense:
		// Use KVector's Comparison method
		return kv.Comparison(VecDenseToKVector(rhsConverted), sense)

	case *mat.VecDense:
		// Use KVector's Comparison method
		return kv.Comparison(VecDenseToKVector(*rhsConverted), sense)

	case VariableVector:
		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsConverted,
			Sense:         sense,
		}

	default:
		// Return an error
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "KVector.Comparison",
				Input:        rightIn,
			},
		)

	}
}

/*
Multiply
Description:

	This method is used to compute the multiplication of the input vector constant with another term.
*/
func (kv KVector) Multiply(rightIn interface{}) Expression {
	// Input Processing
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		// Check rightIn
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check its dimensions
		err = smErrors.CheckDimensionsInMultiplication(kv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	nResultRows := kv.Len()

	// Compute Multiplication
	switch right := rightIn.(type) {
	case float64:
		// Is output a vector or a scalar?
		if nResultRows == 1 {
			return K(float64(kv[0]) * right)
		}

		// Use mat.Vector's multiplication method
		var result mat.VecDense
		kvAsVec := kv.ToVecDense()
		result.ScaleVec(right, &kvAsVec)

		return VecDenseToKVector(result)
	case K:
		// Convert to float64
		eAsFloat := float64(right)

		return kv.Multiply(eAsFloat)
	case Variable:
		// Is the output a vector or a scalar?
		if nResultRows == 1 {
			return right.Multiply(kv[0])
		}

		// Create a new monomial vector
		var mvOut MonomialVector
		for _, element := range kv {
			mvOut = append(mvOut, element.Multiply(right).(Monomial))
		}
		return mvOut
	case Monomial:
		// Is the output a vector or a scalar?
		if nResultRows == 1 {
			return right.Multiply(kv[0])
		}

		// Create a new monomial vector
		var mvOut MonomialVector
		for _, element := range kv {
			mvOut = append(mvOut, element.Multiply(right).(Monomial))
		}
		return mvOut
	case Polynomial:
		// Is the output a vector or a scalar?
		if nResultRows == 1 {
			return right.Multiply(kv[0])
		}

		// Create a new monomial vector
		var pvOut PolynomialVector
		for _, element := range kv {
			pvOut = append(pvOut, element.Multiply(right).(Polynomial))
		}
		return pvOut
	case *mat.VecDense:
		return kv.Multiply(*right)
	case mat.VecDense:
		// If the input is a vector, then use KVector method
		return kv.Multiply(VecDenseToKVector(right))

	case KVector:
		// If the input is a KVector and the dimensions match
		// then right is a (1x1) vector and can use the scalar method.
		return kv.Multiply(right[0])

	case VariableVector:
		// If the input is a KVector and the dimensions match
		// then right is a (1x1) vector and can use the scalar method.
		return kv.Multiply(right[0])

	}

	// If none of the above input types match, then panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KVector.Multiply",
			Input:        rightIn,
		},
	)
}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (kv KVector) Transpose() Expression {
	// Constants
	kvAsVD := kv.ToVecDense()
	kvLen := kv.Len()

	// Create empty matrix and populate
	kvT := ZerosMatrix(1, kvLen)
	for colIndex := 0; colIndex < kvLen; colIndex++ {
		kvT.Set(0, colIndex, kvAsVD.AtVec(colIndex))
	}

	return DenseToKMatrix(kvT)
}

/*
Dims
Description:

	Returns the dimension of the constant vector.
*/
func (kv KVector) Dims() []int {
	return []int{kv.Len(), 1}
}

// Other Functions

/*
OnesVector
Description:

	Returns a vector of ones with length lengthIn.
	Note: this function assumes lengthIn is a positive number.
*/
func OnesVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 1.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

/*
ZerosVector
Description:

	Returns a vector of zeros with length lengthIn.
	Note: this function assumes lengthIn is a positive number.
*/
func ZerosVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 0.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

/*
DerivativeWrt
Description:

	Computes the derivative of the symbolic expression with respect to the
	variable vIn which should be a vector of all zeros.
*/
func (kv KVector) DerivativeWrt(vIn Variable) Expression {
	return VecDenseToKVector(ZerosVector(kv.Len()))
}

/*
String
Description:

	Returns a string representation of the constant vector.
*/
func (kv KVector) String() string {
	// Constants
	lenKV := kv.Len()

	// Assemble string
	stringKV := "["
	for ii, tempK := range kv {
		stringKV += fmt.Sprintf("%v", tempK)
		if ii < lenKV-1 {
			stringKV += ", "
		}
	}
	stringKV += "]"

	return stringKV
}

/*
ToVecDense
Description:

	This method converts the KVector to a mat.VecDense.
*/
func (kv KVector) ToVecDense() mat.VecDense {
	dataIn := make([]float64, kv.Len())
	for ii, tempK := range kv {
		dataIn[ii] = float64(tempK)
	}
	return *mat.NewVecDense(len(kv), dataIn)
}

/*
VecDenseToKVector
Description:

	This method converts the mat.VecDense to a KVector.
*/
func VecDenseToKVector(v mat.VecDense) KVector {
	out := make([]K, v.Len())
	for ii := 0; ii < v.Len(); ii++ {
		out[ii] = K(v.AtVec(ii))
	}
	return out
}

/*
ToMonomialVector
Description:

	This function converts the input expression to a monomial vector.
*/
func (kv KVector) ToMonomialVector() MonomialVector {
	// Input Processing
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var mvOut MonomialVector
	for _, element := range kv {
		mvOut = append(mvOut, element.ToMonomial())
	}

	// Return
	return mvOut
}

/*
ToPolynomialVector
Description:

	This function converts the input expression to a polynomial vector.
*/
func (kv KVector) ToPolynomialVector() PolynomialVector {
	return kv.ToMonomialVector().ToPolynomialVector()
}

/*
Degree
Description:

	The degree of a constant matrix is always 0.
*/
func (kv KVector) Degree() int {
	return 0
}

/*
Substitute
Description:

	Substitutes all occurrences of variable vIn with the expression eIn.
*/
func (kv KVector) Substitute(vIn Variable, eIn Expression) Expression {
	return kv
}

/*
SubstituteAccordingTo
Description:

	Substitutes all occurrences of the variables in the map with the corresponding expressions.
*/
func (kv KVector) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	return kv
}

/*
Power
Description:

	Raises the scalar expression to the power of the input integer.
*/
func (kv KVector) Power(exponent int) Expression {
	return VectorPowerTemplate(kv, exponent)
}
