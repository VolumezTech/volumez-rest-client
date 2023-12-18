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

func (o *PolicyPlanOK) Error() string {
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanOK  %+v", 200, o.Payload)
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

func (o *PolicyPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] policyPlanNotFound  %+v", 404, o.Payload)
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

// Code gets the status code for the policy plan default response
func (o *PolicyPlanDefault) Code() int {
	return o._statusCode
}

func (o *PolicyPlanDefault) Error() string {
	return fmt.Sprintf("[GET /policies/{policy}/size/{size}/zone/{zone}][%d] PolicyPlan default  %+v", o._statusCode, o.Payload)
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
