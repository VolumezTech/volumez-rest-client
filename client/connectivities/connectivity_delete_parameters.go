// Code generated by go-swagger; DO NOT EDIT.

package connectivities

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

// NewConnectivityDeleteParams creates a new ConnectivityDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectivityDeleteParams() *ConnectivityDeleteParams {
	return &ConnectivityDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectivityDeleteParamsWithTimeout creates a new ConnectivityDeleteParams object
// with the ability to set a timeout on a request.
func NewConnectivityDeleteParamsWithTimeout(timeout time.Duration) *ConnectivityDeleteParams {
	return &ConnectivityDeleteParams{
		timeout: timeout,
	}
}

// NewConnectivityDeleteParamsWithContext creates a new ConnectivityDeleteParams object
// with the ability to set a context for a request.
func NewConnectivityDeleteParamsWithContext(ctx context.Context) *ConnectivityDeleteParams {
	return &ConnectivityDeleteParams{
		Context: ctx,
	}
}

// NewConnectivityDeleteParamsWithHTTPClient creates a new ConnectivityDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectivityDeleteParamsWithHTTPClient(client *http.Client) *ConnectivityDeleteParams {
	return &ConnectivityDeleteParams{
		HTTPClient: client,
	}
}

/* ConnectivityDeleteParams contains all the parameters to send to the API endpoint
   for the connectivity delete operation.

   Typically these are written to a http.Request.
*/
type ConnectivityDeleteParams struct {

	// Connectivity.
	Connectivity string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the connectivity delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivityDeleteParams) WithDefaults() *ConnectivityDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connectivity delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivityDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the connectivity delete params
func (o *ConnectivityDeleteParams) WithTimeout(timeout time.Duration) *ConnectivityDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connectivity delete params
func (o *ConnectivityDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connectivity delete params
func (o *ConnectivityDeleteParams) WithContext(ctx context.Context) *ConnectivityDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connectivity delete params
func (o *ConnectivityDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connectivity delete params
func (o *ConnectivityDeleteParams) WithHTTPClient(client *http.Client) *ConnectivityDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connectivity delete params
func (o *ConnectivityDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectivity adds the connectivity to the connectivity delete params
func (o *ConnectivityDeleteParams) WithConnectivity(connectivity string) *ConnectivityDeleteParams {
	o.SetConnectivity(connectivity)
	return o
}

// SetConnectivity adds the connectivity to the connectivity delete params
func (o *ConnectivityDeleteParams) SetConnectivity(connectivity string) {
	o.Connectivity = connectivity
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectivityDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectivity
	if err := r.SetPathParam("connectivity", o.Connectivity); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
