// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NodeGetReader is a Reader for the NodeGet structure.
type NodeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNodeGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNodeGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNodeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeGetOK creates a NodeGetOK with default headers values
func NewNodeGetOK() *NodeGetOK {
	return &NodeGetOK{}
}

/*
	NodeGetOK describes a response with status code 200, with default header values.

Properties of a node
*/
type NodeGetOK struct {
	Payload *models.Node
}

func (o *NodeGetOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{node}][%d] nodeGetOK  %+v", 200, o.Payload)
}
func (o *NodeGetOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *NodeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeGetBadRequest creates a NodeGetBadRequest with default headers values
func NewNodeGetBadRequest() *NodeGetBadRequest {
	return &NodeGetBadRequest{}
}

/*
	NodeGetBadRequest describes a response with status code 400, with default header values.

Invalid node name supplied
*/
type NodeGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *NodeGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /nodes/{node}][%d] nodeGetBadRequest  %+v", 400, o.Payload)
}
func (o *NodeGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeGetNotFound creates a NodeGetNotFound with default headers values
func NewNodeGetNotFound() *NodeGetNotFound {
	return &NodeGetNotFound{}
}

/*
	NodeGetNotFound describes a response with status code 404, with default header values.

Node not found
*/
type NodeGetNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NodeGetNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{node}][%d] nodeGetNotFound  %+v", 404, o.Payload)
}
func (o *NodeGetNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeGetDefault creates a NodeGetDefault with default headers values
func NewNodeGetDefault(code int) *NodeGetDefault {
	return &NodeGetDefault{
		_statusCode: code,
	}
}

/*
	NodeGetDefault describes a response with status code -1, with default header values.

Error getting properties of a node
*/
type NodeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the node get default response
func (o *NodeGetDefault) Code() int {
	return o._statusCode
}

func (o *NodeGetDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{node}][%d] NodeGet default  %+v", o._statusCode, o.Payload)
}
func (o *NodeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
