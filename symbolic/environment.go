package symbolic

/*
environment.go
Description:
	Defines the environment where the symbolic variables are stored.
*/

/*
This interface defines the methods required for an object that tracks symbolic variables.

We plan for several implementations of this interface, including:
- BasicEnvironment: A simple environment that tracks variables in a list.
- OptimizationProblem: An optimization problem that saves variables that have been added in constraints and/or objectives.
*/
type Environment interface {
	// Returns the name of the environment.
	GetName() string

	// Adds a variable to the environment if it is not already present.
	// Returns true if the variable was added, false if adding the variable failed for any reason.
	TrackVariable(v Variable) bool

	// Returns all variables currrently tracked in the environment.
	AllTrackedVariables() []Variable
}
