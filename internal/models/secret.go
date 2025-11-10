package models

import (
	"time"
)

// Defining the "Secret" type
type Secret struct {
	// Secret Key
	Key 			string			`json:"key"`
	// Project ID
	ProjectID 		string			`json:"projectId"`
	// Environment
	Environment 	string			`json:"environment"`

	// Secret Value
	Value 			string			`json:"value"`
	// Secret Version
	Version 		int				`json:"version"`

	// Comment attached to the Secret
	Comment 		string			`json:"comment"`

	// Time when the Secret was Retrieved
	Retrieved 		time.Time		`json:"retrieved"`
}
