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

/*
TestVectorConstraint_Substitute1
Description:

	This function tests that the Substitute method properly panics
	if the input scalar expression is not well defined.
	(In this case, the left hand side is a monomial that is not well-defined.)
*/
func TestVectorConstraint_Substitute1(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariable()
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N)).ToMonomialVector()
	left[6] = symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 0},
		VariableFactors: []symbolic.Variable{x},
	}
	right := symbolic.NewVariableVector(N)

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the function for handling the panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.Substitute() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := left[6].Check()
		if !strings.Contains(rAsError.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected vc.Substitute() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Test
	vc.Substitute(x, x.Power(3).(symbolic.ScalarExpression))

	// Raise an error if the test did not panic
	t.Errorf(
		"Expected vc.Substitute() to panic; received nil",
	)

}

/*
TestVectorConstraint_Substitute2
Description:

	This function tests that the Substitute method properly returns
	a new vector constraint when the input scalar expressions are well-defined.
	In this case, we will start with a vector of monomials on the left and
	a vector of constants on the right.
	We will substitute a variable on the left hand side with a polynomial
	and expect the result to be a vector of polynomials on the left hand side
	and a vector of constants on the right hand side.
*/
func TestVectorConstraint_Substitute2(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	left := x.ToMonomialVector()
	right := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the new expression to replace one of the variables with:
	newExpr := x[0].Plus(x[1]).(symbolic.ScalarExpression)

	// Substitute the variable
	newVC := vc.Substitute(x[0], newExpr).(symbolic.VectorConstraint)

	// Check the left hand side
	// it should now be a vector of polynomials
	if _, ok := newVC.LeftHandSide.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vc.Substitute() to return a vector constraint with a polynomial vector on the left hand side; received %v",
			newVC.LeftHandSide,
		)
	}

	// Check the right hand side
	if _, ok := newVC.RightHandSide.(symbolic.KVector); !ok {
		t.Errorf(
			"Expected vc.Substitute() to return a vector constraint with a constant vector on the right hand side; received %v",
			newVC.RightHandSide,
		)
	}

	// Check the sense
	if newVC.Sense != vc.Sense {
		t.Errorf(
			"Expected vc.Substitute() to return a vector constraint with the same sense; received %v",
			newVC.Sense,
		)
	}

}

/*
TestVectorConstraint_SubstituteAccordingTo1
Description:

	This function tests that the SubstituteAccordingTo method properly panics
	if the input scalar expression is not well defined.
	(In this case, the left hand side is a monomial that is not well-defined.)
*/
func TestVectorConstraint_SubstituteAccordingTo1(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariable()
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N)).ToMonomialVector()
	left[6] = symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 0},
		VariableFactors: []symbolic.Variable{x},
	}
	right := symbolic.NewVariableVector(N)

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the function for handling the panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.SubstituteAccordingTo() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := left[6].Check()
		if !strings.Contains(rAsError.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected vc.SubstituteAccordingTo() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Test
	subMap := map[symbolic.Variable]symbolic.Expression{
		x: x.Power(3),
	}
	vc.SubstituteAccordingTo(subMap)

	// Raise an error if the test did not panic
	t.Errorf(
		"Expected vc.SubstituteAccordingTo() to panic; received nil",
	)

}

/*
TestVectorConstraint_SubstituteAccordingTo2
Description:

	This function tests that the SubstituteAccordingTo method properly returns
	a new vector constraint when the input scalar expressions are well-defined.
	In this case, we will start with a vector of monomials on the left and
	a vector of constants on the right.
	We will substitute a variable on the left hand side with a polynomial
	and expect the result to be a vector of polynomials on the left hand side
	and a vector of constants on the right hand side.
*/
func TestVectorConstraint_SubstituteAccordingTo2(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	left := x.ToMonomialVector()
	right := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the new expression to replace one of the variables with:
	newExpr := x[0].Plus(x[1]).(symbolic.ScalarExpression)

	// Substitute the variable
	subMap := map[symbolic.Variable]symbolic.Expression{
		x[0]: newExpr,
	}
	newVC := vc.SubstituteAccordingTo(subMap).(symbolic.VectorConstraint)

	// Check the left hand side
	// it should now be a vector of polynomials
	if _, ok := newVC.LeftHandSide.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vc.SubstituteAccordingTo() to return a vector constraint with a polynomial vector on the left hand side; received %v",
			newVC.LeftHandSide,
		)
	}

	// Check the right hand side
	if _, ok := newVC.RightHandSide.(symbolic.KVector); !ok {
		t.Errorf(
			"Expected vc.SubstituteAccordingTo() to return a vector constraint with a constant vector on the right hand side; received %v",
			newVC.RightHandSide,
		)
	}

	// Check the sense
	if newVC.Sense != vc.Sense {
		t.Errorf(
			"Expected vc.SubstituteAccordingTo() to return a vector constraint with the same sense; received %v",
			newVC.Sense,
		)
	}
}

/*
TestVectorConstraint_AsSimplifiedConstraint1
Description:

	This function tests that the AsSimplifiedConstraint method properly
	panics if the input vector constraint is not well-defined.
	(In this case, the left and right hand sides have different lengths.)
*/
func TestVectorConstraint_AsSimplifiedConstraint1(t *testing.T) {
	// Setup
	N := 7
	left := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	right := symbolic.NewVariableVector(N + 1)

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the function for handling the panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.AsSimplifiedConstraint() to panic; received nil",
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
				"Expected vc.AsSimplifiedConstraint() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Test
	vc.AsSimplifiedConstraint()

	// Raise an error if the test did not panic
	t.Errorf(
		"Expected vc.AsSimplifiedConstraint() to panic; received nil",
	)

}

/*
TestVectorConstraint_AsSimplifiedConstraint2
Description:

	This function tests that the AsSimplifiedConstraint method properly
	returns a simplified vector constraint when the original vector constraint
	is well-defined. In this case, the original vector constraint will have
	a left hand side that is a vector of monomials and a right hand side
	that is a vector of polynomials.
	After simplification, the left hand side should be a vector of polynomials
	and the right hand side should be a vector of constants.
*/
func TestVectorConstraint_AsSimplifiedConstraint2(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	left := x.ToMonomialVector()
	right := x.Multiply(2.0).Plus(symbolic.VecDenseToKVector(symbolic.OnesVector(N))).(symbolic.PolynomialVector)

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	simplifiedVC := vc.AsSimplifiedConstraint().(symbolic.VectorConstraint)

	// Check the left hand side
	// it should now be a vector of polynomials
	if _, ok := simplifiedVC.LeftHandSide.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vc.AsSimplifiedConstraint() to return a vector constraint with a polynomial vector on the left hand side; received %v",
			simplifiedVC.LeftHandSide,
		)
	}

	// Check the right hand side
	if _, ok := simplifiedVC.RightHandSide.(symbolic.KVector); !ok {
		t.Errorf(
			"Expected vc.AsSimplifiedConstraint() to return a vector constraint with a constant vector on the right hand side; received %v",
			simplifiedVC.RightHandSide,
		)
	}

	// Check the sense
	if simplifiedVC.Sense != vc.Sense {
		t.Errorf(
			"Expected vc.AsSimplifiedConstraint() to return a vector constraint with the same sense; received %v",
			simplifiedVC.Sense,
		)
	}
}

/*
TestVectorConstraint_Variables1
Description:

	This function tests that the Variables method properly
	returns the correct set of variables in a well-defined vector constraint.
	In this case, the left hand side will be a vector of N variables
	and the right hand side will be a vector of N different variables.
	The result should be a set containing all 2N variables.
*/
func TestVectorConstraint_Variables1(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	y := symbolic.NewVariableVector(N)
	left := x
	right := y

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	vars := vc.Variables()

	// Check the length of the variable set
	if len(vars) != 2*N {
		t.Errorf(
			"Expected vc.Variables() to return a set with %v variables; received %v",
			2*N,
			len(vars),
		)
	}

	// Check that all variables are present in the set
	for i := 0; i < N; i++ {
		idx, err := symbolic.FindInSlice(x[i], vars)
		if err != nil {
			t.Errorf(
				"Received an error when searching for variable %v in the set: %v",
				x[i],
				err)
		}
		if idx == -1 {
			t.Errorf(
				"Expected vc.Variables() to return a set containing variable %v; it was not found in the set",
				x[i],
			)
		}
		// Check for y[i]
		idx, err = symbolic.FindInSlice(y[i], vars)
		if err != nil {
			t.Errorf(
				"Received an error when searching for variable %v in the set: %v",
				y[i],
				err)
		}

		if idx == -1 {
			t.Errorf(
				"Expected vc.Variables() to return a set containing variable %v; it was not found in the set",
				y[i],
			)
		}
	}
}

/*
TestVectorConstraint_ImpliesThisIsAlsoSatisfied1
Description:

	This function tests that the ImpliesThisIsAlsoSatisfied method properly
	panics if the receiver vector constraint is not well-defined.
	(In this case, the left and right hand sides have different lengths.)
*/
func TestVectorConstraint_ImpliesThisIsAlsoSatisfied1(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	y := symbolic.NewVariableVector(N)
	left := x
	right := symbolic.NewVariableVector(N - 1)

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the second vector constraint
	vc2 := symbolic.VectorConstraint{x, y, symbolic.SenseLessThanEqual}

	// Test
	expectedError := vc.Check()
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
			)
		}

		rAsError := r.(error)
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.ImpliesThisIsAlsoSatisfied() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.ImpliesThisIsAlsoSatisfied(vc2)
}

/*
TestVectorConstraint_ImpliesThisIsAlsoSatisfied2
Description:

	This function tests that the ImpliesThisIsAlsoSatisfied method properly
	panics if:
	- The receiver vector constraint is well-defined, but
	- The input vector constraint is not well-defined.
	(In this case, the left and right hand sides of the input vector constraint
	have different lengths.)
*/
func TestVectorConstraint_ImpliesThisIsAlsoSatisfied2(t *testing.T) {
	// Setup
	N := 7
	x := symbolic.NewVariableVector(N)
	y := symbolic.NewVariableVector(N)
	left := x
	right := y

	// Create the vector constraint
	vc := symbolic.VectorConstraint{left, right, symbolic.SenseLessThanEqual}

	// Create the second vector constraint
	vc2 := symbolic.VectorConstraint{x, symbolic.NewVariableVector(N - 1), symbolic.SenseLessThanEqual}

	// Test
	expectedError := vc2.Check()
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
			)
		}

		rAsError := r.(error)
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.ImpliesThisIsAlsoSatisfied() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	vc.ImpliesThisIsAlsoSatisfied(vc2)
}

/*
TestVectorConstraint_ImpliesThisIsAlsoSatisfied3
Description:

	This function tests that the ImpliesThisIsAlsoSatisfied method properly
	returns false if the receiver and input vector constraints are well-defined,
	but the input constraint is not implied by the receiver constraint.
	In this case, the receiver constraint will be:
	 [1, 0;
	  0, 1] * x <= [1; 1]
	and the input constraint will be:
	 [1, 0;
	  0, 1] * x <= [2; 2]

	TODO: This constraint should return true, but the implementation
	does not currently handle this case. Let's fix this in a future update.
*/
func TestVectorConstraint_ImpliesThisIsAlsoSatisfied3(t *testing.T) {
	// Setup
	N := 2
	x := symbolic.NewVariableVector(N)
	left := x
	right := mat.NewVecDense(N, []float64{1, 1})

	// Create the vector constraint
	vc := symbolic.VectorConstraint{
		LeftHandSide:  left,
		RightHandSide: symbolic.VecDenseToKVector(*right),
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the second vector constraint
	right2 := mat.NewVecDense(N, []float64{2, 2})
	vc2 := symbolic.VectorConstraint{
		LeftHandSide:  left,
		RightHandSide: symbolic.VecDenseToKVector(*right2),
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Test
	result := vc.ImpliesThisIsAlsoSatisfied(vc2) // TODO: Fix this test to return true
	if result {
		t.Errorf(
			"Expected vc.ImpliesThisIsAlsoSatisfied() to return true; received false",
		)
	}
}

/*
TestVectorConstraint_ImpliesThisIsAlsoSatisfied4
Description:

	This function tests that the ImpliesThisIsAlsoSatisfied method properly
	returns true if:
	- the receiver vector constraints are well-defined,
	- the input is a well-defined scalar constraint, and
	- the input constraint is implied by one of the receiver constraints.
	In this case, the receiver constraint will be:
		[1, 0;
		 0, 1] * x <= [1; 1]
	and the input constraint will be:
		[1, 0] * x <= 1
*/
func TestVectorConstraint_ImpliesThisIsAlsoSatisfied4(t *testing.T) {
	// Setup
	N := 2
	x := symbolic.NewVariableVector(N)
	left := x
	right := mat.NewVecDense(N, []float64{1, 1})

	// Create the vector constraint
	vc := symbolic.VectorConstraint{
		LeftHandSide:  left,
		RightHandSide: symbolic.VecDenseToKVector(*right),
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the scalar constraint
	left2 := x.AtVec(0)
	right2 := symbolic.K(1)
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  left2,
		RightHandSide: right2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Test
	result := vc.ImpliesThisIsAlsoSatisfied(sc)
	if !result {
		t.Errorf(
			"Expected vc.ImpliesThisIsAlsoSatisfied() to return true; received false",
		)
	}
}
