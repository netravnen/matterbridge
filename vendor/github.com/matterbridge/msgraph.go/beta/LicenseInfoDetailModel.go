// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LicenseInfoDetail undocumented
type LicenseInfoDetail struct {
	// Object is the base model of LicenseInfoDetail
	Object
	// LicenseType undocumented
	LicenseType *AzureADLicenseType `json:"licenseType,omitempty"`
	// TotalLicenseCount undocumented
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// TotalAssignedCount undocumented
	TotalAssignedCount *int `json:"totalAssignedCount,omitempty"`
	// TotalUsageCount undocumented
	TotalUsageCount *int `json:"totalUsageCount,omitempty"`
}
