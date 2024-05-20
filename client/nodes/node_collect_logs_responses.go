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

// NodeCollectLogsReader is a Reader for the NodeCollectLogs structure.
type NodeCollectLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeCollectLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeCollectLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNodeCollectLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNodeCollectLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNodeCollectLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeCollectLogsOK creates a NodeCollectLogsOK with default headers values
func NewNodeCollectLogsOK() *NodeCollectLogsOK {
	return &NodeCollectLogsOK{}
}

/*
NodeCollectLogsOK describes a response with status code 200, with default header values.

Node collect logs job  was created successfully
*/
type NodeCollectLogsOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this node collect logs o k response has a 2xx status code
func (o *NodeCollectLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node collect logs o k response has a 3xx status code
func (o *NodeCollectLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node collect logs o k response has a 4xx status code
func (o *NodeCollectLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node collect logs o k response has a 5xx status code
func (o *NodeCollectLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node collect logs o k response a status code equal to that given
func (o *NodeCollectLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node collect logs o k response
func (o *NodeCollectLogsOK) Code() int {
	return 200
}

func (o *NodeCollectLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsOK %s", 200, payload)
}

func (o *NodeCollectLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsOK %s", 200, payload)
}

func (o *NodeCollectLogsOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *NodeCollectLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCollectLogsBadRequest creates a NodeCollectLogsBadRequest with default headers values
func NewNodeCollectLogsBadRequest() *NodeCollectLogsBadRequest {
	return &NodeCollectLogsBadRequest{}
}

/*
NodeCollectLogsBadRequest describes a response with status code 400, with default header values.

Invalid node name supplied
*/
type NodeCollectLogsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node collect logs bad request response has a 2xx status code
func (o *NodeCollectLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node collect logs bad request response has a 3xx status code
func (o *NodeCollectLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node collect logs bad request response has a 4xx status code
func (o *NodeCollectLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this node collect logs bad request response has a 5xx status code
func (o *NodeCollectLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this node collect logs bad request response a status code equal to that given
func (o *NodeCollectLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the node collect logs bad request response
func (o *NodeCollectLogsBadRequest) Code() int {
	return 400
}

func (o *NodeCollectLogsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsBadRequest %s", 400, payload)
}

func (o *NodeCollectLogsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsBadRequest %s", 400, payload)
}

func (o *NodeCollectLogsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCollectLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCollectLogsNotFound creates a NodeCollectLogsNotFound with default headers values
func NewNodeCollectLogsNotFound() *NodeCollectLogsNotFound {
	return &NodeCollectLogsNotFound{}
}

/*
NodeCollectLogsNotFound describes a response with status code 404, with default header values.

Node not found
*/
type NodeCollectLogsNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node collect logs not found response has a 2xx status code
func (o *NodeCollectLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node collect logs not found response has a 3xx status code
func (o *NodeCollectLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node collect logs not found response has a 4xx status code
func (o *NodeCollectLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this node collect logs not found response has a 5xx status code
func (o *NodeCollectLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this node collect logs not found response a status code equal to that given
func (o *NodeCollectLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the node collect logs not found response
func (o *NodeCollectLogsNotFound) Code() int {
	return 404
}

func (o *NodeCollectLogsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsNotFound %s", 404, payload)
}

func (o *NodeCollectLogsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] nodeCollectLogsNotFound %s", 404, payload)
}

func (o *NodeCollectLogsNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCollectLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCollectLogsDefault creates a NodeCollectLogsDefault with default headers values
func NewNodeCollectLogsDefault(code int) *NodeCollectLogsDefault {
	return &NodeCollectLogsDefault{
		_statusCode: code,
	}
}

/*
NodeCollectLogsDefault describes a response with status code -1, with default header values.

Error trying to create node collect logs job
*/
type NodeCollectLogsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node collect logs default response has a 2xx status code
func (o *NodeCollectLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node collect logs default response has a 3xx status code
func (o *NodeCollectLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node collect logs default response has a 4xx status code
func (o *NodeCollectLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node collect logs default response has a 5xx status code
func (o *NodeCollectLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node collect logs default response a status code equal to that given
func (o *NodeCollectLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node collect logs default response
func (o *NodeCollectLogsDefault) Code() int {
	return o._statusCode
}

func (o *NodeCollectLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] NodeCollectLogs default %s", o._statusCode, payload)
}

func (o *NodeCollectLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/logs/{node}/{tenant}][%d] NodeCollectLogs default %s", o._statusCode, payload)
}

func (o *NodeCollectLogsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeCollectLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
