package symbolic_test

/*
constant_matrix_test.go
Description:
	Tests for the functions mentioned in the constant_matrix.go file.
*/

import (
	"fmt"
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"math"
	"reflect"
	"strings"
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
	a KMatrix (of dimension (1 x 3)) by a KMatrix (of dimension 3).
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
TestKMatrix_Multiply9
Description:

	Tests that the Multiply() method properly
	computes the multiplication of a KMatrix by a VariableVector.
	In this case, KMatrix has dimension (2x3) and the KVector has dimension (3x1).
	The result should be a polynomial vector.
*/
func TestKMatrix_Multiply9(t *testing.T) {
	// Constants
	N := 3
	ones1 := symbolic.OnesMatrix(2, N)
	km1 := symbolic.DenseToKMatrix(ones1)

	vv2 := symbolic.NewVariableVector(N)

	// Test
	pv3 := km1.Multiply(vv2)

	// Verify that the result is a polynomial vector
	if _, ok := pv3.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected pv3 to be a symbolic.PolynomialVector; received %T",
			pv3,
		)
	}

	// Verify that the elements of the result is the correct value
	// - Number of monomials in each element of result should
	//	 contain 2 monomials
	nR, _ := ones1.Dims()
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		pv3_ii := pv3.(symbolic.PolynomialVector).AtVec(rowIndex)
		elt := pv3_ii.(symbolic.Polynomial)
		if len(elt.Monomials) != N {
			t.Errorf(
				"Expected pv3.At(%v) to be a degree %v polynomial; received %v",
				rowIndex,
				N,
				pv3.(symbolic.PolynomialVector).AtVec(rowIndex),
			)
		}
	}
}

/*
TestKMatrix_Multiply10
Description:

	Tests that the Multiply() method properly
	computes the multiplication of a KMatrix by a VariableVector.
	In this case, KMatrix has dimension (1x3) and the KVector has dimension (3x1).
	The result should be a scalar polynomial.
*/
func TestKMatrix_Multiply10(t *testing.T) {
	// Constants
	N := 3
	ones1 := symbolic.OnesMatrix(1, N)
	km1 := symbolic.DenseToKMatrix(ones1)

	vv2 := symbolic.NewVariableVector(N)

	// Test
	p3 := km1.Multiply(vv2)

	// Verify that the result is a polynomial
	if _, ok := p3.(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected p3 to be a symbolic.Polynomial; received %T",
			p3,
		)
	}

	// Verify that the elements of the result is the correct value
	// - Number of monomials in each element of result should
	//	 contain 1 monomial
	if len(p3.(symbolic.Polynomial).Monomials) != N {
		t.Errorf(
			"Expected p3 to be a degree %v polynomial; received %v",
			N,
			p3,
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

/*
TestKMatrix_Eq1
Description:

	Tests that the Eq() method properly panics when a well-formed
	KMatrix is compared with a not well-formed expression (in this case, a Monomial).
*/
func TestKMatrix_Eq1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	monom2 := symbolic.Monomial{
		VariableFactors: []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Define Test Handler
	defer func() {
		recoveredVal := recover()
		if recoveredVal == nil {
			t.Errorf("Expected Eq to panic when dimensions are not equal; did not panic")
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
	km1.Eq(monom2)

	t.Errorf("TestKMatrix_Eq1 did not panic as expected")
}

/*
TestKMatrix_Eq2
Description:

	Tests that the Eq() method creates the correct constraint
	when a KMatrix is compared to a Variable matrix of matching
	dimensions.
	The resulting constraint should be of the matrix constraint
	type, it should have the same dimensions as the KMatrix,
	and the sense should be equal.
*/
func TestKMatrix_Eq2(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	vm2 := symbolic.NewVariableMatrix(2, 3)

	// Test
	c3 := km1.Eq(vm2)

	// Verify that c3 is a matrix constraint
	if _, ok := c3.(symbolic.MatrixConstraint); !ok {
		t.Errorf(
			"Expected c3 to be a symbolic.MatrixConstraint; received %T",
			c3,
		)
	}

	// Verify that the dimensions of c3 are the same as that of km1
	c3M, _ := c3.(symbolic.MatrixConstraint)
	if (c3M.Dims()[0] != km1.Dims()[0]) || (c3M.Dims()[1] != km1.Dims()[1]) {
		t.Errorf(
			"Expected c3 to have dimensions (%vx%v); received %v",
			km1.Dims()[0],
			km1.Dims()[1],
			c3M.Dims(),
		)
	}

	// Verify that the sense of c3 is equal
	if c3M.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"Expected c3 to have sense equal; received %v",
			c3M.ConstrSense(),
		)
	}
}

/*
TestKMatrix_Comparison1
Description:

	Tests that the Comparison() method properly returns a constraint
	of the correct type when a KMatrix is compared to a scalar constant
	K with sense SenseGreaterThanEqual.
	The result should be a matrix constraint with the same dimensions
	as the KMatrix and the sense should be SenseGreaterThanEqual.
*/
func TestKMatrix_Comparison1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	k2 := symbolic.K(3)

	// Test
	c3 := km1.Comparison(k2, symbolic.SenseGreaterThanEqual)

	// Verify that c3 is a matrix constraint
	if _, ok := c3.(symbolic.MatrixConstraint); !ok {
		t.Errorf(
			"Expected c3 to be a symbolic.MatrixConstraint; received %T",
			c3,
		)
	}

	// Verify that the dimensions of c3 are the same as that of km1
	c3M, _ := c3.(symbolic.MatrixConstraint)
	if (c3M.Dims()[0] != km1.Dims()[0]) || (c3M.Dims()[1] != km1.Dims()[1]) {
		t.Errorf(
			"Expected c3 to have dimensions (%vx%v); received %v",
			km1.Dims()[0],
			km1.Dims()[1],
			c3M.Dims(),
		)
	}

	// Verify that the sense of c3 is equal
	if c3M.ConstrSense() != symbolic.SenseGreaterThanEqual {
		t.Errorf(
			"Expected c3 to have sense equal; received %v",
			c3M.ConstrSense(),
		)
	}

	// Verify that the right hand side of the constraint is a matrix
	// of all k2
	for rowIndex := 0; rowIndex < c3M.Dims()[0]; rowIndex++ {
		for colIndex := 0; colIndex < c3M.Dims()[1]; colIndex++ {
			if c3M.RightHandSide.At(rowIndex, colIndex) != k2 {
				t.Errorf(
					"Expected c3.At(%v,%v) to be %v; received %v",
					rowIndex,
					colIndex,
					k2,
					c3M.RightHandSide.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestKMatrix_Constant1
Description:

	Tests that the Constant() method properly returns a constant matrix
	when called on a KMatrix.
	The result should contain the same elements as if we called ToDense()
	on the KMatrix.
*/
func TestKMatrix_Constant1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Test
	km2 := km1.Constant()

	// Verify that the elements of the result is the correct value
	nR, nC := km1.Dims()[0], km1.Dims()[1]
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if km2.At(rowIndex, colIndex) != float64(km1.At(rowIndex, colIndex).(symbolic.K)) {
				t.Errorf(
					"Expected km2.At(%v,%v) to be %v; received %v",
					rowIndex,
					colIndex,
					km1.At(rowIndex, colIndex),
					km2.At(rowIndex, colIndex),
				)
			}
		}
	}
}

/*
TestKMatrix_String1
Description:

	Tests that the String() method properly returns a string representation
	of the KMatrix.
	Checks that all elements are properly represented.
*/
func TestKMatrix_String1(t *testing.T) {
	// Constants
	km1 := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Test
	str := km1.String()

	// Verify that the string representation is correct
	// - Check that each element from km1 is contained (i.e.,
	//	 is a substring of the string representation)
	nR, nC := km1.Dims()[0], km1.Dims()[1]
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			elt := km1.At(rowIndex, colIndex)
			if !strings.Contains(str, fmt.Sprintf("%v", elt)) {
				t.Errorf(
					"Expected string representation to contain %v; received %v",
					elt,
					str,
				)
			}
		}
	}
}

/*
TestKMatrix_DerivativeWrt1
Description:

	Tests that the DerivativeWrt() method properly returns a matrix of all
	zeros when the KMatrix is any nonzero constant.
*/
func TestKMatrix_DerivativeWrt1(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Verify that derivative is a KMatrix as well
	derivative := A.DerivativeWrt(symbolic.NewVariable())
	if _, ok := derivative.(symbolic.KMatrix); !ok {
		t.Errorf("Expected derivative to be a KMatrix; got %T", derivative)
	}

	// Verify that the derivative is all zeros
	nR, nC := A.Dims()[0], A.Dims()[1]
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if elt := float64(derivative.(symbolic.KMatrix).At(rowIndex, colIndex).(symbolic.K)); elt != 0.0 {
				t.Errorf("Expected derivative to be all zeros; got %v", derivative)
			}
		}
	}
}

/*
TestKMatrix_Degree1
Description:

	Tests that the Degree() method properly returns the degree of the KMatrix. (zero)
*/
func TestKMatrix_Degree1(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Verify that the degree is zero
	if deg := A.Degree(); deg != 0 {
		t.Errorf("Expected degree to be 0; got %v", deg)
	}
}

/*
TestKMatrix_Substitute1
Description:

	Tests that the Substitute() method properly runs the substitute method when given
	a variable to substite for and an expression to substitute with. In this case, the substitution SHOULD
	NOT CHANGE ANYTHING.
*/
func TestKMatrix_Substitute1(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Substitute
	sub := A.Substitute(symbolic.NewVariable(), symbolic.NewVariable())

	// Verify that the result is the same as the original
	if !reflect.DeepEqual(A, sub) {
		t.Errorf("Expected substitution to not change anything; got %v", sub)
	}
}

/*
TestKMatrix_SubstituteAccordingTo1
Description:

	Tests that the SubstituteAccordingTo() method properly runs the substitute method when given
	a map of substitutions. In this case, the substitution SHOULD NOT CHANGE ANYTHING.
*/
func TestKMatrix_SubstituteAccordingTo1(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	// Substitute
	sub := A.SubstituteAccordingTo(
		map[symbolic.Variable]symbolic.Expression{
			symbolic.NewVariable(): symbolic.NewVariable(),
		})

	// Verify that the result is the same as the original
	if !reflect.DeepEqual(A, sub) {
		t.Errorf("Expected substitution to not change anything; got %v", sub)
	}
}

/*
TestKMatrix_Power1
Description:

	Tests that the Power() method properly raises the KMatrix to the power of 1.
*/
func TestKMatrix_Power1(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Power
	pow := A.Power(1)

	// Verify that the result is the same as the original
	if !reflect.DeepEqual(A, pow) {
		t.Errorf("Expected power to not change anything; got %v", pow)
	}
}

/*
TestKMatrix_Power2
Description:

	Tests that the Power() method properly raises the KMatrix to the power of 2.
*/
func TestKMatrix_Power2(t *testing.T) {
	// Setup
	A := getKMatrix.From([][]float64{
		{1, 0, 0},
		{0, 2, 0},
		{0, 0, 3},
	})

	// Power
	pow := A.Power(2)

	// Verify that the result is a diagonal matrix with
	// the diagonal elements being the squares of the original
	nR, nC := A.Dims()[0], A.Dims()[1]
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			if rowIndex == colIndex {
				if elt := float64(pow.(symbolic.KMatrix).At(rowIndex, colIndex).(symbolic.K)); elt != math.Pow(float64(A.At(rowIndex, colIndex).(symbolic.K)), 2) {
					t.Errorf("Expected diagonal element to be squared; got %v", pow)
				}
			} else {
				if elt := float64(pow.(symbolic.KMatrix).At(rowIndex, colIndex).(symbolic.K)); elt != 0 {
					t.Errorf("Expected off-diagonal element to be zero; got %v", pow)
				}
			}
		}
	}
}
