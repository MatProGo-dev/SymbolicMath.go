package symbolic

/*
environment.go
Description:
	Defines the environment where the symbolic variables are stored.
*/

type Environment struct {
	Name      string
	Variables []Variable
}

/*
Public Functions
*/

/*
BackgroundEnvironment
A variable that exists in the background and used to store information about the variables currently created.
*/
var BackgroundEnvironment = Environment{
	Name:      "Background",
	Variables: []Variable{},
}
