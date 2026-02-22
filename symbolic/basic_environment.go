package symbolic

// BasicEnvironment ...
type BasicEnvironment struct {
	name      string
	Variables []Variable
}

// GetName ...
func (be *BasicEnvironment) GetName() string {
	return be.name
}

// TrackVariable ...
func (be *BasicEnvironment) TrackVariable(v Variable) bool {
	// Check if the variable is already in the environment
	for _, existingVar := range be.Variables {
		if existingVar.ID == v.ID {
			// Variable already exists, do not add it again
			return false
		}
	}

	// Add the variable to the environment
	be.Variables = append(be.Variables, v)
	return true // Variable was added successfully
}

// AllTrackedVariables ...
func (be *BasicEnvironment) AllTrackedVariables() []Variable {
	return be.Variables
}

// MakeBasicEnvironment creates a new BasicEnvironment with the given name.
func MakeBasicEnvironment(nameIn string) BasicEnvironment {
	return BasicEnvironment{
		name:      nameIn,
		Variables: []Variable{},
	}
}

// DefaultEnvironment A variable that exists in the background and used to store information about the variables currently created.
var DefaultEnvironment = MakeBasicEnvironment("DefaultEnvironment")
