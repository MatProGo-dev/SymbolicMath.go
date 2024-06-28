package symbolic

/*
utils_test.go
Description:
	Tests some of the functions in the utils.go file.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestUtils_FindInSlice1
Description:

	This function tests whether the find in slice function works properly for
	a slice of strings when the string is NOT in the slice.
*/
func TestUtils_FindInSlice1(t *testing.T) {
	// Constants
	x := "x"
	slice := []string{"a", "b", "c"}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != -1 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (-1, nil); received (%v, %v)",
			index,
			err,
		)
	}

}

/*
TestUtils_FindInSlice2
Description:

	This test verifies that FindInSlice returns the proper index in a slice
	of strings when the target string is in the slice.
*/
func TestUtils_FindInSlice2(t *testing.T) {
	// Constants
	x := "x"
	slice := []string{"a", "b", "c", "x"}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != 3 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (3, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice3
Description:

	This test verifies that the FindInSlice function returns the expected value
	when searching through a slice of ints for an int that is not in the slice.
*/
func TestUtils_FindInSlice3(t *testing.T) {
	// Constants
	x := 1
	slice := []int{2, 3, 4}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != -1 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (-1, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice4
Description:

	This test verifies that the FindInSlice function returns the expected value
	when searching through a slice of ints for an int that is in the slice.
*/
func TestUtils_FindInSlice4(t *testing.T) {
	// Constants
	x := 1
	slice := []int{2, 3, 4, 1}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != 3 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (3, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice5
Description:

	This test verifies that the FindInSlice function returns the expected value
	when searching through a slice of uint64s for a uint64 that is not in the slice.
*/
func TestUtils_FindInSlice5(t *testing.T) {
	// Constants
	x := uint64(1)
	slice := []uint64{2, 3, 4}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != -1 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (-1, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice6
Description:

	This test verifies that the FindInSlice function returns the expected value
	when searching through a slice of uint64s for a uint64 that is in the slice.
*/
func TestUtils_FindInSlice6(t *testing.T) {
	// Constants
	x := uint64(1)
	slice := []uint64{2, 3, 4, 1}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != 3 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (3, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice7
Description:

	This test verifies that the FindInSlice function returns the expected value
	when searching through a slice of Variables for a Variable that is not in the slice.
*/
func TestUtils_FindInSlice7(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	slice := []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable(), symbolic.NewVariable()}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != -1 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (-1, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice8
Description:

	This test verifies that the FindInSlice function returns the expected value
	when it is searching through a slice of Variables for a Variable that is in the slice.
*/
func TestUtils_FindInSlice8(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	slice := []symbolic.Variable{symbolic.NewVariable(), symbolic.NewVariable(), symbolic.NewVariable(), x}

	// Test
	if index, err := symbolic.FindInSlice(x, slice); index != 3 || err != nil {
		t.Errorf(
			"Expected FindInSlice to return (3, nil); received (%v, %v)",
			index,
			err,
		)
	}
}

/*
TestUtils_FindInSlice9
Description:

	This test verifies that the FindInSlice function returns an error when the
	input slice is not of a type that is supported.
*/
func TestUtils_FindInSlice9(t *testing.T) {
	// Constants
	x := 1
	slice := []float64{2, 3, 4, 1}

	// Test
	if _, err := symbolic.FindInSlice(x, slice); err == nil {
		t.Errorf(
			"Expected FindInSlice to return an error; received nil",
		)
	}
}

/*
TestUtils_FindInSlice10
Description:

	This test verifies that the FindInSlice function returns an error when the
	first input is a string and the second input is a slice of something else (in this case a slice of strings).
*/
func TestUtils_FindInSlice10(t *testing.T) {
	// Constants
	x := "x"
	slice := []int{2, 3, 4, 1}

	// Test
	_, err := symbolic.FindInSlice(x, slice)
	if err == nil {
		t.Errorf(
			"Expected FindInSlice to return an error; received nil",
		)
	} else {
		expectedError := fmt.Errorf(
			"the input slice is of type %T, but the element we're searching for is of type %T",
			slice,
			x,
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err.Error(),
			)
		}
	}
}

/*
TestUtils_FindInSlice11
Description:

	This test verifies that the FindInSlice function returns an error when the
	the first input is a Variable and the second input is a slice of something else (in this case a slice of strings).
*/
func TestUtils_FindInSlice11(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	slice := []int{2, 3, 4, 1}

	// Test
	_, err := symbolic.FindInSlice(x, slice)
	if err == nil {
		t.Errorf(
			"Expected FindInSlice to return an error; received nil",
		)
	} else {
		expectedError := fmt.Errorf(
			"the input slice is of type %T, but the element we're searching for is of type %T",
			slice,
			x,
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err.Error(),
			)
		}
	}
}

/*
TestUtils_FindInSlice12
Description:

	This test verifies that the FindInSlice function returns an error when the
	the first input is an error (an unsupported type). The second input can be anything else.
*/
func TestUtils_FindInSlice12(t *testing.T) {
	// Constants
	x := fmt.Errorf("error")
	slice := []int{2, 3, 4, 1}

	// Test
	_, err := symbolic.FindInSlice(x, slice)
	if err == nil {
		t.Errorf(
			"Expected FindInSlice to return an error; received nil",
		)
	} else {
		allowedTypes := []string{
			"string", "int", "uint64", "Variable",
		}
		expectedError := fmt.Errorf(
			"the FindInSlice() function was only defined for types %v, not type %T:",
			allowedTypes,
			x,
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err.Error(),
			)
		}
	}
}

/*
TestUtils_Unique1
Description:

	This test verifies that when the input slice is empty, the output slice is
	also empty.
*/
func TestUtils_Unique1(t *testing.T) {
	// Constants
	slice := []uint64{}

	// Test
	if len(symbolic.Unique(slice)) != 0 {
		t.Errorf(
			"Expected Unique to return an empty slice; received %v",
			symbolic.Unique(slice),
		)
	}
}

/*
TestUtils_Unique2
Description:

	This test verifies that the unique function returns a slice of length 1
	when the input slice is a slice of uint64's with the same value (13).
*/
func TestUtils_Unique2(t *testing.T) {
	// Constants
	slice := []uint64{13, 13, 13, 13}

	// Test
	if len(symbolic.Unique(slice)) != 1 {
		t.Errorf(
			"Expected Unique to return a slice of length 1; received %v",
			symbolic.Unique(slice),
		)
	}
}

/*
TestUtils_Unique3
Description:

	This test verifies that the unique function returns a slice of length 1
	when the input slice has length 1.
*/
func TestUtils_Unique3(t *testing.T) {
	// Constants
	slice := []uint64{13}

	// Test
	if len(symbolic.Unique(slice)) != 1 {
		t.Errorf(
			"Expected Unique to return a slice of length 1; received %v",
			symbolic.Unique(slice),
		)
	}
}

/*
TestUtils_Unique4
Description:

	This test verifies that the unique function returns a slice of length 3
	when the input slice has length 3 and all elements are unique.
*/
func TestUtils_Unique4(t *testing.T) {
	// Constants
	slice := []uint64{13, 14, 15}

	// Test
	if len(symbolic.Unique(slice)) != 3 {
		t.Errorf(
			"Expected Unique to return a slice of length 3; received %v",
			symbolic.Unique(slice),
		)
	}
}

/*
TestUtils_CheckSubstitutionMap1
Description:

	This test verifies that the CheckSubstitutionMap function returns an error
	when the input map contains a variable that is not well-defined.
*/
func TestUtils_CheckSubstitutionMap1(t *testing.T) {
	// Constants
	badVar := symbolic.Variable{2, -1, -2, symbolic.Binary, "Russ"}
	varMap := map[symbolic.Variable]symbolic.Expression{
		symbolic.NewVariable(): symbolic.K(3),
		badVar:                 symbolic.K(4),
	}

	// Test
	err := symbolic.CheckSubstitutionMap(varMap)
	if err == nil {
		t.Errorf(
			"Expected CheckSubstitutionMap to return an error; received nil",
		)
	} else {
		expectedError := fmt.Errorf(
			"key %v in the substitution map is not a valid variable: %v",
			badVar,
			badVar.Check(),
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err,
			)
		}
	}
}

/*
TestUtils_CheckSubstitutionMap2
Description:

	This test verifies that the CheckSubstitutionMap function returns an error
	when the input map contains a value/mapped item that is not well-defined.
*/
func TestUtils_CheckSubstitutionMap2(t *testing.T) {
	// Constants
	goodVar := symbolic.NewVariable()
	badVar := symbolic.Variable{2, -1, -2, symbolic.Binary, "Russ"}
	varMap := map[symbolic.Variable]symbolic.Expression{
		symbolic.NewVariable(): symbolic.K(3),
		goodVar:                badVar,
	}

	// Test
	err := symbolic.CheckSubstitutionMap(varMap)
	if err == nil {
		t.Errorf(
			"Expected CheckSubstitutionMap to return an error; received nil",
		)
	} else {
		expectedError := fmt.Errorf(
			"value %v in the substitution map[%v] is not a valid expression: %v",
			badVar,
			goodVar,
			badVar.Check(),
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err,
			)
		}
	}
}

/*
TestUtils_CheckSubstitutionMap3
Description:

	This test verifies that the CheckSubstitutionMap function returns an error
	when the input map contains a value/mapped item that is not a scalar expression.
*/
func TestUtils_CheckSubstitutionMap3(t *testing.T) {
	// Constants
	goodVar := symbolic.NewVariable()
	badVar := symbolic.NewVariableMatrix(2, 2)
	varMap := map[symbolic.Variable]symbolic.Expression{
		symbolic.NewVariable(): symbolic.K(3),
		goodVar:                badVar,
	}

	// Test
	err := symbolic.CheckSubstitutionMap(varMap)
	if err == nil {
		t.Errorf(
			"Expected CheckSubstitutionMap to return an error; received nil",
		)
	} else {
		expectedError := fmt.Errorf(
			"value %v in the substitution map[%v] is not a scalar expression (received %T)",
			badVar,
			goodVar,
			badVar,
		)
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error message to be '%v'; received %v",
				expectedError,
				err,
			)
		}
	}
}

/*
TestUtils_CheckSubstitutionMap4
Description:

	This test verifies that the CheckSubstitutionMap function returns no error
	when the input map is well-defined (has valid variable keys and its values are all valid scalar expresssions).
*/
func TestUtils_CheckSubstitutionMap4(t *testing.T) {
	// Constants
	varMap := map[symbolic.Variable]symbolic.Expression{
		symbolic.NewVariable(): symbolic.K(3),
		symbolic.NewVariable(): symbolic.K(4),
	}

	// Test
	if err := symbolic.CheckSubstitutionMap(varMap); err != nil {
		t.Errorf(
			"Expected CheckSubstitutionMap to return nil; received %v",
			err,
		)
	}
}
