package models

// Imports
import "time"

// Definition of the "Result" type structure
type Result struct {
	// ID of the Result (will auto-increment within Postgres)
	ID 				int				`json:"id"`
	// ID of the Monitor the result relates to
	MonitorID 		string			`json:"monitor"`
	
	// Status of the result (Up or Down)
	ResultStatus	Status			`json:"resultStatus"`
	// HTTP status of the result
	HttpStatus 		int				`json:"httpStatus"`
	// Latency of the call in milliseconds
	Latency 		int				`json:"latency"`
	// Error provided (if relevant)
	Error 			string			`json:"error"`

	// Datetime for when the result was made
	CheckedAt 		time.Time		`json:"checkedAt"`
}
