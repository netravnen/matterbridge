// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// IOSLobAppProvisioningConfigurationRequestBuilder is request builder for IOSLobAppProvisioningConfiguration
type IOSLobAppProvisioningConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSLobAppProvisioningConfigurationRequest
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) Request() *IOSLobAppProvisioningConfigurationRequest {
	return &IOSLobAppProvisioningConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSLobAppProvisioningConfigurationRequest is request for IOSLobAppProvisioningConfiguration
type IOSLobAppProvisioningConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Get(ctx context.Context) (resObj *IOSLobAppProvisioningConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Update(ctx context.Context, reqObj *IOSLobAppProvisioningConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for IOSLobAppProvisioningConfigurationAssignment collection
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) Assignments() *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder is request builder for IOSLobAppProvisioningConfigurationAssignment collection
type IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IOSLobAppProvisioningConfigurationAssignment collection
func (b *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder) Request() *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest {
	return &IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IOSLobAppProvisioningConfigurationAssignment item
func (b *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *IOSLobAppProvisioningConfigurationAssignmentRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest is request for IOSLobAppProvisioningConfigurationAssignment collection
type IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IOSLobAppProvisioningConfigurationAssignment collection
func (r *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]IOSLobAppProvisioningConfigurationAssignment, error) {
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
	var values []IOSLobAppProvisioningConfigurationAssignment
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
			value  []IOSLobAppProvisioningConfigurationAssignment
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

// Get performs GET request for IOSLobAppProvisioningConfigurationAssignment collection
func (r *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]IOSLobAppProvisioningConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for IOSLobAppProvisioningConfigurationAssignment collection
func (r *IOSLobAppProvisioningConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *IOSLobAppProvisioningConfigurationAssignment) (resObj *IOSLobAppProvisioningConfigurationAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceStatuses returns request builder for ManagedDeviceMobileAppConfigurationDeviceStatus collection
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) DeviceStatuses() *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationDeviceStatus collection
type IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceMobileAppConfigurationDeviceStatus collection
func (b *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder) Request() *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest {
	return &IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceMobileAppConfigurationDeviceStatus item
func (b *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequestBuilder) ID(id string) *ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder {
	bb := &ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest is request for ManagedDeviceMobileAppConfigurationDeviceStatus collection
type IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceMobileAppConfigurationDeviceStatus collection
func (r *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceMobileAppConfigurationDeviceStatus, error) {
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
	var values []ManagedDeviceMobileAppConfigurationDeviceStatus
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
			value  []ManagedDeviceMobileAppConfigurationDeviceStatus
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

// Get performs GET request for ManagedDeviceMobileAppConfigurationDeviceStatus collection
func (r *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceMobileAppConfigurationDeviceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceMobileAppConfigurationDeviceStatus collection
func (r *IOSLobAppProvisioningConfigurationDeviceStatusesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationDeviceStatus) (resObj *ManagedDeviceMobileAppConfigurationDeviceStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// GroupAssignments returns request builder for MobileAppProvisioningConfigGroupAssignment collection
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) GroupAssignments() *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/groupAssignments"
	return bb
}

// IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder is request builder for MobileAppProvisioningConfigGroupAssignment collection
type IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppProvisioningConfigGroupAssignment collection
func (b *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder) Request() *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest {
	return &IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppProvisioningConfigGroupAssignment item
func (b *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequestBuilder) ID(id string) *MobileAppProvisioningConfigGroupAssignmentRequestBuilder {
	bb := &MobileAppProvisioningConfigGroupAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest is request for MobileAppProvisioningConfigGroupAssignment collection
type IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppProvisioningConfigGroupAssignment collection
func (r *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppProvisioningConfigGroupAssignment, error) {
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
	var values []MobileAppProvisioningConfigGroupAssignment
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
			value  []MobileAppProvisioningConfigGroupAssignment
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

// Get performs GET request for MobileAppProvisioningConfigGroupAssignment collection
func (r *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest) Get(ctx context.Context) ([]MobileAppProvisioningConfigGroupAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppProvisioningConfigGroupAssignment collection
func (r *IOSLobAppProvisioningConfigurationGroupAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppProvisioningConfigGroupAssignment) (resObj *MobileAppProvisioningConfigGroupAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserStatuses returns request builder for ManagedDeviceMobileAppConfigurationUserStatus collection
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) UserStatuses() *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatuses"
	return bb
}

// IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationUserStatus collection
type IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceMobileAppConfigurationUserStatus collection
func (b *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder) Request() *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest {
	return &IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceMobileAppConfigurationUserStatus item
func (b *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequestBuilder) ID(id string) *ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder {
	bb := &ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest is request for ManagedDeviceMobileAppConfigurationUserStatus collection
type IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceMobileAppConfigurationUserStatus collection
func (r *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceMobileAppConfigurationUserStatus, error) {
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
	var values []ManagedDeviceMobileAppConfigurationUserStatus
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
			value  []ManagedDeviceMobileAppConfigurationUserStatus
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

// Get performs GET request for ManagedDeviceMobileAppConfigurationUserStatus collection
func (r *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceMobileAppConfigurationUserStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceMobileAppConfigurationUserStatus collection
func (r *IOSLobAppProvisioningConfigurationUserStatusesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationUserStatus) (resObj *ManagedDeviceMobileAppConfigurationUserStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
