package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
		Degrees:         []int{1, 2},
	}
	mv := symbolic.MonomialVector{m1}
	expectedError := fmt.Errorf(
		"the number of degrees (%v) does not match the number of variables (%v)",
		len(m1.Degrees),
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
		Degrees:         []int{},
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
		Degrees:         []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v2},
		Degrees:         []int{1},
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
		Degrees:         []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1},
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
		Degrees:         []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1},
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
		Degrees:         []int{1},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1},
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
		Degrees:         []int{4},
	}
	m2 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1},
	}
	m3 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Degrees:         []int{1, 1},
	}
	m4 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2, v3, v4},
		Degrees:         []int{1, 1, 1, 1},
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
		Degrees:         []int{},
	}
	m2 := symbolic.Monomial{
		Coefficient:     3.14,
		VariableFactors: []symbolic.Variable{},
		Degrees:         []int{},
	}
	m3 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Degrees:         []int{1, 1},
	}
	m4 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2, v3, v4},
		Degrees:         []int{1, 1, 1, 1},
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
