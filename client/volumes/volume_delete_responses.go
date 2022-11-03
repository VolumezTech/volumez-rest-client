// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"bitbucket.org/volumez/volumez/test/restapi/models"
)

// VolumeDeleteReader is a Reader for the VolumeDelete structure.
type VolumeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeDeleteOK creates a VolumeDeleteOK with default headers values
func NewVolumeDeleteOK() *VolumeDeleteOK {
	return &VolumeDeleteOK{}
}

/* VolumeDeleteOK describes a response with status code 200, with default header values.

A volume was deleted successfully
*/
type VolumeDeleteOK struct {
	Payload *models.RegularResponse
}

func (o *VolumeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volume}][%d] volumeDeleteOK  %+v", 200, o.Payload)
}
func (o *VolumeDeleteOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *VolumeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeDeleteDefault creates a VolumeDeleteDefault with default headers values
func NewVolumeDeleteDefault(code int) *VolumeDeleteDefault {
	return &VolumeDeleteDefault{
		_statusCode: code,
	}
}

/* VolumeDeleteDefault describes a response with status code -1, with default header values.

Error deleting a volume
*/
type VolumeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume delete default response
func (o *VolumeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VolumeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volume}][%d] VolumeDelete default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
