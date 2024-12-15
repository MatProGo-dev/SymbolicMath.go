package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
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

// Constant returns the constant additive value in the expression. For a
// variable, it always returns zero.
func (v Variable) Constant() float64 {
	return 0
}

/*
LinearCoeff
Description:

	Returns the coefficient of the linear term in the expression. For a variable,
	this is always a vector containing exactly 1 value of 1.0 all others are zero.
*/
func (v Variable) LinearCoeff(wrt ...[]Variable) mat.VecDense {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	var wrtVars []Variable
	switch len(wrt) {
	case 0:
		wrtVars = v.Variables()
	case 1:
		wrtVars = wrt[0]
	default:
		panic(
			fmt.Errorf("Too many inputs provided to Variable.LinearCoeff() method."),
		)
	}

	if len(wrtVars) == 0 {
		panic(smErrors.CanNotGetLinearCoeffOfConstantError{v})
	}

	// Constants

	// Algorithm
	coeffOut := ZerosVector(len(wrtVars))
	for i, wrtVar := range wrtVars {
		if v.ID == wrtVar.ID {
			coeffOut.SetVec(i, 1.0)
		}
	}

	return coeffOut
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
	case *mat.VecDense:
		return v.Plus(VecDenseToKVector(*right))
	case KVector, VariableVector, MonomialVector, PolynomialVector:
		ve, _ := ToVectorExpression(rightIn)
		return ve.Plus(v)
	}

	panic(
		fmt.Errorf("there input %v has unexpected type %T given to Variable.Plus()!", rightIn, rightIn),
	)
}

/*
Minus
Description:

	This function subtracts an expression from the current
	variable.
*/
func (v Variable) Minus(rightIn interface{}) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		// Check dimensions
		err = smErrors.CheckDimensionsInSubtraction(v, rightAsE)
		if err != nil {
			panic(err)
		}

		// Use Expression's Minus() method
		return Minus(v, rightAsE)
	}

	// Algorithm
	switch right := rightIn.(type) {
	case float64:
		return v.Minus(K(right))
	}

	panic(
		smErrors.UnsupportedInputError{
			FunctionName: "Variable.Minus",
			Input:        rightIn,
		},
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
		err := rightAsE.Check()
		if err != nil {
			panic(err)
		}

		err = smErrors.CheckDimensionsInMultiplication(v, rightAsE)
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
	case Monomial:
		// Use Monomial method
		return right.Multiply(v)
	case Polynomial:
		// Create a new vector of polynomials.
		return right.Multiply(v)
	case *mat.VecDense:
		return v.Multiply(*right)
	case mat.VecDense:
		// Convert to KVector
		return v.Multiply(VecDenseToKVector(right))
	case KVector:
		// Create a monomial vector and store result in it
		var monomialsOut MonomialVector = make([]Monomial, right.Len())
		for i := 0; i < right.Len(); i++ {
			monomialsOut[i] = Monomial{
				Coefficient:     float64(right[i]),
				VariableFactors: []Variable{v},
				Exponents:       []int{1},
			}
		}
		return monomialsOut
	}

	// Unrecornized response is a panic
	panic(
		fmt.Errorf("Unexpected input to Variable.Multiply(): %T", rightIn),
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

	Checks whether the Variable has a sensible initialization.
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
func NewBinaryVariable(envs ...*Environment) Variable {
	// Constants

	// Input Processing
	var currentEnv = &BackgroundEnvironment
	switch len(envs) {
	case 1:
		currentEnv = envs[0]
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
Degree
Description:

	Returns the degree of the variable (which is always 1).
*/
func (v Variable) Degree() int {
	return 1
}

/*
String
Description:

	This method returns a string representation of the variable.
*/
func (v Variable) String() string {
	return v.Name
}

/*
Substitute
Description:

	Substitutes the variable vIn with the expression eIn.
*/
func (v Variable) Substitute(vIn Variable, seIn ScalarExpression) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	err = seIn.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	if v.ID == vIn.ID {
		return seIn
	} else {
		return v
	}
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variable in the map with the corresponding expression.
*/
func (v Variable) SubstituteAccordingTo(subMap map[Variable]Expression) Expression {
	// Input Processing
	err := v.Check()
	if err != nil {
		panic(err)
	}

	err = CheckSubstitutionMap(subMap)
	if err != nil {
		panic(err)
	}

	// Algorithm
	if e, ok := subMap[v]; ok {
		return e
	} else {
		return v
	}
}

/*
Power
Description:

	Computes the power of the variable.
*/
func (v Variable) Power(exponent int) Expression {
	return ScalarPowerTemplate(v, exponent)
}

/*
At
Description:

	Returns the value at the given row and column index.

Note:

	For a variable, the row and column index should always be 0.
*/
func (v Variable) At(ii, jj int) ScalarExpression {
	// Input Processing

	// Check the variable
	err := v.Check()
	if err != nil {
		panic(err)
	}

	// Check to see whether or not the index is valid.
	err = smErrors.CheckIndexOnMatrix(ii, jj, v)
	if err != nil {
		panic(err)
	}

	// Algorithm
	return v
}
