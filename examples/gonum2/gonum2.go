package main

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
gonum2.go
Description:

	This script is meant to demonstrate the use of the gonum package's
	matrices in an operation with the symbolic package.
*/
func main() {
	// Constants
	N := 3
	vv1 := symbolic.NewVariableVector(N)
	Q := mat.NewDense(N, N, []float64{1, 0, 0, 0, 2, 0, 0, 0, 3})

	// Sum vv1 with a mat.VecDense object from the gonum package
	sum2 := vv1.Transpose().Multiply(Q)

	// Print the result
	fmt.Println(sum2.String())
}
