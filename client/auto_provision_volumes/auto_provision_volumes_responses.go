// Code generated by go-swagger; DO NOT EDIT.

package auto_provision_volumes

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

// AutoProvisionVolumesReader is a Reader for the AutoProvisionVolumes structure.
type AutoProvisionVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoProvisionVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoProvisionVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAutoProvisionVolumesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewAutoProvisionVolumesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAutoProvisionVolumesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoProvisionVolumesOK creates a AutoProvisionVolumesOK with default headers values
func NewAutoProvisionVolumesOK() *AutoProvisionVolumesOK {
	return &AutoProvisionVolumesOK{}
}

/*
AutoProvisionVolumesOK describes a response with status code 200, with default header values.

Auto provision volume has been created successfully
*/
type AutoProvisionVolumesOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this auto provision volumes o k response has a 2xx status code
func (o *AutoProvisionVolumesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto provision volumes o k response has a 3xx status code
func (o *AutoProvisionVolumesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto provision volumes o k response has a 4xx status code
func (o *AutoProvisionVolumesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto provision volumes o k response has a 5xx status code
func (o *AutoProvisionVolumesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto provision volumes o k response a status code equal to that given
func (o *AutoProvisionVolumesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto provision volumes o k response
func (o *AutoProvisionVolumesOK) Code() int {
	return 200
}

func (o *AutoProvisionVolumesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesOK %s", 200, payload)
}

func (o *AutoProvisionVolumesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesOK %s", 200, payload)
}

func (o *AutoProvisionVolumesOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *AutoProvisionVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoProvisionVolumesAccepted creates a AutoProvisionVolumesAccepted with default headers values
func NewAutoProvisionVolumesAccepted() *AutoProvisionVolumesAccepted {
	return &AutoProvisionVolumesAccepted{}
}

/*
AutoProvisionVolumesAccepted describes a response with status code 202, with default header values.

202 response
*/
type AutoProvisionVolumesAccepted struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto provision volumes accepted response has a 2xx status code
func (o *AutoProvisionVolumesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto provision volumes accepted response has a 3xx status code
func (o *AutoProvisionVolumesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto provision volumes accepted response has a 4xx status code
func (o *AutoProvisionVolumesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto provision volumes accepted response has a 5xx status code
func (o *AutoProvisionVolumesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this auto provision volumes accepted response a status code equal to that given
func (o *AutoProvisionVolumesAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the auto provision volumes accepted response
func (o *AutoProvisionVolumesAccepted) Code() int {
	return 202
}

func (o *AutoProvisionVolumesAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesAccepted %s", 202, payload)
}

func (o *AutoProvisionVolumesAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesAccepted %s", 202, payload)
}

func (o *AutoProvisionVolumesAccepted) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoProvisionVolumesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoProvisionVolumesConflict creates a AutoProvisionVolumesConflict with default headers values
func NewAutoProvisionVolumesConflict() *AutoProvisionVolumesConflict {
	return &AutoProvisionVolumesConflict{}
}

/*
AutoProvisionVolumesConflict describes a response with status code 409, with default header values.

409 response
*/
type AutoProvisionVolumesConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto provision volumes conflict response has a 2xx status code
func (o *AutoProvisionVolumesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auto provision volumes conflict response has a 3xx status code
func (o *AutoProvisionVolumesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto provision volumes conflict response has a 4xx status code
func (o *AutoProvisionVolumesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this auto provision volumes conflict response has a 5xx status code
func (o *AutoProvisionVolumesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this auto provision volumes conflict response a status code equal to that given
func (o *AutoProvisionVolumesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the auto provision volumes conflict response
func (o *AutoProvisionVolumesConflict) Code() int {
	return 409
}

func (o *AutoProvisionVolumesConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesConflict %s", 409, payload)
}

func (o *AutoProvisionVolumesConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesConflict %s", 409, payload)
}

func (o *AutoProvisionVolumesConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoProvisionVolumesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoProvisionVolumesDefault creates a AutoProvisionVolumesDefault with default headers values
func NewAutoProvisionVolumesDefault(code int) *AutoProvisionVolumesDefault {
	return &AutoProvisionVolumesDefault{
		_statusCode: code,
	}
}

/*
AutoProvisionVolumesDefault describes a response with status code -1, with default header values.

Error creating new volume
*/
type AutoProvisionVolumesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto provision volumes default response has a 2xx status code
func (o *AutoProvisionVolumesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto provision volumes default response has a 3xx status code
func (o *AutoProvisionVolumesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto provision volumes default response has a 4xx status code
func (o *AutoProvisionVolumesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto provision volumes default response has a 5xx status code
func (o *AutoProvisionVolumesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto provision volumes default response a status code equal to that given
func (o *AutoProvisionVolumesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auto provision volumes default response
func (o *AutoProvisionVolumesDefault) Code() int {
	return o._statusCode
}

func (o *AutoProvisionVolumesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] AutoProvisionVolumes default %s", o._statusCode, payload)
}

func (o *AutoProvisionVolumesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] AutoProvisionVolumes default %s", o._statusCode, payload)
}

func (o *AutoProvisionVolumesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoProvisionVolumesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
