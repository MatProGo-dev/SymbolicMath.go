package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

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
	// Check that both sides are polynomial like
	if !IsPolynomialLike(sc.LeftHandSide) {
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "ScalarConstraint.IsLinear",
				Input:        sc.LeftHandSide,
			},
		)
	}

	if !IsPolynomialLike(sc.RightHandSide) {
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "ScalarConstraint.IsLinear",
				Input:        sc.RightHandSide,
			},
		)
	}

	// Convert left and right hand sides to polynomial like
	leftAsPLS, _ := ToPolynomialLikeScalar(sc.LeftHandSide)
	rightAsPLS, _ := ToPolynomialLikeScalar(sc.RightHandSide)

	return IsLinear(leftAsPLS) && IsLinear(rightAsPLS)
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
