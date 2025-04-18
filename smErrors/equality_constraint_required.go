package smErrors

type EqualityConstraintRequiredError struct {
	Operation string
}

func (e EqualityConstraintRequiredError) Error() string {
	return "Equality constraint required for operation: " + e.Operation
}
