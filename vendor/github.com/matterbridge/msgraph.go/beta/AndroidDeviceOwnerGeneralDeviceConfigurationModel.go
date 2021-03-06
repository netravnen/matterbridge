// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerGeneralDeviceConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the androidDeviceOwnerGeneralDeviceConfiguration resource.
type AndroidDeviceOwnerGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of AndroidDeviceOwnerGeneralDeviceConfiguration
	DeviceConfiguration
	// AccountsBlockModification Indicates whether or not adding or removing accounts is disabled.
	AccountsBlockModification *bool `json:"accountsBlockModification,omitempty"`
	// AppsAllowInstallFromUnknownSources Indicates whether or not the user is allowed to enable to unknown sources setting.
	AppsAllowInstallFromUnknownSources *bool `json:"appsAllowInstallFromUnknownSources,omitempty"`
	// AppsAutoUpdatePolicy Indicates the value of the app auto update policy.
	AppsAutoUpdatePolicy *AndroidDeviceOwnerAppAutoUpdatePolicyType `json:"appsAutoUpdatePolicy,omitempty"`
	// AppsDefaultPermissionPolicy Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically.
	AppsDefaultPermissionPolicy *AndroidDeviceOwnerDefaultAppPermissionPolicyType `json:"appsDefaultPermissionPolicy,omitempty"`
	// AppsRecommendSkippingFirstUseHints Whether or not to recommend all apps skip any first-time-use hints they may have added.
	AppsRecommendSkippingFirstUseHints *bool `json:"appsRecommendSkippingFirstUseHints,omitempty"`
	// BluetoothBlockConfiguration Indicates whether or not to block a user from configuring bluetooth.
	BluetoothBlockConfiguration *bool `json:"bluetoothBlockConfiguration,omitempty"`
	// BluetoothBlockContactSharing Indicates whether or not to block a user from sharing contacts via bluetooth.
	BluetoothBlockContactSharing *bool `json:"bluetoothBlockContactSharing,omitempty"`
	// CameraBlocked Indicates whether or not to disable the use of the camera.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`
	// CellularBlockWiFiTethering Indicates whether or not to block Wi-Fi tethering.
	CellularBlockWiFiTethering *bool `json:"cellularBlockWiFiTethering,omitempty"`
	// DataRoamingBlocked Indicates whether or not to block a user from data roaming.
	DataRoamingBlocked *bool `json:"dataRoamingBlocked,omitempty"`
	// DateTimeConfigurationBlocked Indicates whether or not to block the user from manually changing the date or time on the device
	DateTimeConfigurationBlocked *bool `json:"dateTimeConfigurationBlocked,omitempty"`
	// FactoryResetDeviceAdministratorEmails List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
	FactoryResetDeviceAdministratorEmails []string `json:"factoryResetDeviceAdministratorEmails,omitempty"`
	// FactoryResetBlocked Indicates whether or not the factory reset option in settings is disabled.
	FactoryResetBlocked *bool `json:"factoryResetBlocked,omitempty"`
	// GlobalProxy Proxy is set up directly with host, port and excluded hosts.
	GlobalProxy *AndroidDeviceOwnerGlobalProxy `json:"globalProxy,omitempty"`
	// GoogleAccountsBlocked Indicates whether or not google accounts will be blocked.
	GoogleAccountsBlocked *bool `json:"googleAccountsBlocked,omitempty"`
	// KioskModeScreenSaverConfigurationEnabled Whether or not to enable screen saver mode or not in Kiosk Mode.
	KioskModeScreenSaverConfigurationEnabled *bool `json:"kioskModeScreenSaverConfigurationEnabled,omitempty"`
	// KioskModeScreenSaverImageURL URL for an image that will be the device's screen saver in Kiosk Mode.
	KioskModeScreenSaverImageURL *string `json:"kioskModeScreenSaverImageUrl,omitempty"`
	// KioskModeScreenSaverDisplayTimeInSeconds The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
	KioskModeScreenSaverDisplayTimeInSeconds *int `json:"kioskModeScreenSaverDisplayTimeInSeconds,omitempty"`
	// KioskModeScreenSaverStartDelayInSeconds The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
	KioskModeScreenSaverStartDelayInSeconds *int `json:"kioskModeScreenSaverStartDelayInSeconds,omitempty"`
	// KioskModeScreenSaverDetectMediaDisabled Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
	KioskModeScreenSaverDetectMediaDisabled *bool `json:"kioskModeScreenSaverDetectMediaDisabled,omitempty"`
	// KioskModeApps A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
	KioskModeApps []AppListItem `json:"kioskModeApps,omitempty"`
	// KioskModeWallpaperURL URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
	KioskModeWallpaperURL *string `json:"kioskModeWallpaperUrl,omitempty"`
	// KioskModeExitCode Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
	KioskModeExitCode *string `json:"kioskModeExitCode,omitempty"`
	// KioskModeVirtualHomeButtonEnabled Whether or not to display a virtual home button when the device is in Kiosk Mode.
	KioskModeVirtualHomeButtonEnabled *bool `json:"kioskModeVirtualHomeButtonEnabled,omitempty"`
	// KioskModeVirtualHomeButtonType Indicates whether the virtual home button is a swipe up home button or a floating home button.
	KioskModeVirtualHomeButtonType *AndroidDeviceOwnerVirtualHomeButtonType `json:"kioskModeVirtualHomeButtonType,omitempty"`
	// KioskModeBluetoothConfigurationEnabled Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
	KioskModeBluetoothConfigurationEnabled *bool `json:"kioskModeBluetoothConfigurationEnabled,omitempty"`
	// KioskModeWiFiConfigurationEnabled Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
	KioskModeWiFiConfigurationEnabled *bool `json:"kioskModeWiFiConfigurationEnabled,omitempty"`
	// KioskModeFlashlightConfigurationEnabled Whether or not to allow a user to use the flashlight in Kiosk Mode.
	KioskModeFlashlightConfigurationEnabled *bool `json:"kioskModeFlashlightConfigurationEnabled,omitempty"`
	// KioskModeMediaVolumeConfigurationEnabled Whether or not to allow a user to change the media volume in Kiosk Mode.
	KioskModeMediaVolumeConfigurationEnabled *bool `json:"kioskModeMediaVolumeConfigurationEnabled,omitempty"`
	// MicrophoneForceMute Indicates whether or not to block unmuting the microphone on the device.
	MicrophoneForceMute *bool `json:"microphoneForceMute,omitempty"`
	// NetworkEscapeHatchAllowed Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
	NetworkEscapeHatchAllowed *bool `json:"networkEscapeHatchAllowed,omitempty"`
	// NfcBlockOutgoingBeam Indicates whether or not to block NFC outgoing beam.
	NfcBlockOutgoingBeam *bool `json:"nfcBlockOutgoingBeam,omitempty"`
	// PasswordBlockKeyguard Indicates whether or not the keyguard is disabled.
	PasswordBlockKeyguard *bool `json:"passwordBlockKeyguard,omitempty"`
	// PasswordBlockKeyguardFeatures List of device keyguard features to block. This collection can contain a maximum of 7 elements.
	PasswordBlockKeyguardFeatures []AndroidKeyguardFeature `json:"passwordBlockKeyguardFeatures,omitempty"`
	// PasswordExpirationDays Indicates the amount of time in seconds that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength Indicates the minimum length of the password required on the device. Valid values 4 to 16
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
	// PasswordMinimumUpperCaseCharacters Indicates the minimum number of upper caseletter characters required for device password. Valid values 1 to 16
	PasswordMinimumUpperCaseCharacters *int `json:"passwordMinimumUpperCaseCharacters,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout Milliseconds of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordCountToBlock Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
	PasswordPreviousPasswordCountToBlock *int `json:"passwordPreviousPasswordCountToBlock,omitempty"`
	// PasswordRequiredType Indicates the minimum password quality required on the device.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// PlayStoreMode Indicates the Play Store mode of the device.
	PlayStoreMode *AndroidDeviceOwnerPlayStoreMode `json:"playStoreMode,omitempty"`
	// SafeBootBlocked Indicates whether or not rebooting the device into safe boot is disabled.
	SafeBootBlocked *bool `json:"safeBootBlocked,omitempty"`
	// ScreenCaptureBlocked Indicates whether or not to disable the capability to take screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// SecurityAllowDebuggingFeatures Indicates whether or not to block the user from enabling debugging features on the device.
	SecurityAllowDebuggingFeatures *bool `json:"securityAllowDebuggingFeatures,omitempty"`
	// SecurityRequireVerifyApps Indicates whether or not verify apps is required.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// StatusBarBlocked Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
	StatusBarBlocked *bool `json:"statusBarBlocked,omitempty"`
	// StayOnModes List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
	StayOnModes []AndroidDeviceOwnerBatteryPluggedMode `json:"stayOnModes,omitempty"`
	// StorageAllowUsb Indicates whether or not to allow USB mass storage.
	StorageAllowUsb *bool `json:"storageAllowUsb,omitempty"`
	// StorageBlockExternalMedia Indicates whether or not to block external media.
	StorageBlockExternalMedia *bool `json:"storageBlockExternalMedia,omitempty"`
	// StorageBlockUsbFileTransfer Indicates whether or not to block USB file transfer.
	StorageBlockUsbFileTransfer *bool `json:"storageBlockUsbFileTransfer,omitempty"`
	// SystemUpdateWindowStartMinutesAfterMidnight Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
	SystemUpdateWindowStartMinutesAfterMidnight *int `json:"systemUpdateWindowStartMinutesAfterMidnight,omitempty"`
	// SystemUpdateWindowEndMinutesAfterMidnight Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
	SystemUpdateWindowEndMinutesAfterMidnight *int `json:"systemUpdateWindowEndMinutesAfterMidnight,omitempty"`
	// SystemUpdateInstallType The type of system update configuration.
	SystemUpdateInstallType *AndroidDeviceOwnerSystemUpdateInstallType `json:"systemUpdateInstallType,omitempty"`
	// SystemWindowsBlocked Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
	SystemWindowsBlocked *bool `json:"systemWindowsBlocked,omitempty"`
	// UsersBlockAdd Indicates whether or not adding users and profiles is disabled.
	UsersBlockAdd *bool `json:"usersBlockAdd,omitempty"`
	// UsersBlockRemove Indicates whether or not to disable removing other users from the device.
	UsersBlockRemove *bool `json:"usersBlockRemove,omitempty"`
	// VolumeBlockAdjustment Indicates whether or not adjusting the master volume is disabled.
	VolumeBlockAdjustment *bool `json:"volumeBlockAdjustment,omitempty"`
	// VpnAlwaysOnPackageIdentifier Android app package name for app that will handle an always-on VPN connection.
	VpnAlwaysOnPackageIdentifier *string `json:"vpnAlwaysOnPackageIdentifier,omitempty"`
	// VpnAlwaysOnLockdownMode If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
	VpnAlwaysOnLockdownMode *bool `json:"vpnAlwaysOnLockdownMode,omitempty"`
	// WifiBlockEditConfigurations Indicates whether or not to block the user from editing the wifi connection settings.
	WifiBlockEditConfigurations *bool `json:"wifiBlockEditConfigurations,omitempty"`
	// WifiBlockEditPolicyDefinedConfigurations Indicates whether or not to block the user from editing just the networks defined by the policy.
	WifiBlockEditPolicyDefinedConfigurations *bool `json:"wifiBlockEditPolicyDefinedConfigurations,omitempty"`
}
