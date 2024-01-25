package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
monomial_matrix_test.go
Description:

	Tests for the functions mentioned in the monomial_matrix.go file.
*/

/*
TestMonomialMatrix_Check1
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Monomials (i.e., no monomials are given).
*/
func TestMonomialMatrix_Check1(t *testing.T) {
	// Constants
	mm := symbolic.MonomialMatrix{}
	expectedError := smErrors.EmptyMatrixError{mm}

	// Test
	err := mm.Check()
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}
}

/*
TestMonomialMatrix_Check2
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Monomials (i.e., a monomial is given with an improper number of
	columns in row 2 than in every other row.
*/
func TestMonomialMatrix_Check2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1},
		{m1, m1},
	}
	expectedError := smErrors.MatrixColumnMismatchError{
		ExpectedNColumns: 2,
		ActualNColumns:   1,
		Row:              1,
	}

	// Test
	err := mm.Check()
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}

}

/*
TestMonomialMatrix_Check3
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Monomials (i.e., a monomial is given with an improper number of
	degrees).
*/
func TestMonomialMatrix_Check3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1},
	}
	expectedErrorComponent := m1.Check()

	// Test
	err := mm.Check()
	if !strings.Contains(
		err.Error(),
		expectedErrorComponent.Error(),
	) {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedErrorComponent,
			err,
		)
	}
}

/*
TestMonomialMatrix_Check4
Description:

	Tests that the Check() method properly returns nil when a matrix of
	Monomials is properly initialized. (in this case a 3 x 2 matrix of
	Monomials is given).
*/
func TestMonomialMatrix_Check4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{3},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1, m1},
		{m1, m1},
	}

	// Test
	if mm.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			mm.Check(),
		)
	}
}

/*
TestMonomialMatrix_Variables1
Description:

	Tests that the Variables() method properly returns the variables in a
	matrix of Monomials. (In this case the matrix of monomials
	is 3 x 2 and is composed of only one variable).
*/
func TestMonomialMatrix_Variables1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{3},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1, m1},
		{m1, m1},
	}
	expectedVariables := []symbolic.Variable{v1}

	// Test
	variables := mm.Variables()
	if len(variables) != len(expectedVariables) {
		t.Errorf(
			"expected Variables() to return %v; received %v",
			expectedVariables,
			variables,
		)
	}

}

/*
TestMonomialMatrix_Variables2
Description:

	Tests that the Variables() method properly returns the variables in a
	matrix of Monomials. (In this case the matrix of monomials
	is 3 x 2 and is composed of 5 entries that contain 1 variable
	and one entry that contains 3.).
*/
func TestMonomialMatrix_Variables2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{3},
	}
	m2 := symbolic.Monomial{
		Coefficient:     10.0,
		VariableFactors: []symbolic.Variable{v1, v2, v3},
		Degrees:         []int{3, 5, 9},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1, m1},
		{m2, m2},
	}
	expectedVariables := []symbolic.Variable{v1, v2, v3}

	// Test
	variables := mm.Variables()
	if len(variables) != len(expectedVariables) {
		t.Errorf(
			"expected Variables() to return %v; received %v",
			expectedVariables,
			variables,
		)
	}
}

/*
TestMonomialMatrix_Variables3
Description:

	Tests that the Variables() method properly panics
	when an improperly defined matrix of Monomials is given.
*/
func TestMonomialMatrix_Variables3(t *testing.T) {
	// Constants
	var mm1 symbolic.MonomialMatrix

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected Variables() to panic; it did not",
			)
		}
	}()

	mm1.Variables()

}

/*
TestMonomialMatrix_Dims1
Description:

	Tests that the Dims() method properly returns the dimensions of a
	matrix of Monomials. (In this case the matrix of monomials
	is 3 x 2).
*/
func TestMonomialMatrix_Dims1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{3},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1, m1},
		{m1, m1},
	}
	expectedDims := []int{3, 2}

	// Test
	dims := mm.Dims()
	if len(dims) != len(expectedDims) {
		t.Errorf(
			"expected Dims() to return %v; received %v",
			expectedDims,
			dims,
		)
	}

	// Check that the number of rows is correct
	if dims[0] != expectedDims[0] {
		t.Errorf(
			"expected Dims() to return %v rows; received %v",
			expectedDims,
			dims,
		)
	}

	// Check that the number of columns is correct
	if dims[1] != expectedDims[1] {
		t.Errorf(
			"expected Dims() to return %v columns; received %v",
			expectedDims,
			dims,
		)
	}
}

/*
TestMonomialMatrix_Plus1
Description:

	Tests that the Plus() method panics if the monomial matrix
	that it is called on is not well formed.
*/
func TestMonomialMatrix_Plus1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1},
		{m1, m1},
	}
	mm2 := mm

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected Plus() to panic; it did not",
			)
		}
	}()
	mm.Plus(mm2)
}

/*
TestMonomialMatrix_Plus2
Description:

	Tests that the Plus() method panics if the second expression
	input to the method is not well formed.
*/
func TestMonomialMatrix_Plus2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	var mm2 symbolic.MonomialMatrix

	expectedError := mm2.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Plus() to panic; it did not",
			)
		}

		// Check that the error is correct
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"expected Plus() to panic with an error; it panicked with %v",
				r,
			)
		}

		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"expected Plus() to panic with error \"%v\"; it panicked with \"%v\"",
				expectedError,
				err,
			)
		}

	}()
	mm.Plus(mm2)
}

/*
TestMonomialMatrix_Plus3
Description:

	Tests that the Plus() method panics if the two expressions
	are not the same size. (In this case the first expression
	is 3 x 2 and the second is 2 x 2).
*/
func TestMonomialMatrix_Plus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	var mm2 symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	expectedError := smErrors.DimensionError{
		Arg1:      mm,
		Arg2:      mm2,
		Operation: "Plus",
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Plus() to panic; it did not",
			)
		}

		// Check that the error is correct
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"expected Plus() to panic with an error; it panicked with %v",
				r,
			)
		}

		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"expected Plus() to panic with error \"%v\"; it panicked with \"%v\"",
				expectedError,
				err,
			)
		}

	}()
	mm.Plus(mm2)
}

/*
TestMonomialMatrix_Plus4
Description:

	Tests that the Plus() method properly adds a constant to a matrix
	of Monomials. (In this case, the monomial matrix is a 2 x2 matrix
	of single variables and the constant is 1.0).
*/
func TestMonomialMatrix_Plus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	f2 := 1.0

	// Test
	sum := mm.Plus(f2)

	// Check that the sum is of the PolynomialMatrix type
	_, ok := sum.(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"expected Plus() to return a PolynomialMatrix; received %v",
			sum,
		)
	}

	// Check that each entry in the sum:
	// 1. Contains 2 monomials
	// 2. Contains the original monomial
	// 3. Contains the constant monomial
	for _, row := range sum.(symbolic.PolynomialMatrix) {
		for _, polynomial := range row {
			if len(polynomial.Monomials) != 2 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with 2 monomials; received %v",
					polynomial.Monomials,
				)
			}

			// Check that the first monomial is the original monomial
			if v1Index := polynomial.VariableMonomialIndex(v1); v1Index == -1 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
					polynomial,
				)
			}

			// Check that the second monomial is the constant monomial
			if constIndex := polynomial.ConstantMonomialIndex(); polynomial.Monomials[constIndex].Coefficient != 1.0 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the constant monomial; received %v",
					polynomial.Monomials[1],
				)
			}
		}
	}
}

/*
TestMonomialMatrix_Plus5
Description:

	Tests that the Plus() method properly adds a matrix of Monomials
	to a unique variable object. The result should be a matrix of polynomials
	where each polynomial has two terms. We will check the number of terms
	in each polynomial.
*/
func TestMonomialMatrix_Plus5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}

	v2 := symbolic.NewVariable()

	// Test
	sum := mm.Plus(v2)

	// Check that the sum is of the PolynomialMatrix type
	_, ok := sum.(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"expected Plus() to return a PolynomialMatrix; received %v",
			sum,
		)
	}

	// Check that each entry in the sum:
	// 1. Contains 2 monomials
	// 2. Contains the original monomial
	// 3. Contains the constant monomial
	for _, row := range sum.(symbolic.PolynomialMatrix) {
		for _, polynomial := range row {
			if len(polynomial.Monomials) != 2 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with 2 monomials; received %v",
					polynomial.Monomials,
				)
			}

			// Check that the first monomial is the original monomial
			if v1Index := polynomial.VariableMonomialIndex(v1); v1Index == -1 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
					polynomial,
				)
			}

		}
	}
}

/*
TestMonomialMatrix_Plus6
Description:

	Tests that the Plus() method properly adds a matrix of Monomials
	to a monomial. The result should be a matrix of polynomials
	where each polynomial has two terms. We will check the number of terms
	in each polynomial.
*/
func TestMonomialMatrix_Plus6(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}

	m2 := symbolic.NewVariable().ToMonomial()

	// Test
	sum := mm.Plus(m2)

	// Check that the sum is of the PolynomialMatrix type
	_, ok := sum.(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"expected Plus() to return a PolynomialMatrix; received %v",
			sum,
		)
	}

	// Check that each entry in the sum:
	// 1. Contains 2 monomials
	// 2. Contains the original monomial
	// 3. Contains the constant monomial
	for _, row := range sum.(symbolic.PolynomialMatrix) {
		for _, polynomial := range row {
			if len(polynomial.Monomials) != 2 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with 2 monomials; received %v",
					polynomial.Monomials,
				)
			}

			// Check that the first monomial is the original monomial
			if v1Index := polynomial.VariableMonomialIndex(v1); v1Index == -1 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
					polynomial,
				)
			}

		}
	}
}

/*
TestMonomialMatrix_Plus7
Description:

	Tests that the Plus() method properly adds a matrix of Monomials
	to a polynomial. The result should be a matrix of polynomials
	where each polynomial has three terms. We will check the number
	of terms in each polynomial.
*/
func TestMonomialMatrix_Plus7(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}

	m2 := symbolic.NewVariable().ToMonomial()
	m3 := symbolic.NewVariable().ToMonomial()

	p4 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m2, m3},
	}

	// Test
	sum := mm.Plus(p4)

	// Check that the sum is of the PolynomialMatrix type
	_, ok := sum.(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"expected Plus() to return a PolynomialMatrix; received %v",
			sum,
		)
	}

	// Check that each entry in the sum:
	// 1. Contains 2 monomials
	// 2. Contains the original monomial
	// 3. Contains the constant monomial
	for _, row := range sum.(symbolic.PolynomialMatrix) {
		for _, polynomial := range row {
			if len(polynomial.Monomials) != 3 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with 2 monomials; received %v",
					polynomial.Monomials,
				)
			}

			// Check that the first monomial is the original monomial
			if v1Index := polynomial.VariableMonomialIndex(v1); v1Index == -1 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
					polynomial,
				)
			}

		}
	}
}

/*
TestMonomialMatrix_Plus8
Description:

	Tests that the Plus() method properly adds a matrix of Monomials
	to a polynomial matrix. The result should be a matrix of polynomials
	where each polynomial has three terms. We will check the number
	of terms in each polynomial.
*/
func TestMonomialMatrix_Plus8(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}

	m2 := symbolic.NewVariable().ToMonomial()
	m3 := symbolic.NewVariable().ToMonomial()

	p4 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m2, m3},
	}

	var pm5 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p4, p4},
		{p4, m2.Plus(3.0).(symbolic.Polynomial)},
	}

	// Test
	sum := mm.Plus(pm5)

	// Check that the sum is of the PolynomialMatrix type
	_, ok := sum.(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"expected Plus() to return a PolynomialMatrix; received %v",
			sum,
		)
	}

	// Check that each entry in the sum:
	// 1. Contains 2 monomials
	// 2. Contains the original monomial
	// 3. Contains the constant monomial
	for ii, row := range sum.(symbolic.PolynomialMatrix) {
		for jj, polynomial := range row {
			if len(polynomial.Monomials) != 3 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with 2 monomials; received %v",
					polynomial.Monomials,
				)
			}

			// Check that the first monomial is the original monomial
			if v1Index := polynomial.VariableMonomialIndex(v1); v1Index == -1 {
				t.Errorf(
					"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
					polynomial,
				)
			}

			if ii == 1 && jj == 1 {
				// There should be a constant in this polynomial.
				if constantIndex := polynomial.ConstantMonomialIndex(); constantIndex == -1 {
					t.Errorf(
						"expected Plus() to return a PolynomialMatrix with the original monomial; received %v",
						polynomial,
					)
				}
			}

		}
	}
}

/*
TestMonomialMatrix_Multiply1
Description:

	Tests that the Multiply() method panics if the monomial matrix
	that it is called on is not well formed.
*/
func TestMonomialMatrix_Multiply1(t *testing.T) {
	// Constants
	var mm1 symbolic.MonomialMatrix

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected Multiply() to panic; it did not",
			)
		}
	}()
	mm1.Multiply(mm1)
}

/*
TestMonomialMatrix_Multiply2
Description:

	Tests that the Multiply() method panics if the second expression
	input to the method is not well formed.
*/
func TestMonomialMatrix_Multiply2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	var mm2 symbolic.MonomialMatrix

	expectedError := mm2.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Multiply() to panic; it did not",
			)
		}

		// Check that the error is correct
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"expected Multiply() to panic with an error; it panicked with %v",
				r,
			)
		}

		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"expected Multiply() to panic with error \"%v\"; it panicked with \"%v\"",
				expectedError,
				err,
			)
		}

	}()
	mm.Multiply(mm2)
}

/*
TestMonomialMatrix_Multiply3
Description:

	Tests that the Multiply() method panics if the two expressions
	do not have the matching sizes for matrix multiplication (i.e.,
	the number of columns in the matrix of monomials doesn't match
	the number of rows in the second expression).
*/
func TestMonomialMatrix_Multiply3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{v1.ToMonomial(), v1.ToMonomial()},
		{v1.ToMonomial(), v1.ToMonomial()},
	}
	km2 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 2))

	expectedError := smErrors.DimensionError{
		Arg1:      mm,
		Arg2:      km2,
		Operation: "Multiply",
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Multiply() to panic; it did not",
			)
		}

		// Check that the error is correct
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"expected Multiply() to panic with an error; it panicked with %v",
				r,
			)
		}

		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"expected Multiply() to panic with error \"%v\"; it panicked with \"%v\"",
				expectedError,
				err,
			)
		}

	}()
	mm.Multiply(km2)
}

/*
TestMonomialMatrix_Multiply4
Description:

	Tests that the Multiply() method properly multiplies a matrix
	of Monomials with a constant. The result should be a matrix of
	monomials where each monomial has the scaled coefficients
	of the original monomial.
*/
func TestMonomialMatrix_Multiply4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := v1.ToMonomial()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1},
		{m1, m1},
	}
	f2 := 2.0

	// Test
	product := mm.Multiply(f2)

	// Check that the product is of the MonomialMatrix type
	productMat, ok := product.(symbolic.MonomialMatrix)
	if !ok {
		t.Errorf(
			"expected Multiply() to return a MonomialMatrix; received %v",
			product,
		)
	}

	// Check that each entry in the product:
	// 1. Contains the original monomial
	// 2. Contains the constant monomial
	for ii, row := range productMat {
		for jj, monomial := range row {
			// Check that the coefficient of this monomial is f2 * mm1[ii][jj].Coefficient
			if monomial.Coefficient != f2*m1.Coefficient {
				t.Errorf(
					"expected Multiply() to return a MonomialMatrix with coefficient %v at (%v,%v); received %v",
					f2*m1.Coefficient,
					ii,
					jj,
					monomial.Coefficient,
				)
			}
		}
	}
}

/*
TestMonomialMatrix_Multiply5
Description:

	Tests that the Multiply() method properly multiplies a matrix
	of Monomials with a vector of variables. The result should be a matrix of
	monomials where each polynomial has a number of monomials
	equal to the number of values in the variable vector.
*/
func TestMonomialMatrix_Multiply5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	m1 := v1.ToMonomial()
	m2 := v2.ToMonomial()
	m3 := v3.ToMonomial()
	var mm symbolic.MonomialMatrix = [][]symbolic.Monomial{
		{m1, m1, m1},
		{m2, m2, m2},
		{m3, m3, m3},
	}
	var variableVector symbolic.VariableVector = []symbolic.Variable{v1, v2, v3}

	// Test
	product := mm.Multiply(variableVector)

	// Check that the dimensions of the product are (3,1)
	if dims := product.Dims(); dims[0] != 3 || dims[1] != 1 {
		t.Errorf(
			"expected Multiply() to return a MonomialMatrix with dimensions (3,1); received %v",
			dims,
		)
	}

	// Check that the product is of the PolynomialMatrix type
	productVec, ok := product.(symbolic.PolynomialVector)
	if !ok {
		t.Errorf(
			"expected Multiply() to return a PolynomialMatrix; received %v",
			product,
		)
	}

	// Check that each entry in the product:
	// 1. Does not contains the original monomials
	for jj, polynomial := range productVec {
		// Check that the number of monomials in this polynomial is equal to the number expected
		if len(polynomial.Monomials) != 3 {
			t.Errorf(
				"expected Multiply() to return a PolynomialMatrix with %v monomials at index %v; received %v",
				len(variableVector),
				jj,
				len(polynomial.Monomials),
			)
		}

		// Check that each of the original monomials is NOT in this polynomial
		for _, monomial := range []symbolic.Monomial{m1, m2, m3} {
			if monomialIndex := polynomial.MonomialIndex(monomial); monomialIndex != -1 {
				t.Errorf(
					"expected Multiply() to return a PolynomialMatrix without monomial %v at index %v; received %v",
					monomial,
					monomialIndex,
					polynomial,
				)
			}
		}
	}
}
