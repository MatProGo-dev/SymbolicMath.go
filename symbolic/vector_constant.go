package symbolic

import "gonum.org/v1/gonum/mat"

/*
vector_constant_test.go
Description:
	Creates a vector extension of the constant type K from the original goop.
*/

import (
	"fmt"
)

/*
KVector

	A type which is built on top of the KVector()
	a constant expression type for an MIP. K for short ¯\_(ツ)_/¯
*/
type KVector mat.VecDense // Inherit all methods from mat.VecDense

/*
Len

	Computes the length of the KVector given.
*/
func (kv KVector) Len() int {
	kvAsVector := mat.VecDense(kv)
	return kvAsVector.Len()
}

/*
AtVec
Description:

	This function returns the value at the k index.
*/
func (kv KVector) AtVec(idx int) ScalarExpression {
	kvAsVector := mat.VecDense(kv)
	return K(kvAsVector.AtVec(idx))
}

/*
NumVars
Description:

	This returns the number of variables in the expression. For constants, this is 0.
*/
func (kv KVector) NumVars() int {
	return 0
}

/*
Vars
Description:

	This function returns a slice of the Var ids in the expression. For constants, this is always nil.
*/
func (kv KVector) IDs() []uint64 {
	return nil
}

/*
LinearCoeff
Description:

	This function returns a slice of the coefficients in the expression. For constants, this is always nil.
*/
//func (kv KVector) LinearCoeff() mat.Dense {
//	return ZerosMatrix(kv.Len(), kv.Len())
//}

/*
Constant

	Returns the constant additive value in the expression. For constants, this is just the constants value
*/
func (kv KVector) Constant() mat.VecDense {
	return mat.VecDense(kv)
}

/*
Plus
Description:

	Adds the current expression to another and returns the resulting expression
*/
func (kv KVector) Plus(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return kv, err
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(kv, rightAsE)
		if err != nil {
			return kv, err
		}
	}

	// Constants
	kvLen := kv.Len()

	// Management
	switch right := rightIn.(type) {
	case float64:
		// Create vector
		tempOnes := OnesVector(kvLen)
		var eAsVec mat.VecDense
		eAsVec.ScaleVec(right, &tempOnes)

		// Add the values
		return kv.Plus(KVector(eAsVec))
	case K:
		// Return Addition
		return kv.Plus(float64(right))
	case mat.VecDense:
		// Return Sum
		var result mat.VecDense
		kv2 := mat.VecDense(kv)
		result.AddVec(&kv2, &right)

		return KVector(result), nil
	case KVector:
		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		eAsVec := mat.VecDense(right)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVector(result), nil

	case KVectorTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				right,
			)

	case VarVector:
		return right.Plus(kv)

	case VarVectorTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				right,
			)

	case VectorLinearExpr:
		return right.Plus(kv)

	case VectorLinearExpressionTranspose:
		return kv,
			fmt.Errorf(
				"Cannot add KVector with a transposed vector %T; Try transposing one or the other!",
				right,
			)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVector kv.Plus(%v)!", right, right)
		return KVector{}, fmt.Errorf(errString)
	}
}

/*
Mult
Description:

	This method multiplies the current expression to another and returns the resulting expression.
*/
func (kv KVector) Mult(val float64) (VectorExpression, error) {

	// Use mat.Vector's multiplication method
	var result mat.VecDense
	kvAsVec := mat.VecDense(kv)
	result.ScaleVec(val, &kvAsVec)

	return KVector(result), nil
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kv KVector) LessEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kv.Comparison(rightIn, SenseLessThanEqual, errors...)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kv KVector) GreaterEq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kv.Comparison(rightIn, SenseGreaterThanEqual, errors...)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kv KVector) Eq(rightIn interface{}, errors ...error) (Constraint, error) {
	return kv.Comparison(rightIn, SenseEqual, errors...)
}

func (kv KVector) Comparison(rightIn interface{}, sense ConstrSense, errors ...error) (Constraint, error) {
	switch rhsConverted := rightIn.(type) {
	case KVector:
		// Check Lengths
		if kv.Len() != rhsConverted.Len() {
			return VectorConstraint{},
				fmt.Errorf(
					"The left hand side's dimension (%v) and the left hand side's dimension (%v) do not match!",
					kv.Len(),
					rhsConverted.Len(),
				)
		}

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsConverted,
			Sense:         sense,
		}, nil
	case KVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	case VarVector:
		// Return constraint
		return rhsConverted.Comparison(kv, sense)
	case VarVectorTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	case VectorLinearExpr:
		// Return constraint
		return rhsConverted.Comparison(kv, sense)
	case VectorLinearExpressionTranspose:
		return VectorConstraint{},
			fmt.Errorf(
				"Cannot compare KVector with a transposed vector %T; Try transposing one or the other!",
				rhsConverted,
			)
	default:
		// Return an error
		return VectorConstraint{},
			fmt.Errorf(
				"The input to KVector's '%v' comparison (%v) has unexpected type: %T",
				sense, rightIn, rightIn,
			)

	}
}

/*
Multiply
Description:

	This method is used to compute the multiplication of the input vector constant with another term.
*/
func (kv KVector) Multiply(rightIn interface{}, errors ...error) (Expression, error) {
	// Input Processing
	err := CheckErrors(errors)
	if err != nil {
		return kv, err
	}

	if IsExpression(rightIn) {
		// Check dimensions
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInMultiplication(kv, rightAsE)
		if err != nil {
			return kv, err
		}
	}

	// Compute Multiplication
	switch right := rightIn.(type) {
	case float64:
		// Use mat.Vector's multiplication method
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		result.ScaleVec(right, &kvAsVec)

		return KVector(result), nil
	case K:
		// Convert to float64
		eAsFloat := float64(right)

		return kv.Multiply(eAsFloat)
	case Variable:
		// Create a VectorLinearExpression Output
		vv := VarVector{[]Variable{right}}
		L := ZerosMatrix(kv.Len(), 1)
		for kIndex := 0; kIndex < kv.Len(); kIndex++ {
			L.Set(kIndex, 0, float64(kv.AtVec(kIndex).(K)))
		}

		return VectorLinearExpr{
			L: L,
			X: vv,
			C: ZerosVector(kv.Len()),
		}, nil

	case ScalarLinearExpr:
		// Create a VectorLinearExpression to Output
		L := ZerosMatrix(kv.Len(), right.X.Len())
		for rowIndex := 0; rowIndex < kv.Len(); rowIndex++ {
			for colIndex := 0; colIndex < right.X.Len(); colIndex++ {
				L.Set(
					rowIndex, colIndex,
					right.L.AtVec(colIndex)*float64(kv.AtVec(rowIndex).(K)),
				)
			}
		}
		C := ZerosVector(kv.Len())
		for rowIndex := 0; rowIndex < kv.Len(); rowIndex++ {
			C.SetVec(
				rowIndex,
				right.C*float64(kv.AtVec(rowIndex).(K)),
			)
		}
		return VectorLinearExpr{
			L: L,
			C: C,
			X: right.X.Copy(),
		}, nil

	case mat.VecDense:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case KVector:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		)

	case KVectorTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case VarVector:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		)

	case VarVectorTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	case VectorLinearExpr:
		// Immediately return error.
		return kv, fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		)

	case VectorLinearExpressionTranspose:
		// Send warning until we create matrix type.
		return kv, fmt.Errorf(
			"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
		)

	default:
		return kv, fmt.Errorf(
			"The input to KVectorTranspose's Multiply method (%v) has unexpected type: %T",
			right, right,
		)

	}
}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (kv KVector) Transpose() Expression {
	return KVectorTranspose(kv)
}

/*
Dims
Description:

	Returns the dimension of the constant vector.
*/
func (kv KVector) Dims() []int {
	return []int{kv.Len(), 1}
}
