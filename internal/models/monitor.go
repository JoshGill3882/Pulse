package models

// Imports
import "time"

// Definition of the "Monitor" type structure
type Monitor struct {
	// ID (UUID)
	ID 					string
	// Name (e.g. Portfolio Website)
	Name 				string

	// URL (e.g. "https://www.joshgill.dev")
	URL 				string
	// Method (e.g. "GET")
	Method 				string
	// Headers in a Map Format (e.g. {"Authentication": "Bearer: TOKEN"})
	HeadersJson 		map[string]string
	// Interval between health check requests in seconds
	IntervalSec 		int
	// List of acceptable HTTP Status Codes for the check to return to define an "Up" status
	AcceptableHttpCodes []int
	
	// Datetime when the Monitor was created
	CreatedAt 			time.Time
	// Datetime when the Monitor was most recently updated
	UpdatedAt 			time.Time

	// Boolean of whether the Monitor is enabled
	Enabled 			bool
}
