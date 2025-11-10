package models

// Imports
import "time"

// Defining a "Document" structure type
type Document struct {
	// The ID of the MongoDB Document
	ID 			string				`json:"id"`
	
	// The Database the Document is a part of
	Database	string				`json:"database"`
	// The Collection the Document is a part of
	Collection	string				`json:"collection"`
	// The Content of the MongoDB Document
	Content 	map[string]any		`json:"content"`

	// The timestamp the Document was retrieved at
	Retrieved	time.Time			`json:"retrieved"`
}
