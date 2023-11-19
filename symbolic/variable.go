package symbolic

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type Variable struct {
	ID    uint64
	Lower float64
	Upper float64
	Vtype VarType
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the variable expression v.
*/
func (v Variable) Variables() []Variable {
	return []Variable{v}
}

// Vars returns a slice of the Var ids in the expression. For a variable, it
// always returns a singleton slice with the given variable ID.
func (v Variable) IDs() []uint64 {
	return []uint64{v.ID}
}

// Coeffs returns a slice of the coefficients in the expression. For a variable,
// it always returns a singleton slice containing the value one.
func (v Variable) Coeffs() []float64 {
	return []float64{1}
}

// Constant returns the constant additive value in the expression. For a
// variable, it always returns zero.
func (v Variable) Constant() float64 {
	return 0
}

// Plus adds the current expression to another and returns the resulting
// expression.
func (v Variable) Plus(e interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := v.Check()
	if err != nil {
		return v, err
	}

	err = CheckErrors(errors)
	if err != nil {
		return v, err
	}

	// Algorithm
	switch e.(type) {
	case K:
		eAsK := e.(K)

		// Organize vector variables
		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsK.Variables()...)),
		}

		// Return
		return ScalarLinearExpr{
			L: OnesVector(1),
			X: vv,
			C: float64(eAsK),
		}, nil
	case Variable:
		// Convert
		eAsV := e.(Variable)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsV.Variables()...)),
		}

		// Check to see if this is the same Variable or a different one
		if eAsV.ID == v.ID {
			return ScalarLinearExpr{
				X: vv,
				L: *mat.NewVecDense(1, []float64{2.0}),
				C: 0.0,
			}, nil
		} else {
			return ScalarLinearExpr{
				X: vv,
				L: OnesVector(2),
				C: 0.0,
			}, nil
		}
	case ScalarLinearExpr:
		// Convert
		eAsSLE := e.(ScalarLinearExpr)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsSLE.Variables()...)),
		}

		// Convert SLE to new form
		e2, _ := eAsSLE.RewriteInTermsOf(vv)
		vIndex, _ := FindInSlice(v, vv.Elements)
		e2.L.SetVec(vIndex, e2.L.AtVec(vIndex)+1.0)

		return e2, nil

	case ScalarQuadraticExpression:
		// Convert
		eAsQE := e.(ScalarQuadraticExpression)

		vv := VarVector{
			UniqueVars(append([]Variable{v}, eAsQE.Variables()...)),
		}

		// Convert QE to new form
		e2, _ := eAsQE.RewriteInTermsOf(vv)
		vIndex, _ := FindInSlice(v, vv.Elements)
		e2.L.SetVec(vIndex, e2.L.AtVec(vIndex)+1.0)

		return e2, nil

	default:
		return v, fmt.Errorf("There was an unexpected type (%T) given to Variable.Plus()!", e)
	}
}

//// Mult multiplies the current expression to another and returns the
//// resulting expression
//func (v Variable) Mult(m float64) (ScalarExpression, error) {
//	// Constants
//	// switch m.(type) {
//	// case float64:
//
//	vars := []Variable{v}
//	coeffs := []float64{m * v.Coeffs()[0]}
//
//	// Algorithm
//	newExpr := ScalarLinearExpr{
//		X: VarVector{vars},
//		L: *mat.NewVecDense(1, coeffs),
//		C: 0,
//	}
//	return newExpr, nil
//	// case *Variable:
//	// 	return nil
//	// }
//}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (v Variable) LessEq(rhsIn interface{}, errors ...error) (Constraint, error) {
	return v.Comparison(rhsIn, SenseLessThanEqual, errors...)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (v Variable) GreaterEq(rhsIn interface{}, errors ...error) (Constraint, error) {
	return v.Comparison(rhsIn, SenseGreaterThanEqual, errors...)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (v Variable) Eq(rhsIn interface{}, errors ...error) (Constraint, error) {
	return v.Comparison(rhsIn, SenseEqual, errors...)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := v.Comparison(expr1,SenseGreaterThanEqual)
*/
func (v Variable) Comparison(rhsIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	// Input Processing
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
	return ScalarConstraint{v, rhs, sense}, nil
}

/*
// ID returns the ID of the variable
func (v *Variable) ID() uint64 {
	return v.ID
}

// Lower returns the lower value limit of the variable
func (v *Variable) Lower() float64 {
	return v.Lower
}

// Upper returns the upper value limit of the variable
func (v *Variable) Upper() float64 {
	return v.Upper
}

// Type returns the type of variable (continuous, binary, integer, etc)
func (v *Variable) Type() VarType {
	return v.Vtype
}
*/

// VarType represents the type of the variable (continuous, binary,
// integer, etc) and uses Gurobi's encoding.
type VarType byte

// Multiple common variable types have been included as constants that conform
// to Gurobi's encoding.
const (
	Continuous VarType = 'C'
	Binary             = 'B'
	Integer            = 'I'
)

/*
UniqueVars
Description:

	This function creates a slice of unique variables from the slice given in
	varsIn
*/
func UniqueVars(varsIn []Variable) []Variable {
	// Constants

	// Algorithm
	var varsOut []Variable
	for _, v := range varsIn {
		if vIndex, _ := FindInSlice(v, varsOut); vIndex == -1 { // If v is not yet in varsOut, then add it
			varsOut = append(varsOut, v)
		}
	}

	return varsOut

}

/*
Multiply
Description:

	multiplies the current expression to another and returns the resulting expression
*/
func (v Variable) Multiply(val interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := v.Check()
	if err != nil {
		return v, err
	}

	err = CheckErrors(errors)
	if err != nil {
		return v, err
	}

	if IsExpression(val) {
		rightAsE, _ := ToExpression(val)
		err = CheckDimensionsInMultiplication(v, rightAsE)
		if err != nil {
			return v, err
		}
	}

	// Constants
	switch e := val.(type) {
	case float64:
		return v.Multiply(K(e))
	case K:
		// Algorithm
		return e.Multiply(v)
	case Variable:
		sqeOut := ScalarQuadraticExpression{
			X: VarVector{
				UniqueVars([]Variable{e, v}),
			},
			C: 0.0,
		}
		sqeOut.L = ZerosVector(sqeOut.X.Len())
		if e.ID == v.ID {
			sqeOut.Q = *mat.NewDense(1, 1, []float64{1.0})
		} else {
			sqeOut.Q = ZerosMatrix(2, 2)
			sqeOut.Q.Set(0, 1, 0.5)
			sqeOut.Q.Set(1, 0, 0.5)
		}
		return sqeOut, nil

	case ScalarLinearExpr:
		// Algorithm
		sqeOut := ScalarQuadraticExpression{
			X: VarVector{
				UniqueVars(append(e.X.Elements, v)),
			},
			C: 0.0,
		}
		sqeOut.Q = ZerosMatrix(sqeOut.X.Len(), sqeOut.X.Len())
		sqeOut.L = ZerosVector(sqeOut.X.Len())

		// Update Q
		vIndex, _ := FindInSlice(v, e.X.Elements)           // err should be nil
		vIndexInSQE, _ := FindInSlice(v, sqeOut.X.Elements) // err should be nil
		for xIndex := 0; xIndex < e.L.Len(); xIndex++ {
			// Check to make sure index is not the vIndex
			if vIndex == xIndex {
				// If v is in the original slice (i.e., we now need to represent v^2)

				sqeOut.Q.Set(vIndexInSQE, vIndexInSQE, e.L.AtVec(vIndex))
			} else {
				// If xIndex is not for v, then create off-diagonal elements
				xIndexInSQE, _ := FindInSlice(e.X.AtVec(xIndex), sqeOut.X.Elements)

				// Create a pair of off-diagonal elements
				sqeOut.Q.Set(vIndexInSQE, xIndexInSQE, e.L.AtVec(xIndex)*0.5)
				sqeOut.Q.Set(xIndexInSQE, vIndexInSQE, e.L.AtVec(xIndex)*0.5)

			}
		}

		// Update L
		sqeOut.L.SetVec(vIndexInSQE, e.C)

		return sqeOut, nil

	case ScalarQuadraticExpression:
		// Return error
		return ScalarQuadraticExpression{}, fmt.Errorf("Can not multiply Variable with ScalarQuadraticExpression. MatProInterface can not represent polynomials higher than degree 2.")

	case VectorLinearExpressionTranspose:
		return ScalarQuadraticExpression{}, fmt.Errorf(
			"cannot currently multiply a variable with a vector to create a quadratic expression; file an issue if you are interested in seeing a feature like this.",
		)

	default:
		return v, fmt.Errorf("Unexpected input to v.Multiply(): %T", val)
	}
}

/*
ToScalarLinearExpression
Description:

	Converting the variable into a scalar linear Expression.
*/
func (v Variable) ToScalarLinearExpression() ScalarLinearExpr {
	// Constants

	// Create components
	vars := []Variable{v}
	coeffs := []float64{v.Coeffs()[0]}

	// Create sle
	return ScalarLinearExpr{
		X: VarVector{vars},
		L: *mat.NewVecDense(1, coeffs),
		C: 0,
	}
}

/*
Dims
Description:

	Returns the dimension of the Variable object (should be scalar).
*/
func (v Variable) Dims() []int {
	return []int{1, 1}
}

/*
Check
Description:

	Checks whether or not the Variable has a sensible initialization.
*/
func (v Variable) Check() error {
	// Check that the lower bound is below is the upper bound
	if v.Lower > v.Upper {
		return fmt.Errorf(
			"lower bound (%v) of variable is above upper bound (%v).",
			v.Lower, v.Upper,
		)
	}

	// If nothing was thrown, then return nil!
	return nil
}

func (v Variable) Transpose() Expression {
	return v
}

/*
NewVariable
Description:
*/
func NewVariable(envs ...Environment) Variable {
	return NewContinuousVariable(envs...)
}

/*
NewContinuousVariable
Description:

	Creates a new continuous variable.
*/
func NewContinuousVariable(envs ...Environment) Variable {
	// Constants

	// Input Processing
	var currentEnv Environment
	switch len(envs) {
	case 1:
		currentEnv = envs[0]
	default:
		currentEnv = BackgroundEnvironment
	}

	// Get New Index
	nextIdx := len(currentEnv.Variables)

	return Variable{
		ID:    uint64(nextIdx),
		Lower: float64(-Infinity),
		Upper: float64(+Infinity),
		Vtype: Continuous,
	}

}

/*
NewContinuousVariable
Description:

	Creates a new binary variable.
*/
func NewBinaryVariable(envs ...Environment) Variable {
	// Constants

	// Input Processing
	var currentEnv Environment
	switch len(envs) {
	case 1:
		currentEnv = envs[0]
	default:
		currentEnv = BackgroundEnvironment
	}

	// Get New Index
	nextIdx := len(currentEnv.Variables)

	return Variable{
		ID:    uint64(nextIdx),
		Lower: 0.0,
		Upper: 1.0,
		Vtype: Binary,
	}

}
