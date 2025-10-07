package symbolic

import "fmt"

/*
constraint.go
Description:
	Defines an interface that we are meant to use with the ScalarContraint and VectorConstraint
	objects.
*/

type Constraint interface {
	Left() Expression
	Right() Expression
	ConstrSense() ConstrSense
	Check() error
	IsLinear() bool
	Substitute(vIn Variable, eIn ScalarExpression) Constraint
	SubstituteAccordingTo(subMap map[Variable]Expression) Constraint

	// Variables
	// Returns a slice of all the variables in the constraint.
	Variables() []Variable

	// ImpliesThisIsAlsoSatisfied
	// Returns true if this constraint implies that the other constraint is also satisfied.
	ImpliesThisIsAlsoSatisfied(other Constraint) bool

	// AsSimplifiedConstraint
	// Simplifies the constraint by moving all variables to the left hand side and the constants to the right.
	AsSimplifiedConstraint() Constraint

	// // IsPositivityConstraint
	// // Returns true if the constraint defines a non-negativity constraint of the form x >= 0
	// IsNonnegativityConstraint() bool
}

func IsConstraint(c interface{}) bool {
	switch c.(type) {
	case ScalarConstraint:
		return true
	case *ScalarConstraint:
		return true
	case VectorConstraint:
		return true
	case *VectorConstraint:
		return true
	case MatrixConstraint:
		return true
	case *MatrixConstraint:
		return true
	}

	// Return false, if the constraint is not a scalar or vector constraint.
	return false
}

/*
Variables
Description:

	Returns a slice of all the variables in the constraint.
*/
func VariablesInThisConstraint(c Constraint) []Variable {
	// Setup
	varsMap := make(map[Variable]bool)

	// Input check
	err := c.Check()
	if err != nil {
		panic(err)
	}

	// Get variables from the left hand side
	for _, v := range c.Left().Variables() {
		varsMap[v] = true
	}

	// Get variables from the right hand side
	for _, v := range c.Right().Variables() {
		varsMap[v] = true
	}

	// Convert the map to a slice
	vars := make([]Variable, 0, len(varsMap))
	for v := range varsMap {
		vars = append(vars, v)
	}

	return vars
}

/*
CompileConstraintsIntoScalarConstraints
Description:

	This method analyzes all constraints in an OptimizationProblem and converts them all
	into scalar constraints.
*/
func CompileConstraintsIntoScalarConstraints(constraints []Constraint) []ScalarConstraint {
	// Setup
	var out []ScalarConstraint

	// Iterate through all constraints
	for _, constraint := range constraints {
		// Switch statement based on the type of the constraint
		switch concreteConstraint := constraint.(type) {
		case ScalarConstraint:
			out = append(out, concreteConstraint)
		case VectorConstraint:
			for ii := 0; ii < concreteConstraint.Len(); ii++ {
				out = append(out, concreteConstraint.AtVec(ii))
			}
		case MatrixConstraint:
			dims := concreteConstraint.Dims()
			for rowIdx := 0; rowIdx < dims[0]; rowIdx++ {
				for colIdx := 0; colIdx < dims[1]; colIdx++ {
					out = append(out, concreteConstraint.At(rowIdx, colIdx))
				}
			}
		default:
			panic(
				fmt.Errorf(
					"The received constraint type (%T) is not supported by ExtractScalarConstraints!",
					constraint,
				),
			)
		}
	}

	return out
}
