package main

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
quadratic1.go
Description:

	This script is meant to construct a quadratic polynomial
	using the symbolic package.
*/

func main() {
	// Setup
	t := symbolic.NewVariable()
	t.Name = "t"

	g := 9.81 // m/s^2
	v0 := 3.0 // m/2^2

	// Create quadratic function
	yPosition := symbolic.K(-0.5 * g).Multiply(
		t.Power(2),
	).Plus(
		t.Multiply(v0),
	)

	// Create the derivative
	yVelocity := yPosition.DerivativeWrt(t)

	// Print the polynomial
	fmt.Println(yVelocity.String())
}
