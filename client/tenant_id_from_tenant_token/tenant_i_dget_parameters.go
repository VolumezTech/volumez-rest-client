// Code generated by go-swagger; DO NOT EDIT.

package tenant_id_from_tenant_token

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
)

// NewTenantIDgetParams creates a new TenantIDgetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenantIDgetParams() *TenantIDgetParams {
	return &TenantIDgetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenantIDgetParamsWithTimeout creates a new TenantIDgetParams object
// with the ability to set a timeout on a request.
func NewTenantIDgetParamsWithTimeout(timeout time.Duration) *TenantIDgetParams {
	return &TenantIDgetParams{
		timeout: timeout,
	}
}

// NewTenantIDgetParamsWithContext creates a new TenantIDgetParams object
// with the ability to set a context for a request.
func NewTenantIDgetParamsWithContext(ctx context.Context) *TenantIDgetParams {
	return &TenantIDgetParams{
		Context: ctx,
	}
}

// NewTenantIDgetParamsWithHTTPClient creates a new TenantIDgetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenantIDgetParamsWithHTTPClient(client *http.Client) *TenantIDgetParams {
	return &TenantIDgetParams{
		HTTPClient: client,
	}
}

/*
TenantIDgetParams contains all the parameters to send to the API endpoint

	for the tenant i dget operation.

	Typically these are written to a http.Request.
*/
type TenantIDgetParams struct {

	// Tenanttoken.
	Tenanttoken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenant i dget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantIDgetParams) WithDefaults() *TenantIDgetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenant i dget params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantIDgetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenant i dget params
func (o *TenantIDgetParams) WithTimeout(timeout time.Duration) *TenantIDgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenant i dget params
func (o *TenantIDgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenant i dget params
func (o *TenantIDgetParams) WithContext(ctx context.Context) *TenantIDgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenant i dget params
func (o *TenantIDgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenant i dget params
func (o *TenantIDgetParams) WithHTTPClient(client *http.Client) *TenantIDgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenant i dget params
func (o *TenantIDgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenanttoken adds the tenanttoken to the tenant i dget params
func (o *TenantIDgetParams) WithTenanttoken(tenanttoken string) *TenantIDgetParams {
	o.SetTenanttoken(tenanttoken)
	return o
}

// SetTenanttoken adds the tenanttoken to the tenant i dget params
func (o *TenantIDgetParams) SetTenanttoken(tenanttoken string) {
	o.Tenanttoken = tenanttoken
}

// WriteToRequest writes these params to a swagger request
func (o *TenantIDgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param tenanttoken
	if err := r.SetHeaderParam("tenanttoken", o.Tenanttoken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
