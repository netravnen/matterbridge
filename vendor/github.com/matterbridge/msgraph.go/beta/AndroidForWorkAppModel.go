// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkApp Contains properties and inherited properties for Android for Work (AFW) Apps.
type AndroidForWorkApp struct {
	// MobileApp is the base model of AndroidForWorkApp
	MobileApp
	// PackageID The package identifier.
	PackageID *string `json:"packageId,omitempty"`
	// AppIdentifier The Identity Name.
	AppIdentifier *string `json:"appIdentifier,omitempty"`
	// UsedLicenseCount The number of VPP licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// TotalLicenseCount The total number of VPP licenses.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// AppStoreURL The Play for Work Store app URL.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
}
