package models

// Imports
import "time"

// Definition of the "Result" type structure
type Result struct {
	// ID of the Result (will auto-increment within Postgres)
	ID 				int				`bson:"id,omitempty" json:"id"`
	// ID of the Monitor the result relates to
	MonitorID 		string			`bson:"monitorId" json:"monitor"`
	
	// Status of the result (Up or Down)
	ResultStatus	Status			`bson:"resultStatus" json:"resultStatus"`
	// HTTP status of the result
	HttpStatus 		int				`bson:"httpStatus" json:"httpStatus"`
	// Latency of the call in milliseconds
	Latency 		int				`bson:"latency" json:"latency"`
	// Error provided (if relevant)
	Error 			string			`bson:"error" json:"error"`

	// Datetime for when the result was made
	CheckedAt 		time.Time		`bson:"checkedAt" json:"checkedAt"`
}
