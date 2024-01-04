package symbolic_test

/*
matrix_constraint_test.go
Description:
	Tests the methods defined for the matrix constraint object.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"testing"
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
	left := symbolic.KMatrix(symbolic.Identity(3))
	right := symbolic.KMatrix(symbolic.ZerosMatrix(3, 3))

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
	mcLeftAsD := mat.Dense(mcLeftAsKM)
	leftAsD := mat.Dense(left)
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
	left := symbolic.KMatrix(symbolic.Identity(3))
	right := symbolic.KMatrix(symbolic.ZerosMatrix(3, 3))

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
	mcRightAsD := mat.Dense(mcRightAsKM)
	rightAsD := mat.Dense(right)
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
