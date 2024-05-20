// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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

/*
NodeCleanUpOK describes a response with status code 200, with default header values.

Node cleanup started successfully
*/
type NodeCleanUpOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this node clean up o k response has a 2xx status code
func (o *NodeCleanUpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node clean up o k response has a 3xx status code
func (o *NodeCleanUpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node clean up o k response has a 4xx status code
func (o *NodeCleanUpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node clean up o k response has a 5xx status code
func (o *NodeCleanUpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node clean up o k response a status code equal to that given
func (o *NodeCleanUpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node clean up o k response
func (o *NodeCleanUpOK) Code() int {
	return 200
}

func (o *NodeCleanUpOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpOK %s", 200, payload)
}

func (o *NodeCleanUpOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpOK %s", 200, payload)
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

/*
NodeCleanUpBadRequest describes a response with status code 400, with default header values.

Invalid node name supplied
*/
type NodeCleanUpBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node clean up bad request response has a 2xx status code
func (o *NodeCleanUpBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node clean up bad request response has a 3xx status code
func (o *NodeCleanUpBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node clean up bad request response has a 4xx status code
func (o *NodeCleanUpBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this node clean up bad request response has a 5xx status code
func (o *NodeCleanUpBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this node clean up bad request response a status code equal to that given
func (o *NodeCleanUpBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the node clean up bad request response
func (o *NodeCleanUpBadRequest) Code() int {
	return 400
}

func (o *NodeCleanUpBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpBadRequest %s", 400, payload)
}

func (o *NodeCleanUpBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpBadRequest %s", 400, payload)
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

/*
NodeCleanUpNotFound describes a response with status code 404, with default header values.

Node not found
*/
type NodeCleanUpNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node clean up not found response has a 2xx status code
func (o *NodeCleanUpNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node clean up not found response has a 3xx status code
func (o *NodeCleanUpNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node clean up not found response has a 4xx status code
func (o *NodeCleanUpNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this node clean up not found response has a 5xx status code
func (o *NodeCleanUpNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this node clean up not found response a status code equal to that given
func (o *NodeCleanUpNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the node clean up not found response
func (o *NodeCleanUpNotFound) Code() int {
	return 404
}

func (o *NodeCleanUpNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpNotFound %s", 404, payload)
}

func (o *NodeCleanUpNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] nodeCleanUpNotFound %s", 404, payload)
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

/*
NodeCleanUpDefault describes a response with status code -1, with default header values.

Error trying to cleanup node
*/
type NodeCleanUpDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node clean up default response has a 2xx status code
func (o *NodeCleanUpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node clean up default response has a 3xx status code
func (o *NodeCleanUpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node clean up default response has a 4xx status code
func (o *NodeCleanUpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node clean up default response has a 5xx status code
func (o *NodeCleanUpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node clean up default response a status code equal to that given
func (o *NodeCleanUpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node clean up default response
func (o *NodeCleanUpDefault) Code() int {
	return o._statusCode
}

func (o *NodeCleanUpDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] NodeCleanUp default %s", o._statusCode, payload)
}

func (o *NodeCleanUpDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/cleanup/{node}][%d] NodeCleanUp default %s", o._statusCode, payload)
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
