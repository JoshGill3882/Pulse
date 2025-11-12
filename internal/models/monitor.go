package models

// Imports
import "time"

// Definition of the "Monitor" type structure
type Monitor struct {
	// ID (UUID)
	ID 					string				`bson:"_id,omitempty" json:"id"`
	// Name (e.g. Portfolio Website)
	Name 				string				`bson:"name" json:"name"`

	// URL (e.g. "https://www.joshgill.dev")
	URL 				string				`bson:"url" json:"url"`
	// Method (e.g. "GET")
	Method 				string				`bson:"method" json:"method"`
	// Headers in a Map Format (e.g. {"Authentication": "Bearer: TOKEN"})
	HeadersJson 		map[string]string	`bson:"headers" json:"headers"`
	// Interval between health check requests in seconds
	IntervalSec 		int					`bson:"interval" json:"interval"`
	// List of acceptable HTTP Status Codes for the check to return to define an "Up" status
	AcceptableHttpCodes []int				`bson:"acceptableCodes" json:"acceptableCodes"`
	
	// Datetime when the Monitor was created
	CreatedAt 			time.Time			`bson:"createdAt" json:"createdAt"`
	// Datetime when the Monitor was most recently updated
	UpdatedAt 			time.Time			`bson:"updatedAt" json:"updatedAt"`
	// Datetime when the Monitor was most recently retrieved
	RetrievedAt			time.Time			`bson:"retrievedAt" json:"retrievedAt"`

	// Boolean of whether the Monitor is enabled
	Enabled 			bool				`bson:"enabled" json:"enabled"`
}
