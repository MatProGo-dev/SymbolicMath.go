package symbolic

import "fmt"

// ScalarConstraint represnts a linear constraint of the form x <= y, x >= y, or
// x == y. ScalarConstraint uses a left and right hand side expressions along with a
// constraint sense (<=, >=, ==) to represent a generalized linear constraint
type ScalarConstraint struct {
	LeftHandSide  ScalarExpression
	RightHandSide ScalarExpression
	Sense         ConstrSense
}

func (sc ScalarConstraint) Left() Expression {
	return sc.LeftHandSide
}

func (sc ScalarConstraint) Right() Expression {
	return sc.RightHandSide
}

/*
IsLinear
Description:

	Describes whether or not a given linear constraint is
	linear or not.
*/
func (sc ScalarConstraint) IsLinear() bool {
	return sc.LeftHandSide.IsLinear() && sc.RightHandSide.IsLinear()
}

/*
Simplify
Description:

	Moves all of the variables of the ScalarConstraint to its
	left hand side.
*/
func (sc ScalarConstraint) Simplify() (ScalarConstraint, error) {
	// Create LHS
	newLHS := sc.LeftHandSide

	// Algorithm
	switch right := sc.RightHandSide.(type) {
	case K:
		return sc, nil
	case Variable:
		newLHS := newLHS.Plus(right.Multiply(-1.0))
		newLHSAsSE, _ := ToScalarExpression(newLHS)

		return ScalarConstraint{
			LeftHandSide:  newLHSAsSE,
			RightHandSide: K(0),
			Sense:         sc.Sense,
		}, nil
	default:
		return sc, fmt.Errorf("unexpected type of right hand side: %T", right)
	}

}
