package getKVector_test

import (
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"testing"
)

/*
from_test.go
Description:

	This file contains tests for the functions in from.go.
*/

/*
TestFrom1
Description:

	Tests that the From() method properly panics when given
	an unsupported input (string).
*/
func TestFrom1(t *testing.T) {
	// Constants
	var input string = "This is a test string."

	// Run test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("From() did not panic when given an unsupported input (string).")
		}
	}()

	// Run function
	getKVector.From(input)
}

/*
TestFrom2
Description:

	Tests that the From() method properly returns a KVector
	when given a slice of loats.
*/
func TestFrom2(t *testing.T) {
	// Constants
	var input []float64 = []float64{1, 2, 3}

	// Run function
	kv1 := getKVector.From(input)
	if kv1.Len() != len(input) {
		t.Errorf("From() did not properly convert a slice of floats to a KVector.")
	}
}
