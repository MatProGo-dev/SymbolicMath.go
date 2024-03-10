package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
matrix_expression_test.go
Description:

	Tests the MatrixExpression object.
*/

/*
TestMatrixExpression_ToMatrixExpression1
Description:

	Tests that the conversion method fails if the input is not a valid MatrixExpression.
*/
func TestMatrixExpression_ToMatrixExpression1(t *testing.T) {
	// Constants
	x := 3

	// Test
	_, err := symbolic.ToMatrixExpression(x)
	if err == nil {
		t.Errorf("Expected error when converting %v to a MatrixExpression; received nil", x)
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the input interface is of type %T, which is not recognized as a MatrixExpression.",
				x,
			),
		) {
			t.Errorf("Expected error message to contain %v; received %v",
				fmt.Sprintf(
					"the input interface is of type %T, which is not recognized as a MatrixExpression.",
					x,
				),
				err.Error(),
			)
		}
	}
}
