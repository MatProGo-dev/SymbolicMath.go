package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
)

/*
FindInSlice
Description:

	Identifies if the  input xIn is in the slice sliceIn.
	If it is, then this function returns the index such that xIn = sliceIn[index] and no errors.
	If it is not, then this function returns the index -1 and the boolean value false.
*/
func FindInSlice(xIn interface{}, sliceIn interface{}) (int, error) {
	// Constants
	allowedTypes := []string{"string", "int", "uint64", "Variable"}

	switch x := xIn.(type) {
	case string:
		slice, ok := sliceIn.([]string)
		if !ok {
			return -1, fmt.Errorf(
				"the input slice is of type %T, but the element we're searching for is of type %T",
				sliceIn,
				x,
			)
		}

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case int:
		slice, ok := sliceIn.([]int)
		if !ok {
			return -1, fmt.Errorf(
				"the input slice is of type %T, but the element we're searching for is of type %T",
				sliceIn,
				x,
			)
		}

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case uint64:
		slice := sliceIn.([]uint64)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case Variable:
		slice, ok := sliceIn.([]Variable)
		if !ok {
			return -1, fmt.Errorf(
				"the input slice is of type %T, but the element we're searching for is of type %T",
				sliceIn,
				x,
			)
		}

		// Perform Search
		xLocationInSliceIn := -1
		for sliceIndex, sliceValue := range slice {
			if x.ID == sliceValue.ID {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	default:

		return -1, fmt.Errorf(
			"the FindInSlice() function was only defined for types %v, not type %T:",
			allowedTypes,
			xIn,
		)
	}

}

/*
Unique
Description:

	Returns the unique list of variables in a slice of uint64's.
*/
func Unique(listIn []uint64) []uint64 {
	// Create unique list
	var uniqueList []uint64

	// For each int in the list, determine if it previously existed in the list.
	for listIndex, tempElt := range listIn {
		// Don't do any checks if this is the first element.
		if listIndex == 0 {
			uniqueList = append(uniqueList, tempElt)
			continue
		}

		// check to see if the current element exists in the uniqueList.
		if foundIndex, _ := FindInSlice(tempElt, uniqueList); foundIndex == -1 {
			// tempElt does not exist in uniqueList already. Add it.
			uniqueList = append(uniqueList, tempElt)
		}
		// Otherwise, don't add it.
	}

	return uniqueList
}

/*
CheckDimensionsInComparison
Description:

	Verifies that the two objects being compared in a Comparison method (Comparison, LessEq, Eq, GreaterEq)
	have the same dimensions.
*/
func CheckDimensionsInComparison(left, right Expression, senseIn ConstrSense) error {
	// Check that the size of columns in left and right agree
	dimsAreMatched := (left.Dims()[0] == right.Dims()[0]) && (left.Dims()[1] == right.Dims()[1])
	dimsAreMatched = dimsAreMatched || IsScalarExpression(left)
	dimsAreMatched = dimsAreMatched || IsScalarExpression(right)

	if !dimsAreMatched {
		return smErrors.DimensionError{
			Operation: "Comparison (" + senseIn.String() + ")",
			Arg1:      left,
			Arg2:      right,
		}
	}
	// If dimensions match, then return nothing.
	return nil
}

/*
CheckSubstitutionMap
Description:

	This function verifies that the input substitution map is valid. i.e., the map should contain:
		1. Valid variables as keys
		2. Valid expressions as values
		3. The values should also be scalar expressions
*/
func CheckSubstitutionMap(subMap map[Variable]Expression) error {
	var err error
	// Check each key, value pair in the map
	for tempVar, tempExpr := range subMap {
		// Verify that the key is a valid variable
		err = tempVar.Check()
		if err != nil {
			return fmt.Errorf(
				"key %v in the substitution map is not a valid variable: %v",
				tempVar,
				err,
			)
		}

		// Verify that the value is a valid expression
		err = tempExpr.Check()
		if err != nil {
			return fmt.Errorf(
				"value %v in the substitution map[%v] is not a valid expression: %v",
				tempExpr,
				tempVar,
				err,
			)
		}

		// Verify that the value is a scalar expression
		if !IsScalarExpression(tempExpr) {
			return fmt.Errorf(
				"value %v in the substitution map[%v] is not a scalar expression (received %T)",
				tempExpr,
				tempVar,
				tempExpr,
			)
		}
	}

	// All checks passed
	return nil
}
