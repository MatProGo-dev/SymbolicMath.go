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
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "KMatrix.From",
				Input:        converted,
			},
		)
	}

}
