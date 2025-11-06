package models

// Defining the equivalent of an "enum" in Go for the different types of possible "Status"
// Reference used - https://gobyexample.com/enums

// Defining a "status" type as an integer
type Status int

// Defining the possible values of "Status" using a constant and the "iota" keyword
// "iota" generates "successive constant values automatically" (e.g. 0 and 1 in this case)
const (
	Up Status = iota
	Down
)

// Mapping the different options to string values
var statusName = map[Status]string{
	Up:		"Up",
	Down:	"Down",
}

// Creating a Function which returns the string value for a Status when given the Status
func (s Status) String() string { return statusName[s] }
