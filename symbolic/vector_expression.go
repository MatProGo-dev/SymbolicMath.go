package symbolic

/*
vector_expression.go
Description:
	An improvement/successor to the scalar expr interface.
*/

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
VectorExpression
Description:

	This interface represents any expression written in terms of a
	vector of represents a linear general expression of the form
		c0 * x0 + c1 * x1 + ... + cn * xn + k where ci are coefficients and xi are
	variables and k is a constant. This is a base interface that is implemented
	by single variables, constants, and general linear expressions.
*/
type VectorExpression interface {
	// Check returns an error if the expression is not valid
	Check() error

	// Variables returns the number of variables in the expression.
	Variables() []Variable

	//// Coeffs returns a slice of the coefficients in the expression
	//LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.VecDense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}) Expression

	// Mult multiplies the current expression with another and returns the
	// resulting expression
	Multiply(e interface{}) Expression

	// LessEq returns a less than or equal to (<=) constraint between the
	// current expression and another
	LessEq(rhs interface{}) Constraint

	// GreaterEq returns a greater than or equal to (>=) constraint between the
	// current expression and another
	GreaterEq(rhs interface{}) Constraint

	// Comparison
	// Returns a constraint with respect to the sense (senseIn) between the
	// current expression and another.
	Comparison(rhs interface{}, sense ConstrSense) Constraint

	// Eq returns an equality (==) constraint between the current expression
	// and another
	Eq(rhs interface{}) Constraint

	// Len returns the length of the vector expression.
	Len() int

	//AtVec returns the expression at a given index
	AtVec(idx int) ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression

	// Dims returns the dimensions of the given expression
	Dims() []int

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn Variable) Expression
}

///*
//NewVectorExpression
//Description:
//
//	NewExpr returns a new expression with a single additive constant value, c,
//	and no variables. Creating an expression like sum := NewVectorExpr(0) is useful
//	for creating new empty expressions that you can perform operatotions on later
//*/
//func NewVectorExpression(c mat.VecDense) VectorLinearExpr {
//	return VectorLinearExpr{C: c}
//}

//func (e VectorExpression) getVarsPtr() *uint64 {
//
//	if e.NumVars() > 0 {
//		return &e.IDs()[0]
//	}
//
//	return nil
//}
//
//func (e VectorExpression) getCoeffsPtr() *float64 {
//	if e.NumVars() > 0 {
//		return &e.Coeffs()[0]
//	}
//
//	return nil
//}

/*
IsVectorExpression
Description:

	Determines whether or not an input object is a valid "VectorExpression" according to MatProInterface.
*/
func IsVectorExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case mat.VecDense:
		return true
	case KVector:
		return true
	case VariableVector:
		return true
	default:
		return false

	}
}

/*
ToVectorExpression
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ToVectorExpression(e interface{}) (VectorExpression, error) {
	// Input Processing
	if !IsVectorExpression(e) {
		return KVector(OnesVector(1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a VectorExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case KVector:
		return e2, nil
	case mat.VecDense:
		return KVector(e2), nil
	case VariableVector:
		return e2, nil
	default:
		return KVector(OnesVector(1)), fmt.Errorf(
			"unexpected vector expression conversion requested for type %T!",
			e,
		)
	}
}
