package symbolic_test

/*
constant_vector_test.go
Description:
	Tests for the functions mentioned in the constant_vector.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestConstantVector_Len1
Description:

	Tests that an empty kvector has length 0.
*/
func TestConstantVector_Len1(t *testing.T) {
	// Constants
	kv := symbolic.KVector{}

	// Test
	if kv.Len() != 0 {
		t.Errorf(
			"Expected kv.Len() to be 0; received %v",
			kv.Len(),
		)
	}
}

/*
TestConstantVector_Len2
Description:

	Tests that an vector with over 100 elements has the proper length.
*/
func TestConstantVector_Len2(t *testing.T) {
	// Constants
	N := 101
	kv := symbolic.KVector(symbolic.OnesVector(N))

	// Test
	if kv.Len() != N {
		t.Errorf(
			"Expected kv.Len() to be %v; received %v",
			N,
			kv.Len(),
		)
	}
}

/*
TestConstantVector_Check1
Description:

	This test verifies that Check returns nil for a well-defined constant vector.
*/
func TestConstantVector_Check1(t *testing.T) {
	// Constants
	kv := symbolic.KVector(symbolic.OnesVector(3))

	// Test
	if kv.Check() != nil {
		t.Errorf(
			"Expected kv.Check() to return nil; received %v",
			kv.Check(),
		)
	}
}

/*
TestConstantVector_AtVec1
Description:

	Verifies that this function returns the correct value when an in-bounds value
	is given for the vector.
*/
func TestConstantVector_AtVec1(t *testing.T) {
	// Constants
	kv := symbolic.KVector(symbolic.OnesVector(3))

	// Test
	if float64(kv.AtVec(0).(symbolic.K)) != 1 {
		t.Errorf(
			"Expected kv.AtVec(0) to be 1; received %v",
			kv.AtVec(0),
		)
	}
}

/*
TestConstantVector_AtVec2
Description:

	Verifies that the AtVec function panics when an out-of-bounds value is given.
*/
func TestConstantVector_AtVec2(t *testing.T) {
	// Constants
	kv := symbolic.KVector(symbolic.OnesVector(3))

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected kv.AtVec(3) to panic; received %v",
				kv.AtVec(3),
			)
		}
	}()

	kv.AtVec(3)
}
