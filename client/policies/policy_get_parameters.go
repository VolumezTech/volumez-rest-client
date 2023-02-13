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
)

// NewPolicyGetParams creates a new PolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPolicyGetParams() *PolicyGetParams {
	return &PolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPolicyGetParamsWithTimeout creates a new PolicyGetParams object
// with the ability to set a timeout on a request.
func NewPolicyGetParamsWithTimeout(timeout time.Duration) *PolicyGetParams {
	return &PolicyGetParams{
		timeout: timeout,
	}
}

// NewPolicyGetParamsWithContext creates a new PolicyGetParams object
// with the ability to set a context for a request.
func NewPolicyGetParamsWithContext(ctx context.Context) *PolicyGetParams {
	return &PolicyGetParams{
		Context: ctx,
	}
}

// NewPolicyGetParamsWithHTTPClient creates a new PolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPolicyGetParamsWithHTTPClient(client *http.Client) *PolicyGetParams {
	return &PolicyGetParams{
		HTTPClient: client,
	}
}

/* PolicyGetParams contains all the parameters to send to the API endpoint
   for the policy get operation.

   Typically these are written to a http.Request.
*/
type PolicyGetParams struct {

	// Policy.
	Policy string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyGetParams) WithDefaults() *PolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the policy get params
func (o *PolicyGetParams) WithTimeout(timeout time.Duration) *PolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policy get params
func (o *PolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policy get params
func (o *PolicyGetParams) WithContext(ctx context.Context) *PolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policy get params
func (o *PolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policy get params
func (o *PolicyGetParams) WithHTTPClient(client *http.Client) *PolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policy get params
func (o *PolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the policy get params
func (o *PolicyGetParams) WithPolicy(policy string) *PolicyGetParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the policy get params
func (o *PolicyGetParams) SetPolicy(policy string) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param policy
	if err := r.SetPathParam("policy", o.Policy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
