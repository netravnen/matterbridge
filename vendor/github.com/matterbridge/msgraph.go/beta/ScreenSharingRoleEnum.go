// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ScreenSharingRole undocumented
type ScreenSharingRole int

const (
	// ScreenSharingRoleVViewer undocumented
	ScreenSharingRoleVViewer ScreenSharingRole = 0
	// ScreenSharingRoleVSharer undocumented
	ScreenSharingRoleVSharer ScreenSharingRole = 1
)

// ScreenSharingRolePViewer returns a pointer to ScreenSharingRoleVViewer
func ScreenSharingRolePViewer() *ScreenSharingRole {
	v := ScreenSharingRoleVViewer
	return &v
}

// ScreenSharingRolePSharer returns a pointer to ScreenSharingRoleVSharer
func ScreenSharingRolePSharer() *ScreenSharingRole {
	v := ScreenSharingRoleVSharer
	return &v
}
