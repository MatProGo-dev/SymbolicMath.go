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
