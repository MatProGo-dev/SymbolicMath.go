package symbolic

/*
utils_test.go
Description:
	Tests some of the functions in the utils.go file.
*/

import (
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
