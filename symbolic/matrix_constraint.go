package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
matrix_constraint.go
Description:
	Functions related to the matrix constraint object.
*/

type MatrixConstraint struct {
	LeftHandSide  MatrixExpression
	RightHandSide MatrixExpression
	Sense         ConstrSense
}

func (mc MatrixConstraint) Left() Expression {
	return mc.LeftHandSide
}

func (mc MatrixConstraint) Right() Expression {
	return mc.RightHandSide
}

/*
ConstrSense
Description:

	Returns the sense of the constraint.
*/
func (mc MatrixConstraint) ConstrSense() ConstrSense {
	return mc.Sense
}

/*
Check
Description:

	Verifies that:
	- The left and right hand sides have matching dimensions,
	- The sense is valid, (i.e., it is from the set of allowable senses defined in ConstrSense.
*/
func (mc MatrixConstraint) Check() error {
	// Input Processing
	// Check that the left and right hand sides are well formed.
	err := mc.LeftHandSide.Check()
	if err != nil {
		return err
	}

	err = mc.RightHandSide.Check()
	if err != nil {
		return err
	}

	// Save the dimensions of the left and right hand sides.
	leftDims := mc.LeftHandSide.Dims()
	rightDims := mc.RightHandSide.Dims()

	// Check that the dimensions of the left and right hand sides are the same.
	if len(leftDims) != len(rightDims) {
		return fmt.Errorf("left and right hand sides have different dimensions")
	}

	if leftDims[0] != rightDims[0] {
		return fmt.Errorf(
			"there are a different number of rows in the left (%v) and right (%v) sides of the constraint!",
			leftDims[0],
			rightDims[0],
		)
	}

	if leftDims[1] != rightDims[1] {
		return fmt.Errorf(
			"there are a different number of columns in the left (%v) and right (%v) sides of the constraint!",
			leftDims[1],
			rightDims[1],
		)
	}

	// Check that the sense is valid.
	err = mc.Sense.Check()
	if err != nil {
		return err
	}

	// All checks passed
	return nil
}

/*
Dims
Description:

	The dimension of the matrix constraint (ideally this should be the same as the dimensions
	of the left and right hand sides).
*/
func (mc MatrixConstraint) Dims() []int {
	err := mc.Check()
	if err != nil {
		panic(err)
	}

	// Dimensions of right and left should be the same.
	return mc.LeftHandSide.Dims()

}

/*
AtVec
Description:

	Retrieves the constraint formed by one element of the "vector" constraint.
*/
func (mc MatrixConstraint) At(ii, jj int) ScalarConstraint {
	// Input Processing
	err := mc.Check()
	if err != nil {
		panic(err)
	}

	// Check to see whether or not the index is valid.
	err = smErrors.CheckIndexOnMatrix(ii, jj, mc)
	if err != nil {
		panic(err)
	}

	// Algorithm
	lhsAtIIJJ := mc.LeftHandSide.At(ii, jj)
	rhsAtIIJJ := mc.RightHandSide.At(ii, jj)

	return ScalarConstraint{lhsAtIIJJ, rhsAtIIJJ, mc.Sense}
}

/*
IsLinear
Description:

	Describes whether a given matrix constraint is
	linear or not.
*/
func (mc MatrixConstraint) IsLinear() bool {
	return IsLinear(mc.RightHandSide) && IsLinear(mc.LeftHandSide)
}

/*
Substitute
Description:

	Substitutes the variable vIn with the scalar expression seIn
*/
func (mc MatrixConstraint) Substitute(vIn Variable, seIn ScalarExpression) Constraint {
	// Check that the constraint is well formed.
	err := mc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left hand side
	newLHS := mc.LeftHandSide.Substitute(vIn, seIn).(MatrixExpression)

	// Substitute the variable in the right hand side
	newRHS := mc.RightHandSide.Substitute(vIn, seIn).(MatrixExpression)

	return MatrixConstraint{newLHS, newRHS, mc.Sense}
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the map with the corresponding expressions
	in the given scalar constraint.
*/
func (mc MatrixConstraint) SubstituteAccordingTo(subMap map[Variable]Expression) Constraint {
	// Check that the constraint is well formed.
	err := mc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left hand side
	newLHS := mc.LeftHandSide.SubstituteAccordingTo(subMap).(MatrixExpression)

	// Substitute the variable in the right hand side
	newRHS := mc.RightHandSide.SubstituteAccordingTo(subMap).(MatrixExpression)

	return MatrixConstraint{newLHS, newRHS, mc.Sense}
}
