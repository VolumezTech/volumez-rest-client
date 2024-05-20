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

// NodeUpgradeReader is a Reader for the NodeUpgrade structure.
type NodeUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNodeUpgradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNodeUpgradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewNodeUpgradeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeUpgradeOK creates a NodeUpgradeOK with default headers values
func NewNodeUpgradeOK() *NodeUpgradeOK {
	return &NodeUpgradeOK{}
}

/*
NodeUpgradeOK describes a response with status code 200, with default header values.

Node upgrade started successfully
*/
type NodeUpgradeOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this node upgrade o k response has a 2xx status code
func (o *NodeUpgradeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node upgrade o k response has a 3xx status code
func (o *NodeUpgradeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node upgrade o k response has a 4xx status code
func (o *NodeUpgradeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node upgrade o k response has a 5xx status code
func (o *NodeUpgradeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node upgrade o k response a status code equal to that given
func (o *NodeUpgradeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node upgrade o k response
func (o *NodeUpgradeOK) Code() int {
	return 200
}

func (o *NodeUpgradeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeOK %s", 200, payload)
}

func (o *NodeUpgradeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeOK %s", 200, payload)
}

func (o *NodeUpgradeOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *NodeUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeUpgradeBadRequest creates a NodeUpgradeBadRequest with default headers values
func NewNodeUpgradeBadRequest() *NodeUpgradeBadRequest {
	return &NodeUpgradeBadRequest{}
}

/*
NodeUpgradeBadRequest describes a response with status code 400, with default header values.

Invalid node name supplied
*/
type NodeUpgradeBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node upgrade bad request response has a 2xx status code
func (o *NodeUpgradeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node upgrade bad request response has a 3xx status code
func (o *NodeUpgradeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node upgrade bad request response has a 4xx status code
func (o *NodeUpgradeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this node upgrade bad request response has a 5xx status code
func (o *NodeUpgradeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this node upgrade bad request response a status code equal to that given
func (o *NodeUpgradeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the node upgrade bad request response
func (o *NodeUpgradeBadRequest) Code() int {
	return 400
}

func (o *NodeUpgradeBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeBadRequest %s", 400, payload)
}

func (o *NodeUpgradeBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeBadRequest %s", 400, payload)
}

func (o *NodeUpgradeBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeUpgradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeUpgradeNotFound creates a NodeUpgradeNotFound with default headers values
func NewNodeUpgradeNotFound() *NodeUpgradeNotFound {
	return &NodeUpgradeNotFound{}
}

/*
NodeUpgradeNotFound describes a response with status code 404, with default header values.

Node not found
*/
type NodeUpgradeNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node upgrade not found response has a 2xx status code
func (o *NodeUpgradeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node upgrade not found response has a 3xx status code
func (o *NodeUpgradeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node upgrade not found response has a 4xx status code
func (o *NodeUpgradeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this node upgrade not found response has a 5xx status code
func (o *NodeUpgradeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this node upgrade not found response a status code equal to that given
func (o *NodeUpgradeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the node upgrade not found response
func (o *NodeUpgradeNotFound) Code() int {
	return 404
}

func (o *NodeUpgradeNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeNotFound %s", 404, payload)
}

func (o *NodeUpgradeNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] nodeUpgradeNotFound %s", 404, payload)
}

func (o *NodeUpgradeNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeUpgradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeUpgradeDefault creates a NodeUpgradeDefault with default headers values
func NewNodeUpgradeDefault(code int) *NodeUpgradeDefault {
	return &NodeUpgradeDefault{
		_statusCode: code,
	}
}

/*
NodeUpgradeDefault describes a response with status code -1, with default header values.

Error trying to upgrade node
*/
type NodeUpgradeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node upgrade default response has a 2xx status code
func (o *NodeUpgradeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node upgrade default response has a 3xx status code
func (o *NodeUpgradeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node upgrade default response has a 4xx status code
func (o *NodeUpgradeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node upgrade default response has a 5xx status code
func (o *NodeUpgradeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node upgrade default response a status code equal to that given
func (o *NodeUpgradeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node upgrade default response
func (o *NodeUpgradeDefault) Code() int {
	return o._statusCode
}

func (o *NodeUpgradeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] NodeUpgrade default %s", o._statusCode, payload)
}

func (o *NodeUpgradeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/upgrade/{node}][%d] NodeUpgrade default %s", o._statusCode, payload)
}

func (o *NodeUpgradeDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeUpgradeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
