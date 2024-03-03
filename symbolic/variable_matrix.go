package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
variable_matrix.go
Description:

	This file contains functions that are used to create and
	manipulate matrices of variables (i.e., VariableMatrix) objects.
*/

// Types
// =====

type VariableMatrix [][]Variable

// =======
// Methods
// =======

/*
Check
Description:

	This method is used to make sure that the variable matrix is well-defined.
	It checks:
	- That the matrix is not empty.
	- That all of the rows have the same number of columns.
	- That all of the variables are well-defined.
*/
func (vm VariableMatrix) Check() error {
	// Input Processing

	// Check to see if the matrix is empty
	if len(vm) == 0 {
		return smErrors.EmptyMatrixError{
			Expression: vm,
		}
	}

	// Check the number of columns in each row
	var numCols int = len(vm[0])
	for ii, vmRow := range vm {
		// Check the length of row i
		if len(vmRow) != numCols {
			return smErrors.MatrixColumnMismatchError{
				ExpectedNColumns: numCols,
				ActualNColumns:   len(vmRow),
				Row:              ii,
			}
		}
	}

	// Check the variables
	for ii, vmRow := range vm {
		for jj, v := range vmRow {
			err := v.Check()
			if err != nil {
				return fmt.Errorf(
					"error in entry (%v, %v): %v",
					ii, jj,
					err,
				)
			}
		}
	}

	// All checks passed
	return nil
}

/*
Variables
Description:

	This function returns the list of variables in the matrix.
*/
func (vm VariableMatrix) Variables() []Variable {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var variables []Variable
	for _, vmRow := range vm {
		for _, v := range vmRow {
			variables = append(variables, v)
		}
	}

	return UniqueVars(variables)
}

/*
Dims
Description:

	This function returns the dimensions of the variable matrix.
*/
func (vm VariableMatrix) Dims() []int {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return []int{len(vm), len(vm[0])}
}

/*
Plus
Description:

	This function adds two variable matrices together.
*/
func (vm VariableMatrix) Plus(e interface{}) Expression {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		// Convert e to an expression
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(
				fmt.Errorf("error in second argument to VariableMatrix.Plus: %v", err),
			)
		}

		err := smErrors.CheckDimensionsInAddition(vm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := e.(type) {
	case float64:
		// Use K case
		return vm.Plus(K(right))
	case K:
		// Create K Matrix from this
		var kmOut mat.Dense
		onesMat := OnesMatrix(vm.Dims()[0], vm.Dims()[1])
		kmOut.Scale(float64(right), &onesMat)

		return vm.Plus(DenseToKMatrix(kmOut))

	case KMatrix:
		// Create a new matrix of polynomials.
		var pmOut PolynomialMatrix
		for ii, vmRow := range vm {
			var pmRow []Polynomial
			for jj, v := range vmRow {
				pmRow = append(pmRow, v.Plus(right[ii][jj]).(Polynomial))
			}
			pmOut = append(pmOut, pmRow)
		}
		return pmOut
	}

	// panic if the type is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableMatrix.Plus",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	This function subtracts a variable matrix from another term.
*/
func (vm VariableMatrix) Minus(e interface{}) Expression {
	// Input Processing
	// - Check the variable matrix is well-defined
	// - If e is an expression, then:
	//   + Check that it is a well-defined expression
	//   + Check that the dimensions of the two expressions match

	err := vm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		// Convert e to an expression
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(
				fmt.Errorf("error in second argument to VariableMatrix.Minus: %v", err),
			)
		}

		err := smErrors.CheckDimensionsInSubtraction(vm, eAsE)
		if err != nil {
			panic(err)
		}

		// Use Expression's Minus Function
		return Minus(vm, eAsE)
	}

	// Algorithm for non-expressions
	switch right := e.(type) {
	case float64:
		return vm.Minus(K(right)) // Use K case
	case mat.Dense:
		return vm.Minus(DenseToKMatrix(right)) // Use KMatrix case
	case *mat.Dense:
		return vm.Minus(DenseToKMatrix(*right)) // Use KMatrix case
	}

	// panic if the type is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableMatrix.Minus",
			Input:        e,
		},
	)

}

/*
Multiply
Description:

	This function multiplies a variable matrix by another term.
*/
func (vm VariableMatrix) Multiply(e interface{}) Expression {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		// Convert e to an expression
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(
				fmt.Errorf("error in second argument to VariableMatrix.Multiply: %v", err),
			)
		}

		err := smErrors.CheckDimensionsInMultiplication(vm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Computation
	switch right := e.(type) {
	case float64:
		// Use K case
		return vm.Multiply(K(right))
	case K:
		// Create a new matrix of polynomials.
		var mmOut MonomialMatrix
		for _, vmRow := range vm {
			var mmRow []Monomial
			for _, v := range vmRow {
				mmRow = append(mmRow, v.Multiply(right).(Monomial))
			}
			mmOut = append(mmOut, mmRow)
		}
		return mmOut

	case KVector:
		// Constants
		nResultRows := vm.Dims()[0]

		// Switch on the dimensions of the result
		switch nResultRows {
		case 1:
			// Scalar result
			var result Polynomial = K(0).ToMonomial().ToPolynomial()
			for ii, k := range right {
				result = result.Plus(vm[0][ii].Multiply(k)).(Polynomial)
			}
			return result
		default:
			// Create vector result
			var result PolynomialVector = VecDenseToKVector(ZerosVector(nResultRows)).ToPolynomialVector()
			for ii, vmRow := range vm {
				for jj, v := range vmRow {
					result[ii] = result[ii].Plus(v.Multiply(right[jj])).(Polynomial)
				}
			}
			return result
		}
	case *mat.Dense:
		// Use the mat.Dense case
		return vm.Multiply(DenseToKMatrix(*right))
	case mat.Dense:
		// Use the KMatrix case
		return vm.Multiply(DenseToKMatrix(right))
	case KMatrix:
		// Collect dimensions
		nResultRows, nResultCols := vm.Dims()[0], right.Dims()[1]

		// Switch on the dimensions of the result
		switch {
		case (nResultRows == 1) && (nResultCols == 1):
			// Scalar result
			var result Polynomial = K(0).ToMonomial().ToPolynomial()

			for ii, vmRow := range vm {
				for jj, vIJ := range vmRow {
					result = result.Plus(vIJ.Multiply(right[jj][ii])).(Polynomial)
				}
			}
			return result
		case nResultCols == 1:
			// Vector result
			var result PolynomialVector = VecDenseToKVector(ZerosVector(nResultRows)).ToPolynomialVector()

			for ii, vmRow := range vm {
				for jj, vIJ := range vmRow {
					result[ii] = result[ii].Plus(vIJ.Multiply(right[jj][0])).(Polynomial)
				}
			}

			return result

		default:
			// Create result
			var result PolynomialMatrix

			for ii := 0; ii < nResultRows; ii++ {
				var resultRow []Polynomial
				for jj := 0; jj < nResultCols; jj++ {
					resultRow = append(resultRow, K(0).ToMonomial().ToPolynomial())
				}
				result = append(result, resultRow)
			}

			// Fill in the elements of the new matrix
			for ii := 0; ii < nResultRows; ii++ {
				for jj := 0; jj < nResultCols; jj++ {
					// Compute Sum
					for kk := 0; kk < vm.Dims()[1]; kk++ {
						result[ii][jj] = result[ii][jj].Plus(
							vm[ii][kk].Multiply(right[kk][jj]),
						).(Polynomial)
					}
				}
			}
			return result
		}
	}

	// panic if the type is not recognized
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableMatrix.Multiply",
			Input:        e,
		},
	)
}

/*
Transpose
Description:

	This function returns the transpose of the variable matrix.
*/
func (vm VariableMatrix) Transpose() Expression {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Create blank variable matrix
	nRows, nCols := vm.Dims()[0], vm.Dims()[1]
	var vmOut VariableMatrix
	for ii := 0; ii < nCols; ii++ {
		vmOut = append(vmOut, make([]Variable, nRows))
	}

	// Fill in the elements of the new matrix
	for ii := 0; ii < nRows; ii++ {
		for jj := 0; jj < nCols; jj++ {
			vmOut[jj][ii] = vm[ii][jj]
		}
	}

	return vmOut
}

/*
Comparison
Description:

	This function compares the variable matrix to another expression.
*/
func (vm VariableMatrix) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		// Convert e to an expression
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(
				fmt.Errorf("error in second argument to VariableMatrix.Comparison: %v", err),
			)
		}

		err := CheckDimensionsInComparison(vm, rightAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	// Algorithm
	switch right := rightIn.(type) {
	case float64:
		// Use K case
		return vm.Comparison(K(right), sense)
	case K:
		// Create a new matrix of polynomials.
		onesMat := OnesMatrix(vm.Dims()[0], vm.Dims()[1])
		var KAsDense mat.Dense
		KAsDense.Scale(float64(right), &onesMat)

		return vm.Comparison(KAsDense, sense)
	case mat.Dense:
		// Use the KMatrix case
		return vm.Comparison(DenseToKMatrix(right), sense)
	case KMatrix:
		return MatrixConstraint{
			LeftHandSide:  vm,
			RightHandSide: right,
			Sense:         sense,
		}
	case MonomialMatrix:
		return MatrixConstraint{
			LeftHandSide:  vm,
			RightHandSide: right,
			Sense:         sense,
		}

	}

	// If the type is not recognized, panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "VariableMatrix.Comparison",
			Input:        rightIn,
		},
	)
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between
	the VariableMatrix and another expression.
*/
func (vm VariableMatrix) LessEq(rightIn interface{}) Constraint {
	return vm.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Returns a greater than or equal to (>=) constraint between
	the VariableMatrix and another expression.
*/
func (vm VariableMatrix) GreaterEq(rightIn interface{}) Constraint {
	return vm.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Returns an equality (==) constraint between
	the VariableMatrix and another expression.
*/
func (vm VariableMatrix) Eq(rightIn interface{}) Constraint {
	return vm.Comparison(rightIn, SenseEqual)
}

/*
DerivativeWrt
Description:

	This function returns the derivative of the variable matrix
	with respect to a given variable.
*/
func (vm VariableMatrix) DerivativeWrt(v Variable) Expression {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	kmOut := DenseToKMatrix(
		ZerosMatrix(vm.Dims()[0], vm.Dims()[1]),
	)

	// Iterate through matrix and if there is one
	// element that is equal to the variable, set the
	// corresponding element in the output matrix to 1.
	for ii, vmRow := range vm {
		for jj, vmElt := range vmRow {
			if vmElt.ID == v.ID {
				kmOut[ii][jj] = K(1)
			}
		}
	}

	return kmOut
}

/*
At
Description:

	This function returns the element of the variable matrix at the given indices.
*/
func (vm VariableMatrix) At(ii, jj int) ScalarExpression {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	return vm[ii][jj]
}

/*
Constant
Description:

	This function returns the constant value in the variable matrix.
	this should always be zero.
*/
func (vm VariableMatrix) Constant() mat.Dense {
	return ZerosMatrix(vm.Dims()[0], vm.Dims()[1])
}

/*
String
Description:

	This function returns a string representation of the variable matrix.
*/
func (vm VariableMatrix) String() string {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var out string = "["
	for ii, vmRow := range vm {
		out += "["
		for jj, v := range vmRow {
			out += v.String()
			if jj < len(vmRow)-1 {
				out += ", "
			}
		}
		out += "]"
		if ii < len(vm)-1 {
			out += ", "
		}
	}
	out += "]"
	return out
}

/*
NewVariableMatrix
Description:

	This function creates a new variable matrix
	and properly initializes each element in it.
*/
func NewVariableMatrix(nRows, nCols int, envs ...*Environment) VariableMatrix {
	return NewVariableMatrixClassic(nRows, nCols, envs...)
}

/*
NewVariableMatrixClassic
Description:

	This function creates a new variable matrix
	and properly initializes each element in it
	when given a list of variables.
*/
func NewVariableMatrixClassic(nRows, nCols int, envs ...*Environment) VariableMatrix {
	// Collect an environment if one exists
	var env *Environment
	switch len(envs) {
	case 0:
		env = &BackgroundEnvironment
	case 1:
		env = envs[0]
	default:
		panic(
			fmt.Errorf("Too many inputs provided to NewVariableMatrix() method."),
		)
	}

	// Create a new matrix
	var vmOut VariableMatrix
	for ii := 0; ii < nRows; ii++ {
		vmOut = append(vmOut, make([]Variable, nCols))
		for jj := 0; jj < nCols; jj++ {
			vmOut[ii][jj] = NewContinuousVariable(env)
		}
	}
	return vmOut
}

/*
ToMonomialMatrix
Description:

	This function converts the variable matrix to a monomial matrix.
*/
func (vm VariableMatrix) ToMonomialMatrix() MonomialMatrix {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var mmOut MonomialMatrix
	for _, vmRow := range vm {
		var mmRow []Monomial
		for _, v := range vmRow {
			mmRow = append(mmRow, v.ToMonomial())
		}
		mmOut = append(mmOut, mmRow)
	}
	return mmOut
}

/*
ToPolynomialMatrix
Description:

	This function converts the variable matrix to a polynomial matrix.
*/
func (vm VariableMatrix) ToPolynomialMatrix() PolynomialMatrix {
	// Input Processing
	err := vm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var pmOut PolynomialMatrix
	for _, vmRow := range vm {
		var pmRow []Polynomial
		for _, v := range vmRow {
			pmRow = append(pmRow, v.ToPolynomial())
		}
		pmOut = append(pmOut, pmRow)
	}
	return pmOut
}
