package symbolic_test

import (
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
expression_test.go
Description:
	Tests for the functions mentioned in the expression.go file.
*/

/*
TestExpression_NumVariables1
Description:

	Computes the number of variables in a variable.
*/
func TestExpression_NumVariables1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.NumVariables(x) != 1 {
		t.Errorf(
			"The number of variables in a %T should be 1; received %v",
			x,
			symbolic.NumVariables(x),
		)
	}
}

/*
TestExpression_VariableIDs1
Description:

	Tests the variable IDs function for a single variable.
*/
func TestExpression_VariableIDs1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	ids := symbolic.VariableIDs(x)
	if len(ids) != 1 {
		t.Errorf(
			"Expected the variable IDs to be 1; received %v",
			ids,
		)
	}

	// Verify that the first element of the ids is the same as x's
	if ids[0] != x.ID {
		t.Errorf(
			"Expected the variable ID to be %v; received %v",
			x.ID,
			ids[0],
		)
	}
}

/*
TestExpression_ToExpression1
Description:

	Tests the ToExpression function panics if it is given
	an invalid expression (in this case, a string).
*/
func TestExpression_ToExpression1(t *testing.T) {
	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The ToExpression function should panic when given an invalid expression")
		}
	}()
	symbolic.ToExpression("x")
}

/*
TestExpression_HStack1
Description:

	Tests the HStack function for two variables.
	The result of the stacking should be a variable matrix with one row and two columns.
*/
func TestExpression_HStack1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Test
	result := symbolic.HStack(x, y)
	if result.Dims()[0] != 1 || result.Dims()[1] != 2 {
		t.Errorf(
			"Expected the result to be a 1x2 matrix; received %v",
			result.Dims(),
		)
	}

	// Verify that the result is a variable matrix
	if _, ok := result.(symbolic.VariableMatrix); !ok {
		t.Errorf(
			"Expected the result to be a VariableMatrix; received %T",
			result,
		)
	}
}

/*
TestExpression_HStack2
Description:

	Tests the HStack function for 4 scalar expressions. 3 of the expressions are
	constants and the last one is a variable.
	The result should be a monomial matrix with one row and 4 columns.
*/
func TestExpression_HStack2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	c1 := symbolic.K(1.0)
	c2 := symbolic.K(2.0)
	c3 := symbolic.K(3.0)

	// Test
	result := symbolic.HStack(c1, c2, c3, x)
	if result.Dims()[0] != 1 || result.Dims()[1] != 4 {
		t.Errorf(
			"Expected the result to be a 1x4 matrix; received %v",
			result.Dims(),
		)
	}

	// Verify that the result is a monomial matrix
	if _, ok := result.(symbolic.MonomialMatrix); !ok {
		t.Errorf(
			"Expected the result to be a MonomialMatrix; received %T",
			result,
		)
	}
}

/*
TestExpression_HStack3
Description:

	Tests the HStack function for 2 vector expressions.
	Each vector has 11 elements. One is a constant vector and the other is a variable vector.
	The result should be a monomial matrix with 11 rows and 2 columns.
*/
func TestExpression_HStack3(t *testing.T) {
	// Constants
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(11))
	vv2 := symbolic.NewVariableVector(11)

	// Test
	result := symbolic.HStack(kv1, vv2)
	if result.Dims()[0] != 11 || result.Dims()[1] != 2 {
		t.Errorf(
			"Expected the result to be an 11x2 matrix; received %v",
			result.Dims(),
		)
	}

	// Verify that the result is a monomial matrix
	if _, ok := result.(symbolic.MonomialMatrix); !ok {
		t.Errorf(
			"Expected the result to be a MonomialMatrix; received %T",
			result,
		)
	}
}

/*
TestExpression_HStack4
Description:

	Tests the HStack function for a matrix and a vector expression.
	The matrix is a constant matrix and the vector is a variable vector.
	The matrix is of shape 3x2 and the vector is of length 2.
	The HStack function should panic because the dimensions do not match.
*/
func TestExpression_HStack4(t *testing.T) {
	// Constants
	km1 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 2))
	vv2 := symbolic.NewVariableVector(2)

	// Test
	defer func() {
		err := recover().(error)
		if err == nil {
			t.Errorf("The HStack function should panic when the dimensions do not match")
		}

		// Collect the expected error which should be a dimension error and
		// compare it with the recovered error
		expectedError := smErrors.DimensionError{
			Operation: "HStack",
			Arg1:      km1,
			Arg2:      vv2,
		}
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected the error to be %v; received %v",
				expectedError,
				err,
			)
		}

	}()
	symbolic.HStack(km1, vv2)

	// The function should panic before this point
	t.Errorf("The HStack function should panic before this point")

}

/*
TestExpression_HStack5
Description:

	Tests the HStack function for a matrix and a vector expression.
	The matrix is a constant matrix and the vector is a variable vector.
	The matrix is of shape 3x2 and the vector is of length 3.
	The HStack function should not panic because the dimensions match.
	It should produce a monomial matrix with 3 rows and 3 columns.
*/
func TestExpression_HStack5(t *testing.T) {
	// Constants
	km1 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 2))
	vv3 := symbolic.NewVariableVector(3)

	// Test
	result := symbolic.HStack(km1, vv3)
	if result.Dims()[0] != 3 || result.Dims()[1] != 3 {
		t.Errorf(
			"Expected the result to be a 3x3 matrix; received %v",
			result.Dims(),
		)
	}

	// Verify that the result is a monomial matrix
	if _, ok := result.(symbolic.MonomialMatrix); !ok {
		t.Errorf(
			"Expected the result to be a MonomialMatrix; received %T",
			result,
		)
	}
}
