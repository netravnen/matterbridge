// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TargetedManagedAppProtectionAssignRequestParameter undocumented
type TargetedManagedAppProtectionAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

//
type TargetedManagedAppProtectionAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *TargetedManagedAppProtectionRequestBuilder) Assign(reqObj *TargetedManagedAppProtectionAssignRequestParameter) *TargetedManagedAppProtectionAssignRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppProtectionAssignRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppProtectionAssignRequestBuilder) Request() *TargetedManagedAppProtectionAssignRequest {
	return &TargetedManagedAppProtectionAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppProtectionAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
