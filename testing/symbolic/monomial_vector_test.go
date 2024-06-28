package symbolic_test

import (
	"fmt"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"math"
	"strings"
	"testing"
)

/*
monomial_vector_test.go
Description:

	Tests for the functions mentioned in the monomial_vector.go file.
*/

/*
TestMonomialVector_Check1
Description:

	Tests that the Check() method properly catches an improperly initialized
	vector of Monomials (i.e., no monomials are given).
*/
func TestMonomialVector_Check1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}
	expectedError := smErrors.EmptyVectorError{mv}

	// Test
	err := mv.Check()
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}
}

/*
TestMonomialVector_Check2
Description:

	Tests that the Check() method properly catches an improperly initialized
	vector of Monomials (i.e., a monomial is given with an improper number of
	degrees).
*/
func TestMonomialVector_Check2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1, 2},
	}
	mv := symbolic.MonomialVector{m1}
	expectedError := fmt.Errorf(
		"the number of degrees (%v) does not match the number of variables (%v)",
		len(m1.Exponents),
		len(m1.VariableFactors),
	)

	// Test
	err := mv.Check()
	if !strings.Contains(
		err.Error(),
		expectedError.Error(),
	) {
		t.Errorf(
			"expected Check() to return \"%v\"; received %v",
			expectedError,
			err,
		)
	}
}

/*
TestMonomialVector_Check3
Description:

	Verifies that the Check() method returns nil when a constant is
	given as a monomial.
*/
func TestMonomialVector_Check3(t *testing.T) {
	// Constants
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}
	mv := symbolic.MonomialVector{m1}

	// Test
	if mv.Check() != nil {
		t.Errorf(
			"expected Check() to return nil; received %v",
			mv.Check(),
		)
	}
}

/*
TestMonomialVector_Variables1
Description:

	Verifies that the Variables() method panics when an improperly initialized
	vector of Monomials is given.
*/
func TestMonomialVector_Variables1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Variables() to panic; received %v",
				mv.Variables(),
			)
		}
	}()

	mv.Variables()
}

/*
TestMonomialVector_Variables2
Description:

	Verifies that the Variables() method returns the correct value when a
	vector of monomials of length 2 is given, with no repeated variables.
*/
func TestMonomialVector_Variables2(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	variables := mv.Variables()
	if len(variables) != 2 {
		t.Errorf(
			"expected len(variables) to be 2; received %v",
			len(variables),
		)
	}
}

/*
TestMonomialVector_Variables3
Description:

	Verifies that the Variables() method returns the correct value when a
	vector of monomials of length 2 is given, with repeated variables in
	each monomial.
*/
func TestMonomialVector_Variables3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	variables := mv.Variables()
	if len(variables) != 1 {
		t.Errorf(
			"expected len(variables) to be 1; received %v",
			len(variables),
		)
	}
}

/*
TestMonomialVector_Len1
Description:

	Verifies that the Len() method returns the correct value when a
	vector of monomials of length 2 is given.
*/
func TestMonomialVector_Len1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	if mv.Len() != 2 {
		t.Errorf(
			"expected mv.Len() to be 2; received %v",
			mv.Len(),
		)
	}
}

/*
TestMonomialVector_Dims1
Description:

	Verifies that the Dims() method returns the correct value when a
	vector of monomials of length 2 is given.
*/
func TestMonomialVector_Dims1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	if mv.Dims()[0] != 2 || mv.Dims()[1] != 1 {
		t.Errorf(
			"expected mv.Dims() to be [2,1]; received %v",
			mv.Dims(),
		)
	}
}

/*
TestMonomialVector_Constant1
Description:

	This test verifies that the Constant() method throws a panic
	when an improperly initialized vector of monomials is given.
*/
func TestMonomialVector_Constant1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Constant() to panic; received %v",
				mv.Constant(),
			)
		}
	}()

	mv.Constant()
}

/*
TestMonomialVector_Constant2
Description:

	This test verifies that the Constant() method returns the correct
	value (all zeros) when a vector of 4 monomials (each with nonzero
	number of variablefactors is given).
*/
func TestMonomialVector_Constant2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v3},
		Exponents:       []int{4},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m3 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 1},
	}
	m4 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2, v3, v4},
		Exponents:       []int{1, 1, 1, 1},
	}
	mv := symbolic.MonomialVector{m1, m2, m3, m4}

	// Test
	constant5 := mv.Constant()
	for ii := 0; ii < len(mv); ii++ {
		if constant5.AtVec(ii) != 0 {
			t.Errorf(
				"Expected mv.Constant() to be [0,0,0,0]; received %v at index %v",
				constant5.AtVec(ii),
				ii,
			)
		}

	}
}

/*
TestMonomialVector_Constant3
Description:

	This test verifies that the Constant() method returns the correct
	value (first two element nonzero) when a vector of 4 monomials
	is given and the first two elements are constants.
*/
func TestMonomialVector_Constant3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Exponents:       []int{},
	}
	m3 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Exponents:       []int{1, 1},
	}
	m4 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2, v3, v4},
		Exponents:       []int{1, 1, 1, 1},
	}
	mv := symbolic.MonomialVector{m1, m2, m3, m4}

	// Test
	constant5 := mv.Constant()
	for ii := 0; ii < len(mv); ii++ {
		if ii < 2 {
			if constant5.AtVec(ii) != 3.14 {
				t.Errorf(
					"Expected mv.Constant() to be [1,1,0,0]; received %v at index %v",
					constant5.AtVec(ii),
					ii,
				)
			}
		} else {
			if constant5.AtVec(ii) != 0 {
				t.Errorf(
					"Expected mv.Constant() to be [1,1,0,0]; received %v at index %v",
					constant5.AtVec(ii),
					ii,
				)
			}
		}
	}
}

/*
TestMonomialVector_Plus1
Description:

	Verifies that the Plus() method throws a panic when an improperly
	initialized vector of monomials is given.
*/
func TestMonomialVector_Plus1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Plus(3.14) to panic; received %v",
				mv.Plus(3.14),
			)
		}
	}()

	mv.Plus(3.14)
}

/*
TestMonomialVector_Plus2
Description:

	Verifies that the Plus() method throws a panic when a well-formed
	vector of monomials is added to an improperly initialized expression
	(in this case a monomial matrix).
*/
func TestMonomialVector_Plus2(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	pm := symbolic.MonomialMatrix{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Plus(pm) to panic; received %v",
				mv.Plus(pm),
			)
		}
	}()

	mv.Plus(pm)
}

/*
TestMonomialVector_Plus3
Description:

	Verifies that the Plus() method throws a panic when a well-formed
	vector of monomials is added to an well formed
	matrix of polynomials that do not have identical dimensions.
*/
func TestMonomialVector_Plus3(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	pm := symbolic.PolynomialMatrix{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Plus(pm) to panic; received %v",
				mv.Plus(pm),
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected mv.Plus(pm) to panic with an error; received %v",
				r,
			)
		}

		if !strings.Contains(
			rAsE.Error(),
			smErrors.DimensionError{
				Operation: "Plus",
				Arg1:      mv,
				Arg2:      pm,
			}.Error(),
		) {
			t.Errorf(
				"Expected mv.Plus(pm) to panic with an error containing \"dimensions\"; received %v",
				rAsE.Error(),
			)
		}
	}()

	mv.Plus(pm)
}

/*
TestMonomialVector_Plus4
Description:

	Verifies that the Plus() method returns the correct value when a
	well-formed vector of monomials is added to a well-formed
	vector of monomials.
*/
func TestMonomialVector_Plus4(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m2}
	mv2 := symbolic.MonomialVector{m1, m2}

	// Test
	sum := mv1.Plus(mv2)

	sumAsMV, tf := sum.(symbolic.MonomialVector)
	if !tf {
		t.Errorf(
			"expected sum to be a PolynomialVector; received %T",
			sum,
		)
	}

	// Check that each  monomial vector element contains
	// one variable factor and an exponent of 1
	for _, monomial := range sumAsMV {
		if len(monomial.VariableFactors) != 1 {
			t.Errorf(
				"expected len(monomial.VariableFactors) to be 1; received %v",
				len(monomial.VariableFactors),
			)
		}

		if monomial.Exponents[0] != 1 {
			t.Errorf(
				"expected monomial.Exponents[0] to be 1; received %v",
				monomial.Exponents[0],
			)
		}

	}

}

/*
TestMonomialVector_Plus5
Description:

	This test verifies that the method properly panics if a valid
	vector of monomials is multiplied by an invalid expression
	(in this case a matrix of monomials).
*/
func TestMonomialVector_Plus5(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m2}

	var mm symbolic.MonomialMatrix

	// Setup defer function for catching panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv1.Plus(mm) to panic; received %v",
				mv1.Plus(mm),
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected mv1.Plus(mm) to panic with an error; received %v",
				r,
			)
		}

		if !strings.Contains(
			rAsE.Error(),
			mm.Check().Error(),
		) {
			t.Errorf(
				"Expected mv1.Plus(mm) to panic with an error containing \"%v\"; received %v",
				mm.Check().Error(),
				rAsE.Error(),
			)
		}
	}()

	// Test
	mv1.Plus(mm)
}

/*
TestMonomialVector_Plus6
Description:

	This test verifies that the method properly creates a monomial vector,
	if the current monomial vector contains all constants.
*/
func TestMonomialVector_Plus6(t *testing.T) {
	// Constants
	k2 := symbolic.K(3.14)

	// Create a monomial vector of constants
	kv1 := getKVector.From([]float64{1, 2, 3, 4, 5})
	mv1 := kv1.ToMonomialVector()

	// Compute Sum
	sum := mv1.Plus(k2)

	// Verify that the sum is a monomial vector
	if _, tf := sum.(symbolic.KVector); !tf {
		t.Errorf(
			"expected sum to be a MonomialVector; received %T",
			sum,
		)
	}

}

/*
TestMonomialVector_Plus7
Description:

	This test verifies that the method properly creates a polynomial vector
	when the monomial vector is added to a constant AND the monomial vector
	is not already a constant vector.
*/
func TestMonomialVector_Plus7(t *testing.T) {
	// Constants
	N := 10
	vv1 := symbolic.NewVariableVector(N)
	mv1 := vv1.ToMonomialVector()
	k2 := symbolic.K(3.14)

	// Compute Sum
	sum := mv1.Plus(k2)

	// Verify that the sum is a polynomial vector
	if _, tf := sum.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"expected sum to be a PolynomialVector; received %T",
			sum,
		)
	}

	// Check that each element of the polynomial vector
	// contains 2 monomials
	for _, polynomial := range sum.(symbolic.PolynomialVector) {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"expected len(polynomial.Monomials) to be 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestMonomialVector_Plus8
Description:

	This test verifies that the method properly panics when a
	monomial vector is added to an invalid expression (in this case
	a string).
*/
func TestMonomialVector_Plus8(t *testing.T) {
	// Constants
	mv := symbolic.NewVariableVector(10).ToMonomialVector()
	s2 := "This is a test string."

	// Setup defer function for catching panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Plus(s2) to panic; received %v",
				mv.Plus(s2),
			)
		}
	}()

	// Test
	mv.Plus(s2)
}

/*
TestMonomialVector_Plus9
Description:

	This test verifies that the method properly produces
	a monomial vector when a monomial vector is added to a
	monomial vector that matches each monomial in the vector.
*/
func TestMonomialVector_Plus9(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m2}
	mv2 := symbolic.MonomialVector{m1, m2}

	// Test
	sum := mv1.Plus(mv2)

	// Verify that the sum is a monomial vector
	if _, tf := sum.(symbolic.MonomialVector); !tf {
		t.Errorf(
			"expected sum to be a MonomialVector; received %T",
			sum,
		)
	}
}

/*
TestMonomialVector_Plus10
Description:

	This test verifies that the method properly produces
	a polynomial vector when a monomial vector is added to a
	monomial vector that does not match each monomial in the vector.
*/
func TestMonomialVector_Plus10(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m2}
	mv2 := symbolic.MonomialVector{m1, m1}

	// Test
	sum := mv1.Plus(mv2)

	// Verify that the sum is a polynomial vector
	if _, tf := sum.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"expected sum to be a PolynomialVector; received %T",
			sum,
		)
	}
}

/*
TestMonomialVector_Plus11
Description:

	This test verifies that the method properly produces
	a monomial vector when a monomial vector is added to a
	single monomial that matches each monomial in the vector.
*/
func TestMonomialVector_Plus11(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m1}

	// Test
	sum := mv1.Plus(m1)

	// Verify that the sum is a monomial vector
	if _, tf := sum.(symbolic.MonomialVector); !tf {
		t.Errorf(
			"expected sum to be a MonomialVector; received %T",
			sum,
		)
	}
}

/*
TestMonomialVector_Plus12
Description:

	This test verifies that the method properly produces
	a polynomial vector when a monomial vector is added to a
	single monomial that does not match each monomial in the vector.
*/
func TestMonomialVector_Plus12(t *testing.T) {
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
		Exponents:       []int{1},
	}
	mv1 := symbolic.MonomialVector{m1, m1}

	// Test
	sum := mv1.Plus(m2)

	// Verify that the sum is a polynomial vector
	if _, tf := sum.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"expected sum to be a PolynomialVector; received %T",
			sum,
		)
	}
}

/*
TestMonomialVector_Minus1
Description:

	Verifies that the Minus() method throws a panic when an improperly
	initialized vector of monomials is given.
*/
func TestMonomialVector_Minus1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Minus(3.14) to panic; received %v",
				mv.Minus(3.14),
			)
		}
	}()

	mv.Minus(3.14)
}

/*
TestMonomialVector_Minus2
Description:

	Verifies that the Minus() method throws a panic when a well-formed
	vector of monomials is subtracted from an improperly initialized expression
	(in this case a monomial matrix).
*/
func TestMonomialVector_Minus2(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	pm := symbolic.MonomialMatrix{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Minus(pm) to panic; received %v",
				mv.Minus(pm),
			)
		}
	}()

	mv.Minus(pm)
}

/*
TestMonomialVector_Minus3
Description:

	Verifies that the Minus() method throws a panic when a well-formed
	vector of monomials is subtracted from a well formed vector expression
	OF THE WRONG DIMENSIONs.
*/
func TestMonomialVector_Minus3(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	pm := symbolic.PolynomialMatrix{
		{
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
			symbolic.NewVariable().ToPolynomial(),
		},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Minus(pm) to panic; received %v",
				mv.Minus(pm),
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected mv.Minus(pm) to panic with an error; received %v",
				r,
			)
		}

		if !strings.Contains(
			rAsE.Error(),
			smErrors.DimensionError{
				Operation: "Minus",
				Arg1:      mv,
				Arg2:      pm,
			}.Error(),
		) {
			t.Errorf(
				"Expected mv.Minus(pm) to panic with an error containing \"dimensions\"; received %v",
				rAsE.Error(),
			)
		}
	}()

	mv.Minus(pm)
}

/*
TestMonomialVector_Minus4
Description:

	Verifies that the Minus() method returns the correct value when a
	well-formed vector of monomials is subtracted from a float64.
	The result should be a polynomial vector where each polynomial contains
	two monomials, one that is a constant and one that is a variable.
*/
func TestMonomialVector_Minus4(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	mv := symbolic.MonomialVector{v1.ToMonomial(), v1.ToMonomial()}
	f2 := 3.14

	// Test
	difference := mv.Minus(f2)

	// Verify that the difference is a polynomial vector
	if _, tf := difference.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"expected difference to be a PolynomialVector; received %T",
			difference,
		)
	}

	// Verify that each polynomial contains two monomials
	for _, polynomial := range difference.(symbolic.PolynomialVector) {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"expected len(polynomial.Monomials) to be 2; received %v",
				len(polynomial.Monomials),
			)
		}

		for _, monomial := range polynomial.Monomials {
			if (!monomial.IsConstant()) && (!monomial.IsVariable(v1)) {
				t.Errorf("expected monomial to be a variable or a constant; received %v", monomial)
			}
		}

	}
}

/*
TestMonomialVector_Minus5
Description:

	Verifies that the Minus() method returns the correct value when a
	well-formed vector of monomials is subtracted from a *mat.VecDense object
	of appropriate dimension.
*/
func TestMonomialVector_Minus5(t *testing.T) {
	// Setup
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	mv := symbolic.MonomialVector{v1.ToMonomial(), v2.ToMonomial()}

	// Create a *mat.VecDense object
	vec := mat.NewVecDense(2, []float64{1, 2})

	// Test
	difference := mv.Minus(vec)

	// Verify that the difference is a polynomial vector
	if _, tf := difference.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"expected difference to be a PolynomialVector; received %T",
			difference,
		)
	}

	// Verify that each polynomial contains two monomials
	for ii, polynomial := range difference.(symbolic.PolynomialVector) {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"expected len(polynomial.Monomials) to be 2; received %v",
				len(polynomial.Monomials),
			)
		}

		// Verify that each monomial is the correct value
		for _, monomial := range polynomial.Monomials {
			if monomial.IsConstant() {
				switch ii {
				case 0:
					if monomial.Coefficient != -1.0 {
						t.Errorf(
							"expected monomial.Coefficient to be -1.0; received %v",
							monomial.Coefficient,
						)
					}
				case 1:
					if monomial.Coefficient != -2.0 {
						t.Errorf(
							"expected monomial.Coefficient to be -2.0; received %v",
							monomial.Coefficient,
						)
					}
				}

			}
		}
	}
}

/*
TestMonomialVector_Multiply1
Description:

	Verifies that the Multiply() method throws a panic when an improperly
	initialized vector of monomials is given.
*/
func TestMonomialVector_Multiply1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Multiply(3.14) to panic; received %v",
				mv.Multiply(3.14),
			)
		}
	}()

	mv.Multiply(3.14)
}

/*
TestMonomialVector_Multiply2
Description:

	Verifies that the Multiply() method throws a panic when a well-formed
	vector of monomials is multiplied by an improperly initialized expression
	(in this case a monomial matrix).
*/
func TestMonomialVector_Multiply2(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	mm2 := symbolic.MonomialMatrix{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Multiply(mm2) to panic; received %v",
				mv.Multiply(mm2),
			)
		}
	}()

	mv.Multiply(mm2)
}

/*
TestMonomialVector_Multiply3
Description:

	Verifies that the Multiply() method throws a panic when a well-formed
	vector of monomials is multiplied by a well formed monomial matrix
	that does not have compatible dimensions.
*/
func TestMonomialVector_Multiply3(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	mm2 := symbolic.MonomialMatrix{
		{
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
		},
		{
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
			symbolic.NewVariable().ToMonomial(),
		},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Multiply(mm2) to panic; received %v",
				mv.Multiply(mm2),
			)
		}
	}()

	mv.Multiply(mm2)
}

/*
TestMonomialVector_Multiply4
Description:

	Verifies that the Multiply() method returns the correct value when a
	well-formed vector of monomials is multiplied by a float64.
	We will verify that the result is a monomial vector where each
	monomials coefficient is multiplied by the float64.
*/
func TestMonomialVector_Multiply4(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	f2 := 3.14

	// Test
	product := mv.Multiply(f2)

	// Verify that the product is a monomial vector
	if _, tf := product.(symbolic.MonomialVector); !tf {
		t.Errorf(
			"expected product to be a MonomialVector; received %T",
			product,
		)
	}

	// Verify that the coefficients are correct
	for _, monomial := range product.(symbolic.MonomialVector) {
		if monomial.Coefficient != 3.14 {
			t.Errorf(
				"expected monomial.Coefficient to be 3.14; received %v",
				monomial.Coefficient,
			)
		}
	}
}

/*
TestMonomialVector_Multiply5
Description:

	Verifies that the Multiply() method returns the correct value when a
	well-formed vector of monomials is multiplied by an invalid non-expression
	(string).
*/
func TestMonomialVector_Multiply5(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}
	s2 := "This is a test string."

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Multiply(s2) to panic; received %v",
				mv.Multiply(s2),
			)
		}
	}()

	mv.Multiply(s2)
}

/*
TestMonomialVector_Transpose1
Description:

	Verifies that the Transpose() method throws a panic when an improperly
	initialized vector of monomials is given.
*/
func TestMonomialVector_Transpose1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Transpose() to panic; received %v",
				mv.Transpose(),
			)
		}
	}()

	mv.Transpose()
}

/*
TestMonomialVector_Transpose2
Description:

	Verifies that the Transpose() method returns the correct value when a
	well-formed vector of monomials is given.
	Checks that the dimensions of the transposed vector are correct.
*/
func TestMonomialVector_Transpose2(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
		symbolic.NewVariable().ToMonomial(),
	}

	// Test
	transposed := mv.Transpose()
	mvT, tf := transposed.(symbolic.MonomialMatrix)
	if !tf {
		t.Errorf(
			"expected mvT to be a MonomialMatrix; received %T",
			transposed,
		)
	}

	if mvT.Dims()[0] != mv.Dims()[1] || mvT.Dims()[1] != mv.Dims()[0] {
		t.Errorf(
			"expected mvT.Dims() to be [%v,%v]; received %v",
			mv.Dims()[1],
			mv.Dims()[0],
			mvT.Dims(),
		)
	}
}

/*
TestMonomial_LessEq1
Description:

	Verifies that the LessEq() method returns the correct value when
	compared to a KVector.
*/
func TestMonomial_LessEq1(t *testing.T) {
	// Constants
	N := 4
	vv1 := symbolic.NewVariableVector(N)
	mv1 := vv1.ToMonomialVector()

	kv2 := getKVector.From([]float64{1, 2, 3, 4})

	// Create a vector constraint
	constraint := mv1.LessEq(kv2)

	// verify that the cosntraint is a vector constraint
	vc2, tf := constraint.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf("expected constraint to be a VectorConstraint; received %T", constraint)
	}

	// Verify that the constraint's dimensions are correct
	if vc2.Dims()[0] != N {
		t.Errorf(
			"expected constraint.Dims()[0] to be %v; received %v",
			N,
			vc2.Dims()[0],
		)
	}
}

/*
TestMonomial_GreaterEq1
Description:

	Verifies that the GreaterEq() method returns the correct value when
	compared to a K.
*/
func TestMonomial_GreaterEq1(t *testing.T) {
	// Constants
	N := 5
	vv1 := symbolic.NewVariableVector(N)
	mv1 := vv1.ToMonomialVector()

	k2 := symbolic.K(3.14)

	// Create a vector constraint
	constraint := mv1.GreaterEq(k2)

	// verify that the cosntraint is a vector constraint
	vc2, tf := constraint.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf("expected constraint to be a VectorConstraint; received %T", constraint)
	}

	// Verify that the constraint's dimensions are correct
	if vc2.Dims()[0] != N {
		t.Errorf(
			"expected constraint.Dims()[0] to be %v; received %v",
			N,
			vc2.Dims()[0],
		)
	}
}

/*
TestMonomial_Eq1
Description:

	Verifies that the Eq() method panics when called with an improperly
	initialized vector of monomials.
*/
func TestMonomial_Eq1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Eq(3.14) to panic; received %v",
				mv.Eq(3.14),
			)
		}
	}()

	mv.Eq(3.14)
}

/*
TestMonomialVector_Eq2
Description:

	Verifies that the Eq() method returns the correct value when
	compared to a well-defined and appropriately sized VariableVector.
	We should have a vector constraint with:
	- The matching dimensions to the MonomialVector
	- ConstrSense() of SenseEqual
	- Right hand side of the constraint equal to the VariableVector
*/
func TestMonomialVector_Eq2(t *testing.T) {
	// Constants
	N := 5
	vv1 := symbolic.NewVariableVector(N)
	mv1 := vv1.ToMonomialVector()

	// Create a vector constraint
	constraint := mv1.Eq(vv1)

	// verify that the cosntraint is a vector constraint
	vc2, tf := constraint.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf("expected constraint to be a VectorConstraint; received %T", constraint)
	}

	// Verify that the constraint's dimensions are correct
	if vc2.Dims()[0] != N {
		t.Errorf(
			"expected constraint.Dims()[0] to be %v; received %v",
			N,
			vc2.Dims()[0],
		)
	}

	// Verify that the constraint's sense is correct
	if vc2.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"expected constraint.ConstrSense() to be %v; received %v",
			symbolic.SenseEqual,
			vc2.ConstrSense(),
		)
	}

	// Verify that the constraint's right hand side is of the
	// correct type (VariableVector)
	if _, tf := vc2.Right().(symbolic.VariableVector); !tf {
		t.Errorf(
			"expected vc2.RightHandSide() to be a VariableVector; received %T",
			vc2.Right(),
		)
	}

}

/*
TestMonomialVector_Comparison1
Description:

	Verifies that the Comparison() method panics
	when a well-defined monomial is compared to an object
	that is not an expression (in this case a string).
*/
func TestMonomialVector_Comparison1(t *testing.T) {
	// Constants
	mv := symbolic.NewVariableVector(10).ToMonomialVector()
	s2 := "This is a test string."

	// Setup defer function for catching panic
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Comparison(s2) to panic; received %v",
				mv.Comparison(s2, symbolic.SenseLessThanEqual),
			)
		}
	}()

	// Test
	mv.Comparison(s2, symbolic.SenseLessThanEqual)
}

/*
TestMonomialVector_Comparison2
Description:

	Verifies that the Comparison() method returns the correct value
	when a well-defined monomial vector is compared with a well-defined
	monomial vector.
*/
func TestMonomialVector_Comparison2(t *testing.T) {
	// Constants
	mv1 := symbolic.NewVariableVector(10).ToMonomialVector()
	mv2 := symbolic.NewVariableVector(10).ToMonomialVector()

	// Test
	constraint := mv1.Comparison(mv2, symbolic.SenseLessThanEqual)

	// Verify that the constraint is a vector constraint
	vc2, tf := constraint.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf(
			"expected constraint to be a VectorConstraint; received %T",
			constraint,
		)
	}

	// Verify that the constraint's dimensions are correct
	if vc2.Dims()[0] != 10 {
		t.Errorf(
			"expected constraint.Dims()[0] to be 10; received %v",
			vc2.Dims()[0],
		)
	}

	// Verify that the constraint's sense is correct
	if vc2.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf(
			"expected constraint.ConstrSense() to be %v; received %v",
			symbolic.SenseLessThanEqual,
			vc2.ConstrSense(),
		)
	}

	// Verify that the constraint's right hand side is of the
	// correct type (MonomialVector)
	if _, tf := vc2.Right().(symbolic.MonomialVector); !tf {
		t.Errorf(
			"expected vc2.RightHandSide() to be a MonomialVector; received %T",
			vc2.Right(),
		)
	}
}

/*
TestMonomialVector_Derivative1
Description:

	Verifies that the Derivative() method panics when called with an
	improperly initialized vector of monomials.
*/
func TestMonomialVector_Derivative1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Derivative() to panic; received %v",
				mv.DerivativeWrt(symbolic.NewVariable()),
			)
		}
	}()

	mv.DerivativeWrt(symbolic.NewVariable())
}

/*
TestMonomialVector_Derivative2
Description:

	Verifies that the Derivative() method returns the correct value when
	called with a well-defined vector of monomial, and a variable
	that is NOT contained in any element of th emonomial.
	(Result should be all zeros.)
*/
func TestMonomialVector_Derivative2(t *testing.T) {
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
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	derivative := mv.DerivativeWrt(symbolic.NewVariable())

	// Verify that the derivative is a monomial vector
	if _, tf := derivative.(symbolic.KVector); !tf {
		t.Errorf(
			"expected derivative to be a MonomialVector; received %T",
			derivative,
		)
	}

	// Verify that the derivative is all zeros
	for _, tempK := range derivative.(symbolic.KVector) {
		if tempK != symbolic.K(0) {
			t.Errorf(
				"expected monomial.Coefficient to be 0; received %v",
				tempK,
			)
		}
	}
}

/*
TestMonomialVector_Derivative3
Description:

	Verifies that the Derivative() method returns the correct value when
	called with a well-defined vector of monomial, and a variable
	that is contained in each element of the monomial.
*/
func TestMonomialVector_Derivative3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	derivative := mv.DerivativeWrt(v1)

	// Verify that the derivative is a K vector
	if _, tf := derivative.(symbolic.KVector); !tf {
		t.Errorf(
			"expected derivative to be a MonomialVector; received %T",
			derivative,
		)
	}

	// Verify that each element of the derivative is just the coefficient
	// from the original monomial vector mv
	for ii, d_ii := range derivative.(symbolic.KVector) {
		if float64(d_ii) != mv[ii].Coefficient {
			t.Errorf(
				"expected constant to be %v; received %v",
				mv[ii].Coefficient,
				float64(d_ii),
			)
		}
	}

}

/*
TestMonomialVector_Derivative4
Description:

	Verifies that the Derivative() method returns the correct value when
	called with a well-defined vector of monomial, and a variable
	that is contained in some elements of the monomial.
*/
func TestMonomialVector_Derivative4(t *testing.T) {
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
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	derivative := mv.DerivativeWrt(v1)

	// Verify that the derivative is a K vector
	d_v1, tf := derivative.(symbolic.KVector)
	if !tf {
		t.Errorf(
			"expected derivative to be a MonomialVector; received %T",
			derivative,
		)
	}

	// Verify that the first element of the derivative is a constant and nonzero
	if float64(d_v1[0]) != 3.14 {
		t.Errorf(
			"expected derivative[0].Coefficient to be 3.14; received %v",
			d_v1[0],
		)
	}

	// Verify that the second element of the derivative is a constant and zero
	if float64(d_v1[1]) != 0 {
		t.Errorf(
			"expected derivative[1].Coefficient to be 0; received %v",
			d_v1[1],
		)
	}

}

/*
TestMonomialVector_Degree1
Description:

	Verifies that the Degree() method panics when this is called with a MonomialVector that is not well-defined.
*/
func TestMonomialVector_Degree1(t *testing.T) {
	// Constants
	mv := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Degree() to panic; received %v",
				mv.Degree(),
			)
		}

		err := r.(error)
		expectedError := mv.Check()
		if err.Error() != expectedError.Error() {
			t.Errorf(
				"Expected error to be %v; received %v",
				expectedError,
				err,
			)
		}

	}()

	mv.Degree()
	t.Errorf("Test should panic before this is reached!")
}

/*
TestMonomialVector_Degree2
Description:

	Verifies that the Degree() method returns the correct value when called with a well-defined MonomialVector.
	This should be the maximum of all of the monomials inside the vector.
*/
func TestMonomialVector_Degree2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{2},
	}
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v2},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1, m2}

	// Test
	degree := mv.Degree()

	// Verify that the degree is correct
	if degree != 2 {
		t.Errorf(
			"expected degree to be 2; received %v",
			degree,
		)
	}
}

/*
TestMonomialVector_Power1
Description:

	Verifies that the Power() method panics when called with a vector that has Len() greater
	than 1.
*/
func TestMonomialVector_Power1(t *testing.T) {
	// Constants
	mv := symbolic.NewVariableVector(10).ToMonomialVector()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected mv.Power(3) to panic; received %v",
				mv.Power(3),
			)
		}
	}()

	mv.Power(3)
	t.Errorf("Test should panic before this is reached!")
}

/*
TestMonomialVector_Power2
Description:

	Verifies that the Power() method returns the correct value when called with a vector
	of Len() == 1. The result should be a monomial vector where each monomial is raised to
	the power of the input integer.
*/
func TestMonomialVector_Power2(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	m1 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{v1},
		Exponents:       []int{1},
	}
	mv := symbolic.MonomialVector{m1}

	// Test
	power := mv.Power(3)

	// Verify that the power is a monomial vector
	if _, tf := power.(symbolic.Monomial); !tf {
		t.Errorf(
			"expected power to be a MonomialVector; received %T",
			power,
		)
	}

	// Verify that each monomial is raised to the power of 3
	if power.(symbolic.Monomial).Exponents[0] != 3 {
		t.Errorf(
			"expected monomial.Exponents[0] to be 3; received %v",
			power.(symbolic.Monomial).Exponents[0],
		)
	}

	if power.(symbolic.Monomial).Coefficient != math.Pow(3.14, 3) {
		t.Errorf(
			"expected monomial.Coefficient to be 3.14^3; received %v",
			power.(symbolic.Monomial).Coefficient,
		)
	}
}
