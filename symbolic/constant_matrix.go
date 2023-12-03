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
		// Check dimensions
		rightAsE, _ := ToExpression(e)
		err = CheckDimensionsInAddition(km, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	switch right := e.(type) {
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
		// Check dimensions
		rightAsE, _ := ToExpression(e)
		err = CheckDimensionsInMultiplication(km, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	switch right := e.(type) {
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
func (km KMatrix) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return MatrixConstraint{}, err
	}

	return MatrixConstraint{}, fmt.Errorf("not implemented")

}

/*
GreaterEq
Description:

	Returns a greater equal constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return MatrixConstraint{}, err
	}

	return MatrixConstraint{}, fmt.Errorf("not implemented")

}

/*
Eq
Description:

	Returns an equal constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return MatrixConstraint{}, err
	}

	return MatrixConstraint{}, fmt.Errorf("not implemented")

}

/*
Comparison
Description:

	Returns a constraint between the KMatrix and the
	expression on the right hand side.
*/
func (km KMatrix) Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return MatrixConstraint{}, err
	}

	return MatrixConstraint{}, fmt.Errorf("not implemented")

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
