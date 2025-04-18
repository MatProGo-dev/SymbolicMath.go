package smErrors

type InequalityConstraintRequiredError struct {
	Operation string
}

func (icre InequalityConstraintRequiredError) Error() string {
	return "Inequality constraint required for operation: " + icre.Operation
}
