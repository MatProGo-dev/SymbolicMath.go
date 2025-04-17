package symbolic_test

/*
polynomial_vector_test.go
Description:
	Tests the methods defined in the polynomial_vector.go file.
*/

import (
	"strings"
	"testing"

	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)
	for ii := 0; ii < 20; ii++ {
		if ii != 11 {
			pv[ii] = symbolic.NewVariable().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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
TestPolynomialVector_AtVec3
Description:

	Verifies that the AtVec method panics when called on a polynomial vector
	that has not been properly initialized.
*/
func TestPolynomialVector_AtVec3(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

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
		if rAsE.Error() != pv.Check().Error() {
			t.Errorf(
				"Expected AtVec to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.AtVec(0)

}

/*
TestPolynomialVector_Variables1
Description:

	Verifies that the number of variables found in a polynomial vector containing all constant
	elements is zero.
*/
func TestPolynomialVector_Variables1(t *testing.T) {
	// Constants
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		tempK := symbolic.K(1)
		pv[ii] = tempK.ToMonomial().ToPolynomial()
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

	var pv symbolic.PolynomialVector = []symbolic.Polynomial{
		k1.ToMonomial().ToPolynomial(),
		symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3}, Exponents: []int{1, 2}}.ToPolynomial(),
		symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3, v4}, Exponents: []int{3, 5, 11}}.ToPolynomial(),
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

	var pv symbolic.PolynomialVector = []symbolic.Polynomial{
		k1.ToMonomial().ToPolynomial(),
		symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v3}, Exponents: []int{1, 2}}.ToPolynomial(),
		symbolic.Monomial{VariableFactors: []symbolic.Variable{v2, v4}, Exponents: []int{3, 11}}.ToPolynomial(),
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		pv[ii] = symbolic.NewVariable().ToPolynomial()
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

	pv0 = append(pv0, monom1.ToPolynomial())
	pv0 = append(pv0, monom2.ToPolynomial())

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

	This test verifies that the LinearCoeff method DOES NOT PANIC when a polynomial of all
	constants is provided to the method.
*/
func TestPolynomialVector_LinearCoeff2(t *testing.T) {
	// Constants
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		kII := symbolic.K(float64(ii))
		pv[ii] = kII.ToMonomial().ToPolynomial()
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		vII := symbolic.NewVariable()
		pv[ii] = symbolic.Monomial{
			VariableFactors: []symbolic.Variable{vII},
			Exponents:       []int{2},
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
	var pv symbolic.PolynomialVector = make([]symbolic.Polynomial, 20)

	for ii := 0; ii < 20; ii++ {
		vII := symbolic.NewVariable()
		pv[ii] = symbolic.Monomial{
			Coefficient:     float64(ii),
			VariableFactors: []symbolic.Variable{vII},
			Exponents:       []int{1},
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
TestPolynomialVector_LinearCoeff5
Description:

	This test verifies that the LinearCoeff method returns
	a length 2 slice when we provide a varSlice of length 2.
*/
func TestPolynomialVector_LinearCoeff5(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	pv := symbolic.PolynomialVector{symbolic.NewVariable().ToPolynomial()}

	v2 := symbolic.NewVariable()

	vSlice3 := []symbolic.Variable{v1, v2}

	// Test
	linearCoeff := pv.LinearCoeff(vSlice3)
	nx, ny := linearCoeff.Dims()
	if ny != len(vSlice3) {
		t.Errorf(
			"Expected linearCoeff to have 2 columns; received %v",
			ny,
		)
	}

	if nx != pv.Len() {
		t.Errorf(
			"Expected linearCoeff to have %v rows; received %v",
			pv.Len(),
			nx,
		)
	}
}

/*
TestPolynomialVector_LinearCoeff6
Description:

	This test verifies that the LinearCoeff method panics
	if multiple varSlices are provided to the LinearCoeff.
*/
func TestPolynomialVector_LinearCoeff6(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	pv := symbolic.PolynomialVector{symbolic.NewVariable().ToPolynomial()}

	v2 := symbolic.NewVariable()

	vSlice3 := []symbolic.Variable{v1, v2}

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
		if !strings.Contains(rAsE.Error(), "Too many inputs provided to LinearCoeff()") {
			t.Errorf(
				"Expected LinearCoeff to panic with error 'only one varSlice'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.LinearCoeff(vSlice3, vSlice3)
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
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
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
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pv2 := pv1.Plus(3.14).(symbolic.PolynomialVector)
	for _, polynomial := range pv2 {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Plus4
Description:

	Tests that a polynomial vector added to a constant vector
	produces a polynomial vector.
*/
func TestPolynomialVector_Plus4(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pv2 := pv1.Plus(3.14).(symbolic.PolynomialVector)
	for _, polynomial := range pv2 {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Plus5
Description:

	Tests that a polynomial vector added to a polynomial vector
	produces a polynomial vector.
*/
func TestPolynomialVector_Plus5(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pv2 := pv1.Plus(pv1).(symbolic.PolynomialVector)
	for _, polynomial := range pv2 {
		if len(polynomial.Monomials) != 1 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Plus6
Description:

	Tests that a polynomial vector added to a polynomial
	results in a polynomial vector object.
*/
func TestPolynomialVector_Plus6(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pv2 := pv1.Plus(symbolic.NewVariable().ToPolynomial()).(symbolic.PolynomialVector)
	for _, polynomial := range pv2 {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Plus7
Description:

	Tests that a polynomial vector added to a polynomial vector of incompatible size
	results in a panic.
*/
func TestPolynomialVector_Plus7(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 2; ii++ {
		pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())
	}

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
		expectedError := smErrors.CheckDimensionsInAddition(pv1, pv2)
		if !strings.Contains(rAsE.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected Plus to panic with error 'incompatible sizes'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Plus(pv2)
}

/*
TestPolynomialVector_Plus8
Description:

	Tests that the Plus() method panics when the second input to it (not the receiver) is not
	an expression (or like an expression). In this case, the second input is a string.
*/
func TestPolynomialVector_Plus8(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}

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
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "PolynomialVector.Plus",
			Input:        "hello",
		}
		if !strings.Contains(rAsE.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected Plus to panic with error 'unsupported type'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Plus("hello")
}

/*
TestPolynomialVector_Minus1
Description:

	This test verifies that the Minus method throws a panic
	if it is called with a receiver PolynomialVector that isn't
	propertly initialized.
*/
func TestPolynomialVector_Minus1(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Minus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Minus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != pv1.Check().Error() {
			t.Errorf(
				"Expected Minus to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Minus(pv2)

}

/*
TestPolynomialVector_Minus2
Description:

	This test verifies that the Minus method panics if a well-defined polynomial
	vector is given as the receiver and a second monomial vector is given as the
	second input. The second input is not well-defined.
*/
func TestPolynomialVector_Minus2(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	mv2 := symbolic.MonomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Minus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Minus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(rAsE.Error(), mv2.Check().Error()) {
			t.Errorf(
				"Expected Minus to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Minus(mv2)
}

/*
TestPolynomialVector_Minus3
Description:

	Verifies that the Minus method panics when the inputs to the method do not match
	in terms of dimensions.
*/
func TestPolynomialVector_Minus3(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 2; ii++ {
		pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Minus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Minus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		expectedError := smErrors.CheckDimensionsInSubtraction(pv1, pv2)
		if !strings.Contains(rAsE.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected Minus to panic with error 'incompatible sizes'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Minus(pv2)
}

/*
TestPolynomialVector_Minus4
Description:

	Tests that the Minus() method returns a valid polynomial vector
	when we compute the difference between a polynomial vector and a mat.VecDense object.
*/
func TestPolynomialVector_Minus4(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	v2 := symbolic.OnesVector(20)

	// Test
	pv3, tf := pv1.Minus(v2).(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"Expected Minus to return a polynomial vector; received %T",
			pv3,
		)
	}
	// Check that each polynomial contains two monomials
	for _, polynomial := range pv3 {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}

	// Check the length
	if len(pv3) != 20 {
		t.Errorf(
			"Expected pv3 to have length 20; received %v",
			len(pv3),
		)
	}
}

/*
TestPolynomialVector_Minus5
Description:

	Tests that the Minus() method returns a valid polynomial vector
	when we compute the difference between a polynomial vector and a
	float64.
*/
func TestPolynomialVector_Minus5(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	f2 := 3.14

	// Test
	pv3, tf := pv1.Minus(f2).(symbolic.PolynomialVector)
	if !tf {
		t.Errorf(
			"Expected Minus to return a polynomial vector; received %T",
			pv3,
		)
	}
	// Check that each polynomial contains two monomials
	for _, polynomial := range pv3 {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}

	// Check the length
	if len(pv3) != 20 {
		t.Errorf(
			"Expected pv3 to have length 20; received %v",
			len(pv3),
		)
	}

}

/*
TestPolynomialVector_Minus6
Description:

	Tests that the Minus() method panics when the second input to it (not the receiver) is
	not recognized as an expression or expression-like.
*/
func TestPolynomialVector_Minus6(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Minus to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Minus to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "PolynomialVector.Minus",
			Input:        "hello",
		}
		if !strings.Contains(rAsE.Error(), expectedError.Error()) {
			t.Errorf(
				"Expected Minus to panic with error 'unsupported type'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Minus("hello")
}

/*
TestPolynomialVector_Multiply1
Description:

	This test verifies that the Multiply() method panics
	when called on an improperly initialized polynomial vector.
*/
func TestPolynomialVector_Multiply1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Multiply to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Multiply to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Multiply to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Multiply(3.14)
}

/*
TestPolynomialVector_Multiply2
Description:

	This test verifies that the Multiply() method panics
	when the second input to it (not the receiver) is improperly
	initialized.
*/
func TestPolynomialVector_Multiply2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}
	pv2 := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Multiply to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Multiply to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(rAsE.Error(), "polynomial vector has no polynomials") {
			t.Errorf(
				"Expected Multiply to panic with error \"polynomial vector has no polynomials\"; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Multiply(pv2)
}

/*
TestPolynomialVector_Multiply3
Description:

	This test verifies that the Multiply() method returns a polynomial
	with the correct coefficients when the second input is a constant.
*/
func TestPolynomialVector_Multiply3(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}
	k2 := symbolic.K(3.14)

	// Test
	pv3 := pv.Multiply(k2).(symbolic.PolynomialVector)
	for _, polynomial := range pv3 {
		for _, monomial := range polynomial.Monomials {
			if monomial.Coefficient != 3.14 {
				t.Errorf(
					"Expected monomial.Coefficient to be 3.14; received %v",
					monomial.Coefficient,
				)
			}
		}
	}
}

/*
TestPolynomialVector_Multiply4
Description:

	This test verifies that the Multiply() method returns a polynomial
	with the correct coefficients when the second input is a float64.
*/
func TestPolynomialVector_Multiply4(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}
	f2 := 3.14

	// Test
	pv3 := pv.Multiply(f2).(symbolic.PolynomialVector)
	for _, polynomial := range pv3 {
		for _, monomial := range polynomial.Monomials {
			if monomial.Coefficient != 3.14 {
				t.Errorf(
					"Expected monomial.Coefficient to be 3.14; received %v",
					monomial.Coefficient,
				)
			}
		}
	}
}

/*
TestPolynomialVector_Multiply5
Description:

	This test verifies that a polynomial vector when multiplied
	by a polynomial results in a polynomial vector object.
*/
func TestPolynomialVector_Multiply5(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}
	p2 := symbolic.NewVariable().ToPolynomial()

	// Test
	pv3 := pv.Multiply(p2).(symbolic.PolynomialVector)
	for _, polynomial := range pv3 {
		if len(polynomial.Monomials) != 1 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Multiply6
Description:

	Tests that a polynomial vector when multiplied by a polynomial
	vector of incompatible size results in a panic.
*/
func TestPolynomialVector_Multiply6(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 2; ii++ {
		pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Multiply to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Multiply to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != (smErrors.MatrixDimensionError{
			Arg1:      pv1,
			Arg2:      pv2,
			Operation: "Multiply",
		}).Error() {
			t.Errorf(
				"Expected Multiply to panic with error \"polynomial vector has no polynomials\"; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Multiply(pv2)
}

/*
TestPolynomialVector_Multiply7
Description:

	Tests that a polynomial vector when multiplied by a polynomial
	vector of compatible size (1 x 1) results in a polynomial vector.
*/
func TestPolynomialVector_Multiply7(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}
	pv2 := symbolic.PolynomialVector{}
	pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())

	// Test
	pv3 := pv1.Multiply(pv2).(symbolic.PolynomialVector)
	for _, polynomial := range pv3 {
		if len(polynomial.Monomials) != 1 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Transpose1
Description:

	Tests that the Transpose method returns a properly transposed
	polynomial vector when called on a properly initialized polynomial.
	The result should be a polynomial matrix of dimension (1x3) for an
	input polynomial of dimension (3x1).
*/
func TestPolynomialVector_Transpose1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 3; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	pvT, ok := pv.Transpose().(symbolic.PolynomialMatrix)
	if !ok {
		t.Errorf(
			"Expected pv.Transpose to return a PolynomialMatrix; received %T",
			pv.Transpose(),
		)
	}

	// Check that the dimensions are correct
	nr, nc := pvT.Dims()[0], pvT.Dims()[1]
	if nr != 1 {
		t.Errorf(
			"Expected nr to be 1; received %v",
			nr,
		)
	}

	if nc != 3 {
		t.Errorf(
			"Expected nc to be 3; received %v",
			nc,
		)
	}

}

/*
TestPolynomialVector_Transpose2
Description:

	Tests that the Transpose method panics when
	called on a polynomial vector that is not properly
	initialized.
*/
func TestPolynomialVector_Transpose2(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Transpose to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Transpose to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Transpose to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv.Transpose()
}

/*
TestPolynomialVector_Dims1
Description:

	This test verifies that a length 20 polynomial vector
	returns a slice []int{20,1} from the Dims() method.
*/
func TestPolynomialVector_Dims1(t *testing.T) {
	// Constants
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	dims := pv.Dims()
	nr, nc := dims[0], dims[1]
	if nr != 20 {
		t.Errorf(
			"Expected nr to be 20; received %v",
			nr,
		)
	}
	if nc != 1 {
		t.Errorf(
			"Expected nc to be 1; received %v",
			nc,
		)
	}
}

/*
TestPolynomialVector_Comparison1
Description:

	This test verifies that the Comparison method throws an error
	when the polynomial vector that calls it is improperly
	defined.
*/
func TestPolynomialVector_Comparison1(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Comparison to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Comparison to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != "polynomial vector has no polynomials" {
			t.Errorf(
				"Expected Comparison to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Comparison(pv2, symbolic.SenseGreaterThanEqual)
}

/*
TestPolynomialVector_Comparison2
Description:

	This test verifies that the Comparison method panics
	when the right hand side argument is improperly defined.
*/
func TestPolynomialVector_Comparison2(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected Comparison to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected Comparison to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(rAsE.Error(), "polynomial vector has no polynomials") {
			t.Errorf(
				"Expected Comparison to panic with error 'polynomial vector has no polynomials'; received '%v'",
				rAsE.Error(),
			)
		}
	}()

	pv1.Comparison(pv2, symbolic.SenseGreaterThanEqual)
}

/*
TestPolynomialVector_Comparison3
Description:

	This test verifies that the comparison function returns
	a proper vector constraint when a float64 variable is
	provided to the function.
*/
func TestPolynomialVector_Comparison3(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
		pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())
	}
	f1 := 3.14

	// Test
	comp := pv1.Comparison(f1, symbolic.SenseGreaterThanEqual)
	vectorComparison1, tf := comp.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf(
			"Expected comp to be of type VectorConstraint; received %T",
			comp,
		)
	}

	// Check that the right hand side of the constraint has the length of 20.
	if vectorComparison1.RightHandSide.Len() != 20 {
		t.Errorf(
			"Expected vectorComparison1.RightHandSide.Len() to be 20; received %v",
			vectorComparison1.RightHandSide.Len(),
		)
	}
}

/*
TestPolynomialVector_Comparison4
Description:

	This test verifies that the comparison function returns
	a proper vector constraint when a polynomial vector is
	provided to the function.
*/
func TestPolynomialVector_Comparison4(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	pv2 := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
		pv2 = append(pv2, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	comp := pv1.Comparison(pv2, symbolic.SenseGreaterThanEqual)
	vectorComparison1, tf := comp.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf(
			"Expected comp to be of type VectorConstraint; received %T",
			comp,
		)
	}

	// Check that the right hand side of the constraint has the length of 20.
	if vectorComparison1.RightHandSide.Len() != 20 {
		t.Errorf(
			"Expected vectorComparison1.RightHandSide.Len() to be 20; received %v",
			vectorComparison1.RightHandSide.Len(),
		)
	}
}

/*
TestPolynomialVector_LessEq1
Description:

	This test verifies that the LessEq method panics when
	called on a polynomial vector that is well-defined is
	compared with a monomial vector that is well-defined,
	but has a different size. (Polynomial vector is (3x1)
	and monomial vector that is (4x1)).
*/
func TestPolynomialVector_LessEq1(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 3; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	mv2 := symbolic.NewVariableVector(4).ToMonomialVector()

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected LessEq to panic; received no panic",
			)
		}
	}()

	pv1.LessEq(mv2)
}

/*
TestPolynomialVector_LessEq2
Description:

	This test verifies that the LessEq method produces
	a well-defined constraint when a well-defined monomial
	vector is compared with a constant vector.
*/
func TestPolynomialVector_LessEq2(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 3; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	mv2 := symbolic.NewVariableVector(3).ToMonomialVector()

	// Test
	comp := pv1.LessEq(mv2)
	_, tf := comp.(symbolic.VectorConstraint)
	if !tf {
		t.Errorf(
			"Expected comp to be of type VectorConstraint; received %T",
			comp,
		)
	}
}

/*
TestPolynomialVector_GreaterEq1
Description:

	This test verifies that the GreaterEq method panics when
	a polynomial vector that is well-defined is compared with
	an object that is not an expression (in this case, a string).
*/
func TestPolynomialVector_GreaterEq1(t *testing.T) {
	// Constants
	pv1 := symbolic.PolynomialVector{}
	for ii := 0; ii < 3; ii++ {
		pv1 = append(pv1, symbolic.NewVariable().ToPolynomial())
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected GreaterEq to panic; received no panic",
			)
		}
	}()

	pv1.GreaterEq("test")
}

/*
TestPolynomialVector_Simplify1
Description:

	This test verifies that the Simplify method does not clear out
	all elements of a polynomial vector.
	(Comes from a specific example encountered during testing).
*/
func TestPolynomialVector_Simplify1(t *testing.T) {
	// Create a polynomial vector
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().Plus(3.14).(symbolic.Polynomial))
	}

	// Try to simplify
	pvOut := pv.Simplify()

	// Check each element of pvOut and verify that it has two monomials.
	for _, polynomial := range pvOut {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}

/*
TestPolynomialVector_Simplify2
Description:

	This test verifies that the Simplify method works when simplification
	of terms is possible.
*/
func TestPolynomialVector_Simplify2(t *testing.T) {
	// Create a polynomial vector
	pv := symbolic.PolynomialVector{}
	for ii := 0; ii < 20; ii++ {
		pv = append(pv, symbolic.NewVariable().Plus(3.14).Plus(2.17).Plus(
			symbolic.Monomial{Coefficient: 1.1},
		).(symbolic.Polynomial))
	}

	// Try to simplify
	pvOut := pv.Simplify()

	// Check each element of pvOut and verify that it has two monomials.
	for _, polynomial := range pvOut {
		if len(polynomial.Monomials) != 2 {
			t.Errorf(
				"Expected polynomial.Monomials to have length 2; received %v",
				len(polynomial.Monomials),
			)
		}
	}
}
