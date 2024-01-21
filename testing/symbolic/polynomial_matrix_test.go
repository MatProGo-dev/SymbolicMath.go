package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
polynomial_matrix_test.go
Description:

	Tests for the functions mentioned in the polynomial_matrix.go file.
*/

/*
TestPolynomialMatrix_Check1
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Polynomials (i.e., no polynomials are given).
*/
func TestPolynomialMatrix_Check1(t *testing.T) {
	// Constants
	pm := symbolic.PolynomialMatrix{}
	expectedError := smErrors.EmptyMatrixError{pm}

	// Test
	err := pm.Check()
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}
}

/*
TestPolynomialMatrix_Check2
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Polynomials (i.e., a polynomial is given with an improper number of
	columns in row 2 than in every other row.)
*/
func TestPolynomialMatrix_Check2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}
	p1 := symbolic.Polynomial{
		[]symbolic.Monomial{m1},
	}
	var pm symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1},
		{p1, p1},
	}

	expectedError := smErrors.MatrixColumnMismatchError{
		ExpectedNColumns: 2,
		ActualNColumns:   1,
		Row:              1,
	}

	// Test
	err := pm.Check()
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}
}

/*
TestPolynomialMatrix_Check3
Description:

	Tests that the Check() method properly catches an improperly initialized
	matrix of Polynomials (i.e., a polynomial is given with an improper number of
	degrees in a monomial in the third row.)
*/
func TestPolynomialMatrix_Check3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}

	// Create polynomials
	p1 := v1.ToPolynomial()
	p2 := symbolic.Polynomial{[]symbolic.Monomial{m2}}

	// Construct matrix
	var pm symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p2},
	}

	// Test
	err := pm.Check()
	if !strings.Contains(
		err.Error(),
		"2,1", // coordinate of bad polynomial
	) {
		t.Errorf(
			"expected Check() to return error containing \"%v\"; received %v",
			"2,1",
			err,
		)
	}
}

/*
TestPolynomialMatrix_Check4
Description:

	Tests that the Check() method properly returns nil when a matrix of
	Polynomials is properly initialized. (in this case a 3 x 2 matrix of
	Polynomials is given).
*/
func TestPolynomialMatrix_Check4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()

	// Construct matrix
	var pm symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	if pm.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			pm.Check(),
		)
	}
}

/*
TestPolynomialMatrix_Variables1
Description:

	Tests that the Variables() method properly returns the variables in a
	PolynomialMatrix. In this case, the polynomials all contain
	a single variable and so the length of Variables should be 1.
*/
func TestPolynomialMatrix_Variables1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()

	// Construct matrix
	var pm symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	vars := pm.Variables()
	if len(vars) != 1 {
		t.Errorf(
			"expected len(vars) to be 1; received %v",
			len(vars),
		)
	}
}

/*
TestPolynomialMatrix_Dims1
Description:

	Tests that the Dims() method properly returns the dimensions of a
	PolynomialMatrix. In this case, the matrix is 3 x 2.
*/
func TestPolynomialMatrix_Dims1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()

	// Construct matrix
	var pm symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	dim1, dim2 := pm.Dims()[0], pm.Dims()[1]
	if dim1 != 3 {
		t.Errorf(
			"expected pm.Dims() to be [3,2]; received %v rows",
			dim1,
		)
	}

	if dim2 != 2 {
		t.Errorf(
			"expected pm.Dims() to be [3,2]; received %v columns",
			dim2,
		)
	}
}

/*
TestPolynomialMatrix_Plus1
Description:

	Tests that the Plus() panics when an improperly initialized
	matrix of polynomials is used to call it..
*/
func TestPolynomialMatrix_Plus1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Plus(pm) to panic; received %v",
				pm.Plus(pm),
			)
		}
	}()

	pm.Plus(pm)
}

/*
TestPolynomialMatrix_Plus2
Description:

	Tests that the Plus() method properly panics an error if two
	matrices of different dimensions are given to the method.
*/
func TestPolynomialMatrix_Plus2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}
	var pm2 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Plus(pm2) to panic; received %v",
				pm1.Plus(pm2),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if rAsError.Error() != (smErrors.DimensionError{
			Operation: "Plus",
			Arg1:      pm1,
			Arg2:      pm2,
		}).Error() {
			t.Errorf(
				"expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	// Test
	pm1.Plus(pm2)

}
