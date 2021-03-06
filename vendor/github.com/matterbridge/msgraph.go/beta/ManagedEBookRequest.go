// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ManagedEBookRequestBuilder is request builder for ManagedEBook
type ManagedEBookRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookRequest
func (b *ManagedEBookRequestBuilder) Request() *ManagedEBookRequest {
	return &ManagedEBookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookRequest is request for ManagedEBook
type ManagedEBookRequest struct{ BaseRequest }

// Get performs GET request for ManagedEBook
func (r *ManagedEBookRequest) Get(ctx context.Context) (resObj *ManagedEBook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedEBook
func (r *ManagedEBookRequest) Update(ctx context.Context, reqObj *ManagedEBook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedEBook
func (r *ManagedEBookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for ManagedEBookAssignment collection
func (b *ManagedEBookRequestBuilder) Assignments() *ManagedEBookAssignmentsCollectionRequestBuilder {
	bb := &ManagedEBookAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// ManagedEBookAssignmentsCollectionRequestBuilder is request builder for ManagedEBookAssignment collection
type ManagedEBookAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedEBookAssignment collection
func (b *ManagedEBookAssignmentsCollectionRequestBuilder) Request() *ManagedEBookAssignmentsCollectionRequest {
	return &ManagedEBookAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedEBookAssignment item
func (b *ManagedEBookAssignmentsCollectionRequestBuilder) ID(id string) *ManagedEBookAssignmentRequestBuilder {
	bb := &ManagedEBookAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedEBookAssignmentsCollectionRequest is request for ManagedEBookAssignment collection
type ManagedEBookAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedEBookAssignment collection
func (r *ManagedEBookAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedEBookAssignment, error) {
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
	var values []ManagedEBookAssignment
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
			value  []ManagedEBookAssignment
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

// Get performs GET request for ManagedEBookAssignment collection
func (r *ManagedEBookAssignmentsCollectionRequest) Get(ctx context.Context) ([]ManagedEBookAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedEBookAssignment collection
func (r *ManagedEBookAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *ManagedEBookAssignment) (resObj *ManagedEBookAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Categories returns request builder for ManagedEBookCategory collection
func (b *ManagedEBookRequestBuilder) Categories() *ManagedEBookCategoriesCollectionRequestBuilder {
	bb := &ManagedEBookCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// ManagedEBookCategoriesCollectionRequestBuilder is request builder for ManagedEBookCategory collection
type ManagedEBookCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedEBookCategory collection
func (b *ManagedEBookCategoriesCollectionRequestBuilder) Request() *ManagedEBookCategoriesCollectionRequest {
	return &ManagedEBookCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedEBookCategory item
func (b *ManagedEBookCategoriesCollectionRequestBuilder) ID(id string) *ManagedEBookCategoryRequestBuilder {
	bb := &ManagedEBookCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedEBookCategoriesCollectionRequest is request for ManagedEBookCategory collection
type ManagedEBookCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedEBookCategory collection
func (r *ManagedEBookCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedEBookCategory, error) {
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
	var values []ManagedEBookCategory
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
			value  []ManagedEBookCategory
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

// Get performs GET request for ManagedEBookCategory collection
func (r *ManagedEBookCategoriesCollectionRequest) Get(ctx context.Context) ([]ManagedEBookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedEBookCategory collection
func (r *ManagedEBookCategoriesCollectionRequest) Add(ctx context.Context, reqObj *ManagedEBookCategory) (resObj *ManagedEBookCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceStates returns request builder for DeviceInstallState collection
func (b *ManagedEBookRequestBuilder) DeviceStates() *ManagedEBookDeviceStatesCollectionRequestBuilder {
	bb := &ManagedEBookDeviceStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStates"
	return bb
}

// ManagedEBookDeviceStatesCollectionRequestBuilder is request builder for DeviceInstallState collection
type ManagedEBookDeviceStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceInstallState collection
func (b *ManagedEBookDeviceStatesCollectionRequestBuilder) Request() *ManagedEBookDeviceStatesCollectionRequest {
	return &ManagedEBookDeviceStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceInstallState item
func (b *ManagedEBookDeviceStatesCollectionRequestBuilder) ID(id string) *DeviceInstallStateRequestBuilder {
	bb := &DeviceInstallStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedEBookDeviceStatesCollectionRequest is request for DeviceInstallState collection
type ManagedEBookDeviceStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceInstallState collection
func (r *ManagedEBookDeviceStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceInstallState, error) {
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
	var values []DeviceInstallState
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
			value  []DeviceInstallState
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

// Get performs GET request for DeviceInstallState collection
func (r *ManagedEBookDeviceStatesCollectionRequest) Get(ctx context.Context) ([]DeviceInstallState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceInstallState collection
func (r *ManagedEBookDeviceStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceInstallState) (resObj *DeviceInstallState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// InstallSummary is navigation property
func (b *ManagedEBookRequestBuilder) InstallSummary() *EBookInstallSummaryRequestBuilder {
	bb := &EBookInstallSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installSummary"
	return bb
}

// UserStateSummary returns request builder for UserInstallStateSummary collection
func (b *ManagedEBookRequestBuilder) UserStateSummary() *ManagedEBookUserStateSummaryCollectionRequestBuilder {
	bb := &ManagedEBookUserStateSummaryCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStateSummary"
	return bb
}

// ManagedEBookUserStateSummaryCollectionRequestBuilder is request builder for UserInstallStateSummary collection
type ManagedEBookUserStateSummaryCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserInstallStateSummary collection
func (b *ManagedEBookUserStateSummaryCollectionRequestBuilder) Request() *ManagedEBookUserStateSummaryCollectionRequest {
	return &ManagedEBookUserStateSummaryCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserInstallStateSummary item
func (b *ManagedEBookUserStateSummaryCollectionRequestBuilder) ID(id string) *UserInstallStateSummaryRequestBuilder {
	bb := &UserInstallStateSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedEBookUserStateSummaryCollectionRequest is request for UserInstallStateSummary collection
type ManagedEBookUserStateSummaryCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserInstallStateSummary collection
func (r *ManagedEBookUserStateSummaryCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]UserInstallStateSummary, error) {
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
	var values []UserInstallStateSummary
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
			value  []UserInstallStateSummary
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

// Get performs GET request for UserInstallStateSummary collection
func (r *ManagedEBookUserStateSummaryCollectionRequest) Get(ctx context.Context) ([]UserInstallStateSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for UserInstallStateSummary collection
func (r *ManagedEBookUserStateSummaryCollectionRequest) Add(ctx context.Context, reqObj *UserInstallStateSummary) (resObj *UserInstallStateSummary, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
