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

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NewAttachmentModifyParams creates a new AttachmentModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachmentModifyParams() *AttachmentModifyParams {
	return &AttachmentModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachmentModifyParamsWithTimeout creates a new AttachmentModifyParams object
// with the ability to set a timeout on a request.
func NewAttachmentModifyParamsWithTimeout(timeout time.Duration) *AttachmentModifyParams {
	return &AttachmentModifyParams{
		timeout: timeout,
	}
}

// NewAttachmentModifyParamsWithContext creates a new AttachmentModifyParams object
// with the ability to set a context for a request.
func NewAttachmentModifyParamsWithContext(ctx context.Context) *AttachmentModifyParams {
	return &AttachmentModifyParams{
		Context: ctx,
	}
}

// NewAttachmentModifyParamsWithHTTPClient creates a new AttachmentModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachmentModifyParamsWithHTTPClient(client *http.Client) *AttachmentModifyParams {
	return &AttachmentModifyParams{
		HTTPClient: client,
	}
}

/*
AttachmentModifyParams contains all the parameters to send to the API endpoint

	for the attachment modify operation.

	Typically these are written to a http.Request.
*/
type AttachmentModifyParams struct {

	/* Body.

	   An Attachment object
	*/
	Body *models.Attachment

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

// WithDefaults hydrates default values in the attachment modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentModifyParams) WithDefaults() *AttachmentModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attachment modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attachment modify params
func (o *AttachmentModifyParams) WithTimeout(timeout time.Duration) *AttachmentModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attachment modify params
func (o *AttachmentModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attachment modify params
func (o *AttachmentModifyParams) WithContext(ctx context.Context) *AttachmentModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attachment modify params
func (o *AttachmentModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attachment modify params
func (o *AttachmentModifyParams) WithHTTPClient(client *http.Client) *AttachmentModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attachment modify params
func (o *AttachmentModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the attachment modify params
func (o *AttachmentModifyParams) WithBody(body *models.Attachment) *AttachmentModifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the attachment modify params
func (o *AttachmentModifyParams) SetBody(body *models.Attachment) {
	o.Body = body
}

// WithNode adds the node to the attachment modify params
func (o *AttachmentModifyParams) WithNode(node string) *AttachmentModifyParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the attachment modify params
func (o *AttachmentModifyParams) SetNode(node string) {
	o.Node = node
}

// WithSnapshot adds the snapshot to the attachment modify params
func (o *AttachmentModifyParams) WithSnapshot(snapshot string) *AttachmentModifyParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the attachment modify params
func (o *AttachmentModifyParams) SetSnapshot(snapshot string) {
	o.Snapshot = snapshot
}

// WithVolume adds the volume to the attachment modify params
func (o *AttachmentModifyParams) WithVolume(volume string) *AttachmentModifyParams {
	o.SetVolume(volume)
	return o
}

// SetVolume adds the volume to the attachment modify params
func (o *AttachmentModifyParams) SetVolume(volume string) {
	o.Volume = volume
}

// WriteToRequest writes these params to a swagger request
func (o *AttachmentModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
