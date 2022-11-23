// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// VolumeCreateReader is a Reader for the VolumeCreate structure.
type VolumeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewVolumeCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewVolumeCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeCreateOK creates a VolumeCreateOK with default headers values
func NewVolumeCreateOK() *VolumeCreateOK {
	return &VolumeCreateOK{}
}

/* VolumeCreateOK describes a response with status code 200, with default header values.

Volume has been created successfully
*/
type VolumeCreateOK struct {
	Payload *models.RegularResponse
}

func (o *VolumeCreateOK) Error() string {
	return fmt.Sprintf("[POST /volumes][%d] volumeCreateOK  %+v", 200, o.Payload)
}
func (o *VolumeCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *VolumeCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateAccepted creates a VolumeCreateAccepted with default headers values
func NewVolumeCreateAccepted() *VolumeCreateAccepted {
	return &VolumeCreateAccepted{}
}

/* VolumeCreateAccepted describes a response with status code 202, with default header values.

202 response
*/
type VolumeCreateAccepted struct {
	Payload *models.ErrorResponse
}

func (o *VolumeCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /volumes][%d] volumeCreateAccepted  %+v", 202, o.Payload)
}
func (o *VolumeCreateAccepted) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateConflict creates a VolumeCreateConflict with default headers values
func NewVolumeCreateConflict() *VolumeCreateConflict {
	return &VolumeCreateConflict{}
}

/* VolumeCreateConflict describes a response with status code 409, with default header values.

409 response
*/
type VolumeCreateConflict struct {
	Payload *models.ErrorResponse
}

func (o *VolumeCreateConflict) Error() string {
	return fmt.Sprintf("[POST /volumes][%d] volumeCreateConflict  %+v", 409, o.Payload)
}
func (o *VolumeCreateConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateDefault creates a VolumeCreateDefault with default headers values
func NewVolumeCreateDefault(code int) *VolumeCreateDefault {
	return &VolumeCreateDefault{
		_statusCode: code,
	}
}

/* VolumeCreateDefault describes a response with status code -1, with default header values.

Error creating new volume
*/
type VolumeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume create default response
func (o *VolumeCreateDefault) Code() int {
	return o._statusCode
}

func (o *VolumeCreateDefault) Error() string {
	return fmt.Sprintf("[POST /volumes][%d] VolumeCreate default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
