package smErrors

import "fmt"

type LinearExpressionRequiredError struct {
	Operation  string
	Expression interface{}
}

func (lere LinearExpressionRequiredError) Error() string {
	return fmt.Sprintf(
		"Linear expression required for operation %v; received an expression which is not linear (%T).",
		lere.Operation,
		lere.Expression,
	)
}
