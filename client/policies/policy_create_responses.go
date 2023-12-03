// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// PolicyCreateReader is a Reader for the PolicyCreate structure.
type PolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyCreateOK creates a PolicyCreateOK with default headers values
func NewPolicyCreateOK() *PolicyCreateOK {
	return &PolicyCreateOK{}
}

/*
PolicyCreateOK describes a response with status code 200, with default header values.

New policy was created successfully
*/
type PolicyCreateOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this policy create o k response has a 2xx status code
func (o *PolicyCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy create o k response has a 3xx status code
func (o *PolicyCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy create o k response has a 4xx status code
func (o *PolicyCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy create o k response has a 5xx status code
func (o *PolicyCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy create o k response a status code equal to that given
func (o *PolicyCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the policy create o k response
func (o *PolicyCreateOK) Code() int {
	return 200
}

func (o *PolicyCreateOK) Error() string {
	return fmt.Sprintf("[POST /policies][%d] policyCreateOK  %+v", 200, o.Payload)
}

func (o *PolicyCreateOK) String() string {
	return fmt.Sprintf("[POST /policies][%d] policyCreateOK  %+v", 200, o.Payload)
}

func (o *PolicyCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *PolicyCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyCreateDefault creates a PolicyCreateDefault with default headers values
func NewPolicyCreateDefault(code int) *PolicyCreateDefault {
	return &PolicyCreateDefault{
		_statusCode: code,
	}
}

/*
PolicyCreateDefault describes a response with status code -1, with default header values.

Error creating new policy
*/
type PolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this policy create default response has a 2xx status code
func (o *PolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy create default response has a 3xx status code
func (o *PolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy create default response has a 4xx status code
func (o *PolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy create default response has a 5xx status code
func (o *PolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy create default response a status code equal to that given
func (o *PolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the policy create default response
func (o *PolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *PolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /policies][%d] PolicyCreate default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyCreateDefault) String() string {
	return fmt.Sprintf("[POST /policies][%d] PolicyCreate default  %+v", o._statusCode, o.Payload)
}

func (o *PolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
