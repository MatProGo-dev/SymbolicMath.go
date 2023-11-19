package symbolic

import (
	"fmt"
)

/*
dimension.go
*/

/* Type Definitions */

type DimensionError struct {
	Arg1      Expression
	Arg2      Expression
	Operation string // Either multiply or Plus
}

type UnexpectedInputError struct {
	InputInQuestion interface{}
	Operation       string
}

/* Methods */

func (de DimensionError) Error() string {
	dimStrings := de.ArgDimsAsStrings()
	return fmt.Sprintf(
		"dimension error: Cannot perform %v between expression of dimension %v and expression of dimension %v",
		de.Operation,
		dimStrings[0],
		dimStrings[1],
	)
}

func (de DimensionError) ArgDimsAsStrings() []string {
	// Create string for arg 1
	arg1DimsAsString := "("
	for ii, dimValue := range de.Arg1.Dims() {
		arg1DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg1.Dims())-1 { // If this isn't the last element of size
			arg1DimsAsString += ","
		}
	}
	arg1DimsAsString += ")"

	// Create string for arg 2
	arg2DimsAsString := "("
	for ii, dimValue := range de.Arg2.Dims() {
		arg2DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg2.Dims())-1 { // If this isn't the last element of size
			arg2DimsAsString += ","
		}
	}
	arg2DimsAsString += ")"

	return []string{arg1DimsAsString, arg2DimsAsString}

}

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

func CheckDimensionsInMultiplication(left, right Expression) error {
	// Check that the # of columns in left
	// matches the # of rows in right
	if left.Dims()[1] != right.Dims()[0] {
		return DimensionError{
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
		return DimensionError{
			Operation: "Plus",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}
