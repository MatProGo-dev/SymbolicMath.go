package symbolic_test

/*
matrix_constraint_test.go
Description:
	Tests the methods defined for the matrix constraint object.
*/

import (
	"fmt"
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
TestMatrixConstraint_Left1
Description:

	Tests that the Left() method returns the correct value.
	When the constraint is made of two constant matrices,
	the Left() method should return the left hand side of the constraint
	which will be an identity matrix.
*/
func TestMatrixConstraint_Left1(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	mcLeft := mc.Left()

	// Verify that mcLeft is a KMatrix type
	mcLeftAsKM, ok := mcLeft.(symbolic.KMatrix)
	if !ok {
		t.Errorf(
			"Expected mcLeft to be of type KMatrix; received %T",
			mcLeft,
		)
	}

	// Verify that each of mcLeftAsKM's elements match left's elements
	mcLeftAsD := mcLeftAsKM.ToDense()
	leftAsD := left.ToDense()
	for rIndex := 0; rIndex < 3; rIndex++ {
		for cIndex := 0; cIndex < 3; cIndex++ {
			if mcLeftAsD.At(rIndex, cIndex) != leftAsD.At(rIndex, cIndex) {
				t.Errorf(
					"Expected mcLeftAsKM.At(%v, %v) to be %v; received %v",
					rIndex, cIndex,
					leftAsD.At(rIndex, cIndex),
					mcLeftAsD.At(rIndex, cIndex),
				)
			}
		}
	}
}

/*
TestMatrixConstraint_Right1
Description:

	Tests that the Right() method returns the correct value.
	When the constraint is made of two constant matrices,
	the Right() method should return the left hand side of the constraint
	which will be a zeros matrix.
*/
func TestMatrixConstraint_Right1(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	mcRight := mc.Right()

	// Verify that mcRight is a KMatrix type
	mcRightAsKM, ok := mcRight.(symbolic.KMatrix)
	if !ok {
		t.Errorf(
			"Expected mcRight to be of type KMatrix; received %T",
			mcRight,
		)
	}

	// Verify that each of mcRightAsKM's elements match right's elements
	mcRightAsD := mcRightAsKM.ToDense()
	rightAsD := right.ToDense()
	for rIndex := 0; rIndex < 3; rIndex++ {
		for cIndex := 0; cIndex < 3; cIndex++ {
			if mcRightAsD.At(rIndex, cIndex) != rightAsD.At(rIndex, cIndex) {
				t.Errorf(
					"Expected mcRightAsKM.At(%v, %v) to be %v; received %v",
					rIndex, cIndex,
					rightAsD.At(rIndex, cIndex),
					mcRightAsD.At(rIndex, cIndex),
				)
			}
		}
	}
}

/*
TestMatrixConstraint_Check1
Description:

	Tests that the Check() method returns an error when the
	left and right hand sides have different dimensions.
	(Left side has dimension 3x3, right side has dimension 3x2)
*/
func TestMatrixConstraint_Check1(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 2))
	expectedError := fmt.Errorf(
		"there are a different number of columns in the left (%v) and right (%v) sides of the constraint!",
		left.Dims()[1],
		right.Dims()[1],
	)

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	if mc.Check().Error() != expectedError.Error() {
		t.Errorf(
			"Expected mc.Check() to return \"%v\"; received %v",
			expectedError,
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Check2
Description:

	Tests that the Check() method returns an error when the
	left side is not a well-formed matrix expression.
*/
func TestMatrixConstraint_Check2(t *testing.T) {
	// Constants
	left := symbolic.MonomialMatrix{}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	if mc.Check().Error() != left.Check().Error() {
		t.Errorf(
			"Expected mc.Check() to return \"%v\"; received %v",
			left.Check().Error(),
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Check3
Description:

	Tests that the Check() method returns an error when the
	right side is not a well-formed matrix expression.
*/
func TestMatrixConstraint_Check3(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.MonomialMatrix{}

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	if mc.Check().Error() != right.Check().Error() {
		t.Errorf(
			"Expected mc.Check() to return \"%v\"; received %v",
			right.Check().Error(),
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Check4
Description:

	Tests that the Check() method returns an error when the
	dimensions of the left and right do not match (there are
	more rows in left, 4, than there are in the right, 3).
*/
func TestMatrixConstraint_Check4(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(4))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))
	expectedError := fmt.Errorf(
		"there are a different number of rows in the left (%v) and right (%v) sides of the constraint!",
		left.Dims()[0],
		right.Dims()[0],
	)

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	if mc.Check().Error() != expectedError.Error() {
		t.Errorf(
			"Expected mc.Check() to return \"%v\"; received %v",
			expectedError,
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Check5
Description:

	Tests that the Check() method returns an error when the
	sense is not a well-formed sense.
*/
func TestMatrixConstraint_Check5(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))
	var sense symbolic.ConstrSense = 12

	// Test
	mc := symbolic.MatrixConstraint{left, right, sense}

	if mc.Check().Error() != sense.Check().Error() {
		t.Errorf(
			"Expected mc.Check() to return \"%v\"; received %v",
			sense.Check().Error(),
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Check6
Description:

	Tests that the Check() method returns nil when the
	left and right hand sides are well-formed and the
	sense is well-formed.
*/
func TestMatrixConstraint_Check6(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))

	// Test
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	if mc.Check() != nil {
		t.Errorf(
			"Expected mc.Check() to return nil; received %v",
			mc.Check(),
		)
	}
}

/*
TestMatrixConstraint_Dims1
Description:

	Tests that the Dims() method returns the correct value
	on a well-formed matrix constraint each of shape (3,4).
	Dims() should be [3,4].
*/
func TestMatrixConstraint_Dims1(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	dims := mc.Dims()

	if len(dims) != 2 {
		t.Errorf(
			"Expected len(dims) to be 2; received %v",
			len(dims),
		)
	}

	// Check rows
	if dims[0] != 3 {
		t.Errorf(
			"Expected dims[0] to be 3; received %v",
			dims[0],
		)
	}

	// Check columns
	if dims[1] != 4 {
		t.Errorf(
			"Expected dims[1] to be 3; received %v",
			dims[1],
		)
	}
}

/*
TestMatrixConstraint_At1
Description:

	Tests that the At() method returns the correct value
	on a well-formed matrix constraint each of shape (3,4).
*/
func TestMatrixConstraint_At1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	left := symbolic.MonomialMatrix{
		{v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial()},
	}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(1, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	constrElt := mc.At(0, 2)

	_, tf := constrElt.LeftHandSide.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"Expected sConstr.LeftHandSide to be of type Monomial; received %T",
			constrElt.LeftHandSide,
		)
	}

	_, tf = constrElt.RightHandSide.(symbolic.K)
	if !tf {
		t.Errorf(
			"Expected sConstr.RightHandSide to be of type K; received %T",
			constrElt.RightHandSide,
		)
	}

}

/*
TestMatrixConstraint_At2
Description:

	Tests that the At() method panics when the row index is out of bounds.
*/
func TestMatrixConstraint_At2(t *testing.T) {
	// Constants
	left := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := smErrors.InvalidMatrixIndexError{
		RowIndex:   3,
		ColIndex:   0,
		Expression: mc,
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mc.At(3, 0) to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(smErrors.InvalidMatrixIndexError)
		if !ok {
			t.Errorf(
				"Expected mc.At(3, 0) to panic with type InvalidMatrixIndexError; received %T",
				r,
			)
		}

		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mc.At(3, 0) to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	mc.At(3, 0)
}

/*
TestMatrixConstraint_At3
Description:

	Tests that the At() method panics when one of the matrices
	is not well-formed.
*/
func TestMatrixConstraint_At3(t *testing.T) {
	// Constants
	left := symbolic.MonomialMatrix{}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := left.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mc.At(0, 0) to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mc.At(0, 0) to panic with type error; received %T",
				r,
			)
		}

		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mc.At(0, 0) to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	mc.At(0, 0)
}

/*
TestMatrixConstraint_Substitute1
Description:

	Tests that the Substitute() method properly panics
	when the left hand side is not a well-formed matrix expression.
	(In this case, the left hand side is a monomial matrix that is not well-formed)
*/
func TestMatrixConstraint_Substitute1(t *testing.T) {
	// Constants
	left := symbolic.MonomialMatrix{}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := left.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mc.Substitute() to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mc.Substitute() to panic with type error; received %T",
				r,
			)
		}

		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mc.Substitute() to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	mc.Substitute(symbolic.NewVariable(), symbolic.NewVariable())

	t.Errorf(
		"Expected mc.Substitute() to panic; did not panic",
	)
}

/*
TestMatrixConstraint_Substitute2
Description:

	Tests that the Substitute() method properly returns a
	well-formed matrix constraint when the left and right hand sides
	are well-formed and the sense is well-formed.
	The left hand side is a monomial matrix and the right hand side
	is a constant matrix.
	After substitution, the left hand side should be a polynomial matrix
	and the right hand side should be a constant matrix.
*/
func TestMatrixConstraint_Substitute2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	left := symbolic.MonomialMatrix{
		{v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial()},
	}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(1, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	mcSubstituted := mc.Substitute(v1, v2.Plus(v2).(symbolic.ScalarExpression))

	// Verify that mcSubstituted is a MatrixConstraint type
	mcSubstitutedAsMC, ok := mcSubstituted.(symbolic.MatrixConstraint)
	if !ok {
		t.Errorf(
			"Expected mcSubstituted to be of type MatrixConstraint; received %T",
			mcSubstituted,
		)
	}

	// Verify that the left hand side is a polynomial matrix
	if _, ok := mcSubstitutedAsMC.LeftHandSide.(symbolic.PolynomialMatrix); !ok {
		t.Errorf(
			"Expected mcSubstituted.LeftHandSide to be of type PolynomialMatrix; received %T",
			mcSubstitutedAsMC.LeftHandSide,
		)
	}

	// Verify that the right hand side is a constant matrix
	if _, ok := mcSubstitutedAsMC.RightHandSide.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected mcSubstituted.RightHandSide to be of type KMatrix; received %T",
			mcSubstitutedAsMC.RightHandSide,
		)
	}
}

/*
TestMatrixConstraint_SubstituteAccordingTo1
Description:

	Tests that the SubstituteAccordingTo() method properly panics
	when the left hand side is not a well-formed matrix expression.
	(In this case, the left hand side is a monomial matrix that is not well-formed)
*/
func TestMatrixConstraint_SubstituteAccordingTo1(t *testing.T) {
	// Constants
	left := symbolic.MonomialMatrix{}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := left.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic with type error; received %T",
				r,
			)
		}

		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	mc.SubstituteAccordingTo(map[symbolic.Variable]symbolic.Expression{})

	t.Errorf(
		"Expected mc.SubstituteAccordingTo() to panic; did not panic",
	)
}

/*
TestMatrixConstraint_SubstituteAccordingTo2
Description:

	Tests that the SubstituteAccordingTo() method properly returns a
	well-formed matrix constraint when the left and right hand sides
	are well-formed and the sense is well-formed.
	The left hand side is a monomial matrix and the right hand side
	is a constant matrix.
	After substitution, the left hand side should be a polynomial matrix
	and the right hand side should be a constant matrix.
*/
func TestMatrixConstraint_SubstituteAccordingTo2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	left := symbolic.MonomialMatrix{
		{v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial(), v1.ToMonomial()},
	}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(1, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	// Test
	mcSubstituted := mc.SubstituteAccordingTo(
		map[symbolic.Variable]symbolic.Expression{
			v1: v2.Plus(v2).(symbolic.ScalarExpression),
		},
	)

	// Verify that mcSubstituted is a MatrixConstraint type
	mcSubstitutedAsMC, ok := mcSubstituted.(symbolic.MatrixConstraint)
	if !ok {
		t.Errorf(
			"Expected mcSubstituted to be of type MatrixConstraint; received %T",
			mcSubstituted,
		)
	}

	// Verify that the left hand side is a polynomial matrix
	if _, ok := mcSubstitutedAsMC.LeftHandSide.(symbolic.PolynomialMatrix); !ok {
		t.Errorf(
			"Expected mcSubstituted.LeftHandSide to be of type PolynomialMatrix; received %T",
			mcSubstitutedAsMC.LeftHandSide,
		)
	}

	// Verify that the right hand side is a constant matrix
	if _, ok := mcSubstitutedAsMC.RightHandSide.(symbolic.KMatrix); !ok {
		t.Errorf(
			"Expected mcSubstituted.RightHandSide to be of type KMatrix; received %T",
			mcSubstitutedAsMC.RightHandSide,
		)
	}
}

/*
TestMatrixConstraint_AsSimplifiedConstraint1
Description:

	This test verifies that AsSimplifiedConstraint() properly panics
	whenever a malformed MatrixConstraint is used to call it.
*/
func TestMatrixConstraint_AsSimplifiedConstraint1(t *testing.T) {
	// Create constraint
	left := symbolic.MonomialMatrix{}
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 4))
	mc := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := left.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic with type error; received %T",
				r,
			)
		}

		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mc.SubstituteAccordingTo() to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	// Call AsSimplifiedConstraint()
	mc.AsSimplifiedConstraint()

	t.Errorf(
		"Expected mc.AsSimplifiedConstraint() to panic; did not panic",
	)

}

/*
TestMatrixConstraint_Variables1
Description:

	Verifies that the Variables() method works as expected.
	We verify that the method properly returns 9 unique variables when
	we create a constraint between a 3x3 matrix of variables and a
	3x3 matrix of constants.
*/
func TestMatrixConstraint_Variables1(t *testing.T) {
	// Constants
	vm1 := symbolic.NewVariableMatrix(3, 3)
	km2 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 3))

	mConstr := symbolic.MatrixConstraint{vm1, km2, symbolic.SenseLessThanEqual}

	// Test
	vars := mConstr.Variables()

	if len(vars) != 9 {
		t.Errorf(
			"Expected len(vars) to be 9; received %v",
			len(vars),
		)
	}
	// Check that each variable is unique
	varsMap := make(map[symbolic.Variable]bool)
	for _, v := range vars {
		if _, ok := varsMap[v]; ok {
			t.Errorf(
				"Expected all variables to be unique; received duplicate variable %v",
				v,
			)
		}
		varsMap[v] = true
	}
}

/*
TestMatrixConstraint_ImpliesThisIsAlsoSatisfied1
Description:

	Verifies that the ImpliesThisIsAlsoSatisfied method works as expected
	when the input constraint is a scalar constraint that matches one of
	the constraints in the original matrix constraint (i.e., the scalar constraint
	is just the (0,1)-th constraint in the original matrix constraint).
*/
func TestMatrixConstraint_ImpliesThisIsAlsoSatisfied1(t *testing.T) {
	// Constants
	vm1 := symbolic.NewVariableMatrix(3, 3)
	km2 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 3))

	mConstr := symbolic.MatrixConstraint{vm1, km2, symbolic.SenseLessThanEqual}

	// Extract the (0,1)-th constraint from mConstr
	scalarConstraint := mConstr.At(0, 1)

	// Test
	if !mConstr.ImpliesThisIsAlsoSatisfied(scalarConstraint) {
		t.Errorf(
			"Expected mConstr.ImpliesThisIsAlsoSatisfied(scalarConstraint) to be true; received false",
		)
	}
}

/*
TestMatrixConstraint_ImpliesThisIsAlsoSatisfied2
Description:

	Verifies that the ImpliesThisIsAlsoSatisfied method correctly panics
	when the input matrix constraint is malformed.
*/
func TestMatrixConstraint_ImpliesThisIsAlsoSatisfied2(t *testing.T) {
	// Constants
	vm1 := symbolic.NewVariableMatrix(3, 2)
	km2 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 3))

	// Create matrix constraint
	mConstr := symbolic.MatrixConstraint{vm1, km2, symbolic.SenseLessThanEqual}

	// Create normal matrix constraint to use as input
	left := symbolic.DenseToKMatrix(symbolic.Identity(3))
	right := symbolic.DenseToKMatrix(symbolic.ZerosMatrix(3, 3))
	normalMC := symbolic.MatrixConstraint{left, right, symbolic.SenseLessThanEqual}

	expectedError := mConstr.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(normalMC) to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(normalMC) to panic with type error; received %T",
				r,
			)
		}
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(normalMC) to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	// Call ImpliesThisIsAlsoSatisfied
	mConstr.ImpliesThisIsAlsoSatisfied(normalMC)

	t.Errorf(
		"Expected mConstr.ImpliesThisIsAlsoSatisfied(normalMC) to panic; did not panic",
	)
}

/*
TestConstraint_ImpliesThisIsAlsoSatisfied3
Description:

	Verifies that the ImpliesThisIsAlsoSatisfied method correctly panics
	when the receiver is well-defined but the input constraint is not well-formed.
*/
func TestConstraint_ImpliesThisIsAlsoSatisfied3(t *testing.T) {
	// Constants
	vm1 := symbolic.NewVariableMatrix(3, 3)
	km2 := symbolic.DenseToKMatrix(symbolic.OnesMatrix(3, 3))

	// Create matrix constraint
	mConstr := symbolic.MatrixConstraint{vm1, km2, symbolic.SenseLessThanEqual}

	// Create malformed scalar constraint to use as input
	scLeft := symbolic.Monomial{
		Exponents: []int{1},
	}
	scRight := symbolic.K(5)
	sc := symbolic.ScalarConstraint{scLeft, scRight, symbolic.SenseEqual}
	expectedError := sc.Check()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(sc) to panic; did not panic",
			)
		}

		// Check that the error is the expected error
		err, ok := r.(error)
		if !ok {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(sc) to panic with type error; received %T",
				r,
			)
		}
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected mConstr.ImpliesThisIsAlsoSatisfied(sc) to panic with error \"%v\"; received \"%v\"",
				expectedError,
				err,
			)
		}

	}()

	// Call ImpliesThisIsAlsoSatisfied
	mConstr.ImpliesThisIsAlsoSatisfied(sc)

	t.Errorf(
		"Expected mConstr.ImpliesThisIsAlsoSatisfied(sc) to panic; did not panic",
	)
}
