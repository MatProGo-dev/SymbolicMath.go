package symbolic_test

/*
constraint_test.go
Description:
	Tests for the functions mentioned in the constraint.go file.
*/

import (
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
TestConstraint_IsConstraint1
Description:

	Tests to make sure that an expression object is identified as not a constraint.
*/
func TestConstraint_IsConstraint1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.IsConstraint(x) {
		t.Errorf(
			"Expected IsConstraint(%T) to be false; received true",
			x,
		)
	}
}

/*
TestConstraint_IsConstraint2
Description:

	Verifies that a scalar constraint is identified as a constraint.
*/
func TestConstraint_IsConstraint2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 2},
		VariableFactors: []symbolic.Variable{x, y},
	}
	sc := symbolic.ScalarConstraint{x, m, symbolic.SenseLessThanEqual}

	// Test
	if !symbolic.IsConstraint(sc) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			sc,
		)
	}
}

/*
TestConstraint_IsConstraint3
Description:

	Verifies that a vector constraint is identified as a constraint.
*/
func TestConstraint_IsConstraint3(t *testing.T) {
	// Constants
	N := 11
	x := symbolic.NewVariableVector(N)
	kv2 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	vConstr := symbolic.VectorConstraint{x, kv2, symbolic.SenseLessThanEqual}

	// Test
	if !symbolic.IsConstraint(vConstr) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			vConstr,
		)
	}

}

/*
TestConstraint_IsConstraint4
Description:

	Verifies that a matrix constraint is identified as a constraint.
*/
func TestConstraint_IsConstraint4(t *testing.T) {
	// Constants
	N := 11
	mk1 := symbolic.DenseToKMatrix(symbolic.Identity(N))
	mk2 := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(N, N))

	mConstr := symbolic.MatrixConstraint{mk1, mk2, symbolic.SenseGreaterThanEqual}

	// Test
	if !symbolic.IsConstraint(mConstr) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			mConstr,
		)
	}
}

/*
TestConstraint_IsConstraint5
Description:

	Verifies that a scalar constraint pointer is identified as a constraint.
*/
func TestConstraint_IsConstraint5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 2},
		VariableFactors: []symbolic.Variable{x, y},
	}
	sc := symbolic.ScalarConstraint{x, m, symbolic.SenseLessThanEqual}

	// Test
	if !symbolic.IsConstraint(&sc) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			sc,
		)
	}
}

/*
TestConstraint_IsConstraint6
Description:

	Verifies that a vector constraint pointer is identified as a constraint.
*/
func TestConstraint_IsConstraint6(t *testing.T) {
	// Constants
	N := 11
	x := symbolic.NewVariableVector(N)
	kv2 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	vConstr := symbolic.VectorConstraint{x, kv2, symbolic.SenseLessThanEqual}

	// Test
	if !symbolic.IsConstraint(&vConstr) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			vConstr,
		)
	}
}

/*
TestConstraint_IsConstraint7
Description:

	Verifies that a matrix constraint pointer is identified as a constraint.
*/
func TestConstraint_IsConstraint7(t *testing.T) {
	// Constants
	N := 11
	mk1 := symbolic.DenseToKMatrix(symbolic.Identity(N))
	mk2 := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(N, N))

	mConstr := symbolic.MatrixConstraint{mk1, mk2, symbolic.SenseGreaterThanEqual}

	// Test
	if !symbolic.IsConstraint(&mConstr) {
		t.Errorf(
			"Expected IsConstraint(%T) to be true; received false",
			mConstr,
		)
	}
}

/*
TestConstraint_VariablesInThisConstraint1
Description:

	Verifies that the VariablesInThisConstraint function works as expected.
	We verify that the method properly returns 5 unique variables when
	there are 3 variables on the left hand side and 3 variables on the right hand side,
	but one of the variables is shared between the two sides.
*/
func TestConstraint_VariablesInThisConstraint1(t *testing.T) {
	// Constants
	x1 := symbolic.NewVariable()
	x2 := symbolic.NewVariable()
	x3 := symbolic.NewVariable()
	x4 := symbolic.NewVariable()
	x5 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 2, 1},
		VariableFactors: []symbolic.Variable{x1, x2, x3},
	}

	m2 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{3, 1, 1},
		VariableFactors: []symbolic.Variable{x3, x4, x5},
	}

	sc := symbolic.ScalarConstraint{m1, m2, symbolic.SenseEqual}

	// Test
	vars := symbolic.VariablesInThisConstraint(sc)
	if len(vars) != 5 {
		t.Errorf(
			"Expected VariablesInThisConstraint to return 5 unique variables; received %d",
			len(vars),
		)
	}
}
