package getKMatrix_test

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"testing"
)

/*
from_test.go
Description:

	This file contains tests for the From() function in get/KMatrix/from.go.
*/

/*
TestFrom1
Description:

	Tests that the From() function returns a KMatrix object when given a Dense object.
*/
func TestFrom1(t *testing.T) {
	// Constants
	dense := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Test
	km := getKMatrix.From(dense)
	if km.Dims()[0] != 3 || km.Dims()[1] != 3 {
		t.Errorf(
			"Expected km to have dimensions 3x3; received %vx%v",
			km.Dims()[0],
			km.Dims()[1],
		)
	}

	// Test that the elements are the same
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if float64(km.At(i, j).(symbolic.K)) != dense.At(i, j) {
				t.Errorf(
					"Expected km.At(%v, %v) to be %v; received %v",
					i,
					j,
					dense.At(i, j),
					km.At(i, j),
				)
			}
		}
	}

}

/*
TestFrom2
Description:

	Tests that the From() function returns a KMatrix object when given a Dense object.
*/
func TestFrom2(t *testing.T) {
	// Constants
	dense := symbolic.OnesMatrix(3, 2)

	// Test
	km := getKMatrix.From(dense)
	if km.Dims()[0] != 3 || km.Dims()[1] != 2 {
		t.Errorf(
			"Expected km to have dimensions 3x2; received %vx%v",
			km.Dims()[0],
			km.Dims()[1],
		)
	}

	// Test that the elements are the same
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if float64(km.At(i, j).(symbolic.K)) != dense.At(i, j) {
				t.Errorf(
					"Expected km.At(%v, %v) to be %v; received %v",
					i,
					j,
					dense.At(i, j),
					km.At(i, j),
				)
			}
		}
	}
}

/*
TestFrom3
Description:

	Tests that the From() function returns a KMatrix object when given a KMatrix object.
*/
func TestFrom3(t *testing.T) {
	// Constants
	d1 := symbolic.OnesMatrix(3, 2)
	km := symbolic.DenseToKMatrix(d1)

	// Test
	km2 := getKMatrix.From(km)
	if km2.Dims()[0] != 3 || km2.Dims()[1] != 2 {
		t.Errorf(
			"Expected km to have dimensions 3x2; received %vx%v",
			km2.Dims()[0],
			km2.Dims()[1],
		)
	}

	// Test that the elements are the same
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if float64(km2.At(i, j).(symbolic.K)) != d1.At(i, j) {
				t.Errorf(
					"Expected km.At(%v, %v) to be %v; received %v",
					i,
					j,
					km.At(i, j),
					km2.At(i, j),
				)
			}
		}
	}
}

/*
TestFrom4
Description:

	Tests that the From() function panics when given an unsupported type.
*/
func TestFrom4(t *testing.T) {
	// Constants
	unsupportedType := "unsupported type"

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(
				"Expected From() to panic when given an unsupported type; did not panic!",
			)
		}

		rAsE, tf := r.(error)
		if !tf {
			t.Errorf(
				"Expected From() to panic with an error; received %v",
				r,
			)
		}

		if rAsE.Error() != (smErrors.UnsupportedInputError{
			FunctionName: "getKMatrix.From",
			Input:        unsupportedType,
		}).Error() {
			t.Errorf(
				"Expected From() to panic with an error; received %v",
				r,
			)
		}
	}()

	getKMatrix.From(unsupportedType)
}

/*
TestFrom5
Description:

	Tests that the From() function returns a KMatrix object when given a
	[][]float64 object.
*/
func TestFrom5(t *testing.T) {
	// Constants
	floatSliceSlice1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Test
	km := getKMatrix.From(floatSliceSlice1)
	if km.Dims()[0] != 2 || km.Dims()[1] != 3 {
		t.Errorf(
			"Expected km to have dimensions 2x3; received %vx%v",
			km.Dims()[0],
			km.Dims()[1],
		)
	}

	// Test that the elements are the same
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if float64(km.At(i, j).(symbolic.K)) != floatSliceSlice1[i][j] {
				t.Errorf(
					"Expected km.At(%v, %v) to be %v; received %v",
					i,
					j,
					floatSliceSlice1[i][j],
					km.At(i, j),
				)
			}
		}
	}
}
