package smErrors

import "fmt"

/*
negative_exponent.go
Description:

	Functions related to the negative exponent error.
*/

// Type Definition
type NegativeExponentError struct {
	Exponent int
}

// Error
func (e NegativeExponentError) Error() string {
	return fmt.Sprintf("received negative exponent (%v); expected non-negative exponent", e.Exponent)
}
