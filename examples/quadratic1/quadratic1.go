package main

import (
	"fmt"
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
quadratic1.go
Description:

	This script is meant to construct a quadratic polynomial
	using the symbolic package.
*/

func main() {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	Q := getKMatrix.From(
		[][]float64{
			{1, 0},
			{0, 2.0},
		})

	fmt.Println(x.Transpose().Multiply(Q))

	// Create the quadratic polynomial
	quadPoly := x.Transpose().Multiply(Q).Multiply(x)

	// Print the polynomial
	fmt.Println(quadPoly.String())
}
