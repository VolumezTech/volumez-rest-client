// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewNetworkGetParams creates a new NetworkGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkGetParams() *NetworkGetParams {
	return &NetworkGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkGetParamsWithTimeout creates a new NetworkGetParams object
// with the ability to set a timeout on a request.
func NewNetworkGetParamsWithTimeout(timeout time.Duration) *NetworkGetParams {
	return &NetworkGetParams{
		timeout: timeout,
	}
}

// NewNetworkGetParamsWithContext creates a new NetworkGetParams object
// with the ability to set a context for a request.
func NewNetworkGetParamsWithContext(ctx context.Context) *NetworkGetParams {
	return &NetworkGetParams{
		Context: ctx,
	}
}

// NewNetworkGetParamsWithHTTPClient creates a new NetworkGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkGetParamsWithHTTPClient(client *http.Client) *NetworkGetParams {
	return &NetworkGetParams{
		HTTPClient: client,
	}
}

/*
NetworkGetParams contains all the parameters to send to the API endpoint

	for the network get operation.

	Typically these are written to a http.Request.
*/
type NetworkGetParams struct {

	// Network.
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkGetParams) WithDefaults() *NetworkGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network get params
func (o *NetworkGetParams) WithTimeout(timeout time.Duration) *NetworkGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network get params
func (o *NetworkGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network get params
func (o *NetworkGetParams) WithContext(ctx context.Context) *NetworkGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network get params
func (o *NetworkGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network get params
func (o *NetworkGetParams) WithHTTPClient(client *http.Client) *NetworkGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network get params
func (o *NetworkGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the network get params
func (o *NetworkGetParams) WithNetwork(network string) *NetworkGetParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the network get params
func (o *NetworkGetParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network
	if err := r.SetPathParam("network", o.Network); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
