package symbolic_test

/*
scalar_expression_test.go
Description:
	Tests for the functions mentioned in the variable_vector.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestVariableVector_Len1
Description:

	Verifiest that a variable vector with 111 elements has the proper length.
*/
func TestVariableVector_Len1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	if vv.Len() != N {
		t.Errorf(
			"Expected vv.Len() to be %v; received %v",
			N,
			vv.Len(),
		)
	}
}

/*
TestVariableVector_AtVec1
Description:

	This test verifies that the AtVec function returns a Variable object.
*/
func TestVariableVector_AtVec1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	for ii := 0; ii < N; ii++ {
		if _, ok := vv.AtVec(ii).(symbolic.Variable); !ok {
			t.Errorf(
				"Expected vv.AtVec(%v) to be a Variable object; received %T",
				ii,
				vv.AtVec(ii),
			)
		}
	}
}

/*
TestVariableVector_Variables1
Description:

	Verifies that the variables function returns a slice of unique variables
	that has length equal to the original vector's length.
*/
func TestVariableVector_Variables1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	vars := vv.Variables()
	if len(vars) != N {
		t.Errorf(
			"Expected len(vars) to be %v; received %v",
			N,
			len(vars),
		)
	}

	// Check that all variables are unique
	for ii := 0; ii < N; ii++ {
		for jj := ii + 1; jj < N; jj++ {
			if vars[ii].ID == vars[jj].ID {
				t.Errorf(
					"Expected vars[%v].ID to be unique; received vars[%v].ID= %v",
					ii,
					jj,
					vars[ii].ID,
				)
			}
		}
	}
}
