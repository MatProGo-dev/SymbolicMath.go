package symbolic_test

/*
polyomial_test.go
Description:
	Tests the functions implemented in the polynomial.go file.
*/

import (
	"reflect"
	"strings"
	"testing"

	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
TestPolynomial_Plus5
Description:

	Verifies that calling Plus() using a well-defined polynomial
	and a not well-defined expression causes a panic to occur.
*/
func TestPolynomial_Plus5(t *testing.T) {
	// Setup
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
	p1.Plus(symbolic.Polynomial{})

	t.Errorf("expected Plus to panic when called with a non-expression; received nil")
}

/*
TestPolynomial_Plus6
Description:

	This function verifies that the plus method returns the correct
	output type when called with a polynomial vector (of
	length 1) and a constant.
*/
func TestPolynomial_Plus6(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	k1 := getKVector.From([]float64{3.14})

	// Test
	sum := p1.Plus(k1)

	// Check that the output is a scalar polynomial
	if _, ok := sum.(symbolic.Polynomial); !ok {
		t.Errorf(
			"expected Plus to return a polynomial; received %T",
			sum,
		)
	}
}

/*
TestPolynomial_Plus7
Description:

	Verifies that the sum of a polynomial and a well-defined polynomial
	vector of length 4.
*/
func TestPolynomial_Plus7(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	p2 := symbolic.NewVariable().ToPolynomial()
	p3 := symbolic.NewVariable().ToPolynomial()
	p4 := symbolic.NewVariable().ToPolynomial()

	pv5 := symbolic.PolynomialVector{p2, p3, p4}

	// Test
	sum := p1.Plus(pv5)

	// Check that the output is a polynomial vector
	sumAsPV, ok := sum.(symbolic.PolynomialVector)
	if !ok {
		t.Errorf(
			"expected Plus to return a polynomial; received %T",
			sum,
		)
	}

	// Check that the length of the output is 3
	if len(sumAsPV) != 3 {
		t.Errorf(
			"expected Plus to return a polynomial vector of length 4; received %v",
			len(sumAsPV),
		)
	}
}

/*
TestPolynomial_Plus8
Description:

	Verifies that the sum of a polynomial and a well-defined KMatrix
	(of dimension 1 x 1) returns a polynomial.
*/
func TestPolynomial_Plus8(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	km1 := getKMatrix.From([][]float64{{3.14}})

	// Test
	sum := p1.Plus(km1)

	// Check that the output is a polynomial
	if _, ok := sum.(symbolic.Polynomial); !ok {
		t.Errorf(
			"expected Plus to return a polynomial; received %T",
			sum,
		)
	}
}

/*
TestPolynomial_Plus9
Description:

	Verifies that the sum of a polynomial and a well-defined KMatrix
	(of dimension 1 x 4) returns a polynomial matrix of dimension 1 x 4.
*/
func TestPolynomial_Plus9(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	km1 := getKMatrix.From([][]float64{{3.14, 2.71, 1.0, 0.0}})

	// Test
	sum := p1.Plus(km1)

	// Check that the output is a polynomial matrix
	if _, ok := sum.(symbolic.PolynomialMatrix); !ok {
		t.Errorf(
			"expected Plus to return a polynomial matrix; received %T",
			sum,
		)
	}

	// Check that the dimensions of the output are correct
	if sum.Dims()[0] != 1 || sum.Dims()[1] != 4 {
		t.Errorf(
			"expected Plus to return a polynomial matrix of dimensions [1,4]; received %v",
			sum.Dims(),
		)
	}
}

/*
TestPolynomial_Minus1
Description:

	Verifies that the Polynomial.Minus function panics
	when called with a polynomial that is not well-defined.
*/
func TestPolynomial_Minus1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Minus to panic when called with an undefined polynomial; received nil",
			)
		}
	}()

	// Call the Minus method
	p1.Minus(p1)
}

/*
TestPolynomial_Minus2
Description:

	Verifies that the Polynomial.Minus() method panics when called
	with a well-defined polynomial and a not well-defined monomial.
*/
func TestPolynomial_Minus2(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{2},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Minus to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the Minus method
	p1.Minus(m1)
}

/*
TestPolynomial_Minus3
Description:

	Verifies that the Polynomial.Minus() method when applied
	to a well-defined polynomial and a well-defined variable
	returns a polynomial. The new polynomial should have
	one more extra monomial than the original.
*/
func TestPolynomial_Minus3(t *testing.T) {
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
	diff := p1.Minus(symbolic.NewVariable())
	if len(diff.(symbolic.Polynomial).Monomials) != len(p1.Monomials)+1 {
		t.Errorf(
			"expected %v - %v to have 2 monomials; received %v",
			p1,
			v1,
			len(diff.(symbolic.Polynomial).Monomials),
		)
	}
}

/*
TestPolynomial_Minus4
Description:

	Verifies that the Polynomial.Minus() method produces
	the correct output when called with a well-defined polynomial
	and a float64.
*/
func TestPolynomial_Minus4(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	diff := p1.Minus(3.14)
	if len(diff.(symbolic.Polynomial).Monomials) != 2 {
		t.Errorf(
			"expected %v - %v to have 1 monomial; received %v",
			p1,
			3.14,
			len(diff.(symbolic.Polynomial).Monomials),
		)
	}
}

/*
TestPolynomial_Minus5
Description:

	Verifies that the Polynomial.Minus() method panics
	when called with a well-defined polynomial and a non-expression
	(in this case a string).
*/
func TestPolynomial_Minus5(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Minus to panic when called with a non-expression; received nil",
			)
		}
	}()

	// Call the Minus method
	p1.Minus("string")
}

/*
TestPolynomial_ConstantMonomialIndex1
Description:

	Verifies that ConstantMonomialIndex() panics
	when called with a polynomial that is not well-defined.
*/
func TestPolynomial_ConstantMonomialIndex1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected ConstantMonomialIndex to panic when called with an undefined polynomial; received nil",
			)
		}
	}()

	// Call the ConstantMonomialIndex method
	p1.ConstantMonomialIndex()

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
TestPolynomial_VariableMonomialIndex2
Description:

	Verifies that VariableMonomialIndex panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_VariableMonomialIndex2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected VariableMonomialIndex to panic when called with an undefined polynomial; received nil",
			)
		}
	}()

	// Call the VariableMonomialIndex method
	p1.VariableMonomialIndex(symbolic.NewVariable())
}

/*
TestPolynomial_VariableMonomialIndex3
Description:

	Verifies that VariableMonomialIndex panics
	when called with a polynomial that is well-defined
	and a variable that is NOT well-defined.
*/
func TestPolynomial_VariableMonomialIndex3(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected VariableMonomialIndex to panic when called with an undefined variable; received nil",
			)
		}
	}()

	// Call the VariableMonomialIndex method
	p1.VariableMonomialIndex(symbolic.Variable{})
}

/*
TestPolynomial_MonomialIndex1
Description:

	Verifies that MonomialIndex returns the correct index when the
	monomial is in the polynomial.
*/
func TestPolynomial_MonomialIndex1(t *testing.T) {
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

	// Test that the index is 0
	if p1.MonomialIndex(m1) != 0 {
		t.Errorf(
			"expected index of %v to be 0; received %v",
			m1,
			p1.MonomialIndex(m1),
		)
	}
}

/*
TestPolynomial_MonomialIndex2
Description:

	Verifies that MonomialIndex panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_MonomialIndex2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected MonomialIndex to panic when called with an undefined polynomial; received nil",
			)
		}
	}()

	// Call the MonomialIndex method
	p1.MonomialIndex(symbolic.NewVariable().ToMonomial())

}

/*
TestPolynomial_MonomialIndex3
Description:

	Verifies that MonomialIndex panics
	when called with a polynomial that is well-defined
	and a monomial that is NOT well-defined.
*/
func TestPolynomial_MonomialIndex3(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected MonomialIndex to panic when called with an undefined monomial; received nil",
			)
		}
	}()

	// Call the MonomialIndex method
	p1.MonomialIndex(symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{2},
	})
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
TestPolynomial_Multiply4
Description:

	Verifies that the Polynomial.Multiply method panics
	when called with a polynomial that is well-defined
	and a non-expression (in this case a string).
*/
func TestPolynomial_Multiply4(t *testing.T) {
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
TestPolynomial_Multiply5
Description:

	Verifies that the Polynomial.Multiply method produced
	the correct output when called with a well-defined polynomial
	and a float64.
	The resulting polynomial should have the same number of polynomials
	as the original polynomial.
*/
func TestPolynomial_Multiply5(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	prod := p1.Multiply(3.14)
	if len(prod.(symbolic.Polynomial).Monomials) != 1 {
		t.Errorf(
			"expected %v * %v to have 1 monomial; received %v",
			p1,
			3.14,
			len(prod.(symbolic.Polynomial).Monomials),
		)
	}
}

/*
TestPolynomial_Multiply6
Description:

	Verifies that the Polynomial.Multiply method returns the correct output
	when called with a well-defined polynomial (with 2 monomials)
	and a well-defined polynomial (with 2 unique monomials).
	In this case, the resulting polynomial should have
	4 = 2 x 2 monomials.
*/
func TestPolynomial_Multiply6(t *testing.T) {
	// Setup
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
		},
	}
	p2 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
		},
	}

	// Test
	prod := p1.Multiply(p2)

	// Verify that the product is a polynomial
	prodAsP, tf := prod.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v * %v to return a polynomial; received %T",
			p1,
			p2,
			prod,
		)
	}

	// Verify that the product has the correct number of monomials
	if len(prodAsP.Monomials) != 4 {
		t.Errorf(
			"expected (%v) * (%v) to have 4 monomials; received %v (%v)",
			p1,
			p2,
			len(prodAsP.Monomials),
			prodAsP,
		)
	}
}

/*
TestPolynomial_Multiply7
Description:

	Verifies that the Polynomial.Multiply properly panics when the input polynomial is not well-defined.
*/
func TestPolynomial_Multiply7(t *testing.T) {
	// Setup
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Multiply to panic when called with an invalid polynomial; received nil",
			)
		}

		rAsE := r.(error)
		if rAsE.Error() != p1.Check().Error() {
			t.Errorf(
				"expected Multiply to panic with error %v; received %v",
				p1.Check(),
				rAsE,
			)
		}
	}()

	// Call the Multiply method
	p1.Multiply(p1)
}

/*
TestPolynomial_Multiply8
Description:

	Verifies that the Polynomial.Multiply method returns the correct output
	when called with a well-defined polynomial and a well-defined constant matrix.
	The resulting polynomial should be a polynomial matrix.
*/
func TestPolynomial_Multiply8(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()
	km1 := getKMatrix.From([][]float64{
		{3.14, 2.71},
		{1.0, 0.0},
	})

	// Test
	prod := p1.Multiply(km1)

	// Verify that the product is a polynomial matrix
	prodAsPM, tf := prod.(symbolic.PolynomialMatrix)
	if !tf {
		t.Errorf(
			"expected %v * %v to return a polynomial matrix; received %T",
			p1,
			km1,
			prod,
		)
	}

	// Verify that the coefficients of the product are correct
	for ii, pRow := range prodAsPM {
		for jj, p := range pRow {
			if len(p.Monomials) != 1 {
				t.Errorf(
					"expected %v * %v to have 1 monomial; received %v",
					p1,
					km1,
					len(p.Monomials),
				)
			}

			if prodAsPM[ii][jj].Monomials[0].Coefficient != float64(km1.At(ii, jj).(symbolic.K)) {
				t.Errorf(
					"expected %v * %v to have coefficient %v; received %v",
					p1,
					km1,
					km1.At(ii, jj),
					prodAsPM[ii][jj].Monomials[0].Coefficient,
				)
			}

		}
	}
}

/*
TestPolynomial_Transpose1
Description:

	Verifies that the output of the transpose of a polynomial
	is the same as the original polynomial when the polynomial
	is well-defined.
*/
func TestPolynomial_Transpose1(t *testing.T) {
	// Setup
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	pT := p1.Transpose()

	// Verify that the transpose is the same as the original
	if !reflect.DeepEqual(p1, pT) {
		t.Errorf(
			"expected %v^T to be %v; received %v",
			p1,
			p1,
			pT,
		)

	}
}

/*
TestPolynomial_Transpose2
Description:

	Verifies that the Transpose() method properly panics when the polynomial is not well-defined.
*/
func TestPolynomial_Transpose2(t *testing.T) {
	// Setup
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Transpose to panic when called with an invalid polynomial; received nil",
			)
		}

		rAsE := r.(error)
		if rAsE.Error() != p1.Check().Error() {
			t.Errorf(
				"expected Transpose to panic with error %v; received %v",
				p1.Check(),
				rAsE,
			)
		}
	}()

	// Call the Transpose method
	p1.Transpose()
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
TestPolynomial_Constant2
Description:

	Verifies that the Polynomial.Constant method panics when called on a polynomial
	that is not well-defined.
*/
func TestPolynomial_Constant2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Constant to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the Constant method
	p1.Constant()
}

/*
TestPolynomial_Simplify1
Description:

	Verifies that the Polynomial.Simplify method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_Simplify1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Simplify to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the Simplify method
	p1.Simplify()
}

/*
TestPolynomial_Simplify2
Description:

	Verifies that the Polynomial.Simplify method returns the same polynomial,
	when the polynomial contains one monomial with a coefficient of 0.
*/
func TestPolynomial_Simplify2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 0, VariableFactors: []symbolic.Variable{}, Exponents: []int{}},
		},
	}

	// Test
	simp := p1.Simplify()
	if !reflect.DeepEqual(p1, simp) {
		t.Errorf(
			"expected %v to simplify to %v; received %v",
			p1,
			p1,
			simp,
		)
	}
}

/*
TestPolynomial_DerivativeWrt1
Description:

	Verifies that the Polynomial.DerivativeWrt method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_DerivativeWrt1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected DerivativeWrt to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the DerivativeWrt method
	p1.DerivativeWrt(symbolic.NewVariable())
}

/*
TestPolynomial_DerivativeWrt2
Description:

	Verifies that the Polynomial.DerivativeWrt method panics when the polynomial
	is well-defined, but the variable wrt is not well-defined.
*/
func TestPolynomial_DerivativeWrt2(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	v1 := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected DerivativeWrt to panic when called with an invalid variable; received nil",
			)
		}

		rAsE := r.(error)
		if rAsE.Error() != v1.Check().Error() {

		}
	}()

	// Call the DerivativeWrt method
	p1.DerivativeWrt(v1)
}

/*
TestPolynomial_DerivativeWrt3
Description:

	Verifies that the Polynomial.DerivativeWrt method returns the correct output
	of 0 when the well-defined polynomial contains variables and the wrt variable
	is not in the polynomial.
*/
func TestPolynomial_DerivativeWrt3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{1}},
			symbolic.Monomial{Coefficient: 2, VariableFactors: []symbolic.Variable{v2}, Exponents: []int{2}},
		},
	}

	// Test
	derivative := p1.DerivativeWrt(symbolic.NewVariable())
	expected := symbolic.K(0.0)
	if !reflect.DeepEqual(expected, derivative) {
		t.Errorf(
			"expected %v.derivative(%v) to be %v; received %v",
			p1,
			symbolic.NewVariable(),
			expected,
			derivative,
		)
	}
}

/*
TestPolynomial_Degree1
Description:

	Verifies that the Polynomial.Degree method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_Degree1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Degree to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the Degree method
	p1.Degree()
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

/*
TestPolynomial_IsConstant1
Description:

	Verifies that the Polynomial.IsConstant method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_IsConstant1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected IsConstant to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the IsConstant method
	p1.IsConstant()
}

/*
TestPolynomial_IsConstant2
Description:

	Verifies that the Polynomial.IsConstant method returns true
	when called with a polynomial that is a constant (i.e., only contains one monomial that has just a coefficient).
*/
func TestPolynomial_IsConstant2(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 3.14, VariableFactors: []symbolic.Variable{}, Exponents: []int{}},
		},
	}

	// Test
	if !p1.IsConstant() {
		t.Errorf(
			"expected %v to be constant; received %v",
			p1,
			p1.IsConstant(),
		)
	}
}

/*
TestPolynomial_String1
Description:

	Verifies that the Polynomial.String method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_String1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected String to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the String method
	p1.String()
}

/*
TestPolynomial_Substitute1
Description:

	Verifies that the Polynomial.Substitute method panics when called with a polynomial
	that is not well-defined.
*/
func TestPolynomial_Substitute1(t *testing.T) {
	// Constants
	p1 := symbolic.Polynomial{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Substitute to panic when called with an invalid polynomial; received nil",
			)
		}
	}()

	// Call the Substitute method
	p1.Substitute(symbolic.NewVariable(), symbolic.NewVariable())
}

/*
TestPolynomial_Substitute2
Description:

	Verifies that the Polynomial.Substitute method returns the correct output
	when called with a well-defined polynomial, a well-defined variable and a well-defined expression to use for substitution.
*/
func TestPolynomial_Substitute2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{Coefficient: 1, VariableFactors: []symbolic.Variable{v1}, Exponents: []int{1}},
		},
	}

	// Test
	sub := p1.Substitute(v1, v2.Multiply(3.0).(symbolic.ScalarExpression))
	if sub.(symbolic.Polynomial).Monomials[0].Coefficient != 3.0 {
		t.Errorf(
			"expected %v.substitute(%v, %v) to have coefficient 3.0; received %v",
			p1,
			v1,
			v2.Multiply(3.0),
			sub.(symbolic.Polynomial).Monomials[0].Coefficient,
		)
	}
}

/*
TestPolynomial_Substitute3
Description:

	Verifies that the Polynomial.Substitute method panics when the variable used for substitution
	is not well-defiend, but the polynomial is well-defined.
*/
func TestPolynomial_Substitute3(t *testing.T) {
	// Constants
	v1 := symbolic.Variable{}
	p1 := symbolic.NewVariable().ToPolynomial()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Substitute to panic when called with an invalid variable; received nil",
			)
		}

		rAsE := r.(error)
		if rAsE.Error() != v1.Check().Error() {
			t.Errorf(
				"expected Substitute to panic with error %v; received %v",
				v1.Check(),
				rAsE,
			)
		}
	}()

	// Call the Substitute method
	p1.Substitute(v1, symbolic.NewVariable())
}

/*
TestPolynomial_SubstituteWith4
Description:

	Verifies that the Polynomial.Substitute method correctly computes the substitution
	when the polynomial is well-defined and the expression used for substitution is well-defined.
	We make the polynomial very long and complex to replicate a bug that occurred in one
	of the downstream projects.

	p1 = 1 + x1 + x2 + x3 + ... + x20
	substitute x1 with 2.
*/
func TestPolynomial_SubstituteWith4(t *testing.T) {
	// Constants
	N := 20
	x := symbolic.NewVariableVector(N)

	// Create a polynomial that is the sum of 1 and all variables in x
	sum1 := symbolic.K(1).Plus(
		x.Transpose().Multiply(symbolic.OnesVector(N)),
	)
	p1, tf := sum1.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v to be a polynomial; received %T",
			sum1,
			sum1,
		)
	}

	// Test
	substitution := p1.Substitute(x[0], symbolic.K(2.0))

	// Search for a constant element in the monomials
	for _, m := range substitution.(symbolic.Polynomial).Monomials {
		if m.IsConstant() {
			if m.Coefficient != 3.0 {
				t.Errorf(
					"expected (%v).substitute(%v, %v) to have constant 3.0; received %v",
					p1,
					x[0],
					symbolic.K(2.0),
					m.Coefficient,
				)
			}
		}
	}
}
