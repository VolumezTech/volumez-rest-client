// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

// NewAttachmentDeleteParams creates a new AttachmentDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachmentDeleteParams() *AttachmentDeleteParams {
	return &AttachmentDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachmentDeleteParamsWithTimeout creates a new AttachmentDeleteParams object
// with the ability to set a timeout on a request.
func NewAttachmentDeleteParamsWithTimeout(timeout time.Duration) *AttachmentDeleteParams {
	return &AttachmentDeleteParams{
		timeout: timeout,
	}
}

// NewAttachmentDeleteParamsWithContext creates a new AttachmentDeleteParams object
// with the ability to set a context for a request.
func NewAttachmentDeleteParamsWithContext(ctx context.Context) *AttachmentDeleteParams {
	return &AttachmentDeleteParams{
		Context: ctx,
	}
}

// NewAttachmentDeleteParamsWithHTTPClient creates a new AttachmentDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachmentDeleteParamsWithHTTPClient(client *http.Client) *AttachmentDeleteParams {
	return &AttachmentDeleteParams{
		HTTPClient: client,
	}
}

/* AttachmentDeleteParams contains all the parameters to send to the API endpoint
   for the attachment delete operation.

   Typically these are written to a http.Request.
*/
type AttachmentDeleteParams struct {

	// Authorization.
	Authorization *string

	// Force.
	Force bool

	// Node.
	Node string

	// Snapshot.
	Snapshot string

	// Volume.
	Volume string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attachment delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentDeleteParams) WithDefaults() *AttachmentDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attachment delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentDeleteParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := AttachmentDeleteParams{
		Force: forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the attachment delete params
func (o *AttachmentDeleteParams) WithTimeout(timeout time.Duration) *AttachmentDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attachment delete params
func (o *AttachmentDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attachment delete params
func (o *AttachmentDeleteParams) WithContext(ctx context.Context) *AttachmentDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attachment delete params
func (o *AttachmentDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attachment delete params
func (o *AttachmentDeleteParams) WithHTTPClient(client *http.Client) *AttachmentDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attachment delete params
func (o *AttachmentDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the attachment delete params
func (o *AttachmentDeleteParams) WithAuthorization(authorization *string) *AttachmentDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the attachment delete params
func (o *AttachmentDeleteParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithForce adds the force to the attachment delete params
func (o *AttachmentDeleteParams) WithForce(force bool) *AttachmentDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the attachment delete params
func (o *AttachmentDeleteParams) SetForce(force bool) {
	o.Force = force
}

// WithNode adds the node to the attachment delete params
func (o *AttachmentDeleteParams) WithNode(node string) *AttachmentDeleteParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the attachment delete params
func (o *AttachmentDeleteParams) SetNode(node string) {
	o.Node = node
}

// WithSnapshot adds the snapshot to the attachment delete params
func (o *AttachmentDeleteParams) WithSnapshot(snapshot string) *AttachmentDeleteParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the attachment delete params
func (o *AttachmentDeleteParams) SetSnapshot(snapshot string) {
	o.Snapshot = snapshot
}

// WithVolume adds the volume to the attachment delete params
func (o *AttachmentDeleteParams) WithVolume(volume string) *AttachmentDeleteParams {
	o.SetVolume(volume)
	return o
}

// SetVolume adds the volume to the attachment delete params
func (o *AttachmentDeleteParams) SetVolume(volume string) {
	o.Volume = volume
}

// WriteToRequest writes these params to a swagger request
func (o *AttachmentDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param force
	qrForce := o.Force
	qForce := swag.FormatBool(qrForce)

	if err := r.SetQueryParam("force", qForce); err != nil {
		return err
	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
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
