package symbolic

/*
A simple implementation of the Environment interface.
This is used to track variables, nothing more and nothing less.

The "Background" environment (used automatically when none is specified) is
a BasicEnvironment.
*/
type BasicEnvironment struct {
	// The name of this environment.
	name string

	// The variables currently tracked in this environment.
	Variables []Variable
}

/*
Returns the name of the environment

Note: This is required to implement the Environment interface.
*/
func (be *BasicEnvironment) GetName() string {
	return be.name
}

/*
Tracks a variable in the environment.

Returns true if the variable was added, false if it was already present.

Note: This is required to implement the Environment interface.
*/
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

/*
Returns all variables currently tracked in the environment.

Note: This is required to implement the Environment interface.
*/
func (be *BasicEnvironment) AllTrackedVariables() []Variable {
	return be.Variables
}

/*
Public Functions
*/

/*
Defines an empty, new BasicEnvironment with the given name.
*/
func MakeBasicEnvironment(nameIn string) BasicEnvironment {
	return BasicEnvironment{
		name:      nameIn,
		Variables: []Variable{},
	}
}

/*
This variable that exists in the background and is used to store information about the variables currently created
when no environment is specified.
*/
var DefaultEnvironment = MakeBasicEnvironment("DefaultEnvironment")
