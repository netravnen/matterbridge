// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PowerActionType undocumented
type PowerActionType int

const (
	// PowerActionTypeVNotConfigured undocumented
	PowerActionTypeVNotConfigured PowerActionType = 0
	// PowerActionTypeVNoAction undocumented
	PowerActionTypeVNoAction PowerActionType = 1
	// PowerActionTypeVSleep undocumented
	PowerActionTypeVSleep PowerActionType = 2
	// PowerActionTypeVHibernate undocumented
	PowerActionTypeVHibernate PowerActionType = 3
	// PowerActionTypeVShutdown undocumented
	PowerActionTypeVShutdown PowerActionType = 4
)

// PowerActionTypePNotConfigured returns a pointer to PowerActionTypeVNotConfigured
func PowerActionTypePNotConfigured() *PowerActionType {
	v := PowerActionTypeVNotConfigured
	return &v
}

// PowerActionTypePNoAction returns a pointer to PowerActionTypeVNoAction
func PowerActionTypePNoAction() *PowerActionType {
	v := PowerActionTypeVNoAction
	return &v
}

// PowerActionTypePSleep returns a pointer to PowerActionTypeVSleep
func PowerActionTypePSleep() *PowerActionType {
	v := PowerActionTypeVSleep
	return &v
}

// PowerActionTypePHibernate returns a pointer to PowerActionTypeVHibernate
func PowerActionTypePHibernate() *PowerActionType {
	v := PowerActionTypeVHibernate
	return &v
}

// PowerActionTypePShutdown returns a pointer to PowerActionTypeVShutdown
func PowerActionTypePShutdown() *PowerActionType {
	v := PowerActionTypeVShutdown
	return &v
}
