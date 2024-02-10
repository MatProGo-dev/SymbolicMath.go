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

/*
TestKMatrix_Multiply5
Description:

	Tests that the Multiply() method panics when a well-defined
	KMatrix is multiplied by an unsupported input (in this case, a string).
*/
func TestKMatrix_Multiply5(t *testing.T) {
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

		// Check that error is the UnsupportedInputError error defined in KMatrix.Multiply()
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Multiply",
			Input:        "test",
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.Multiply("test")

	t.Errorf("TestKMatrix_Multiply5 did not panic as expected")
}

/*
TestKMatrix_Multiply6
Description:

	Tests that the Multiply() method properly multiplies a KMatrix by a
	KMatrix with the appropriate dimensions.
	The first KMatrix is a 3x3 identity matrix, and the second is a 3x2 ones matrix.
*/
func TestKMatrix_Multiply6(t *testing.T) {
	// Constants
	eye1 := symbolic.Identity(3)
	km1 := symbolic.DenseToKMatrix(eye1)

	ones2 := symbolic.OnesMatrix(3, 2)
	km2 := symbolic.DenseToKMatrix(ones2)

	// Test
	km3 := km1.Multiply(km2)

	// Verify that the result is a constant matrix
	if _, ok := km3.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected km3 to be a symbolic.KMatrix; received %T",
			km3,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, nC := km3.Dims()[0], km3.Dims()[1]
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			km3_ii_jj := km3.(symbolic.KMatrix).At(rowIndex, colIndex)
			elt := km3_ii_jj.(symbolic.K)
			if float64(elt) != 1 {
				t.Errorf(
					"Expected km3.At(%v,%v) to be 1; received %v",
					rowIndex,
					colIndex,
					km3.(symbolic.KMatrix).At(rowIndex, colIndex),
				)
			}
		}

	}
}

/*
TestKMatrix_Multiply7
Description:

	Tests that the Multiply() method properly multiplies a KMatrix by
	a mat.Dense of the appropriate dimensions. When the KMatrix
	has dimension (3x2) and the mat.Dense has dimension (2x1),
	the result should be a 3x1 KVector.
*/
func TestKMatrix_Multiply7(t *testing.T) {
	// Constants
	ones1 := symbolic.OnesMatrix(3, 2)
	km1 := symbolic.DenseToKMatrix(ones1)

	ones2 := symbolic.OnesMatrix(2, 1)

	// Test
	kv3 := km1.Multiply(ones2)

	// Verify that the result is a constant matrix
	if _, ok := kv3.(symbolic.KVector); !ok {
		t.Errorf(
			"Expected kv3 to be a symbolic.KVector; received %T",
			kv3,
		)
	}

	// Verify that the elements of the result is the correct value
	nR, _ := ones1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		kv3_ii := kv3.(symbolic.KVector).AtVec(rowIndex)
		elt := kv3_ii.(symbolic.K)
		if float64(elt) != 2 {
			t.Errorf(
				"Expected kv3.At(%v) to be 2; received %v",
				rowIndex,
				kv3.(symbolic.KVector).AtVec(rowIndex),
			)
		}
	}
}

/*
TestKMatrix_Multiply8
Description:

	Tests that the Multiply() method properly multiplies
	a KMatrix (of dimension (1 x 3)) by a KVector (of dimension 3).
	The result should be a single scalar K.
*/
func TestKMatrix_Multiply8(t *testing.T) {
	// Constants
	ones1 := symbolic.OnesMatrix(1, 3)
	km1 := symbolic.DenseToKMatrix(ones1)

	ones2 := symbolic.OnesMatrix(3, 1)

	// Test
	k3 := km1.Multiply(ones2)

	// Verify that the result is a constant matrix
	if _, ok := k3.(symbolic.K); !ok {
		t.Errorf(
			"Expected k3 to be a symbolic.K; received %T",
			k3,
		)
	}

	// Verify that the elements of the result is the correct value
	if float64(k3.(symbolic.K)) != 3 {
		t.Errorf(
			"Expected k3 to be 3; received %v",
			k3,
		)
	}
}

/*
TestKMatrix_Transpose1
Description:

	Tests that the Transpose() method properly transposes a KMatrix.
*/
func TestKMatrix_Transpose1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Transpose
	km2 := km1.Transpose()

	// Verify that the dimensions have swapped
	if (km2.Dims()[0] != km1.Dims()[1]) || (km2.Dims()[1] != km1.Dims()[0]) {
		t.Errorf(
			"Expected km2 to have dimensions (%vx%v); received %v",
			km1.Dims()[1],
			km1.Dims()[0],
			km2.Dims(),
		)
	}

	// Verify that the elements have swapped correctly
	kmT, _ := km2.(symbolic.KMatrix)
	for rowIndex := 0; rowIndex < kmT.Dims()[0]; rowIndex++ {
		for colIndex := 0; colIndex < kmT.Dims()[1]; colIndex++ {
			if kmT.At(rowIndex, colIndex) != km1.At(colIndex, rowIndex) {
				t.Errorf(
					"Expected km2.At(%v,%v) to be %v; received %v",
					rowIndex,
					colIndex,
					km1.At(colIndex, rowIndex),
					kmT.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestKMatrix_LessEq1
Description:

	Tests that the LessEq() method properly compares two KMatrices.
	The result should be a matrix constraint with the correct dimensions.
*/
func TestKMatrix_LessEq1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	km2 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Test
	mc := km1.LessEq(km2)

	// Verify that the result is a matrix constraint
	if _, ok := mc.(symbolic.MatrixConstraint); !ok {
		t.Errorf(
			"Expected mc to be a symbolic.MatrixConstraint; received %T",
			mc,
		)
	}

	// Verify that the dimensions of the constraint are the same
	// as that of km1
	mcM, _ := mc.(symbolic.MatrixConstraint)
	if (mcM.Dims()[0] != km1.Dims()[0]) || (mcM.Dims()[1] != km1.Dims()[1]) {
		t.Errorf(
			"Expected mc to have dimensions (%vx%v); received %v",
			km1.Dims()[0],
			km1.Dims()[1],
			mcM.Dims(),
		)
	}
}

/*
TestKMatrix_LessEq2
Description:

	Tests that LessEq() properly panics when the dimensions of the KMatrix receiver
	and the variablematrix are not equal. In this case the KMatrix has dimension
	(2x3) and the variable matrix has dimension (3x2).
*/
func TestKMatrix_LessEq2(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Define Test Handler
	defer func() {
		recoveredVal := recover()
		if recoveredVal == nil {
			t.Errorf("Expected LessEq to panic when dimensions are not equal; did not panic")
		}

		err, ok := recoveredVal.(error)
		if !ok {
			t.Errorf("Expected recovered value to be an error; received %T", recoveredVal)
		}

		var expectedSense symbolic.ConstrSense = symbolic.SenseLessThanEqual
		err2 := smErrors.DimensionError{
			Operation: "Comparison (" + expectedSense.String() + ")",
			Arg1:      km1,
			Arg2:      symbolic.NewVariableMatrix(3, 2),
		}
		if err.Error() != err2.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.LessEq(symbolic.NewVariableMatrix(3, 2))

	t.Errorf("TestKMatrix_LessEq2 did not panic as expected")
}

/*
TestKMatrix_GreaterEq1
Description:

	Tests that the GreaterEq() method properly compares a KMatrix
	to a variable matrix of appropriate dimension.
	This function checks that the result is a matrix constraint
	with the appropriate dimensions.
*/
func TestKMatrix_GreaterEq1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	km2 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Test
	mc := km1.GreaterEq(km2)

	// Verify that the result is a matrix constraint
	if _, ok := mc.(symbolic.MatrixConstraint); !ok {
		t.Errorf(
			"Expected mc to be a symbolic.MatrixConstraint; received %T",
			mc,
		)
	}

	// Verify that the dimensions of the constraint are the same
	// as that of km1
	mcM, _ := mc.(symbolic.MatrixConstraint)
	if (mcM.Dims()[0] != km1.Dims()[0]) || (mcM.Dims()[1] != km1.Dims()[1]) {
		t.Errorf(
			"Expected mc to have dimensions (%vx%v); received %v",
			km1.Dims()[0],
			km1.Dims()[1],
			mcM.Dims(),
		)
	}
}

/*
TestKMatrix_GreaterEq2
Description:

	Tests that GreaterEq() properly panics when the KMatrix
	is compared to an unexpected input (in this case, a string).
*/
func TestKMatrix_GreaterEq2(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Define Test Handler
	defer func() {
		recoveredVal := recover()
		if recoveredVal == nil {
			t.Errorf("Expected GreaterEq to panic when dimensions are not equal; did not panic")
		}

		err, ok := recoveredVal.(error)
		if !ok {
			t.Errorf("Expected recovered value to be an error; received %T", recoveredVal)
		}

		var expectedSense symbolic.ConstrSense = symbolic.SenseGreaterThanEqual
		err2 := smErrors.UnsupportedInputError{
			FunctionName: "KMatrix.Comparison (" + expectedSense.String() + ")",
			Input:        "test",
		}
		if err.Error() != err2.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Test
	km1.GreaterEq("test")

	t.Errorf("TestKMatrix_GreaterEq2 did not panic as expected")
}
