// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidDeviceOwnerVpnConfigurationRequestBuilder is request builder for AndroidDeviceOwnerVpnConfiguration
type AndroidDeviceOwnerVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidDeviceOwnerVpnConfigurationRequest
func (b *AndroidDeviceOwnerVpnConfigurationRequestBuilder) Request() *AndroidDeviceOwnerVpnConfigurationRequest {
	return &AndroidDeviceOwnerVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidDeviceOwnerVpnConfigurationRequest is request for AndroidDeviceOwnerVpnConfiguration
type AndroidDeviceOwnerVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Get(ctx context.Context) (resObj *AndroidDeviceOwnerVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Update(ctx context.Context, reqObj *AndroidDeviceOwnerVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidDeviceOwnerVpnConfiguration
func (r *AndroidDeviceOwnerVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *AndroidDeviceOwnerVpnConfigurationRequestBuilder) IdentityCertificate() *AndroidDeviceOwnerCertificateProfileBaseRequestBuilder {
	bb := &AndroidDeviceOwnerCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}
