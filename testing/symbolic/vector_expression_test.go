package symbolic_test

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
vector_expression_test.go
Description:

	Tests for the functions mentioned in the vector_expression.go file.
*/

/*
TestVectorExpression_ToVectorExpression1
Description:

	Tests the conversion of a K() to a vector expression.
	An error should be returned along with a small KVector.
*/
func TestVectorExpression_ToVectorExpression1(t *testing.T) {
	// Constants
	k := symbolic.K(2)

	// Test
	_, err := symbolic.ToVectorExpression(k)
	if err == nil {
		t.Errorf(
			"Expected an error to be returned; received nil",
		)
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf(
				"the input interface is of type %T, which is not recognized as a VectorExpression.",
				k,
			),
		) {
			t.Errorf("expected error message to contain %v; received %v",
				fmt.Sprintf("the input interface is of type %T, which is not recognized as a VectorExpression.",
					k,
				),
				err.Error(),
			)

		}
	}
}
