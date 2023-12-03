package symbolic

/*
matrix_constraint.go
Description:
	Functions related to the matrix constraint object.
*/

type MatrixConstraint struct {
	LeftHandSide  MatrixExpression
	RightHandSide MatrixExpression
	Sense         ConstrSense
}

func (mc MatrixConstraint) Left() Expression {
	return mc.LeftHandSide
}

func (mc MatrixConstraint) Right() Expression {
	return mc.RightHandSide
}
