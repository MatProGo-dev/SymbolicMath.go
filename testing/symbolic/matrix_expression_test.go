package symbolic_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
matrix_expression_test.go
Description:

	Tests the MatrixExpression object.
*/

/*
TestMatrixExpression_ToMatrixExpression1
Description:

	Tests that the conversion method fails if the input is not a valid MatrixExpression.
*/
func TestMatrixExpression_ToMatrixExpression1(t *testing.T) {
	// Constants
	x := 3

	// Test
	_, err := symbolic.ToMatrixExpression(x)
	if err == nil {
		t.Errorf("Expected error when converting %v to a MatrixExpression; received nil", x)
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the input interface is of type %T, which is not recognized as a MatrixExpression.",
				x,
			),
		) {
			t.Errorf("Expected error message to contain %v; received %v",
				fmt.Sprintf(
					"the input interface is of type %T, which is not recognized as a MatrixExpression.",
					x,
				),
				err.Error(),
			)
		}
	}
}

/*
TestMatrixExpression_ConcretizeMatrixExpression1
Description:

	Tests the conversion of a slice of slices of constants (K) to a KMatrix.
*/
func TestMatrixExpression_ConcretizeMatrixExpression1(t *testing.T) {
	// Setup
	x := [][]symbolic.ScalarExpression{
		{symbolic.K(1.0), symbolic.K(2), symbolic.K(3)},
		{symbolic.K(4), symbolic.K(5), symbolic.K(6)},
		{symbolic.K(7), symbolic.K(8), symbolic.K(9)},
	}

	// Test
	me := symbolic.ConcretizeMatrixExpression(x)

	// Check that me is a KMatrix
	if _, ok := me.(symbolic.KMatrix); !ok {
		t.Errorf("Expected a KMatrix; received %T", me)
	}

	// Check that me is the correct KMatrix (i.e., it has the correct dimensions)
	if me.Dims()[0] != 3 || me.Dims()[1] != 3 {
		t.Errorf("Expected a 3x3 KMatrix; received %dx%d KMatrix", me.Dims()[0], me.Dims()[1])
	}

}

/*
TestMatrixExpression_ConcretizeMatrixExpression2
Description:

	Tests the conversion of a slice of slices of Monomials to a MonomialMatrix.
*/
func TestMatrixExpression_ConcretizeMatrixExpression2(t *testing.T) {
	// Setup
	m := symbolic.NewVariable().ToMonomial()
	x := [][]symbolic.ScalarExpression{
		{m, m},
		{m, m},
	}

	// Test
	me := symbolic.ConcretizeMatrixExpression(x)

	// Check that me is a MonomialMatrix
	if _, ok := me.(symbolic.MonomialMatrix); !ok {
		t.Errorf("Expected a MonomialMatrix; received %T", me)
	}

	// Check that me is the correct MonomialMatrix (i.e., it has the correct dimensions)
	if me.Dims()[0] != 2 || me.Dims()[1] != 2 {
		t.Errorf("Expected a 2x2 MonomialMatrix; received %dx%d MonomialMatrix", me.Dims()[0], me.Dims()[1])
	}
}

/*
TestMatrixExpression_ConcretizeMatrixExpression3
Description:

	Tests the conversion of a slice of slices of Polynomials to a PolynomialMatrix.
*/
func TestMatrixExpression_ConcretizeMatrixExpression3(t *testing.T) {
	// Setup
	p := symbolic.NewVariable().ToPolynomial()
	x := [][]symbolic.ScalarExpression{
		{p, p},
		{p, p},
	}

	// Test
	me := symbolic.ConcretizeMatrixExpression(x)

	// Check that me is a PolynomialMatrix
	if _, ok := me.(symbolic.PolynomialMatrix); !ok {
		t.Errorf("Expected a PolynomialMatrix; received %T", me)
	}

	// Check that me is the correct PolynomialMatrix (i.e., it has the correct dimensions)
	if me.Dims()[0] != 2 || me.Dims()[1] != 2 {
		t.Errorf("Expected a 2x2 PolynomialMatrix; received %dx%d PolynomialMatrix", me.Dims()[0], me.Dims()[1])
	}
}

/*
TestMatrixExpression_ConcretizeMatrixExpression4
Description:

	Tests the conversion of a slice of slices of Variable objects to a VariableMatrix.
*/
func TestMatrixExpression_ConcretizeMatrixExpression4(t *testing.T) {
	// Setup
	v := symbolic.NewVariable()
	x := [][]symbolic.ScalarExpression{
		{v, v},
		{v, v},
		{v, v},
	}

	// Test
	me := symbolic.ConcretizeMatrixExpression(x)

	// Check that me is a VariableMatrix
	if _, ok := me.(symbolic.VariableMatrix); !ok {
		t.Errorf("Expected a VariableMatrix; received %T", me)
	}

	// Check that me is the correct VariableMatrix (i.e., it has the correct dimensions)
	if me.Dims()[0] != 3 || me.Dims()[1] != 2 {
		t.Errorf("Expected a 3x2 VariableMatrix; received %dx%d VariableMatrix", me.Dims()[0], me.Dims()[1])
	}

}

/*
TestMatrixExpression_ConcretizeMatrixExpression5
Description:

	Tests the conversion of a slice of slices constaining constants (K) and variables (symbolic.variable) expressions
	to a MonomialMatrix.
*/
func TestMatrixExpression_ConcretizeMatrixExpression5(t *testing.T) {
	// Setup
	k := symbolic.K(2)
	v := symbolic.NewVariable()
	x := [][]symbolic.ScalarExpression{
		{k, v},
		{v, k},
	}

	// Test
	me := symbolic.ConcretizeMatrixExpression(x)

	// Check that me is a MonomialMatrix
	if _, ok := me.(symbolic.MonomialMatrix); !ok {
		t.Errorf("Expected a MonomialMatrix; received %T", me)
	}

	// Check that me is the correct MonomialMatrix (i.e., it has the correct dimensions)
	if me.Dims()[0] != 2 || me.Dims()[1] != 2 {
		t.Errorf("Expected a 2x2 MonomialMatrix; received %dx%d MonomialMatrix", me.Dims()[0], me.Dims()[1])
	}
}

/*
TestMatrixExpression_MatrixPowerTemplate1
Description:

	Tests that the matrix power template properly panics when called with a MatrixExpression that is not
	well-defined (in this case, a MonomialMatrix).
*/
func TestMatrixExpression_MatrixPowerTemplate1(t *testing.T) {
	// Setup
	m := symbolic.Monomial{
		Coefficient:     1.2,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
		Exponents:       []int{1},
	}
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixPowerTemplate on a MonomialMatrix; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), m.Check().Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", m.Check().Error(), rAsE.Error())
		}
	}()
	symbolic.MatrixPowerTemplate(x, 2)
}

/*
TestMatrixExpression_MatrixPowerTemplate2
Description:

	Tests that the matrix power template properly panics when called with a MatrixExpression that is well-defined
	but is not square.
*/
func TestMatrixExpression_MatrixPowerTemplate2(t *testing.T) {
	// Setup
	m := symbolic.NewVariable().ToMonomial()
	x := symbolic.MonomialMatrix{
		{m, m},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixPowerTemplate on a non-square MonomialMatrix; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), "matrix is not square") {
			t.Errorf("Expected the panic to contain the error message %v; received %v", "Matrix is not square", rAsE.Error())
		}
	}()
	symbolic.MatrixPowerTemplate(x, 2)
}

/*
TestMatrixExpression_MatrixPowerTemplate3
Description:

	Tests that the matrix power template properly panics when called with a MatrixExpression that is well-defined
	and is square but with a negative exponent.
*/
func TestMatrixExpression_MatrixPowerTemplate3(t *testing.T) {
	// Setup
	m := symbolic.NewVariable().ToMonomial()
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixPowerTemplate with a negative exponent; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		expectedError := smErrors.NegativeExponentError{
			Exponent: -2,
		}
		if !strings.Contains(rAsE.Error(), expectedError.Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", "Negative exponent", rAsE.Error())
		}
	}()
	symbolic.MatrixPowerTemplate(x, -2)
}

/*
TestMatrixExpression_MatrixSubstituteTemplate1
Description:

	Tests that the matrix substitute template properly panics when called with a MatrixExpression that is not
	well-defined (in this case, a MonomialMatrix).
*/
func TestMatrixExpression_MatrixSubstituteTemplate1(t *testing.T) {
	// Setup
	m := symbolic.Monomial{
		Coefficient:     1.2,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
		Exponents:       []int{1},
	}
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixSubstituteTemplate on a MonomialMatrix; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), m.Check().Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", m.Check().Error(), rAsE.Error())
		}
	}()
	symbolic.MatrixSubstituteTemplate(x, symbolic.NewVariable(), symbolic.NewVariable())
}

/*
TestMatrixExpression_MatrixSubstituteTemplate2
Description:

	Tests that the matrix substitute template properly panics when called with a MatrixExpression that is well-defined
	and a variable vIn that is not well-defined.
*/
func TestMatrixExpression_MatrixSubstituteTemplate2(t *testing.T) {
	// Setup
	m := symbolic.NewVariable().ToMonomial()
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}
	v1 := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixSubstituteTemplate with an invalid variable; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), v1.Check().Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", "the input variable is not well-defined", rAsE.Error())
		}
	}()
	symbolic.MatrixSubstituteTemplate(x, v1, symbolic.NewVariable())
}

/*
TestMatrixExpression_MatrixSubstituteTemplate3
Description:

	Tests that the matrix substitute template properly panics when called with a MatrixExpression that is well-defined,
	a variable vIn that is well-defined, but an expression eIn that is not well-defined (in this case a monomial).
*/
func TestMatrixExpression_MatrixSubstituteTemplate3(t *testing.T) {
	// Setup
	m := symbolic.NewVariable().ToMonomial()
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.2,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixSubstituteTemplate with an invalid expression; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), m1.Check().Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", m1.Check().Error(), rAsE.Error())
		}
	}()
	symbolic.MatrixSubstituteTemplate(x, v1, m1)
}

/*
TestMatrixExpression_MatrixMultiplyTemplate1
Description:

	Tests that the matrix multiply template properly panics when called with a MatrixExpression that is not
	well-defined (in this case, a MonomialMatrix).
*/
func TestMatrixExpression_MatrixMultiplyTemplate1(t *testing.T) {
	// Setup
	m := symbolic.Monomial{
		Coefficient:     1.2,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
		Exponents:       []int{1},
	}
	x := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}
	y := symbolic.MonomialMatrix{
		{m, m},
		{m, m},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixMultiplyTemplate on a MonomialMatrix; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), m.Check().Error()) {
			t.Errorf("Expected the panic to contain the error message %v; received %v", m.Check().Error(), rAsE.Error())
		}
	}()
	symbolic.MatrixMultiplyTemplate(x, y)
}

/*
TestMatrixExpression_MatrixMultiplyTemplate2
Description:

	Tests that the matrix multiply template properly panics when called with a MatrixExpression that is well-defined
	but with incompatible dimensions.
	In this case, we use two KMatrix objects with incompatible dimensions.
	The first KMatrix is 2x3 and the second KMatrix is 2x2.
*/
func TestMatrixExpression_MatrixMultiplyTemplate2(t *testing.T) {
	// Setup
	x := symbolic.KMatrix{
		{symbolic.K(1), symbolic.K(2), symbolic.K(3)},
		{symbolic.K(4), symbolic.K(5), symbolic.K(6)},
	}
	y := symbolic.KMatrix{
		{symbolic.K(7), symbolic.K(8)},
		{symbolic.K(9), symbolic.K(10)},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected a panic when calling MatrixMultiplyTemplate with incompatible dimensions; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("Expected the panic to be an error; received %T", r)
		}

		if !strings.Contains(rAsE.Error(), "dimension error") {
			t.Errorf("Expected the panic to contain the error message %v; received %v", "incompatible dimensions", rAsE.Error())
		}
	}()
	symbolic.MatrixMultiplyTemplate(x, y)
}
