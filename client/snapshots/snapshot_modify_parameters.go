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

	"bitbucket.org/volumez/volumez/test/restapi/models"
)

// NewSnapshotModifyParams creates a new SnapshotModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotModifyParams() *SnapshotModifyParams {
	return &SnapshotModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotModifyParamsWithTimeout creates a new SnapshotModifyParams object
// with the ability to set a timeout on a request.
func NewSnapshotModifyParamsWithTimeout(timeout time.Duration) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		timeout: timeout,
	}
}

// NewSnapshotModifyParamsWithContext creates a new SnapshotModifyParams object
// with the ability to set a context for a request.
func NewSnapshotModifyParamsWithContext(ctx context.Context) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		Context: ctx,
	}
}

// NewSnapshotModifyParamsWithHTTPClient creates a new SnapshotModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotModifyParamsWithHTTPClient(client *http.Client) *SnapshotModifyParams {
	return &SnapshotModifyParams{
		HTTPClient: client,
	}
}

/* SnapshotModifyParams contains all the parameters to send to the API endpoint
   for the snapshot modify operation.

   Typically these are written to a http.Request.
*/
type SnapshotModifyParams struct {

	/* Body.

	   A Snapshot object
	*/
	Body *models.Snapshot

	// Snapshot.
	Snapshot string

	// Volume.
	Volume string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotModifyParams) WithDefaults() *SnapshotModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot modify params
func (o *SnapshotModifyParams) WithTimeout(timeout time.Duration) *SnapshotModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot modify params
func (o *SnapshotModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot modify params
func (o *SnapshotModifyParams) WithContext(ctx context.Context) *SnapshotModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot modify params
func (o *SnapshotModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot modify params
func (o *SnapshotModifyParams) WithHTTPClient(client *http.Client) *SnapshotModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot modify params
func (o *SnapshotModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the snapshot modify params
func (o *SnapshotModifyParams) WithBody(body *models.Snapshot) *SnapshotModifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the snapshot modify params
func (o *SnapshotModifyParams) SetBody(body *models.Snapshot) {
	o.Body = body
}

// WithSnapshot adds the snapshot to the snapshot modify params
func (o *SnapshotModifyParams) WithSnapshot(snapshot string) *SnapshotModifyParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the snapshot modify params
func (o *SnapshotModifyParams) SetSnapshot(snapshot string) {
	o.Snapshot = snapshot
}

// WithVolume adds the volume to the snapshot modify params
func (o *SnapshotModifyParams) WithVolume(volume string) *SnapshotModifyParams {
	o.SetVolume(volume)
	return o
}

// SetVolume adds the volume to the snapshot modify params
func (o *SnapshotModifyParams) SetVolume(volume string) {
	o.Volume = volume
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param snapshot
	if err := r.SetPathParam("snapshot", o.Snapshot); err != nil {
		return err
	}

	// path param volume
	if err := r.SetPathParam("volume", o.Volume); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
