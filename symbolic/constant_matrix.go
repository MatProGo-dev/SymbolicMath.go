package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
constant_matrix.go
Description:
	Defines all methods related to the constant matrix type.
*/

type KMatrix [][]K

/*
Check
Description:

	Checks to make sure that the constant is initialized properly.
	ConstantMatrix objects are always initialized properly, so this should always return
	no error.
*/
func (km KMatrix) Check() error {
	return nil
}

/*
ToDense
Description:

	Converts the constant matrix to a dense matrix.
*/
func (km KMatrix) ToDense() mat.Dense {
	// Input Checking
	err := km.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	nR, nC := km.Dims()[0], km.Dims()[1]

	// Construct Dense
	kmAsDense := mat.NewDense(nR, nC, make([]float64, nR*nC))
	for rIndex := 0; rIndex < nR; rIndex++ {
		for cIndex := 0; cIndex < nC; cIndex++ {
			kmAsDense.Set(rIndex, cIndex, float64(km[rIndex][cIndex]))
		}
	}

	// Return
	return *kmAsDense
}

/*
Variables
Description:

	There are no variables in the constant matrix.
*/
func (km KMatrix) Variables() []Variable {
	return []Variable{}
}

/*
Dims
Description:

	The dimensions of the given matrix.
*/
func (km KMatrix) Dims() []int {
	// Input Checking
	err := km.Check()
	if err != nil {
		panic(err)
	}
	return []int{len(km), len(km[0])}
}

/*
Plus
Description:

	Addition of the constant matrix with another expression.
*/
func (km KMatrix) Plus(e interface{}) Expression {
	// Input Processing
	err := km.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		rightAsE, _ := ToExpression(e)
		// Check expression
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}
		// Check Dimensions
		err = smErrors.CheckDimensionsInAddition(km, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	dims := km.Dims()
	nR, nC := dims[0], dims[1]

	switch right := e.(type) {
	case float64:
		// Create a matrix of all elements with value right
		ones := OnesMatrix(nR, nC)
		var rightAsDense mat.Dense
		rightAsDense.Scale(right, &ones)

		// Create copy of km as a dense matrix
		kmAsDense := km.ToDense()

		// Compute sum
		var sumAsDense mat.Dense
		sumAsDense.Add(&rightAsDense, &kmAsDense)

		return DenseToKMatrix(sumAsDense)

	case K:
		return km.Plus(float64(right)) // Reuse float64 case

	case Variable:
		// Create a matrix of variables where each element has
		// the value of the variable
		var rightAsVM VariableMatrix = make([][]Variable, nR)
		for rIndex := 0; rIndex < nR; rIndex++ {
			rightAsVM[rIndex] = make([]Variable, nC)
			for cIndex := 0; cIndex < nC; cIndex++ {
				rightAsVM[rIndex][cIndex] = right
			}
		}

		return km.Plus(rightAsVM) // Reuse VariableMatrix case

	case mat.Dense:
		return km.Plus(DenseToKMatrix(right)) // Reuse KMatrix case

	case *mat.Dense:
		return km.Plus(*right) // Reuse mat.Dense case

	case KMatrix:
		// Create the result matrix
		var result KMatrix = make([][]K, nR)
		for rIndex := 0; rIndex < nR; rIndex++ {
			result[rIndex] = make([]K, nC)
			for cIndex := 0; cIndex < nC; cIndex++ {
				result[rIndex][cIndex] = km[rIndex][cIndex] + right[rIndex][cIndex]
			}
		}
		return result

	case VariableMatrix:
		// Create the result matrix
		var result PolynomialMatrix = make([][]Polynomial, nR)
		for rIndex := 0; rIndex < nR; rIndex++ {
			result[rIndex] = make([]Polynomial, nC)
			for cIndex := 0; cIndex < nC; cIndex++ {
				result[rIndex][cIndex] = km[rIndex][cIndex].Plus(right[rIndex][cIndex]).(Polynomial)
				// Each addition should create a polynomial
			}
		}
		return result
	case PolynomialMatrix:
		return right.Plus(km) // Reuse PolynomialMatrix case
	}

	// If we reach this point, the input is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Plus",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	Subtraction of the constant matrix with another expression.
*/
func (km KMatrix) Minus(e interface{}) Expression {
	// Input Processing
	err := km.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		rightAsE, _ := ToExpression(e)

		// Check expression
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check Dimensions
		err = smErrors.CheckDimensionsInSubtraction(km, rightAsE)
		if err != nil {
			panic(err)
		}

		// Use Expression method
		return Minus(km, rightAsE)
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		return km.Minus(K(right)) // Reuse K case
	case mat.Dense:
		return km.Minus(DenseToKMatrix(right)) // Reuse KMatrix case
	case *mat.Dense:
		return km.Minus(*right) // Reuse mat.Dense case
	}

	// If we reach this point, the input is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Minus",
			Input:        e,
		},
	)
}

/*
Multiply
Description:

	Multiplication of the constant matrix with another expression.
*/
func (km KMatrix) Multiply(e interface{}) Expression {
	// Input Processing
	err := km.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		rightAsE, _ := ToExpression(e)

		// Check expressions
		err = rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInMultiplication(km, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	switch right := e.(type) {
	case float64:
		// Use gonum's built-in scale function
		kmAsDense := km.ToDense()
		var product mat.Dense
		product.Scale(right, &kmAsDense)

		return DenseToKMatrix(product)

	case K:
		return km.Multiply(float64(right)) // Reuse float64 case
	case Polynomial:
		// Choose the correct output type based on the size of km
		nR, nC := km.Dims()[0], km.Dims()[1]
		switch {
		case (nR == 1) && (nC == 1):
			// If the output is a scalar, return a scalar
			return km[0][0].Multiply(right)
		case nC == 1:
			// If the output is a vector, return a vector
			var outputVec PolynomialVector = make([]Polynomial, nR)
			for rIndex := 0; rIndex < nR; rIndex++ {
				outputVec[rIndex] = km[rIndex][0].Multiply(right.Copy()).(Polynomial)
			}
			return outputVec
		default:
			// If the output is a matrix, return a matrix
			var outputMat PolynomialMatrix = make([][]Polynomial, nR)
			for rIndex := 0; rIndex < nR; rIndex++ {
				outputMat[rIndex] = make([]Polynomial, nC)
				for cIndex := 0; cIndex < nC; cIndex++ {
					outputMat[rIndex][cIndex] = km[rIndex][cIndex].Multiply(right.Copy()).(Polynomial)
				}
			}
			return outputMat
		}

	case *mat.VecDense:
		// Use gonum's built-in multiplication function
		var product mat.VecDense
		kmAsDense := km.ToDense()
		product.MulVec(&kmAsDense, right)

		// Check output dimensions
		nOutput := product.Len()
		if nOutput == 1 {
			// If the output is a scalar, return a scalar
			return K(product.AtVec(0))
		} else {
			// Otherwsie return a KVector
			return VecDenseToKVector(product)
		}

	case VariableVector:
		// Choose the correct output type based on the size of km
		nR := km.Dims()[0]
		if nR == 1 {
			// If the output is a scalar, return a scalar
			var out Polynomial = K(0).ToPolynomial()
			for cIndex := 0; cIndex < len(right); cIndex++ {
				out = out.Plus(
					right[cIndex].Multiply(km[0][cIndex]),
				).(Polynomial)
			}
			return out
		} else {
			nC := km.Dims()[1]
			// If the output is a vector, return a vector
			var outputVec PolynomialVector = VecDenseToKVector(
				ZerosVector(nR),
			).ToPolynomialVector()
			for colIndex := 0; colIndex < nC; colIndex++ {
				kmAsDense := km.ToDense()
				tempCol := (&kmAsDense).ColView(colIndex)
				outputVec = outputVec.Plus(
					right[colIndex].Multiply(tempCol),
				).(PolynomialVector)
			}
			return outputVec
		}
	case *mat.Dense:
		// Check output dimensions
		nOutputR := km.Dims()[0]
		_, nOutputCols := right.Dims()

		kmAsDense := km.ToDense()
		var out mat.Dense
		out.Mul(&kmAsDense, right)

		switch {
		case nOutputR == 1 && nOutputCols == 1:
			// If the constant matrix is a scalar, return the scalar
			return K(out.At(0, 0))
		case nOutputCols == 1:
			// If the output is a vector, return a vector
			var outputVec KVector = make([]K, nOutputR)
			for rIndex := 0; rIndex < nOutputR; rIndex++ {
				outputVec[rIndex] = K(out.At(rIndex, 0))
			}
			return outputVec
		default:
			// If the output is a matrix, return a matrix
			return DenseToKMatrix(out)
		}
	case mat.Dense:
		// Use *mat.Dense method
		return km.Multiply(&right) // Reuse *mat.Dense case
	case KMatrix:
		return km.Multiply(right.ToDense()) // Reuse *mat.Dense case
	case VariableMatrix:
		return MatrixMultiplyTemplate(km, right)
	case MonomialMatrix:
		return MatrixMultiplyTemplate(km, right)
	case PolynomialMatrix:
		return MatrixMultiplyTemplate(km, right)
	}

	// If we reach this point, the input is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Multiply",
			Input:        e,
		},
	)
}

/*
Transpose
Description:

	Transposes the constant matrix and returns a new
	matrix.
*/
func (km KMatrix) Transpose() Expression {
	// Constants

	// Transpose KM after converting it to dense
	kmAsDense := km.ToDense()
	kmDTransposed := kmAsDense.T()

	// Copy
	nR, nC := kmDTransposed.Dims()
	kmT := ZerosMatrix(nR, nC)

	for rIndex := 0; rIndex < nR; rIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			kmT.Set(rIndex, colIndex, kmDTransposed.At(rIndex, colIndex))
		}
	}

	// Return
	return DenseToKMatrix(kmT)

}

/*
LessEq
Description:

	Returns a constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) LessEq(rightIn interface{}) Constraint {
	return km.Comparison(rightIn, SenseLessThanEqual)

}

/*
GreaterEq
Description:

	Returns a greater equal constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) GreaterEq(rightIn interface{}) Constraint {
	return km.Comparison(rightIn, SenseGreaterThanEqual)

}

/*
Eq
Description:

	Returns an equal constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) Eq(rightIn interface{}) Constraint {
	return km.Comparison(rightIn, SenseEqual)

}

/*
Comparison
Description:

	Returns a constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	//err := km.Check()
	//if err != nil {
	//	panic(err)
	//}

	if IsExpression(rightIn) {
		// Check rightIn
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = CheckDimensionsInComparison(km, rightAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := rightIn.(type) {
	case K:
		// Create a matrix of all elements with value right
		ones := OnesMatrix(km.Dims()[0], km.Dims()[1])
		var rightAsDense mat.Dense
		rightAsDense.Scale(float64(right), &ones)

		// Call the *mat.Dense case
		return km.Comparison(rightAsDense, sense)
	case *mat.Dense:
		return km.Comparison(*right, sense) // Call the mat.Dense case
	case mat.Dense:
		return km.Comparison(DenseToKMatrix(right), sense) // Call the KMatrix case
	case KMatrix, VariableMatrix, MonomialMatrix, PolynomialMatrix:
		// Convert to matrix expression
		rightAsME, _ := ToMatrixExpression(rightIn)

		// Return constraint
		return MatrixConstraint{
			LeftHandSide:  km,
			RightHandSide: rightAsME,
			Sense:         sense,
		}
	}

	// If we reach this point, the input is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Comparison (" + sense.String() + ")",
			Input:        rightIn,
		},
	)

}

/*
At
Description:

	Retrieves element at the specified indices.
*/
func (km KMatrix) At(i, j int) ScalarExpression {
	kmAsD := km.ToDense()
	return K(kmAsD.At(i, j))
}

/*
Constant
Description:

	Retrieves the constant component.
*/
func (km KMatrix) Constant() mat.Dense {
	return km.ToDense()
}

// Other Functions

/*
ZerosMatrix
Description:

	Returns a dense matrix of all zeros.
*/
func ZerosMatrix(nR, nC int) mat.Dense {
	// Create empty slice
	elts := make([]float64, nR*nC)
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			elts[rowIndex*nC+colIndex] = 0.0
		}
	}

	return *mat.NewDense(nR, nC, elts)
}

/*
OnesMatrix
Description:

	Returns a dense matrix of all ones.
*/
func OnesMatrix(nR, nC int) mat.Dense {
	// Create empty slice
	elts := make([]float64, nR*nC)
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			elts[rowIndex*nC+colIndex] = 1.0
		}
	}

	return *mat.NewDense(nR, nC, elts)
}

/*
Identity
Description:

	Returns a symmetric matrix that is the identity matrix.
	Note: this function assumes lengthIn is a positive number.
*/
func Identity(dim int) mat.Dense {
	// Create the empty matrix.
	zeroBase := ZerosMatrix(dim, dim)

	// Populate Diagonal
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		zeroBase.Set(rowIndex, rowIndex, 1.0)
	}

	return zeroBase
}

/*
DerivativeWrt
Description:

	Computes the derivative of the constant matrix with respect to the variable
	v. For a constant matrix, this should create a matrix of all zeros (ZerosMatrix).
*/
func (km KMatrix) DerivativeWrt(vIn Variable) Expression {
	return DenseToKMatrix(ZerosMatrix(km.Dims()[0], km.Dims()[1]))
}

/*
String
Description:

	Returns a string representation of the constant matrix.
*/
func (km KMatrix) String() string {
	// Constants
	nR, nC := km.Dims()[0], km.Dims()[1]

	// Construct String
	out := "["
	for ii, row := range km {
		out += "["
		for jj, tempK := range row {
			out += fmt.Sprintf("%v", tempK)
			if jj != nC-1 {
				out += ","
			}
		}
		// Add "]," and newline if not last row
		out += "]"
		if ii != nR-1 {
			out += ",\n ["
		} else {
			out += "]\n"
		}

	}

	// Return string
	return out
}

/*
DenseToKMatrix
Description:

	Converts a dense matrix to a KMatrix.
*/
func DenseToKMatrix(denseIn mat.Dense) KMatrix {
	// Constants
	nR, nC := denseIn.Dims()

	// Create KMatrix
	var km KMatrix = make([][]K, nR)

	// Populate
	for rIndex := 0; rIndex < nR; rIndex++ {
		km[rIndex] = make([]K, nC)
		for cIndex := 0; cIndex < nC; cIndex++ {
			km[rIndex][cIndex] = K(denseIn.At(rIndex, cIndex))
		}
	}

	// Return
	return km
}

/*
ToMonomialMatrix
Description:

	Converts the constant matrix to a monomial matrix.
*/
func (km KMatrix) ToMonomialMatrix() MonomialMatrix {
	// Constants
	nR, nC := km.Dims()[0], km.Dims()[1]

	// Create MonomialMatrix
	var mm MonomialMatrix = make([][]Monomial, nR)
	for rIndex := 0; rIndex < nR; rIndex++ {
		mm[rIndex] = make([]Monomial, nC)
		for cIndex := 0; cIndex < nC; cIndex++ {
			mm[rIndex][cIndex] = km[rIndex][cIndex].ToMonomial()
		}
	}

	// Return
	return mm
}

/*
ToPolynomialMatrix
Description:

	Converts the constant matrix to a polynomial matrix.
*/
func (km KMatrix) ToPolynomialMatrix() PolynomialMatrix {
	// Constants
	nR, nC := km.Dims()[0], km.Dims()[1]

	// Create PolynomialMatrix
	var pm PolynomialMatrix = make([][]Polynomial, nR)
	for rIndex := 0; rIndex < nR; rIndex++ {
		pm[rIndex] = make([]Polynomial, nC)
		for cIndex := 0; cIndex < nC; cIndex++ {
			pm[rIndex][cIndex] = km[rIndex][cIndex].ToPolynomial()
		}
	}

	// Return
	return pm
}

/*
Degree
Description:

	The degree of a constant matrix is always 0.
*/
func (km KMatrix) Degree() int {
	return 0
}

/*
Substitute
Description:

	Substitutes all occurrences of variable vIn with the expression eIn.
*/
func (km KMatrix) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	return km
}

/*
SubstituteAccordingTo
Description:

	Substitutes all occurrences of the variables in the map with the corresponding expressions.
*/
func (km KMatrix) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	return km
}

/*
Power
Description:

	Raises the constant matrix to the power of the input integer.
*/
func (km KMatrix) Power(exponent int) Expression {
	return MatrixPowerTemplate(km, exponent)
}

/*
AsSimplifiedExpression
Description:

	Simplifies the constant matrix. Since the constant matrix is always in simplest form,
	this function simply returns the original constant matrix.
*/
func (km KMatrix) AsSimplifiedExpression() Expression {
	return km
}
