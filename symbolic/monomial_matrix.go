package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
monomial_matrix.go
Description:

	Defines a matrix of monomials and its associated methods.
*/

// ===============
// Type Definition
// ===============

type MonomialMatrix [][]Monomial

// =========
// Functions
// =========

/*
Check
Description:

	Verifies that:
	- The matrix has at least one row
	- The number of columns is the same in each row.
	- Each of the monomials in the matrix are valid.
*/
func (mm MonomialMatrix) Check() error {
	// Check that the matrix has at least one row
	if len(mm) == 0 {
		return smErrors.EmptyMatrixError{mm}
	}

	// Check that the number of columns is the same in each row
	numColumns := len(mm[0])
	for ii, row := range mm {
		if len(row) != numColumns {
			return smErrors.MatrixColumnMismatchError{
				ExpectedNColumns: numColumns,
				ActualNColumns:   len(row),
				Row:              ii,
			}

		}
	}

	// Check that each of the monomials are well formed
	for ii, row := range mm {
		for jj, monomial := range row {
			err := monomial.Check()
			if err != nil {
				return fmt.Errorf("error in monomial %v,%v: %v", ii, jj, err)
			}
		}
	}

	// All checks passed
	return nil
}

/*
Variables
Description:

	Returns the variables in the matrix.
*/
func (mm MonomialMatrix) Variables() []Variable {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	// Get the variables
	variables := []Variable{}
	for _, row := range mm {
		for _, monomial := range row {
			variables = append(variables, monomial.Variables()...)
		}
	}

	// Return the variables
	return UniqueVars(variables)
}

/*
Dims
Description:

	Returns the dimensions of the matrix.
*/
func (mm MonomialMatrix) Dims() []int {
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	return []int{len(mm), len(mm[0])}
}

/*
Plus
Description:

	Addition of the monomial matrix with another expression.
*/
func (mm MonomialMatrix) Plus(e interface{}) Expression {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Plus: %v", err))
		}

		err := CheckDimensionsInAddition(mm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := e.(type) {
	case float64:
		return mm.Plus(K(right))
	case K:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range mm {
			sumRow := make([]Polynomial, len(row))
			for jj, monomial := range row {
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, monomial, right.ToMonomial())
				// Simplify
				sumRow[jj].Simplify()
			}
			sum = append(sum, sumRow)
		}
		return sum
	case Variable:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range mm {
			sumRow := make([]Polynomial, len(row))
			for jj, monomial := range row {
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, monomial, right.ToMonomial())
				// Simplify
				sumRow[jj].Simplify()
			}
			sum = append(sum, sumRow)
		}
		return sum

	case Monomial:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range mm {
			sumRow := make([]Polynomial, len(row))
			for jj, monomial := range row {
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, monomial, right)
				// Simplify
				sumRow[jj].Simplify()
			}
			sum = append(sum, sumRow)
		}

		return sum

	case Polynomial:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range mm {
			sumRow := make([]Polynomial, len(row))
			for jj, monomial := range row {
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, monomial)
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, right.Monomials...)
				// Simplify
				sumRow[jj].Simplify()
			}
			sum = append(sum, sumRow)
		}
		return sum
	case PolynomialMatrix:
		// Create containers
		var sum PolynomialMatrix

		for ii, row := range mm {
			sumRow := make([]Polynomial, len(row))
			for jj, monomial := range row {
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, monomial)
				sumRow[jj].Monomials = append(sumRow[jj].Monomials, right[ii][jj].Monomials...)
				// Simplify
				sumRow[jj].Simplify()
			}
			sum = append(sum, sumRow)
		}

		return sum

	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Plus",
				Input:        right,
			},
		)
	}
}

/*
Multiply
Description:

	Multiplication of the monomial matrix with another expression.
*/
func (mm MonomialMatrix) Multiply(e interface{}) Expression {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	nRows := mm.Dims()[0]
	//nCols := -1 // Will be reassigned for a valid expression.

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Multiply: %v", err))
		}

		err := CheckDimensionsInMultiplication(mm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants

	// Algorithm
	switch right := e.(type) {
	case float64:
		return mm.Multiply(K(right))
	case K:
		// Create containers
		var product MonomialMatrix

		for _, row := range mm {
			productRow := make([]Monomial, len(row))
			for jj, monomial := range row {
				productRow[jj] = monomial.Multiply(right).(Monomial)
			}
			product = append(product, productRow)
		}
		return product
	case VariableVector:
		if nRows == 1 {
			// Output will be a polynomial
			var product Polynomial
			for ii, monomial := range mm[0] {
				product.Monomials = append(product.Monomials, monomial.Multiply(right[ii]).(Monomial))
			}
			return product.Simplify()

		} else {
			// Output will be a polynomial matrix
			var product PolynomialVector
			for _, row := range mm {
				product_ii := row[0].ToPolynomial().Multiply(right[0]).(Polynomial)
				for jj := 1; jj < len(row); jj++ {
					product_ii = product_ii.Plus(
						row[jj].ToPolynomial().Multiply(right[jj]),
					).(Polynomial)
				}
				product = append(product, product_ii)
			}

			return product

		}
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Multiply",
				Input:        right,
			},
		)
	}
}

/*
Transpose
Description:

	Returns the transpose of the monomial matrix.
*/
func (mm MonomialMatrix) Transpose() Expression {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	dims := mm.Dims()
	n_row, n_cols := dims[0], dims[1]

	// Create blank matrix monomial
	var mmOut MonomialMatrix
	for ii := 0; ii < n_cols; ii++ {
		mmOut = append(mmOut, make([]Monomial, n_row))
	}

	// Algorithm
	for ii := 0; ii < len(mm[0]); ii++ {
		for jj := 0; jj < len(mm); jj++ {
			mmOut[ii][jj] = mm[jj][ii]
		}
	}

	return mmOut
}

/*
Comparison
Description:

	Compares the monomial matrix to another expression according to the sense `sense`.
*/
func (mm MonomialMatrix) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = rightAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Comparison: %v", err))
		}

		err := CheckDimensionsInComparison(mm, rightAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := rightIn.(type) {
	case float64:
		return mm.Comparison(K(right), sense)
	case K:
		// Create containers
		onesMat := OnesMatrix(mm.Dims()[0], mm.Dims()[1])
		var KAsDense mat.Dense
		KAsDense.Scale(float64(right), &onesMat)

		return MatrixConstraint{
			LeftHandSide:  mm,
			RightHandSide: DenseToKMatrix(KAsDense),
			Sense:         sense,
		}
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "MonomialVector.Comparison",
				Input:        right,
			},
		)
	}
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the
	current expression and another.
*/
func (mm MonomialMatrix) LessEq(rightIn interface{}) Constraint {
	return mm.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Returns a greater than or equal to (>=) constraint between the
	current expression and another.
*/
func (mm MonomialMatrix) GreaterEq(rightIn interface{}) Constraint {
	return mm.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Returns an equality (==) constraint between the current expression
	and another.
*/
func (mm MonomialMatrix) Eq(rightIn interface{}) Constraint {
	return mm.Comparison(rightIn, SenseEqual)
}

/*
DerivativeWrt
Description:

	Returns the derivative of the monomial matrix with respect to the input variable.
*/
func (mm MonomialMatrix) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Compute the Derivative of each monomial
	var dmm MonomialMatrix
	for _, row := range mm {
		var dmmRow []Monomial
		for _, monomial := range row {
			dMonomial := monomial.DerivativeWrt(vIn)
			var dMonomialAsMonomial Monomial
			switch dMonomial.(type) {
			case Monomial:
				dMonomialAsMonomial = dMonomial.(Monomial)
			case K:
				dMonomialAsMonomial = dMonomial.(K).ToMonomial()
			default:
				panic(
					fmt.Errorf("unexpected type of derivative: %T (%v)", dMonomial, dMonomial),
				)
			}
			// Add the converted dMonomial to dmmRow
			dmmRow = append(dmmRow, dMonomialAsMonomial)
		}
		dmm = append(dmm, dmmRow)
	}

	// Return the derivative
	return dmm
}

/*
At
Description:

	Returns the (ii,jj)-th value of the monomial matrix.
*/
func (mm MonomialMatrix) At(ii, jj int) ScalarExpression {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	// Return the value
	return mm[ii][jj]
}

/*
Constant
Description:

	Returns the components of the monomial matrix which are constant-valued.
*/
func (mm MonomialMatrix) Constant() mat.Dense {
	// Input Processing
	err := mm.Check()
	if err != nil {
		panic(err)
	}

	// Create container for constant
	dims := mm.Dims()
	nRows, nCols := dims[0], dims[1]
	constant := ZerosMatrix(nRows, nCols)
	for ii := 0; ii < nRows; ii++ {
		for jj := 0; jj < nCols; jj++ {
			if mm[ii][jj].IsConstant() {
				constant.Set(ii, jj, mm[ii][jj].Coefficient)
			}
		}
	}

	// Return the constant
	return constant
}
