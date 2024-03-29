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

// NewConnectivitiesListParams creates a new ConnectivitiesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectivitiesListParams() *ConnectivitiesListParams {
	return &ConnectivitiesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectivitiesListParamsWithTimeout creates a new ConnectivitiesListParams object
// with the ability to set a timeout on a request.
func NewConnectivitiesListParamsWithTimeout(timeout time.Duration) *ConnectivitiesListParams {
	return &ConnectivitiesListParams{
		timeout: timeout,
	}
}

// NewConnectivitiesListParamsWithContext creates a new ConnectivitiesListParams object
// with the ability to set a context for a request.
func NewConnectivitiesListParamsWithContext(ctx context.Context) *ConnectivitiesListParams {
	return &ConnectivitiesListParams{
		Context: ctx,
	}
}

// NewConnectivitiesListParamsWithHTTPClient creates a new ConnectivitiesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectivitiesListParamsWithHTTPClient(client *http.Client) *ConnectivitiesListParams {
	return &ConnectivitiesListParams{
		HTTPClient: client,
	}
}

/*
ConnectivitiesListParams contains all the parameters to send to the API endpoint

	for the connectivities list operation.

	Typically these are written to a http.Request.
*/
type ConnectivitiesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the connectivities list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivitiesListParams) WithDefaults() *ConnectivitiesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connectivities list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivitiesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the connectivities list params
func (o *ConnectivitiesListParams) WithTimeout(timeout time.Duration) *ConnectivitiesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connectivities list params
func (o *ConnectivitiesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connectivities list params
func (o *ConnectivitiesListParams) WithContext(ctx context.Context) *ConnectivitiesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connectivities list params
func (o *ConnectivitiesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connectivities list params
func (o *ConnectivitiesListParams) WithHTTPClient(client *http.Client) *ConnectivitiesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connectivities list params
func (o *ConnectivitiesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectivitiesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
