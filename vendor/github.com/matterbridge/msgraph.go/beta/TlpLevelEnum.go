// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TlpLevel undocumented
type TlpLevel int

const (
	// TlpLevelVUnknown undocumented
	TlpLevelVUnknown TlpLevel = 0
	// TlpLevelVWhite undocumented
	TlpLevelVWhite TlpLevel = 1
	// TlpLevelVGreen undocumented
	TlpLevelVGreen TlpLevel = 2
	// TlpLevelVAmber undocumented
	TlpLevelVAmber TlpLevel = 3
	// TlpLevelVRed undocumented
	TlpLevelVRed TlpLevel = 4
	// TlpLevelVUnknownFutureValue undocumented
	TlpLevelVUnknownFutureValue TlpLevel = 127
)

// TlpLevelPUnknown returns a pointer to TlpLevelVUnknown
func TlpLevelPUnknown() *TlpLevel {
	v := TlpLevelVUnknown
	return &v
}

// TlpLevelPWhite returns a pointer to TlpLevelVWhite
func TlpLevelPWhite() *TlpLevel {
	v := TlpLevelVWhite
	return &v
}

// TlpLevelPGreen returns a pointer to TlpLevelVGreen
func TlpLevelPGreen() *TlpLevel {
	v := TlpLevelVGreen
	return &v
}

// TlpLevelPAmber returns a pointer to TlpLevelVAmber
func TlpLevelPAmber() *TlpLevel {
	v := TlpLevelVAmber
	return &v
}

// TlpLevelPRed returns a pointer to TlpLevelVRed
func TlpLevelPRed() *TlpLevel {
	v := TlpLevelVRed
	return &v
}

// TlpLevelPUnknownFutureValue returns a pointer to TlpLevelVUnknownFutureValue
func TlpLevelPUnknownFutureValue() *TlpLevel {
	v := TlpLevelVUnknownFutureValue
	return &v
}
