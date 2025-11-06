package models

// Imports
import "time"

// Definition of the "Result" type structure
type Result struct {
	// ID of the Result (will auto-increment within Postgres)
	ID 				int
	// ID of the Monitor the result relates to
	MonitorID 		string
	
	// Status of the result (Up or Down)
	ResultStatus	Status
	// HTTP status of the result
	HttpStatus 		int
	// Latency of the call in milliseconds
	Latency 		int
	// Error provided (if relevant)
	Error 			string

	// Datetime for when the result was made
	CheckedAt 		time.Time
}
