package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
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
	err = smErrors.CheckDimensionsInComparison(vc.Left(), vc.Right(), vc.ConstrSense().String())
	if err != nil {
		return err
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

/*
LinearInequalityConstraintRepresentation
Description:

	Returns the linear constraint representation of the scalar constraint.
	Returns a tuple of the form (A, b) where A is a vector and b is a constant such that:
	A.Dot(x) <= b
*/
func (vc VectorConstraint) LinearInequalityConstraintRepresentation(wrt ...[]Variable) (A mat.Dense, b mat.VecDense) {
	// Check that the constraint is well formed.
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Check that the constraint is linear.
	if !vc.IsLinear() {
		if !IsLinear(vc.LeftHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearInequalityConstraintRepresentation",
				Expression: vc.LeftHandSide,
			})
		}

		if !IsLinear(vc.RightHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearInequalityConstraintRepresentation",
				Expression: vc.RightHandSide,
			})
		}
	}

	// Check that the sense is inequality.
	if vc.Sense == SenseEqual {
		panic(
			smErrors.InequalityConstraintRequiredError{
				Operation: "LinearInequalityConstraintRepresentation",
			},
		)
	}

	// Create A
	newLHS := vc.Left().(PolynomialLikeVector)
	rhsWithoutConst := vc.Right().(PolynomialLikeVector)
	rhsWithoutConst = rhsWithoutConst.Minus(rhsWithoutConst.Constant()).(PolynomialLikeVector)
	newLHS = newLHS.Minus(rhsWithoutConst).(PolynomialLikeVector)

	A = newLHS.LinearCoeff(wrt...)

	if vc.Sense == SenseGreaterThanEqual {
		A.Scale(-1, &A)
	}

	// Create b
	N := vc.Left().(VectorExpression).Len()
	var newRHS *mat.VecDense = mat.NewVecDense(N, make([]float64, N))
	rightConst := vc.Right().(VectorExpression).Constant()
	leftConst := vc.Left().(VectorExpression).Constant()

	newRHS.SubVec(&rightConst, &leftConst)
	b = *newRHS

	if vc.Sense == SenseGreaterThanEqual {
		b.ScaleVec(-1, &b)
	}

	// Return the tuple
	return A, b
}

/*
LinearEqualityConstraintRepresentation
Description:

	Returns the representation of the constraint as a linear equality constraint.
	Returns a tuple of the form (C, d) where C is a matrix and d is a vector such that:
	C*x = d
*/
func (vc VectorConstraint) LinearEqualityConstraintRepresentation(wrt ...[]Variable) (C mat.Dense, d mat.VecDense) {
	// Check that the constraint is well formed.
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Check that the constraint is linear.
	if !vc.IsLinear() {
		if !IsLinear(vc.LeftHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearEqualityConstraintRepresentation",
				Expression: vc.LeftHandSide,
			})
		}

		if !IsLinear(vc.RightHandSide) {
			panic(smErrors.LinearExpressionRequiredError{
				Operation:  "LinearEqualityConstraintRepresentation",
				Expression: vc.RightHandSide,
			})
		}
	}

	// Check that the sense is equality.
	if vc.Sense != SenseEqual {
		panic(
			smErrors.EqualityConstraintRequiredError{
				Operation: "LinearEqualityConstraintRepresentation",
			},
		)
	}

	// Create C
	newLHS := vc.Left().(PolynomialLikeVector)
	rhsWithoutConst := vc.Right().(PolynomialLikeVector)
	rhsWithoutConst = rhsWithoutConst.Minus(rhsWithoutConst.Constant()).(PolynomialLikeVector)
	newLHS = newLHS.Minus(rhsWithoutConst).(PolynomialLikeVector)

	C = newLHS.LinearCoeff(wrt...)

	// Create d
	N := vc.Left().(VectorExpression).Len()
	var newRHS *mat.VecDense = mat.NewVecDense(N, make([]float64, N))
	rightConst := vc.Right().(VectorExpression).Constant()
	leftConst := vc.Left().(VectorExpression).Constant()
	newRHS.SubVec(&rightConst, &leftConst)
	d = *newRHS

	// Return the tuple
	return C, d
}

/*
Substitute
Description:

	Substitutes the variable vIn with the scalar expression seIn in the vector constraint.
*/
func (vc VectorConstraint) Substitute(vIn Variable, seIn ScalarExpression) Constraint {
	// Check that the constraint is well formed.
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left and right hand sides.
	newLHS := vc.LeftHandSide.Substitute(vIn, seIn).(VectorExpression)
	newRHS := vc.RightHandSide.Substitute(vIn, seIn).(VectorExpression)

	// Return the new constraint
	return VectorConstraint{newLHS, newRHS, vc.Sense}
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the map with the corresponding expressions
*/
func (vc VectorConstraint) SubstituteAccordingTo(subMap map[Variable]Expression) Constraint {
	// Check that the constraint is well formed.
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Substitute the variable in the left and right hand sides.
	newLHS := vc.LeftHandSide.SubstituteAccordingTo(subMap).(VectorExpression)
	newRHS := vc.RightHandSide.SubstituteAccordingTo(subMap).(VectorExpression)

	// Return the new constraint
	return VectorConstraint{newLHS, newRHS, vc.Sense}
}

/*
Len
Description:

	Returns the length of the vector constraint.
*/
func (vc VectorConstraint) Len() int {
	// Check that the constraint is well formed.
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Return the length of the vector constraint.
	return vc.LeftHandSide.Len()
}

/*
AsSimplifiedConstraint
Description:

	Simplifies the constraint by moving all variables to the left hand side and the constants to the right.
*/
func (vc VectorConstraint) AsSimplifiedConstraint() Constraint {
	// Input Checking
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Create Left Hand side of all of the expressions
	var newLHS Expression = vc.LeftHandSide.Minus(vc.LeftHandSide.Constant())
	newLHS = newLHS.Minus(
		vc.RightHandSide.Minus(vc.RightHandSide.Constant()),
	)

	// Create Right Hand Side of only constants
	var newRHS Expression = VecDenseToKVector(vc.RightHandSide.Constant()).Minus(
		vc.LeftHandSide.Constant(),
	)

	// Return new constraint
	return VectorConstraint{
		LeftHandSide:  newLHS.(VectorExpression),
		RightHandSide: newRHS.(VectorExpression),
		Sense:         vc.Sense,
	}
}

/*
Variables
Description:

	Returns a slice of all the variables in the constraint.
*/
func (vc VectorConstraint) Variables() []Variable {
	return VariablesInThisConstraint(vc)
}

/*
ImpliesThisIsAlsoSatisfied
Description:

	Returns true if this constraint implies that the other constraint is also satisfied.
*/
func (vc VectorConstraint) ImpliesThisIsAlsoSatisfied(other Constraint) bool {
	// Input Processing
	err := vc.Check()
	if err != nil {
		panic(err)
	}

	// Check the other constraint
	err = other.Check()
	if err != nil {
		panic(err)
	}

	switch otherC := other.(type) {
	case ScalarConstraint:
		// Continue
		// Naive implication check:
		// 1. One of the scalar constraints produces the correct implications.
		for i := 0; i < vc.Len(); i++ {
			if vc.AtVec(i).ImpliesThisIsAlsoSatisfied(otherC) {
				return true
			}
		}
	case VectorConstraint:
		// TODO: Implement more advanced implication checks.
		return false
	case MatrixConstraint:
		// TODO: Implement more advanced implication checks.
		return false
	default:
		// Other types of constraints are not currently supported.
		panic(
			fmt.Errorf("implication checking between VectorConstraint and %T is not currently supported", other),
		)
	}

	return false
}
