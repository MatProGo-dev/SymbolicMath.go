package getKVector

/*
from.go
Description:

	This file contains functions that convert from other types to KVector objects.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
From
Description:

	This function converts from an interface to a KVector object.
*/
func From(expr interface{}) symbolic.KVector {
	switch converted := expr.(type) {
	case symbolic.KVector:
		return converted
	case *mat.VecDense:
		// Return
		return symbolic.VecDenseToKVector(*converted)
	case mat.VecDense:
		// Return
		return symbolic.VecDenseToKVector(converted)
	case []float64:
		// Transform float64 into VecDense and then use mat.VecDense
		// constructor.
		newVD := mat.NewVecDense(len(converted), converted)
		return symbolic.VecDenseToKVector(*newVD)
	default:
		panic(
			smErrors.UnsupportedInputError{
				FunctionName: "KVector.From",
				Input:        converted,
			},
		)
	}
}
