package symbolic_test

/*
polyomial_test.go
Description:
	Tests the functions implemented in the polynomial.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
TestPolynomial_Check1
Description:

	Verifies that a polynomial with NO monomials given is not valid.
*/
func TestPolynomial_Check1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	err := p1.Check()
	if err == nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			p1.Check(),
		)
	} else {
		if err.Error() != "polynomial has no monomials" {
			t.Errorf(
				"expected Check() to return nil; received %v",
				p1.Check(),
			)
		}
	}
}

/*
TestPolynomial_Check2
Description:

	Tests that a polynomial with a single invalid monomial is invalid.
*/
func TestPolynomial_Check2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m1},
	}

	// Test
	err := p1.Check()
	if err == nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			p1.Check(),
		)
	} else {
		if !strings.Contains(
			err.Error(),
			"the number of degrees (2) does not match the number of variables (1)",
		) {
			t.Errorf(
				"expected Check() to return nil; received %v",
				p1.Check(),
			)
		}
	}
}

/*
TestPolynomial_Check3
Description:

	Tests that a polynomial with one valid monomial is valid.
*/
func TestPolynomial_Check3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m1},
	}

	// Test
	err := p1.Check()
	if err != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			p1.Check(),
		)
	}
}

/*
TestPolynomial_Variables1
Description:

	Verifies that a polynomial containing a single monomial (that represents a constant)
	returns no variables.
*/
func TestPolynomial_Variables1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{},
				Exponents:       []int{},
			},
		},
	}

	// Test
	if symbolic.NumVariables(p1) != 0 {
		t.Errorf(
			"expected %v to have no variables; received %v",
			p1,
			p1.Variables(),
		)
	}
}

/*
TestPolynomial_Variables2
Description:

	Verifies that a polynomial containing a single monomial
	(that represents a variable) returns the correct variable.
*/
func TestPolynomial_Variables2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v1},
				Exponents:       []int{1},
			},
		},
	}

	// Test
	if symbolic.NumVariables(p1) != 1 {
		t.Errorf(
			"expected %v to have 1 variable; received %v",
			p1,
			p1.Variables(),
		)
	}

	if p1.Variables()[0].ID != v1.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v1,
			p1.Variables()[0],
		)
	}
}

/*
TestPolynomial_Variables3
Description:

	Verifies that a polynomial containing a single monomial
	of three variables returns the correct variables.
*/
func TestPolynomial_Variables3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v1, v2, v3},
				Exponents:       []int{1, 2, 3},
			},
		},
	}

	// Test
	if symbolic.NumVariables(p1) != 3 {
		t.Errorf(
			"expected %v to have 3 variables; received %v",
			p1,
			p1.Variables(),
		)
	}

	if p1.Variables()[0].ID != v1.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v1,
			p1.Variables()[0],
		)
	}

	if p1.Variables()[1].ID != v2.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v2,
			p1.Variables()[1],
		)
	}

	if p1.Variables()[2].ID != v3.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v3,
			p1.Variables()[2],
		)
	}
}

/*
TestPolynomial_Variables4
Description:

	Verifies that a polynomial containing two monomials
	each composed of the same three variables returns the correct variables.
*/
func TestPolynomial_Variables4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v1, v2, v3},
				Exponents:       []int{1, 2, 3},
			},
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v1, v2, v3},
				Exponents:       []int{5, 6, 7},
			},
		},
	}

	// Test
	if symbolic.NumVariables(p1) != 3 {
		t.Errorf(
			"expected %v to have 3 variables; received %v",
			p1,
			p1.Variables(),
		)
	}

	if p1.Variables()[0].ID != v1.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v1,
			p1.Variables()[0],
		)
	}

	if p1.Variables()[1].ID != v2.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v2,
			p1.Variables()[1],
		)
	}

	if p1.Variables()[2].ID != v3.ID {
		t.Errorf(
			"expected %v to have variable %v; received %v",
			p1,
			v3,
			p1.Variables()[2],
		)
	}
}

/*
TestPolynomial_Check4
Description:

	Tests that a polynomial containing one valid monomial and one invalid monomial
	is invalid.
*/
func TestPolynomial_Check4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1, 2},
	}

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m1, m2},
	}

	// Test
	err := p1.Check()
	if err == nil {
		t.Errorf(
			"expected Check() to return error; received none",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			"the number of degrees (2) does not match the number of variables (1)",
		) {
			t.Errorf(
				"expected Check() to return error containing %v; received \"%v\"",
				"the number of degrees (2) does not match the number of variables (1)",
				p1.Check(),
			)
		}
	}
}

/*
TestPolynomial_Dims1
Description:

	Verifies that a polynomial composed of three monomials
	each with different variables returns the correct dimensions (1,1) because
	this is a scalar.
*/
func TestPolynomial_Dims1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v1},
				Exponents:       []int{1},
			},
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v2},
				Exponents:       []int{1},
			},
			symbolic.Monomial{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{v3},
				Exponents:       []int{1},
			},
		},
	}

	// Test
	if p1.Dims()[0] != 1 || p1.Dims()[1] != 1 {
		t.Errorf(
			"expected %v to have dimensions [1,1]; received %v",
			p1,
			p1.Dims(),
		)
	}
}

/*
TestPolynomial_Plus1
Description:

	Verifies that the sum of a polynomial and a constant
	does not change the number of monomials when the polynomial contains
	a constant already.
*/
func TestPolynomial_Plus1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{},
				Exponents:       []int{},
			},
		},
	}

	k1 := symbolic.K(2.71)

	// Test
	sum := p1.Plus(k1)
	if len(sum.(symbolic.Polynomial).Monomials) != 1 {
		t.Errorf(
			"expected %v + %v to have 1 monomial; received %v",
			p1,
			k1,
			len(sum.(symbolic.Polynomial).Monomials),
		)
	}

	// Verify that the sum is a polynomial
	if _, tf := sum.(symbolic.Polynomial); !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			p1,
			k1,
			sum,
		)
	}

	//Verify that the sum's value matches what we expect
	if sum.(symbolic.Polynomial).Monomials[0].Coefficient != 5.85 {
		t.Errorf(
			"expected %v + %v to have coefficient 5.85; received %v",
			p1,
			k1,
			sum.(symbolic.Polynomial).Monomials[0].Coefficient,
		)
	}
}

/*
TestPolynomial_VariableMonomialIndex1
Description:

	Tests that VariableMonomialIndex returns the correct index when the
	variable is in the monomial.
*/
func TestPolynomial_VariableMonomialIndex1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 2},
	}

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m1, v2.ToMonomial()},
	}

	// Test that the index is 1
	if p1.VariableMonomialIndex(v2) != 1 {
		t.Errorf(
			"expected index of %v to be 1; received %v",
			v2,
			p1.VariableMonomialIndex(v2),
		)
	}

}

/*
TestPolynomial_LinearCoeff1
Description:

	This test verifies that the LinearCoeff function returns a VecDense of all zeros
	when the provided polynomial is a constant.
*/
func TestPolynomial_LinearCoeff1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{},
				Exponents:       []int{},
			},
		},
	}

	// Check that a panic is thrown when the LinearCoeff method is called
	// on such a polynomial
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected LinearCoeff to panic when called on a constant polynomial; received nil",
			)
		}

		// Check that the recovered value is an error
		rAsErr, tf := r.(error)
		if !tf {
			t.Errorf(
				"expected LinearCoeff to panic with an error; received %v of type %T",
				r, r,
			)
		}

		// Check that the panic message is correct
		if rAsErr.Error() != (smErrors.CanNotGetLinearCoeffOfConstantError{p1}).Error() {
			t.Errorf(
				"expected LinearCoeff to panic with an error %v; received %v",
				smErrors.CanNotGetLinearCoeffOfConstantError{p1},
				rAsErr,
			)
		}
	}()

	// Check that coeff has the same length as the number of variables in p1
	p1.LinearCoeff()
}

/*
TestPolynomial_LinearCoeff2
Description:

	This test verifies that the LinearCoeff function applied to a polynomial with all
	quadratic terms is a vector of all zeros.
*/
func TestPolynomial_LinearCoeff2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{2}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{2}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1, v3}, Exponents: []int{1, 1}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v2, v3}, Exponents: []int{1, 1}},
		},
	}

	// Collect coefficient and check each of its elements
	coeff := p1.LinearCoeff()
	for i := 0; i < coeff.Len(); i++ {
		if coeff.At(i, 0) != 0 {
			t.Errorf(
				"expected LinearCoeff to return a vector of all zeros; received %v",
				coeff,
			)
		}
	}
}
