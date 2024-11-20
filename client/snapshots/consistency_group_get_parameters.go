// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// NewConsistencyGroupGetParams creates a new ConsistencyGroupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConsistencyGroupGetParams() *ConsistencyGroupGetParams {
	return &ConsistencyGroupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConsistencyGroupGetParamsWithTimeout creates a new ConsistencyGroupGetParams object
// with the ability to set a timeout on a request.
func NewConsistencyGroupGetParamsWithTimeout(timeout time.Duration) *ConsistencyGroupGetParams {
	return &ConsistencyGroupGetParams{
		timeout: timeout,
	}
}

// NewConsistencyGroupGetParamsWithContext creates a new ConsistencyGroupGetParams object
// with the ability to set a context for a request.
func NewConsistencyGroupGetParamsWithContext(ctx context.Context) *ConsistencyGroupGetParams {
	return &ConsistencyGroupGetParams{
		Context: ctx,
	}
}

// NewConsistencyGroupGetParamsWithHTTPClient creates a new ConsistencyGroupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewConsistencyGroupGetParamsWithHTTPClient(client *http.Client) *ConsistencyGroupGetParams {
	return &ConsistencyGroupGetParams{
		HTTPClient: client,
	}
}

/*
ConsistencyGroupGetParams contains all the parameters to send to the API endpoint

	for the consistency group get operation.

	Typically these are written to a http.Request.
*/
type ConsistencyGroupGetParams struct {

	// Authorization.
	Authorization *string

	// SnapshotGroupName.
	SnapshotGroupName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the consistency group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupGetParams) WithDefaults() *ConsistencyGroupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the consistency group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the consistency group get params
func (o *ConsistencyGroupGetParams) WithTimeout(timeout time.Duration) *ConsistencyGroupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consistency group get params
func (o *ConsistencyGroupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consistency group get params
func (o *ConsistencyGroupGetParams) WithContext(ctx context.Context) *ConsistencyGroupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consistency group get params
func (o *ConsistencyGroupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consistency group get params
func (o *ConsistencyGroupGetParams) WithHTTPClient(client *http.Client) *ConsistencyGroupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consistency group get params
func (o *ConsistencyGroupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the consistency group get params
func (o *ConsistencyGroupGetParams) WithAuthorization(authorization *string) *ConsistencyGroupGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the consistency group get params
func (o *ConsistencyGroupGetParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithSnapshotGroupName adds the snapshotGroupName to the consistency group get params
func (o *ConsistencyGroupGetParams) WithSnapshotGroupName(snapshotGroupName string) *ConsistencyGroupGetParams {
	o.SetSnapshotGroupName(snapshotGroupName)
	return o
}

// SetSnapshotGroupName adds the snapshotGroupName to the consistency group get params
func (o *ConsistencyGroupGetParams) SetSnapshotGroupName(snapshotGroupName string) {
	o.SnapshotGroupName = snapshotGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *ConsistencyGroupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param snapshot_group_name
	if err := r.SetPathParam("snapshot_group_name", o.SnapshotGroupName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
