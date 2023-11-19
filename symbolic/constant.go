package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
Integer constants representing commonly used numbers. Makes for better
readability
*/
const (
	Zero     = K(0)
	One      = K(1)
	Infinity = K(1e18)
)

// K is a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
type K float64

/*
Variables
Description:

	Shares all variables included in the expression that is K.
	It is a constant, so there are none.
*/
func (c K) Variables() []Variable {
	return []Variable{}
}

// Vars returns a slice of the Var ids in the expression. For constants,
// this is always nil
func (c K) IDs() []uint64 {
	return nil
}

// Coeffs returns a slice of the coefficients in the expression. For constants,
// this is always nil
func (c K) Coeffs() []float64 {
	return nil
}

// Constant returns the constant additive value in the expression. For
// constants, this is just the constants value
func (c K) Constant() float64 {
	return float64(c)
}

// Plus adds the current expression to another and returns the resulting
// expression
func (c K) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(c, rightAsE)
		if err != nil {
			return c, err
		}
	}

	// Switching based on input type
	switch right := rightIn.(type) {
	case K:
		return K(c.Constant() + right.Constant()), nil
	case Variable:
		return right.Plus(c)
	case ScalarLinearExpr:
		return right.Plus(c)
	case ScalarQuadraticExpression:
		return right.Plus(c) // Very compact, but potentially confusing to read?
	default:
		return c, fmt.Errorf("Unexpected type in K.Plus() for constant %v: %T", right, right)
	}
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (c K) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (c K) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (c K) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return c.Comparison(rightIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.
*/
func (c K) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// InputProcessing
	err := CheckErrors(errors)
	if err != nil {
		return ScalarConstraint{}, err
	}

	rhs, err := ToScalarExpression(rhsIn)
	if err != nil {
		return ScalarConstraint{}, err
	}

	// Constants

	// Algorithm
	return ScalarConstraint{c, rhs, sense}, nil
}

/*
Multiply
Description:

	This method multiplies the input constant by another expression.
*/
func (c K) Multiply(term1 interface{}, errors ...error) (Expression, error) {
	// Constants

	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return c, err
	}

	if IsExpression(term1) {
		// Check dimensions
		term1AsE, _ := ToExpression(term1)
		err = CheckDimensionsInMultiplication(c, term1AsE)
		if err != nil {
			return c, err
		}
	}

	// Algorithm
	switch right := term1.(type) {
	case float64:
		return c.Multiply(K(right))
	case K:
		return c * right, nil
	case Variable:
		// Algorithm
		term1AsSLE := right.ToScalarLinearExpression()

		return c.Multiply(term1AsSLE)
	case ScalarLinearExpr:
		// Scale all vectors and constants
		sleOut := right.Copy()
		sleOut.L.ScaleVec(float64(c), &sleOut.L)
		sleOut.C = right.C * float64(c)

		return sleOut, nil
	case ScalarQuadraticExpression:
		// Scale all matrices and constants
		var sqeOut ScalarQuadraticExpression
		sqeOut.Q.Scale(float64(c), &right.Q)
		sqeOut.L.ScaleVec(float64(c), &right.L)
		sqeOut.C = float64(c) * right.C

		return sqeOut, nil
	case KVector:
		var prod mat.VecDense = ZerosVector(right.Len())
		term1AsVecDense := mat.VecDense(right)

		prod.ScaleVec(float64(c), &term1AsVecDense)

		return KVector(prod), nil
	case KVectorTranspose:
		var prod mat.VecDense = ZerosVector(right.Len())
		term1AsVecDense := mat.VecDense(right)

		prod.ScaleVec(float64(c), &term1AsVecDense)

		return KVectorTranspose(prod), nil
	case VarVector:
		// VarVector is of unit length.
		return ScalarLinearExpr{
			L: OnesVector(1),
			X: right.Copy(),
			C: 0.0,
		}, nil
	case VarVectorTranspose:
		if right.Len() == 1 {
			rightTransposed := right.Transpose().(VarVector)
			prod := ScalarLinearExpr{
				L: OnesVector(1),
				X: rightTransposed.Copy(),
				C: 0.0,
			}
			prod.L.ScaleVec(float64(c), &prod.L)

			return prod, nil
		} else {
			var vleOut VectorLinearExpressionTranspose
			vleOut.X = right.Copy().Transpose().(VarVector)
			tempIdentity := Identity(right.Len()) // Is this needed?
			vleOut.L.Scale(float64(c), &tempIdentity)
			vleOut.C = ZerosVector(right.Len())

			return vleOut, nil
		}
	case VectorLinearExpr:
		var vleOut VectorLinearExpr
		vleOut.L.Scale(float64(c), &right.L)
		vleOut.C.ScaleVec(float64(c), &right.C)
		vleOut.X = right.X.Copy()

		return vleOut, nil
	case VectorLinearExpressionTranspose:
		var vletOut VectorLinearExpressionTranspose
		vletOut.L.Scale(float64(c), &right.L)
		vletOut.C.ScaleVec(float64(c), &right.C)
		vletOut.X = right.X.Copy()

		return vletOut, nil
	default:
		return K(0), fmt.Errorf("Unexpected type of term1 in the Multiply() method: %T (%v)", term1, term1)

	}
}

func (c K) Dims() []int {
	return []int{1, 1} // Signifies scalar
}

func (c K) Check() error {
	return nil
}

func (c K) Transpose() Expression {
	return c
}
