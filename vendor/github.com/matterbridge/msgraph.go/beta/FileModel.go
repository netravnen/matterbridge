// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// File undocumented
type File struct {
	// Object is the base model of File
	Object
	// Hashes undocumented
	Hashes *Hashes `json:"hashes,omitempty"`
	// MimeType undocumented
	MimeType *string `json:"mimeType,omitempty"`
	// ProcessingMetadata undocumented
	ProcessingMetadata *bool `json:"processingMetadata,omitempty"`
}
