// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// EnrollmentProfileRequestBuilder is request builder for EnrollmentProfile
type EnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnrollmentProfileRequest
func (b *EnrollmentProfileRequestBuilder) Request() *EnrollmentProfileRequest {
	return &EnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EnrollmentProfileRequest is request for EnrollmentProfile
type EnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Get(ctx context.Context) (resObj *EnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Update(ctx context.Context, reqObj *EnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnrollmentProfile
func (r *EnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
