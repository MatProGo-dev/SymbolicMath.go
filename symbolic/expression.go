package symbolic

import "fmt"

/*
Expression
Description:

	This interface should be implemented by and ScalarExpression and VectorExpression
*/
type Expression interface {
	// NumVars returns the number of variables in the expression
	Variables() []Variable

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

/*
NumVariables
Description:

	The number of distinct variables.
*/
func NumVariables(e Expression) int {
	return len(e.Variables())
}

/*
VariableIDs
Description:

	Returns a list of ids associated with each variable.
*/
func VariableIDs(e Expression) []uint64 {
	vSlice := e.Variables()

	var idSlice []uint64
	for _, v := range vSlice {
		idSlice = append(idSlice, v.ID)
	}

	return idSlice
}

/*
IsExpression
Description:

	Tests whether or not the input variable is one of the expression types.
*/
func IsExpression(e interface{}) bool {
	return IsScalarExpression(e) || IsVectorExpression(e)
}

func ToExpression(e interface{}) (Expression, error) {
	switch {
	case IsScalarExpression(e):
		return ToScalarExpression(e)
	case IsVectorExpression(e):
		return ToVectorExpression(e)
	default:
		return K(Infinity), fmt.Errorf("the input expression is not recognized as a scalar or vector expression.")
	}
}
