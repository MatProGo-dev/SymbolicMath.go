package main

import (
	"fmt"

	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
substitute1.go
Description:

	This script is meant to show how the substitution
	method can be used to evaluate the value of your
	expressions. In this case, we will use it to evaluate
	a quadratic expression at the point [-3.0, 1.0].

	The result should be a value of 10.0.
*/

func main() {
	// Construct quadratic expression
	N := 2
	x := symbolic.NewVariableVector(N)
	Q := getKMatrix.From(
		symbolic.Identity(N))

	// Create the quadratic polynomial
	quadPoly := x.Transpose().Multiply(Q).Multiply(x)
	fmt.Println("Polynomial is:", quadPoly)

	// Create the "map" used for substitution
	x0 := x[0]
	x1 := x[1]

	subMap := map[symbolic.Variable]symbolic.Expression{
		x0: symbolic.K(-3.0),
		x1: symbolic.K(1.0),
	}

	quadEvaluated := quadPoly.SubstituteAccordingTo(subMap)

	// Print the polynomial
	fmt.Println(
		"Polynomial at the point x = [",
		subMap[x0],
		",",
		subMap[x1],
		"] is",
		quadEvaluated.String())
}
