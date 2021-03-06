// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PositionDetail undocumented
type PositionDetail struct {
	// Object is the base model of PositionDetail
	Object
	// Company undocumented
	Company *CompanyDetail `json:"company,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// EndMonthYear undocumented
	EndMonthYear *time.Time `json:"endMonthYear,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// Role undocumented
	Role *string `json:"role,omitempty"`
	// StartMonthYear undocumented
	StartMonthYear *time.Time `json:"startMonthYear,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
}
