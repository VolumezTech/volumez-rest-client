// Code generated by go-swagger; DO NOT EDIT.

package media

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

// NewVirtualMediaDeleteParams creates a new VirtualMediaDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualMediaDeleteParams() *VirtualMediaDeleteParams {
	return &VirtualMediaDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMediaDeleteParamsWithTimeout creates a new VirtualMediaDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualMediaDeleteParamsWithTimeout(timeout time.Duration) *VirtualMediaDeleteParams {
	return &VirtualMediaDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualMediaDeleteParamsWithContext creates a new VirtualMediaDeleteParams object
// with the ability to set a context for a request.
func NewVirtualMediaDeleteParamsWithContext(ctx context.Context) *VirtualMediaDeleteParams {
	return &VirtualMediaDeleteParams{
		Context: ctx,
	}
}

// NewVirtualMediaDeleteParamsWithHTTPClient creates a new VirtualMediaDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualMediaDeleteParamsWithHTTPClient(client *http.Client) *VirtualMediaDeleteParams {
	return &VirtualMediaDeleteParams{
		HTTPClient: client,
	}
}

/*
VirtualMediaDeleteParams contains all the parameters to send to the API endpoint

	for the virtual media delete operation.

	Typically these are written to a http.Request.
*/
type VirtualMediaDeleteParams struct {

	// Authorization.
	Authorization *string

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtual media delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualMediaDeleteParams) WithDefaults() *VirtualMediaDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtual media delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualMediaDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtual media delete params
func (o *VirtualMediaDeleteParams) WithTimeout(timeout time.Duration) *VirtualMediaDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual media delete params
func (o *VirtualMediaDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual media delete params
func (o *VirtualMediaDeleteParams) WithContext(ctx context.Context) *VirtualMediaDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual media delete params
func (o *VirtualMediaDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtual media delete params
func (o *VirtualMediaDeleteParams) WithHTTPClient(client *http.Client) *VirtualMediaDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtual media delete params
func (o *VirtualMediaDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the virtual media delete params
func (o *VirtualMediaDeleteParams) WithAuthorization(authorization *string) *VirtualMediaDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the virtual media delete params
func (o *VirtualMediaDeleteParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithMedia adds the media to the virtual media delete params
func (o *VirtualMediaDeleteParams) WithMedia(media string) *VirtualMediaDeleteParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the virtual media delete params
func (o *VirtualMediaDeleteParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMediaDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param authorization
		if err := r.SetHeaderParam("authorization", *o.Authorization); err != nil {
			return err
		}
	}

	// path param media
	if err := r.SetPathParam("media", o.Media); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
