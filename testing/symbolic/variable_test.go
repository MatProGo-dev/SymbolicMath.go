package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
variable_test.go
Description:
	Testing functions relevant to the Var() object. (Scalar Variable)
*/

/*
TestVariable_NumVars1
Description:

	Tests whether or not NumVars returns 1 for a single variable.
*/
func TestVariable_NumVars1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.NumVars(x) != 1 {
		t.Errorf(
			"The number of variables in a %T should be 1; received %v",
			x,
			x.NumVars(),
		)
	}

}

/*
TestVariable_Constant1
Description:

	Tests whether or not NumVars returns 0 as the constant included in the a single variable.
*/
func TestVariable_Constant1(t *testing.T) {
	// Constants
	m := optim.NewModel("Var-Constant1")
	x := m.AddVariable()

	// Test
	if x.Constant() != 0.0 {
		t.Errorf(
			"The number of variables in a %T should be 1; received %v",
			x,
			x.NumVars(),
		)
	}

}

/*
TestVariable_Plus1
Description:

	Tests the approach of performing addition of a var with a constant.
*/
func TestVariable_Plus1(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus1")
	x := m.AddVariable()

	k1 := optim.K(1.0)

	// Algorithm
	tempSum, err := x.Plus(k1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 1 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.C != float64(k1) {
		t.Errorf("Expected sum's constant to be %v; received %v.", k1, sumAsSLE.C)
	}
}

/*
TestVariable_Plus2
Description:

	Tests the approach of performing addition of a var with a var.
*/
func TestVariable_Plus2(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus2")
	x := m.AddVariable()
	y := m.AddVariable()

	// Algorithm
	tempSum, err := x.Plus(y)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 2 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.C != 0.0 {
		t.Errorf("Expected sum's constant to be %v; received %v.", 0.0, sumAsSLE.C)
	}
}

/*
TestVariable_Plus3
Description:

	Tests the approach of performing addition of a var with a scalar linear expression.
*/
func TestVariable_Plus3(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus3")
	x := m.AddVariable()
	y := m.AddVariable()
	z := m.AddVariable()

	vv := optim.VarVector{[]optim.Variable{y, z}}
	L := optim.OnesVector(2)
	C := 3.0
	sle1 := optim.ScalarLinearExpr{
		X: vv,
		L: L,
		C: C,
	}

	// Algorithm
	tempSum, err := x.Plus(sle1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.X.AtVec(2).(optim.Variable).ID != z.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(2).(optim.Variable).ID,
			z.ID,
		)
	}

	// Checking linear terms
	if sumAsSLE.L.AtVec(0) != 1.0 {
		t.Errorf(
			"Expected first element of L (%v) to have value 1.0.",
			sumAsSLE.L.AtVec(2),
		)
	}

	if sumAsSLE.L.AtVec(1) != sle1.L.AtVec(0) {
		t.Errorf(
			"Expected second element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(1),
			sle1.L.AtVec(0),
		)
	}

	if sumAsSLE.L.AtVec(2) != sle1.L.AtVec(1) {
		t.Errorf(
			"Expected first element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(2),
			sle1.L.AtVec(1),
		)
	}

	// Checking Constant terms
	if sumAsSLE.C != sle1.C {
		t.Errorf("Expected sum's constant to be %v; received %v.", sle1.C, sumAsSLE.C)
	}
}

/*
TestVariable_Plus4
Description:

	Tests the approach of performing addition of a var with a scalar quadratic expression.
*/
func TestVariable_Plus4(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus4")
	x := m.AddVariable()
	y := m.AddVariable()
	z := m.AddVariable()

	vv := optim.VarVector{[]optim.Variable{y, z}}
	Q := optim.Identity(2)
	L := optim.OnesVector(2)
	C := 3.0
	qe1 := optim.ScalarQuadraticExpression{
		Q: Q,
		X: vv,
		L: L,
		C: C,
	}

	// Algorithm
	tempSum, err := x.Plus(qe1)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarQuadraticExpression)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 3 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.X.AtVec(1).(optim.Variable).ID != y.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(1).(optim.Variable).ID,
			y.ID,
		)
	}

	if sumAsSLE.X.AtVec(2).(optim.Variable).ID != z.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(2).(optim.Variable).ID,
			z.ID,
		)
	}

	// Checking linear terms
	if sumAsSLE.L.AtVec(0) != 1.0 {
		t.Errorf(
			"Expected first element of L (%v) to have value 1.0.",
			sumAsSLE.L.AtVec(2),
		)
	}

	if sumAsSLE.L.AtVec(1) != qe1.L.AtVec(0) {
		t.Errorf(
			"Expected second element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(1),
			qe1.L.AtVec(0),
		)
	}

	if sumAsSLE.L.AtVec(2) != qe1.L.AtVec(1) {
		t.Errorf(
			"Expected first element of L (%v) to have value %v.",
			sumAsSLE.L.AtVec(2),
			qe1.L.AtVec(1),
		)
	}

	// Checking Constant terms
	if sumAsSLE.C != qe1.C {
		t.Errorf("Expected sum's constant to be %v; received %v.", qe1.C, sumAsSLE.C)
	}
}

/*
TestVariable_Plus5
Description:

	Tests that the Plus method properly throws an error.
*/
func TestVariable_Plus5(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus1")
	x := m.AddVariable()

	k1 := optim.K(1.0)
	err0 := fmt.Errorf("Test")

	// Algorithm
	_, err := x.Plus(k1, err0)
	if err == nil {
		t.Errorf("There was not an error, when we expected!")
	} else {
		if !strings.Contains(
			err.Error(),
			err0.Error(),
		) {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

/*
TestVariable_Plus6
Description:

	Tests the approach of performing addition of a var with a var
	when the var is the same.
*/
func TestVariable_Plus6(t *testing.T) {
	// Constants
	m := optim.NewModel("Plus2")
	x := m.AddVariable()

	// Algorithm
	tempSum, err := x.Plus(x)
	if err != nil {
		t.Errorf("There was an issue computing sum: %v", err)
	}

	sumAsSLE, ok1 := tempSum.(optim.ScalarLinearExpr)
	if !ok1 {
		t.Errorf("The sum is expected to be a ScalarLinearExpression but it was not! (Type = %T)", tempSum)
	}

	if sumAsSLE.X.Len() != 1 {
		t.Errorf("Expected sum to have scalar variable, but found %v items.", sumAsSLE.X.Len())
	}

	if sumAsSLE.X.AtVec(0).(optim.Variable).ID != x.ID {
		t.Errorf(
			"Expected scalar variable (ID %v) to have same ID as x's ID (%v), but they are different.",
			sumAsSLE.X.AtVec(0).(optim.Variable).ID,
			x.ID,
		)
	}

	if sumAsSLE.L.AtVec(0) != 2.0 {
		t.Errorf(
			"Expected L[0] = 2.0; received %v",
			sumAsSLE.L.AtVec(0),
		)
	}

	if sumAsSLE.C != 0.0 {
		t.Errorf("Expected sum's constant to be %v; received %v.", 0.0, sumAsSLE.C)
	}
}

/*
TestVariable_Multiply1
Description:

	Tests how well the Multiply() function works between a variable and a float.
*/
func TestVariable_Multiply1(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()
	f1 := 3.14

	// Algorithm
	prod, err := v1.Multiply(f1)
	if err != nil {
		t.Errorf("Error multiplying variable with float: %v", err)
	}

	prodAsSLE, ok := prod.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSLE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSLE.X.Len(),
		)
	}

	if prodAsSLE.X.AtVec(0).IDs()[0] != v1.ID {
		t.Errorf(
			"Expected the variable in X be %v; received %v",
			v1, prodAsSLE.X.AtVec(0),
		)
	}

	if prodAsSLE.L.AtVec(0) != f1 {
		t.Errorf(
			"Expected linear coefficient to be %v; received %v.",
			f1,
			prodAsSLE.L.AtVec(0),
		)
	}
}

/*
TestVariable_Multiply2
Description:

	Tests how well the Multiply() function works between a variable and a K.
*/
func TestVariable_Multiply2(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply2")
	v1 := m.AddVariable()
	k1 := optim.K(3.14)

	// Algorithm
	prod, err := v1.Multiply(k1)
	if err != nil {
		t.Errorf("Error multiplying variable with constant K: %v", err)
	}

	prodAsSLE, ok := prod.(optim.ScalarLinearExpr)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarLinearExpr; received %T",
			prod,
		)
	}

	if prodAsSLE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSLE.X.Len(),
		)
	}

	if prodAsSLE.X.AtVec(0).IDs()[0] != v1.ID {
		t.Errorf(
			"Expected the variable in X be %v; received %v",
			v1, prodAsSLE.X.AtVec(0),
		)
	}

	if prodAsSLE.L.AtVec(0) != float64(k1) {
		t.Errorf(
			"Expected linear coefficient to be %v; received %v.",
			k1,
			prodAsSLE.L.AtVec(0),
		)
	}
}

/*
TestVariable_Multiply3
Description:

	Tests how well the Multiply() function works between a variable and a variable (different).
*/
func TestVariable_Multiply3(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()
	v2 := m.AddVariable()

	// Algorithm
	prod, err := v1.Multiply(v2)
	if err != nil {
		t.Errorf("Error multiplying variable with variable: %v", err)
	}

	prodAsSQE, ok := prod.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarQuadraticExpression; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 2 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSQE.X.Len(),
		)
	}

	// Verify that v1 and v2 are in the X
	if v1Index, _ := optim.FindInSlice(v1, prodAsSQE.X.Elements); v1Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v1,
		)
	}

	if v2Index, _ := optim.FindInSlice(v2, prodAsSQE.X.Elements); v2Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v2,
		)
	}

	for xIndex := 0; xIndex < prodAsSQE.X.Len(); xIndex++ {
		if prodAsSQE.L.AtVec(xIndex) != 0.0 {
			t.Errorf(
				"Expected linear coefficient to be %v; received %v.",
				0.0,
				prodAsSQE.L.AtVec(xIndex),
			)
		}
	}

}

/*
TestVariable_Multiply4
Description:

	Tests how well the Multiply() function works between a variable and a variable (same as original).
*/
func TestVariable_Multiply4(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()

	// Algorithm
	prod, err := v1.Multiply(v1)
	if err != nil {
		t.Errorf("Error multiplying vector with float: %v", err)
	}

	prodAsSQE, ok := prod.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarQuadraticExpression; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSQE.X.Len(),
		)
	}

	// Verify that v1 and v2 are in the X
	if v1Index, _ := optim.FindInSlice(v1, prodAsSQE.X.Elements); v1Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v1,
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 1.0 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v.",
			3.14,
			prodAsSQE.Q.At(0, 0),
		)
	}

	if prodAsSQE.L.AtVec(0) != 0.0 {
		t.Errorf(
			"Expected linear coefficient to be %v; received %v.",
			0.0,
			prodAsSQE.L.AtVec(0),
		)
	}

}

/*
TestVariable_Multiply5
Description:

	Tests how well the Multiply() function works between a variable and a scalar linear expression.
*/
func TestVariable_Multiply5(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()

	// Algorithm
	prod, err := v1.Multiply(v1.Multiply(3.14))
	if err != nil {
		t.Errorf("Error multiplying variable with sle: %v", err)
	}

	prodAsSQE, ok := prod.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarQuadraticExpression; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 1 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSQE.X.Len(),
		)
	}

	// Verify that v1 and v2 are in the X
	if v1Index, _ := optim.FindInSlice(v1, prodAsSQE.X.Elements); v1Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v1,
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 3.14 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v.",
			3.14,
			prodAsSQE.Q.At(0, 0),
		)
	}

	if prodAsSQE.L.AtVec(0) != 0.0 {
		t.Errorf(
			"Expected linear coefficient to be %v; received %v.",
			0.0,
			prodAsSQE.L.AtVec(0),
		)
	}

}

/*
TestVariable_Multiply6
Description:

	Tests how well the Multiply() function works between a variable and a scalar linear expression.
*/
func TestVariable_Multiply6(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()
	v2 := m.AddVariable()

	// Algorithm
	prod, err := v1.Multiply(v2.Multiply(3.14))
	if err != nil {
		t.Errorf("Error multiplying variable with sle: %v", err)
	}

	prodAsSQE, ok := prod.(optim.ScalarQuadraticExpression)
	if !ok {
		t.Errorf(
			"Expected product to be ScalarQuadraticExpression; received %T",
			prod,
		)
	}

	if prodAsSQE.X.Len() != 2 {
		t.Errorf(
			"Expected product to contain a single variable; received %v.",
			prodAsSQE.X.Len(),
		)
	}

	// Verify that v1 and v2 are in the X
	if v1Index, _ := optim.FindInSlice(v1, prodAsSQE.X.Elements); v1Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v1,
		)
	}

	if v2Index, _ := optim.FindInSlice(v2, prodAsSQE.X.Elements); v2Index == -1 {
		t.Errorf(
			"Expected for %v to be in X; it was not!",
			v2,
		)
	}

	// Check constants
	if prodAsSQE.Q.At(0, 0) != 0.0 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v.",
			0.0,
			prodAsSQE.Q.At(0, 0),
		)
	}

	if prodAsSQE.Q.At(1, 1) != 0.0 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v.",
			0.0,
			prodAsSQE.Q.At(1, 1),
		)
	}

	if prodAsSQE.Q.At(0, 1) != 3.14*0.5 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v",
			3.14*0.5,
			prodAsSQE.Q.At(0, 1),
		)
	}

	if prodAsSQE.Q.At(1, 0) != 3.14*0.5 {
		t.Errorf(
			"Expected quadratic coefficient to be %v; received %v",
			3.14*0.5,
			prodAsSQE.Q.At(1, 0),
		)
	}

	if prodAsSQE.L.AtVec(0) != 0.0 {
		t.Errorf(
			"Expected linear coefficient to be %v; received %v.",
			0.0,
			prodAsSQE.L.AtVec(0),
		)
	}

}

/*
TestVariable_Multiply7
Description:

	Tests how well the Multiply() function works between a variable and
	a scalar quadratic expression. Should produce an error
*/
func TestVariable_Multiply7(t *testing.T) {
	//Constants
	m := optim.NewModel("Test-Variable-Multiply1")
	v1 := m.AddVariable()
	v2 := m.AddVariable()

	// Algorithm
	_, err := v1.Multiply(v1.Multiply(v2.Multiply(3.14)))
	if err == nil {
		t.Errorf("Error should occur multiplying variable with sqe, but none did!")
	}

	if strings.Compare(
		err.Error(),
		"Can not multiply Variable with ScalarQuadraticExpression. MatProInterface can not represent polynomials higher than degree 2.",
	) != 0 {
		t.Errorf(
			"Expected for specific error to occur, but received %v", err)
	}

}
