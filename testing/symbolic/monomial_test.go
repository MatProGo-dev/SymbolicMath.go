package symbolic

/*
monomial_test.go
Description:
	Tests the monomial object.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestMonomial_Check1
Description:

	Tests that the Check() method properly catches an improperly initialized
	monomial.
*/
func TestMonomial_Check1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	err := m1.Check()
	if err.Error() != fmt.Errorf(
		"the number of degrees (%v) does not match the number of variables (%v)",
		len(m1.Exponents),
		len(m1.VariableFactors),
	).Error() {
		t.Errorf(
			"expected Check() to return false; received %v",
			err,
		)
	}
}

/*
TestMonomial_Check2
Description:

	Verifies that the Check() method returns nil when a constant is
	given as a monomial.
*/
func TestMonomial_Check2(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}

	// Test
	if m1.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			m1.Check(),
		)
	}
}

/*
TestMonomial_Check3
Description:

	Verifies that the Check() method returns nil when a complex monomial
	is given.
*/
func TestMonomial_Check3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 2},
	}

	// Test
	if m1.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			m1.Check(),
		)
	}
}

/*
TestMonomial_Plus1
Description:

	Verifies that the addition of a monomial and a constant is still a monomial,
	if the monomial is a constant.
*/
func TestMonomial_Plus1(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}
	f1 := 3.14

	// Compute Sum
	sum := m1.Plus(f1)

	// Verify that the sum is a monomial
	_, tf := sum.(symbolic.K)
	if !tf {
		t.Errorf(
			"expected sum to be a K; received %T",
			sum,
		)
	}

}

/*
TestMonomial_Plus2
Description:

	Verifies that the sum of a monomial and a constant is a polynomial,
*/
func TestMonomial_Plus2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 7},
	}
	f1 := 3.14

	// Compute Sum
	sum := m1.Plus(f1)

	// Verify that the sum is a monomial
	sumAsP, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected sum to be a polynomial; received %T",
			sum,
		)
	}

	// Test that the sum is a constant
	if len(sumAsP.Monomials) != 2 {
		t.Errorf(
			"expected sum to have 2 monomials; received %v",
			len(sumAsP.Monomials),
		)
	}

	// Check that the monomial is well formed with Check
	if sumAsP.Check() != nil {
		t.Errorf(
			"expected sum to be well formed; received %v",
			sumAsP.Check(),
		)
	}

}

/*
TestMonomial_Plus3
Description:

	Tests that the addition of a monomial with a variable is a monomial,
	when the monomial contains just the variable.
*/
func TestMonomial_Plus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute Sum
	sum := m2.Plus(v1)

	// Verify that the sum is a monomial
	sumAsM, tf := sum.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected sum to be a monomial; received %T",
			sum,
		)
	}

	// Check that the monomial is well formed with Check
	if sumAsM.Check() != nil {
		t.Errorf(
			"expected sum to be well formed; received %v",
			sumAsM.Check(),
		)
	}

	// Check that the variable in the sum is the same as the one v1
	if sumAsM.VariableFactors[0].ID != v1.ID {
		t.Errorf(
			"expected sum to have the same variable as v1; received %v",
			sumAsM.VariableFactors[0],
		)
	}
}

/*
TestMonomial_Plus4
Description:

	Verifies that the sum of a monomial of one variable and a variable
	is a polynomial, when the variables are different.
*/
func TestMonomial_Plus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute Sum
	sum := m2.Plus(v2)

	// Verify that the sum is a polynomial
	sumAsP, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected sum to be a polynomial; received %T",
			sum,
		)
	}

	// Check that the polynomial is well formed with Check
	if sumAsP.Check() != nil {
		t.Errorf(
			"expected sum to be well formed; received %v",
			sumAsP.Check(),
		)
	}

	// Check that the polynomial has 2 monomials
	if len(sumAsP.Monomials) != 2 {
		t.Errorf(
			"expected sum to have 2 monomials; received %v",
			len(sumAsP.Monomials),
		)
	}
}
