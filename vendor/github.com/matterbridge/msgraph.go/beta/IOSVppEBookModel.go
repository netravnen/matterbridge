// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSVppEBook A class containing the properties for iOS Vpp eBook.
type IOSVppEBook struct {
	// ManagedEBook is the base model of IOSVppEBook
	ManagedEBook
	// VppTokenID The Vpp token ID.
	VppTokenID *UUID `json:"vppTokenId,omitempty"`
	// AppleID The Apple ID associated with Vpp token.
	AppleID *string `json:"appleId,omitempty"`
	// VppOrganizationName The Vpp token's organization name.
	VppOrganizationName *string `json:"vppOrganizationName,omitempty"`
	// Genres Genres.
	Genres []string `json:"genres,omitempty"`
	// Language Language.
	Language *string `json:"language,omitempty"`
	// Seller Seller.
	Seller *string `json:"seller,omitempty"`
	// TotalLicenseCount Total license count.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// UsedLicenseCount Used license count.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this Entity instance.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
}
