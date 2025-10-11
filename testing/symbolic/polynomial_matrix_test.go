package symbolic_test

import (
	"strings"
	"testing"

	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
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
		Exponents:       []int{1, 2},
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
		Exponents:       []int{1, 2},
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
TestPolynomialMatrix_Variables2
Description:

	Test that the variables() method properly panics
	when applied to an improperly initialized matrix of polynomials.
*/
func TestPolynomialMatrix_Variables2(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Variables() to panic; received %v",
				pm.Variables(),
			)
		}
	}()

	pm.Variables()
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
TestPolynomialMatrix_Dims2
Description:

	Tests that the Dims() method properly panics when applied to an
	improperly initialized matrix of polynomials.
*/
func TestPolynomialMatrix_Dims2(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Dims() to panic; received %v",
				pm.Dims(),
			)
		}
	}()

	pm.Dims()
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

		if rAsError.Error() != (smErrors.MatrixDimensionError{
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

/*
TestPolynomialMatrix_Plus3
Description:

	Tests that the Plus() method properly adds a polynomial matrix
	with a float64. The result should be a polynomial matrix with
	each polynomial containing two monomials.
*/
func TestPolynomialMatrix_Plus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	pm2 := pm1.Plus(1.0)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	for _, pmRow := range pm2AsPM {
		for _, p := range pmRow {
			if len(p.Monomials) != 2 {
				t.Errorf(
					"expected len(p.Monomials) to be 2; received %v",
					len(p.Monomials),
				)
			}
		}
	}
}

/*
TestPolynomialMatrix_Plus4
Description:

	Tests that the Plus() method properly adds a polynomial matrix (really a variable matrix)
	with THE SAME polynomial matrix. The result should be a monomial
	matrix.
*/
func TestPolynomialMatrix_Plus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	pm2 := pm1.Plus(pm1)

	_, tf := pm2.(symbolic.MonomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a MonomialMatrix; received %v",
			pm2,
		)
	}
}

/*
TestPolynomialMatrix_Plus5
Description:

	Tests that the Plus() method properly adds a polynomial matrix
	with a polynomial. The result should be a polynomial matrix with
	each polynomial containing three monomials.
*/
func TestPolynomialMatrix_Plus5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	p2 := v1.Plus(symbolic.NewVariable()).Plus(symbolic.NewVariable()).(symbolic.Polynomial)

	// Test
	pm2 := pm1.Plus(p2)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	for _, pmRow := range pm2AsPM {
		for _, p := range pmRow {
			if len(p.Monomials) != 3 {
				t.Errorf(
					"expected len(p.Monomials) to be 3; received %v",
					len(p.Monomials),
				)
			}
		}
	}
}

/*
TestPolynomialMatrix_Plus6
Description:

	Tests that the Plus() method properly adds a polynomial matrix
	to a monomial. The result should be a polynomial matrix with
	each polynomial containing two monomials.
*/
func TestPolynomialMatrix_Plus6(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	m1 := symbolic.NewVariable().ToMonomial()

	// Test
	pm2 := pm1.Plus(m1)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	for _, pmRow := range pm2AsPM {
		for _, p := range pmRow {
			if len(p.Monomials) != 2 {
				t.Errorf(
					"expected len(p.Monomials) to be 2; received %v",
					len(p.Monomials),
				)
			}
		}
	}
}

/*
TestPolynomialMatrix_Plus7
Description:

	Tests that the Plus() method properly adds a polynomial matrix
	to a matrix of constants (K). The result should be a polynomial
	matrix with each polynomial containing two monomials.
*/
func TestPolynomialMatrix_Plus7(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	km1 := getKMatrix.From(symbolic.OnesMatrix(3, 2))

	// Test
	pm2 := pm1.Plus(km1)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	for _, pmRow := range pm2AsPM {
		for _, p := range pmRow {
			if len(p.Monomials) != 2 {
				t.Errorf(
					"expected len(p.Monomials) to be 2; received %v",
					len(p.Monomials),
				)
			}
		}
	}
}

/*
TestPolynomialMatrix_Plus8
Description:

	Tests that the Plus() method properly panics
	when a valid polynomial matrix is added to an
	invalid expression (in this case, a variable).
*/
func TestPolynomialMatrix_Plus8(t *testing.T) {
	// Constants
	v1 := symbolic.Variable{}
	p1 := symbolic.NewVariable().ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Plus(v1) to panic; received %v",
				pm1.Plus(v1),
			)
		}

		// Check that the panic includes a Check() error
		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		expectedError := v1.Check()
		if expectedError == nil {
			t.Errorf(
				"expected v1.Check() to return an error; received nil",
			)
			return
		}

		if !strings.Contains(
			rAsError.Error(),
			v1.Check().Error(),
		) {
			t.Errorf(
				"expected r to be a Check() error; received %v",
				r,
			)
		}
	}()

	// Test
	pm1.Plus(v1)
}

/*
TestPolynomialMatrix_Plus9
Description:

	Tests that the Plus() method properly panics
	when a valid polynomial matrix is added to an
	invalid object (a string).
*/
func TestPolynomialMatrix_Plus9(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Plus(\"hi\") to panic; received %v",
				pm1.Plus("hi"),
			)
		}
	}()

	// Test
	pm1.Plus("hi")
}

/*
TestPolynomialMatrix_Minus1
Description:

	Tests that the Minus() method properly panics when an improperly
	initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_Minus1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Minus(pm) to panic; received %v",
				pm.Minus(pm),
			)
		}
	}()

	pm.Minus(pm)
}

/*
TestPolynomialMatrix_Minus2
Description:

	Tests that the Minus() method properly panics an error if two
	matrices of different dimensions are given to the method.
*/
func TestPolynomialMatrix_Minus2(t *testing.T) {
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
				"expected pm1.Minus(pm2) to panic; received %v",
				pm1.Minus(pm2),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if rAsError.Error() != (smErrors.MatrixDimensionError{
			Operation: "Minus",
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
	pm1.Minus(pm2)
}

/*
TestPolynomialMatrix_Minus3
Description:

	Tests that the Minus() method properly panics when a well-defined PolynomialMatrix
	calls Miinus() on a Monomial that is NOT well-defined.
*/
func TestPolynomialMatrix_Minus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	m1 := symbolic.Monomial{
		VariableFactors: []symbolic.Variable{v1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Minus(m1) to panic; received %v",
				pm1.Minus(m1),
			)
		}
	}()

	// Test
	pm1.Minus(m1)
}

/*
TestPolynomialMatrix_Minus4
Description:

	Tests that the Minus() method properly subtracts a polynomial matrix (Each with 1 monomial)
	with a float64. The result should be a polynomial vector with
	each polynomial containing two monomials.
*/
func TestPolynomialMatrix_Minus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1},
		{p1},
		{p1},
	}

	// Test
	pm2 := pm1.Minus(1.0)

	// Check that the result is a polynomial vector
	pv3, tf := pm2.(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"expected pv3 to be a PolynomialVector; received %T",
			pv3,
		)
	}

	for _, p := range pv3 {
		if len(p.Monomials) != 2 {
			t.Errorf(
				"expected len(p.Monomials) to be 2; received %v",
				len(p.Monomials),
			)
		}
	}
}

/*
TestPolynomialMatrix_Minus5
Description:

	Tests that the Minus() method properly panics when a well-defined polynomial matrix
	calls the Minus() method on an object that is not an expression (in this case, a string).
*/
func TestPolynomialMatrix_Minus5(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1},
		{p1},
		{p1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Minus(\"hi\") to panic; received %v",
				pm1.Minus("hi"),
			)
		}
	}()

	// Test
	pm1.Minus("hi")
}

/*
TestPolynomialMatrix_Multiply1
Description:

	Tests that the Multiply() method properly panics when an improperly
	initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_Multiply1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Multiply(pm) to panic; received %v",
				pm.Multiply(pm),
			)
		}
	}()

	pm.Multiply(pm)
}

/*
TestPolynomialMatrix_Multiply2
Description:

	Tests that the Multiply() method properly panics an error if it
	is called with a properly defined matrix of Polynomials and an improperly
	initialized matrix of polynomials.
*/
func TestPolynomialMatrix_Multiply2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
	}

	var pm2 symbolic.PolynomialMatrix

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Multiply(pm2) to panic; received %v",
				pm1.Multiply(pm2),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if !strings.Contains(
			rAsError.Error(),
			pm2.Check().Error(),
		) {
			t.Errorf(
				"expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	// Test
	pm1.Multiply(pm2)
}

/*
TestPolynomialMatrix_Multiply3
Description:

	Tests that the Multiply() method properly panics when a polynomial matrix
	with a KVector (vector of constants) that does not match in size.
	A DimensionError should be thrown.
*/
func TestPolynomialMatrix_Multiply3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
	}

	kv1 := getKVector.From([]float64{1.0, 2.1, 3.14})

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Multiply(km1) to panic; received %v",
				pm1.Multiply(kv1),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if !strings.Contains(
			rAsError.Error(),
			smErrors.MatrixDimensionError{
				Operation: "Multiply",
				Arg1:      pm1,
				Arg2:      kv1,
			}.Error(),
		) {
			t.Errorf(
				"expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	// Test
	pm1.Multiply(kv1)
}

/*
TestPolynomialMatrix_Multiply4
Description:

	Tests that the Multiply() method properly multiplies a polynomial matrix
	with a single constant. The result should be a monomial vector with
	the same number of monomials as the original. But the coefficient
	of each monomial should be multiplied by the constant.
*/
func TestPolynomialMatrix_Multiply4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1},
		{p1},
		{p1},
	}

	// Test
	prod := pm1.Multiply(3.14)

	// Verify that this is a monomial Vector
	mv3, tf := prod.(symbolic.MonomialVector)
	if !tf {
		t.Errorf(
			"expected mv3 to be a MonomialVector; received %T",
			mv3,
		)
	}

	// Check the elements of each monomial
	for _, monomial := range mv3 {
		if monomial.Coefficient != 3.14 {
			t.Errorf(
				"expected monomial.Coefficient to be 3.14; received %v",
				monomial.Coefficient,
			)
		}
	}
}

/*
TestPolynomialMatrix_Multiply5
Description:

	Tests that the Multiply() method properly multiplies a polynomial matrix
	(1 x 3) with a vector of variables (VariableVector).
	The output should be a Polynomial. It should have three terms.
*/
func TestPolynomialMatrix_Multiply5(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	N := 3
	vv1 := symbolic.NewVariableVector(N)

	// Test
	p2 := pm1.Multiply(vv1)

	p2AsP, tf := p2.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected p2 to be a Polynomial; received %v",
			p2,
		)
	}

	if len(p2AsP.Monomials) != 3 {
		t.Errorf(
			"expected len(p2.Monomials) to be 3; received %v",
			len(p2AsP.Monomials),
		)
	}

	// Check each monomial and verify that each monomial has
	// - 2 variable factors
	// - 2 indices (each are 1)
	for _, m := range p2AsP.Monomials {
		// 2 VariableFactors
		if len(m.VariableFactors) != 2 {
			t.Errorf(
				"expected len(m.VariableFactors) to be 2; received %v",
				len(m.VariableFactors),
			)
		}
		// 2 indices (both are 1)
		if m.Exponents[0] != 1 || m.Exponents[1] != 1 {
			t.Errorf(
				"expected m.Exponents to be [1,1]; received %v",
				m.Exponents,
			)
		}
	}
}

/*
TestPolynomialMatrix_Multiply6
Description:

	Tests that the Multiply() method properly multiplies a polynomial matrix
	(2 x 3) with a vector of variables (VariableVector).
	The output should be a PolynomialVector.
	It should have two rows and each row should contain three monomials.
*/
func TestPolynomialMatrix_Multiply6(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	N := 3
	vv1 := symbolic.NewVariableVector(N)

	// Test
	pv2 := pm1.Multiply(vv1)

	pv2AsPV, tf := pv2.(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"expected pv2 to be a PolynomialVector; received %v",
			pv2,
		)
	}

	if len(pv2AsPV) != 2 {
		t.Errorf(
			"expected len(pv2) to be 2; received %v",
			len(pv2AsPV),
		)
	}

	for _, p := range pv2AsPV {
		if len(p.Monomials) != 3 {
			t.Errorf(
				"expected len(p.Monomials) to be 3; received %v",
				len(p.Monomials),
			)
		}
	}
}

/*
TestPolynomialMatrix_Multiply7
Description:

	Tests that the Multiply() method properly panics when a
	well-defined polynomial matrix is multiplied by an invalid
	object (in this case, a string).
*/
func TestPolynomialMatrix_Multiply7(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Multiply(\"hi\") to panic; received %v",
				pm1.Multiply("hi"),
			)
		}
	}()

	// Test
	pm1.Multiply("hi")
}

/*
TestPolynomialMatrix_Multiply8
Description:

	Tests that the Multiply() method properly panics when a
	well-defined polynomial matrix is multiplied by an invalid
	object (in this case, a VariableMatrix).
*/
func TestPolynomialMatrix_Multiply8(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	var vm1 symbolic.VariableMatrix

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Multiply(vm1) to panic; received %v",
				pm1.Multiply(vm1),
			)
		}
	}()

	// Test
	pm1.Multiply(vm1)
}

/*
TestPolynomialMatrix_Multiply9
Description:

	Tests that the Multiply() method properly computes the product of a
	polynomial matrix (2 x 3) with a constant matrix (3 x 2).
	The output should be a polynomial matrix (2 x 2).
*/
func TestPolynomialMatrix_Multiply9(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	km1 := getKMatrix.From([][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
		{5.0, 6.0},
	})

	// Test
	pm2 := pm1.Multiply(km1)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	if pm2AsPM.Dims()[0] != 2 || pm2AsPM.Dims()[1] != 2 {
		t.Errorf(
			"expected pm2.Dims() to be [2,2]; received %v",
			pm2AsPM.Dims(),
		)
	}
}

/*
TestPolynomialMatrix_Transpose1
Description:

	Tests that the Transpose() method properly transposes a 3 x 2
	polynomial matrix into a 2 x 3 polynomial matrix.
	We'll also check the unique elements of each of these
	matrices to ensure that the transpose was done properly.
*/
func TestPolynomialMatrix_Transpose1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{v1.Plus(1.0).(symbolic.Polynomial), v1.Plus(2.0).(symbolic.Polynomial)},
		{v1.Plus(3.0).(symbolic.Polynomial), v1.Plus(4.0).(symbolic.Polynomial)},
		{v1.Plus(5.0).(symbolic.Polynomial), v1.Plus(6.0).(symbolic.Polynomial)},
	}

	// Test
	pm2 := pm1.Transpose()

	// Verify that pm2 is a PolynomialMatrix
	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	// Verify that the dimensions are correct
	if pm2AsPM.Dims()[0] != 2 || pm2AsPM.Dims()[1] != 3 {
		t.Errorf(
			"expected pm2.Dims() to be [2,3]; received %v",
			pm2AsPM.Dims(),
		)
	}

	// Verify that the each element is correct
	for ii, pmRow := range pm2AsPM {
		for jj, p := range pmRow {
			expectedConstant := float64(ii + 2*jj + 1)
			if p.Constant() != expectedConstant {
				t.Errorf(
					"expected p.Constant() to be %v; received %v",
					expectedConstant,
					p.Constant(),
				)
			}
		}
	}
}

/*
TestPolynomialMatrix_Transpose2
Description:

	Tests that the Transpose() method properly panics when an improperly
	initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_Transpose2(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Transpose() to panic; received %v",
				pm.Transpose(),
			)
		}
	}()

	pm.Transpose()
}

/*
TestPolynomialMatrix_LessEq1
Description:

	Tests that the LessEq() method panics
	when an improperly initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_LessEq1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.LessEq(pm) to panic; received %v",
				pm.LessEq(pm),
			)
		}
	}()

	pm.LessEq(pm)
}

/*
TestPolynomialMatrix_GreaterEq1
Description:

	Tests that the GreaterEq() method properly returns
	a MatrixConstraint when a valid polynomial matrix is used to call it
	with a constant matrix.
*/
func TestPolynomialMatrix_GreaterEq1(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Test
	mc := pm1.GreaterEq(getKMatrix.From(symbolic.ZerosMatrix(2, 2)))

	// Check the type of mc
	_, tf := mc.(symbolic.MatrixConstraint)
	if !tf {
		t.Errorf(
			"expected mc to be a MatrixConstraint; received %v",
			mc,
		)
	}
}

/*
TestPolynomialMatrix_GreaterEq2
Description:

	Tests that the GreaterEq() method properly panics
	when a well-defined PolynomialMatrix is compared
	with an improperly-defined expression (in this case,
	a VariableMatrix).
*/
func TestPolynomialMatrix_GreaterEq2(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	var vm2 symbolic.VariableMatrix

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.GreaterEq(pm1) to panic; received %v",
				pm1.GreaterEq(pm1),
			)
		}
	}()

	// Test
	pm1.GreaterEq(vm2)
}

/*
TestPolynomialMatrix_Eq1
Description:

	Tests that the Eq() method properly returns a
	MatrixConstraint when a valid polynomial matrix is used to call it
	with a *mat.Dense object.
*/
func TestPolynomialMatrix_Eq1(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	d2 := mat.NewDense(
		2, 2,
		[]float64{1, 2, 3, 4},
	)

	// Test
	mc := pm1.Eq(d2)

	// Check the type of mc
	_, tf := mc.(symbolic.MatrixConstraint)
	if !tf {
		t.Errorf(
			"expected mc to be a MatrixConstraint; received %v",
			mc,
		)
	}

	// Verify that the type of the right hand side of mc
	// is of type KMatrix
	mcAsMC := mc.(symbolic.MatrixConstraint)
	_, tf = mcAsMC.RightHandSide.(symbolic.KMatrix)
	if !tf {
		t.Errorf(
			"expected mc.Right to be a KMatrix; received %v",
			mcAsMC.RightHandSide,
		)
	}
}

/*
TestPolynomialMatrix_Eq2
Description:

	Tests that the Eq() method properly panics
	when a well-defined PolynomialMatrix is compared
	with a well-defined matrix of constants when
	the dimensions do not match.
*/
func TestPolynomialMatrix_Eq2(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	d2 := mat.NewDense(
		3, 2,
		[]float64{1, 2, 3, 4, 5, 6},
	)

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Eq(d2) to panic; received %v",
				pm1.Eq(d2),
			)
		}
	}()

	// Test
	pm1.Eq(d2)
}

/*
TestPolynomialMatrix_Eq3
Description:

	Tests that the Eq() method properly panics
	when a well-defined PolynomialMatrix is compared
	with an invalid object (in this case a string)
*/
func TestPolynomialMatrix_Eq3(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.Eq(\"hi\") to panic; received %v",
				pm1.Eq("hi"),
			)
		}
	}()

	// Test
	pm1.Eq("hi")
}

/*
TestPolynomialMatrix_Comparison1
Description:

	Tests that the Comparison() method properly returns
	a MatrixConstraint when a valid polynomial matrix is used to call it
	with a float64.
*/
func TestPolynomialMatrix_Comparison1(t *testing.T) {
	// Constants
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Test
	mc := pm1.Comparison(3.14, symbolic.SenseLessThanEqual)

	// Check the type of mc
	_, tf := mc.(symbolic.MatrixConstraint)
	if !tf {
		t.Errorf(
			"expected mc to be a MatrixConstraint; received %v",
			mc,
		)
	}

	// Verifies that the right hand side of mc is a KMatrix
	mcAsMC := mc.(symbolic.MatrixConstraint)
	_, tf = mcAsMC.RightHandSide.(symbolic.KMatrix)
	if !tf {
		t.Errorf(
			"expected mc.Right to be a KMatrix; received %v",
			mcAsMC.RightHandSide,
		)
	}
}

/*
TestPolynomialMatrix_DerivativeWrt1
Description:

	Tests that the DerivativeWrt() method properly panics
	when an improperly initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_DerivativeWrt1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix
	v1 := symbolic.NewVariable()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.DerivativeWrt(pm) to panic; received %v",
				pm.DerivativeWrt(v1),
			)
		}
	}()

	pm.DerivativeWrt(v1)
}

/*
TestPolynomialMatrix_DerivativeWrt2
Description:

	Tests that the DerivativeWrt() method properly panics when a
	well-defined polynomial matrix is used to call it with an
	a variable that is not well-defiend.
*/
func TestPolynomialMatrix_DerivativeWrt2(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix = symbolic.NewVariableMatrix(3, 2).ToPolynomialMatrix()
	v1 := symbolic.Variable{}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.DerivativeWrt(\"hi\") to panic; received %v",
				pm.DerivativeWrt(v1),
			)
		}
	}()

	// Test
	pm.DerivativeWrt(v1)
}

/*
TestPolynomialMatrix_DerivativeWrt3
Description:

	Tests that the DerivativeWrt() method properly returns a
	PolynomialMatrix when a valid polynomial matrix is used to call it
	with a valid variable. In this case, each polynomial in the matrix
	contains two monomials. The result should be a polynomial matrix
	with each polynomial containing one monomial.
*/
func TestPolynomialMatrix_DerivativeWrt3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
	}

	// Test
	pm2 := pm1.DerivativeWrt(v1)

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	// Check that each polynomial in pm2 contains one monomial
	for _, pm2Row := range pm2AsPM {
		for _, p := range pm2Row {
			if len(p.Monomials) != 1 {
				t.Errorf(
					"expected len(p.Monomials) to be 1; received %v",
					len(p.Monomials),
				)
			}
		}
	}

}

/*
TestPolynomialMatrix_At1
Description:

	Tests that the At() method properly panics when an improperly
	initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_At1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.At(0,0) to panic; received %v",
				pm.At(0, 0),
			)
		}
	}()

	pm.At(0, 0)
}

/*
TestPolynomialMatrix_At2
Description:

	Tests that the At() method properly panics when one attempts to get
	an element from a well-defined matrix of polynomials that is out of bounds.
*/
func TestPolynomialMatrix_At2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Create panic checking logic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm1.At(3,0) to panic; received %v",
				pm1.At(3, 0),
			)
		}
	}()

	// Test
	pm1.At(3, 0)
}

/*
TestPolynomialMatrix_Constant1
Description:

	Tests that the Constant() method properly panics when an improperly
	initialized matrix of polynomials is used to call it.
*/
func TestPolynomialMatrix_Constant1(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected pm.Constant() to panic; received %v",
				pm.Constant(),
			)
		}
	}()

	pm.Constant()
}

/*
TestPolynomialMatrix_Constant2
Description:

	Tests that the Constant() method properly returns the constant
	value of a well-defined polynomial matrix containing ones.
*/
func TestPolynomialMatrix_Constant2(t *testing.T) {
	// Constants
	var pm symbolic.PolynomialMatrix = symbolic.NewVariableMatrix(3, 2).ToPolynomialMatrix().Plus(
		1.0,
	).(symbolic.PolynomialMatrix)

	// Test
	constant0 := pm.Constant()

	// Verify that matrix is almost equal to matrix of all ones.
	ones := symbolic.OnesMatrix(3, 2)
	if !mat.EqualApprox(&constant0, &ones, 1e-14) {
		t.Errorf(
			"expected pm.Constant() to be a matrix of all ones; received %v",
			constant0,
		)
	}
}

/*
TestPolynomialMatrix_String1
Description:

	Tests that the String() method properly returns a string
	that is composed of each of the polynomials in the matrix.
*/
func TestPolynomialMatrix_String1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1, p1},
		{p1, p1},
		{p1, p1},
	}

	// Test
	pm1String := pm1.String()

	// Verify that the string contains the string representation of each polynomial
	for _, pRow := range pm1 {
		for _, p := range pRow {
			if !strings.Contains(pm1String, p.String()) {
				t.Errorf(
					"expected pm1String to contain p.String(); received %v",
					pm1String,
				)
			}
		}

	}
}

/*
TestPolynomialMatrix_AsSimplifiedExpression1
Description:

	Tests that the AsSimplifiedExpression() method properly returns
	a PolynomialMatrix when a well-defined polynomial matrix
	calls it. The result should be a polynomial matrix with
	the same dimensions and number of monomials in each polynomial.
*/
func TestPolynomialMatrix_AsSimplifiedExpression1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	p1 := v1.ToPolynomial()
	var pm1 symbolic.PolynomialMatrix = [][]symbolic.Polynomial{
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
		{p1.Minus(3.14).(symbolic.Polynomial), p1.Minus(3.14).(symbolic.Polynomial)},
	}

	// Test
	pm2 := pm1.AsSimplifiedExpression()

	pm2AsPM, tf := pm2.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected pm2 to be a PolynomialMatrix; received %v",
			pm2,
		)
	}

	// Check that the dimensions are the same
	if pm1.Dims()[0] != pm2AsPM.Dims()[0] || pm1.Dims()[1] != pm2AsPM.Dims()[1] {
		t.Errorf(
			"expected pm2.Dims() to be %v; received %v",
			pm1.Dims(),
			pm2AsPM.Dims(),
		)
	}

	// Check that each polynomial in pm2 contains two monomials
	for _, pm2Row := range pm2AsPM {
		for _, p := range pm2Row {
			if len(p.Monomials) != 2 {
				t.Errorf(
					"expected len(p.Monomials) to be 2; received %v",
					len(p.Monomials),
				)
			}
		}
	}

}
