// Code generated by go-swagger; DO NOT EDIT.

package tenant_refresh_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NewTenantRefreshTokenParams creates a new TenantRefreshTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenantRefreshTokenParams() *TenantRefreshTokenParams {
	return &TenantRefreshTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenantRefreshTokenParamsWithTimeout creates a new TenantRefreshTokenParams object
// with the ability to set a timeout on a request.
func NewTenantRefreshTokenParamsWithTimeout(timeout time.Duration) *TenantRefreshTokenParams {
	return &TenantRefreshTokenParams{
		timeout: timeout,
	}
}

// NewTenantRefreshTokenParamsWithContext creates a new TenantRefreshTokenParams object
// with the ability to set a context for a request.
func NewTenantRefreshTokenParamsWithContext(ctx context.Context) *TenantRefreshTokenParams {
	return &TenantRefreshTokenParams{
		Context: ctx,
	}
}

// NewTenantRefreshTokenParamsWithHTTPClient creates a new TenantRefreshTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenantRefreshTokenParamsWithHTTPClient(client *http.Client) *TenantRefreshTokenParams {
	return &TenantRefreshTokenParams{
		HTTPClient: client,
	}
}

/* TenantRefreshTokenParams contains all the parameters to send to the API endpoint
   for the tenant refresh token operation.

   Typically these are written to a http.Request.
*/
type TenantRefreshTokenParams struct {

	/* Body.

	   Tenant's Refresh Token
	*/
	Body *models.GetTenantRefreshTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenant refresh token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantRefreshTokenParams) WithDefaults() *TenantRefreshTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenant refresh token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantRefreshTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenant refresh token params
func (o *TenantRefreshTokenParams) WithTimeout(timeout time.Duration) *TenantRefreshTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenant refresh token params
func (o *TenantRefreshTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenant refresh token params
func (o *TenantRefreshTokenParams) WithContext(ctx context.Context) *TenantRefreshTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenant refresh token params
func (o *TenantRefreshTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenant refresh token params
func (o *TenantRefreshTokenParams) WithHTTPClient(client *http.Client) *TenantRefreshTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenant refresh token params
func (o *TenantRefreshTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tenant refresh token params
func (o *TenantRefreshTokenParams) WithBody(body *models.GetTenantRefreshTokenRequest) *TenantRefreshTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tenant refresh token params
func (o *TenantRefreshTokenParams) SetBody(body *models.GetTenantRefreshTokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TenantRefreshTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
