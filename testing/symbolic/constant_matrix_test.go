package symbolic_test

/*
constant_matrix_test.go
Description:
	Tests for the functions mentioned in the constant_matrix.go file.
*/

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
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
	km1 := symbolic.DenseToKMatrix(eye1)

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
	km1 := symbolic.DenseToKMatrix(eye1)

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

		expectedError := smErrors.DimensionError{
			Operation: "Plus",
			Arg1:      km1,
			Arg2:      symbolic.DenseToKMatrix(eye2),
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Plus(eye2)

	t.Errorf("TestConstantMatrix_Plus1 did not panic as expected")

}

/*
TestConstantMatrix_Plus2
Description:

	Tests that the Plus() method properly adds a float64 to the constant matrix.
*/
func TestConstantMatrix_Plus2(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	// Test
	km2 := km1.Plus(3.14)

	// Verify that the result is a constant matrix
	if _, ok := km2.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected km2 to be a symbolic.KMatrix; received %T",
			km2,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			km2_ii_jj := km2.(symbolic.KMatrix).At(rowIndex, colIndex)
			elt := km2_ii_jj.(symbolic.K)
			if float64(elt) != eye1.At(rowIndex, colIndex)+3.14 {
				t.Errorf(
					"Expected km2.At(0,0) to be 4.14; received %v",
					km2.(symbolic.KMatrix).At(0, 0),
				)
			}
		}

	}

}

/*
TestConstantMatrix_Plus3
Description:

	Tests that the Plus() method properly adds a symbolic.K to the constant matrix.
*/
func TestConstantMatrix_Plus3(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	// Test
	km2 := km1.Plus(symbolic.K(3.14))

	// Verify that the result is a constant matrix
	if _, ok := km2.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected km2 to be a symbolic.KMatrix; received %T",
			km2,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			km2_ii_jj := km2.(symbolic.KMatrix).At(rowIndex, colIndex)
			elt := km2_ii_jj.(symbolic.K)
			if float64(elt) != eye1.At(rowIndex, colIndex)+3.14 {
				t.Errorf(
					"Expected km2.At(0,0) to be 4.14; received %v",
					km2.(symbolic.KMatrix).At(0, 0),
				)
			}
		}

	}

}

/*
TestConstantMatrix_Plus4
Description:

	Tests that the Plus() method properly panics
	when called with a valid KMatrix and a not well-defined
	expression (in this case, a Monomial).
*/
func TestConstantMatrix_Plus4(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	monom2 := symbolic.Monomial{
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
	}

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

		err2 := monom2.Check()
		if err.Error() != err2.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Plus(monom2)

	t.Errorf("TestConstantMatrix_Plus4 did not panic as expected")
}

/*
TestConstantMatrix_Plus5
Description:

	Tests that the Plus() method properly adds a KMatrix to
	a polynomial matrix. The result should be a polynomial matrix.
*/
func TestConstantMatrix_Plus5(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	ones2 := symbolic.OnesMatrix(3, 3)
	pm2 := getKMatrix.From(ones2).ToPolynomialMatrix()

	// Test
	pm3 := km1.Plus(pm2)

	// Verify that the result is a polynomial matrix
	if _, ok := pm3.(symbolic.PolynomialMatrix); !ok {
		t.Errorf(
			"Expected pm3 to be a symbolic.PolynomialMatrix; received %T",
			pm3,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			pm3_ii_jj := pm3.(symbolic.PolynomialMatrix).At(rowIndex, colIndex)
			elt := pm3_ii_jj.(symbolic.Polynomial)
			if len(elt.Monomials) != 1 {
				t.Errorf(
					"Expected pm3.At(0,0) to be a degree 0 polynomial; received %v",
					pm3.(symbolic.PolynomialMatrix).At(0, 0),
				)
			}
		}

	}
}

/*
TestConstantMatrix_Plus6
Description:

	Tests that the Plus() method properly adds a KMatrix to
	a variable. The result should be a matrix of polynomials.
*/
func TestConstantMatrix_Plus6(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	x2 := symbolic.NewVariable()

	// Test
	pm3 := km1.Plus(x2)

	// Verify that the result is a polynomial matrix
	if _, ok := pm3.(symbolic.PolynomialMatrix); !ok {
		t.Errorf(
			"Expected pm3 to be a symbolic.PolynomialMatrix; received %T",
			pm3,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			pm3_ii_jj := pm3.(symbolic.PolynomialMatrix).At(rowIndex, colIndex)
			elt := pm3_ii_jj.(symbolic.Polynomial)
			if len(elt.Monomials) != 2 {
				t.Errorf(
					"Expected pm3.At(0,0) to be a polynomial with 2 monomials; received %v",
					pm3.(symbolic.PolynomialMatrix).At(0, 0),
				)
			}
		}

	}
}

/*
TestKMatrix_Plus7
Description:

	Tests that the Plus() method properly panics
	when the a well-defined KMatrix is added to
	an unsupported input (in this case, a string).
*/
func TestKMatrix_Plus7(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

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

		// Check that error is the UnsupportedInputError error defined in KMatrix.Plus()
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Plus",
			Input:        "test",
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Plus("test")

	t.Errorf("TestKMatrix_Plus7 did not panic as expected")
}

/*
TestKMatrix_Multiply1
Description:

	Tests that the Multiply() method properly panics when the KMatrix is added
	to an expression that is improperly intiialized.
*/
func TestKMatrix_Multiply1(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	monom2 := symbolic.Monomial{
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
	}

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

		err2 := monom2.Check()
		if err.Error() != err2.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Multiply(monom2)

	t.Errorf("TestKMatrix_Multiply1 did not panic as expected")
}

/*
TestKMatrix_Multiply2
Description:

	Tests that the Multiply() method properly panics when the dimensions of km
	and the expression are not compatible.
*/
func TestKMatrix_Multiply2(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	ones2 := symbolic.OnesMatrix(2, 4)
	km2 := symbolic.DenseToKMatrix(ones2)

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

		err2 := smErrors.DimensionError{
			Operation: "Multiply",
			Arg1:      km1,
			Arg2:      km2,
		}
		if err.Error() != err2.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Multiply(km2)

	t.Errorf("TestKMatrix_Multiply2 did not panic as expected")
}

/*
TestKMatrix_Multiply3
Description:

	Tests that the Multiply() method properly multiplies a KMatrix by a float64.
*/
func TestKMatrix_Multiply3(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	// Test
	km2 := km1.Multiply(3.14)

	// Verify that the result is a constant matrix
	if _, ok := km2.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected km2 to be a symbolic.KMatrix; received %T",
			km2,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			km2_ii_jj := km2.(symbolic.KMatrix).At(rowIndex, colIndex)
			elt := km2_ii_jj.(symbolic.K)
			if float64(elt) != eye1.At(rowIndex, colIndex)*3.14 {
				t.Errorf(
					"Expected km2.At(0,0) to be 3.14; received %v",
					km2.(symbolic.KMatrix).At(0, 0),
				)
			}
		}

	}
}

/*
TestKMatrix_Multiply4
Description:

	Tests that the Multiply() method properly multiplies a KMatrix by a symbolic.K.
*/
func TestKMatrix_Multiply4(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	// Test
	km2 := km1.Multiply(symbolic.K(3.14))

	// Verify that the result is a constant matrix
	if _, ok := km2.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected km2 to be a symbolic.KMatrix; received %T",
			km2,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := eye1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			km2_ii_jj := km2.(symbolic.KMatrix).At(rowIndex, colIndex)
			elt := km2_ii_jj.(symbolic.K)
			if float64(elt) != eye1.At(rowIndex, colIndex)*3.14 {
				t.Errorf(
					"Expected km2.At(0,0) to be 3.14; received %v",
					km2.(symbolic.KMatrix).At(0, 0),
				)
			}
		}

	}
}
