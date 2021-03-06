// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSUpdateDeviceStatusRequestBuilder is request builder for IOSUpdateDeviceStatus
type IOSUpdateDeviceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSUpdateDeviceStatusRequest
func (b *IOSUpdateDeviceStatusRequestBuilder) Request() *IOSUpdateDeviceStatusRequest {
	return &IOSUpdateDeviceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSUpdateDeviceStatusRequest is request for IOSUpdateDeviceStatus
type IOSUpdateDeviceStatusRequest struct{ BaseRequest }

// Get performs GET request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Get(ctx context.Context) (resObj *IOSUpdateDeviceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Update(ctx context.Context, reqObj *IOSUpdateDeviceStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
