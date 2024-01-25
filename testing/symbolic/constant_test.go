package symbolic

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
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

/*
TestConstant_Plus3
Description:

	This test verifies that the Plus() method panics when it is given an improperly defined
	expression. In this case, the expression is a symbolic.Variable.
*/
func TestConstant_Plus3(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v1 := symbolic.Variable{
		Lower: 1,
		Upper: 0,
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"expected Plus() to panic when given a bad symbolic.Variable; received nothing",
			)
		}

		rAsError := r.(error)
		expectedError := fmt.Errorf("lower bound (%v) of variable must be less than upper bound (%v)", v1.Lower, v1.Upper)
		if !strings.Contains(
			rAsError.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"expected Plus() to panic with error \"%v\"; received %v",
				expectedError,
				rAsError,
			)
		}

	}()

	k1.Plus(v1)
}

/*
TestConstant_Plus4
Description:

	This test verifies that the K.Plus() method properly returns a Polynomial when given a
	well-defined Variable as input.
*/
func TestConstant_Plus4(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v1 := symbolic.NewVariable()

	// Test
	sum := k1.Plus(v1)

	sumAsPolynomial, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected Plus() to return a Polynomial; received %T",
			sum,
		)
	}

	if len(sumAsPolynomial.Monomials) != 2 {
		t.Errorf(
			"expected Plus() to return a Polynomial with 2 monomials; received %v",
			len(sumAsPolynomial.Monomials),
		)
	}

	// Check that one of the monomials is the constant 3.14 (i.e., has coefficient 3.14 and has no variables)
	var constantIndex int = sumAsPolynomial.ConstantMonomialIndex()
	if constantIndex == -1 {
		t.Errorf(
			"expected Plus() to return a Polynomial with a constant monomial; found no constant terms",
		)
	}

	if sumAsPolynomial.Monomials[constantIndex].Coefficient != 3.14 {
		t.Errorf(
			"expected Plus() to return a Polynomial with a monomial with coefficient 3.14; received %v",
			sumAsPolynomial.Monomials[0],
		)
	}

	// Check that the other monomial is the variable
	var variableIndex int = 1 - constantIndex
	if len(sumAsPolynomial.Monomials[variableIndex].VariableFactors) != 1 {
		t.Errorf(
			"expected Plus() to return a Polynomial with a monomial with 1 variable; received %v",
			sumAsPolynomial.Monomials[variableIndex],
		)
	}
}

/*
TestConstant_Plus5
Description:

	This test verifies that a constant plus a monomial returns a monomial ONLY if the monomial
	represents a constant.
*/
func TestConstant_Plus5(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{},
		Degrees:         []int{},
	}

	// Test
	sum := k1.Plus(m1)

	sumAsK, tf := sum.(symbolic.K)
	if !tf {
		t.Errorf(
			"expected Plus() to return a K; received %T",
			sum,
		)
	}

	// Check that the proper constant is revealed.
	if (float64(sumAsK) < 4.14-0.01) || (float64(sumAsK) > 4.14+0.01) {
		t.Errorf(
			"expected Plus() to return a K 4.14; received %v",
			sumAsK,
		)
	}
}

/*
TestConstant_Plus6
Description:

	This test verifies that when a constant is added to a polynomial, the result is a polynomial.
*/
func TestConstant_Plus6(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{
			symbolic.NewVariable().ToMonomial(),
			{
				Coefficient:     1.0,
				VariableFactors: []symbolic.Variable{},
				Degrees:         []int{},
			},
		},
	}

	// Test
	sum := k1.Plus(p1)

	sumAsP, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected Plus() to return a Polynomial; received %T",
			sum,
		)
	}

	if len(sumAsP.Monomials) != 2 {
		t.Errorf(
			"expected Plus() to return a Polynomial with 2 monomials; received %v",
			len(sumAsP.Monomials),
		)
	}

	// Check that one of the monomials is the constant 4.14 (i.e., has coefficient 4.14 and has no variables)
	var constantIndex int = sumAsP.ConstantMonomialIndex()
	if constantIndex == -1 {
		t.Errorf(
			"expected Plus() to return a Polynomial with a constant monomial; found no constant terms",
		)
	}

	if (sumAsP.Monomials[constantIndex].Coefficient > 4.14+1e-5) || (sumAsP.Monomials[constantIndex].Coefficient < 4.14-1e-5) {
		t.Errorf(
			"expected Plus() to return a Polynomial with a monomial with coefficient 4.14; received %v",
			sumAsP.Monomials[constantIndex],
		)
	}

	// Check that the other monomial is the variable
	var variableIndex int = 1 - constantIndex
	if len(sumAsP.Monomials[variableIndex].VariableFactors) != 1 {
		t.Errorf(
			"expected Plus() to return a Polynomial with a monomial with 1 variable; received %v",
			sumAsP.Monomials[variableIndex],
		)
	}
}

/*
TestConstant_Plus7
Description:

	Tests that when a constant is added to a mat.VecDense object,
	the result is a constant vector.
*/
func TestConstant_Plus7(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v1 := symbolic.OnesVector(4)

	// Test
	sum := k1.Plus(v1)

	sumAsKV, tf := sum.(symbolic.KVector)
	if !tf {
		t.Errorf(
			"expected Plus() to return a KVectir; received %T",
			sum,
		)
	}

	// Check that the proper constant is revealed.
	for ii := 0; ii < v1.Len(); ii++ {

		if (float64(sumAsKV.AtVec(ii).(symbolic.K)) < 4.14-0.01) || (float64(sumAsKV.AtVec(ii).(symbolic.K)) > 4.14+0.01) {
			t.Errorf(
				"expected Plus() to return a K 4.14; received %v",
				sumAsKV.AtVec(ii),
			)
		}
	}
}

/*
TestConstant_Plus8
Description:

	Tests that when a constant is added to a *mat.VecDense object,
	the result is a constant vector.
*/
func TestConstant_Plus8(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v1 := symbolic.OnesVector(4)

	// Test
	sum := k1.Plus(&v1)

	sumAsKV, tf := sum.(symbolic.KVector)
	if !tf {
		t.Errorf(
			"expected Plus() to return a KVectir; received %T",
			sum,
		)
	}

	// Check that the proper constant is revealed.
	for ii := 0; ii < v1.Len(); ii++ {

		if (float64(sumAsKV.AtVec(ii).(symbolic.K)) < 4.14-0.01) || (float64(sumAsKV.AtVec(ii).(symbolic.K)) > 4.14+0.01) {
			t.Errorf(
				"expected Plus() to return a K 4.14; received %v",
				sumAsKV.AtVec(ii),
			)
		}
	}
}

/*
TestConstant_Plus9
Description:

	Tests that when a constant is added to a KVector object,
	the result is a constant vector.
*/
func TestConstant_Plus9(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v2 := symbolic.OnesVector(4)
	kv2 := symbolic.VecDenseToKVector(v2)

	// Test
	sum := k1.Plus(kv2)

	sumAsKV, tf := sum.(symbolic.KVector)
	if !tf {
		t.Errorf(
			"expected Plus() to return a KVectir; received %T",
			sum,
		)
	}

	// Check that the proper constant is revealed.
	for ii := 0; ii < kv2.Len(); ii++ {

		if (float64(sumAsKV.AtVec(ii).(symbolic.K)) < 4.14-0.01) || (float64(sumAsKV.AtVec(ii).(symbolic.K)) > 4.14+0.01) {
			t.Errorf(
				"expected Plus() to return a K 4.14; received %v",
				sumAsKV.AtVec(ii),
			)
		}
	}
}

/*
TestConstant_Plus10
Description:

	Tests that when a constant is added to a PolynomialVector object,
	a PolynomialVector object is returned.
*/
func TestConstant_Plus10(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	N := 5
	vv2 := symbolic.NewVariableVector(N)
	var pv3 symbolic.PolynomialVector
	for ii := 0; ii < N; ii++ {
		pv3 = append(pv3, vv2[ii].ToPolynomial())
	}

	// Test
	sum := k1.Plus(pv3)

	sumAsPV, tf := sum.(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"expected Plus() to return a PolynomialVector; received %T",
			sum,
		)
	}

	// Check that the sum's elements are all correct
	for ii := 0; ii < pv3.Len(); ii++ {
		// Collect the polynomial
		p_ii := sumAsPV[ii]

		// Check that the polynomial has the correct number of monomials (2)
		if len(p_ii.Monomials) != 2 {
			t.Errorf(
				"expected Plus() to return a PolynomialVector with Polynomial elements with 2 monomials; received %v",
				len(p_ii.Monomials),
			)
		}

		// Check that a constant monomial is present
		if p_ii.ConstantMonomialIndex() == -1 {
			t.Errorf(
				"expected Plus() to return a PolynomialVector with Polynomial elements with a constant monomial; received %v",
				p_ii,
			)
		}
	}
}
