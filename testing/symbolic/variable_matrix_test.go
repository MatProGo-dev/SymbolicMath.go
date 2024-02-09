package symbolic_test

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
variable_matrix_test.go
Description:

	Tests the functions in the variable_matrix.go file.
*/

/*
TestVariableMatrix_Check1
Description:

	Tests the Check method for a VariableMatrix object that is empty.
*/
func TestVariableMatrix_Check1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Test
	expectedError := smErrors.EmptyMatrixError{
		Expression: vm,
	}
	err := vm.Check()
	if err == nil {
		t.Errorf("Expected Check to return error; received nil")
		return
	}

	if err.Error() != expectedError.Error() {
		t.Errorf(
			"Expected Check to return \"%v\"; received \"%v\"",
			expectedError.Error(),
			err.Error(),
		)
	}
}

/*
TestVariableMatrix_Check2
Description:

	Tests the Check method for a VariableMatrix object that has a mismatch in the number of columns.
*/
func TestVariableMatrix_Check2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.Variable{}, symbolic.Variable{}},
		{symbolic.Variable{}},
	}

	// Test
	expectedError := smErrors.MatrixColumnMismatchError{
		ExpectedNColumns: 2,
		ActualNColumns:   1,
		Row:              1,
	}
	err := vm.Check()
	if err == nil {
		t.Errorf("Expected Check to return error; received nil")
		return
	}

	if err.Error() != expectedError.Error() {
		t.Errorf(
			"Expected Check to return \"%v\"; received \"%v\"",
			expectedError.Error(),
			err.Error(),
		)
	}
}

/*
TestVariableMatrix_Check3
Description:

	Tests the Check method for a VariableMatrix object that has
	a single variable that is not well-defined.
*/
func TestVariableMatrix_Check3(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.Variable{}},
	}

	// Test
	expectedError := vm[0][0].Check()
	err := vm.Check()
	if err == nil {
		t.Errorf("Expected Check to return error; received nil")
		return
	}

	if !strings.Contains(
		err.Error(),
		expectedError.Error(),
	) {
		t.Errorf(
			"Expected Check to return \"%v\"; received \"%v\"",
			expectedError.Error(),
			err.Error(),
		)
	}
}

/*
TestVariableMatrix_Check4
Description:

	Tests the Check method for a VariableMatrix object that is well-defined.
*/
func TestVariableMatrix_Check4(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	err := vm.Check()
	if err != nil {
		t.Errorf("Expected Check to return nil; received \"%v\"", err.Error())
	}
}

/*
TestVariableMatrix_Variables1
Description:

	Tests the Variables method for a VariableMatrix object that is empty.
	A panic should be thrown.
*/
func TestVariableMatrix_Variables1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Variables to panic; received nil")
		}
	}()

	vm.Variables()
}

/*
TestVariableMatrix_Variables2
Description:

	Tests the Variables method for a VariableMatrix object that is well-defined.
	For this 2x2 matrix, the expected output is a slice of 4 variables.
*/
func TestVariableMatrix_Variables2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	variables := vm.Variables()
	if len(variables) != 4 {
		t.Errorf(
			"Expected Variables to return a slice of length 4; received %v",
			len(variables),
		)
	}
}

/*
TestVariableMatrix_Variables3
Description:

	Tests the Variables method for a VariableMatrix object that has a single variable.
*/
func TestVariableMatrix_Variables3(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	vm := symbolic.VariableMatrix{
		{v1, v1},
		{v1, v1},
	}

	// Test
	variables := vm.Variables()
	if len(variables) != 1 {
		t.Errorf(
			"Expected Variables to return a slice of length 1; received %v",
			len(variables),
		)
	}
}

/*
TestVariableMatrix_Dims1
Description:

	Tests the Dims method for a VariableMatrix object that is empty.
	A panic should be thrown.
*/
func TestVariableMatrix_Dims1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Dims to panic; received nil")
		}
	}()

	vm.Dims()
}

/*
TestVariableMatrix_Dims2
Description:

	Tests the Dims method for a VariableMatrix object that is well-defined.
	For this 2x2 matrix, the expected output is (2, 2).
*/
func TestVariableMatrix_Dims2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	dims := vm.Dims()
	if dims[0] != 2 || dims[1] != 2 {
		t.Errorf(
			"Expected Dims to return (2, 2); received (%v, %v)",
			dims[0],
			dims[1],
		)
	}
}

/*
TestVariableMatrix_Plus1
Description:

	Tests the Plus method for a VariableMatrix object that is empty.
	A panic should be thrown.
*/
func TestVariableMatrix_Plus1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Plus to panic; received nil")
		}
	}()

	vm.Plus(vm)
}

/*
TestVariableMatrix_Plus2
Description:

	Tests the Plus method for a VariableMatrix object when
	added to a not well-defined variable matrix.
*/
func TestVariableMatrix_Plus2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	vm2 := symbolic.VariableMatrix{
		{symbolic.Variable{}, symbolic.Variable{}},
		{symbolic.Variable{}, symbolic.Variable{}},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Plus to panic; received nil")
		}
	}()

	vm.Plus(vm2)
}

/*
TestVariableMatrix_Plus3
Description:

	Tests the Plus method for a VariableMatrix object that is well-defined
	being added to a well-defined expression (a constant matrix).
	The two matrices will be mismatched in dimensions which should
	cause a panic.
*/
func TestVariableMatrix_Plus3(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1, 2},
			{3, 4},
			{5, 6},
		})

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Plus to panic; received nil")
		}
	}()

	vm.Plus(cm)
}

/*
TestVariableMatrix_Plus4
Description:

	Tests the Plus method for a VariableMatrix object that is well-defined
	being added to a well-defined expression (a constant matrix).
	The two matrices will be matched in dimensions.
	Checks that the result is a polynomial matrix.
*/
func TestVariableMatrix_Plus4(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1, 2},
			{3, 4},
		})

	// Test
	result := vm.Plus(cm)

	// Check that object is a polynomial matrix
	if _, ok := result.(symbolic.PolynomialMatrix); !ok {
		t.Errorf("Expected Plus to return a PolynomialMatrix; received %T", result)
	}

	// Check that each polynomial in the result contains two monomials.
	pm := result.(symbolic.PolynomialMatrix)
	for i := 0; i < pm.Dims()[0]; i++ {
		for j := 0; j < pm.Dims()[1]; j++ {
			if len(pm[i][j].Monomials) != 2 {
				t.Errorf("Expected each polynomial to contain 2 monomials; received %v", len(pm[i][j].Monomials))
			}
		}
	}
}

/*
TestVariableMatrix_Plus5
Description:

	Tests the Plus method for a VariableMatrix object that is well-defined
	being added to a well-defined expression (a float64).
	Checks that the result is a PolynomialMatrix.
*/
func TestVariableMatrix_Plus5(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	result := vm.Plus(1.0)

	// Check that object is a PolynomialMatrix
	if _, ok := result.(symbolic.PolynomialMatrix); !ok {
		t.Errorf("Expected Plus to return a KMatrix; received %T", result)
	}

	// Check that each polynomial in the result contains two monomials.
	pm := result.(symbolic.PolynomialMatrix)
	for i := 0; i < pm.Dims()[0]; i++ {
		for j := 0; j < pm.Dims()[1]; j++ {
			if len(pm[i][j].Monomials) != 2 {
				t.Errorf("Expected each polynomial to contain 2 monomial; received %v", len(pm[i][j].Monomials))
			}
		}
	}
}

/*
TestVariableMatrix_Plus6
Description:

	Tests the Plus method for a VariableMatrix object that is well-defined
	being added to a non-expression (a string).
	Checks that the method panics.
*/
func TestVariableMatrix_Plus6(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Plus to panic; received nil")
		}
	}()

	vm.Plus("hello")
}

/*
TestVariableMatrix_Multiply1
Description:

	Tests the Multiply method for a VariableMatrix object that is empty.
	A panic should be thrown.
*/
func TestVariableMatrix_Multiply1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Multiply to panic; received nil")
		}
	}()

	vm.Multiply(vm)
}

/*
TestVariableMatrix_Multiply2
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	being multiplied by a NOT well-defined expression (a variable matrix).
	A panic should be thrown.
*/
func TestVariableMatrix_Multiply2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	vm2 := symbolic.VariableMatrix{
		{symbolic.Variable{}, symbolic.Variable{}},
		{symbolic.Variable{}, symbolic.Variable{}},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Multiply to panic; received nil")
		}
	}()

	vm.Multiply(vm2)
}

/*
TestVariableMatrix_Multiply3
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	being multiplied by a well-defined expression (a constant matrix).
	The two matrices will be mismatched in dimensions which should
	cause a panic.
*/
func TestVariableMatrix_Multiply3(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1, 2},
			{3, 4},
			{5, 6},
		})

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Multiply to panic; received nil")
		}
	}()

	vm.Multiply(cm)
}

/*
TestVariableMatrix_Multiply4
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	being multiplied by a well-defined expression (a constant matrix).
	The two matrices will be matched in dimensions.
	Checks that the result is a polynomial matrix when the second expression has more than one column.
*/
func TestVariableMatrix_Multiply4(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1, 2},
			{3, 4},
		})

	// Test
	result := vm.Multiply(cm)

	// Check that object is a polynomial matrix
	if _, ok := result.(symbolic.PolynomialMatrix); !ok {
		t.Errorf("Expected Multiply to return a PolynomialMatrix; received %T", result)
	}

	// Check that each polynomial in the result contains two monomials.
	pm := result.(symbolic.PolynomialMatrix)
	for i := 0; i < pm.Dims()[0]; i++ {
		for j := 0; j < pm.Dims()[1]; j++ {
			if len(pm[i][j].Monomials) != 2 {
				t.Errorf("Expected each polynomial to contain 2 monomials; received %v", len(pm[i][j].Monomials))
			}
		}
	}
}

/*
TestVariableMatrix_Multiply5
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	being multiplied by a float64.
	Checks that the result is a MonomialMatrix.
*/
func TestVariableMatrix_Multiply5(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	result := vm.Multiply(2.0)

	// Check that object is a MonomialMatrix
	if _, ok := result.(symbolic.MonomialMatrix); !ok {
		t.Errorf("Expected Multiply to return a MonomialMatrix; received %T", result)
	}

	// Check that each monomial in the result contains one variable factor
	// and a coefficient of 2.0
	mm := result.(symbolic.MonomialMatrix)
	for _, mmRow := range mm {
		for _, m := range mmRow {
			if len(m.VariableFactors) != 1 {
				t.Errorf("Expected each monomial to contain 1 factor; received %v", len(m.VariableFactors))
			}

			if m.Coefficient != 2.0 {
				t.Errorf("Expected each monomial to contain a coefficient of 2.0; received %v", m.Coefficient)
			}
		}
	}
}

/*
TestVariableMatrix_Multiply6
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	being multiplied by a non-expression (a string).
	Checks that the method panics.
*/
func TestVariableMatrix_Multiply6(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Multiply to panic; received nil")
		}
	}()

	vm.Multiply("hello")
}

/*
TestVariableMatrix_Multiply7
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	of a dimension (3, 2) being multiplied by a KMatrix of dimension (2, 1).
	Checks that the result is a PolynomialVector.
*/
func TestVariableMatrix_Multiply7(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1},
			{2},
		})

	// Test
	result := vm.Multiply(cm)

	// Check that object is a PolynomialVector
	if _, ok := result.(symbolic.PolynomialVector); !ok {
		t.Errorf("Expected Multiply to return a PolynomialVector; received %T", result)
	}

	// Check that each polynomial in the result contains two monomials.
	pv := result.(symbolic.PolynomialVector)
	for i := 0; i < pv.Dims()[0]; i++ {
		if len(pv[i].Monomials) != 2 {
			t.Errorf("Expected each polynomial to contain 2 monomials; received %v", len(pv[i].Monomials))
		}
	}
}

/*
TestVariableMatrix_Multiply8
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	of a dimension (1, 3) being multiplied by a KMatrix of dimension (3, 1).
	Checks that the result is a Polynomial.
*/
func TestVariableMatrix_Multiply8(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cm := getKMatrix.From(
		[][]float64{
			{1},
			{2},
			{3},
		})

	// Test
	result := vm.Multiply(cm)

	// Check that object is a Polynomial
	if _, ok := result.(symbolic.Polynomial); !ok {
		t.Errorf("Expected Multiply to return a Polynomial; received %T", result)
	}

	// Check that the polynomial contains three monomials.
	p := result.(symbolic.Polynomial)
	if len(p.Monomials) != 3 {
		t.Errorf("Expected the polynomial to contain 3 monomials; received %v", len(p.Monomials))
	}
}

/*
TestVariableMatrix_Multiply9
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	of a dimension (3, 2) being multiplied by a KVector of dimension (2, 1).
	The product should be a PolynomialVector of dimension (3, 1)
	with two monomials in each polynomial.
*/
func TestVariableMatrix_Multiply9(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cv := getKMatrix.From(
		[][]float64{
			{1},
			{2},
		})

	// Compute Product
	result := vm.Multiply(cv)

	// Check that object is a PolynomialVector
	if _, ok := result.(symbolic.PolynomialVector); !ok {
		t.Errorf("Expected Multiply to return a PolynomialVector; received %T", result)
	}

	// Check that each polynomial in the result contains two monomials.
	pv := result.(symbolic.PolynomialVector)
	for i := 0; i < pv.Dims()[0]; i++ {
		if len(pv[i].Monomials) != 2 {
			t.Errorf("Expected each polynomial to contain 2 monomials; received %v", len(pv[i].Monomials))
		}
	}
}

/*
TestVariableMatrix_Multiply10
Description:

	Tests the Multiply method for a VariableMatrix object that is well-defined
	of a dimension (1, 3) being multiplied by a KVector of dimension (3, 1).
	The product should be a Polynomial of dimension (1, 1) with three monomials.
*/
func TestVariableMatrix_Multiply10(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable(), symbolic.NewVariable()},
	}
	cv := getKMatrix.From(
		[][]float64{
			{1},
			{2},
			{3},
		})

	// Compute Product
	result := vm.Multiply(cv)

	// Check that object is a Polynomial
	if _, ok := result.(symbolic.Polynomial); !ok {
		t.Errorf("Expected Multiply to return a Polynomial; received %T", result)
	}

	// Check that the polynomial contains three monomials.
	p := result.(symbolic.Polynomial)
	if len(p.Monomials) != 3 {
		t.Errorf("Expected the polynomial to contain 3 monomials; received %v", len(p.Monomials))
	}
}

/*
TestVariableMatrix_Transpose1
Description:

	Tests the Transpose method for a VariableMatrix object that is empty.
	A panic should be thrown.
*/
func TestVariableMatrix_Transpose1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{}

	// Panic handling
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Transpose to panic; received nil")
		}
	}()

	// Test
	vm.Transpose()
}

/*
TestVariableMatrix_Transpose2
Description:

	Tests the Transpose method for a VariableMatrix object that is well-defined.
	Checks that the result is a VariableMatrix with the correct dimensions.
*/
func TestVariableMatrix_Transpose2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	result := vm.Transpose()

	// Check that the result is a VariableMatrix
	if _, ok := result.(symbolic.VariableMatrix); !ok {
		t.Errorf("Expected Transpose to return a VariableMatrix; received %T", result)
	}

	// Check that the dimensions are correct
	dims := result.Dims()
	if dims[0] != vm.Dims()[1] || dims[1] != vm.Dims()[0] {
		t.Errorf(
			"Expected Transpose to return a VariableMatrix with dimensions (%v, %v); received (%v, %v)",
			vm.Dims()[1],
			vm.Dims()[0],
			dims[0],
			dims[1],
		)
	}

	// Check that the result is the transpose of the input
	result2 := result.(symbolic.VariableMatrix)
	for i := 0; i < vm.Dims()[0]; i++ {
		for j := 0; j < vm.Dims()[1]; j++ {
			if vm[i][j].ID != result2[j][i].ID {
				t.Errorf(
					"Expected the result to be the transpose of the input; received %v",
					result,
				)
			}
		}
	}
}

/*
TestVariableMatrix_LessEq1
Description:

	Tests the LessEq method panics for a variable matrix that is not
	well-defined.
*/
func TestVariableMatrix_LessEq1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.Variable{}},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected LessEq to panic; received nil")
		}
	}()

	vm.LessEq(vm)
}

/*
TestVariableMatrix_LessEq2
Description:

	Tests the LessEq method for a VariableMatrix object that is well-defined
	being compared to a float64.
	Verifies that the resulting constraint has:
	- A VariableMatrix on the LeftHandSide
	- A KMatrix on the RightHandSide
	- The sense SenseLessThanEqual
*/
func TestVariableMatrix_LessEq2(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}

	// Test
	constraint := vm.LessEq(2.0)

	// Check that the LeftHandSide is a VariableMatrix
	if _, ok := constraint.Left().(symbolic.VariableMatrix); !ok {
		t.Errorf("Expected the LeftHandSide to be a VariableMatrix; received %T", constraint.Left())
	}

	// Check that the RightHandSide is a KMatrix
	if _, ok := constraint.Right().(symbolic.KMatrix); !ok {
		t.Errorf("Expected the RightHandSide to be a KMatrix; received %T", constraint.Right())
	}

	// Check that the Sense is SenseLessThanEqual
	if constraint.ConstrSense() != symbolic.SenseLessThanEqual {
		t.Errorf("Expected the Sense to be SenseLessThanEqual; received %v", constraint.ConstrSense())
	}
}

/*
TestVariableMatrix_GreaterEq1
Description:

	Tests the GreaterEq method panics
	for a VariableMatrix object that is well-defined
	and a monomial matrix of incompatible dimensions.
	Initial VariableMatrix has dimension (3,2) and the
	monomial matrix has dimension (2,1).
*/
func TestVariableMatrix_GreaterEq1(t *testing.T) {
	// Constants
	vm := symbolic.VariableMatrix{
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
		{symbolic.NewVariable(), symbolic.NewVariable()},
	}
	mm := symbolic.MonomialMatrix{
		{symbolic.NewVariable().ToMonomial()},
		{symbolic.NewVariable().ToMonomial()},
	}

	// Test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected GreaterEq to panic; received nil")
		}
	}()

	vm.GreaterEq(mm)
}
