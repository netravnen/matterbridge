// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CompanyDetail undocumented
type CompanyDetail struct {
	// Object is the base model of CompanyDetail
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Pronunciation undocumented
	Pronunciation *string `json:"pronunciation,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
