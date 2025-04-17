package symbolic_test

/*
scalar_constraint_test.go
Description:
	Tests for the functions mentioned in the scalar_constraint.go file.
*/

import (
	"strings"
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
	sc = sc.Simplify()

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
	sc = sc.Simplify()

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

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation1
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. Here, we provide a constraint with two variables that are LessThanEqual
	to a constant. The expected output is a vector of all ones and a constant 2.5.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation1(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(symbolic.OnesVector(2)).LessEq(c2)

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation()

	// Verify that the vector is all ones
	for ii := 0; ii < A.Len(); ii++ {
		if A.AtVec(ii) != 1 {
			t.Errorf(
				"Expected A[%v] to be 1; received %v",
				ii,
				A.AtVec(ii),
			)
		}
	}

	// Verify that the constant is 2.5
	if b != 2.5 {
		t.Errorf(
			"Expected b to be 2.5; received %v",
			b,
		)
	}
}

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation2
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. Here, we provide a constraint with two variables that are GreaterThanEqual
	to a constant. The expected output is a vector of all negative ones and a constant -2.5.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation2(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(symbolic.OnesVector(2)).GreaterEq(c2)

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation()

	// Verify that the vector is all negative ones
	for ii := 0; ii < A.Len(); ii++ {
		if A.AtVec(ii) != -1 {
			t.Errorf(
				"Expected A[%v] to be -1; received %v",
				ii,
				A.AtVec(ii),
			)
		}
	}

	// Verify that the constant is -2.5
	if b != -2.5 {
		t.Errorf(
			"Expected b to be -2.5; received %v",
			b,
		)
	}
}

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation3
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. Here, we provide a constraint with two variable that is LessThanEqual
	to one variable added to a constant. The expected output is a vector with a
	single one and a single zero and a constant 2.5.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation3(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(symbolic.OnesVector(2)).LessEq(x.AtVec(1).Plus(c2))

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation()

	// Verify that the vector is all ones

	if A.AtVec(0) != 1 {
		t.Errorf("Expected A[0] to be 1; received %v", A.AtVec(0))
	}
	if A.AtVec(1) != 0.0 {
		t.Errorf("Expected A[1] to be 0; received %v", A.AtVec(1))
	}

	// Verify that the constant is 2.5
	if b != 2.5 {
		t.Errorf("Expected b to be 2.5; received %v", b)
	}
}

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation4
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method panics if the left hand side
	is not a polynomial like scalar.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation4(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(x).LessEq(c2)

	// Verify that the constraint is linear
	if sc.IsLinear() {
		t.Errorf(
			"Expected sc to be nonlinear; received IsLinear() = %v",
			sc.IsLinear(),
		)
	}

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
			Expression: sc.Left(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Get linear representation
	sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation()
}

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation5
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method panics if the RIGHT hand side
	is not a polynomial like scalar.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation5(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := c2.LessEq(x.Transpose().Multiply(x))

	// Verify that the constraint is linear
	if sc.IsLinear() {
		t.Errorf(
			"Expected sc to be nonlinear; received IsLinear() = %v",
			sc.IsLinear(),
		)
	}

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
			Expression: sc.Right(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearInequalityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Get linear representation
	sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation()
}
