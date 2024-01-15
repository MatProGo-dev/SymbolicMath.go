package smErrors

import "fmt"

type MatrixColumnMismatchError struct {
	ExpectedNColumns int
	ActualNColumns   int
	Row              int
}

func (mcme MatrixColumnMismatchError) Error() string {
	return fmt.Sprintf(
		"matrix column mismatch error: expected %v columns, received %v in row %v",
		mcme.ExpectedNColumns,
		mcme.ActualNColumns,
		mcme.Row,
	)
}
