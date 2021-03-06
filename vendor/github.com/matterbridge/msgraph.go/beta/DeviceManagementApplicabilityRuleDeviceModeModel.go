// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementApplicabilityRuleDeviceMode undocumented
type DeviceManagementApplicabilityRuleDeviceMode struct {
	// Object is the base model of DeviceManagementApplicabilityRuleDeviceMode
	Object
	// DeviceMode Applicability rule for device mode.
	DeviceMode *Windows10DeviceModeType `json:"deviceMode,omitempty"`
	// Name Name for object.
	Name *string `json:"name,omitempty"`
	// RuleType Applicability Rule type.
	RuleType *DeviceManagementApplicabilityRuleType `json:"ruleType,omitempty"`
}
