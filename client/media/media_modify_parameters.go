// Code generated by go-swagger; DO NOT EDIT.

package media

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

// NewMediaModifyParams creates a new MediaModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaModifyParams() *MediaModifyParams {
	return &MediaModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaModifyParamsWithTimeout creates a new MediaModifyParams object
// with the ability to set a timeout on a request.
func NewMediaModifyParamsWithTimeout(timeout time.Duration) *MediaModifyParams {
	return &MediaModifyParams{
		timeout: timeout,
	}
}

// NewMediaModifyParamsWithContext creates a new MediaModifyParams object
// with the ability to set a context for a request.
func NewMediaModifyParamsWithContext(ctx context.Context) *MediaModifyParams {
	return &MediaModifyParams{
		Context: ctx,
	}
}

// NewMediaModifyParamsWithHTTPClient creates a new MediaModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaModifyParamsWithHTTPClient(client *http.Client) *MediaModifyParams {
	return &MediaModifyParams{
		HTTPClient: client,
	}
}

/*
MediaModifyParams contains all the parameters to send to the API endpoint

	for the media modify operation.

	Typically these are written to a http.Request.
*/
type MediaModifyParams struct {

	// Authorization.
	Authorization *string

	/* Body.

	   A Media Modify Object
	*/
	Body *models.MediaModify

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaModifyParams) WithDefaults() *MediaModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media modify params
func (o *MediaModifyParams) WithTimeout(timeout time.Duration) *MediaModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media modify params
func (o *MediaModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media modify params
func (o *MediaModifyParams) WithContext(ctx context.Context) *MediaModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media modify params
func (o *MediaModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media modify params
func (o *MediaModifyParams) WithHTTPClient(client *http.Client) *MediaModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media modify params
func (o *MediaModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media modify params
func (o *MediaModifyParams) WithAuthorization(authorization *string) *MediaModifyParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media modify params
func (o *MediaModifyParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the media modify params
func (o *MediaModifyParams) WithBody(body *models.MediaModify) *MediaModifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the media modify params
func (o *MediaModifyParams) SetBody(body *models.MediaModify) {
	o.Body = body
}

// WithMedia adds the media to the media modify params
func (o *MediaModifyParams) WithMedia(media string) *MediaModifyParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media modify params
func (o *MediaModifyParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param media
	if err := r.SetPathParam("media", o.Media); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
