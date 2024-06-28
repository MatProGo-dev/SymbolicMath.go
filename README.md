[![Go Reference](https://pkg.go.dev/badge/github.com/MatProGo-dev/SymbolicMath.go.svg)](https://pkg.go.dev/github.com/MatProGo-dev/SymbolicMath.go)
[![codecov](https://codecov.io/gh/MatProGo-dev/SymbolicMath.go/graph/badge.svg?token=CO7oq7ZZ9l)](https://codecov.io/gh/MatProGo-dev/SymbolicMath.go)

# SymbolicMath.go
A symbolic math module for the Go (Golang) Programming Language.
It features a simple, composable API for performing
math operations on symbolic expressions and is built to
express everything from scalar constants to matrices of polynomials.

Some key features include:
- Composable operations which allow complex mathematics to be expressed
  in as few lines as you wish (e.g., `x.Transpose().Multiply(Q).Multiply(x)`)
- Simple API for defining constants and variables (e.g., `Q := getKMatrix.From([][]float64{{1, 0}, {0, 2.0}})`)
- Recognition of [gonum](https://www.gonum.org/) matrices and vectors
  in most operations (e.g., `vv1.Plus(mat.NewVecDense(N, []float64{1, 2, 3}))`)
- Producing panics instead of errors (i.e., if something strange is written
in symbolic math, like the multiplication of two mismatched matrices, then the
program will panic instead of returning an error. This way, you can worry
about the math and not the error handling.)

Some documentation can be found by clicking the "reference" badge above.

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

## Related Projects

This project was motivated by the need for a symbolic math package for defining
optimization and control theory problems in Go, but symbolic mathematics is a topic that covers
a wide range of applications. If this tool is not useful for your purpose, then you might
find one of the following projects more helpful:

While other symbolic math libraries exist for Go, they typically focus on:
- Computer Algebra Systems that will help you get a final expression (in text) from
arbitrary math input (often in text) \[[expreduce](https://github.com/corywalker/expreduce),[sm](https://github.com/Konstantin8105/sm)\]
- Implementing Algorithms from [Domain-Specific Languages of Mathematics](https://github.com/DSLsofMath/DSLsofMath)
  Course \[[gosymbol](https://github.com/victorbrun/gosymbol/tree/main)\]
- Machine Learning and algorithms needed to perform automatic differentiation \[[gorgonia](https://github.com/gorgonia/gorgonia)\]