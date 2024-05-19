package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
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
