// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserRegistrationCount undocumented
type UserRegistrationCount struct {
	// Object is the base model of UserRegistrationCount
	Object
	// RegistrationStatus undocumented
	RegistrationStatus *RegistrationStatusType `json:"registrationStatus,omitempty"`
	// RegistrationCount undocumented
	RegistrationCount *int `json:"registrationCount,omitempty"`
}
