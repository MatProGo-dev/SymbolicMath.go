package symbolic

/*
monomial_test.go
Description:
	Tests the monomial object.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
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

/*
TestMonomial_Plus5
Description:

	Verifies that the Monomial.Plus panics when
	the monomial used to call it is not well-defined.
*/
func TestMonomial_Plus5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Plus to panic; received nil")
		}
	}()
	m1.Plus(m1)
}

/*
TestMonomial_Plus6
Description:

	Verifies that the Monomial.Plus panics when
	the monomial used to call it is well-defined, but it is
	added to an expression that is well-defined (a variable).
*/
func TestMonomial_Plus6(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Plus to panic; received nil")
		}
	}()
	m1.Plus(symbolic.Variable{})
}

/*
TestMonomial_Plus7
Description:

	This test verifies that the Monomial.Plus panics when
	the monomial used to call it is well-defined, but it is
	added to something that is not an expression (in this case,
	a string).
*/
func TestMonomial_Plus7(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Plus to panic; received nil")
		}
	}()
	m1.Plus("x")
}

/*
TestMonomial_Plus8
Description:

	Verifies that the Monomial.Plus for a valid monomial and a valid
	polynomial returns a valid polynomial.
	For a unique monomial and polynomial, the sum should contain
	n+1 monomials, where n is the number of monomials in the
	original xwpolynomial.
*/
func TestMonomial_Plus8(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1},
	}
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			m2,
			symbolic.K(3.14).ToMonomial(),
		},
	}

	// Compute Sum
	sum := m1.Plus(p1)

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

	// Check that the polynomial has 3 monomials
	if len(sumAsP.Monomials) != len(p1.Monomials)+1 {
		t.Errorf(
			"expected sum to have 3 monomials; received %v",
			len(sumAsP.Monomials),
		)
	}
}

/*
TestMonomial_Minus1
Description:

	Verifies that the Minus() method panics when called using a monomial that is
	not well-defined.
*/
func TestMonomial_Minus1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected Minus to panic; received nil")
		}

		// Check that error is as expected:
		rAsE := r.(error)
		if !strings.Contains(
			rAsE.Error(),
			m1.Check().Error(),
		) {
			t.Errorf(
				"expected error message to contain %v; received %v",
				m1.Check(),
				rAsE.Error(),
			)
		}
	}()
	m1.Minus(m1)
}

/*
TestMonomial_Minus2
Description:

	Verifies that the Monomial.Minus function panics when the monomial used to call it
	is well-defined but the second input expression is not well-defined (a variable).
*/
func TestMonomial_Minus2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Minus to panic; received nil")
		}
	}()
	m1.Minus(symbolic.Variable{})
}

/*
TestMonomial_Minus3
Description:

	Verifies that the Monomial.Minus function returns a valid polynomial when the input to it
	is a floating point number (when the original monomial is NOT a constant).
*/
func TestMonomial_Minus3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{2},
	}
	f1 := 3.14

	// Compute Difference
	difference := m1.Minus(f1)

	// Verify that the difference is a polynomial
	diffAsP, tf := difference.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected difference to be a K; received %T",
			difference,
		)
	}

	// Verify that the polynomial contains 2 monomials
	if len(diffAsP.Monomials) != 2 {
		t.Errorf(
			"expected difference to have 2 monomials; received %v",
			len(diffAsP.Monomials),
		)
	}
}

/*
TestMonomial_Minus4
Description:

	Verifies that the Monomial.Minus function panics when the monomial used to call it
	is well-defined but the second input is not an expression at all (in this case, it's a string).
*/
func TestMonomial_Minus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	s2 := "x"

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected Minus to panic; received nil")
		}

		// Collect error and compare it to expectation
		rAse := r.(error)
		expectedError := smErrors.UnsupportedInputError{
			Input:        s2,
			FunctionName: "Monomial.Minus",
		}

		if !strings.Contains(rAse.Error(), expectedError.Error()) {
			t.Errorf(
				"expected error message to contain %v; received %v",
				expectedError,
				rAse.Error(),
			)
		}
	}()
	_ = m1.Minus(s2)
}

/*
TestMonomial_Multiply1
Description:

	Verifies that the Monomial.Multiply function panics
	when the monomial used to call it is not well-defined.
*/
func TestMonomial_Multiply1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Multiply to panic; received nil")
		}
	}()
	m1.Multiply(m1)
}

/*
TestMonomial_Multiply2
Description:

	Verifies that the Monomial.Multiply function panics
	when the monomial used to call it is well-defined, but it is
	multiplied by an expression that is NOT well-defined (a variable).
*/
func TestMonomial_Multiply2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Multiply to panic; received nil")
		}
	}()
	m1.Multiply(symbolic.Variable{})
}

/*
TestMonomial_Multiply3
Description:

	Verifies that the Monomial.Multiply function panics
	when the monomial used to call it is well-defined, but it is
	multiplied by something that is not an expression (in this case,
	a string).
*/
func TestMonomial_Multiply3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected Multiply to panic; received nil")
		}
	}()
	m1.Multiply("x")
}

/*
TestMonomial_Transpose1
Description:

	Verifies that the Monomial.Transpose function panics
	when the monomial used to call it is not well-defined.
*/
func TestMonomial_Transpose1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected Transpose to panic; received nil",
			)
		}
	}()

	m1.Transpose()

}

/*
TestMonomial_Transpose2
Description:

	Verifies that the Monomial.Transpose function returns
	a monomial expression. The same one that was used to call it.
*/
func TestMonomial_Transpose2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute Transpose
	transposed := m1.Transpose()

	// Verify that the transposed is a monomial
	transposedAsM, tf := transposed.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected transposed to be a monomial; received %T",
			transposed,
		)
	}

	// Verify that the coefficient is the same
	if transposedAsM.Coefficient != m1.Coefficient {
		t.Errorf(
			"expected transposed coefficient to be %v; received %v",
			m1.Coefficient,
			transposedAsM.Coefficient,
		)
	}
}

/*
TestMonomial_LessEq1
Description:

	Verifies that the Monomial.LessEq function panics
	when the monomial used to call it is not well-defined.
*/
func TestMonomial_LessEq1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected LessEq to panic; received nil",
			)
		}
	}()

	m1.LessEq(m1)
}

/*
TestMonomial_LessEq2
Description:

	Verifies that the Monomial.LessEq function produces a valid
	scalar constraint with the correct sense (SenseLessThanEqual).
*/
func TestMonomial_LessEq2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute LessEq
	lessEq := m1.LessEq(3.14)

	// Verify that the lessEq is a scalar constraint
	lessEqAsS, tf := lessEq.(symbolic.ScalarConstraint)
	if !tf {
		t.Errorf(
			"expected lessEq to be a scalar constraint; received %T",
			lessEq,
		)
	}

	// Verify that the sense is correct
	if lessEqAsS.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf(
			"expected lessEq to have sense <=; received %v",
			lessEqAsS.ConstrSense(),
		)
	}
}

/*
TestMonomial_GreaterEq1
Description:

	Verifies that the Monomial.GreaterEq function panics
	when the monomial used to call it is well-defined,
	but the input expression to it is not well-defined.
*/
func TestMonomial_GreaterEq1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected GreaterEq to panic; received nil",
			)
		}
	}()

	m1.GreaterEq(symbolic.Variable{})
}

/*
TestMonomial_GreaterEq2
Description:

	Verifies that the Monomial.GreaterEq function produces a valid
	scalar constraint with the correct sense (SenseGreaterThanEqual)
	when called using a valid monomial and a second input expression
	that is also valid.
*/
func TestMonomial_GreaterEq2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute GreaterEq
	greaterEq := m1.GreaterEq(v1)

	// Verify that the greaterEq is a scalar constraint
	greaterEqAsS, tf := greaterEq.(symbolic.ScalarConstraint)
	if !tf {
		t.Errorf(
			"expected greaterEq to be a scalar constraint; received %T",
			greaterEq,
		)
	}

	// Verify that the sense is correct
	if greaterEqAsS.ConstrSense() != symbolic.SenseGreaterThanEqual {
		t.Errorf(
			"expected greaterEq to have sense >=; received %v",
			greaterEqAsS.ConstrSense(),
		)
	}
}

/*
TestMonomial_Eq1
Description:

	Verifies that the Monomial.Eq function produces a valid
	scalar constraint with the correct sense (SenseEqual)
	when called using a valid monomial and a second input expression
	that is also valid (a polynomial).
*/
func TestMonomial_Eq1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute Eq
	eq := m1.Eq(v1.ToPolynomial())

	// Verify that the eq is a scalar constraint
	eqAsS, tf := eq.(symbolic.ScalarConstraint)
	if !tf {
		t.Errorf(
			"expected eq to be a scalar constraint; received %T",
			eq,
		)
	}

	// Verify that the sense is correct
	if eqAsS.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"expected eq to have sense ==; received %v",
			eqAsS.ConstrSense(),
		)
	}
}

/*
TestMonomial_Eq2
Description:

	Verifies that the Monomial.Eq function returns a valid
	scalar constraint with the correct sense (SenseEqual)
	when called using a valid monomial and a second input expression
	that is also valid (a monomial).
*/
func TestMonomial_Eq2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute Eq
	eq := m1.Eq(v1.ToMonomial())

	// Verify that the eq is a scalar constraint
	eqAsS, tf := eq.(symbolic.ScalarConstraint)
	if !tf {
		t.Errorf(
			"expected eq to be a scalar constraint; received %T",
			eq,
		)
	}

	// Verify that the sense is correct
	if eqAsS.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"expected eq to have sense ==; received %v",
			eqAsS.ConstrSense(),
		)
	}
}

/*
TestMonomial_Comparison1
Description:

	Verifies that the Monomial.Comparison function panics
	when the comparison is called using a valid monomial and a second input
	that is not an expression (but a string).
*/
func TestMonomial_Comparison1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected Comparison to panic; received nil",
			)
		}
	}()

	m1.Comparison("x", symbolic.SenseEqual)
}

/*
TestMonomial_Constant1
Description:

	Verifies that the Monomial.Constant function returns
	zero when the monomial is not a constant.
*/
func TestMonomial_Constant1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	if m1.Constant() != 0 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.Constant(),
		)
	}
}

/*
TestMonomial_Constant2
Description:

	Verifies that the Monomial.Constant function returns
	the coefficient when the monomial is a constant.
*/
func TestMonomial_Constant2(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}

	// Test
	if m1.Constant() != 3.14 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.Constant(),
		)
	}
}

/*
TestMonomial_LinearCoeff1
Description:

	Verifies that the Monomial.LinearCoeff function returns
	zero when the monomial is not a linear monomial.
*/
func TestMonomial_LinearCoeff1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 2},
	}

	// Test
	coeffs := m1.LinearCoeff()

	// Check that the length of the coeffs vector is 2
	if coeffs.Len() != 2 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.LinearCoeff(),
		)
	}

	// Check that all elements of the coeffs vector are 0
	for _, val := range coeffs.RawVector().Data {
		if val != 0 {
			t.Errorf(
				"expected m1 to be a constant; received %v",
				m1.LinearCoeff(),
			)
		}
	}

}

/*
TestMonomial_LinearCoeff2
Description:

	Verifies that the Monomial.LinearCoeff function returns
	the correct linear coefficients when the monomial is a linear monomial.
	When no variable slice is given, this should be a length 1 vector.
*/
func TestMonomial_LinearCoeff2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	coeffs := m1.LinearCoeff()

	// Check that the length of the coeffs vector is 1
	if coeffs.Len() != 1 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.LinearCoeff(),
		)
	}

	// Check that all elements of the coeffs vector are 0
	if coeffs.AtVec(0) != 3.14 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.LinearCoeff(),
		)
	}
}

/*
TestMonomial_LinearCoeff3
Description:

	Verifies that the Monomial.LinearCoeff function returns
	the correct linear coefficients when the monomial is a linear monomial.
	When a variable slice is given of length 5,
	this should be a length 5 vector.
*/
func TestMonomial_LinearCoeff3(t *testing.T) {
	// Constants
	N := 5
	vv1 := symbolic.NewVariableVector(N)
	k := 3
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{vv1[k]},
		Exponents:       []int{1},
	}

	// Test
	coeffs := m1.LinearCoeff(vv1)

	// Check that the length of the coeffs vector is 5
	if coeffs.Len() != 5 {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.LinearCoeff(),
		)
	}

	// Check that all elements of the coeffs vector are 0
	// except for the k-th element, which should be 3.14
	for i := 0; i < N; i++ {
		if i == k {
			if coeffs.AtVec(i) != 3.14 {
				t.Errorf(
					"expected m1 to be a constant; received %v",
					m1.LinearCoeff(),
				)
			}
		} else {
			if coeffs.AtVec(i) != 0 {
				t.Errorf(
					"expected m1 to be a constant; received %v",
					m1.LinearCoeff(),
				)
			}
		}
	}
}

/*
TestMonomial_IsConstant1
Description:

	Verifies that the Monomial.IsConstant function returns
	false when the monomial is not a constant.
*/
func TestMonomial_IsConstant1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Test
	if m1.IsConstant() {
		t.Errorf(
			"expected m1 to be a constant; received %v",
			m1.IsConstant(),
		)
	}
}

/*
TestMonomial_IsConstant2
Description:

	Verifies that the Monomial.IsConstant function
	panics if the monomial is not well-defined.
*/
func TestMonomial_IsConstant2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected IsConstant to panic; received nil",
			)
		}
	}()

	m1.IsConstant()
}

/*
TestMonomial_IsVariable1
Description:

	Verifies that the Monomial.IsVariable function returns
	false when the monomial is not the same as a separate variable
	v2.
*/
func TestMonomial_IsVariable1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1},
	}

	// Test
	if m1.IsVariable(v1) {
		t.Errorf(
			"expected m1 to be a variable; received %v",
			m1.IsVariable(v1),
		)
	}
}

/*
TestMonomial_IsVariable2
Description:

	Verifies that the Monomial.IsVariable function
	panics if the monomial is not well-defined.
*/
func TestMonomial_IsVariable2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected IsVariable to panic; received nil",
			)
		}
	}()

	m1.IsVariable(v1)
}

/*
TestMonomial_IsVariable3
Description:

	Verifies that the Monomial.IsVariable function
	panics if the input variable is not well-defined.
*/
func TestMonomial_IsVariable3(t *testing.T) {
	// Constants
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected IsVariable to panic; received nil",
			)
		}
	}()

	m1.IsVariable(symbolic.Variable{})
}

/*
TestMonomial_DerivativeWrt1
Description:

	Verifies that the Monomial.DerivativeWrt function
	panics if the monomial is not well-defined.
*/
func TestMonomial_DerivativeWrt1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected DerivativeWrt to panic; received nil",
			)
		}
	}()

	m1.DerivativeWrt(v1)
}

/*
TestMonomial_DerivativeWrt2
Description:

	Verifies that the Monomial.DerivativeWrt function
	returns a valid monomial when the input variable is contained
	in the monomial.
*/
func TestMonomial_DerivativeWrt2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}

	// Compute DerivativeWrt
	derivative := m1.DerivativeWrt(v1)

	// Verify that the derivative is a monomial
	derivativeAsM, tf := derivative.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected derivative to be a monomial; received %T",
			derivative,
		)
	}

	// Verify that the derivative is a constant
	if derivativeAsM.Coefficient != 3.14 {
		t.Errorf(
			"expected derivative to be a constant; received %v",
			derivativeAsM.Coefficient,
		)
	}
}

/*
TestMonomial_DerivativeWrt3
Description:

	Verifies that the Monomial.DerivativeWrt function
	returns 0 when the monomial does not contain the
	input variable.
*/
func TestMonomial_DerivativeWrt3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1},
	}

	// Compute DerivativeWrt
	derivative := m1.DerivativeWrt(v1)

	// Verify that the derivative is a constant
	if derivative.(symbolic.K) != symbolic.K(0) {
		t.Errorf(
			"expected derivative to be a constant; received %v",
			derivative,
		)
	}
}

/*
TestMonomial_String1
Description:

	Verifies that the Monomial.String function returns
	a string representation of the monomial.
	Verify that the variable factors in the monomial string contains the strings
	of the variables.
*/
func TestMonomial_String1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 2},
	}

	// Test
	if !strings.Contains(m1.String(), v1.String()) {
		t.Errorf(
			"expected string to contain %v; received %v",
			v1.String(),
			m1.String(),
		)
	}

	if !strings.Contains(m1.String(), v2.String()) {
		t.Errorf(
			"expected string to contain %v; received %v",
			v2.String(),
			m1.String(),
		)
	}
}

/*
TestMonomial_String2
Description:

	Verifies that the Monomial.String function panics
	when the monomial is not well-defined.
*/
func TestMonomial_String2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"expected String to panic; received nil",
			)
		}
	}()

	_ = m1.String()
}
