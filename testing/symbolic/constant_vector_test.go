package symbolic_test

/*
constant_vector_test.go
Description:
	Tests for the functions mentioned in the constant_vector.go file.
*/

import (
	"fmt"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
TestConstantVector_Len1
Description:

	Tests that an empty kvector has length 0.
*/
func TestConstantVector_Len1(t *testing.T) {
	// Constants
	kv := symbolic.KVector{}

	// Test
	if kv.Len() != 0 {
		t.Errorf(
			"Expected kv.Len() to be 0; received %v",
			kv.Len(),
		)
	}
}

/*
TestConstantVector_Len2
Description:

	Tests that an vector with over 100 elements has the proper length.
*/
func TestConstantVector_Len2(t *testing.T) {
	// Constants
	N := 101
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Test
	if kv.Len() != N {
		t.Errorf(
			"Expected kv.Len() to be %v; received %v",
			N,
			kv.Len(),
		)
	}
}

/*
TestConstantVector_Check1
Description:

	This test verifies that Check returns nil for a well-defined constant vector.
*/
func TestConstantVector_Check1(t *testing.T) {
	// Constants
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(3))

	// Test
	if kv.Check() != nil {
		t.Errorf(
			"Expected kv.Check() to return nil; received %v",
			kv.Check(),
		)
	}
}

/*
TestConstantVector_AtVec1
Description:

	Verifies that this function returns the correct value when an in-bounds value
	is given for the vector.
*/
func TestConstantVector_AtVec1(t *testing.T) {
	// Constants
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(3))

	// Test
	if float64(kv.AtVec(0).(symbolic.K)) != 1 {
		t.Errorf(
			"Expected kv.AtVec(0) to be 1; received %v",
			kv.AtVec(0),
		)
	}
}

/*
TestConstantVector_AtVec2
Description:

	Verifies that the AtVec function panics when an out-of-bounds value is given.
*/
func TestConstantVector_AtVec2(t *testing.T) {
	// Constants
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(3))

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected kv.AtVec(3) to panic; received %v",
				kv.AtVec(3),
			)
		}
	}()

	kv.AtVec(3)
}

/*
TestConstantVector_Variables1
Description:

	Verifies that the variables function returns an empty slice for this constant.
*/
func TestConstantVector_Variables1(t *testing.T) {
	// Constants
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(3))

	// Test
	vars := kv.Variables()
	if len(vars) != 0 {
		t.Errorf(
			"Expected len(vars) to be 0; received %v",
			len(vars),
		)
	}
}

/*
TestConstantVector_LinearCoeff1
Description:

	Verifies that the LinearCoeff method returns a matrix of all zeros for a long
	constant matrix vector.
*/
func TestConstantVector_LinearCoeff1(t *testing.T) {
	// Constants
	N := 11
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Test
	for ii := 0; ii < N; ii++ {
		for jj := 0; jj < N; jj++ {
			if L := kv.LinearCoeff(); L.At(ii, jj) != 0 {
				t.Errorf(
					"Expected kv.LinearCoeff().At(%v,%v) to be 0; received %v",
					ii, jj,
					L.At(ii, jj),
				)
			}
		}
	}
}

/*
TestConstantVector_Constant1
Description:

	This function verifies that the constant method produces a mat.VecDense containing
	the same elements used to initialize the KVector.
*/
func TestConstantVector_Constant1(t *testing.T) {
	// Constants
	N := 11
	kv := symbolic.VecDenseToKVector(symbolic.OnesVector(N))

	// Test
	for ii := 0; ii < N; ii++ {
		if L := kv.Constant(); L.AtVec(ii) != 1 {
			t.Errorf(
				"Expected kv.Constant().AtVec(%v) to be 1; received %v",
				ii,
				L.AtVec(ii),
			)
		}
	}
}

/*
TestConstantVector_Plus1
Description:

	Verifies that the Plus method panics when two vectors are added together with
	different lengths.
*/
func TestConstantVector_Plus1(t *testing.T) {
	// Constants
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(3))
	kv2 := symbolic.VecDenseToKVector(symbolic.OnesVector(4))

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected kv1.Plus(kv2) to panic; received %v",
				kv1.Plus(kv2),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if rAsError.Error() != (smErrors.DimensionError{
			Operation: "Plus",
			Arg1:      kv1,
			Arg2:      kv2,
		}).Error() {
			t.Errorf(
				"Expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	kv1.Plus(kv2)
}

/*
TestConstantVector_Plus2
Description:

	This test verifies that the Plus() method panics if the
	rightIn input to the Plus() method is improperly defined.
*/
func TestConstantVector_Plus2(t *testing.T) {
	// Constants
	N := 3
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	vv2 := symbolic.NewVariableVector(N)
	vv2[0] = symbolic.Variable{
		ID:    1001,
		Lower: 0.0,
		Upper: -1.0,
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected kv1.Plus(kv2) to panic; received none",
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		if !strings.Contains(rAsError.Error(), (fmt.Errorf(
			"lower bound (%v) of variable must be less than upper bound (%v).",
			0.0, -1.0,
		)).Error()) {
			t.Errorf(
				"Expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	kv1.Plus(vv2)
}

/*
TestConstantVector_Plus3
Description:

	This test verifies that the Plus() method correctly adds
	a KVector with a float64 (3.14) by checking all elements
	of the vector.
*/
func TestConstantVector_Plus3(t *testing.T) {
	// Constants
	N := 3
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	f2 := 3.14

	// Test
	kv3 := kv1.Plus(f2)
	for ii := 0; ii < N; ii++ {
		if float64(kv3.(symbolic.KVector).AtVec(ii).(symbolic.K)) != float64(kv1.AtVec(ii).(symbolic.K))+f2 {
			t.Errorf(
				"Expected kv3.AtVec(%v) to be %v; received %v",
				ii,
				kv3.(symbolic.KVector).AtVec(ii),
				float64(kv1.AtVec(ii).(symbolic.K))+f2,
			)
		}
	}
}

/*
TestConstantVector_Plus4
Description:

	This test verifies that the Plus() method correctly adds
	a KVector with a K (3.14) by checking all elements
	of the vector.
*/
func TestConstantVector_Plus4(t *testing.T) {
	// Constants
	N := 3
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	k2 := symbolic.K(3.14)

	// Test
	kv3 := kv1.Plus(k2)
	for ii := 0; ii < N; ii++ {
		if float64(kv3.(symbolic.KVector).AtVec(ii).(symbolic.K)) != float64(kv1.AtVec(ii).(symbolic.K))+float64(k2) {
			t.Errorf(
				"Expected kv3.AtVec(%v) to be %v; received %v",
				ii,
				kv3.(symbolic.KVector).AtVec(ii),
				float64(kv1.AtVec(ii).(symbolic.K))+float64(k2),
			)
		}
	}
}

/*
TestConstantVector_Plus5
Description:

	Tests that the Plus() method correctly computes the sum of a
	KVector and a mat.VecDense.
*/
func TestConstantVector_Plus5(t *testing.T) {
	// Constants
	N := 3
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	vd2 := mat.NewVecDense(N, []float64{3.14, 2.71, 1.41})

	// Test
	kv3 := kv1.Plus(vd2)
	for ii := 0; ii < N; ii++ {
		if float64(kv3.(symbolic.KVector).AtVec(ii).(symbolic.K)) != float64(kv1.AtVec(ii).(symbolic.K))+vd2.AtVec(ii) {
			t.Errorf(
				"Expected kv3.AtVec(%v) to be %v; received %v",
				ii,
				kv3.(symbolic.KVector).AtVec(ii),
				float64(kv1.AtVec(ii).(symbolic.K))+vd2.AtVec(ii),
			)
		}
	}
}

/*
TestConstantVector_Plus6
Description:

	Tests that the Plus() method correctly computes the sum of a
	KVector and a KVector.
*/
func TestConstantVector_Plus6(t *testing.T) {
	// Constants
	N := 3
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(N))
	vd2 := mat.NewVecDense(N, []float64{3.14, 2.71, 1.41})
	kv2 := symbolic.VecDenseToKVector(*vd2)

	// Test
	kv3 := kv1.Plus(vd2)
	for ii := 0; ii < N; ii++ {
		if float64(kv3.(symbolic.KVector).AtVec(ii).(symbolic.K)) != float64(kv1.AtVec(ii).(symbolic.K))+float64(kv2.AtVec(ii).(symbolic.K)) {
			t.Errorf(
				"Expected kv3.AtVec(%v) to be %v; received %v",
				ii,
				kv3.(symbolic.KVector).AtVec(ii),
				float64(kv1.AtVec(ii).(symbolic.K))+float64(kv2.AtVec(ii).(symbolic.K)),
			)
		}
	}
}

/*
TestConstantVector_Plus7
Description:

	Tests that the Plus() method correctly computes the sum of a
	KVector and a MonomialVector.
*/
func TestConstantVector_Plus7(t *testing.T) {
	// Constants
	N := 3
	kv1 := getKVector.From(symbolic.OnesVector(N))
	kv2 := getKVector.From([]float64{3.14, 2.71, 1.41})
	mv2 := kv2.ToMonomialVector()

	// Test
	kv3 := kv1.Plus(mv2)

	// Check that the sum is of type Kvector
	if _, tf := kv3.(symbolic.KVector); !tf {
		t.Errorf(
			"Expected kv3 to be of type MonomialVector; received %v",
			kv3,
		)
	}

	// Check that the sum is correct
	for ii := 0; ii < N; ii++ {
		if float64(kv3.(symbolic.KVector).AtVec(ii).(symbolic.K)) != float64(kv1.AtVec(ii).(symbolic.K))+float64(kv2.AtVec(ii).(symbolic.K)) {
			t.Errorf(
				"Expected kv3.AtVec(%v) to be %v; received %v",
				ii,
				kv3.(symbolic.KVector).AtVec(ii),
				float64(kv1.AtVec(ii).(symbolic.K))+float64(kv2.AtVec(ii).(symbolic.K)),
			)
		}
	}
}

/*
TestConstantVector_Plus8
Description:

	Tests that the Plus() method correctly computes the sum of a
	KVector and a vector of Polynomials (PolynomialVector).
*/
func TestConstantVector_Plus8(t *testing.T) {
	// Constants
	N := 3
	kv1 := getKVector.From(symbolic.OnesVector(N))
	pv2 := symbolic.PolynomialVector{
		symbolic.Polynomial{
			Monomials: []symbolic.Monomial{
				symbolic.Monomial{
					Coefficient:     3.14,
					VariableFactors: []symbolic.Variable{symbolic.NewVariable()},
					Exponents:       []int{1},
				},
			},
		},
		symbolic.Polynomial{
			Monomials: []symbolic.Monomial{
				symbolic.Monomial{
					Coefficient:     2.71,
					VariableFactors: []symbolic.Variable{symbolic.NewVariable()},
					Exponents:       []int{1},
				},
			},
		},
		symbolic.Polynomial{
			Monomials: []symbolic.Monomial{
				symbolic.Monomial{
					Coefficient:     1.41,
					VariableFactors: []symbolic.Variable{symbolic.NewVariable()},
					Exponents:       []int{1},
				},
			},
		},
	}

	// Test
	sum := kv1.Plus(pv2)

	// Check that the sum is of type polynomial vector
	if _, tf := sum.(symbolic.PolynomialVector); !tf {
		t.Errorf(
			"Expected sum to be of type MonomialVector; received %v",
			sum,
		)
	}

	// Check that the sum contains two monomials in every
	// polynomial.
	for ii := 0; ii < N; ii++ {
		if len(sum.(symbolic.PolynomialVector)[ii].Monomials) != 2 {
			t.Errorf(
				"Expected len(sum[%v].Monomials) to be 2; received %v",
				ii,
				len(sum.(symbolic.PolynomialVector)[ii].Monomials),
			)
		}
	}
}

/*
TestConstantVector_Plus9
Description:

	Tests that the Plus() method properly panics
	when given an unsupported input (string).
*/
func TestConstantVector_Plus9(t *testing.T) {
	// Constants
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(3))
	var input string = "This is a test string."

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Plus() did not panic when given an unsupported input (string).")
		}
	}()

	// Run function
	kv1.Plus(input)
}

/*
TestConstantVector_LessEq1
Description:

	Verifies that the LessEq method panics when two vectors are compared with
	different lengths.
*/
func TestConstantVector_LessEq1(t *testing.T) {
	// Constants
	kv1 := symbolic.VecDenseToKVector(symbolic.OnesVector(3))
	kv2 := symbolic.VecDenseToKVector(symbolic.OnesVector(4))

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected kv1.LessEq(kv2) to panic; received %v",
				kv1.LessEq(kv2),
			)
		}

		rAsError, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected r to be an error; received %v of type %T",
				r, r,
			)
		}

		var sense0 symbolic.ConstrSense = symbolic.SenseLessThanEqual
		if rAsError.Error() != (smErrors.DimensionError{
			Operation: "Comparison (" + sense0.String() + ")",
			Arg1:      kv1,
			Arg2:      kv2,
		}).Error() {
			t.Errorf(
				"Expected r to be a DimensionError; received %v",
				r,
			)
		}
	}()

	kv1.LessEq(kv2)
}
