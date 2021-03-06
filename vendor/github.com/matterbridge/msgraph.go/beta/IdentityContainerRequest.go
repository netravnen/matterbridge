// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// IdentityContainerRequestBuilder is request builder for IdentityContainer
type IdentityContainerRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityContainerRequest
func (b *IdentityContainerRequestBuilder) Request() *IdentityContainerRequest {
	return &IdentityContainerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityContainerRequest is request for IdentityContainer
type IdentityContainerRequest struct{ BaseRequest }

// Get performs GET request for IdentityContainer
func (r *IdentityContainerRequest) Get(ctx context.Context) (resObj *IdentityContainer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IdentityContainer
func (r *IdentityContainerRequest) Update(ctx context.Context, reqObj *IdentityContainer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IdentityContainer
func (r *IdentityContainerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserFlows returns request builder for IdentityUserFlow collection
func (b *IdentityContainerRequestBuilder) UserFlows() *IdentityContainerUserFlowsCollectionRequestBuilder {
	bb := &IdentityContainerUserFlowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userFlows"
	return bb
}

// IdentityContainerUserFlowsCollectionRequestBuilder is request builder for IdentityUserFlow collection
type IdentityContainerUserFlowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityUserFlow collection
func (b *IdentityContainerUserFlowsCollectionRequestBuilder) Request() *IdentityContainerUserFlowsCollectionRequest {
	return &IdentityContainerUserFlowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityUserFlow item
func (b *IdentityContainerUserFlowsCollectionRequestBuilder) ID(id string) *IdentityUserFlowRequestBuilder {
	bb := &IdentityUserFlowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IdentityContainerUserFlowsCollectionRequest is request for IdentityUserFlow collection
type IdentityContainerUserFlowsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]IdentityUserFlow, error) {
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
	var values []IdentityUserFlow
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
			value  []IdentityUserFlow
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

// Get performs GET request for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Get(ctx context.Context) ([]IdentityUserFlow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for IdentityUserFlow collection
func (r *IdentityContainerUserFlowsCollectionRequest) Add(ctx context.Context, reqObj *IdentityUserFlow) (resObj *IdentityUserFlow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
