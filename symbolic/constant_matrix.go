package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
constant_matrix.go
Description:
	Defines all methods related to the constant matrix type.
*/

type KMatrix mat.Dense

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
	kmAsDense := mat.Dense(km)

	nR, nC := kmAsDense.Dims()

	return []int{nR, nC}
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
		err = CheckDimensionsInAddition(km, rightAsE)
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
		kmAsDense := mat.Dense(km)

		// Compute sum
		var sumAsDense mat.Dense
		sumAsDense.Add(&rightAsDense, &kmAsDense)

		return KMatrix(sumAsDense)

	case K:
		return km.Plus(float64(right)) // Reuse float64 case

	case PolynomialMatrix:
		return right.Plus(km) // Reuse PolynomialMatrix case

	default:
		panic(
			fmt.Errorf(
				"The input to KMatrix's Plus() method (%v) has unexpected type: %T",
				right, right,
			),
		)
	}
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
		err = CheckDimensionsInMultiplication(km, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	switch right := e.(type) {
	case float64:
		// Use gonum's built-in scale function
		kmAsDense := mat.Dense(km)
		var product mat.Dense
		product.Scale(right, &kmAsDense)

		return KMatrix(product)

	case K:
		return km.Multiply(float64(right)) // Reuse float64 case

	default:
		panic(
			fmt.Errorf(
				"The input to KMatrix's Multiply method (%v) has unexpected type: %T",
				right, right,
			),
		)
	}
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
	kmAsDense := mat.Dense(km)
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
	return KMatrix(kmT)

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

	// Algorithm
	return MatrixConstraint{}

}

/*
At
Description:

	Retrieves element at the specified indices.
*/
func (km KMatrix) At(i, j int) ScalarExpression {
	kmAsD := mat.Dense(km)
	return K(kmAsD.At(i, j))
}

/*
Constant
Description:

	Retrieves the constant component.
*/
func (km KMatrix) Constant() mat.Dense {
	return mat.Dense(km)
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
	return KMatrix(ZerosMatrix(km.Dims()[0], km.Dims()[1]))
}
