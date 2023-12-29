package symbolic_test

/*
scalar_expression_test.go
Description:
	Tests for the functions mentioned in the variable_vector.go file.
*/

import (
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
			vv1.Elements = append(vv1.Elements, symbolic.NewVariable())
		} else {
			vv1.Elements = append(vv1.Elements, symbolic.Variable{})
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
			vv2.Elements = append(vv2.Elements, symbolic.NewVariable())
		} else {
			vv2.Elements = append(vv2.Elements, symbolic.Variable{})
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
	kv2 := symbolic.KVector(symbolic.OnesVector(N))

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
			vv1.Elements = append(vv1.Elements, symbolic.NewVariable())
		} else {
			vv1.Elements = append(vv1.Elements, symbolic.Variable{})
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
	vv2.Elements = append(vv2.Elements, symbolic.Variable{})

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
			vv1.Elements = append(vv1.Elements, symbolic.NewVariable())
		} else {
			vv1.Elements = append(vv1.Elements, symbolic.Variable{})
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
	vv2.Elements = append(vv2.Elements, symbolic.Variable{})

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
