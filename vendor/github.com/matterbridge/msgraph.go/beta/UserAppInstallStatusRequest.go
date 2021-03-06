// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// UserAppInstallStatusRequestBuilder is request builder for UserAppInstallStatus
type UserAppInstallStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAppInstallStatusRequest
func (b *UserAppInstallStatusRequestBuilder) Request() *UserAppInstallStatusRequest {
	return &UserAppInstallStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAppInstallStatusRequest is request for UserAppInstallStatus
type UserAppInstallStatusRequest struct{ BaseRequest }

// Get performs GET request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Get(ctx context.Context) (resObj *UserAppInstallStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Update(ctx context.Context, reqObj *UserAppInstallStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// App is navigation property
func (b *UserAppInstallStatusRequestBuilder) App() *MobileAppRequestBuilder {
	bb := &MobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/app"
	return bb
}

// DeviceStatuses returns request builder for MobileAppInstallStatus collection
func (b *UserAppInstallStatusRequestBuilder) DeviceStatuses() *UserAppInstallStatusDeviceStatusesCollectionRequestBuilder {
	bb := &UserAppInstallStatusDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// UserAppInstallStatusDeviceStatusesCollectionRequestBuilder is request builder for MobileAppInstallStatus collection
type UserAppInstallStatusDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppInstallStatus collection
func (b *UserAppInstallStatusDeviceStatusesCollectionRequestBuilder) Request() *UserAppInstallStatusDeviceStatusesCollectionRequest {
	return &UserAppInstallStatusDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppInstallStatus item
func (b *UserAppInstallStatusDeviceStatusesCollectionRequestBuilder) ID(id string) *MobileAppInstallStatusRequestBuilder {
	bb := &MobileAppInstallStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserAppInstallStatusDeviceStatusesCollectionRequest is request for MobileAppInstallStatus collection
type UserAppInstallStatusDeviceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppInstallStatus collection
func (r *UserAppInstallStatusDeviceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppInstallStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppInstallStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MobileAppInstallStatus
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for MobileAppInstallStatus collection
func (r *UserAppInstallStatusDeviceStatusesCollectionRequest) Get(ctx context.Context) ([]MobileAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppInstallStatus collection
func (r *UserAppInstallStatusDeviceStatusesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppInstallStatus) (resObj *MobileAppInstallStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
