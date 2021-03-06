// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DepEnrollmentBaseProfile The DepEnrollmentBaseProfile resource represents an Apple Device Enrollment Program (DEP) enrollment profile. This type of profile must be assigned to Apple DEP serial numbers before the corresponding devices can enroll via DEP.
type DepEnrollmentBaseProfile struct {
	// EnrollmentProfile is the base model of DepEnrollmentBaseProfile
	EnrollmentProfile
	// IsDefault Indicates if this is the default profile
	IsDefault *bool `json:"isDefault,omitempty"`
	// SupervisedModeEnabled Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/en-us/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
	SupervisedModeEnabled *bool `json:"supervisedModeEnabled,omitempty"`
	// SupportDepartment Support department information
	SupportDepartment *string `json:"supportDepartment,omitempty"`
	// PassCodeDisabled Indicates if Passcode setup pane is disabled
	PassCodeDisabled *bool `json:"passCodeDisabled,omitempty"`
	// IsMandatory Indicates if the profile is mandatory
	IsMandatory *bool `json:"isMandatory,omitempty"`
	// LocationDisabled Indicates if Location service setup pane is disabled
	LocationDisabled *bool `json:"locationDisabled,omitempty"`
	// SupportPhoneNumber Support phone number
	SupportPhoneNumber *string `json:"supportPhoneNumber,omitempty"`
	// ProfileRemovalDisabled Indicates if the profile removal option is disabled
	ProfileRemovalDisabled *bool `json:"profileRemovalDisabled,omitempty"`
	// RestoreBlocked Indicates if Restore setup pane is blocked
	RestoreBlocked *bool `json:"restoreBlocked,omitempty"`
	// AppleIDDisabled Indicates if Apple id setup pane is disabled
	AppleIDDisabled *bool `json:"appleIdDisabled,omitempty"`
	// TermsAndConditionsDisabled Indicates if 'Terms and Conditions' setup pane is disabled
	TermsAndConditionsDisabled *bool `json:"termsAndConditionsDisabled,omitempty"`
	// TouchIDDisabled Indicates if touch id setup pane is disabled
	TouchIDDisabled *bool `json:"touchIdDisabled,omitempty"`
	// ApplePayDisabled Indicates if Apple pay setup pane is disabled
	ApplePayDisabled *bool `json:"applePayDisabled,omitempty"`
	// ZoomDisabled Indicates if zoom setup pane is disabled
	ZoomDisabled *bool `json:"zoomDisabled,omitempty"`
	// SiriDisabled Indicates if siri setup pane is disabled
	SiriDisabled *bool `json:"siriDisabled,omitempty"`
	// DiagnosticsDisabled Indicates if diagnostics setup pane is disabled
	DiagnosticsDisabled *bool `json:"diagnosticsDisabled,omitempty"`
	// DisplayToneSetupDisabled Indicates if displaytone setup screen is disabled
	DisplayToneSetupDisabled *bool `json:"displayToneSetupDisabled,omitempty"`
	// PrivacyPaneDisabled Indicates if privacy screen is disabled
	PrivacyPaneDisabled *bool `json:"privacyPaneDisabled,omitempty"`
	// ScreenTimeScreenDisabled Indicates if screen timeout setup is disabled
	ScreenTimeScreenDisabled *bool `json:"screenTimeScreenDisabled,omitempty"`
	// DeviceNameTemplate Sets a literal or name pattern.
	DeviceNameTemplate *string `json:"deviceNameTemplate,omitempty"`
	// ConfigurationWebURL URL for setup assistant login
	ConfigurationWebURL *bool `json:"configurationWebUrl,omitempty"`
}
