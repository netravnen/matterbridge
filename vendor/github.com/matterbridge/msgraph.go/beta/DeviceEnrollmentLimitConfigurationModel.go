// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceEnrollmentLimitConfiguration Device Enrollment Configuration that restricts the number of devices a user can enroll
type DeviceEnrollmentLimitConfiguration struct {
	// DeviceEnrollmentConfiguration is the base model of DeviceEnrollmentLimitConfiguration
	DeviceEnrollmentConfiguration
	// Limit The maximum number of devices that a user can enroll
	Limit *int `json:"limit,omitempty"`
}
