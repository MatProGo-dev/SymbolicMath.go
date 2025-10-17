package symbolic

/*
environment.go
Description:
	Defines the environment where the symbolic variables are stored.
*/

type Environment interface {
	GetName() string
	TrackVariable(v Variable)
	AllTrackedVariables() []Variable
}
