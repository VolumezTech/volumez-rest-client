// Code generated by go-swagger; DO NOT EDIT.

package associations

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
	"github.com/go-openapi/swag"
)

// NewAssociationsListParams creates a new AssociationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociationsListParams() *AssociationsListParams {
	return &AssociationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociationsListParamsWithTimeout creates a new AssociationsListParams object
// with the ability to set a timeout on a request.
func NewAssociationsListParamsWithTimeout(timeout time.Duration) *AssociationsListParams {
	return &AssociationsListParams{
		timeout: timeout,
	}
}

// NewAssociationsListParamsWithContext creates a new AssociationsListParams object
// with the ability to set a context for a request.
func NewAssociationsListParamsWithContext(ctx context.Context) *AssociationsListParams {
	return &AssociationsListParams{
		Context: ctx,
	}
}

// NewAssociationsListParamsWithHTTPClient creates a new AssociationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociationsListParamsWithHTTPClient(client *http.Client) *AssociationsListParams {
	return &AssociationsListParams{
		HTTPClient: client,
	}
}

/*
AssociationsListParams contains all the parameters to send to the API endpoint

	for the associations list operation.

	Typically these are written to a http.Request.
*/
type AssociationsListParams struct {

	// Authorization.
	Authorization *string

	// Count.
	Count *int64

	// Startfrom.
	Startfrom *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsListParams) WithDefaults() *AssociationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associations list params
func (o *AssociationsListParams) WithTimeout(timeout time.Duration) *AssociationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associations list params
func (o *AssociationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associations list params
func (o *AssociationsListParams) WithContext(ctx context.Context) *AssociationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associations list params
func (o *AssociationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associations list params
func (o *AssociationsListParams) WithHTTPClient(client *http.Client) *AssociationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associations list params
func (o *AssociationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the associations list params
func (o *AssociationsListParams) WithAuthorization(authorization *string) *AssociationsListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the associations list params
func (o *AssociationsListParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithCount adds the count to the associations list params
func (o *AssociationsListParams) WithCount(count *int64) *AssociationsListParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the associations list params
func (o *AssociationsListParams) SetCount(count *int64) {
	o.Count = count
}

// WithStartfrom adds the startfrom to the associations list params
func (o *AssociationsListParams) WithStartfrom(startfrom *string) *AssociationsListParams {
	o.SetStartfrom(startfrom)
	return o
}

// SetStartfrom adds the startfrom to the associations list params
func (o *AssociationsListParams) SetStartfrom(startfrom *string) {
	o.Startfrom = startfrom
}

// WriteToRequest writes these params to a swagger request
func (o *AssociationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Startfrom != nil {

		// query param startfrom
		var qrStartfrom string

		if o.Startfrom != nil {
			qrStartfrom = *o.Startfrom
		}
		qStartfrom := qrStartfrom
		if qStartfrom != "" {

			if err := r.SetQueryParam("startfrom", qStartfrom); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
