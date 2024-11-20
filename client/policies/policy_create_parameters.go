// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewPolicyCreateParams creates a new PolicyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPolicyCreateParams() *PolicyCreateParams {
	return &PolicyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPolicyCreateParamsWithTimeout creates a new PolicyCreateParams object
// with the ability to set a timeout on a request.
func NewPolicyCreateParamsWithTimeout(timeout time.Duration) *PolicyCreateParams {
	return &PolicyCreateParams{
		timeout: timeout,
	}
}

// NewPolicyCreateParamsWithContext creates a new PolicyCreateParams object
// with the ability to set a context for a request.
func NewPolicyCreateParamsWithContext(ctx context.Context) *PolicyCreateParams {
	return &PolicyCreateParams{
		Context: ctx,
	}
}

// NewPolicyCreateParamsWithHTTPClient creates a new PolicyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPolicyCreateParamsWithHTTPClient(client *http.Client) *PolicyCreateParams {
	return &PolicyCreateParams{
		HTTPClient: client,
	}
}

/*
PolicyCreateParams contains all the parameters to send to the API endpoint

	for the policy create operation.

	Typically these are written to a http.Request.
*/
type PolicyCreateParams struct {

	// Authorization.
	Authorization *string

	/* Body.

	   A Policy object
	*/
	Body *models.Policy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the policy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyCreateParams) WithDefaults() *PolicyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the policy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the policy create params
func (o *PolicyCreateParams) WithTimeout(timeout time.Duration) *PolicyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policy create params
func (o *PolicyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policy create params
func (o *PolicyCreateParams) WithContext(ctx context.Context) *PolicyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policy create params
func (o *PolicyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policy create params
func (o *PolicyCreateParams) WithHTTPClient(client *http.Client) *PolicyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policy create params
func (o *PolicyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the policy create params
func (o *PolicyCreateParams) WithAuthorization(authorization *string) *PolicyCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the policy create params
func (o *PolicyCreateParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the policy create params
func (o *PolicyCreateParams) WithBody(body *models.Policy) *PolicyCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the policy create params
func (o *PolicyCreateParams) SetBody(body *models.Policy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PolicyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
