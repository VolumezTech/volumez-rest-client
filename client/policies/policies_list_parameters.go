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

// NewPoliciesListParams creates a new PoliciesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPoliciesListParams() *PoliciesListParams {
	return &PoliciesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPoliciesListParamsWithTimeout creates a new PoliciesListParams object
// with the ability to set a timeout on a request.
func NewPoliciesListParamsWithTimeout(timeout time.Duration) *PoliciesListParams {
	return &PoliciesListParams{
		timeout: timeout,
	}
}

// NewPoliciesListParamsWithContext creates a new PoliciesListParams object
// with the ability to set a context for a request.
func NewPoliciesListParamsWithContext(ctx context.Context) *PoliciesListParams {
	return &PoliciesListParams{
		Context: ctx,
	}
}

// NewPoliciesListParamsWithHTTPClient creates a new PoliciesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPoliciesListParamsWithHTTPClient(client *http.Client) *PoliciesListParams {
	return &PoliciesListParams{
		HTTPClient: client,
	}
}

/*
PoliciesListParams contains all the parameters to send to the API endpoint

	for the policies list operation.

	Typically these are written to a http.Request.
*/
type PoliciesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the policies list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoliciesListParams) WithDefaults() *PoliciesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the policies list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoliciesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the policies list params
func (o *PoliciesListParams) WithTimeout(timeout time.Duration) *PoliciesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policies list params
func (o *PoliciesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policies list params
func (o *PoliciesListParams) WithContext(ctx context.Context) *PoliciesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policies list params
func (o *PoliciesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policies list params
func (o *PoliciesListParams) WithHTTPClient(client *http.Client) *PoliciesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policies list params
func (o *PoliciesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PoliciesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
