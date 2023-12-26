package symbolic_test

/*
polynomial_vector_test.go
Description:
	Tests the methods defined in the polynomial_vector.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
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

/*
TestPolynomialVector_Variables1
Description:

	Verifies that the number of variables found in a polynomial vector containing all constant
	elements is zero.
*/
func TestPolynomialVector_Variables1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		tempK := symbolic.K(1)
		pv.Elements[ii] = tempK.ToMonomial().ToPolynomial()
	}

	// Test
	if len(pv.Variables()) != 0 {
		t.Errorf(
			"Expected Variables to return 0; received %v",
			pv.Variables(),
		)
	}
}

/*
TestPolynomialVector_Variables2
Description:

	Verifies that the number of variables found in a polynomial vector containing a number of variables
	that matches the second polynomial.
*/
func TestPolynomialVector_Variables2(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()

	pv := symbolic.PolynomialVector{
		Elements: []symbolic.Polynomial{
			k1.ToMonomial().ToPolynomial(),
			symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3}, Degrees: []int{1, 2}}.ToPolynomial(),
			symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3, v4}, Degrees: []int{3, 5, 11}}.ToPolynomial(),
		},
	}

	// Check that there are 3 variables in pv
	if len(pv.Variables()) != 3 {
		t.Errorf(
			"Expected Variables to return 3; received %v",
			pv.Variables(),
		)
	}
}

/*
TestPolynomialVector_Variables3
Description:

	Verifies that the number of variables found in a polynomial vector containing a number of variables
	that doesn't match any individual polynomial but correctly captures the union of all variables.
*/
func TestPolynomialVector_Variables3(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)
	v2 := symbolic.NewVariable()
	v3 := symbolic.NewVariable()
	v4 := symbolic.NewVariable()

	pv := symbolic.PolynomialVector{
		Elements: []symbolic.Polynomial{
			k1.ToMonomial().ToPolynomial(),
			symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3}, Degrees: []int{1, 2}}.ToPolynomial(),
			symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v4}, Degrees: []int{3, 11}}.ToPolynomial(),
		},
	}

	// Check that there are 3 variables in pv
	if len(pv.Variables()) != 3 {
		t.Errorf(
			"Expected Variables to return 3; received %v",
			pv.Variables(),
		)
	}
}

/*
TestPolynomialVector_Constant1
Description:

	Verifies that the Constant method panics if the polynomial vector is not properly
	initialized.
*/
func TestPolynomialVector_Constant1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Constant to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Constant to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Constant to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Constant()
}

/*
TestPolynomialVector_Constant2
Description:

	Tests that the constant method returns all zeros when the polynomial vector
	contains ALL monomials.
*/
func TestPolynomialVector_Constant2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		pv.Elements[ii] = symbolic.NewVariable().ToPolynomial()
	}

	// Check each element of the constant vector
	constantOut := pv.Constant()
	for ii := 0; ii < pv.Len(); ii++ {
		// check that constantOut.AtVec(ii) is zer
		if constantOut.AtVec(ii) != 0 {
			t.Errorf(
				"Expected constantOut.AtVec(%v) to be 0; received %v",
				ii,
				constantOut.AtVec(ii),
			)
		}
	}
}

/*
TestPolynomialVector_Constant3
Description:

	Verifies that a small polynomial vector containing a mixture of nonzero constants
	and non-constant polynomials returns the correct constant vector.
*/
func TestPolynomialVector_Constant3(t *testing.T) {
	// Constants
	pv0 := symbolic.PolynomialVector{}
	monom1 := symbolic.Monomial{Coefficient: 3.14}
	monom2 := symbolic.NewVariable().ToMonomial()

	pv0.Elements = append(pv0.Elements, monom1.ToPolynomial())
	pv0.Elements = append(pv0.Elements, monom2.ToPolynomial())

	// Test that the constant vector contains a 3.14 at the first position and a 0 at the second position
	constantOut := pv0.Constant()
	if constantOut.AtVec(0) != 3.14 {
		t.Errorf(
			"Expected constantOut.AtVec(0) to be 3.14; received %v",
			constantOut.AtVec(0),
		)
	}
	if constantOut.AtVec(1) != 0 {
		t.Errorf(
			"Expected constantOut.AtVec(1) to be 0; received %v",
			constantOut.AtVec(1),
		)
	}

}

/*
TestPolynomialVector_LinearCoeff1
Description:

	This test verifies that a panic is thrown if the LinearCoeff method is called on a polynomial vector
	that is not properly initialized.
*/
func TestPolynomialVector_LinearCoeff1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected LinearCoeff to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected LinearCoeff to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected LinearCoeff to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.LinearCoeff()
}

/*
TestPolynomialVector_LinearCoeff2
Description:

	This test verifies that the LinearCoeff method panics when a polynomial of all
	constants is provided to the method.
*/
func TestPolynomialVector_LinearCoeff2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		kII := symbolic.K(float64(ii))
		pv.Elements[ii] = kII.ToMonomial().ToPolynomial()
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected LinearCoeff to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected LinearCoeff to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != (smErrors.CanNotGetLinearCoeffOfConstantError{pv}).Error() {
			t.Errorf(
				"Expected LinearCoeff to panic with error 'polynomial vector has no linear coefficients'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.LinearCoeff()
}

/*
TestPolynomialVector_LinearCoeff3
Description:

	This test verifies that the LinearCoeff method returns
	a matrix of zeros when it contains terms that are all of
	degree 2 or higher.
*/
func TestPolynomialVector_LinearCoeff3(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		vII := symbolic.NewVariable()
		pv.Elements[ii] = symbolic.Monomial{
			VariableFactors: []symbolic.Variable{vII},
			Degrees:         []int{2},
		}.ToPolynomial()
	}

	// Test
	linearCoeff := pv.LinearCoeff()
	nr, nc := linearCoeff.Dims()
	for ii := 0; ii < nr; ii++ {
		for jj := 0; jj < nc; jj++ {
			if linearCoeff.At(ii, jj) != 0 {
				t.Errorf(
					"Expected linearCoeff.At(%v, %v) to be 0; received %v",
					ii,
					jj,
					linearCoeff.At(ii, jj),
				)
			}
		}
	}
}

/*
TestPolynomialVector_LinearCoeff4
Description:

	This test verifies that the LinearCoeff method returns
	an identity matrix when each polynomial in the vector
	contains a linear term containing just that variable.
*/
func TestPolynomialVector_LinearCoeff4(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{
		Elements: make([]symbolic.Polynomial, 20),
	}

	for ii := 0; ii < 20; ii++ {
		vII := symbolic.NewVariable()
		pv.Elements[ii] = symbolic.Monomial{
			Coefficient:     float64(ii),
			VariableFactors: []symbolic.Variable{vII},
			Degrees:         []int{1},
		}.ToPolynomial()
	}

	// Test
	linearCoeff := pv.LinearCoeff()
	nr, nc := linearCoeff.Dims()
	for ii := 0; ii < nr; ii++ {
		for jj := 0; jj < nc; jj++ {
			if ii == jj {
				if linearCoeff.At(ii, jj) != float64(ii) {
					t.Errorf(
						"Expected linearCoeff.At(%v, %v) to be 1; received %v",
						ii,
						jj,
						linearCoeff.At(ii, jj),
					)
				}
			} else {
				if linearCoeff.At(ii, jj) != 0 {
					t.Errorf(
						"Expected linearCoeff.At(%v, %v) to be 0; received %v",
						ii,
						jj,
						linearCoeff.At(ii, jj),
					)
				}
			}
		}
	}
}

/*
TestPolynomialVector_Plus1
Description:

	This test verifies that the plus method throws a panic
	if it is called with a receiver PolynomialVector that isn't
	propertly initialized.
*/
func TestPolynomialVector_Plus1(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Plus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Plus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Plus to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Plus(pv2)
}

/*
TestPolynomialVector_Plus2
Description:

	This test verifies that the Plus method throws an error if
	a correctly initialized PolynomialVector object is added to
	a second expression that is not properly defined.
*/
func TestPolynomialVector_Plus2(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1.Elements = append(pv1.Elements, symbolic.NewVariable().ToPolynomial())
	}

	pv2 := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Plus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Plus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(rAsE.Error(), "polynomial vector has no polynomials") {
			t.Errorf(
				"Expected Plus to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Plus(pv2)
}

/*
TestPolynomialVector_Plus3
Description:

	Tests that a polynomial vector, created by a large number of
	variables, remains a polynomial vector after summing with a
	constant of type float64.
	Then, we test that all polynomials in the vector get an added
	term (for th econstant).
*/
func TestPolynomialVector_Plus3(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1.Elements = append(pv1.Elements, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pv2 := pv1.Plus(3.14).(symbolic.PolynomialVector)
	for _, polynomial := range pv2.Elements {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}
