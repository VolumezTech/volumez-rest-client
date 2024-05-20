// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// PolicyPlanReader is a Reader for the PolicyPlan structure.
type PolicyPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PolicyPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPolicyPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPolicyPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPolicyPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPolicyPlanOK creates a PolicyPlanOK with default headers values
func NewPolicyPlanOK() *PolicyPlanOK {
	return &PolicyPlanOK{}
}

/*
PolicyPlanOK describes a response with status code 200, with default header values.

Policy volume group plan
*/
type PolicyPlanOK struct {
	Payload *models.Plan
}

// IsSuccess returns true when this policy plan o k response has a 2xx status code
func (o *PolicyPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this policy plan o k response has a 3xx status code
func (o *PolicyPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy plan o k response has a 4xx status code
func (o *PolicyPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this policy plan o k response has a 5xx status code
func (o *PolicyPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this policy plan o k response a status code equal to that given
func (o *PolicyPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the policy plan o k response
func (o *PolicyPlanOK) Code() int {
	return 200
}

func (o *PolicyPlanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanOK %s", 200, payload)
}

func (o *PolicyPlanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanOK %s", 200, payload)
}

func (o *PolicyPlanOK) GetPayload() *models.Plan {
	return o.Payload
}

func (o *PolicyPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyPlanNotFound creates a PolicyPlanNotFound with default headers values
func NewPolicyPlanNotFound() *PolicyPlanNotFound {
	return &PolicyPlanNotFound{}
}

/*
PolicyPlanNotFound describes a response with status code 404, with default header values.

Plan failed
*/
type PolicyPlanNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this policy plan not found response has a 2xx status code
func (o *PolicyPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this policy plan not found response has a 3xx status code
func (o *PolicyPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this policy plan not found response has a 4xx status code
func (o *PolicyPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this policy plan not found response has a 5xx status code
func (o *PolicyPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this policy plan not found response a status code equal to that given
func (o *PolicyPlanNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the policy plan not found response
func (o *PolicyPlanNotFound) Code() int {
	return 404
}

func (o *PolicyPlanNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanNotFound %s", 404, payload)
}

func (o *PolicyPlanNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanNotFound %s", 404, payload)
}

func (o *PolicyPlanNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PolicyPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPolicyPlanDefault creates a PolicyPlanDefault with default headers values
func NewPolicyPlanDefault(code int) *PolicyPlanDefault {
	return &PolicyPlanDefault{
		_statusCode: code,
	}
}

/*
PolicyPlanDefault describes a response with status code -1, with default header values.

Error getting policy volume group plan
*/
type PolicyPlanDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this policy plan default response has a 2xx status code
func (o *PolicyPlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this policy plan default response has a 3xx status code
func (o *PolicyPlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this policy plan default response has a 4xx status code
func (o *PolicyPlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this policy plan default response has a 5xx status code
func (o *PolicyPlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this policy plan default response a status code equal to that given
func (o *PolicyPlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the policy plan default response
func (o *PolicyPlanDefault) Code() int {
	return o._statusCode
}

func (o *PolicyPlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] PolicyPlan default %s", o._statusCode, payload)
}

func (o *PolicyPlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] PolicyPlan default %s", o._statusCode, payload)
}

func (o *PolicyPlanDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PolicyPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
