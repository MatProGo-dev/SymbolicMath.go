package symbolic

import (
	"fmt"

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

	A = newLHS.LinearCoeff(wrt...)

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
	C = newLHS.LinearCoeff(wrt...)

	// Create d
	newRHS := sc.Right().(ScalarExpression).Constant() - sc.Left().(ScalarExpression).Constant()
	d = newRHS

	// Return
	return C, d
}

/*
Substitute
Description:

	Substitutes the variable vIn with the scalar expression seIn in the
	given scalar constraint.
*/
func (sc ScalarConstraint) Substitute(vIn Variable, seIn ScalarExpression) Constraint {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left hand side
	newLHS := sc.LeftHandSide.Substitute(vIn, seIn).(ScalarExpression)

	// Substitute the variable in the right hand side
	newRHS := sc.RightHandSide.Substitute(vIn, seIn).(ScalarExpression)

	// Return the new constraint
	return ScalarConstraint{
		LeftHandSide:  newLHS,
		RightHandSide: newRHS,
		Sense:         sc.Sense,
	}
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the map with the corresponding expressions
	in the given scalar constraint.
*/
func (sc ScalarConstraint) SubstituteAccordingTo(subMap map[Variable]Expression) Constraint {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left hand side
	newLHS := sc.LeftHandSide.SubstituteAccordingTo(subMap).(ScalarExpression)

	// Substitute the variable in the right hand side
	newRHS := sc.RightHandSide.SubstituteAccordingTo(subMap).(ScalarExpression)

	// Return the new constraint
	return ScalarConstraint{
		LeftHandSide:  newLHS,
		RightHandSide: newRHS,
		Sense:         sc.Sense,
	}
}

/*
String
Description:

	Returns a string representation of the scalar constraint.
*/
func (sc ScalarConstraint) String() string {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Create the string representation
	return sc.LeftHandSide.String() + " " + sc.Sense.String() + " " + sc.RightHandSide.String()
}

/*
AsSimplifiedConstraint
Description:

	Simplifies the constraint by moving all variables to the left hand side and the constants to the right.
*/
func (sc ScalarConstraint) AsSimplifiedConstraint() Constraint {
	return sc.Simplify()
}

func (sc ScalarConstraint) Variables() []Variable {
	return VariablesInThisConstraint(sc)
}

func (sc ScalarConstraint) ScaleBy(factor float64) Constraint {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Scale the left hand side
	newLHS := sc.LeftHandSide.Multiply(factor).(ScalarExpression)

	// Scale the right hand side
	newRHS := sc.RightHandSide.Multiply(factor).(ScalarExpression)

	// If the factor is negative, then flip the sense of the constraint
	newSense := sc.Sense
	if factor < 0 {
		if sc.Sense == SenseLessThanEqual {
			newSense = SenseGreaterThanEqual
		} else if sc.Sense == SenseGreaterThanEqual {
			newSense = SenseLessThanEqual
		}
	}

	// Return the new constraint
	return ScalarConstraint{
		LeftHandSide:  newLHS,
		RightHandSide: newRHS,
		Sense:         newSense,
	}
}

/*
ImpliesThisIsAlsoSatisfied
Description:

	Returns true if this constraint implies that the other constraint is also satisfied.
*/
func (sc ScalarConstraint) ImpliesThisIsAlsoSatisfied(other Constraint) bool {
	// Check that the constraint is well formed.
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	// Check that the other constraint is well formed.
	err = other.Check()
	if err != nil {
		panic(err)
	}

	// Simplify both constraints
	sc = sc.Simplify()

	switch otherC := other.(type) {
	case ScalarConstraint:
		otherC = otherC.Simplify()

		// Naive implication check:
		// 1. Both constraints contain only 1 variable AND it is the same variable. Then, simply check the bounds.
		containsOneVar := len(sc.Variables()) == 1 && len(otherC.Variables()) == 1
		scAndOtherShareSameVar := len(UnionOfVariables(sc.Variables(), otherC.Variables())) == 1

		if containsOneVar && scAndOtherShareSameVar {
			// Get the coefficient of the single variable
			scCoeffVector := sc.LeftHandSide.LinearCoeff(sc.Variables())
			scCoeff := scCoeffVector.AtVec(0)
			otherCCoeffVector := otherC.LeftHandSide.LinearCoeff(otherC.Variables())
			otherCCoeff := otherCCoeffVector.AtVec(0)

			// If the coefficient of scCoeff is < 0,
			// then flip the signs of both sides of the constraint
			if scCoeff < 0 {
				sc = sc.ScaleBy(-1).(ScalarConstraint)
			}

			if otherCCoeff < 0 {
				otherC = otherC.ScaleBy(-1).(ScalarConstraint)
			}

			// The implication holds if all of the following are true:
			// 1. The sense of sc and otherC are either the same (or one is equality)
			// 2. The bounds of the constraint with the LessThanEqual or GreaterThanEqual sense are within the bounds of the other constraint.
			sensesAreCompatible := sc.Sense == otherC.Sense ||
				sc.Sense == SenseEqual ||
				otherC.Sense == SenseEqual

			if !sensesAreCompatible {
				return false
			}

			switch sc.Sense {
			case SenseLessThanEqual:
				// Check the senses of otherC
				switch otherC.Sense {
				case SenseLessThanEqual:
					// Both are <=
					// Then the implication holds if the upper bound of sc is <= the upper bound of otherC
					return sc.RightHandSide.Constant() <= otherC.RightHandSide.Constant()
				default:
					// sc is <= and otherC is either >= or ==
					// Then the implication holds if the upper bound of sc is <= the lower bound of otherC
					return false
				}
			case SenseGreaterThanEqual:
				// Check the senses of otherC
				switch otherC.Sense {
				case SenseGreaterThanEqual:
					// Both are >=
					// Then the implication holds if the lower bound of sc is >= the lower bound of otherC
					return sc.RightHandSide.Constant() >= otherC.RightHandSide.Constant()
				default:
					// sc is >= and otherC is either <= or ==
					// Then the implication holds if the lower bound of sc is >= the upper bound of otherC
					return false
				}
			case SenseEqual:
				// Check the senses of otherC
				switch otherC.Sense {
				case SenseEqual:
					// Both are ==
					// Then the implication holds if the bounds are equal
					return sc.RightHandSide.Constant() == otherC.RightHandSide.Constant()
				case SenseLessThanEqual:
					// sc is == and otherC is <=
					// Then the implication holds if the bound of sc is <= the upper bound of otherC
					return sc.RightHandSide.Constant() <= otherC.RightHandSide.Constant()
				case SenseGreaterThanEqual:
					// sc is == and otherC is >=
					// Then the implication holds if the bound of sc is >= the lower bound of otherC
					return sc.RightHandSide.Constant() >= otherC.RightHandSide.Constant()
				}
			default:
				panic("unreachable code reached in ScalarConstraint.ImpliesThisIsAlsoSatisfied")
			}
		}
	case VectorConstraint, MatrixConstraint:
		// TODO: Implement more advanced implication checks.
		return false
	default:
		// Other types of constraints are not currently supported.
		panic(
			fmt.Errorf("implication checking between ScalarConstraint and %T is not currently supported", other),
		)
	}

	return false
}

/*
IsNonnegativityConstraint
Description:

	Checks to see if the constraint is of the form:
	- x >= 0, or
	- 0 <= x
*/
func (sc ScalarConstraint) IsNonnegativityConstraint() bool {
	// Setup
	err := sc.Check()
	if err != nil {
		panic(err)
	}

	simplified := sc.AsSimplifiedConstraint().(ScalarConstraint)

	// Check to see if constraint contains more than 1 variable
	if len(simplified.Variables()) != 1 {
		return false
	}

	// Otherwise, the sense is SenseGreaterThanEqual, and this is a non-negativity
	// constraint if:
	// - LHS is Variable-Like
	// - RHS is Zero

	lhsIsVariableLike := false

	// LHS Is Variable Like if:
	// - It is a variable
	// - It is a monomial with a positive coefficient
	simplifiedAsPL, tf := simplified.LeftHandSide.(PolynomialLikeScalar)
	if !tf {
		return false // If lhs is not polynomial like, then return false.
	}

	lhsIsVariableLike = simplifiedAsPL.Degree() == 1

	if !lhsIsVariableLike {
		return false // If lhs is still not variable like, then return false.
	}

	// Check to see if rhs is zero
	rhsAsK := simplified.RightHandSide.(K)
	rhsIsZero := float64(rhsAsK) == 0

	if !rhsIsZero {
		return false
	}

	// Finally, the constraint is non-negativie if:
	// - LHS has positive coefficient AND sense is GreaterThanEqual
	// - LHS has negative coefficient AND sense is LessThanEqual
	coeffs := simplified.LeftHandSide.LinearCoeff()

	return (coeffs.AtVec(0) > 0 && simplified.Sense == SenseGreaterThanEqual) ||
		(coeffs.AtVec(0) < 0 && simplified.Sense == SenseLessThanEqual)
}
