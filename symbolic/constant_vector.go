package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
vector_constant_test.go
Description:
	Creates a vector extension of the constant type K from the original goop.
*/

// KVector A slice of constant (K) values, to be used as a VectorExpression. (and also an Expression)
type KVector []K // Inherit all methods from mat.VecDense

// Len Computes the length of the KVector given.
func (kv KVector) Len() int {
	return len(kv)
}

// Check This method is used to make sure that the variable is well-defined.
// For a constant vector, the vecdense should always be well-defined.
func (kv KVector) Check() error {
	return nil
}

// At This function returns the value at the ii, jj index.
// Note:
//
// For a constant vector, the jj index should always be 0.
func (kv KVector) At(ii, jj int) ScalarExpression {
	// Input Processing

	// Check to see whether or not the index is valid.
	err := smErrors.CheckIndexOnMatrix(ii, jj, kv)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return kv.AtVec(ii)

}

// AtVec This function returns the value at the k index.
func (kv KVector) AtVec(idx int) ScalarExpression {
	// Input Processing

	// Check to see whether or not the index is valid.
	err := smErrors.CheckIndexOnMatrix(idx, 0, kv)
	if err != nil {
		panic(err)
	}

	// Algorithm
	kvAsVector := kv.ToVecDense()
	return K(kvAsVector.AtVec(idx))
}

// Variables This function returns the empty slice because no variables are in a constant vector.
func (kv KVector) Variables() []Variable {
	return []Variable{}
}

// LinearCoeff This function returns a slice of the coefficients in the expression. For constants, this is always nil.
func (kv KVector) LinearCoeff(wrt ...[]Variable) mat.Dense {
	return PolynomialLikeVector_SharedLinearCoeffCalc(kv, wrt...)
}

// Constant Returns the constant additive value in the expression. For constants, this is just the constants value
func (kv KVector) Constant() mat.VecDense {
	return kv.ToVecDense()
}

// Plus Adds the current expression to another and returns the resulting expression
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

	// Algorithm
	var out Expression
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
		out = kv.Plus(VecDenseToKVector(eAsVec))
	case Expression:
		out = VectorPlusTemplate(kv, right)

	case *mat.VecDense:
		out = kv.Plus(VecDenseToKVector(*right)) // Convert to KVector
	case mat.VecDense:
		out = kv.Plus(VecDenseToKVector(right)) // Convert to KVector

	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "KVector.Plus",
				Input:        rightIn,
			},
		)
	}

	// Simplify and return
	return out.AsSimplifiedExpression()
}

// Minus Subtracts the current expression from another and returns the resulting expression
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
	case K, Variable, Monomial, Polynomial:
		rightAsSE := right.(ScalarExpression)
		return kv.Plus(rightAsSE.Multiply(-1.0)) // Reuse K case
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		// Force the right hand side to be a VectorExpression
		rhsAsVE := right.(VectorExpression)

		// Compute Subtraction using our Multiply method
		return kv.Plus(
			rhsAsVE.Multiply(-1.0),
		)
	}

	// Default response is a panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KVector.Minus",
			Input:        e,
		},
	)
}

// LessEq Returns a less than or equal to (<=) constraint between the current expression and another
func (kv KVector) LessEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseLessThanEqual)
}

// GreaterEq This method returns a greater than or equal to (>=) constraint between the current expression and another
func (kv KVector) GreaterEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseGreaterThanEqual)
}

// Eq This method returns an equality (==) constraint between the current expression and another
func (kv KVector) Eq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseEqual)
}

// Comparison creates a constraint comparing the KVector with the given
// expression in the sense provided by sense.
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
	case float64:
		var rhsAsVector mat.VecDense
		onesVector := OnesVector(kv.Len())
		rhsAsVector.ScaleVec(rhsConverted, &onesVector)
		return kv.Comparison(rhsAsVector, sense)
	case int:
		return kv.Comparison(float64(rhsConverted), sense)
	case K:
		return kv.Comparison(float64(rhsConverted), sense)
	case mat.VecDense:
		// Use KVector's Comparison method
		return kv.Comparison(VecDenseToKVector(rhsConverted), sense)

	case *mat.VecDense:
		// Use KVector's Comparison method
		return kv.Comparison(VecDenseToKVector(*rhsConverted), sense)

	case KVector, VariableVector, MonomialVector, PolynomialVector:
		// Pass the rhsConverted object into a container marked as a "VectorExpression" interface
		rhsAsVE := rhsConverted.(VectorExpression)

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsAsVE,
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

// Multiply returns the product of the input vector constant with another term.
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
	var out Expression
	switch right := rightIn.(type) {
	case float64:
		// If the input is a float64, then use the K method
		out = kv.Multiply(K(right))
	case K:
		// Iterate through each element
		var prod []ScalarExpression
		for i := 0; i < nResultRows; i++ {
			prod = append(prod, kv[i].Multiply(right).(ScalarExpression))
		}

		out = ConcretizeExpression(prod)
	case Variable:
		// Iterate through each element
		var prod []ScalarExpression
		for i := 0; i < nResultRows; i++ {
			prod = append(prod, kv[i].Multiply(right).(ScalarExpression))
		}

		out = ConcretizeExpression(prod)
	case Monomial:
		// Iterate through each element
		var prod []ScalarExpression
		for i := 0; i < nResultRows; i++ {
			prod = append(prod, kv[i].Multiply(right).(ScalarExpression))
		}

		out = ConcretizeExpression(prod)
	case Polynomial:
		// Iterate through each element
		var prod []ScalarExpression
		for i := 0; i < nResultRows; i++ {
			prod = append(prod, kv[i].Multiply(right).(ScalarExpression))
		}
		out = ConcretizeExpression(prod)
	case *mat.VecDense:
		out = kv.Multiply(*right)
	case mat.VecDense:
		// If the input is a vector, then use KVector method
		out = kv.Multiply(VecDenseToKVector(right))

	case KVector:
		// If the input is a KVector and the dimensions match
		// then right is a (1x1) vector and can use the scalar method.
		out = kv.Multiply(right[0])

	case VariableVector:
		// If the input is a KVector and the dimensions match
		// then right is a (1x1) vector and can use the scalar method.
		out = kv.Multiply(right[0])

	default:
		// Panic if the input type is not recognized
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "KVector.Multiply",
				Input:        rightIn,
			},
		)

	}

	// return
	return out.AsSimplifiedExpression()
}

// Transpose This method creates the transpose of the current vector and returns it.
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

// Dims Returns the dimension of the constant vector.
func (kv KVector) Dims() []int {
	return []int{kv.Len(), 1}
}

// Other Functions

// OnesVector Returns a vector of ones with length lengthIn.
// Note: this function assumes lengthIn is a positive number.
func OnesVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 1.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

// ZerosVector Returns a vector of zeros with length lengthIn.
// Note: this function assumes lengthIn is a positive number.
func ZerosVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 0.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

// DerivativeWrt Computes the derivative of the symbolic expression with respect to the
// variable vIn which should be a vector of all zeros.
func (kv KVector) DerivativeWrt(vIn Variable) Expression {
	return VecDenseToKVector(ZerosVector(kv.Len()))
}

// String Returns a string representation of the constant vector.
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

// ToVecDense This method converts the KVector to a mat.VecDense.
func (kv KVector) ToVecDense() mat.VecDense {
	dataIn := make([]float64, kv.Len())
	for ii, tempK := range kv {
		dataIn[ii] = float64(tempK)
	}
	return *mat.NewVecDense(len(kv), dataIn)
}

// VecDenseToKVector This method converts the mat.VecDense to a KVector.
func VecDenseToKVector(v mat.VecDense) KVector {
	out := make([]K, v.Len())
	for ii := 0; ii < v.Len(); ii++ {
		out[ii] = K(v.AtVec(ii))
	}
	return out
}

// ToMonomialVector This function converts the input expression to a monomial vector.
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

// ToPolynomialVector This function converts the input expression to a polynomial vector.
func (kv KVector) ToPolynomialVector() PolynomialVector {
	return kv.ToMonomialVector().ToPolynomialVector()
}

/*
ToKMatrix
Description:
*/

// Degree The degree of a constant matrix is always 0.
func (kv KVector) Degree() int {
	return 0
}

// Substitute Substitutes all occurrences of variable vIn with the expression eIn.
func (kv KVector) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	return kv
}

// SubstituteAccordingTo Substitutes all occurrences of the variables in the map with the corresponding expressions.
func (kv KVector) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	return kv
}

// Power Raises the scalar expression to the power of the input integer.
func (kv KVector) Power(exponent int) Expression {
	return VectorPowerTemplate(kv, exponent)
}

// AsSimplifiedExpression Returns the simplest form of the expression.
func (kv KVector) AsSimplifiedExpression() Expression {
	return kv
}

// ToScalarExpressions Converts the KVector into a slice of ScalarExpression type objects.
func (kv KVector) ToScalarExpressions() []ScalarExpression {
	var out []ScalarExpression
	for _, k := range kv {
		out = append(out, k)
	}
	return out
}
