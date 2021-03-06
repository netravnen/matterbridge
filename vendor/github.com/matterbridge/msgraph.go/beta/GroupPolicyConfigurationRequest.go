// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// GroupPolicyConfigurationRequestBuilder is request builder for GroupPolicyConfiguration
type GroupPolicyConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyConfigurationRequest
func (b *GroupPolicyConfigurationRequestBuilder) Request() *GroupPolicyConfigurationRequest {
	return &GroupPolicyConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyConfigurationRequest is request for GroupPolicyConfiguration
type GroupPolicyConfigurationRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyConfiguration
func (r *GroupPolicyConfigurationRequest) Get(ctx context.Context) (resObj *GroupPolicyConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyConfiguration
func (r *GroupPolicyConfigurationRequest) Update(ctx context.Context, reqObj *GroupPolicyConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyConfiguration
func (r *GroupPolicyConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for GroupPolicyConfigurationAssignment collection
func (b *GroupPolicyConfigurationRequestBuilder) Assignments() *GroupPolicyConfigurationAssignmentsCollectionRequestBuilder {
	bb := &GroupPolicyConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// GroupPolicyConfigurationAssignmentsCollectionRequestBuilder is request builder for GroupPolicyConfigurationAssignment collection
type GroupPolicyConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyConfigurationAssignment collection
func (b *GroupPolicyConfigurationAssignmentsCollectionRequestBuilder) Request() *GroupPolicyConfigurationAssignmentsCollectionRequest {
	return &GroupPolicyConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyConfigurationAssignment item
func (b *GroupPolicyConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *GroupPolicyConfigurationAssignmentRequestBuilder {
	bb := &GroupPolicyConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyConfigurationAssignmentsCollectionRequest is request for GroupPolicyConfigurationAssignment collection
type GroupPolicyConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GroupPolicyConfigurationAssignment collection
func (r *GroupPolicyConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GroupPolicyConfigurationAssignment, error) {
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
	var values []GroupPolicyConfigurationAssignment
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
			value  []GroupPolicyConfigurationAssignment
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

// Get performs GET request for GroupPolicyConfigurationAssignment collection
func (r *GroupPolicyConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]GroupPolicyConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GroupPolicyConfigurationAssignment collection
func (r *GroupPolicyConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *GroupPolicyConfigurationAssignment) (resObj *GroupPolicyConfigurationAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DefinitionValues returns request builder for GroupPolicyDefinitionValue collection
func (b *GroupPolicyConfigurationRequestBuilder) DefinitionValues() *GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder {
	bb := &GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitionValues"
	return bb
}

// GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder is request builder for GroupPolicyDefinitionValue collection
type GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyDefinitionValue collection
func (b *GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder) Request() *GroupPolicyConfigurationDefinitionValuesCollectionRequest {
	return &GroupPolicyConfigurationDefinitionValuesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyDefinitionValue item
func (b *GroupPolicyConfigurationDefinitionValuesCollectionRequestBuilder) ID(id string) *GroupPolicyDefinitionValueRequestBuilder {
	bb := &GroupPolicyDefinitionValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyConfigurationDefinitionValuesCollectionRequest is request for GroupPolicyDefinitionValue collection
type GroupPolicyConfigurationDefinitionValuesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GroupPolicyDefinitionValue collection
func (r *GroupPolicyConfigurationDefinitionValuesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GroupPolicyDefinitionValue, error) {
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
	var values []GroupPolicyDefinitionValue
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
			value  []GroupPolicyDefinitionValue
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

// Get performs GET request for GroupPolicyDefinitionValue collection
func (r *GroupPolicyConfigurationDefinitionValuesCollectionRequest) Get(ctx context.Context) ([]GroupPolicyDefinitionValue, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GroupPolicyDefinitionValue collection
func (r *GroupPolicyConfigurationDefinitionValuesCollectionRequest) Add(ctx context.Context, reqObj *GroupPolicyDefinitionValue) (resObj *GroupPolicyDefinitionValue, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
