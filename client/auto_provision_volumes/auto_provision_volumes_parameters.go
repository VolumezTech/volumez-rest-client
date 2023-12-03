// Code generated by go-swagger; DO NOT EDIT.

package auto_provision_volumes

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

// NewAutoProvisionVolumesParams creates a new AutoProvisionVolumesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutoProvisionVolumesParams() *AutoProvisionVolumesParams {
	return &AutoProvisionVolumesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutoProvisionVolumesParamsWithTimeout creates a new AutoProvisionVolumesParams object
// with the ability to set a timeout on a request.
func NewAutoProvisionVolumesParamsWithTimeout(timeout time.Duration) *AutoProvisionVolumesParams {
	return &AutoProvisionVolumesParams{
		timeout: timeout,
	}
}

// NewAutoProvisionVolumesParamsWithContext creates a new AutoProvisionVolumesParams object
// with the ability to set a context for a request.
func NewAutoProvisionVolumesParamsWithContext(ctx context.Context) *AutoProvisionVolumesParams {
	return &AutoProvisionVolumesParams{
		Context: ctx,
	}
}

// NewAutoProvisionVolumesParamsWithHTTPClient creates a new AutoProvisionVolumesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutoProvisionVolumesParamsWithHTTPClient(client *http.Client) *AutoProvisionVolumesParams {
	return &AutoProvisionVolumesParams{
		HTTPClient: client,
	}
}

/*
AutoProvisionVolumesParams contains all the parameters to send to the API endpoint

	for the auto provision volumes operation.

	Typically these are written to a http.Request.
*/
type AutoProvisionVolumesParams struct {

	// Authorization.
	Authorization *string

	/* Body.

	   Auto Provision Volume object
	*/
	Body *models.AutoProvisionVolume

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auto provision volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoProvisionVolumesParams) WithDefaults() *AutoProvisionVolumesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auto provision volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoProvisionVolumesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auto provision volumes params
func (o *AutoProvisionVolumesParams) WithTimeout(timeout time.Duration) *AutoProvisionVolumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auto provision volumes params
func (o *AutoProvisionVolumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auto provision volumes params
func (o *AutoProvisionVolumesParams) WithContext(ctx context.Context) *AutoProvisionVolumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auto provision volumes params
func (o *AutoProvisionVolumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auto provision volumes params
func (o *AutoProvisionVolumesParams) WithHTTPClient(client *http.Client) *AutoProvisionVolumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auto provision volumes params
func (o *AutoProvisionVolumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the auto provision volumes params
func (o *AutoProvisionVolumesParams) WithAuthorization(authorization *string) *AutoProvisionVolumesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the auto provision volumes params
func (o *AutoProvisionVolumesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the auto provision volumes params
func (o *AutoProvisionVolumesParams) WithBody(body *models.AutoProvisionVolume) *AutoProvisionVolumesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auto provision volumes params
func (o *AutoProvisionVolumesParams) SetBody(body *models.AutoProvisionVolume) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AutoProvisionVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
