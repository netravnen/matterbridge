// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// UpdateWindow undocumented
type UpdateWindow struct {
	// Object is the base model of UpdateWindow
	Object
	// UpdateWindowStartTime undocumented
	UpdateWindowStartTime *time.Time `json:"updateWindowStartTime,omitempty"`
	// UpdateWindowEndTime undocumented
	UpdateWindowEndTime *time.Time `json:"updateWindowEndTime,omitempty"`
}
