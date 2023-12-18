// Code generated by go-swagger; DO NOT EDIT.

package auto_provision_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

func (o *AutoProvisionVolumesOK) Error() string {
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesOK  %+v", 200, o.Payload)
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

func (o *AutoProvisionVolumesAccepted) Error() string {
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesAccepted  %+v", 202, o.Payload)
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

func (o *AutoProvisionVolumesConflict) Error() string {
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] autoProvisionVolumesConflict  %+v", 409, o.Payload)
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

// Code gets the status code for the auto provision volumes default response
func (o *AutoProvisionVolumesDefault) Code() int {
	return o._statusCode
}

func (o *AutoProvisionVolumesDefault) Error() string {
	return fmt.Sprintf("[POST /autoprovisionvolumes][%d] AutoProvisionVolumes default  %+v", o._statusCode, o.Payload)
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
