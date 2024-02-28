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
TestPolynomial_Plus2
Description:

	Verifies that the Polynomial.Plus method panics
	when the polynomial used to call it is not well-defined.
*/
func TestPolynomial_Plus2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Plus to panic when called on an undefined polynomial; received nil",
			)
		}
	}()

	// Call the Plus method
	p1.Plus(p1)
}

/*
TestPolynomial_Plus3
Description:

	Verifies that the sum of a well-defined polynomial
	and a well-defined variable (one that is in the polynomial)
	returns a polynomial with the same number of monomials.
*/
func TestPolynomial_Plus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     2.0,
				VariableFactors: []symbolic.Variable{v1},
				Exponents:       []int{1},
			},
		},
	}

	// Test
	sum := p1.Plus(v1)
	if len(sum.(symbolic.Polynomial).Monomials) != 1 {
		t.Errorf(
			"expected %v + %v to have 1 monomial; received %v",
			p1,
			v1,
			len(sum.(symbolic.Polynomial).Monomials),
		)
	}

	// Verify that the coefficient of the sum's monomial
	// is 1.0 more than the original
	if sum.(symbolic.Polynomial).Monomials[0].Coefficient != 3.0 {
		t.Errorf(
			"expected %v + %v to have coefficient 3.0; received %v",
			p1,
			v1,
			sum.(symbolic.Polynomial).Monomials[0].Coefficient,
		)
	}
}

/*
TestPolynomial_Plus4
Description:

	Verifies that the sum of a well-defined polynomial
	and an object that is not an expression (a string) panics.
*/
func TestPolynomial_Plus4(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Plus to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the Plus method
	p1.Plus("string")
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

/*
TestPolynomial_LinearCoeff3
Description:

	Verifies that the Polynomial.LinearCoeff method panics
	when the underlying polynomial is not well-defined.
*/
func TestPolynomial_LinearCoeff3(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected LinearCoeff to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the LinearCoeff method
	p1.LinearCoeff()
}

/*
TestPolynomial_LinearCoeff4
Description:

	Verifies that the Polynomial.LinearCoeff method panics
	when the underlying polynomial is a constant AND
	there are no inputs provided to LinearCoeff().
*/
func TestPolynomial_LinearCoeff4(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 3.14, VariableFactors: []symbolic.Variable{}, Exponents: []int{}},
		},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected LinearCoeff to panic when called with a constant polynomial; received nil",
			)
		}
	}()

	// Call the LinearCoeff method
	p1.LinearCoeff()
}

/*
TestPolynomial_LinearCoeff5
Description:

	Verifies that the Polynomial.LinearCoeff method returns the correct
	linear coefficients (all zeros) when called with a polynomial that
	is a constant and a slice of variables is provided.
*/
func TestPolynomial_LinearCoeff5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 3.14, VariableFactors: []symbolic.Variable{}, Exponents: []int{}},
		},
	}

	// Test
	coeff := p1.LinearCoeff([]symbolic.Variable{v1, v2})

	// Tests that the length of the coefficient vector is the same as the number of variables
	if coeff.Len() != 2 {
		t.Errorf(
			"expected LinearCoeff to return a vector of length 2; received %v",
			coeff,
		)
	}

	// Check that all elements are zero
	for i := 0; i < coeff.Len(); i++ {
		if coeff.AtVec(i) != 0 {
			t.Errorf(
				"expected LinearCoeff to return a vector of all zeros; received %v",
				coeff,
			)
		}
	}
}

/*
TestPolynomial_LinearCoeff6
Description:

	Verifies that the Polynomial.LinearCoeff method returns the correct
	linear coefficients when called with a polynomial that contains SOME linear terms.
*/
func TestPolynomial_LinearCoeff6(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v3}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 3.14, VariableFactors: []symbolic.Variable{v4}, Exponents: []int{2}},
		},
	}

	// Test
	coeff := p1.LinearCoeff([]symbolic.Variable{v1, v2, v3, v4})

	// Tests that the length of the coefficient vector is the same as the number of variables
	if coeff.Len() != 4 {
		t.Errorf(
			"expected LinearCoeff to return a vector of length 3; received %v",
			coeff,
		)
	}

	// Check that all elements are zero
	for i := 0; i < coeff.Len(); i++ {
		if i < 3 {
			if coeff.AtVec(i) != 1 {
				t.Errorf(
					"expected LinearCoeff to return a vector of all ones; received %v",
					coeff,
				)
			}
		} else {
			if coeff.AtVec(i) != 0 {
				t.Errorf(
					"expected LinearCoeff to return a vector of all zeros; received %v",
					coeff,
				)
			}
		}
	}
}

/*
TestPolynomial_Multiply1
Description:

	Verifies that the Polynomial.Multiply method panics when called with an invalid polynomial.
*/
func TestPolynomial_Multiply1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Multiply to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the Multiply method
	p1.Multiply(p1)
}

/*
TestPolynomial_Multiply2
Description:

	Verifies that the Polynomial.Multiply method panics when called with a non-expression.
*/
func TestPolynomial_Multiply2(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Multiply to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the Multiply method
	p1.Multiply("string")
}

/*
TestPolynomial_Multiply3
Description:

	Verifies that the Polynomial.Multiply method returns a polynomial with the correct number of monomials
	after multiplying it with a polynomial.
*/
func TestPolynomial_Multiply3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 2, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{2}},
		},
	}

	p2 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{3}},
		},
	}

	// Test
	product := p1.Multiply(p2)
	if len(product.(symbolic.Polynomial).Monomials) != len(p1.Monomials)*len(p2.Monomials) {
		t.Errorf(
			"expected %v * %v to have %v monomial; received %v",
			p1,
			p2,
			len(p1.Monomials)*len(p2.Monomials),
			len(product.(symbolic.Polynomial).Monomials),
		)
	}
}

/*
TestPolynomial_LessEq1
Description:

	Verifies that the Polynomial.LessEq method panics when called
	with polynomial that is not well-defined.
*/
func TestPolynomial_LessEq1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected LessEq to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the LessEq method
	p1.LessEq(p1)
}

/*
TestPolynomial_LessEq2
Description:

	Verifies that the Polynomial.LessEq method returns a valid
	ScalarConstraint when called with a valid polynomial
	and a float64.
*/
func TestPolynomial_LessEq2(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	lessEq := p1.LessEq(3.14)

	// Verify that the constraint is a scalar constraint
	if _, tf := lessEq.(symbolic.ScalarConstraint); !tf {
		t.Errorf(
			"expected %v <= %v to return a ScalarConstraint; received %T",
			p1,
			3.14,
			lessEq,
		)
	}
}

/*
TestPolynomial_GreaterEq1
Description:

	Verifies that the Polynomial.GreaterEq method panics when called
	with a well-defined polynomial and a variable that is not well-defined.
*/
func TestPolynomial_GreaterEq1(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected GreaterEq to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the GreaterEq method
	p1.GreaterEq(symbolic.Variable{})
}

/*
TestPolynomial_GreaterEq2
Description:

	Verifies that the Polynomial.GreaterEq method returns a valid
	ScalarConstraint when called with a well-defined polynomial
	and a well-defined monomial.
*/
func TestPolynomial_GreaterEq2(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	m1 := symbolic.NewVariable().ToMonomial()

	// Test
	greaterEq := p1.GreaterEq(m1)

	// Verify that the constraint is a scalar constraint
	if _, tf := greaterEq.(symbolic.ScalarConstraint); !tf {
		t.Errorf(
			"expected %v >= %v to return a ScalarConstraint; received %T",
			p1,
			m1,
			greaterEq,
		)
	}
}

/*
TestPolynomial_Eq1
Description:

	Verifies that the Polynomial.Eq method panics when called
	with a well-defined polynomial and a variable that is not
	an expression (in this case, it's a string).
*/
func TestPolynomial_Eq1(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Eq to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the Eq method
	p1.Eq("string")
}

/*
TestPolynomial_Eq2
Description:

	Verifies that the Polynomial.Eq method returns a valid
	ScalarConstraint when called with a well-defined polynomial
	and a well-defined polynomial.
*/
func TestPolynomial_Eq2(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	p2 := symbolic.NewVariable().ToPolynomial()

	// Test
	eq := p1.Eq(p2)

	// Verify that the constraint is a scalar constraint
	if _, tf := eq.(symbolic.ScalarConstraint); !tf {
		t.Errorf(
			"expected %v == %v to return a ScalarConstraint; received %T",
			p1,
			p2,
			eq,
		)
	}
}

/*
TestPolynomial_Comparison1
Description:

	Verifies that the Polynomial.Comparison method
	returns a valid scalar constraint when a well-defined
	polynomial is compared to a well-defined symbolic.Variable
*/
func TestPolynomial_Comparison1(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	v1 := symbolic.NewVariable()

	// Test
	comp := p1.Comparison(v1, symbolic.SenseEqual)

	// Verify that the constraint is a scalar constraint
	if _, tf := comp.(symbolic.ScalarConstraint); !tf {
		t.Errorf(
			"expected %v == %v to return a ScalarConstraint; received %T",
			p1,
			v1,
			comp,
		)
	}
}

/*
TestPolynomial_Constant1
Description:

	Verifies that the Polynomial.Constant method returns a valid
	nonzero constant when a polynomial with two monomials (one of
	which is a constant) is used.
*/
func TestPolynomial_Constant1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{}, Exponents: []int{}},
			symbolic.Monomial{Coefficient: 2, VariableFactors: []symbolic.Variable{symbolic.NewVariable()}, Exponents: []int{1}},
		},
	}

	// Test
	constant := p1.Constant()
	if constant != 1 {
		t.Errorf(
			"expected %v to have constant 1; received %v",
			p1,
			constant,
		)
	}
}

/*
TestPolynomial_IsLinear1
Description:

	Verifies that the Polynomial.IsLinear method returns true
	when a polynomial is a sum of three variables.
*/
func TestPolynomial_IsLinear1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v3}, Exponents: []int{1}},
		},
	}

	// Test
	if !symbolic.IsLinear(p1) {
		t.Errorf(
			"expected %v to be linear; received %v",
			p1,
			symbolic.IsLinear(p1),
		)
	}
}
