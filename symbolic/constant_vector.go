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
Check
Description:

	This method is used to make sure that the variable is well-defined.
	For a constant vector, the vecdense should always be well-defined.
*/
func (kv KVector) Check() error {
	return nil
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
Variables
Description:

	This function returns the empty slice because no variables are in a constant vector.
*/
func (kv KVector) Variables() []Variable {
	return []Variable{}
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
func (kv KVector) LinearCoeff() mat.Dense {
	return ZerosMatrix(kv.Len(), kv.Len())
}

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
func (kv KVector) Plus(rightIn interface{}) Expression {
	// Input Processing
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInAddition(kv, rightAsE)
		if err != nil {
			panic(err)
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

		return KVector(result)
	case KVector:
		// Compute Addition
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		eAsVec := mat.VecDense(right)
		result.AddVec(&kvAsVec, &eAsVec)

		return KVector(result)

	case VariableVector:
		return right.Plus(kv)

	default:
		errString := fmt.Sprintf("Unrecognized expression type %T for addition of KVector kv.Plus(%v)!", right, right)
		panic(fmt.Errorf(errString))
	}
}

/*
LessEq
Description:

	Returns a less than or equal to (<=) constraint between the current expression and another
*/
func (kv KVector) LessEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseLessThanEqual)
}

/*
GreaterEq
Description:

	This method returns a greater than or equal to (>=) constraint between the current expression and another
*/
func (kv KVector) GreaterEq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseGreaterThanEqual)
}

/*
Eq
Description:

	This method returns an equality (==) constraint between the current expression and another
*/
func (kv KVector) Eq(rightIn interface{}) Constraint {
	return kv.Comparison(rightIn, SenseEqual)
}

func (kv KVector) Comparison(rightIn interface{}, sense ConstrSense) Constraint {
	switch rhsConverted := rightIn.(type) {
	case KVector:
		// Check Lengths
		if kv.Len() != rhsConverted.Len() {
			panic(
				fmt.Errorf(
					"The left hand side's dimension (%v) and the left hand side's dimension (%v) do not match!",
					kv.Len(),
					rhsConverted.Len(),
				),
			)
		}

		// Return constraint
		return VectorConstraint{
			LeftHandSide:  kv,
			RightHandSide: rhsConverted,
			Sense:         sense,
		}

	case VariableVector:
		// Return constraint
		return rhsConverted.Comparison(kv, sense)

	default:
		// Return an error
		panic(
			fmt.Errorf(
				"The input to KVector's '%v' comparison (%v) has unexpected type: %T",
				sense, rightIn, rightIn,
			),
		)

	}
}

/*
Multiply
Description:

	This method is used to compute the multiplication of the input vector constant with another term.
*/
func (kv KVector) Multiply(rightIn interface{}) Expression {
	// Input Processing
	err := kv.Check()
	if err != nil {
		panic(err)
	}

	if IsExpression(rightIn) {
		// Check dimensions
		rightAsE, _ := ToExpression(rightIn)
		err = CheckDimensionsInMultiplication(kv, rightAsE)
		if err != nil {
			panic(err)
		}
	}

	// Compute Multiplication
	switch right := rightIn.(type) {
	case float64:
		// Use mat.Vector's multiplication method
		var result mat.VecDense
		kvAsVec := mat.VecDense(kv)
		result.ScaleVec(right, &kvAsVec)

		return KVector(result)
	case K:
		// Convert to float64
		eAsFloat := float64(right)

		return kv.Multiply(eAsFloat)

	case mat.VecDense:
		// Send warning until we create matrix type.
		panic(
			fmt.Errorf(
				"MatProInterface does not currently support operations that result in matrices! if you want this feature, create an issue!",
			),
		)

	case KVector:
		// Immediately return error.
		panic(fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		))

	case VariableVector:
		// Immediately return error.
		panic(fmt.Errorf(
			"dimension mismatch! Cannot multiply KVector with a vector of type %T; Try transposing one or the other!",
			right,
		))

	default:
		panic(fmt.Errorf(
			"The input to KVectorTranspose's Multiply method (%v) has unexpected type: %T",
			right, right,
		))

	}
}

/*
Transpose
Description:

	This method creates the transpose of the current vector and returns it.
*/
func (kv KVector) Transpose() Expression {
	// Constants
	kvAsVD := mat.VecDense(kv)
	kvLen := kv.Len()

	// Create empty matrix and populate
	kvT := ZerosMatrix(1, kvLen)
	for colIndex := 0; colIndex < kvLen; colIndex++ {
		kvT.Set(0, colIndex, kvAsVD.AtVec(colIndex))
	}

	return KMatrix(kvT)
}

/*
Dims
Description:

	Returns the dimension of the constant vector.
*/
func (kv KVector) Dims() []int {
	return []int{kv.Len(), 1}
}

// Other Functions

/*
OnesVector
Description:

	Returns a vector of ones with length lengthIn.
	Note: this function assumes lengthIn is a positive number.
*/
func OnesVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 1.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

/*
ZerosVector
Description:

	Returns a vector of zeros with length lengthIn.
	Note: this function assumes lengthIn is a positive number.
*/
func ZerosVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 0.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}
