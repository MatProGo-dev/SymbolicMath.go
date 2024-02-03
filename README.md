[![codecov](https://codecov.io/gh/MatProGo-dev/SymbolicMath.go/graph/badge.svg?token=CO7oq7ZZ9l)](https://codecov.io/gh/MatProGo-dev/SymbolicMath.go)

# SymbolicMath.go
A symbolic math module for the Go (Golang) Programming Language.

## Installation
```bash
go get github.com/MatProGo-dev/SymbolicMath.go
```

## Usage
```go
package main

import (
    "fmt"
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
    "github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

func main() {
	// Constants
	N := 2
	x := symbolic.NewVariableVector(N)
	Q := getKMatrix.From(
		[][]float64{
			{1, 0},
			{0, 2.0},
		})
	
	// Create the quadratic polynomial
	quadPoly := x.Transpose().Multiply(Q).Multiply(x)

	// Print the polynomial
	fmt.Println(quadPoly.String())
	
	/* Other stuff... */
}

```

Further examples can be found in the `examples` directory.