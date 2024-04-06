package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
variable_test.go
Description:
	Testing functions relevant to the Var() object. (Scalar Variable)
*/

/*
TestVariable_Constant1
Description:

	Tests whether or not NumVars returns 0 as the constant included in the a single variable.
*/
func TestVariable_Constant1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if x.Constant() != 0.0 {
		t.Errorf(
			"The constant component of a variable (%T) should be 0.0; received %v",
			x,
			x.Constant(),
		)
	}

}

/*
TestVariable_LinearCoeff1
Description:

	Tests whether the LinearCoeff() method returns a vector of 1.0 when called on a variable
	with no inputs.
*/
func TestVariable_LinearCoeff1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	coeff := x.LinearCoeff()
	// Check that the length of the vector is 1
	if coeff.Len() != 1 {
		t.Errorf(
			"expected the linear coefficient of %v to be a vector of length 1; received %v",
			x,
			coeff.Len(),
		)
	}

	// Check that the only element is 1.0
	if coeff.AtVec(0) != 1.0 {
		t.Errorf(
			"expected the linear coefficient of %v to be 1.0; received %v",
			x,
			coeff.AtVec(0),
		)
	}
}

/*
TestVariable_LinearCoeff2
Description:

	Tests whether the LinearCoeff() method panics when called on a variable
	with an input that is an empty slice.
*/
func TestVariable_LinearCoeff2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.LinearCoeff([]symbolic.Variable{})
}

/*
TestVariable_LinearCoeff3
Description:

	Tests whether the LinearCoeff() method returns a vector that
	contains a 1.0 in one of its elements when called on a variable
	with an input that is a slice containing the same variable.
*/
func TestVariable_LinearCoeff3(t *testing.T) {
	// Constants
	N := 5
	vv1 := symbolic.NewVariableVector(N)
	k := 3

	// Test
	coeff := vv1[k].LinearCoeff(vv1)
	// Check that the length of the vector is 1
	if coeff.Len() != N {
		t.Errorf(
			"expected the linear coefficient of %v to be a vector of length 1; received %v",
			vv1[k],
			coeff.Len(),
		)
	}

	// Check all elements of coeff
	// - almost all entries should be 0.0
	// - the k-th entry should be 1.0
	for i := 0; i < N; i++ {
		if i == k {
			if coeff.AtVec(i) != 1.0 {
				t.Errorf(
					"expected the linear coefficient of %v to be 1.0; received %v",
					vv1[k],
					coeff.AtVec(i),
				)
			}
		} else {
			if coeff.AtVec(i) != 0.0 {
				t.Errorf(
					"expected the linear coefficient of %v to be 0.0; received %v",
					vv1[k],
					coeff.AtVec(i),
				)
			}
		}
	}
}

/*
TestVariable_LinearCoeff4
Description:

	Verifies that LinearCoeff panics when called with
	a Variable that is not well-formed.
*/
func TestVariable_LinearCoeff4(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.LinearCoeff()
}

/*
TestVariable_LinearCoeff5
Description:

	Verifies that LinearCoeff panics when called with
	more than one input.
*/
func TestVariable_LinearCoeff5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.LinearCoeff([]symbolic.Variable{x}, []symbolic.Variable{x})
}

/*
TestVariable_Plus1
Description:

	Tests that the Plus() method works properly when adding a float64 to a variable.
*/
func TestVariable_Plus1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(3.14)
	if sum.(symbolic.ScalarExpression).Constant() != 3.14 {
		t.Errorf(
			"expected %v + 3.14 to have constant component 3.14; received %v",
			x,
			x.Plus(3.14),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + 3.14 to be a polynomial; received %T",
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + 3.14 to have 2 terms; received %v",
			x,
			len(sumAsPoly.Monomials),
		)
	}

}

/*
TestVariable_Plus2
Description:

	Tests that the Plus() method works properly when adding a constant to a variable.
*/
func TestVariable_Plus2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(symbolic.K(3.14))
	if sum.(symbolic.ScalarExpression).Constant() != 3.14 {
		t.Errorf(
			"expected %v + 3.14 to have constant component 3.14; received %v",
			x,
			x.Plus(3.14),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + 3.14 to be a polynomial; received %T",
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + 3.14 to have 2 terms; received %v",
			x,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus3
Description:

	Tests that the Plus() method works properly when adding a variable to
	a DIFFERENT variable.
*/
func TestVariable_Plus3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Test
	sum := x.Plus(y)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			y,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			y,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + %v to have 2 terms; received %v",
			x,
			y,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus4
Description:

	Tests that the Plus() method works properly when adding a variable to
	the same variable.
*/
func TestVariable_Plus4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(x)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			x,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 1 {
		t.Errorf(
			"expected %v + %v to have 1 term; received %v",
			x,
			x,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus5
Description:

	This test verifies that a "panic" is thrown when a variable
	is used in the Plus() method that is not well-formed.
*/
func TestVariable_Plus5(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.Plus("hello")
}

/*
TestVariable_Plus6
Description:

	This test verifies that a variable added to a monomial
	produces a polynomial with two monomials.
*/
func TestVariable_Plus6(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{x},
		Exponents:       []int{2},
	}

	// Test
	sum := x.Plus(m)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			m,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			m,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + %v to have 2 terms; received %v",
			x,
			m,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus7
Description:

	This test verifies that a variable added to a polynomial
	produces a polynomial with two monomials.
*/
func TestVariable_Plus7(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	p := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{2},
			},
			symbolic.Monomial{
				Coefficient:     2.71,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{4},
			},
		},
	}

	// Test
	sum := x.Plus(p)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			p,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			p,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != len(p.Monomials)+1 {
		t.Errorf(
			"expected %v + %v to have 2 terms; received %v",
			x,
			p,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus8
Description:

	This test verifies that the variable.Plus method
	properly panics when given an unexpected type (in this case, a string).
*/
func TestVariable_Plus8(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.Plus("hello")
}

/*
TestVariable_Plus9
Description:

	This test verifies that the variable.Plus method
	returns a polynomialvector with the correct length
	when a well-formed variable is added to a mat.VecDense.
*/
func TestVariable_Plus9(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	v := mat.NewVecDense(3, []float64{1, 2, 3})

	// Test
	sum := x.Plus(v)
	sumAsPV, tf := sum.(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a KVector; received %T",
			x,
			v,
			sum,
		)
	}

	// Check that the length of the vector is 3
	if sumAsPV.Len() != 3 {
		t.Errorf(
			"expected %v + %v to be a vector of length 3; received %v",
			x,
			v,
			sumAsPV.Len(),
		)
	}
}

/*
TestVariable_Minus1
Description:

	Verifies that the Minus() method works properly when subtracting a float64 from a variable.
*/
func TestVariable_Minus1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	diff := x.Minus(3.14)
	if diff.(symbolic.ScalarExpression).Constant() != -3.14 {
		t.Errorf(
			"expected %v - 3.14 to have constant component -3.14; received %v",
			x,
			x.Minus(3.14),
		)
	}

	// Test that diff is a polynomial with 2 terms
	diffAsPoly, tf := diff.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v - 3.14 to be a polynomial; received %T",
			x,
			diff,
		)
	}

	if len(diffAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v - 3.14 to have 2 terms; received %v",
			x,
			len(diffAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Minus2
Description:

	Verifies that the Minus() method works properly when subtracting a float64 from a variable.
*/
func TestVariable_Minus2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	diff := x.Minus(3.14)
	if diff.(symbolic.ScalarExpression).Constant() != -3.14 {
		t.Errorf(
			"expected %v - 3.14 to have constant component -3.14; received %v",
			x,
			x.Minus(3.14),
		)
	}

	// Test that diff is a polynomial with 2 terms
	diffAsPoly, tf := diff.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v - 3.14 to be a polynomial; received %T",
			x,
			diff,
		)
	}

	if len(diffAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v - 3.14 to have 2 terms; received %v",
			x,
			len(diffAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Minus3
Description:

	Verifies that the Minus() method properly panics when called on a variable
	that is not well-defined.
*/
func TestVariable_Minus3(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.Minus("hello")
}

/*
TestVariable_LessEq1
Description:

	Tests that the LessEq() method works properly when comparing a variable to a float64.
*/
func TestVariable_LessEq1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Create comparison
	le := x.LessEq(3.14)

	// Verify that the sense of the constraint is SenseLessThanEqual
	if le.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf(
			"expected %v <= 3.14 to have a sense of LessThanEq; received %v",
			x,
			le.ConstrSense(),
		)
	}

	// Get Left side and verify that it is a variable object
	left := le.Left()
	if _, tf := left.(symbolic.Variable); !tf {
		t.Errorf(
			"expected %v <= 3.14 to have a variable component; received %T",
			x,
			left,
		)
	}

	// Get Right side and verify that it is a constant object
	right := le.Right()
	if _, tf := right.(symbolic.K); !tf {
		t.Errorf(
			"expected %v <= 3.14 to have a constant component on RHS; received %T",
			x,
			right,
		)
	}
}

/*
TestVariable_LessEq2
Description:

	Tests that the LessEq() method properly panics
	if the variable used to call it is not well-defined.
*/
func TestVariable_LessEq2(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.LessEq(3.14)
}

/*
TestVariable_GreaterEq1
Description:

	Tests that the GreaterEq() method works properly when comparing a variable to a monomial.
*/
func TestVariable_GreaterEq1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{x},
		Exponents:       []int{2},
	}

	// Create comparison
	ge := x.GreaterEq(m)

	// Verify that the constraints sense is SenseGreaterThanEqual
	if ge.ConstrSense() != symbolic.SenseGreaterThanEqual {
		t.Errorf(
			"expected %v >= %v to have a sense of GreaterThanEq; received %v",
			x,
			m,
			ge.ConstrSense(),
		)
	}

	// Get Left side and verify that it is a variable object
	left := ge.Left()
	if _, tf := left.(symbolic.Variable); !tf {
		t.Errorf(
			"expected %v >= %v to have a variable component; received %T",
			x,
			m,
			left,
		)
	}

	// Get Right side and verify that it is a monomial object
	right := ge.Right()
	if _, tf := right.(symbolic.Monomial); !tf {
		t.Errorf(
			"expected %v >= %v to have a monomial component on RHS; received %T",
			x,
			m,
			right,
		)
	}
}

/*
TestVariable_GreaterEq2
Description:

	Tests that the GreaterEq() method properly panics
	if the second input is not a well-defined expression
	(in this case, it's a bad monomial).
*/
func TestVariable_GreaterEq2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	m := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{x},
		Exponents:       []int{2, 3},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.GreaterEq(m)
}

/*
TestVariable_Eq1
Description:

	Tests that the Eq() method works properly when comparing a variable to a
	well-defined variable.
*/
func TestVariable_Eq1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Create comparison
	eq := x.Eq(y)

	// Verify that the constraints sense is SenseEqual
	if eq.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"expected %v == %v to have a sense of Equal; received %v",
			x,
			y,
			eq.ConstrSense(),
		)
	}

	// Get Left side and verify that it is a variable object
	left := eq.Left()
	if _, tf := left.(symbolic.Variable); !tf {
		t.Errorf(
			"expected %v == %v to have a variable component; received %T",
			x,
			y,
			left,
		)
	}

	// Get Right side and verify that it is a variable object
	right := eq.Right()
	if _, tf := right.(symbolic.Variable); !tf {
		t.Errorf(
			"expected %v == %v to have a variable component on RHS; received %T",
			x,
			y,
			right,
		)
	}
}

/*
TestVariable_Eq2
Description:

	Tests that the Eq() method properly panics
	if the second input is not a valid object
	(in this case, it's a string).
*/
func TestVariable_Eq2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.Eq("hello")
}

/*
TestVariable_Comparison1
Description:

	Tests that the Comparison() method works properly when comparing a variable to a
	well-defined polynomial.
*/
func TestVariable_Comparison1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	p := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{2},
			},
			symbolic.Monomial{
				Coefficient:     2.71,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{4},
			},
		},
	}

	// Create comparison
	comp := x.Comparison(p, symbolic.SenseLessThanEqual)

	// Get Left side and verify that it is a variable object
	left := comp.Left()
	if _, tf := left.(symbolic.Variable); !tf {
		t.Errorf(
			"expected %v >= %v to have a variable component; received %T",
			x,
			p,
			left,
		)
	}

	// Get Right side and verify that it is a polynomial object
	right := comp.Right()
	if _, tf := right.(symbolic.Polynomial); !tf {
		t.Errorf(
			"expected %v >= %v to have a polynomial component on RHS; received %T",
			x,
			p,
			right,
		)
	}
}

/*
TestVariable_Multiply1
Description:

	Tests that the Multiply() method works properly
	panics when the variable used to call it is not
	well-defined.
*/
func TestVariable_Multiply1(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	x.Multiply(3.14)
}

/*
TestVariable_Multiply2
Description:

	Tests that the Multiply() method works properly
	when a variable multiplies THE SAME variable.
	We will check that the output is a monomial
	with one variable factor and an exponent of 2.
*/
func TestVariable_Multiply2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	prod := x.Multiply(x)
	mon, tf := prod.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected %v * %v to be a monomial; received %T",
			x,
			x,
			prod,
		)
	}

	if mon.Exponents[0] != 2 {
		t.Errorf(
			"expected %v * %v to have an exponent of 2; received %v",
			x,
			x,
			mon.Exponents[0],
		)
	}

	if mon.Coefficient != 1.0 {
		t.Errorf(
			"expected %v * %v to have a coefficient of 1.0; received %v",
			x,
			x,
			mon.Coefficient,
		)
	}

}

/*
TestVariable_Multiply3
Description:

	Tests that the Multiply() method works properly
	when a variable multiplies a DIFFERENT variable.
	We will check that the output is a monomial
	with two variable factors and an exponent of 1
	on each.
*/
func TestVariable_Multiply3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Test
	prod := x.Multiply(y)
	mon, tf := prod.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected %v * %v to be a monomial; received %T",
			x,
			y,
			prod,
		)
	}

	if mon.Exponents[0] != 1 {
		t.Errorf(
			"expected %v * %v to have an exponent of 1; received %v",
			x,
			y,
			mon.Exponents[0],
		)
	}

	if mon.Exponents[1] != 1 {
		t.Errorf(
			"expected %v * %v to have an exponent of 1; received %v",
			x,
			y,
			mon.Exponents[1],
		)
	}

	if mon.Coefficient != 1.0 {
		t.Errorf(
			"expected %v * %v to have a coefficient of 1.0; received %v",
			x,
			y,
			mon.Coefficient,
		)
	}
}

/*
TestVariable_Multiply4
Description:

	Tests that the Multiply() method works properly
	when a variable multiplies a *mat.VecDense vector.
*/
func TestVariable_Multiply4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	v := mat.NewVecDense(3, []float64{1, 2, 3})

	// Test
	prod := x.Multiply(v)
	if _, tf := prod.(symbolic.MonomialVector); !tf {
		t.Errorf(
			"expected %v * %v to be a KVector; received %T",
			x,
			v,
			prod,
		)
	}

	// Check that the length of the vector is 3
	if prod.(symbolic.MonomialVector).Len() != 3 {
		t.Errorf(
			"expected the linear coefficient of %v to be a vector of length 3; received %v",
			x,
			prod.(symbolic.MonomialVector).Len(),
		)
	}

	// Check that each element of the vector contains a coefficient
	// that matches v
	for ii := 0; ii < 3; ii++ {
		if prod.(symbolic.MonomialVector)[ii].Coefficient != v.AtVec(ii) {
			t.Errorf(
				"expected the linear coefficient of %v to be %v; received %v",
				x,
				v.AtVec(ii),
				prod.(symbolic.MonomialVector)[ii].Coefficient,
			)
		}
	}
}

/*
TestVariable_Multiply5
Description:

	Tests that the Multiply() method works properly
	when a variable multiplies a polynomial.
*/
func TestVariable_Multiply5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	p := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.Monomial{
				Coefficient:     3.14,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{2},
			},
			symbolic.Monomial{
				Coefficient:     2.71,
				VariableFactors: []symbolic.Variable{x},
				Exponents:       []int{4},
			},
		},
	}

	// Test
	prod := x.Multiply(p)
	if _, tf := prod.(symbolic.Polynomial); !tf {
		t.Errorf(
			"expected %v * %v to be a polynomial; received %T",
			x,
			p,
			prod,
		)
	}

	// Check that the product is a polynomial
	p2, ok := prod.(symbolic.Polynomial)
	if !ok {
		t.Errorf(
			"expected %v * %v to be a polynomial; received %T",
			x,
			p,
			prod,
		)
	}

	// Check that the length of the vector is 3
	if len(p2.Monomials) != 2 {
		t.Errorf(
			"expected the polynomial %v * %v to have 2 terms; received %v",
			x,
			p,
			len(p2.Monomials),
		)
	}

	// Check that each element of the vector contains a coefficient
	// that matches v
	for ii := 0; ii < 2; ii++ {
		if prod.(symbolic.Polynomial).Monomials[ii].Coefficient != p.Monomials[ii].Coefficient {
			t.Errorf(
				"expected the polynomial %v * %v to have a coefficient of %v; received %v",
				x,
				p,
				p.Monomials[ii].Coefficient,
				prod.(symbolic.Polynomial).Monomials[ii].Coefficient,
			)
		}
	}
}

/*
TestVariable_Multiply6
Description:

	Tests that the Multiply() method properly panics when called with a
	well-defined variable and an expression that is not well-defined
	(in this case, a variable).
*/
func TestVariable_Multiply6(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	uninit := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), uninit.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()
	x.Multiply(symbolic.Variable{})
}

/*
TestVariable_String1
Description:

	Tests that the String() method works properly.
*/
func TestVariable_String1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if !strings.Contains(x.String(), "x_") {
		t.Errorf(
			"expected %v to be \"x\"; received %v",
			x,
			x.String(),
		)
	}
}

/*
TestVariable_Transpose1
Description:

	Tests that the Transpose() method returns the same variable that
	it was called on.
*/
func TestVariable_Transpose1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Algorthm
	xT := x.Transpose()

	xTAsV, tf := xT.(symbolic.Variable)
	if !tf {
		t.Errorf(
			"expected %v.T to be a variable; received %T",
			x,
			xT,
		)
	}

	if xTAsV.ID != x.ID {
		t.Errorf(
			"expected %v.T to be %v; received %v",
			x,
			x,
			xT,
		)
	}
}

/*
TestVariable_NewBinaryVariable1
Description:

	Tests that the NewBinaryVariable() method properly creates
	a variable with binary type.
*/
func TestVariable_NewBinaryVariable1(t *testing.T) {
	// Constants
	x := symbolic.NewBinaryVariable()

	// Test
	if x.Type != symbolic.Binary {
		t.Errorf(
			"expected %v to be a binary variable; received %v",
			x,
			x.Type,
		)
	}
}

/*
TestVariable_DerivativeWrt1
Description:

	Tests that the DerivativeWrt() method properly panics when called
	on a variable that is not well-defined.
*/
func TestVariable_DerivativeWrt1(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), x.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()
	x.DerivativeWrt(x)
}

/*
TestVariable_DerivativeWrt2
Description:

	Tests that the DerivativeWrt() method properly returns zero when
	called on a variable with a w.r.t. variable that is different
	from the receiver.
*/
func TestVariable_DerivativeWrt2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Test
	der := x.DerivativeWrt(y)

	derAsScalar, tf := der.(symbolic.ScalarExpression)
	if !tf {
		t.Errorf(
			"expected d(%v)/d(%v) to be a scalar expression; received %T",
			x,
			y,
			der,
		)

	}
	if derAsScalar.Constant() != 0.0 {
		t.Errorf(
			"expected d(%v)/d(%v) to be 0.0; received %v",
			x,
			y,
			derAsScalar.Constant(),
		)
	}
}

/*
TestVariable_DerivativeWrt3
Description:

	Tests that the DerivativeWrt() method properly returns 1.0 when
	called on a variable with a w.r.t. variable that is the same
	as the receiver.
*/
func TestVariable_DerivativeWrt3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	der := x.DerivativeWrt(x)

	derAsScalar, tf := der.(symbolic.ScalarExpression)
	if !tf {
		t.Errorf(
			"expected d(%v)/d(%v) to be a scalar expression; received %T",
			x,
			x,
			der,
		)

	}
	if derAsScalar.Constant() != 1.0 {
		t.Errorf(
			"expected d(%v)/d(%v) to be 1.0; received %v",
			x,
			x,
			derAsScalar.Constant(),
		)
	}
}

/*
TestVariable_DerivativeWrt4
Description:

	Tests that the DerivativeWrt() method properly panics when called
	with a well-defined variable receiver and an input that is not
	well-defined.
*/
func TestVariable_DerivativeWrt4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	uninit := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), uninit.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()
	x.DerivativeWrt(symbolic.Variable{})
}

/*
TestVariable_Substitute1
Description:

	Tests that the Substitute() method properly panics when called
	on a variable that is not well-defined.
*/
func TestVariable_Substitute1(t *testing.T) {
	// Constants
	var x symbolic.Variable

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), x.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()
	x.Substitute(x, x)
}

/*
TestVariable_Substitute2
Description:

	Tests that the Substitute() method properly panics when called with a variable input vIn
	that is not well-defined.
*/
func TestVariable_Substitute2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	var vIn symbolic.Variable

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), vIn.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()

	x.Substitute(vIn, x)
}

/*
TestVariable_Substitute3
Description:

	Tests that the Substitute() method properly panics when called with:
	- A well-defined variable receiver (v)
	- A well-defined variable input vIn
	- An expression that is not well-defined
*/
func TestVariable_Substitute3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()
	var uninit symbolic.Variable

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("The panic was not an error")
		}

		if !strings.Contains(rAsE.Error(), uninit.Check().Error()) {
			t.Errorf("The panic was not due to a bad Variable")
		}
	}()
	x.Substitute(y, uninit)
}

/*
TestVariable_Substitute4
Description:

	Tests that the Substitute() method properly returns the variable
	that it was called on when the input variable is different.
	(And everything is well-defined.)
*/
func TestVariable_Substitute4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	sumAsScalar := y.Plus(3.114).(symbolic.ScalarExpression)

	// Test
	sub := x.Substitute(y, sumAsScalar)

	subAsV, tf := sub.(symbolic.Variable)
	if !tf {
		t.Errorf(
			"expected %v.sub(%v, %v) to be a variable; received %T",
			x,
			y,
			x,
			sub,
		)
	}

	if subAsV.ID != x.ID {
		t.Errorf(
			"expected %v.sub(%v, %v) to be %v; received %v",
			x,
			y,
			x,
			x,
			sub,
		)
	}
}

/*
TestVariable_Substitute5
Description:

	Tests that the Substitute() method properly returns the expression
	that it was called on when the input variable is the same as the replacement variable
	vIn (and everything is well-defined).
*/
func TestVariable_Substitute5(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient: 3.14,
	}

	// Test
	sub := x.Substitute(x, m1)

	_, tf := sub.(symbolic.Monomial)
	if !tf {
		t.Errorf(
			"expected %v.sub(%v, %v) to be a variable; received %T",
			x,
			x,
			x,
			sub,
		)
	}
}
