package symbolic_test

/*
scalar_constraint_test.go
Description:
	Tests for the functions mentioned in the scalar_constraint.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestScalarConstraint_Left1
Description:

	Tests that the left side of a scalar constraint is properly returned.
	In this example, where there is a variable on the left and a monomial on
	the right, the left side should be the variable.
*/
func TestScalarConstraint_Left1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Degrees:         []int{1, 2},
		VariableFactors: []symbolic.Variable{x, y},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{x, m, symbolic.SenseLessThanEqual}

	// Cast left to variable
	leftAsV, ok := sc.Left().(symbolic.Variable)
	if !ok {
		t.Errorf(
			"Expected sc.Left() to be a symbolic.Variable; received %T",
			sc.Left(),
		)
	}

	// Verify that left has the same id as x
	if leftAsV.ID != x.ID {
		t.Errorf(
			"Expected sc.Left() to have id %v; received %v",
			x.ID,
			leftAsV.ID,
		)
	}
}

/*
TestScalarConstraint_Right1
Description:

	Verifies that the right side of a scalar constraint is properly returned.
	In this example, where there is a variable on the left and a monomial on
	the right, the right side should be the monomial.
*/
func TestScalarConstraint_Right1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Degrees:         []int{1, 2},
		VariableFactors: []symbolic.Variable{x, y},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{x, m, symbolic.SenseLessThanEqual}

	// Cast right to monomial
	rightAsM, ok := sc.Right().(symbolic.Monomial)
	if !ok {
		t.Errorf(
			"Expected sc.Right() to be a symbolic.Monomial; received %T",
			sc.Right(),
		)
	}

	// Verify that right has the same id as x
	if symbolic.NumVariables(rightAsM) != 2 {
		t.Errorf(
			"Expected sc.Right() to have 2 variables; received %v",
			symbolic.NumVariables(rightAsM),
		)
	}
}

/*
TestScalarConstraint_IsLinear1
Description:

	Verifies that a scalar constraint containing a constant on the left and a
	variable on the right is linear.
*/
func TestScalarConstraint_IsLinear1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{c2, x, symbolic.SenseLessThanEqual}

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}
}

/*
TestScalarConstraint_IsLinear2
Description:

	Verifies that a scalar constraint containing a constant on the left
	and a monomial with two unique degrees on the right is NOT linear.
*/
func TestScalarConstraint_IsLinear2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Degrees:         []int{1, 2},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{c2, m, symbolic.SenseLessThanEqual}

	// Verify that the constraint is linear
	if sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}
}
