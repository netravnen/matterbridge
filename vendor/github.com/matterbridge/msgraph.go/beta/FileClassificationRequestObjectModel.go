// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FileClassificationRequestObject undocumented
type FileClassificationRequestObject struct {
	// Entity is the base model of FileClassificationRequestObject
	Entity
	// File undocumented
	File *Stream `json:"file,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}
