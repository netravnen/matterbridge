// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// FinancialsRequestBuilder is request builder for Financials
type FinancialsRequestBuilder struct{ BaseRequestBuilder }

// Request returns FinancialsRequest
func (b *FinancialsRequestBuilder) Request() *FinancialsRequest {
	return &FinancialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FinancialsRequest is request for Financials
type FinancialsRequest struct{ BaseRequest }

// Get performs GET request for Financials
func (r *FinancialsRequest) Get(ctx context.Context) (resObj *Financials, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Financials
func (r *FinancialsRequest) Update(ctx context.Context, reqObj *Financials) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Financials
func (r *FinancialsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Companies returns request builder for Company collection
func (b *FinancialsRequestBuilder) Companies() *FinancialsCompaniesCollectionRequestBuilder {
	bb := &FinancialsCompaniesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/companies"
	return bb
}

// FinancialsCompaniesCollectionRequestBuilder is request builder for Company collection
type FinancialsCompaniesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Company collection
func (b *FinancialsCompaniesCollectionRequestBuilder) Request() *FinancialsCompaniesCollectionRequest {
	return &FinancialsCompaniesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Company item
func (b *FinancialsCompaniesCollectionRequestBuilder) ID(id string) *CompanyRequestBuilder {
	bb := &CompanyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// FinancialsCompaniesCollectionRequest is request for Company collection
type FinancialsCompaniesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Company collection
func (r *FinancialsCompaniesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Company, error) {
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
	var values []Company
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
			value  []Company
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

// Get performs GET request for Company collection
func (r *FinancialsCompaniesCollectionRequest) Get(ctx context.Context) ([]Company, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Company collection
func (r *FinancialsCompaniesCollectionRequest) Add(ctx context.Context, reqObj *Company) (resObj *Company, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
