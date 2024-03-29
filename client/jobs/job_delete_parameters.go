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
)

// NewJobDeleteParams creates a new JobDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobDeleteParams() *JobDeleteParams {
	return &JobDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobDeleteParamsWithTimeout creates a new JobDeleteParams object
// with the ability to set a timeout on a request.
func NewJobDeleteParamsWithTimeout(timeout time.Duration) *JobDeleteParams {
	return &JobDeleteParams{
		timeout: timeout,
	}
}

// NewJobDeleteParamsWithContext creates a new JobDeleteParams object
// with the ability to set a context for a request.
func NewJobDeleteParamsWithContext(ctx context.Context) *JobDeleteParams {
	return &JobDeleteParams{
		Context: ctx,
	}
}

// NewJobDeleteParamsWithHTTPClient creates a new JobDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobDeleteParamsWithHTTPClient(client *http.Client) *JobDeleteParams {
	return &JobDeleteParams{
		HTTPClient: client,
	}
}

/*
JobDeleteParams contains all the parameters to send to the API endpoint

	for the job delete operation.

	Typically these are written to a http.Request.
*/
type JobDeleteParams struct {

	// Job.
	Job string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobDeleteParams) WithDefaults() *JobDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job delete params
func (o *JobDeleteParams) WithTimeout(timeout time.Duration) *JobDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job delete params
func (o *JobDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job delete params
func (o *JobDeleteParams) WithContext(ctx context.Context) *JobDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job delete params
func (o *JobDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job delete params
func (o *JobDeleteParams) WithHTTPClient(client *http.Client) *JobDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job delete params
func (o *JobDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJob adds the job to the job delete params
func (o *JobDeleteParams) WithJob(job string) *JobDeleteParams {
	o.SetJob(job)
	return o
}

// SetJob adds the job to the job delete params
func (o *JobDeleteParams) SetJob(job string) {
	o.Job = job
}

// WriteToRequest writes these params to a swagger request
func (o *JobDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job
	if err := r.SetPathParam("job", o.Job); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
