package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

// Var represnts a variable in a optimization problem. The variable is
// identified with an uint64.
type Variable struct {
	ID    uint64
	Lower float64
	Upper float64
	Type  VarType
	Name  string
}

/*
Variables
Description:

	This function returns a slice containing all unique variables in the variable expression v.
*/
func (v Variable) Variables() []Variable {
	return []Variable{v}
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
func (v Variable) Plus(rightIn interface{}) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	switch right := rightIn.(type) {
	case float64:
		return v.Plus(K(right))
	case K:
		return Polynomial{
			Monomials: []Monomial{
				v.ToMonomial(),
				right.ToMonomial(),
			},
		}
	case Variable:
		if v.ID == right.ID {
			return Polynomial{
				Monomials: []Monomial{
					Monomial{
						Coefficient:     2.0,
						VariableFactors: []Variable{v},
						Exponents:       []int{1},
					},
				},
			}
		} else {
			return Polynomial{
				Monomials: []Monomial{
					v.ToMonomial(),
					right.ToMonomial(),
				},
			}
		}
	case Monomial:
		return right.Plus(v)
	case Polynomial:
		return right.Plus(v)
	}

	panic(
		fmt.Errorf("there input %v has unexpected type %T given to Variable.Plus()!", rightIn, rightIn),
	)
}

// LessEq returns a less than or equal to (<=) constraint between the
// current expression and another
func (v Variable) LessEq(rhsIn interface{}) Constraint {
	return v.Comparison(rhsIn, SenseLessThanEqual)
}

// GreaterEq returns a greater than or equal to (>=) constraint between the
// current expression and another
func (v Variable) GreaterEq(rhsIn interface{}) Constraint {
	return v.Comparison(rhsIn, SenseGreaterThanEqual)
}

// Eq returns an equality (==) constraint between the current expression
// and another
func (v Variable) Eq(rhsIn interface{}) Constraint {
	return v.Comparison(rhsIn, SenseEqual)
}

/*
Comparison
Description:

	This method compares the receiver with expression rhs in the sense provided by sense.

Usage:

	constr, err := v.Comparison(expr1,SenseGreaterThanEqual)
*/
func (v Variable) Comparison(rhsIn interface{}, sense ConstrSense) Constraint {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rhsIn) {
		rhsAsE, _ := ToExpression(rhsIn)
		err = rhsAsE.Check()
		if err != nil {
			panic(err)
		}

		// No need to check dimensions here, as the comparison is
		// a scalar and thus valid for any dimension of rhsIn
	}

	// Algorithm
	switch rhs := rhsIn.(type) {
	case float64:
		// Use version of comparison for K
		return v.Comparison(K(rhs), sense)
	case K:
		// Create a new constraint
		return ScalarConstraint{v, rhs, sense}
	case Variable:
		// Create a new constraint
		return ScalarConstraint{v, rhs, sense}
	case Monomial:
		// Create a new constraint
		return ScalarConstraint{v, rhs, sense}
	case Polynomial:
		// Create a new constraint
		return ScalarConstraint{v, rhs, sense}
	}

	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Variable.Comparison",
			Input:        rhsIn,
		},
	)
}

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
func (v Variable) Multiply(rightIn interface{}) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := CheckDimensionsInMultiplication(v, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Constants
	switch right := rightIn.(type) {
	case float64:
		return v.Multiply(K(right))
	case K:
		// Create a new monomial
		monomialOut := Monomial{
			Coefficient:     float64(right),
			VariableFactors: []Variable{v},
			Exponents:       []int{1},
		}
		return monomialOut
	case Variable:
		var monomialOut Monomial
		if right.ID == v.ID {
			monomialOut = Monomial{
				Coefficient:     1.0,
				VariableFactors: []Variable{v},
				Exponents:       []int{2},
			}
		} else {
			monomialOut = Monomial{
				Coefficient:     1.0,
				VariableFactors: []Variable{v, right},
				Exponents:       []int{1, 1},
			}
		}
		return monomialOut
	}

	// Unrecornized response is a panic
	panic(
		fmt.Errorf("Unexpected input to v.Multiply(): %T", rightIn),
	)
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
	if v.Lower >= v.Upper {
		return fmt.Errorf(
			"lower bound (%v) of variable must be less than upper bound (%v).",
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
func NewVariable(envs ...*Environment) Variable {
	return NewContinuousVariable(envs...)
}

/*
NewContinuousVariable
Description:

	Creates a new continuous variable.
*/
func NewContinuousVariable(envs ...*Environment) Variable {
	// Constants

	// Input Processing
	var currentEnv = &BackgroundEnvironment
	switch len(envs) {
	case 1:
		currentEnv = envs[0]
	}

	// Get New Index
	nextIdx := len(currentEnv.Variables)

	// Create variable
	variableOut := Variable{
		ID:    uint64(nextIdx),
		Lower: float64(-Infinity),
		Upper: float64(+Infinity),
		Type:  Continuous,
		Name:  fmt.Sprintf("x_%v", nextIdx),
	}

	// Update environment
	currentEnv.Variables = append(currentEnv.Variables, variableOut)

	return variableOut

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

	// Get New Variable Object and add it to environment
	variableOut := Variable{
		ID:    uint64(nextIdx),
		Lower: 0.0,
		Upper: 1.0,
		Type:  Binary,
		Name:  fmt.Sprintf("x_%v", nextIdx),
	}

	// Update env
	currentEnv.Variables = append(currentEnv.Variables, variableOut)

	return variableOut

}

/*
ToMonomial
Description:

	Converts the variable into a monomial.
*/
func (v Variable) ToMonomial() Monomial {
	return Monomial{
		Coefficient:     1.0,
		VariableFactors: []Variable{v},
		Exponents:       []int{1},
	}
}

/*
ToPolynomial
Description:

	Converts the variable into a monomial and then into a polynomial.
*/
func (v Variable) ToPolynomial() Polynomial {
	return Polynomial{
		Monomials: []Monomial{v.ToMonomial()},
	}
}

/*
DerivativeWrt
Description:

	Computes the derivative of the Variable with respect to vIn.
	If vIn is the same as the Variable, then this returns 1.0. Otherwise, it
	returns 0.0.
*/
func (v Variable) DerivativeWrt(vIn Variable) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	if v.ID == vIn.ID {
		return K(1.0)
	} else {
		return K(0.0)
	}
}

/*
IsLinear
Description:

	This function always returns true. A single variable is always a linear expression.
*/
func (v Variable) IsLinear() bool {
	return true
}

/*
String
Description:

	This method returns a string representation of the variable.
*/
func (v Variable) String() string {
	return v.Name
}
