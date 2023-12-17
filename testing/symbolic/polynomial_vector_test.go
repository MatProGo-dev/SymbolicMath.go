package symbolic_test

/*
polynomial_vector_test.go
Description:
	Tests the methods defined in the polynomial_vector.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestPolynomialVector_Check1
Description:

	Verifies that the Check function returns an error when the polynomial vector
	is empty.
*/
func TestPolynomialVector_Check1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	err := pv.Check()
	if err == nil {
		t.Errorf(
			"Expected Check to return an error; received nil",
		)
	} else {
		if err.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Check to return error 'polynomial vector has no polynomials'; received '%v'",
				err.Error(),
			)
		}
	}
}

/*
TestPolynomialVector_Check2
Description:

	Verifies that the Check function returns an error when the polynomial vector
	in the twelfth index of a twenty-length polynomial vector is not properly
	initialized.
*/
func TestPolynomialVector_Check2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}
	for ii := 0; ii < 20; ii++ {
		if ii != 11 {
			pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
		}
	}

	// Test
	err := pv.Check()
	if err == nil {
		t.Errorf(
			"Expected Check to return an error; received nil",
		)
	} else {
		if err.Error() != "error in polynomial 11: polynomial has no monomials" {
			t.Errorf(
				"Expected Check to return error 'error in polynomial 11: polynomial has no monomials'; received '%v'",
				err.Error(),
			)
		}
	}
}

/*
TestPolynomialVector_Check3
Description:

	Verifies that a properly initialized polynomial vector returns no error when
	the Check function is called.
*/
func TestPolynomialVector_Check3(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}
	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Test
	err := pv.Check()
	if err != nil {
		t.Errorf(
			"Expected Check to return nil; received '%v'",
			err.Error(),
		)
	}
}

/*
TestPolynomialVector_Length1
Description:

	Tests that the Length method returns the correct value when the polynomial
	vector was properly defined with 20 elements.
*/
func TestPolynomialVector_Length1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Test
	if pv.Length() != 20 {
		t.Errorf(
			"Expected Length to return 20; received %v",
			pv.Length(),
		)
	}
}

/*
TestPolynomialVector_Length2
Description:

	Verifies that a panic is thrown if the Length method is called on a
	polynomial vector that has not been properly initialized.
*/
func TestPolynomialVector_Length2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Length to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Length to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Length to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Length()
}

/*
TestPolynomialVector_Len1
Description:

	Verifies that this produces the same result as the Length method
	for a properly defined polynomial vector.
*/
func TestPolynomialVector_Len1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Test
	if pv.Len() != 20 {
		t.Errorf(
			"Expected Len to return 20; received %v",
			pv.Len(),
		)
	}
}

/*
TestPolynomialVector_AtVec1
Description:

	Verifies that the AtVec method returns a polynomial type object when the
	method is called on a properly initialized object.
*/
func TestPolynomialVector_AtVec1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Test
	pvAtI := pv.AtVec(0)
	if _, ok := pvAtI.(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected AtVec to return a polynomial; received object of type '%T'",
			pvAtI,
		)
	}
}

/*
TestPolynomialVector_AtVec2
Description:

	Verifies that the AtVec method throws an error when a poorly chosen index is given.
	Matches the panic error produced with the expected one.
*/
func TestPolynomialVector_AtVec2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected AtVec to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected AtVec to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != smErrors.CheckIndexOnVector(20, pv).Error() {
			t.Errorf(
				"Expected AtVec to panic with error 'index out of range'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.AtVec(20)
}
