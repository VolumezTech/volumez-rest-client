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

// NodeCleanUpReader is a Reader for the NodeCleanUp structure.
type NodeCleanUpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeCleanUpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeCleanUpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNodeCleanUpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNodeCleanUpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNodeCleanUpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeCleanUpOK creates a NodeCleanUpOK with default headers values
func NewNodeCleanUpOK() *NodeCleanUpOK {
	return &NodeCleanUpOK{}
}

/* NodeCleanUpOK describes a response with status code 200, with default header values.

Node cleanup started successfully
*/
type NodeCleanUpOK struct {
	Payload *models.RegularResponse
}

func (o *NodeCleanUpOK) Error() string {
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpOK  %+v", 200, o.Payload)
}
func (o *NodeCleanUpOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *NodeCleanUpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCleanUpBadRequest creates a NodeCleanUpBadRequest with default headers values
func NewNodeCleanUpBadRequest() *NodeCleanUpBadRequest {
	return &NodeCleanUpBadRequest{}
}

/* NodeCleanUpBadRequest describes a response with status code 400, with default header values.

Invalid node name supplied
*/
type NodeCleanUpBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *NodeCleanUpBadRequest) Error() string {
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpBadRequest  %+v", 400, o.Payload)
}
func (o *NodeCleanUpBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCleanUpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCleanUpNotFound creates a NodeCleanUpNotFound with default headers values
func NewNodeCleanUpNotFound() *NodeCleanUpNotFound {
	return &NodeCleanUpNotFound{}
}

/* NodeCleanUpNotFound describes a response with status code 404, with default header values.

Node not found
*/
type NodeCleanUpNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NodeCleanUpNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpNotFound  %+v", 404, o.Payload)
}
func (o *NodeCleanUpNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCleanUpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCleanUpDefault creates a NodeCleanUpDefault with default headers values
func NewNodeCleanUpDefault(code int) *NodeCleanUpDefault {
	return &NodeCleanUpDefault{
		_statusCode: code,
	}
}

/* NodeCleanUpDefault describes a response with status code -1, with default header values.

Error trying to cleanup node
*/
type NodeCleanUpDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the node clean up default response
func (o *NodeCleanUpDefault) Code() int {
	return o._statusCode
}

func (o *NodeCleanUpDefault) Error() string {
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] NodeCleanUp default  %+v", o._statusCode, o.Payload)
}
func (o *NodeCleanUpDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCleanUpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
