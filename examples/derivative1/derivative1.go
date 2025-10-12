package main

import (
	"fmt"

	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
derivative1.go
Description:

	This script is meant to construct a simple polynomial
	representing an object in free fall within standard
	earth gravity, `g`, with an initial upward velocity of `v0`.

	By differentiating the position equation, you should get the
	velocity equation.
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
