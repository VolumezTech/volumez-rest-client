// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewNodeDeleteParams creates a new NodeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeDeleteParams() *NodeDeleteParams {
	return &NodeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeDeleteParamsWithTimeout creates a new NodeDeleteParams object
// with the ability to set a timeout on a request.
func NewNodeDeleteParamsWithTimeout(timeout time.Duration) *NodeDeleteParams {
	return &NodeDeleteParams{
		timeout: timeout,
	}
}

// NewNodeDeleteParamsWithContext creates a new NodeDeleteParams object
// with the ability to set a context for a request.
func NewNodeDeleteParamsWithContext(ctx context.Context) *NodeDeleteParams {
	return &NodeDeleteParams{
		Context: ctx,
	}
}

// NewNodeDeleteParamsWithHTTPClient creates a new NodeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeDeleteParamsWithHTTPClient(client *http.Client) *NodeDeleteParams {
	return &NodeDeleteParams{
		HTTPClient: client,
	}
}

/* NodeDeleteParams contains all the parameters to send to the API endpoint
   for the node delete operation.

   Typically these are written to a http.Request.
*/
type NodeDeleteParams struct {

	// Force.
	Force *bool

	/* Node.

	   node to delete
	*/
	Node string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeDeleteParams) WithDefaults() *NodeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeDeleteParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := NodeDeleteParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) WithTimeout(timeout time.Duration) *NodeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node delete params
func (o *NodeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node delete params
func (o *NodeDeleteParams) WithContext(ctx context.Context) *NodeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node delete params
func (o *NodeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) WithHTTPClient(client *http.Client) *NodeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node delete params
func (o *NodeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the node delete params
func (o *NodeDeleteParams) WithForce(force *bool) *NodeDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the node delete params
func (o *NodeDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithNode adds the node to the node delete params
func (o *NodeDeleteParams) WithNode(node string) *NodeDeleteParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the node delete params
func (o *NodeDeleteParams) SetNode(node string) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *NodeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
