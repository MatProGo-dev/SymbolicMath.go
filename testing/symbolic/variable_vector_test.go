package symbolic_test

/*
scalar_expression_test.go
Description:
	Tests for the functions mentioned in the variable_vector.go file.
*/

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
TestVariableVector_Len1
Description:

	Verifiest that a variable vector with 111 elements has the proper length.
*/
func TestVariableVector_Len1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	if vv.Len() != N {
		t.Errorf(
			"Expected vv.Len() to be %v; received %v",
			N,
			vv.Len(),
		)
	}
}

/*
TestVariableVector_AtVec1
Description:

	This test verifies that the AtVec function returns a Variable object.
*/
func TestVariableVector_AtVec1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	for ii := 0; ii < N; ii++ {
		if _, ok := vv.AtVec(ii).(symbolic.Variable); !ok {
			t.Errorf(
				"Expected vv.AtVec(%v) to be a Variable object; received %T",
				ii,
				vv.AtVec(ii),
			)
		}
	}
}

/*
TestVariableVector_Variables1
Description:

	Verifies that the variables function returns a slice of unique variables
	that has length equal to the original vector's length.
*/
func TestVariableVector_Variables1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	vars := vv.Variables()
	if len(vars) != N {
		t.Errorf(
			"Expected len(vars) to be %v; received %v",
			N,
			len(vars),
		)
	}

	// Check that all variables are unique
	for ii := 0; ii < N; ii++ {
		for jj := ii + 1; jj < N; jj++ {
			if vars[ii].ID == vars[jj].ID {
				t.Errorf(
					"Expected vars[%v].ID to be unique; received vars[%v].ID= %v",
					ii,
					jj,
					vars[ii].ID,
				)
			}
		}
	}
}

/*
TestVariableVector_Constant1
Description:

	Verifies that the VariableVector's constant always returns a vector of all zeros.
*/
func TestVariableVector_Constant1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	for ii := 0; ii < N; ii++ {
		if const0 := vv.Constant(); const0.AtVec(ii) != 0 {
			t.Errorf(
				"Expected vv.Constant().AtVec(%v) to be 0; received %v",
				ii,
				const0.AtVec(ii),
			)
		}
	}
}

/*
TestVariableVector_LinearCoeff1
Description:

	Verifies that the LinearCoeff method returns an identity matrix for a long
	variable vector.
*/
func TestVariableVector_LinearCoeff1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	for ii := 0; ii < N; ii++ {
		for jj := 0; jj < N; jj++ {
			if ii == jj {
				if L := vv.LinearCoeff(); L.At(ii, jj) != 1 {
					t.Errorf(
						"Expected vv.LinearCoeff().At(%v,%v) to be 1; received %v",
						ii, jj,
						L.At(ii, jj),
					)
				}
				continue
			} else {
				if L := vv.LinearCoeff(); L.At(ii, jj) != 0 {
					t.Errorf(
						"Expected vv.LinearCoeff().At(%v,%v) to be 0; received %v",
						ii, jj,
						L.At(ii, jj),
					)
				}
			}
		}
	}
}

/*
TestVariableVector_Plus1
Description:

	Tests that the Plus method correctly panics when a vector of length 11 to
	a vector of length 111.
*/
func TestVariableVector_Plus1(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	vv2 := symbolic.NewVariableVector(N - 100)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != (smErrors.DimensionError{
			Operation: "Plus",
			Arg1:      vv1,
			Arg2:      vv2,
		}).Error() {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with a DimensionError; received %v",
				r,
			)
		}
	}()

	vv1.Plus(vv2)
}

/*
TestVariableVector_Plus2
Description:

	Verifies that the Plus method will throw an error if an improperly defined variable
	vector was given as the receiver.
*/
func TestVariableVector_Plus2(t *testing.T) {
	// Constants
	N := 111
	vv2 := symbolic.NewVariableVector(N)
	var vv1 symbolic.VariableVector
	for ii := 0; ii < N; ii++ {
		if ii != 100 {
			vv1 = append(vv1, symbolic.NewVariable())
		} else {
			vv1 = append(vv1, symbolic.Variable{})
		}
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 100 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Plus(vv2)
}

/*
TestVariableVector_Plus3
Description:

	This test verifies that a panic is thrown when the rightIn input to the Plus method
	(i.e., not the receiver) is not properly defined.
*/
func TestVariableVector_Plus3(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	var vv2 symbolic.VariableVector
	for ii := 0; ii < N; ii++ {
		if ii != 100 {
			vv2 = append(vv2, symbolic.NewVariable())
		} else {
			vv2 = append(vv2, symbolic.Variable{})
		}
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 100 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Plus(vv2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Plus(vv2)
}

/*
TestVariableVector_Plus4
Description:

	This test verifies that the Plus() method returns a polynomial vector object
	when rightIn is a KVector object.
*/
func TestVariableVector_Plus4(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Test
	r := vv1.Plus(kv2)
	if _, ok := r.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vv1.Plus(kv2) to return a PolynomialVector object; received %T",
			r,
		)
	}
}

/*
TestVariableVector_Plus5
Description:

	This test verifies that the Plus() method returns a polynomial vector object
	when rightIn is a mat.VecDense object.
*/
func TestVariableVector_Plus5(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := symbolic.OnesVector(N)

	// Test
	r := vv1.Plus(kv2)
	if _, ok := r.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vv1.Plus(kv2) to return a PolynomialVector object; received %T",
			r,
		)
	}
}

/*
TestVariableVector_Plus6
Description:

	This test verifies that the sum of a variable vector
	and a valid *mat.VecDense object is a PolynomialVector object.
*/
func TestVariableVector_Plus6(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := symbolic.OnesVector(N)

	// Test
	r := vv1.Plus(kv2)
	if _, ok := r.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vv1.Plus(kv2) to return a PolynomialVector object; received %T",
			r,
		)
	}
}

/*
TestVariableVector_Plus7
Description:

	This test verifies that the Plus() method
	panics when a well-defined VariableVector is summed with
	an invalid object (a string).
*/
func TestVariableVector_Plus7(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	s2 := "Hello, World!"

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Plus(s2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Plus(s2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Plus",
			Input:        s2,
		}
		if !strings.Contains(
			rAsE.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"Expected vv1.Plus(s2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Plus(s2)
}

/*
TestVariableVector_Multiply1
Description:

	This test verifies that the Multiply() method panics when the receiver is
	not properly defined.
*/
func TestVariableVector_Multiply1(t *testing.T) {
	// Constants
	N := 111
	var vv1 symbolic.VariableVector
	for ii := 0; ii < N; ii++ {
		if ii != 100 {
			vv1 = append(vv1, symbolic.NewVariable())
		} else {
			vv1 = append(vv1, symbolic.Variable{})
		}
	}
	k2 := symbolic.K(3.14)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Multiply(k2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Multiply(k2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 100 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Multiply(k2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Multiply(k2)
}

/*
TestVariableVector_Multiply2
Description:

	This test verifies that the Multiply() method panics when the rightIn is
	not properly defined.
*/
func TestVariableVector_Multiply2(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	var vv2 symbolic.VariableVector
	vv2 = append(vv2, symbolic.Variable{})

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 0 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Multiply(vv2)
}

/*
TestVariableVector_Multiply3
Description:

	This test verifies that the Multiply method panics when a vector of improper length
	is provided in the multiplication. A DimensionError should be given in the panic.
*/
func TestVariableVector_Multiply3(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	vv2 := symbolic.NewVariableVector(N - 100)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if rAsE.Error() != (smErrors.DimensionError{
			Operation: "Multiply",
			Arg1:      vv1,
			Arg2:      vv2,
		}).Error() {
			t.Errorf(
				"Expected vv1.Multiply(vv2) to panic with a DimensionError; received %v",
				r,
			)
		}
	}()

	vv1.Multiply(vv2)
}

/*
TestVariableVector_Multiply4
Description:

	This test verifies that the Multiply method, after multiplying
	a vector of variables with a constant, returns a vector of
	vector of monomials.
*/
func TestVariableVector_Multiply4(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	k2 := symbolic.K(3.14)

	// Test
	r := vv1.Multiply(k2)
	if _, ok := r.(symbolic.MonomialVector); !ok {
		t.Errorf(
			"Expected vv1.Multiply(k2) to return a PolynomialVector object; received %T",
			r,
		)
		return
	}

	// Check that all of the coefficients are equal to k2
	rAsMV, _ := r.(symbolic.MonomialVector)
	for ii := 0; ii < N; ii++ {
		if rAsMV.AtVec(ii).(symbolic.Monomial).Coefficient != float64(k2) {
			t.Errorf(
				"Expected r.AtVec(%v) to be k2; received %v",
				ii,
				rAsMV.AtVec(ii),
			)
		}
	}
}

/*
TestVariableVector_Multiply5
Description:

	This test verifies that the Multiply method, after multiplying
	a vector of variables with an invalid object (string), panics.
*/
func TestVariableVector_Multiply5(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	s2 := "Hello, World!"

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Multiply(s2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Multiply(s2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Multiply",
			Input:        s2,
		}
		if !strings.Contains(
			rAsE.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"Expected vv1.Multiply(s2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Multiply(s2)
}

/*
TestVariableVector_Multiply6
Description:

	This test verifies that the Multiply method properly
	multiplies a vector of variables with a polynomial.
*/
func TestVariableVector_Multiply6(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	p2 := symbolic.NewVariable().ToPolynomial()

	// Test
	r := vv1.Multiply(p2)
	if _, ok := r.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vv1.Multiply(pv2) to return a PolynomialVector object; received %T",
			r,
		)
	}

	// Check that all of the coefficients are the same as they were in p2
	rAsPV, _ := r.(symbolic.PolynomialVector)
	for ii := 0; ii < N; ii++ {
		for jj, monomial := range rAsPV.AtVec(ii).(symbolic.Polynomial).Monomials {
			if monomial.Coefficient != p2.Monomials[jj].Coefficient {
				t.Errorf(
					"Expected %v^th monomial in product's %v^th element to have coefficient %v; received %v",
					jj,
					ii,
					p2.Monomials[jj].Coefficient,
					monomial.Coefficient,
				)
			}
		}
	}
}

/*
TestVariableVector_Multiply7
Description:

	This test verifies that the Multiply method properly
	multiplies a vector of variables with a polynomial vector.
	(For this to be valid the polynomial vector must be a scalar)
*/
func TestVariableVector_Multiply7(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	pv2 := symbolic.NewVariableVector(1).ToPolynomialVector()

	// Test
	r := vv1.Multiply(pv2)
	if _, ok := r.(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected vv1.Multiply(pv2) to return a PolynomialVector object; received %T",
			r,
		)
	}

	// Check that all of the coefficients are the same as they were in pv2
	rAsPV, _ := r.(symbolic.PolynomialVector)
	for ii := 0; ii < N; ii++ {
		for jj, monomial := range rAsPV.AtVec(ii).(symbolic.Polynomial).Monomials {
			if monomial.Coefficient != pv2.AtVec(0).(symbolic.Polynomial).Monomials[jj].Coefficient {
				t.Errorf(
					"Expected %v^th monomial in product's %v^th element to have coefficient %v; received %v",
					jj,
					ii,
					pv2.AtVec(0).(symbolic.Polynomial).Monomials[jj].Coefficient,
					monomial.Coefficient,
				)
			}
		}
	}
}

/*
TestVariableVector_Multiply8
Description:

	This test verifies that the multiplication of a variable vector
	with a square matrix and another variable vector produces a
	scalar expression.
*/
func TestVariableVector_Multiply8(t *testing.T) {
	// Constants
	N := 1
	vv1 := symbolic.NewVariableVector(N)
	m2 := getKMatrix.From(symbolic.Identity(N))

	L4 := getKVector.From(symbolic.OnesVector(N))

	// Test
	r := vv1.Transpose().Multiply(m2).Multiply(vv1).Plus(
		L4.Transpose().Multiply(vv1),
	).Plus(3.14)
	if _, ok := r.(symbolic.ScalarExpression); !ok {
		t.Errorf(
			"Expected vv1.Multiply(m2, vv3) to return a scalar expression; received %T",
			r,
		)
	}

	// Check that the output is a scalar polynomial
	if _, ok := r.(symbolic.Polynomial); !ok {
		t.Errorf(
			"Expected vv1.Multiply(m2, vv3) to return a scalar polynomial; received %T",
			r,
		)
	}
}

/*
TestVariableVector_Comparison1
Description:

	This test verifies that the Comparison method panics when the receiver is
	not properly defined.
*/
func TestVariableVector_Comparison1(t *testing.T) {
	// Constants
	N := 111
	var vv1 symbolic.VariableVector
	for ii := 0; ii < N; ii++ {
		if ii != 100 {
			vv1 = append(vv1, symbolic.NewVariable())
		} else {
			vv1 = append(vv1, symbolic.Variable{})
		}
	}
	k2 := symbolic.K(3.14)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Comparison(k2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Comparison(k2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 100 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Comparison(k2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Comparison(k2, symbolic.SenseLessThanEqual)
}

/*
TestVariableVector_Comparison2
Description:

	This test verifies that the Comparison() method panics when the rightIn is
	not properly defined.
*/
func TestVariableVector_Comparison2(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	var vv2 symbolic.VariableVector
	vv2 = append(vv2, symbolic.Variable{})

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 0 has an issue:",
		) {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Comparison(vv2, symbolic.SenseLessThanEqual)
}

/*
TestVariableVector_Comparison3
Description:

	This test verifies that the Comparison method panics when a vector of improper length
	is provided in the comparison. A DimensionError should be given in the panic.
*/
func TestVariableVector_Comparison3(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	vv2 := symbolic.NewVariableVector(N - 100)

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		var tempSense symbolic.ConstrSense = symbolic.SenseLessThanEqual
		if rAsE.Error() != (smErrors.DimensionError{
			Operation: "Comparison (" + tempSense.String() + ")",
			Arg1:      vv1,
			Arg2:      vv2,
		}).Error() {
			t.Errorf(
				"Expected vv1.Comparison(vv2) to panic with a DimensionError; received %v",
				r,
			)
		}
	}()

	vv1.Comparison(vv2, symbolic.SenseLessThanEqual)
}

/*
TestVariableVector_Comparison4
Description:

	This test verifies that the Comparison method panics when the rightIn is
	an invalid object (string).
*/
func TestVariableVector_Comparison4(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	s2 := "Hello, World!"

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv1.Comparison(s2) to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv1.Comparison(s2) to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		expectedError := smErrors.UnsupportedInputError{
			FunctionName: "VariableVector.Comparison",
			Input:        s2,
		}
		if !strings.Contains(
			rAsE.Error(),
			expectedError.Error(),
		) {
			t.Errorf(
				"Expected vv1.Comparison(s2) to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv1.Comparison(s2, symbolic.SenseLessThanEqual)
}

/*
TestVariableVector_LessEq1
Description:

	This test verifies that the LessEq method returns a constraint object.
	We verify that the constraint:
	- Has sense SenseLessThanEq
	- Has the correct left and right hand sides (of type VariableVector and KVector,
		respectively)
*/
func TestVariableVector_LessEq1(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := symbolic.OnesVector(N)

	// Test
	r := vv1.LessEq(kv2)
	if r.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf(
			"Expected r.Sense() to be SenseLessThanEqual; received %v",
			r.ConstrSense(),
		)
	}

	// Check that the type of the Left() is VariableVector
	if _, ok := r.Left().(symbolic.VariableVector); !ok {
		t.Errorf(
			"Expected r.Left() to be a VariableVector; received %T",
			r.Left(),
		)

	}

	// Check that the right hand side is a KVector
	if _, ok := r.Right().(symbolic.KVector); !ok {
		t.Errorf(
			"Expected r.Right() to be a KVector; received %T",
			r.Right(),
		)
	}
}

/*
TestVariableVector_GreaterEq1
Description:

	This test verifies that the GreaterEq method returns a constraint object.
	We verify that the constraint:
	- Has sense SenseGreaterThanEq
	- Has the correct left and right hand sides (of type VariableVector and KVector),
*/
func TestVariableVector_GreaterEq1(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := symbolic.OnesVector(N)

	// Test
	r := vv1.GreaterEq(kv2)
	if r.ConstrSense() != symbolic.SenseGreaterThanEqual {
		t.Errorf(
			"Expected r.Sense() to be SenseGreaterThanEqual; received %v",
			r.ConstrSense(),
		)
	}

	// Check that the type of the Left() is VariableVector
	if _, ok := r.Left().(symbolic.VariableVector); !ok {
		t.Errorf(
			"Expected r.Left() to be a VariableVector; received %T",
			r.Left(),
		)

	}

	// Check that the right hand side is a KVector
	if _, ok := r.Right().(symbolic.KVector); !ok {
		t.Errorf(
			"Expected r.Right() to be a KVector; received %T",
			r.Right(),
		)
	}
}

/*
TestVariableVector_Eq1
Description:

	This test verifies that the Eq method returns a constraint object.
	We verify that the constraint:
	- Has sense SenseEqual
	- Has the correct left and right hand sides (of type VariableVector and PolynomialVector),
*/
func TestVariableVector_Eq1(t *testing.T) {
	// Constants
	N := 111
	vv1 := symbolic.NewVariableVector(N)
	kv2 := getKVector.From(symbolic.OnesVector(N))
	pv3 := kv2.ToPolynomialVector()

	// Test
	r := vv1.Eq(pv3)
	if r.ConstrSense() != symbolic.SenseEqual {
		t.Errorf(
			"Expected r.Sense() to be SenseEqual; received %v",
			r.ConstrSense(),
		)
	}

	// Check that the type of the Left() is VariableVector
	if _, ok := r.Left().(symbolic.VariableVector); !ok {
		t.Errorf(
			"Expected r.Left() to be a VariableVector; received %T",
			r.Left(),
		)

	}

	// Check that the right hand side is a PolynomialVector
	if _, ok := r.Right().(symbolic.PolynomialVector); !ok {
		t.Errorf(
			"Expected r.Right() to be a PolynomialVector; received %T",
			r.Right(),
		)
	}
}

/*
TestVariableVector_Copy1
Description:

	Verifies that the Copy method returns a new VariableVector object.
*/
func TestVariableVector_Copy1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	vv2 := vv.Copy()

	// Replace v2[10] with a new variable
	vv2[10] = symbolic.NewVariable()

	// Verify that vv[10] != vv2[10]
	if vv[10].ID == vv2[10].ID {
		t.Errorf(
			"Expected vv[10].ID to not be equal to vv2[10].ID; received %v",
			vv[10].ID,
		)
	}
}

/*
TestVariableVector_Transpose1
Description:

	Verifies that the Transpose method returns a VariableMatrix
	object with the appropriate size.
*/
func TestVariableVector_Transpose1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	vm := vv.Transpose()
	if vm.Dims()[0] != 1 || vm.Dims()[1] != N {
		t.Errorf(
			"Expected vm.Rows() to be 1 and vm.Cols() to be %v; received (%v, %v)",
			N,
			vm.Dims()[0],
			vm.Dims()[1],
		)
	}
}

/*
TestVariableVector_Transpose2
Description:

	Verifies that the Transpose method panics if
	the variable vector used to call it is not
	well-defined.
*/
func TestVariableVector_Transpose2(t *testing.T) {
	// Constants
	N := 111
	var vv symbolic.VariableVector
	for ii := 0; ii < N; ii++ {
		if ii != 100 {
			vv = append(vv, symbolic.NewVariable())
		} else {
			vv = append(vv, symbolic.Variable{})
		}
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected vv.Transpose() to panic; received no panic",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected vv.Transpose() to panic with an error; received %v",
				r,
			)
		}

		// Check that the error message is correct
		if !strings.Contains(
			rAsE.Error(),
			"element 100 has an issue:",
		) {
			t.Errorf(
				"Expected vv.Transpose() to panic with a specific error; instead received %v",
				r,
			)
		}
	}()

	vv.Transpose()
}

/*
TestVariableVector_String1
Description:

	Verifies that the String method returns a string which contains
	each of the variables in the vector.
*/
func TestVariableVector_String1(t *testing.T) {
	// Constants
	N := 111
	vv := symbolic.NewVariableVector(N)

	// Test
	str := vv.String()
	for ii := 0; ii < N; ii++ {
		if !strings.Contains(str, vv[ii].String()) {
			t.Errorf(
				"Expected vv.String() to contain \"%v\"; received \"%v\"",
				vv[ii].String(),
				str,
			)
		}
	}
}
