package symbolic_test

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
polynomial_like_matrix_test.go
Description:

	Tests for the functions mentioned in the polynomial_like_matrix.go file.
*/

/*
TestPolynomialLikeMatrix_IsPolynomialLikeMatrix1
Description:

	Tests that a mat.Dense object is identified as a polynomial like matrix object.
*/
func TestPolynomialLikeMatrix_IsPolynomialLikeMatrix1(t *testing.T) {
	// Constants
	x := symbolic.Identity(10)

	// Test
	if !symbolic.IsPolynomialLikeMatrix(x) {
		t.Errorf(
			"Expected IsPolynomialLikeMatrix(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeMatrix_IsPolynomialLikeMatrix2
Description:

	Tests that a VariableMatrix object is identified as a polynomial like matrix object.
*/
func TestPolynomialLikeMatrix_IsPolynomialLikeMatrix2(t *testing.T) {
	// Constants
	x := symbolic.NewVariableMatrix(3, 3)

	// Test
	if !symbolic.IsPolynomialLikeMatrix(x) {
		t.Errorf(
			"Expected IsPolynomialLikeMatrix(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeMatrix_IsPolynomialLikeMatrix3
Description:

	Tests that a MonomialMatrix object is identified as a polynomial like matrix object.
*/
func TestPolynomialLikeMatrix_IsPolynomialLikeMatrix3(t *testing.T) {
	// Constants
	x := getKMatrix.From(symbolic.Identity(3)).ToMonomialMatrix()

	// Test
	if !symbolic.IsPolynomialLikeMatrix(x) {
		t.Errorf(
			"Expected IsPolynomialLikeMatrix(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeMatrix_IsPolynomialLikeMatrix4
Description:

	Tests that a PolynomialMatrix object is identified as a polynomial like matrix object.
*/
func TestPolynomialLikeMatrix_IsPolynomialLikeMatrix4(t *testing.T) {
	// Constants
	x := symbolic.NewVariableMatrix(3, 3).ToPolynomialMatrix()

	// Test
	if !symbolic.IsPolynomialLikeMatrix(x) {
		t.Errorf(
			"Expected IsPolynomialLikeMatrix(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeMatrix_ToPolynomialLikeMatrix1
Description:

	Tests to make sure that a symbolic.Variable object creates
	an error when attempting to convert to a PolynomialLikeMatrix.
*/
func TestPolynomialLikeMatrix_ToPolynomialLikeMatrix1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if _, err := symbolic.ToPolynomialLikeMatrix(x); err == nil {
		t.Errorf(
			"Expected ToPolynomialLikeMatrix(%T) to return an error; received nil",
			x,
		)
	}
}

/*
TestPolynomialLikeMatrix_ToPolynomialLikeMatrix2
Description:

	Tests to make sure that a mat.Dense object is properly
	converted to a PolynomialLikeMatrix object.
	Verifies that the output has no error and is of the type KMatrix.
*/
func TestPolynomialLikeMatrix_ToPolynomialLikeMatrix2(t *testing.T) {
	// Constants
	x := symbolic.Identity(3)

	// Test
	xAsPLM, err := symbolic.ToPolynomialLikeMatrix(x)
	if err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeMatrix(%T) to return no error; received %v",
			x,
			err,
		)

	}

	// Check that xAsPLM is a KMatrix object
	if _, ok := xAsPLM.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected the output to be a KMatrix; received %T",
			xAsPLM,
		)
	}
}

/*
TestPolynomialLikeMatrix_ToPolynomialLikeMatrix3
Description:

	Tests to make sure that a VariableMatrix object is properly
	converted to a PolynomialLikeMatrix object.
*/
func TestPolynomialLikeMatrix_ToPolynomialLikeMatrix3(t *testing.T) {
	// Constants
	x := symbolic.NewVariableMatrix(3, 3)

	// Test
	_, err := symbolic.ToPolynomialLikeMatrix(x)
	if err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeMatrix(%T) to return no error; received %v",
			x,
			err,
		)
	}
}

/*
TestPolynomialLikeMatrix_ToPolynomialLikeMatrix4
Description:

	Tests to make sure that a MonomialMatrix object is properly
	converted to a PolynomialLikeMatrix object.
*/
func TestPolynomialLikeMatrix_ToPolynomialLikeMatrix4(t *testing.T) {
	// Constants
	x := getKMatrix.From(symbolic.Identity(3)).ToMonomialMatrix()

	// Test
	_, err := symbolic.ToPolynomialLikeMatrix(x)
	if err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeMatrix(%T) to return no error; received %v",
			x,
			err,
		)
	}
}

/*
TestPolynomialLikeMatrix_ToPolynomialLikeMatrix5
Description:

	Tests to make sure that a PolynomialMatrix object is properly
	converted to a PolynomialLikeMatrix object.
*/
func TestPolynomialLikeMatrix_ToPolynomialLikeMatrix5(t *testing.T) {
	// Constants
	x := symbolic.NewVariableMatrix(3, 3).ToPolynomialMatrix()

	// Test
	_, err := symbolic.ToPolynomialLikeMatrix(x)
	if err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeMatrix(%T) to return no error; received %v",
			x,
			err,
		)
	}
}
