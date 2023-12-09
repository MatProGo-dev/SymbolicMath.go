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
    "github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

func main() {
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()
	
	sum := v1.Plus(v2)
	
	/* Other stuff... */
}

```