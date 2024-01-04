package symbolic_test

/*
constr_sense_test.go
Description:
	Tests for the functions mentioned in the constr_sense.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestConstrSense_String1
Description:

	Tests that the string of a SenseEqual object is equal to "=".
*/
func TestConstrSense_String1(t *testing.T) {
	// Constants
	sense0 := symbolic.SenseEqual

	// Test
	if sense0.String() != "=" {
		t.Errorf(
			"Expected sense.String() to be \"=\"; received %v",
			sense0.String(),
		)
	}
}

/*
TestConstrSense_String2
Description:

	Tests that the string of a SenseLessThanEqual object is equal to "<".
*/
func TestConstrSense_String2(t *testing.T) {
	// Constants
	var sense0 symbolic.ConstrSense = symbolic.SenseLessThanEqual

	// Test
	if sense0.String() != "<=" {
		t.Errorf(
			"Expected sense.String() to be \"<\"; received %v",
			sense0.String(),
		)
	}
}

/*
TestConstrSense_String3
Description:

	Tests that the string of a SenseGreaterThanEqual object is equal to ">".
*/
func TestConstrSense_String3(t *testing.T) {
	// Constants
	var sense0 symbolic.ConstrSense = symbolic.SenseGreaterThanEqual

	// Test
	if sense0.String() != ">=" {
		t.Errorf(
			"Expected sense.String() to be \">\"; received %v",
			sense0.String(),
		)
	}
}
