// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ImportedAppleDeviceIdentityResult The importedAppleDeviceIdentityResult resource represents the result of attempting to import Apple devices identities.
type ImportedAppleDeviceIdentityResult struct {
	// ImportedAppleDeviceIdentity is the base model of ImportedAppleDeviceIdentityResult
	ImportedAppleDeviceIdentity
	// Status Status of imported device identity
	Status *bool `json:"status,omitempty"`
}
