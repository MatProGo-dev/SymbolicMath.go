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
	A, b := sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation([]symbolic.Variable(x))

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

/*
TestScalarConstraint_LinearInequalityConstraintRepresentation6
Description:

	Tests the LinearInequalityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method correctly produces a vector with
	length 2 and a constant of value 2.1 when using a small cosntraint:
	x1 <= 2.1
	but when calculating the representation with respect to a vector of 2 variables.
*/
func TestScalarConstraint_LinearInequalityConstraintRepresentation6(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.1)

	// Create constraint
	sc := x.AtVec(0).LessEq(c2)

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearInequalityConstraintRepresentation(x)

	// Verify that the vector is all ones
	if A.AtVec(0) != 1 {
		t.Errorf("Expected A[0] to be 1; received %v", A.AtVec(0))
	}

	if A.AtVec(1) != 0 {
		t.Errorf("Expected A[1] to be 0; received %v", A.AtVec(1))
	}

	// Verify that the constant is 2.5
	if b != 2.1 {
		t.Errorf(
			"Expected b to be 2.5; received %v",
			b,
		)
	}
}

/*
TestScalarConstraint_LinearEqualityConstraintRepresentation1
Description:

	Tests the LinearEqualityConstraintRepresentation() method of a scalar
	constraint. Here, we provide a constraint with two variables that are LessThanEqual
	to a constant. The expected output is a vector of all ones and a constant 2.5.
*/
func TestScalarConstraint_LinearEqualityConstraintRepresentation1(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(symbolic.OnesVector(2)).Eq(c2)

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearEqualityConstraintRepresentation()

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
TestScalarConstraint_LinearEqualityConstraintRepresentation2
Description:

	Tests the LinearEqualityConstraintRepresentation() method of a scalar
	constraint. Here, we check that the method properly panics when we call
	the method on a constraint that is not an equality constraint.
*/
func TestScalarConstraint_LinearEqualityConstraintRepresentation2(t *testing.T) {
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

	sc.(symbolic.ScalarConstraint).LinearEqualityConstraintRepresentation()
}

/*
TestScalarConstraint_LinearEqualityConstraintRepresentation3
Description:

	Tests the LinearEqualityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method panics if the left hand side
	is not a polynomial like scalar (in this case, it will be nonlinear).
*/
func TestScalarConstraint_LinearEqualityConstraintRepresentation3(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := x.Transpose().Multiply(x).Eq(c2)

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
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearEqualityConstraintRepresentation",
			Expression: sc.Left(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.(symbolic.ScalarConstraint).LinearEqualityConstraintRepresentation()
}

/*
TestScalarConstraint_LinearEqualityConstraintRepresentation4
Description:

	Tests the LinearEqualityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method panics if the RIGHT hand side
	is not a polynomial like scalar (in this case, it will be nonlinear).
*/
func TestScalarConstraint_LinearEqualityConstraintRepresentation4(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.5)

	// Create constraint
	sc := c2.Eq(x.Transpose().Multiply(x))

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
				"Expected vc.LinearEqualityConstraintRepresentation() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := smErrors.LinearExpressionRequiredError{
			Operation:  "LinearEqualityConstraintRepresentation",
			Expression: sc.Right(),
		}
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected vc.LinearEqualityConstraintRepresentation() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.(symbolic.ScalarConstraint).LinearEqualityConstraintRepresentation()
}

/*
TestScalarConstraint_LinearEqualityConstraintRepresentation5
Description:

	Tests the LinearEqualityConstraintRepresentation() method of a scalar
	constraint. This test verifies that the method correctly produces a vector with
	length 2 and a constant of value 2.1 when using a small cosntraint:
	x1 = 2.1
	but when calculating the representation with respect to a vector of 2 variables.
*/
func TestScalarConstraint_LinearEqualityConstraintRepresentation5(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)
	c2 := symbolic.K(2.1)

	// Create constraint
	sc := x.AtVec(0).Eq(c2)

	// Verify that the constraint is linear
	if !sc.IsLinear() {
		t.Errorf(
			"Expected sc to be linear; received %v",
			sc.IsLinear(),
		)
	}

	// Get linear representation
	A, b := sc.(symbolic.ScalarConstraint).LinearEqualityConstraintRepresentation(x)

	// Verify that the vector is all ones
	if A.AtVec(0) != 1 {
		t.Errorf("Expected A[0] to be 1; received %v", A.AtVec(0))
	}

	if A.AtVec(1) != 0 {
		t.Errorf("Expected A[1] to be 0; received %v", A.AtVec(1))
	}

	// Verify that the constant is 2.5
	if b != 2.1 {
		t.Errorf(
			"Expected b to be 2.5; received %v",
			b,
		)
	}
}

/*
TestScalarConstraint_Substitute1
Description:

	This tests that the ScalarConstraint.Substitute() method properly panics
	when the left hand side is not a valid monomial.
*/
func TestScalarConstraint_Substitute1(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: v2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.Substitute() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m1.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.Substitute() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	// Call the method
	sc.Substitute(v2, symbolic.K(1))

	t.Errorf(
		"Expected sc.Substitute() to panic; received nil",
	)
}

/*
TestScalarConstraint_Substitute2
Description:

	This tests that the ScalarConstraint.Substitute() method properly
	returns a new ScalarConstraint when the left hand side is a valid monomial.
	In this case, we substitute a variable for a sum of two variables,
	which should lead to a new constraint with the same sense.
*/
func TestScalarConstraint_Substitute2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Substitute
	sum := y.Plus(x)
	sumAsSE, ok := sum.(symbolic.ScalarExpression)
	if !ok {
		t.Errorf(
			"Expected sum to be a symbolic.ScalarExpression; received %T",
			sum,
		)
	}

	newSc := sc.Substitute(x, sumAsSE)

	// Verify that the new left hand side is a polynomial and NOT a monomial
	if _, ok := newSc.Left().(symbolic.Monomial); ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Polynomial; received %T",
			newSc.Left(),
		)
	}

	if _, ok := newSc.Left().(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Polynomial; received %T",
			newSc.Left(),
		)
	}

	// Verify that the new right hand side is a constant 3.14
	m2, ok := newSc.Right().(symbolic.K)
	if !ok {
		t.Errorf(
			"Expected newSc.RightHandSide to be a symbolic.K; received %T",
			newSc.Right(),
		)
	}

	// Verify that the new right hand side is a constant 3.14
	if float64(m2) != 3.14 {
		t.Errorf(
			"Expected newSc.RightHandSide to be 3.14; received %v",
			m2,
		)
	}

	// Verify that the new constraint has the same sense
	if newSc.ConstrSense() != sc.Sense {
		t.Errorf(
			"Expected newSc.Sense to be %v; received %v",
			sc.Sense,
			newSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_SubstituteAccordingTo1
Description:

	This tests that the ScalarConstraint.SubstituteAccordingTo() method
	properly panics when the left hand side is not a valid monomial.
*/
func TestScalarConstraint_SubstituteAccordingTo1(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: v2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.SubstituteAccordingTo() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m1.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.SubstituteAccordingTo() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.SubstituteAccordingTo(nil)

	t.Errorf(
		"Expected sc.SubstituteAccordingTo() to panic; received nil",
	)
}

/*
TestScalarConstraint_SubstituteAccordingTo2
Description:

	This tests that the ScalarConstraint.SubstituteAccordingTo() method
	properly returns a new ScalarConstraint when the left hand side is a valid monomial.
	In this case, we substitute a variable for a sum of two variables,
	which should lead to a new constraint with the same sense.
*/
func TestScalarConstraint_SubstituteAccordingTo2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Substitute
	sum := y.Plus(x)
	sumAsSE, ok := sum.(symbolic.ScalarExpression)
	if !ok {
		t.Errorf(
			"Expected sum to be a symbolic.ScalarExpression; received %T",
			sum,
		)
	}

	newSc := sc.SubstituteAccordingTo(
		map[symbolic.Variable]symbolic.Expression{
			x: sumAsSE,
		},
	)

	// Verify that the new left hand side is a polynomial and NOT a monomial
	if _, ok := newSc.Left().(symbolic.Monomial); ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Polynomial; received %T",
			newSc.Left(),
		)
	}

	if _, ok := newSc.Left().(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Polynomial; received %T",
			newSc.Left(),
		)
	}

	// Verify that the new right hand side is a constant 3.14
	m2, ok := newSc.Right().(symbolic.K)
	if !ok {
		t.Errorf(
			"Expected newSc.RightHandSide to be a symbolic.K; received %T",
			newSc.Right(),
		)
	}

	if float64(m2) != 3.14 {
		t.Errorf(
			"Expected newSc.RightHandSide to be 3.14; received %v",
			m2,
		)
	}

	if newSc.ConstrSense() != sc.Sense {
		t.Errorf(
			"Expected newSc.Sense to be %v; received %v",
			sc.Sense,
			newSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_String1
Description:

	Tests the String() method of a scalar constraint. This test verifies
	that the method properly returns a string representation of the
	constraint.
	We will double check that the string contains
		- the left hand side's string representation
		- the right hand side's string representation
		- the sense of the constraint
*/
func TestScalarConstraint_String1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{x, c2, symbolic.SenseLessThanEqual}

	// Get string representation
	s := sc.String()

	// Verify that the string contains the left hand side
	if !strings.Contains(s, sc.Left().String()) {
		t.Errorf(
			"Expected s to contain %v; received %v",
			sc.Left().String(),
			s,
		)
	}

	// Verify that the string contains the right hand side
	if !strings.Contains(s, sc.Right().String()) {
		t.Errorf(
			"Expected s to contain %v; received %v",
			sc.Right().String(),
			s,
		)
	}

	// Verify that the string contains the sense
	if !strings.Contains(s, sc.ConstrSense().String()) {
		t.Errorf(
			"Expected s to contain %v; received %v",
			sc.ConstrSense().String(),
			s,
		)
	}
}

/*
TestScalarConstraint_ScaleBy1
Description:

	This tests that the ScalarConstraint.ScaleBy() method properly panics
	when the left hand side is not a valid monomial.
*/
func TestScalarConstraint_ScaleBy1(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: v2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.ScaleBy() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m1.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.ScaleBy() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.ScaleBy(2)

	t.Errorf(
		"Expected sc.ScaleBy() to panic; received nil",
	)
}

/*
TestScalarConstraint_ScaleBy2
Description:

	This tests that the ScalarConstraint.ScaleBy() method properly
	returns a new ScalarConstraint when the left hand side is a valid monomial.
	In this case, we scale by a negative number, which should lead to a new
	constraint with the opposite sense.
*/
func TestScalarConstraint_ScaleBy2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Scale
	newSc := sc.ScaleBy(-2)

	// Verify that the new left hand side is a monomial
	if _, ok := newSc.Left().(symbolic.Monomial); !ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Monomial; received %T",
			newSc.Left(),
		)
	}

	// Verify that the new right hand side is a constant 6.28
	m2, ok := newSc.Right().(symbolic.K)
	if !ok {
		t.Errorf(
			"Expected newSc.RightHandSide to be a symbolic.K; received %T",
			newSc.Right(),
		)
	}

	// Verify that the new right hand side is a constant 6.28
	if float64(m2) != -6.28 {
		t.Errorf(
			"Expected newSc.RightHandSide to be 6.28; received %v",
			m2,
		)
	}

	// Verify that the new constraint has the opposite sense
	if newSc.ConstrSense() != symbolic.SenseGreaterThanEqual {
		t.Errorf(
			"Expected newSc.Sense to be different from %v; received %v",
			sc.Sense,
			newSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_ScaleBy3
Description:

	This tests that the ScalarConstraint.ScaleBy() method properly
	returns a new ScalarConstraint when the left hand side is a valid monomial.
	In this case, we scale by a positive number, which should lead to a new
	constraint with the same sense.
*/
func TestScalarConstraint_ScaleBy3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Scale
	newSc := sc.ScaleBy(2)

	// Verify that the new left hand side is a monomial
	if _, ok := newSc.Left().(symbolic.Monomial); !ok {
		t.Errorf(
			"Expected newSc.LeftHandSide to be a symbolic.Monomial; received %T",
			newSc.Left(),
		)
	}

	// Verify that the new right hand side is a constant 6.28
	m2, ok := newSc.Right().(symbolic.K)
	if !ok {
		t.Errorf(
			"Expected newSc.RightHandSide to be a symbolic.K; received %T",
			newSc.Right(),
		)
	}

	// Verify that the new right hand side is a constant 6.28
	if float64(m2) != 6.28 {
		t.Errorf(
			"Expected newSc.RightHandSide to be 6.28; received %v",
			m2,
		)
	}

	// Verify that the new constraint has the same sense
	if newSc.ConstrSense() != sc.Sense {
		t.Errorf(
			"Expected newSc.Sense to be %v; received %v",
			sc.Sense,
			newSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_ScaleBy4
Description:

	This tests that the ScalarConstraint.ScaleBy() method properly
	flips the sign of a constraint with the SenseGreaterThanEqual sense
	when the scale factor is negative.
*/
func TestScalarConstraint_ScaleBy4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{x},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseGreaterThanEqual,
	}

	// Scale
	newSc := sc.ScaleBy(-2)

	// Verify that the new constraint has the opposite sense
	if newSc.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf(
			"Expected newSc.Sense to be different from %v; received %v",
			sc.Sense,
			newSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_Variables1
Description:

	Tests the Variables() method of a scalar constraint. This test verifies
	that the method properly returns a slice of all variables in the constraint.
	We will create a constraint with 2 variables (one in a monomial on the left
	and a constant on the right).
*/
func TestScalarConstraint_Variables1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Verify variables
	vars := sc.Variables()
	if len(vars) != 2 {
		t.Errorf(
			"Expected 2 variables; received %d",
			len(vars),
		)
	}
}

/*
TestScalarConstraint_Variables2
Description:

	Tests the Variables() method of a scalar constraint. This test verifies
	that the method properly panics when the left hand side is not a valid monomial.
*/
func TestScalarConstraint_Variables2(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: v2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.Variables() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m1.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.Variables() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.Variables()

	t.Errorf(
		"Expected sc.Variables() to panic; received nil",
	)
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied1
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly panics when the given
	scalar constraint is not valid (in this case, the left hand side is not
	a valid monomial).
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied1(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: v2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create a second constraint
	sc2 := v2.GreaterEq(symbolic.K(1))

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m1.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.ImpliesThisIsAlsoSatisfied() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.ImpliesThisIsAlsoSatisfied(sc2)

	t.Errorf(
		"Expected sc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
	)
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied2
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly panics when the constraint
	given as an argument is not valid (in this case, the left hand side is not
	a valid monomial).
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	c2 := symbolic.K(3.14)

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: c2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Create a second constraint
	m2 := symbolic.Monomial{
		Coefficient:     1,
		Exponents:       []int{1},
		VariableFactors: []symbolic.Variable{},
	}
	v2 := symbolic.NewVariable()
	sc2 := symbolic.ScalarConstraint{
		LeftHandSide:  m2,
		RightHandSide: v2,
		Sense:         symbolic.SenseGreaterThanEqual,
	}

	// Create the panic handling function
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected sc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
			)
		}

		rAsError := r.(error)
		expectedError := m2.Check()
		if rAsError.Error() != expectedError.Error() {
			t.Errorf(
				"Expected sc.ImpliesThisIsAlsoSatisfied() to panic with error \"%v\"; received \"%v\"",
				expectedError.Error(),
				rAsError.Error(),
			)
		}
	}()

	sc.ImpliesThisIsAlsoSatisfied(sc2)

	t.Errorf(
		"Expected sc.ImpliesThisIsAlsoSatisfied() to panic; received nil",
	)
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied3
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly returns false when the
	constraint does not imply the other constraint. In this case, we have
	x + y <= 3 and y >= 1. The first constraint does NOT imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	c2 := symbolic.K(3.0)

	// Create constraint
	sc := x.Plus(y).LessEq(c2)

	// Create a second constraint
	sc2 := y.GreaterEq(symbolic.K(1))

	// Verify that the first constraint does NOT imply the second
	if sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be false; received true",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied4
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly returns true when the
	constraint DOES imply the other constraint AND they are both single-variable
	constraints. In this case, we have
	x <= 3 and x <= 4. The first constraint DOES imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.LessEq(3.0)

	// Create a second constraint
	sc2 := x.LessEq(4.0)

	// Verify that the first constraint DOES imply the second
	if !sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be true; received false",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied5
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly returns true when the
	constraint DOES imply the other constraint AND they are both single-variable
	constraints. In this case, we have
	x >= 4 and x >= 3. The first constraint DOES imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.GreaterEq(4.0)

	// Create a second constraint
	sc2 := x.GreaterEq(3.0)

	// Verify that the first constraint DOES imply the second
	if !sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be true; received false",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied6
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly returns true when the
	constraint DOES imply the other constraint AND they are both single-variable
	constraints. In this case, we have
	x = 3 and x <= 4. The first constraint DOES imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied6(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Eq(3.0)

	// Create a second constraint
	sc2 := x.LessEq(4.0)

	// Verify that the first constraint DOES imply the second
	if !sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be true; received false",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied7
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly returns true when the
	constraint DOES imply the other constraint AND they are both single-variable
	constraints. In this case, we have
	x = 3 and x >= 2. The first constraint DOES imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied7(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Eq(3.0)

	// Create a second constraint
	sc2 := x.GreaterEq(2.0)

	// Verify that the first constraint DOES imply the second
	if !sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be true; received false",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied8
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly handles the case when both
	constraints contain negative coefficients on the same variable.
	In this case, we have:
		-2 x <= 4 and
		-2 x <= 5
	as the input constraints. The first constraint DOES imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied8(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Multiply(-2).LessEq(4.0)

	// Create a second constraint
	sc2 := x.Multiply(-2).LessEq(5.0)

	// Verify that the first constraint DOES imply the second
	if !sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be true; received false",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied9
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly handles the case when both
	constraints are SenseGreaterThanEqual and have positive coefficients
	on the same variable, but they do NOT imply each other.
	In this case, we have:
		2 x >= 4 and
		2 x >= 5
	as the input constraints. The first constraint does NOT imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied9(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Multiply(2).GreaterEq(4.0)

	// Create a second constraint
	sc2 := x.Multiply(2).GreaterEq(5.0)

	// Verify that the first constraint does NOT imply the second
	if sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be false; received true",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied10
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test verifies that the method properly handles the case when both
	cosntraints are SenseEqual and have positive coefficients
	on the same variable, but they do NOT imply each other.
	In this case, we have:
		2 x = 4 and
		2 x = 5
	as the input constraints. The first constraint does NOT imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied10(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Multiply(2).Eq(4.0)

	// Create a second constraint
	sc2 := x.Multiply(2).Eq(5.0)

	// Verify that the first constraint does NOT imply the second
	if sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be false; received true",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied11
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test attempts to catch a bug where the method would note that some
	constraints imply others that they should not actually imply.
	In this case, it seems like we have constraints where, if you compare
	the constants on the right hand side, then they might seem to imply one
	another: (1 <= 2), but when you consider the left hand side's
	coefficients, you see that they do NOT actually imply one another.
	In this case, we have:
		2 x <= 1 and
		10 x <= 2
	as the input constraints. The first constraint does NOT imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied11(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Multiply(2).LessEq(1.0)

	// Create a second constraint
	sc2 := x.Multiply(10).LessEq(2.0)

	// Verify that the first constraint does NOT imply the second
	if sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be false; received true",
		)
	}
}

/*
TestScalarConstraint_ImpliesThisIsAlsoSatisfied12
Description:

	Tests the ImpliesThisIsAlsoSatisfied() method of a scalar constraint.
	This test attempts to catch a bug where the method would note that some
	constraints imply others that they should not actually imply.
	In this case, it seems like we have constraints where, if you compare
	the constants on the right hand side, then they might seem to imply one
	another: (-2 >= -4), but when you consider the left hand side's
	coefficients, you see that they do NOT actually imply one another.
	In this case, we have:
		-2 x <= 2 and
		-10 x <= 4
	as the input constraints. The first constraint does NOT imply the second.
*/
func TestScalarConstraint_ImpliesThisIsAlsoSatisfied12(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	sc := x.Multiply(-2).LessEq(2.0)

	// Create a second constraint
	sc2 := x.Multiply(-10).LessEq(4.0)

	// Verify that the first constraint does NOT imply the second
	if sc.ImpliesThisIsAlsoSatisfied(sc2) {
		t.Errorf(
			"Expected sc.ImpliesThisIsAlsoSatisfied(sc2) to be false; received true",
		)
	}
}

/*
TestScalarConstraint_AsSimplifiedConstraint1
Description:

	This function tests the AsSimplifiedConstraint() method of a scalar constraint.
	We will create a scalar constraint with a monomial on the left hand side
	and a polynomial on the right hand side. The expected output is a new
	constraint with a polynomial on the left hand side and a constant on the
	right hand side.
*/
func TestScalarConstraint_AsSimplifiedConstraint1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     2,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x, y},
	}
	p2 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			{
				Coefficient:     1,
				Exponents:       []int{1},
				VariableFactors: []symbolic.Variable{x},
			},
			{
				Coefficient:     3,
				Exponents:       []int{1},
				VariableFactors: []symbolic.Variable{y},
			},
			{
				Coefficient:     4,
				Exponents:       []int{},
				VariableFactors: []symbolic.Variable{},
			},
		},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: p2,
		Sense:         symbolic.SenseLessThanEqual,
	}

	// Get simplified constraint
	simplifiedSc := sc.AsSimplifiedConstraint()

	// Verify that the left hand side is a polynomial
	if _, ok := simplifiedSc.Left().(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected simplifiedSc.LeftHandSide to be a symbolic.Polynomial; received %T",
			simplifiedSc.Left(),
		)
	}

	// Verify that the right hand side is a constant
	k, ok := simplifiedSc.Right().(symbolic.K)
	if !ok {
		t.Errorf(
			"Expected simplifiedSc.RightHandSide to be a symbolic.K; received %T",
			simplifiedSc.Right(),
		)
	}

	if float64(k) != 4 {
		t.Errorf(
			"Expected simplifiedSc.RightHandSide to be 4; received %v",
			k,
		)
	}

	// Verify that the sense is the same
	if simplifiedSc.ConstrSense() != sc.ConstrSense() {
		t.Errorf(
			"Expected simplifiedSc.Sense to be %v; received %v",
			sc.ConstrSense(),
			simplifiedSc.ConstrSense(),
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint1
Description:

	Tests the IsNonnegativityConstraint() method of a scalar constraint.
	This test verifies that the method correctly identifies a non-negativity
	constraint of the form: x >= 0.
*/
func TestScalarConstraint_IsNonnegativityConstraint1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := x.GreaterEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as a non-negativity constraint
	if !sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be true; received false",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint2
Description:

	This test verifies that the new method panics if the
	input ScalarConstraint is not well-defined.
*/
func TestScalarConstraint_IsNonnegativityConstraint2(t *testing.T) {
	// Setup
	x := symbolic.NewVariable()

	// Create a malformed monomial
	m1 := symbolic.Monomial{
		Coefficient:     2,
		Exponents:       []int{1, 1},
		VariableFactors: []symbolic.Variable{x},
	}

	// Create constraint
	sc := symbolic.ScalarConstraint{
		LeftHandSide:  m1,
		RightHandSide: symbolic.K(0),
		Sense:         symbolic.SenseGreaterThanEqual,
	}

	// Verify that the IsNonnegativityConstraint() method
	// panics
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected IsNonnegativityConstraint() to panic; received no panic",
			)
		}

		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected IsNonnegativityConstraint() to panic with an error; received %v",
				r,
			)
		}

		if err.Error() != m1.Check().Error() {
			t.Errorf(
				"Expected IsNonnegativityConstraint() to panic with error %v; received %v",
				m1.Check().Error(),
				err.Error(),
			)
		}
	}()

	sc.IsNonnegativityConstraint()
	t.Errorf("Expected IsNonnegativityConstraint() to panic; received no panic")
}

/*
TestScalarConstraint_IsNonnegativityConstraint3
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly
	returns false when the scalar constraint contains more than one variable.
*/
func TestScalarConstraint_IsNonnegativityConstraint3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Create constraint
	constr := x.Plus(y).GreaterEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as NOT a non-negativity constraint
	if sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be false; received true",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint4
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly identifies
	that a non-positivity constraint (i.e., x <= 0) is NOT a non-negativity constraint.
*/
func TestScalarConstraint_IsNonnegativityConstraint4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := x.LessEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as NOT a non-negativity constraint
	if sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be false; received true",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint5
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly identifies
	a constraint of the form 2 * x >= 0 IS ALSO A non-negativity constraint.
*/
func TestScalarConstraint_IsNonnegativityConstraint5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := symbolic.K(2).Multiply(x).GreaterEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as a non-negativity constraint
	if !sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be true; received false",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint6
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly
	identifies that a constraint of the form x >= 3 is NOT a non-negativity constraint.
*/
func TestScalarConstraint_IsNonnegativityConstraint6(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := x.GreaterEq(3)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as NOT a non-negativity constraint
	if sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be false; received true",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint7
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly identifies
	that a constraint of the form - x <= 0 IS a non-negativity constraint.
*/
func TestScalarConstraint_IsNonnegativityConstraint7(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := symbolic.K(-1).Multiply(x).LessEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as a non-negativity constraint
	if !sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be true; received false",
		)
	}
}

/*
TestScalarConstraint_IsNonnegativityConstraint8
Description:

	This test verifies that the IsNonnegativityConstraint() method correctly identifies
	that a constraint of the form x^2 >= 0 IS NOT a non-negativity constraint.
*/
func TestScalarConstraint_IsNonnegativityConstraint8(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create constraint
	constr := x.Power(2).GreaterEq(0)
	sc, ok := constr.(symbolic.ScalarConstraint)
	if !ok {
		t.Errorf(
			"Expected constr to be a symbolic.ScalarConstraint; received %T",
			constr,
		)
	}

	// Verify that the constraint is identified as NOT a non-negativity constraint
	if sc.IsNonnegativityConstraint() {
		t.Errorf(
			"Expected sc.IsNonnegativityConstraint() to be false; received true",
		)
	}
}
