package symbolic_test

/*
scalar_constraint_test.go
Description:
	Tests for the functions mentioned in the scalar_constraint.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
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
		Exponents:       []int{1, 2},
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
		Exponents:       []int{1, 2},
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
		Exponents:       []int{1, 2},
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

/*
TestScalarConstraint_IsLinear3
Description:

	Verifies that the method panics if the left hand side
	is not a polynomial like scalar.
	TODO: Finish this test
*/

/*
TestScalarConstraint_Simplify1
Description:

	This function tests that the Simplify() method of a ScalarConstraint
	formed between a variable on the left hand side and a constant on the
	right.
*/
func TestScalarConstraint_Simplify1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{x, c2, symbolic.SenseLessThanEqual}

	// Simplify
	sc, err := sc.Simplify()
	if err != nil {
		t.Errorf(
			"Expected no error from sc.Simplify(); received %v",
			err,
		)
	}

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Verify that the left hand side is a variable
	if _, ok := sc.Left().(symbolic.Variable); !ok {
		t.Errorf(
			"Expected sc.Left() to be a symbolic.Variable; received %T",
			sc.Left(),
		)
	}

	// Verify that the right hand side is a constant
	if _, ok := sc.Right().(symbolic.K); !ok {
		t.Errorf(
			"Expected sc.Right() to be a symbolic.K; received %T",
			sc.Right(),
		)
	}
}

/*
TestScalarConstraint_Simplify2
Description:

	This function tests that the Simplify() method of a ScalarConstraint
	formed between a constant on the left hand side and a variable on the
	right.
	This test verifies that the simplified function has a polynomial on the
	left and zero on the right.
*/
func TestScalarConstraint_Simplify2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{c2, x, symbolic.SenseLessThanEqual}

	// Simplify
	sc, err := sc.Simplify()
	if err != nil {
		t.Errorf(
			"Expected no error from sc.Simplify(); received %v",
			err,
		)
	}

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Verify that the left hand side is a variable
	if _, ok := sc.Left().(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected sc.Left() to be a symbolic.Polynomial; received %T",
			sc.Left(),
		)
	}

	// Verify that the right hand side is a constant
	if _, ok := sc.Right().(symbolic.K); !ok {
		t.Errorf(
			"Expected sc.Right() to be a symbolic.K; received %T",
			sc.Right(),
		)
	}

	if float64(sc.Right().(symbolic.K)) != 0 {
		t.Errorf(
			"Expected sc.Right() to be 0; received %v",
			sc.Right(),
		)
	}
}

/*
TestScalarConstraint_Check1
Description:

	Tests the Check() method of a scalar constraint.
	Verifies that an error is returned when the left hand side
	is a NOT well-defined variable and the right hand side is a
	constant.
*/
func TestScalarConstraint_Check1(t *testing.T) {
	// Constants
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		symbolic.Variable{},
		c2,
		symbolic.SenseLessThanEqual,
	}

	// Check
	err := sc.Check()
	if err == nil {
		t.Errorf(
			"Expected an error from sc.Check(); received nil",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			sc.LeftHandSide.Check().Error(),
		) {
			t.Errorf(
				"Expected error message to contain %v; received %v",
				sc.LeftHandSide.Check(),
				err.Error(),
			)
		}
	}
}

/*
TestScalarConstraint_Check2
Description:

	Tests the Check() method of a scalar constraint.
	Verifies that an error is returned when the left hand side
	is a constant and the right hand side is a NOT well-defined
	variable.
*/
func TestScalarConstraint_Check2(t *testing.T) {
	// Constants
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		c2,
		symbolic.Variable{},
		symbolic.SenseLessThanEqual,
	}

	// Check
	err := sc.Check()
	if err == nil {
		t.Errorf(
			"Expected an error from sc.Check(); received nil",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			sc.RightHandSide.Check().Error(),
		) {
			t.Errorf(
				"Expected error message to contain %v; received %v",
				sc.RightHandSide.Check(),
				err.Error(),
			)
		}
	}
}

/*
TestScalarConstraint_Check3
Description:

	Tests the Check() method of a scalar constraint.
	Verifies that an error is returned when the left hand side
	is a constant and the right hand side is a well-defined monomial,
	BUT the sense is not an allowed value.
*/
func TestScalarConstraint_Check3(t *testing.T) {
	// Constants
	c2 := symbolic.K(3.14)
	x := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{x},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{
		c2,
		m,
		'?',
	}

	// Check
	err := sc.Check()
	if err == nil {
		t.Errorf(
			"Expected an error from sc.Check(); received nil",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			sc.ConstrSense().Check().Error(),
		) {
			t.Errorf(
				"Expected error message to contain '%v'; received %v",
				sc.ConstrSense().Check(),
				err.Error(),
			)
		}
	}
}

/*
TestScalarConstraint_Check4
Description:

	Tests the Check() method of a scalar constraint.
	Verifies that a completely well-defined constraint returns no error.
*/
func TestScalarConstraint_Check4(t *testing.T) {
	// Constants
	c2 := symbolic.K(3.14)
	x := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{x},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{
		c2,
		m,
		symbolic.SenseLessThanEqual,
	}

	// Check
	err := sc.Check()
	if err != nil {
		t.Errorf(
			"Expected no error from sc.Check(); received %v",
			err,
		)
	}
}
