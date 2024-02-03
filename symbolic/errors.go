package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
dimension.go
*/

/* Type Definitions */

type UnexpectedInputError struct {
	InputInQuestion interface{}
	Operation       string
}

/* Methods */

func (uie UnexpectedInputError) Error() string {
	return fmt.Sprintf(
		"Unexpected input to \"%v\" operation: %T",
		uie.Operation,
		uie.InputInQuestion,
	)
}

/*
CheckErrors
Description:
*/
func CheckErrors(extras []error) error {
	// Constants

	// Check all of the extras to see if one of them contains an error
	switch {
	case len(extras) == 1:
		return extras[0]

	case len(extras) > 1:
		return fmt.Errorf(
			"did not expect to receive more than one element in 'extras' input; received %v",
			len(extras),
		)
	}

	// If extras has length 0, then return nil
	return nil
}

/*
CheckDimensionsInMultiplication
Description:

	This function checks that the dimensions of the left and right expressions
	are compatible for multiplication.
	We allow:
	- Multiplication if Dimensions Match OR
	- Multiplication if one of the expressions is a scalar
*/
func CheckDimensionsInMultiplication(left, right Expression) error {
	// Check that dimensions match
	dimsMatch := left.Dims()[1] == right.Dims()[0]

	// Check that one of the expressions is a scalar
	leftIsScalar := IsScalarExpression(left)
	rightIsScalar := IsScalarExpression(right)

	multiplicationIsAllowed := dimsMatch || leftIsScalar || rightIsScalar

	// Check that the # of columns in left
	// matches the # of rows in right
	if !multiplicationIsAllowed {
		return smErrors.DimensionError{
			Operation: "Multiply",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}

func CheckDimensionsInAddition(left, right Expression) error {
	// Check that the size of columns in left and right agree
	dimsAreMatched := (left.Dims()[0] == right.Dims()[0]) && (left.Dims()[1] == right.Dims()[1])
	dimsAreMatched = dimsAreMatched || IsScalarExpression(left)
	dimsAreMatched = dimsAreMatched || IsScalarExpression(right)

	if !dimsAreMatched {
		return smErrors.DimensionError{
			Operation: "Plus",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}
