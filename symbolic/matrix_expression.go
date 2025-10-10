package symbolic

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"gonum.org/v1/gonum/mat"
)

/*
   matrix_expression.go
   Description:

*/

type MatrixExpression interface {
	// Check returns an error if the expression is not initialized properly
	Check() error

	// Variables returns the number of variables in the expression.
	Variables() []Variable

	//// Coeffs returns a slice of the coefficients in the expression
	//LinearCoeff() mat.Dense

	// Constant returns the constant additive value in the expression
	Constant() mat.Dense

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

	//AtVec returns the expression at a given index
	At(i int, j int) ScalarExpression

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

	// Power
	// Raises the scalar expression to the power of the input integer
	Power(exponent int) Expression

	// Simplify simplifies the expression and returns the simplified version
	AsSimplifiedExpression() Expression
}

/*
IsMatrixExpression
Description:

	Determines whether or not an input object is a valid "VectorExpression" according to MatProInterface.
*/
func IsMatrixExpression(e interface{}) bool {
	// Check each type
	switch e.(type) {
	case *mat.Dense:
		return true
	case mat.Dense:
		return true
	case KMatrix:
		return true
	case VariableMatrix:
		return true
	case MonomialMatrix:
		return true
	case PolynomialMatrix:
		return true
	default:
		return false

	}
}

/*
ToMatrixExpression
Description:

	Converts the input expression to a valid type that implements "VectorExpression".
*/
func ToMatrixExpression(e interface{}) (MatrixExpression, error) {
	// Input Processing
	if !IsMatrixExpression(e) {
		return DenseToKMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"the input interface is of type %T, which is not recognized as a MatrixExpression.",
			e,
		)
	}

	// Convert
	switch e2 := e.(type) {
	case *mat.Dense:
		return DenseToKMatrix(*e2), nil
	case mat.Dense:
		return DenseToKMatrix(e2), nil
	case KMatrix:
		return e2, nil
	case VariableMatrix:
		return e2, nil
	case MonomialMatrix:
		return e2, nil
	case PolynomialMatrix:
		return e2, nil
	default:
		return DenseToKMatrix(ZerosMatrix(1, 1)), fmt.Errorf(
			"unexpected matrix expression conversion requested for object \"%v\" of type %T!",
			e,
			e,
		)
	}
}

/*
IsSquare
Description:

	Determines whether the input matrix expression is square.
*/
func IsSquare(e MatrixExpression) bool {
	dims := e.Dims()
	return dims[0] == dims[1]
}

/*
MatrixPowerTemplate
Description:

	Template for the matrix power function.
*/
func MatrixPowerTemplate(me MatrixExpression, exponent int) MatrixExpression {
	// Input Processing
	err := me.Check()
	if err != nil {
		panic(err)
	}

	// Check if the matrix is square
	if !IsSquare(me) {
		panic(fmt.Errorf("matrix is not square; cannot raise to power"))
	}

	// Check if the power is non-negative
	if exponent < 0 {
		panic(smErrors.NegativeExponentError{
			Exponent: exponent,
		})
	}

	// Algorithm
	var out MatrixExpression = K(0).Plus(Identity(me.Dims()[0])).(MatrixExpression)
	for i := 0; i < exponent; i++ {
		out = out.Multiply(me).(MatrixExpression)
	}
	return out
}

/*
MatrixMultiplyTemplate
Description:

	Template for the matrix multiply function.
*/
func MatrixMultiplyTemplate(left MatrixExpression, right MatrixExpression) Expression {
	// Input Processing
	err := left.Check()
	if err != nil {
		panic(err)
	}

	err = right.Check()
	if err != nil {
		panic(err)
	}

	// Check dimensions
	leftDims := left.Dims()
	rightDims := right.Dims()

	if leftDims[1] != rightDims[0] {
		panic(
			smErrors.MatrixDimensionError{
				Arg1:      left,
				Arg2:      right,
				Operation: "MatrixMultiplyTemplate",
			},
		)
	}

	// Algorithm
	var out [][]ScalarExpression
	for ii := 0; ii < leftDims[0]; ii++ {
		var tempRow []ScalarExpression
		for jj := 0; jj < rightDims[1]; jj++ {
			// Compute the (ii,jj) element of the product
			var sum Expression = K(0.0)
			for kk := 0; kk < leftDims[1]; kk++ {
				sum = sum.Plus(left.At(ii, kk).Multiply(right.At(kk, jj)))
				fmt.Printf("sum after adding entry %v: %v\n", kk, sum)
			}
			sumAsSE, tf := sum.AsSimplifiedExpression().(ScalarExpression)
			if !tf {
				panic(
					fmt.Errorf(
						"unexpected expression type in MatrixMultiplyTemplate at entry [%v,%v]: %T",
						ii, jj,
						sum,
					),
				)
			}
			tempRow = append(tempRow, sumAsSE)
		}
		out = append(out, tempRow)
	}

	// Use the general concretization function (not the matrix-specific one)
	// because it will also convert matrices to scalars or vectors if needed.
	return ConcretizeExpression(out)
}

/*
MatrixSubstituteTemplate
Description:

	Template for the matrix substitute function.
*/
func MatrixSubstituteTemplate(me MatrixExpression, vIn Variable, seIn ScalarExpression) MatrixExpression {
	// Input Processing
	err := me.Check()
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
	var out [][]ScalarExpression
	for ii := 0; ii < me.Dims()[0]; ii++ {
		var tempRow []ScalarExpression
		for jj := 0; jj < me.Dims()[1]; jj++ {
			newElt := me.At(ii, jj).Substitute(vIn, seIn)
			tempRow = append(tempRow, newElt.(ScalarExpression))
		}
		out = append(out, tempRow)
	}
	return ConcretizeMatrixExpression(out)
}

/*
ConcretizeMatrixExpression
Description:

	Converts the input expression to a valid type that implements "MatrixExpression".
*/
func ConcretizeMatrixExpression(sliceIn [][]ScalarExpression) MatrixExpression {
	// Input Processing
	// - Check that the input slice is not empty
	if len(sliceIn) == 0 {
		panic(
			fmt.Errorf(
				"the input slice is empty, which is not recognized as a VectorExpression.",
			),
		)
	}

	// - Check the number of columns in each row is the same
	numCols := len(sliceIn[0])
	for ii, row := range sliceIn {
		if len(row) != numCols {
			panic(
				fmt.Errorf(
					"all rows in the input slice must have the same number of columns, but row %v has %v columns (expected %v).",
					ii,
					len(row),
					numCols,
				),
			)

		}
	}

	// Check the type of all expressions
	var (
		containsConstant   bool = false
		isAllVariables     bool = true
		containsVariable   bool = false
		containsMonomial   bool = false
		containsPolynomial bool = false
	)

	for ii, row := range sliceIn {
		for jj, elt := range row {
			// Check each element in row
			if _, tf := elt.(Variable); !tf {
				isAllVariables = false
			}

			switch elt.(type) {
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
					fmt.Errorf(
						"unexpected expression type in matrix expression at [%v,%v]: %T",
						ii, jj,
						elt,
					),
				)
			}
		}
	}

	// Convert
	switch {
	case containsPolynomial:
		// Convert to a polynomial vector
		var out PolynomialMatrix
		for _, row_ii := range sliceIn {
			var tempRow []Polynomial
			for _, elt := range row_ii {
				switch tempE := elt.(type) {
				case Polynomial:
					tempRow = append(tempRow, tempE)
				case Monomial:
					tempRow = append(tempRow, tempE.ToPolynomial())
				case Variable:
					tempRow = append(tempRow, tempE.ToPolynomial())
				case K:
					tempRow = append(tempRow, tempE.ToPolynomial())
				default:
					panic(
						smErrors.UnsupportedInputError{
							FunctionName: "ConcretizeVectorExpression",
							Input:        tempE,
						},
					)
				}
			}
			out = append(out, tempRow)
		}

		return out

	case containsMonomial || (containsVariable && containsConstant):
		// Convert to a monomial vector
		var out MonomialMatrix
		for _, row_ii := range sliceIn {
			var tempRow []Monomial
			for _, elt := range row_ii {
				switch tempE := elt.(type) {
				case Monomial:
					tempRow = append(tempRow, tempE)
				case Variable:
					tempRow = append(tempRow, tempE.ToMonomial())
				case K:
					tempRow = append(tempRow, tempE.ToMonomial())
				default:
					panic(
						smErrors.UnsupportedInputError{
							FunctionName: "ConcretizeVectorExpression",
							Input:        tempE,
						},
					)
				}
			}
			out = append(out, tempRow)
		}

		return out

	case isAllVariables:
		// Convert to a variable matrix
		var out VariableMatrix
		for _, row_ii := range sliceIn {
			var tempRow []Variable
			for _, elt := range row_ii {
				switch tempE := elt.(type) {
				case Variable:
					tempRow = append(tempRow, tempE)
				default:
					panic(
						smErrors.UnsupportedInputError{
							FunctionName: "ConcretizeVectorExpression",
							Input:        tempE,
						},
					)
				}
			}
			out = append(out, tempRow)
		}

		return out

	case containsConstant:
		// Convert to a constant vector
		var out KMatrix
		for ii, row_ii := range sliceIn {
			var tempRow []K
			for jj, elt := range row_ii {
				eltAsK, tf := elt.(K)
				if !tf {
					panic(
						fmt.Errorf(
							"unexpected expression type in vector expression at entry [%v,%v]: %T",
							ii, jj,
							elt,
						),
					)
				}
				tempRow = append(tempRow, eltAsK)
			}

			out = append(out, tempRow)
		}

		return out

	default:
		panic(
			fmt.Errorf(
				"unrecognized vector expression type in ConcretizeMatrixExpression.\n"+
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
