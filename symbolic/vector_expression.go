package symbolic

/*
vector_expression.go
Description:
	An improvement/successor to the scalar expr interface.
*/

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
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

	// LinearCoeffs returns a slice of the coefficients in the expression
	LinearCoeff(wrt ...[]Variable) mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.VecDense

	// Plus adds the current expression to another and returns the resulting
	// expression
	Plus(e interface{}) Expression

	// Minus subtracts an expression from the current one and returns the resulting
	// expression
	Minus(rightIn interface{}) Expression

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

	// At returns the expression at a given indices
	At(ii, jj int) ScalarExpression

	//AtVec returns the expression at a given index
	AtVec(idx int) ScalarExpression

	//Transpose returns the transpose of the given vector expression
	Transpose() Expression

	// Dims returns the dimensions of the given expression
	Dims() []int

	// DerivativeWrt returns the derivative of the expression with respect to the input variable vIn.
	DerivativeWrt(vIn Variable) Expression

	// String returns a string representation of the expression
	String() string

	// Substitute returns the expression with the variable vIn replaced with the expression eIn
	Substitute(vIn Variable, eIn ScalarExpression) Expression

	// SubstituteAccordingTo returns the expression with the variables in the map replaced with the corresponding expressions
	SubstituteAccordingTo(subMap map[Variable]Expression) Expression

	// Power returns the expression raised to the power of the input exponent
	Power(exponent int) Expression

	// Simplify simplifies the expression and returns the simplified version
	AsSimplifiedExpression() Expression
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
	case *mat.VecDense:
		return true
	case mat.VecDense:
		return true
	case KVector:
		return true
	case VariableVector:
		return true
	case MonomialVector:
		return true
	case PolynomialVector:
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
		return VecDenseToKVector(OnesVector(1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a VectorExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case KVector:
		return e2, nil
	case *mat.VecDense:
		return VecDenseToKVector(*e2), nil
	case mat.VecDense:
		return VecDenseToKVector(e2), nil
	case VariableVector:
		return e2, nil
	case MonomialVector:
		return e2, nil
	case PolynomialVector:
		return e2, nil
	default:
		return VecDenseToKVector(OnesVector(1)), fmt.Errorf(
			"unexpected vector expression conversion requested for type %T!",
			e,
		)
	}
}

/*
ConcretizeVectorExpression
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ConcretizeVectorExpression(sliceIn []ScalarExpression) VectorExpression {
	// Input Processing
	if len(sliceIn) == 0 {
		panic(
			fmt.Errorf(
				"the input slice is empty, which is not recognized as a VectorExpression.",
			),
		)
	}

	// Check the type of all expressions
	var (
		containsConstant   bool = false
		isAllVariables     bool = true
		containsVariable   bool = false
		containsMonomial   bool = false
		containsPolynomial bool = false
	)

	for _, expr := range sliceIn {
		if _, tf := expr.(Variable); !tf {
			isAllVariables = false
		}

		switch expr.(type) {
		case K:
			containsConstant = true
		case Variable:
			containsVariable = true
		case Monomial:
			containsMonomial = true
		case Polynomial:
			containsPolynomial = true
		default:
			panic(
				fmt.Errorf("unexpected expression type in vector expression: %T", expr),
			)
		}
	}

	// Convert
	switch {
	case containsPolynomial:
		// Convert to a polynomial vector
		var out PolynomialVector
		for _, e_ii := range sliceIn {
			switch tempE := e_ii.(type) {
			case Polynomial:
				out = append(out, tempE)
			case Monomial:
				out = append(out, tempE.ToPolynomial())
			case Variable:
				out = append(out, tempE.ToPolynomial())
			case K:
				out = append(out, tempE.ToPolynomial())
			default:
				panic(
					smErrors.UnsupportedInputError{
						FunctionName: "ConcretizeVectorExpression",
						Input:        tempE,
					},
				)
			}
		}

		return out

	case containsMonomial || (containsConstant && containsVariable):
		// Convert to a monomial vector
		var out MonomialVector
		for _, e_ii := range sliceIn {
			switch tempE := e_ii.(type) {
			case Monomial:
				out = append(out, tempE)
			case Variable:
				out = append(out, tempE.ToMonomial())
			case K:
				out = append(out, tempE.ToMonomial())
			default:
				panic(
					smErrors.UnsupportedInputError{
						FunctionName: "ConcretizeVectorExpression",
						Input:        tempE,
					},
				)
			}
		}

		return out

	case isAllVariables:
		// Convert to a variable vector
		var out VariableVector
		for _, e_ii := range sliceIn {
			switch tempE := e_ii.(type) {
			case Variable:
				out = append(out, tempE)
			default:
				panic(
					smErrors.UnsupportedInputError{
						FunctionName: "ConcretizeVectorExpression",
						Input:        tempE,
					},
				)
			}
		}

		return out

	case containsConstant:
		// Convert to a constant vector
		var out KVector
		for ii, i2 := range sliceIn {
			i2AsK, tf := i2.(K)
			if !tf {
				panic(
					fmt.Errorf(
						"unexpected expression type in vector expression at %v: %T",
						ii,
						i2,
					),
				)
			}
			out = append(out, i2AsK)
		}

		return out

	default:
		panic(
			fmt.Errorf(
				"unrecognized vector expression type in ConcretizeVectorExpression.\n"+
					"containsConstant = %v\n"+
					"isAllVariables = %v\n"+
					"containsMonomial = %v\n"+
					"containsPolynomial = %v\n",
				containsConstant,
				isAllVariables,
				containsMonomial,
				containsPolynomial,
			),
		)
	}
}

/*
VectorSubstituteTemplate
Description:

	Defines the template for the vector substitution operation.
*/
func VectorSubstituteTemplate(ve VectorExpression, vIn Variable, se ScalarExpression) VectorExpression {
	// Input Processing
	err := ve.Check()
	if err != nil {
		panic(err)
	}

	err = vIn.Check()
	if err != nil {
		panic(err)
	}

	err = se.Check()
	if err != nil {
		panic(err)
	}

	// Algorithm
	var result []ScalarExpression
	for ii := 0; ii < ve.Len(); ii++ {
		eltII := ve.AtVec(ii)
		postSub := eltII.Substitute(vIn, se)
		result = append(result, postSub.(ScalarExpression))
	}

	return ConcretizeVectorExpression(result)
}

/*
VectorPowerTemplate
Description:

	Defines the template for the vector power operation.
*/
func VectorPowerTemplate(base VectorExpression, exponent int) Expression {
	// Setup

	// Input Processing
	err := base.Check()
	if err != nil {
		panic(err)
	}

	if exponent < 0 {
		panic(smErrors.NegativeExponentError{Exponent: exponent})
	}

	if base.Len() != 1 {
		panic(
			fmt.Errorf(
				"the Power operation is only defined for vectors of length 1, but the input vector has length %v.",
				base.Len(),
			),
		)

	}

	// Algorithm
	var result Expression = K(1.0)
	for i := 0; i < exponent; i++ {
		result = result.Multiply(base.AtVec(0))
	}

	return result
}
