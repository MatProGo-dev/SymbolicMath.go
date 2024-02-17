package symbolic

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
