package main

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
gonum1.go
Description:

	This script is meant to demonstrate the use of the gonum package
	within some operations of the symbolic package.
*/

func main() {
	// Constants
	N := 3
	vv1 := symbolic.NewVariableVector(N)

	// Sum vv1 with a mat.VecDense object from the gonum package
	sum2 := vv1.Plus(mat.NewVecDense(N, []float64{1, 2, 3}))

	// Print the result
	fmt.Println(sum2.String())
}
