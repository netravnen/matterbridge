// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IOSLobAppProvisioningConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the iOS Lob App Provisioning Configuration resource.
type IOSLobAppProvisioningConfiguration struct {
	// Entity is the base model of IOSLobAppProvisioningConfiguration
	Entity
	// ExpirationDateTime Optional profile expiration date and time.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// PayloadFileName Payload file name (*.mobileprovision | *.xml).
	PayloadFileName *string `json:"payloadFileName,omitempty"`
	// Payload Payload. (UTF8 encoded byte array)
	Payload *Binary `json:"payload,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this iOS LOB app provisioning configuration entity.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// CreatedDateTime DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description Admin provided description of the Device Configuration.
	Description *string `json:"description,omitempty"`
	// LastModifiedDateTime DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// DisplayName Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`
	// Version Version of the device configuration.
	Version *int `json:"version,omitempty"`
	// GroupAssignments undocumented
	GroupAssignments []MobileAppProvisioningConfigGroupAssignment `json:"groupAssignments,omitempty"`
	// Assignments undocumented
	Assignments []IOSLobAppProvisioningConfigurationAssignment `json:"assignments,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	// UserStatuses undocumented
	UserStatuses []ManagedDeviceMobileAppConfigurationUserStatus `json:"userStatuses,omitempty"`
}
