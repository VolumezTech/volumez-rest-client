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
)

// NewAttachmentsListAllParams creates a new AttachmentsListAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttachmentsListAllParams() *AttachmentsListAllParams {
	return &AttachmentsListAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttachmentsListAllParamsWithTimeout creates a new AttachmentsListAllParams object
// with the ability to set a timeout on a request.
func NewAttachmentsListAllParamsWithTimeout(timeout time.Duration) *AttachmentsListAllParams {
	return &AttachmentsListAllParams{
		timeout: timeout,
	}
}

// NewAttachmentsListAllParamsWithContext creates a new AttachmentsListAllParams object
// with the ability to set a context for a request.
func NewAttachmentsListAllParamsWithContext(ctx context.Context) *AttachmentsListAllParams {
	return &AttachmentsListAllParams{
		Context: ctx,
	}
}

// NewAttachmentsListAllParamsWithHTTPClient creates a new AttachmentsListAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttachmentsListAllParamsWithHTTPClient(client *http.Client) *AttachmentsListAllParams {
	return &AttachmentsListAllParams{
		HTTPClient: client,
	}
}

/* AttachmentsListAllParams contains all the parameters to send to the API endpoint
   for the attachments list all operation.

   Typically these are written to a http.Request.
*/
type AttachmentsListAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attachments list all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentsListAllParams) WithDefaults() *AttachmentsListAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attachments list all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttachmentsListAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attachments list all params
func (o *AttachmentsListAllParams) WithTimeout(timeout time.Duration) *AttachmentsListAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attachments list all params
func (o *AttachmentsListAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attachments list all params
func (o *AttachmentsListAllParams) WithContext(ctx context.Context) *AttachmentsListAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attachments list all params
func (o *AttachmentsListAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attachments list all params
func (o *AttachmentsListAllParams) WithHTTPClient(client *http.Client) *AttachmentsListAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attachments list all params
func (o *AttachmentsListAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AttachmentsListAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
