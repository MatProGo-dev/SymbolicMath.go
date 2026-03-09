package symbolic

// Environment defines the environment where the symbolic variables are stored.
type Environment interface {
	GetName() string
	TrackVariable(v Variable) bool
	AllTrackedVariables() []Variable
}
