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

// VolumeRecoverInitiateReader is a Reader for the VolumeRecoverInitiate structure.
type VolumeRecoverInitiateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeRecoverInitiateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeRecoverInitiateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeRecoverInitiateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeRecoverInitiateOK creates a VolumeRecoverInitiateOK with default headers values
func NewVolumeRecoverInitiateOK() *VolumeRecoverInitiateOK {
	return &VolumeRecoverInitiateOK{}
}

/*
VolumeRecoverInitiateOK describes a response with status code 200, with default header values.

Volume recovery was initiated successfully
*/
type VolumeRecoverInitiateOK struct {
	Payload *models.RegularResponse
}

func (o *VolumeRecoverInitiateOK) Error() string {
	return fmt.Sprintf("[POST /volumes/{volume}/recover][%d] volumeRecoverInitiateOK  %+v", 200, o.Payload)
}
func (o *VolumeRecoverInitiateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *VolumeRecoverInitiateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeRecoverInitiateDefault creates a VolumeRecoverInitiateDefault with default headers values
func NewVolumeRecoverInitiateDefault(code int) *VolumeRecoverInitiateDefault {
	return &VolumeRecoverInitiateDefault{
		_statusCode: code,
	}
}

/*
VolumeRecoverInitiateDefault describes a response with status code -1, with default header values.

Error initiating recovery on volume
*/
type VolumeRecoverInitiateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume recover initiate default response
func (o *VolumeRecoverInitiateDefault) Code() int {
	return o._statusCode
}

func (o *VolumeRecoverInitiateDefault) Error() string {
	return fmt.Sprintf("[POST /volumes/{volume}/recover][%d] VolumeRecoverInitiate default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeRecoverInitiateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeRecoverInitiateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
