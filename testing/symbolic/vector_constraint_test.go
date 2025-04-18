package symbolic_test

/*
vector_constraint_test.go
Description:
	Tests for the functions mentioned in the vector_constraint.go file.
*/

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
TestVectorConstraint_Check1
Description:

	This function verifies that Check returns an error is returned when the left
	and right hand sides of the vector constraint have different lengths.
*/
func TestVectorConstraint_Check1(t *testing.T) {
	// Constants
	N := 3
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)

	// Test
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}
	err := vc.Check()
	if err == nil {
		t.Errorf(
			"Expected vc.Check() to return an error; received nil",
		)
	} else {
		expectedError := smErrors.VectorDimensionError{
			Arg1:      left,
			Arg2:      right,
			Operation: fmt.Sprintf("Comparison (%v)", vc.Sense),
		}
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.Check() to return error \"%v\"; received \"%v\"",
				expectedError.Error(),
				err.Error(),
			)
		}
	}
}

/*
TestVectorConstraint_Check2
Description:

	Tests that the Check method returns nil for a well-defined vector constraint.
	i.e., both vectors have the same length.
*/
func TestVectorConstraint_Check2(t *testing.T) {
	// Constants
	N := 3
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)

	// Test
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}
	err := vc.Check()
	if err != nil {
		t.Errorf(
			"Expected vc.Check() to return nil; received \"%v\"",
			err.Error(),
		)
	}
}

/*
TestVectorConstraint_Dims1
Description:

	This function tests that a vector constraint composed of two length 10 vectors
	returns dims of (10, 1).
*/
func TestVectorConstraint_Dims1(t *testing.T) {
	// Constants
	N := 10
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)

	// Test
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}
	dims := vc.Dims()
	if dims[0] != 10 || dims[1] != 1 {
		t.Errorf(
			"Expected vc.Dims() to return [10, 1]; received %v",
			dims,
		)
	}
}

/*
TestVectorConstraint_Dims2
Description:

	This test verifies that the Dims() method panics if the VectorConstraint is improperly
	defined.
*/
func TestVectorConstraint_Dims2(t *testing.T) {
	// Constants
	N := 10
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)

	// Test
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.Dims() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.VectorDimensionError{
			Arg1:      vc.LeftHandSide,
			Arg2:      vc.RightHandSide,
			Operation: fmt.Sprintf("Comparison (%v)", vc.Sense),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.Dims() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()
	vc.Dims()
}

/*
TestVectorConstraint_AtVec1
Description:

	This function tests that the AtVec method returns the correct scalar constraint
	for a vector constraint of dimension (7, 1).
*/
func TestVectorConstraint_AtVec1(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	for i := 0; i < N; i++ {
		sc := vc.AtVec(i)
		if sc.LeftHandSide.Dims()[0] != 1 || sc.RightHandSide.Dims()[0] != 1 || sc.Sense != symbolic.SenseLessThanEqual {
			t.Errorf(
				"Expected vc.AtVec(%v) to return a scalar constraint with LHS = 1, RHS = 1, and Sense = SenseLessThanEqual; received %v",
				i,
				sc,
			)
		}

		lhsAsK := sc.LeftHandSide.(symbolic.K)
		if float64(lhsAsK) != 1 {
			t.Errorf(
				"Expected vc.AtVec(%v).LeftHandSide to be 1; received %v",
				i,
				lhsAsK,
			)
		}
	}
}

/*
TestVectorConstraint_AtVec2
Description:

	This function tests that the AtVec method panics if the index is out of bounds.
*/
func TestVectorConstraint_AtVec2(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.AtVec(%v) to panic; received nil",
				N,
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.InvalidVectorIndexError{Index: N, Expression: vc}
		if !strings.Contains(
			rAsError.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"Expected vc.AtVec(%v) to panic with error \"Index %v is out of bounds for vector constraint of length %v\"; received \"%v\"",
				N,
				N,
				N,
				rAsError.Error(),
			)
		}
	}()
	vc.AtVec(N)
}

/*
TestVectorConstraint_AtVec3
Description:

	This test verifies that the AtVec method panics if the index is fine but the
	VectorConstraint is improperly defined.
*/
func TestVectorConstraint_AtVec3(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.AtVec(%v) to panic; received nil",
				N,
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.VectorDimensionError{
			Arg1:      vc.LeftHandSide,
			Arg2:      vc.RightHandSide,
			Operation: fmt.Sprintf("Comparison (%v)", vc.Sense),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.Dims() to panic with error \"Left hand side has dimension %v, but right hand side has dimension %v!\"; received \"%v\"",
				vc.LeftHandSide.Len(),
				vc.RightHandSide.Len(),
				rAsError.Error(),
			)
		}
	}()

	vc.AtVec(N - 1)
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation1
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly panics if the input VectorConstraint is not well defined.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation1(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.VectorDimensionError{
			Operation: fmt.Sprintf("Comparison (%v)", vc.Sense),
			Arg1:      vc.LeftHandSide,
			Arg2:      vc.RightHandSide,
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearInequalityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation2
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly panics when the left hand side is not linear.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation2(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	x := symbolic.NewVariableVector(N)
	right := x.Plus(x.Transpose().Multiply(x))
	vc := left.LessEq(right).(symbolic.VectorConstraint)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearInequalityConstraintRepresentation",
			Expression: vc.Right(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearInequalityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation3
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	handles a well-defined, lienar vector constraint properly.
	In this case, the left hand side should be a 2x2 matrix of ones and the right hand side
	should be a vector containing a 1 and a 2. This will be a LessEq constraint.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation3(t *testing.T) {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	left := symbolic.VecDenseToKVector(symbolic.ZerosVector(N)).Plus(x.AtVec(0)).Plus(x.AtVec(1))
	right := mat.NewVecDense(N, []float64{1, 2})
	vc := left.LessEq(right).(symbolic.VectorConstraint)
	fmt.Printf("vc_0: %v\n", vc)

	// Test
	A, b := vc.LinearInequalityConstraintRepresentation()

	nRowsA, nColsA := A.Dims()
	if nRowsA != N || nColsA != 2 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation() to return a matrix of dimension %v; received dimension (%v, %v)",
			[]int{N, 2},
			nRowsA, nColsA,
		)
	}

	if b.AtVec(0) != 1 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation()'s b vector to contain a 1 at the %v-th index; received %v",
			0,
			b.AtVec(0),
		)
	}

	if b.AtVec(1) != 2 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation()'s b vector to contain a 2 at the %v-th index; received %v",
			1,
			b.AtVec(1),
		)
	}
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation4
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	returns the proper matrix and vector for a well-defined, lienar vector constraint.
	We will construct a vector constraint of the form:
	 [1, 0;
	  0, 1] * x <= [1; 2]
	where x is a vector of variables.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation4(t *testing.T) {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	left := x
	right := mat.NewVecDense(N, []float64{1, 2})
	vc := left.LessEq(right).(symbolic.VectorConstraint)

	// Test
	A, b := vc.LinearInequalityConstraintRepresentation()

	nRowsA, nColsA := A.Dims()
	if nRowsA != N || nColsA != 2 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation() to return a matrix of dimension %v; received dimension (%v, %v)",
			[]int{N, 1},
			nRowsA, nColsA,
		)
	}

	if b.AtVec(0) != 1 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s b vector to contain a 1 at the %v-th index; received %v",
			0,
			b.AtVec(0),
		)
	}

	if b.AtVec(1) != 2 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s b vector to contain a 2 at the %v-th index; received %v",
			1,
			b.AtVec(1),
		)
	}
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation5
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly panics if the left hand side is not linear.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation5(t *testing.T) {
	// Constants
	N := 7
	right := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	x := symbolic.NewVariableVector(N)
	left := x.Plus(x.Transpose().Multiply(x))
	vc := left.GreaterEq(right).(symbolic.VectorConstraint)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearInequalityConstraintRepresentation",
			Expression: vc.Left(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearInequalityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation6
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly flips the sign of the matrices and vectors.
	In this case, we will create the same expressions as
	TestVectorConstraint_LinearInequalityConstraintRepresentation3 but with a GreaterEq constraint
	sense.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation6(t *testing.T) {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	left := symbolic.VecDenseToKVector(symbolic.ZerosVector(N)).Plus(x.AtVec(0)).Plus(x.AtVec(1))
	right := mat.NewVecDense(N, []float64{1, 2})
	vc := left.GreaterEq(right).(symbolic.VectorConstraint)

	// Test
	A, b := vc.LinearInequalityConstraintRepresentation()

	nRowsA, nColsA := A.Dims()
	if nRowsA != N || nColsA != 2 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation() to return a matrix of dimension %v; received dimension (%v, %v)",
			[]int{N, 2},
			nRowsA, nColsA,
		)
	}

	if b.AtVec(0) != -1 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation()'s b vector to contain a -1 at the %v-th index; received %v",
			0,
			b.AtVec(0),
		)
	}

	if b.AtVec(1) != -2 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation()'s b vector to contain a -2 at the %v-th index; received %v",
			1,
			b.AtVec(1),
		)
	}
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation7
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly panics when called with a constraint that is not based on inequality.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation7(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.InequalityConstraintRequiredError{
			Operation: "LinearInequalityConstraintRepresentation",
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearInequalityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearInequalityConstraintRepresentation8
Description:

	This function tests that the LinearInequalityConstraintRepresentation method
	properly returns a matrix of shape (2, 2) and a vector of shape (2, 1)
	for a well-defined, lienar vector constraint. This constraint will only contain
	1 variable, but the w.r.t. variable will contain 2 variables.
*/
func TestVectorConstraint_LinearInequalityConstraintRepresentation8(t *testing.T) {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	left := x.AtVec(0).Plus(symbolic.ZerosVector(N))
	right := mat.NewVecDense(N, []float64{1, 2})
	vc := left.LessEq(right).(symbolic.VectorConstraint)

	// Test
	A, b := vc.LinearInequalityConstraintRepresentation(x)

	nRowsA, nColsA := A.Dims()
	if nRowsA != N || nColsA != 2 {
		t.Errorf(
			"Expected vc.LinearInequalityConstraintRepresentation() to return a matrix of dimension %v; received dimension (%v, %v)",
			[]int{N, 2},
			nRowsA, nColsA,
		)
	}

	if b.AtVec(0) != 1 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s b vector to contain a 1 at the %v-th index; received %v",
			0,
			b.AtVec(0),
		)
	}

	if b.AtVec(1) != 2 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s b vector to contain a 2 at the %v-th index; received %v",
			1,
			b.AtVec(1),
		)
	}
}

/*
TestVectorConstraint_LinearEqualityConstraintRepresentation1
Description:

	This function tests that the LinearEqualityConstraintRepresentation method
	properly panics if the input VectorConstraint is not well defined.
*/
func TestVectorConstraint_LinearEqualityConstraintRepresentation1(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.VectorDimensionError{
			Operation: fmt.Sprintf("Comparison (%v)", vc.Sense),
			Arg1:      vc.LeftHandSide,
			Arg2:      vc.RightHandSide,
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearEqualityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearEqualityConstraintRepresentation2
Description:

	This function tests that the LinearEqualityConstraintRepresentation method
	properly panics when the LEFT hand side is not linear.
*/
func TestVectorConstraint_LinearEqualityConstraintRepresentation2(t *testing.T) {
	// Constants
	N := 7
	right := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	x := symbolic.NewVariableVector(N)
	left := x.Plus(x.Transpose().Multiply(x))
	vc := left.Eq(right).(symbolic.VectorConstraint)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearEqualityConstraintRepresentation",
			Expression: vc.Left(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearEqualityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearEqualityConstraintRepresentation3
Description:

	This function tests that the LinearEqualityConstraintRepresentation method
	properly panics when the RIGHT hand side is not linear.
*/
func TestVectorConstraint_LinearEqualityConstraintRepresentation3(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	x := symbolic.NewVariableVector(N)
	right := x.Plus(x.Transpose().Multiply(x))
	vc := left.Eq(right).(symbolic.VectorConstraint)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearEqualityConstraintRepresentation",
			Expression: vc.Right(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearEqualityConstraintRepresentation()
}

/*
TestVectorConstraint_LinearEqualityConstraintRepresentation4
Description:

	This function tests that the LinearEqualityConstraintRepresentation method
	properly returns the matrix and vector for a well-defined, lienar vector constraint.
	In this case, we will create a vector constraint of the form:
	 [1, 0;
	  0, 1] * x = [1; 2]
	where x is a vector of variables.
*/
func TestVectorConstraint_LinearEqualityConstraintRepresentation4(t *testing.T) {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	left := x
	right := mat.NewVecDense(N, []float64{1, 2})
	vc := left.Eq(right).(symbolic.VectorConstraint)

	// Test
	C, d := vc.LinearEqualityConstraintRepresentation()

	nRowsC, nColsC := C.Dims()
	if nRowsC != N || nColsC != N {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation() to return a matrix of dimension %v; received dimension (%v, %v)",
			[]int{N, N},
			nRowsC, nColsC,
		)
	}

	if d.AtVec(0) != 1 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s d vector to contain a 1 at the %v-th index; received %v",
			0,
			d.AtVec(0),
		)
	}

	if d.AtVec(1) != 2 {
		t.Errorf(
			"Expected vc.LinearEqualityConstraintRepresentation()'s d vector to contain a 2 at the %v-th index; received %v",
			1,
			d.AtVec(1),
		)
	}
}

/*
TestVectorConstraint_LinearEqualityConstraintRepresentation5
Description:

	This function tests that the LinearEqualityConstraintRepresentation method
	properly panics when called with a constraint that is not based on equality.
*/
func TestVectorConstraint_LinearEqualityConstraintRepresentation5(t *testing.T) {
	// Constants
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N)
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.EqualityConstraintRequiredError{
			Operation: "LinearEqualityConstraintRepresentation",
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.LinearEqualityConstraintRepresentation()
}
