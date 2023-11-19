package symbolic

/*
Expression
Description:

	This interface should be implemented by and ScalarExpression and VectorExpression
*/
type Expression interface {
	// NumVars returns the number of variables in the expression
	Variables() []Variable

	// Vars returns a slice of the Var ids in the expression
	IDs() []uint64

	// Dims returns a slice describing the true dimensions of a given expression (scalar, vector, or matrix)
	Dims() []int

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}, errors ...error) (Expression, error)

	// Multiply multiplies the current expression to another and returns the
	// resulting expression
	Multiply(c interface{}, errors ...error) (Expression, error)

	// Transpose transposes the given expression
	Transpose() Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rightIn interface{}, errors ...error) (Constraint, error)

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rightIn interface{}, errors ...error) (Constraint, error)

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rightIn interface{}, errors ...error) (Constraint, error)

	// Comparison
	Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error)
}

func NumVariables(e Expression) int {
	return len(e.Variables())
}
