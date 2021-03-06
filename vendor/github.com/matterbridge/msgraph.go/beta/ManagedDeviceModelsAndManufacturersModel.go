// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceModelsAndManufacturers undocumented
type ManagedDeviceModelsAndManufacturers struct {
	// Object is the base model of ManagedDeviceModelsAndManufacturers
	Object
	// DeviceModels List of Models for managed devices in the account
	DeviceModels []string `json:"deviceModels,omitempty"`
	// DeviceManufacturers List of Manufactures for managed devices in the account
	DeviceManufacturers []string `json:"deviceManufacturers,omitempty"`
}
