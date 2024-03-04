package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
vector_constraint.go
Description:

*/

type VectorConstraint struct {
	LeftHandSide  VectorExpression
	RightHandSide VectorExpression
	Sense         ConstrSense
}

/*
Dims
Description:

	The dimension of the vector constraint (ideally this should be the same as the dimensions
	of the left and right hand sides).
*/
func (vc VectorConstraint) Dims() []int {
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Dimensions of right and left should be the same.
	return vc.LeftHandSide.Dims()

}

/*
AtVec
Description:

	Retrieves the constraint formed by one element of the "vector" constraint.
*/
func (vc VectorConstraint) AtVec(i int) ScalarConstraint {
	// Input Processing
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Check to see whether or not the index is valid.
	err = smErrors.CheckIndexOnVector(i, vc)
	if err != nil {
		panic(err)
	}

	// Algorithm
	lhsAtI := vc.LeftHandSide.AtVec(i)
	rhsAtI := vc.RightHandSide.AtVec(i)

	return ScalarConstraint{lhsAtI, rhsAtI, vc.Sense}
}

/*
Check
Description:

	Checks that the VectorConstraint is valid.
*/
func (vc VectorConstraint) Check() error {
	// Constants

	// Input Processing
	// Check that the left and right hand sides are well formed.
	err := vc.LeftHandSide.Check()
	if err != nil {
		return err
	}

	err = vc.RightHandSide.Check()
	if err != nil {
		return err
	}

	// Check the sense
	err = vc.Sense.Check()
	if err != nil {
		return err
	}

	// Check dimensions of left and right hand sides.
	if vc.LeftHandSide.Len() != vc.RightHandSide.Len() {
		return fmt.Errorf(
			"Left hand side has dimension %v, but right hand side has dimension %v!",
			vc.LeftHandSide.Len(),
			vc.RightHandSide.Len(),
		)
	}

	// All Checks Passed!
	return nil
}

func (vc VectorConstraint) Left() Expression {
	return vc.LeftHandSide
}

func (vc VectorConstraint) Right() Expression {
	return vc.RightHandSide
}

/*
ConstrSense
Description:

	Returns the sense of the constraint.
*/
func (vc VectorConstraint) ConstrSense() ConstrSense {
	return vc.Sense
}

/*
IsLinear
Description:

	Describes whether a given vector constraint is
	linear or not.
*/
func (vc VectorConstraint) IsLinear() bool {
	return IsLinear(vc.RightHandSide) && IsLinear(vc.LeftHandSide)
}
