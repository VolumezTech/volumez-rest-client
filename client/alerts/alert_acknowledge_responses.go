// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// AlertAcknowledgeReader is a Reader for the AlertAcknowledge structure.
type AlertAcknowledgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertAcknowledgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertAcknowledgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAlertAcknowledgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAlertAcknowledgeOK creates a AlertAcknowledgeOK with default headers values
func NewAlertAcknowledgeOK() *AlertAcknowledgeOK {
	return &AlertAcknowledgeOK{}
}

/*
AlertAcknowledgeOK describes a response with status code 200, with default header values.

Alert was successfully acknowledged
*/
type AlertAcknowledgeOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this alert acknowledge o k response has a 2xx status code
func (o *AlertAcknowledgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alert acknowledge o k response has a 3xx status code
func (o *AlertAcknowledgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alert acknowledge o k response has a 4xx status code
func (o *AlertAcknowledgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alert acknowledge o k response has a 5xx status code
func (o *AlertAcknowledgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alert acknowledge o k response a status code equal to that given
func (o *AlertAcknowledgeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the alert acknowledge o k response
func (o *AlertAcknowledgeOK) Code() int {
	return 200
}

func (o *AlertAcknowledgeOK) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert}/acknowledge][%d] alertAcknowledgeOK  %+v", 200, o.Payload)
}

func (o *AlertAcknowledgeOK) String() string {
	return fmt.Sprintf("[POST /alerts/{alert}/acknowledge][%d] alertAcknowledgeOK  %+v", 200, o.Payload)
}

func (o *AlertAcknowledgeOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *AlertAcknowledgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertAcknowledgeDefault creates a AlertAcknowledgeDefault with default headers values
func NewAlertAcknowledgeDefault(code int) *AlertAcknowledgeDefault {
	return &AlertAcknowledgeDefault{
		_statusCode: code,
	}
}

/*
AlertAcknowledgeDefault describes a response with status code -1, with default header values.

Error acknowledging alert
*/
type AlertAcknowledgeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this alert acknowledge default response has a 2xx status code
func (o *AlertAcknowledgeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this alert acknowledge default response has a 3xx status code
func (o *AlertAcknowledgeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this alert acknowledge default response has a 4xx status code
func (o *AlertAcknowledgeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this alert acknowledge default response has a 5xx status code
func (o *AlertAcknowledgeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this alert acknowledge default response a status code equal to that given
func (o *AlertAcknowledgeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the alert acknowledge default response
func (o *AlertAcknowledgeDefault) Code() int {
	return o._statusCode
}

func (o *AlertAcknowledgeDefault) Error() string {
	return fmt.Sprintf("[POST /alerts/{alert}/acknowledge][%d] AlertAcknowledge default  %+v", o._statusCode, o.Payload)
}

func (o *AlertAcknowledgeDefault) String() string {
	return fmt.Sprintf("[POST /alerts/{alert}/acknowledge][%d] AlertAcknowledge default  %+v", o._statusCode, o.Payload)
}

func (o *AlertAcknowledgeDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AlertAcknowledgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}