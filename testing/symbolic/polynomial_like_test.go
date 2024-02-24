package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
polynomial_like_test.go
Description:

	Tests for the functions mentioned in the polynomial_like.go file.
*/

/*
TestPolynomialLike_IsPolynomialLike1
Description:

	Tests to make sure that an expression object is identified as not a polynomial like object.
*/
func TestPolynomialLike_IsPolynomialLike1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.IsPolynomialLike(x) {
		t.Errorf(
			"Expected IsPolynomialLike(%T) to be false; received true",
			x,
		)
	}
}
