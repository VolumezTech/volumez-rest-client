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

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NewConnectivityCreateParams creates a new ConnectivityCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectivityCreateParams() *ConnectivityCreateParams {
	return &ConnectivityCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectivityCreateParamsWithTimeout creates a new ConnectivityCreateParams object
// with the ability to set a timeout on a request.
func NewConnectivityCreateParamsWithTimeout(timeout time.Duration) *ConnectivityCreateParams {
	return &ConnectivityCreateParams{
		timeout: timeout,
	}
}

// NewConnectivityCreateParamsWithContext creates a new ConnectivityCreateParams object
// with the ability to set a context for a request.
func NewConnectivityCreateParamsWithContext(ctx context.Context) *ConnectivityCreateParams {
	return &ConnectivityCreateParams{
		Context: ctx,
	}
}

// NewConnectivityCreateParamsWithHTTPClient creates a new ConnectivityCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectivityCreateParamsWithHTTPClient(client *http.Client) *ConnectivityCreateParams {
	return &ConnectivityCreateParams{
		HTTPClient: client,
	}
}

/*
ConnectivityCreateParams contains all the parameters to send to the API endpoint

	for the connectivity create operation.

	Typically these are written to a http.Request.
*/
type ConnectivityCreateParams struct {

	/* Body.

	   A Connectivity object
	*/
	Body *models.Connectivity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the connectivity create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivityCreateParams) WithDefaults() *ConnectivityCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connectivity create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectivityCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the connectivity create params
func (o *ConnectivityCreateParams) WithTimeout(timeout time.Duration) *ConnectivityCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connectivity create params
func (o *ConnectivityCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connectivity create params
func (o *ConnectivityCreateParams) WithContext(ctx context.Context) *ConnectivityCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connectivity create params
func (o *ConnectivityCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connectivity create params
func (o *ConnectivityCreateParams) WithHTTPClient(client *http.Client) *ConnectivityCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connectivity create params
func (o *ConnectivityCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the connectivity create params
func (o *ConnectivityCreateParams) WithBody(body *models.Connectivity) *ConnectivityCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the connectivity create params
func (o *ConnectivityCreateParams) SetBody(body *models.Connectivity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectivityCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
