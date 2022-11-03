// Code generated by go-swagger; DO NOT EDIT.

package version

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

// NewVersionGetParams creates a new VersionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionGetParams() *VersionGetParams {
	return &VersionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionGetParamsWithTimeout creates a new VersionGetParams object
// with the ability to set a timeout on a request.
func NewVersionGetParamsWithTimeout(timeout time.Duration) *VersionGetParams {
	return &VersionGetParams{
		timeout: timeout,
	}
}

// NewVersionGetParamsWithContext creates a new VersionGetParams object
// with the ability to set a context for a request.
func NewVersionGetParamsWithContext(ctx context.Context) *VersionGetParams {
	return &VersionGetParams{
		Context: ctx,
	}
}

// NewVersionGetParamsWithHTTPClient creates a new VersionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionGetParamsWithHTTPClient(client *http.Client) *VersionGetParams {
	return &VersionGetParams{
		HTTPClient: client,
	}
}

/* VersionGetParams contains all the parameters to send to the API endpoint
   for the version get operation.

   Typically these are written to a http.Request.
*/
type VersionGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the version get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionGetParams) WithDefaults() *VersionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the version get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the version get params
func (o *VersionGetParams) WithTimeout(timeout time.Duration) *VersionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version get params
func (o *VersionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version get params
func (o *VersionGetParams) WithContext(ctx context.Context) *VersionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version get params
func (o *VersionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version get params
func (o *VersionGetParams) WithHTTPClient(client *http.Client) *VersionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version get params
func (o *VersionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VersionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
