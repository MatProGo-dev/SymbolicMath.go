package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
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

/*
LinearInequalityConstraintRepresentation
Description:

	Returns the linear constraint representation of the scalar constraint.
	Returns a tuple of the form (A, b) where A is a vector and b is a constant such that:
	A.Dot(x) <= b
*/
func (sc ScalarConstraint) LinearInequalityConstraintRepresentation(wrt ...[]Variable) (A mat.VecDense, b float64) {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Check that the constraint is linear.
	if !sc.IsLinear() {
		if !IsLinear(sc.LeftHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearInequalityConstraintRepresentation",
				Expression: sc.LeftHandSide,
			})
		}

		if !IsLinear(sc.RightHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearInequalityConstraintRepresentation",
				Expression: sc.RightHandSide,
			})
		}
	}

	// Create A
	newLHS := sc.Left().(ScalarExpression)
	newLHS = newLHS.Minus(sc.Right()).(ScalarExpression)

	A = newLHS.LinearCoeff()

	if sc.Sense == SenseGreaterThanEqual {
		A.ScaleVec(-1, &A)
	}

	// Create b
	newRHS := sc.Right().(ScalarExpression).Constant() - sc.Left().(ScalarExpression).Constant()
	b = newRHS

	if sc.Sense == SenseGreaterThanEqual {
		b = -b
	}

	// Return the tuple
	return A, b
}

/*
LinearEqualityConstraintRepresentation
Description:

	Returns the linear constraint representation of the scalar constraint.
	Returns a tuple of the form (C, d) where C is a vector and d is a constant such that:
	C.Dot(x) == d
*/
func (sc ScalarConstraint) LinearEqualityConstraintRepresentation(wrt ...[]Variable) (C mat.VecDense, d float64) {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Check that the constraint is linear.
	if !sc.IsLinear() {
		if !IsLinear(sc.LeftHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearEqualityConstraintRepresentation",
				Expression: sc.LeftHandSide,
			})
		}

		if !IsLinear(sc.RightHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearEqualityConstraintRepresentation",
				Expression: sc.RightHandSide,
			})
		}
	}

	// Check that the sense is equality.
	if sc.Sense != SenseEqual {
		panic(
			smErrors.EqualityConstraintRequiredError{
				Operation: "LinearEqualityConstraintRepresentation",
			},
		)
	}

	// Create C
	newLHS := sc.Left().(ScalarExpression)
	newLHS = newLHS.Minus(sc.Right()).(ScalarExpression)
	C = newLHS.LinearCoeff()

	// Create d
	newRHS := sc.Right().(ScalarExpression).Constant() - sc.Left().(ScalarExpression).Constant()
	d = newRHS

	// Return
	return C, d
}
