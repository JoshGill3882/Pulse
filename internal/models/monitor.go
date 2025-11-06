package models

// Imports
import "time"

// Definition of the "Monitor" type structure
type Monitor struct {
	// ID (UUID)
	ID 					string				`json:"id"`
	// Name (e.g. Portfolio Website)
	Name 				string				`json:"name"`

	// URL (e.g. "https://www.joshgill.dev")
	URL 				string				`json:"url"`
	// Method (e.g. "GET")
	Method 				string				`json:"method"`
	// Headers in a Map Format (e.g. {"Authentication": "Bearer: TOKEN"})
	HeadersJson 		map[string]string	`json:"headers"`
	// Interval between health check requests in seconds
	IntervalSec 		int					`json:"interval"`
	// List of acceptable HTTP Status Codes for the check to return to define an "Up" status
	AcceptableHttpCodes []int				`json:"acceptableCodes"`
	
	// Datetime when the Monitor was created
	CreatedAt 			time.Time			`json:"createdAt"`
	// Datetime when the Monitor was most recently updated
	UpdatedAt 			time.Time			`json:"updatedAt"`

	// Boolean of whether the Monitor is enabled
	Enabled 			bool				`json:"enabled"`
}
