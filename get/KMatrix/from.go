package getKMatrix

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
from.go
Description:

	This file contains functions that convert from other types to KMatrix objects.
*/
func From(expr interface{}) symbolic.KMatrix {
	switch converted := expr.(type) {
	case symbolic.KMatrix:
		return converted
	case mat.Dense:
		// Return
		return symbolic.DenseToKMatrix(converted)
	case *mat.Dense:
		// Return
		return symbolic.DenseToKMatrix(*converted)
	case [][]float64:
		// Input Processing

		// Check that the input has the same number of columns
		// in each row.

		nCols := len(converted[0])
		for ii, row := range converted {
			if len(row) != nCols {
				panic(
					smErrors.MatrixColumnMismatchError{
						ExpectedNColumns: nCols,
						ActualNColumns:   len(row),
						Row:              ii,
					},
				)
			}
		}

		// Create a container for each of the rows
		var out symbolic.KMatrix

		for _, row := range converted {
			newRowOut := make([]symbolic.K, len(row))
			for ii, element := range row {
				newRowOut[ii] = symbolic.K(element)
			}
			out = append(out, newRowOut)
		}
		return out
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "getKMatrix.From",
				Input:        converted,
			},
		)
	}

}
