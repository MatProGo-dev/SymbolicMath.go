package symbolic_test

/*
constant_matrix_test.go
Description:
	Tests for the functions mentioned in the constant_matrix.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestConstantMatrix_Variables1
Description:

	Tests that the Variables() method returns NO variables when called.
*/
func TestConstantMatrix_Variables1(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.KMatrix(eye1)

	// Test that variables returns empty list
	if len(km1.Variables()) != 0 {
		t.Errorf(
			"Expected no variables to be in constant matrix; received %v",
			len(km1.Variables()),
		)
	}
}

/*
TestConstantMatrix_Plus1
Description:

	Tests that the Plus() method properly panics when a matrix of improper
	dimensions is added to the constant matrix.
*/
func TestConstantMatrix_Plus1(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.KMatrix(eye1)

	eye2 := symbolic.Identity(4)

	// Define Test Handler
	defer func() {
		recoveredVal := recover()
		if recoveredVal == nil {
			t.Errorf("Expected constant matrix addition to panic when dimensions are not equal; did not panic")
		}

		err, ok := recoveredVal.(error)
		if !ok {
			t.Errorf("Expected recovered value to be an error; received %T", recoveredVal)
		}

		expectedError := symbolic.DimensionError{
			Operation: "Plus",
			Arg1:      km1,
			Arg2:      symbolic.KMatrix(eye2),
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	_, err := km1.Plus(eye2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	t.Errorf("TestConstantMatrix_Plus1 did not panic as expected")

}
