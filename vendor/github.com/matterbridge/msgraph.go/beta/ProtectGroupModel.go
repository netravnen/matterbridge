// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProtectGroup undocumented
type ProtectGroup struct {
	// LabelActionBase is the base model of ProtectGroup
	LabelActionBase
	// AllowEmailFromGuestUsers undocumented
	AllowEmailFromGuestUsers *bool `json:"allowEmailFromGuestUsers,omitempty"`
	// AllowGuestUsers undocumented
	AllowGuestUsers *bool `json:"allowGuestUsers,omitempty"`
	// Privacy undocumented
	Privacy *GroupPrivacy `json:"privacy,omitempty"`
}
