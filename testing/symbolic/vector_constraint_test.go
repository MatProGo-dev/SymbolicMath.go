package symbolic_test

/*
vector_constraint_test.go
Description:
	Tests for the functions mentioned in the vector_constraint.go file.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
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
		if err.Error() != "Left hand side has dimension 3, but right hand side has dimension 4!" {
			t.Errorf(
				"Expected vc.Check() to return error \"Left hand side has dimension 3, but right hand side has dimension 4!\"; received \"%v\"",
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
		if rAsError.Error() != (fmt.Errorf(
			"Left hand side has dimension %v, but right hand side has dimension %v!",
			vc.LeftHandSide.Len(),
			vc.RightHandSide.Len(),
		)).Error() {
			t.Errorf(
				"Expected vc.Dims() to panic with error \"Left hand side has dimension %v, but right hand side has dimension %v!\"; received \"%v\"",
				vc.LeftHandSide.Len(),
				vc.RightHandSide.Len(),
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
		if rAsError.Error() != (fmt.Errorf(
			"Left hand side has dimension %v, but right hand side has dimension %v!",
			vc.LeftHandSide.Len(),
			vc.RightHandSide.Len(),
		)).Error() {
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
