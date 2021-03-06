// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerCompliancePolicy This topic provides descriptions of the declared methods, properties and relationships exposed by the AndroidDeviceOwnerCompliancePolicy resource.
type AndroidDeviceOwnerCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of AndroidDeviceOwnerCompliancePolicy
	DeviceCompliancePolicy
	// DeviceThreatProtectionEnabled Require that devices have enabled device threat protection.
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel Require Mobile Threat Protection minimum risk level to report noncompliance.
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// SecurityRequireSafetyNetAttestationBasicIntegrity Require the device to pass the SafetyNet basic integrity check.
	SecurityRequireSafetyNetAttestationBasicIntegrity *bool `json:"securityRequireSafetyNetAttestationBasicIntegrity,omitempty"`
	// SecurityRequireSafetyNetAttestationCertifiedDevice Require the device to pass the SafetyNet certified device check.
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool `json:"securityRequireSafetyNetAttestationCertifiedDevice,omitempty"`
	// OsMinimumVersion Minimum Android version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// OsMaximumVersion Maximum Android version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// MinAndroidSecurityPatchLevel Minimum Android security patch level.
	MinAndroidSecurityPatchLevel *string `json:"minAndroidSecurityPatchLevel,omitempty"`
	// PasswordRequired Require a password to unlock device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordMinimumLength Minimum password length. Valid values 4 to 16
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinimumLetterCharacters Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
	PasswordMinimumLetterCharacters *int `json:"passwordMinimumLetterCharacters,omitempty"`
	// PasswordMinimumLowerCaseCharacters Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
	PasswordMinimumLowerCaseCharacters *int `json:"passwordMinimumLowerCaseCharacters,omitempty"`
	// PasswordMinimumNonLetterCharacters Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
	PasswordMinimumNonLetterCharacters *int `json:"passwordMinimumNonLetterCharacters,omitempty"`
	// PasswordMinimumNumericCharacters Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
	PasswordMinimumNumericCharacters *int `json:"passwordMinimumNumericCharacters,omitempty"`
	// PasswordMinimumSymbolCharacters Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
	PasswordMinimumSymbolCharacters *int `json:"passwordMinimumSymbolCharacters,omitempty"`
	// PasswordMinimumUpperCaseCharacters Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
	PasswordMinimumUpperCaseCharacters *int `json:"passwordMinimumUpperCaseCharacters,omitempty"`
	// PasswordRequiredType Type of characters in password
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordExpirationDays Number of days before the password expires. Valid values 1 to 365
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordPreviousPasswordCountToBlock Number of previous passwords to block. Valid values 1 to 24
	PasswordPreviousPasswordCountToBlock *int `json:"passwordPreviousPasswordCountToBlock,omitempty"`
	// StorageRequireEncryption Require encryption on Android devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}
