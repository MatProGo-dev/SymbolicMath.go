package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
polynomial_matrix.go
Description:

	Defines a matrix of polynomials and its associated methods.
*/

// ===============
// Type Definition
// ===============

type PolynomialMatrix [][]Polynomial

// =========
// Functions
// =========

/*
Check
Description:

	Verifies that:
	- The matrix has at least one row
	- The number of columns is the same in each row.
	- Each of the polynomials in the matrix are valid.
*/
func (pm PolynomialMatrix) Check() error {
	// Check that the matrix has at least one row
	if len(pm) == 0 {
		return smErrors.EmptyMatrixError{Expression: pm}
	}

	// Check that the number of columns is the same in each row
	numColumns := len(pm[0])
	for ii, row := range pm {
		if len(row) != numColumns {
			return smErrors.MatrixColumnMismatchError{
				ExpectedNColumns: numColumns,
				ActualNColumns:   len(row),
				Row:              ii,
			}

		}
	}

	// Check that each of the polynomials are well formed
	for ii, row := range pm {
		for jj, polynomial := range row {
			err := polynomial.Check()
			if err != nil {
				return fmt.Errorf("error in polynomial %v,%v: %v", ii, jj, err)
			}
		}
	}

	// All checks passed
	return nil
}

/*
Variables
Description:

	Returns the variables in the polynomial matrix.
*/
func (pm PolynomialMatrix) Variables() []Variable {
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	variables := []Variable{}
	for _, row := range pm {
		for _, polynomial := range row {
			variables = append(variables, polynomial.Variables()...)
		}
	}

	return UniqueVars(variables)
}

/*
Dims
Description:

	Returns the dimensions of the matrix of polynomials.
*/
func (pm PolynomialMatrix) Dims() []int {
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	return []int{len(pm), len(pm[0])}
}

/*
Plus
Description:

	Addition of the polynomial matrix with another expression.
*/
func (pm PolynomialMatrix) Plus(e interface{}) Expression {
	// Input Processing
	// - Check that pm is valid
	// - Check that the input expression (if it is an expression)
	//   + is valid
	//	 + has the same dimensions as pm

	err := pm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Plus: %v", err))
		}

		err = smErrors.CheckDimensionsInAddition(pm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Perform the addition
	switch right := e.(type) {
	case float64:
		return pm.Plus(K(right))
	case K:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range pm {
			var sumRow []Polynomial
			for _, polynomial := range row {
				sumRow = append(sumRow, polynomial.Plus(right).(Polynomial))
			}
			sum = append(sum, sumRow)
		}
		return sum
	case Monomial:
		return pm.Plus(right.ToPolynomial())
	case Polynomial:
		// Create containers
		var sum PolynomialMatrix

		for _, row := range pm {
			var sumRow []Polynomial
			for _, polynomial := range row {
				sumRow = append(sumRow, polynomial.Plus(right).(Polynomial))
			}
			sum = append(sum, sumRow)
		}
		return sum
	case KMatrix:
		// Create containers
		var sum PolynomialMatrix

		for ii, row := range pm {
			var sumRow []Polynomial
			for jj, polynomial := range row {
				sumRow = append(sumRow, polynomial.Plus(right.At(ii, jj).(K)).(Polynomial))
			}
			sum = append(sum, sumRow)
		}

		return sum

	case PolynomialMatrix:
		// Create containers
		var sum PolynomialMatrix

		for ii, row := range pm {
			var sumRow []Polynomial
			for jj, polynomial := range row {
				sumRow = append(sumRow, polynomial.Plus(right[ii][jj]).(Polynomial))
			}
			sum = append(sum, sumRow)
		}

		return sum.Simplify()
	}

	// If type isn't recognized, then panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "PolynomialMatrix.Plus",
			Input:        e,
		},
	)
}

/*
Minus
Description:

	Subtracts another expression from the given polynomial matrix
	and returns the result.
*/
func (pm PolynomialMatrix) Minus(e interface{}) Expression {
	// Input Processing
	// - Check that pm is valid
	// - Check that the input expression (if it is an expression)
	//   + is valid
	//	 + has the same dimensions as pm

	err := pm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Minus: %v", err))
		}

		err = smErrors.CheckDimensionsInSubtraction(pm, eAsE)
		if err != nil {
			panic(err)
		}

		// Perform the subtraction with Expression's Minus() function
		return Minus(pm, eAsE)
	}

	// If type isn't recognized, then panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "PolynomialMatrix.Minus",
			Input:        e,
		},
	)
}

/*
Multiply
Description:

	Multiplication of the polynomial matrix with another expression.
*/
func (pm PolynomialMatrix) Multiply(e interface{}) Expression {
	// Input Processing
	// - Check that pm is valid
	// - Check that the input expression (if it is an expression)
	//   + is valid
	//	 + has the matching dimensions for pm

	err := pm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Multiply: %v", err))
		}

		err = smErrors.CheckDimensionsInMultiplication(pm, eAsE)
		if err != nil {
			panic(err)
		}
	}

	// Perform the multiplication
	switch right := e.(type) {
	case float64:
		return pm.Multiply(K(right))
	case K:
		// Create containers
		var product PolynomialMatrix

		for _, row := range pm {
			var productRow []Polynomial
			for _, polynomial := range row {
				polynomialCopy := Polynomial{Monomials: make([]Monomial, len(polynomial.Monomials))}
				copy(polynomialCopy.Monomials, polynomial.Monomials)
				productRow = append(productRow, polynomialCopy.Multiply(right).(Polynomial))
			}
			product = append(product, productRow)
		}
		return product
	case VariableVector:
		// Identify output dimensions
		nResultRows := pm.Dims()[0]

		// Create product based on the number of Resulting rows
		if nResultRows == 1 {
			// Create container
			var product Polynomial = K(0).ToPolynomial()
			for ii, tempPolynomial := range pm[0] {
				product = product.Plus(
					tempPolynomial.Multiply(right[ii]).(Polynomial),
				).(Polynomial)
			}
			return product
		} else {
			// Create container
			var product PolynomialVector = VecDenseToKVector(
				ZerosVector(nResultRows),
			).ToPolynomialVector()

			// Fill container
			for ii := 0; ii < nResultRows; ii++ {
				// Construct the ii-th element of the product
				for jj, polynomial := range pm[ii] {
					product[ii] = product[ii].Plus(
						polynomial.Multiply(right[jj]).(Polynomial),
					).(Polynomial)
				}
			}
			return product
		}
	}

	// If type isn't recognized, then panic
	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "PolynomialMatrix.Multiply",
			Input:        e,
		},
	)
}

/*
Transpose
Description:

	Transposes the polynomial matrix.
*/
func (pm PolynomialMatrix) Transpose() Expression {
	// Input Processing
	// - Check that pm is valid
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	nRows, nCols := pm.Dims()[0], pm.Dims()[1]

	// Perform the transpose
	var pmT PolynomialMatrix
	for rowIndex := 0; rowIndex < nCols; rowIndex++ {
		tempRow := make([]Polynomial, nRows)
		for colIndex := 0; colIndex < nRows; colIndex++ {
			tempRow[colIndex] = pm[colIndex][rowIndex]
		}
		// Append row to container
		pmT = append(pmT, tempRow)
	}

	// Return the transpose
	return pmT
}

/*
Comparison
Description:

	Compares the polynomial matrix to another expression.
*/
func (pm PolynomialMatrix) Comparison(e interface{}, sense ConstrSense) Constraint {
	// Input Checking
	// - Check that pm is valid
	// - Check that the input expression (if it is an expression)
	//   + is valid
	//	 + has the same dimensions for pm

	err := pm.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(e) {
		eAsE, _ := ToExpression(e)
		err = eAsE.Check()
		if err != nil {
			panic(fmt.Errorf("error in second argument to Comparison: %v", err))
		}

		err = CheckDimensionsInComparison(pm, eAsE, sense)
		if err != nil {
			panic(err)
		}
	}

	// Perform Comparison
	switch right := e.(type) {
	case float64:
		return pm.Comparison(K(right), sense)
	case K:
		// Create containers
		onesMat := OnesMatrix(pm.Dims()[0], pm.Dims()[1])
		var KAsDense mat.Dense
		KAsDense.Scale(float64(right), &onesMat)

		return pm.Comparison(KAsDense, sense)
	case mat.Dense:
		return pm.Comparison(DenseToKMatrix(right), sense)
	case *mat.Dense:
		return pm.Comparison(*right, sense)
	case KMatrix:
		return MatrixConstraint{
			LeftHandSide:  pm,
			RightHandSide: right,
			Sense:         sense,
		}
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "PolynomialMatrix.Comparison",
				Input:        right,
			},
		)
	}
}

/*
LessEq
Description:

	Compares the polynomial matrix to another expression using
	the SenseLessThanEqual sense.
*/
func (pm PolynomialMatrix) LessEq(e interface{}) Constraint {
	return pm.Comparison(e, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	Compares the polynomial matrix to another expression using
	the SenseGreaterThanEqual sense.
*/
func (pm PolynomialMatrix) GreaterEq(e interface{}) Constraint {
	return pm.Comparison(e, SenseGreaterThanEqual)
}

/*
Eq
Description:

	Compares the polynomial matrix to another expression using
	the SenseEqual sense.
*/
func (pm PolynomialMatrix) Eq(e interface{}) Constraint {
	return pm.Comparison(e, SenseEqual)
}

/*
DerivativeWrt
Description:

	Returns the derivative of the polynomial matrix with respect to the
	input variable.
*/
func (pm PolynomialMatrix) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	// - Check that pm is valid
	// - Check that vIn is valid
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Create containers
	var dpm PolynomialMatrix

	for _, row := range pm {
		var dpmRow []Polynomial
		for _, polynomial := range row {
			dpmRow = append(dpmRow, polynomial.DerivativeWrt(vIn).(Polynomial))
		}
		dpm = append(dpm, dpmRow)
	}

	return dpm
}

/*
At
Description:

	Returns the (ii, jj)-th element of the polynomial matrix.
*/
func (pm PolynomialMatrix) At(ii int, jj int) ScalarExpression {
	// Input Processing
	// - Check that pm is valid
	// - Check that ii, jj is in bounds

	err := pm.Check()
	if err != nil {
		panic(err)
	}

	err = smErrors.CheckIndexOnMatrix(ii, jj, pm)
	if err != nil {
		panic(err)
	}

	// Return the element
	return pm[ii][jj]
}

/*
Constant
Description:

	Returns the components of the polynomial matrix which are
	constant-valued.
*/
func (pm PolynomialMatrix) Constant() mat.Dense {
	// Input Processing
	// - Check that pm is valid
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	// Create containers
	nRows, nCols := pm.Dims()[0], pm.Dims()[1]
	constant := ZerosMatrix(nRows, nCols)
	for ii := 0; ii < nRows; ii++ {
		for jj := 0; jj < nCols; jj++ {
			constant.Set(ii, jj, pm.At(ii, jj).Constant())
		}
	}

	return constant
}

/*
Simplify
Description:

	Simplifies the polynomial matrix, if possible.
*/
func (pm PolynomialMatrix) Simplify() PolynomialMatrix {
	// Constants
	nRows, nCols := pm.Dims()[0], pm.Dims()[1]

	// Fill container with simplified polynomials
	var simplified PolynomialMatrix
	for rowIndex := 0; rowIndex < nRows; rowIndex++ {
		tempRow := make([]Polynomial, nCols)
		for colIndex := 0; colIndex < nCols; colIndex++ {
			tempRow[colIndex] = pm[rowIndex][colIndex].Simplify()
		}
		simplified = append(simplified, tempRow)
	}

	// Return simplified polynomial
	return simplified
}

/*
String
Description:

	Returns a string representation of the polynomial matrix.
*/
func (pm PolynomialMatrix) String() string {
	// Input Processing
	// - Check that pm is valid
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	// Constants
	nRows, nCols := pm.Dims()[0], pm.Dims()[1]

	// Create the string
	var out string = "PolynomialMatrix =\n["
	for ii, row := range pm {
		out += "["
		for jj, polynomial := range row {
			out += fmt.Sprintf("%v", polynomial)
			if jj != nCols-1 {
				out += ", "
			}
		}
		out += "]"
		if ii != nRows-1 {
			out += ",\n"
		}
	}
	out += "]"

	// Return the string
	return out
}

/*
Degree
Description:

	The degree of the polynomial matrix
	is the maximum degree of the entries.
*/
func (pm PolynomialMatrix) Degree() int {
	// Input Processing
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	maxDegree := 0
	for _, row := range pm {
		for _, p := range row {
			if p.Degree() > maxDegree {
				maxDegree = p.Degree()
			}
		}
	}
	return maxDegree
}

/*
Substitute
Description:

	Substitutes the variable vIn with the expression eIn in the polynomial matrix.
*/
func (pm PolynomialMatrix) Substitute(vIn Variable, eIn ScalarExpression) Expression {
	return MatrixSubstituteTemplate(pm, vIn, eIn)
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the polynomial matrix with the corresponding expressions in the map.
*/
func (pm PolynomialMatrix) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	// Input Processing
	err := pm.Check()
	if err != nil {
		panic(err)
	}

	err = CheckSubstitutionMap(subMap)
	if err != nil {
		panic(err)
	}

	// Algorithm
	var out MatrixExpression = pm
	for v, e := range subMap {
		outSubbed := out.Substitute(v, e.(ScalarExpression))
		out = outSubbed.(MatrixExpression)
	}

	return out
}

/*
Power
Description:

	Raises the polynomial matrix to the power of the input integer.
*/
func (pm PolynomialMatrix) Power(exponent int) Expression {
	return MatrixPowerTemplate(pm, exponent)
}
