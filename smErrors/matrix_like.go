package smErrors

type MatrixLike interface {
	Dims() []int
}

/*
IsScalarExpression
Description:

	Determines whether or not an input object is a valid "ScalarExpression" according to Dims().
*/
func IsScalarExpression(e MatrixLike) bool {
	nRows, nCols := e.Dims()[0], e.Dims()[1]

	return (nRows == 1) && (nCols == 1)
}
