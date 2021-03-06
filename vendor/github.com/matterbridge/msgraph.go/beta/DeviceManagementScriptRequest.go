// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// DeviceManagementScriptRequestBuilder is request builder for DeviceManagementScript
type DeviceManagementScriptRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementScriptRequest
func (b *DeviceManagementScriptRequestBuilder) Request() *DeviceManagementScriptRequest {
	return &DeviceManagementScriptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementScriptRequest is request for DeviceManagementScript
type DeviceManagementScriptRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Get(ctx context.Context) (resObj *DeviceManagementScript, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Update(ctx context.Context, reqObj *DeviceManagementScript) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for DeviceManagementScriptAssignment collection
func (b *DeviceManagementScriptRequestBuilder) Assignments() *DeviceManagementScriptAssignmentsCollectionRequestBuilder {
	bb := &DeviceManagementScriptAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// DeviceManagementScriptAssignmentsCollectionRequestBuilder is request builder for DeviceManagementScriptAssignment collection
type DeviceManagementScriptAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptAssignment collection
func (b *DeviceManagementScriptAssignmentsCollectionRequestBuilder) Request() *DeviceManagementScriptAssignmentsCollectionRequest {
	return &DeviceManagementScriptAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptAssignment item
func (b *DeviceManagementScriptAssignmentsCollectionRequestBuilder) ID(id string) *DeviceManagementScriptAssignmentRequestBuilder {
	bb := &DeviceManagementScriptAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptAssignmentsCollectionRequest is request for DeviceManagementScriptAssignment collection
type DeviceManagementScriptAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementScriptAssignment, error) {
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
	var values []DeviceManagementScriptAssignment
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
			value  []DeviceManagementScriptAssignment
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

// Get performs GET request for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Get(ctx context.Context) ([]DeviceManagementScriptAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementScriptAssignment) (resObj *DeviceManagementScriptAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceRunStates returns request builder for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptRequestBuilder) DeviceRunStates() *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder {
	bb := &DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceRunStates"
	return bb
}

// DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder is request builder for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder) Request() *DeviceManagementScriptDeviceRunStatesCollectionRequest {
	return &DeviceManagementScriptDeviceRunStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptDeviceState item
func (b *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder) ID(id string) *DeviceManagementScriptDeviceStateRequestBuilder {
	bb := &DeviceManagementScriptDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptDeviceRunStatesCollectionRequest is request for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptDeviceRunStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementScriptDeviceState, error) {
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
	var values []DeviceManagementScriptDeviceState
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
			value  []DeviceManagementScriptDeviceState
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

// Get performs GET request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Get(ctx context.Context) ([]DeviceManagementScriptDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementScriptDeviceState) (resObj *DeviceManagementScriptDeviceState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// GroupAssignments returns request builder for DeviceManagementScriptGroupAssignment collection
func (b *DeviceManagementScriptRequestBuilder) GroupAssignments() *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder {
	bb := &DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/groupAssignments"
	return bb
}

// DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder is request builder for DeviceManagementScriptGroupAssignment collection
type DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptGroupAssignment collection
func (b *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder) Request() *DeviceManagementScriptGroupAssignmentsCollectionRequest {
	return &DeviceManagementScriptGroupAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptGroupAssignment item
func (b *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder) ID(id string) *DeviceManagementScriptGroupAssignmentRequestBuilder {
	bb := &DeviceManagementScriptGroupAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptGroupAssignmentsCollectionRequest is request for DeviceManagementScriptGroupAssignment collection
type DeviceManagementScriptGroupAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementScriptGroupAssignment, error) {
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
	var values []DeviceManagementScriptGroupAssignment
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
			value  []DeviceManagementScriptGroupAssignment
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

// Get performs GET request for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Get(ctx context.Context) ([]DeviceManagementScriptGroupAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementScriptGroupAssignment) (resObj *DeviceManagementScriptGroupAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RunSummary is navigation property
func (b *DeviceManagementScriptRequestBuilder) RunSummary() *DeviceManagementScriptRunSummaryRequestBuilder {
	bb := &DeviceManagementScriptRunSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/runSummary"
	return bb
}

// UserRunStates returns request builder for DeviceManagementScriptUserState collection
func (b *DeviceManagementScriptRequestBuilder) UserRunStates() *DeviceManagementScriptUserRunStatesCollectionRequestBuilder {
	bb := &DeviceManagementScriptUserRunStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userRunStates"
	return bb
}

// DeviceManagementScriptUserRunStatesCollectionRequestBuilder is request builder for DeviceManagementScriptUserState collection
type DeviceManagementScriptUserRunStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptUserState collection
func (b *DeviceManagementScriptUserRunStatesCollectionRequestBuilder) Request() *DeviceManagementScriptUserRunStatesCollectionRequest {
	return &DeviceManagementScriptUserRunStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptUserState item
func (b *DeviceManagementScriptUserRunStatesCollectionRequestBuilder) ID(id string) *DeviceManagementScriptUserStateRequestBuilder {
	bb := &DeviceManagementScriptUserStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptUserRunStatesCollectionRequest is request for DeviceManagementScriptUserState collection
type DeviceManagementScriptUserRunStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementScriptUserState, error) {
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
	var values []DeviceManagementScriptUserState
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
			value  []DeviceManagementScriptUserState
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

// Get performs GET request for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Get(ctx context.Context) ([]DeviceManagementScriptUserState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementScriptUserState) (resObj *DeviceManagementScriptUserState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
