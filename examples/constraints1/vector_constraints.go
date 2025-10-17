package main

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Create the variables
	N := 2
	x := symbolic.NewVariableVector(N)

	// Create the constraints
	// - x >= -1
	c1 := x.GreaterEq(mat.NewVecDense(N, []float64{-1, -1}))

	// - x <= 1
	c2 := x.LessEq(symbolic.OnesVector(N))

	// Print the constraints
	fmt.Println("Constraint 1:", c1)
	fmt.Println("Constraint 2:", c2)
}
