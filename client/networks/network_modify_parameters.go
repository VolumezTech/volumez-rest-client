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

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NewNetworkModifyParams creates a new NetworkModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkModifyParams() *NetworkModifyParams {
	return &NetworkModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkModifyParamsWithTimeout creates a new NetworkModifyParams object
// with the ability to set a timeout on a request.
func NewNetworkModifyParamsWithTimeout(timeout time.Duration) *NetworkModifyParams {
	return &NetworkModifyParams{
		timeout: timeout,
	}
}

// NewNetworkModifyParamsWithContext creates a new NetworkModifyParams object
// with the ability to set a context for a request.
func NewNetworkModifyParamsWithContext(ctx context.Context) *NetworkModifyParams {
	return &NetworkModifyParams{
		Context: ctx,
	}
}

// NewNetworkModifyParamsWithHTTPClient creates a new NetworkModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkModifyParamsWithHTTPClient(client *http.Client) *NetworkModifyParams {
	return &NetworkModifyParams{
		HTTPClient: client,
	}
}

/* NetworkModifyParams contains all the parameters to send to the API endpoint
   for the network modify operation.

   Typically these are written to a http.Request.
*/
type NetworkModifyParams struct {

	/* Body.

	   A Network object
	*/
	Body *models.Network

	// Network.
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkModifyParams) WithDefaults() *NetworkModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network modify params
func (o *NetworkModifyParams) WithTimeout(timeout time.Duration) *NetworkModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network modify params
func (o *NetworkModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network modify params
func (o *NetworkModifyParams) WithContext(ctx context.Context) *NetworkModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network modify params
func (o *NetworkModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network modify params
func (o *NetworkModifyParams) WithHTTPClient(client *http.Client) *NetworkModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network modify params
func (o *NetworkModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the network modify params
func (o *NetworkModifyParams) WithBody(body *models.Network) *NetworkModifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the network modify params
func (o *NetworkModifyParams) SetBody(body *models.Network) {
	o.Body = body
}

// WithNetwork adds the network to the network modify params
func (o *NetworkModifyParams) WithNetwork(network string) *NetworkModifyParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the network modify params
func (o *NetworkModifyParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network
	if err := r.SetPathParam("network", o.Network); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
