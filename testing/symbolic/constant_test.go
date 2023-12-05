package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
constant_test.go
Description:
	Tests the constant object.
*/

/*
TestConstant_ToMonomial1
Description:

	Tests that this works correctly.
*/
func TestConstant_ToMonomial1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)

	// Test
	m1 := k1.ToMonomial()
	if float64(k1) != m1.Coefficient {
		t.Errorf(
			"expected monomial coefficient to be %v; received %v",
			k1,
			m1,
		)
	}

	if len(m1.VariableFactors) != 0 {
		t.Errorf(
			"expected there to be 0 variables in monomial; received %v",
			len(m1.VariableFactors),
		)
	}

	if len(m1.Degrees) != 0 {
		t.Errorf(
			"expected there to be 0 degrees in monomial; received %v",
			len(m1.Degrees),
		)
	}

}

/*
TestConstant_Check1
Description:

	The check function should always produce no errors.
*/
func TestConstant_Check1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)

	// Test
	if k1.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			k1.Check(),
		)
	}
}

/*
TestConstant_Variables1
Description:

	The Variables() method should return an empty list.
*/
func TestConstant_Variables1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)

	// Test
	if len(k1.Variables()) != 0 {
		t.Errorf(
			"expected Variables() to return empty list; received %v",
			symbolic.NumVariables(k1),
		)
	}
}

/*
TestConstant_Constant1
Description:

	Tests that the constant method returns the float version of the
	constant.
*/
func TestConstant_Constant1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)

	// Test
	if float64(k1) != k1.Constant() {
		t.Errorf(
			"expected constant to be %v; received %v",
			float64(k1),
			k1.Constant(),
		)
	}
}

/*
TestConstant_Plus1
Description:

	Tests that the Plus() method properly adds with a symbolic.K.
*/
func TestConstant_Plus1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	k2 := symbolic.K(2.71)

	// Test
	if float64(k1.Plus(k2).(symbolic.K)) != 5.85 {
		t.Errorf(
			"expected constant to be %v; received %v",
			5.85,
			k1.Plus(k2),
		)
	}
}

/*
TestConstant_Plus1
Description:

	Tests that the Plus() method properly adds with a float.
*/
func TestConstant_Plus2(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	f2 := 2.71

	// Test
	if float64(k1.Plus(f2).(symbolic.K)) != 5.85 {
		t.Errorf(
			"expected constant to be %v; received %v",
			5.85,
			k1.Plus(f2),
		)
	}
}
