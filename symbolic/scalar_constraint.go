package symbolic

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

	Describes whether a given scalar constraint is
	linear or not.
*/
func (sc ScalarConstraint) IsLinear() bool {
	return IsLinear(sc.RightHandSide) && IsLinear(sc.LeftHandSide)
}

/*
Simplify
Description:

	Moves all of the variables of the ScalarConstraint to its
	left hand side.
*/
func (sc ScalarConstraint) Simplify() ScalarConstraint {
	// Create LHS
	newLHS := sc.LeftHandSide

	// If RHS is a constant, then simply return the constraint
	if _, ok := sc.RightHandSide.(K); ok {
		return sc
	}

	// Algorithm
	newLHS = newLHS.Minus(sc.RightHandSide).Plus(
		sc.RightHandSide.Constant(),
	).(ScalarExpression) // This should be a scalar expression

	return ScalarConstraint{
		LeftHandSide:  newLHS,
		RightHandSide: K(sc.RightHandSide.Constant()),
		Sense:         sc.Sense,
	}

}

/*
ConstrSense
Description:

	Returns the sense of the constraint.
*/
func (sc ScalarConstraint) ConstrSense() ConstrSense {
	return sc.Sense
}

/*
Check
Description:

	Checks that the ScalarConstraint is valid.
*/
func (sc ScalarConstraint) Check() error {
	// Input Processing
	// Check that the left and right hand sides are well formed.
	err := sc.LeftHandSide.Check()
	if err != nil {
		return err
	}

	err = sc.RightHandSide.Check()
	if err != nil {
		return err
	}

	// Check that the sense is valid.
	err = sc.Sense.Check()
	if err != nil {
		return err
	}

	// All Checks Passed!
	return nil
}
