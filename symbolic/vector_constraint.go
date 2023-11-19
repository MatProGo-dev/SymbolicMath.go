package symbolic

import "fmt"

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
AtVec
Description:

	Retrieves the constraint formed by one element of the "vector" constraint.
*/
func (vc VectorConstraint) AtVec(i int) (ScalarConstraint, error) {
	// Input Processing
	if vc.Check() != nil {
		return ScalarConstraint{}, vc.Check()
	}

	vcLen := vc.LeftHandSide.Len()
	if i >= vcLen {
		return ScalarConstraint{},
			fmt.Errorf(
				"Cannot extract VectorConstraint element at %v; VectorConstraint has length %v.",
				i, vcLen,
			)
	}

	// Algorithm
	lhsAtI := vc.LeftHandSide.AtVec(i)
	rhsAtI := vc.RightHandSide.AtVec(i)

	return ScalarConstraint{lhsAtI, rhsAtI, vc.Sense}, nil
}

/*
Check
Description:

	Checks that the VectorConstraint is valid.
*/
func (vc VectorConstraint) Check() error {
	// Constants

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
