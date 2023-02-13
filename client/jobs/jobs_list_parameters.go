// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// NewJobsListParams creates a new JobsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobsListParams() *JobsListParams {
	return &JobsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobsListParamsWithTimeout creates a new JobsListParams object
// with the ability to set a timeout on a request.
func NewJobsListParamsWithTimeout(timeout time.Duration) *JobsListParams {
	return &JobsListParams{
		timeout: timeout,
	}
}

// NewJobsListParamsWithContext creates a new JobsListParams object
// with the ability to set a context for a request.
func NewJobsListParamsWithContext(ctx context.Context) *JobsListParams {
	return &JobsListParams{
		Context: ctx,
	}
}

// NewJobsListParamsWithHTTPClient creates a new JobsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobsListParamsWithHTTPClient(client *http.Client) *JobsListParams {
	return &JobsListParams{
		HTTPClient: client,
	}
}

/* JobsListParams contains all the parameters to send to the API endpoint
   for the jobs list operation.

   Typically these are written to a http.Request.
*/
type JobsListParams struct {

	// Count.
	Count *int64

	// Internal.
	Internal *bool

	// Page.
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobsListParams) WithDefaults() *JobsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the jobs list params
func (o *JobsListParams) WithTimeout(timeout time.Duration) *JobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs list params
func (o *JobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs list params
func (o *JobsListParams) WithContext(ctx context.Context) *JobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs list params
func (o *JobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs list params
func (o *JobsListParams) WithHTTPClient(client *http.Client) *JobsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs list params
func (o *JobsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the jobs list params
func (o *JobsListParams) WithCount(count *int64) *JobsListParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the jobs list params
func (o *JobsListParams) SetCount(count *int64) {
	o.Count = count
}

// WithInternal adds the internal to the jobs list params
func (o *JobsListParams) WithInternal(internal *bool) *JobsListParams {
	o.SetInternal(internal)
	return o
}

// SetInternal adds the internal to the jobs list params
func (o *JobsListParams) SetInternal(internal *bool) {
	o.Internal = internal
}

// WithPage adds the page to the jobs list params
func (o *JobsListParams) WithPage(page *int64) *JobsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the jobs list params
func (o *JobsListParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *JobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Internal != nil {

		// query param internal
		var qrInternal bool

		if o.Internal != nil {
			qrInternal = *o.Internal
		}
		qInternal := swag.FormatBool(qrInternal)
		if qInternal != "" {

			if err := r.SetQueryParam("internal", qInternal); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
