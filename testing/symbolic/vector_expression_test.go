package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
vector_expression_test.go
Description:

	Tests for the functions mentioned in the vector_expression.go file.
*/

/*
TestVectorExpression_ToVectorExpression1
Description:

	Tests the conversion of a K() to a vector expression.
	An error should be returned along with a small KVector.
*/
func TestVectorExpression_ToVectorExpression1(t *testing.T) {
	// Constants
	k := symbolic.K(2)

	// Test
	_, err := symbolic.ToVectorExpression(k)
	if err == nil {
		t.Errorf(
			"Expected an error to be returned; received nil",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the input interface is of type %T, which is not recognized as a VectorExpression.",
				k,
			),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				fmt.Sprintf("the input interface is of type %T, which is not recognized as a VectorExpression.",
					k,
				),
				err.Error(),
			)

		}
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression1
Description:

	Tests the conversion of a slice of constants (K) to a KVector.
*/
func TestVectorExpression_ConcretizeVectorExpression1(t *testing.T) {
	// Constants
	k1 := symbolic.K(2)
	k2 := symbolic.K(3)
	k3 := symbolic.K(4)
	k4 := symbolic.K(5)
	slice := []symbolic.ScalarExpression{k1, k2, k3, k4}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.KVector); !tf {
		t.Errorf("expected a KVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression2
Description:

	Tests the conversion of a slice of all variables to a VariableVector.
*/
func TestVectorExpression_ConcretizeVectorExpression2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()
	slice := []symbolic.ScalarExpression{v1, v2, v3, v4}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.VariableVector); !tf {
		t.Errorf("expected a VariableVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression3
Description:

	Tests the conversion of a slice of all monomials to a MonomialVector.
*/
func TestVectorExpression_ConcretizeVectorExpression3(t *testing.T) {
	// Constants
	m1 := symbolic.NewVariable().ToMonomial()
	m2 := symbolic.NewVariable().ToMonomial()
	m3 := symbolic.NewVariable().ToMonomial()
	m4 := symbolic.NewVariable().ToMonomial()
	slice := []symbolic.ScalarExpression{m1, m2, m3, m4}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.MonomialVector); !tf {
		t.Errorf("expected a MonomialVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression4
Description:

	Tests the conversion of a slice of all polynomials to a PolynomialVector.
*/
func TestVectorExpression_ConcretizeVectorExpression4(t *testing.T) {
	// Constants
	p1 := symbolic.NewVariable().ToPolynomial()
	p2 := symbolic.NewVariable().ToPolynomial()
	p3 := symbolic.NewVariable().ToPolynomial()
	p4 := symbolic.NewVariable().ToPolynomial()
	slice := []symbolic.ScalarExpression{p1, p2, p3, p4}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.PolynomialVector); !tf {
		t.Errorf("expected a PolynomialVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression5
Description:

	Tests the conversion of a slice containing constant (K) and variable expressions to a
	vector of monomials (MonomialVector).
*/
func TestVectorExpression_ConcretizeVectorExpression5(t *testing.T) {
	// Constants
	k1 := symbolic.K(2)
	v1 := symbolic.NewVariable()
	k2 := symbolic.K(3)
	v2 := symbolic.NewVariable()
	k3 := symbolic.K(4)
	v3 := symbolic.NewVariable()
	k4 := symbolic.K(5)
	v4 := symbolic.NewVariable()
	slice := []symbolic.ScalarExpression{k1, v1, k2, v2, k3, v3, k4, v4}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.MonomialVector); !tf {
		t.Errorf("expected a MonomialVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression6
Description:

	Tests the conversion of a slice containing a constant and a polynomial to a PolynomialVector.
*/
func TestVectorExpression_ConcretizeVectorExpression6(t *testing.T) {
	// Constants
	k := symbolic.K(2)
	p := symbolic.NewVariable().ToPolynomial()
	slice := []symbolic.ScalarExpression{k, p}

	// Test
	v := symbolic.ConcretizeVectorExpression(slice)
	if _, tf := v.(symbolic.PolynomialVector); !tf {
		t.Errorf("expected a PolynomialVector; received %T", v)
	}
}

/*
TestVectorExpression_ConcretizeVectorExpression7
Description:

	Tests that the function panics when the input slice is empty.
*/
func TestVectorExpression_ConcretizeVectorExpression7(t *testing.T) {
	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic when the input slice is empty; received nil")
		}
	}()
	symbolic.ConcretizeVectorExpression([]symbolic.ScalarExpression{})
	t.Errorf("Problem! The function did not panic when the input slice was empty")
}

/*
TestVectorExpression_ConcretizeVectorExpression8
Description:

	Tests that the function correctly generates a polynomial vector from a slice of ScalarExpression objects,
	one is a polynomial, one is a monomial, and the other is a variable.
*/
func TestVectorExpression_ConcretizeVectorExpression8(t *testing.T) {
	// Constants
	p := symbolic.NewVariable().ToPolynomial()
	m := symbolic.NewVariable().ToMonomial()
	v := symbolic.NewVariable()
	slice := []symbolic.ScalarExpression{p, m, v}

	// Test
	pv := symbolic.ConcretizeVectorExpression(slice)

	vOut, tf := pv.(symbolic.PolynomialVector)
	if !tf {
		t.Errorf("expected a PolynomialVector; received %T", pv)
	}

	if len(vOut) != 3 {
		t.Errorf("expected a PolynomialVector of length 3; received %v", len(vOut))
	}
}

/*
TestVectorExpression_VectorSubstituteTemplate1
Description:

	Verify that the VectorSubstituteTemplate() function panics when the input vector is
	not a well-defined VectorExpression.
*/
func TestVectorExpression_VectorSubstituteTemplate1(t *testing.T) {
	// Constants
	testVec := symbolic.VariableVector{
		symbolic.NewVariable(),
		symbolic.Variable{},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic when the input vector is not a well-defined VectorExpression; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		if !strings.Contains(
			rAsE.Error(),
			testVec.Check().Error(),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				"unexpected expression type in vector expression: symbolic.Variable",
				rAsE.Error(),
			)
		}
	}()
	symbolic.VectorSubstituteTemplate(testVec, symbolic.NewVariable(), symbolic.NewVariable())
	t.Errorf("Problem! The function did not panic when the input vector was not a well-defined VectorExpression")
}

/*
TestVectorExpression_VectorSubstituteTemplate2
Description:

	Verify that the VectorSubstituteTemplate() function panics when the input vector is
	well-defined but the input target variable is not well-defined.
*/
func TestVectorExpression_VectorSubstituteTemplate2(t *testing.T) {
	// Constants
	testVec := symbolic.VariableVector{
		symbolic.NewVariable(),
		symbolic.NewVariable(),
	}
	badVar := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic when the input target variable is not well-defined; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		if !strings.Contains(
			rAsE.Error(),
			badVar.Check().Error(),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				badVar.Check().Error(),
				rAsE.Error(),
			)
		}
	}()
	symbolic.VectorSubstituteTemplate(testVec, badVar, symbolic.NewVariable())
	t.Errorf("Problem! The function did not panic when the input target variable was not well-defined")
}

/*
TestVectorExpression_VectorSubstituteTemplate3
Description:

	Verify that the VectorSubstituteTemplate() function panics when the input vector is
	well-defined, the input target variable is well-defined, but the input scalar expression
	is not well-defined.
*/
func TestVectorExpression_VectorSubstituteTemplate3(t *testing.T) {
	// Constants
	testVec := symbolic.VariableVector{
		symbolic.NewVariable(),
		symbolic.NewVariable(),
	}
	badSE := symbolic.Variable{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic when the input scalar expression is not well-defined; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		if !strings.Contains(
			rAsE.Error(),
			badSE.Check().Error(),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				badSE.Check().Error(),
				rAsE.Error(),
			)
		}
	}()
	symbolic.VectorSubstituteTemplate(testVec, symbolic.NewVariable(), badSE)
	t.Errorf("Problem! The function did not panic when the input scalar expression was not well-defined")
}

/*
TestVectorExpression_VectorPowerTemplate1
Description:

	Verify that the VectorPowerTemplate() function panics when the input vector is
	not a well-defined VectorExpression.
*/
func TestVectorExpression_VectorPowerTemplate1(t *testing.T) {
	// Constants
	testVec := symbolic.VariableVector{
		symbolic.NewVariable(),
		symbolic.Variable{},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic when the input vector is not a well-defined VectorExpression; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		if !strings.Contains(
			rAsE.Error(),
			testVec.Check().Error(),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				"unexpected expression type in vector expression: symbolic.Variable",
				rAsE.Error(),
			)
		}
	}()
	symbolic.VectorPowerTemplate(testVec, 2)
	t.Errorf("Problem! The function did not panic when the input vector was not a well-defined VectorExpression")
}

/*
TestVectorExpression_VectorPowerTemplate2
Description:

	Verify that the VectorPowerTemplate() function panics when the input vector is
	well-defined but the input power is less than 0.
*/
func TestVectorExpression_VectorPowerTemplate2(t *testing.T) {
	// Constants
	testVec := symbolic.VariableVector{
		symbolic.NewVariable(),
		symbolic.NewVariable(),
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic when the input power is less than 0; received nil")
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		expectedError := smErrors.NegativeExponentError{
			-1,
		}
		if !strings.Contains(
			rAsE.Error(),
			expectedError.Error(),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				expectedError.Error(),
				rAsE.Error(),
			)
		}
	}()
	symbolic.VectorPowerTemplate(testVec, -1)
	t.Errorf("Problem! The function did not panic when the input power was less than 0")
}
