package smErrors

import "fmt"

/*
out_of_bounds_error.go
Description:
	Defining the out of bounds error object and all of its associated functions.
*/

type InvalidVectorIndexError struct {
	Index      int
	Expression MatrixLike
}

type InvalidMatrixIndexError struct {
	RowIndex   int
	ColIndex   int
	Expression MatrixLike
}

func (ivi InvalidVectorIndexError) Error() string {
	errOut := fmt.Sprintf(
		"out of bounds error: index %v is out of bounds for %T object of dimension %v",
		ivi.Index,
		ivi.Expression,
		ivi.Expression.Dims(),
	)
	return errOut
}

func (imi InvalidMatrixIndexError) Error() string {
	errOut := fmt.Sprintf(
		"out of bounds error: index (%v, %v) is out of bounds for %T object of dimension %v",
		imi.RowIndex,
		imi.ColIndex,
		imi.Expression,
		imi.Expression.Dims(),
	)
	return errOut
}

// Explicit Functions
func CheckIndexOnVector(index int, vector MatrixLike) error {
	if index < 0 || index >= vector.Dims()[0] {
		return InvalidVectorIndexError{
			Index:      index,
			Expression: vector,
		}
	}
	return nil
}

/*
CheckIndexOnMatrix
Description:

	Checks that the index is valid for the matrix.
*/
func CheckIndexOnMatrix(rowIndex int, colIndex int, matrix MatrixLike) error {
	// Check that row index is in bounds
	if rowIndex < 0 || rowIndex >= matrix.Dims()[0] {
		return InvalidMatrixIndexError{
			RowIndex:   rowIndex,
			ColIndex:   colIndex,
			Expression: matrix,
		}
	}

	// Check that column index is in bounds
	if colIndex < 0 || colIndex >= matrix.Dims()[1] {
		return InvalidMatrixIndexError{
			RowIndex:   rowIndex,
			ColIndex:   colIndex,
			Expression: matrix,
		}
	}

	// Otherwise, return nil
	return nil
}
