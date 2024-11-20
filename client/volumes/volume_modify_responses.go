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

// VolumeModifyReader is a Reader for the VolumeModify structure.
type VolumeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeModifyOK creates a VolumeModifyOK with default headers values
func NewVolumeModifyOK() *VolumeModifyOK {
	return &VolumeModifyOK{}
}

/*
	VolumeModifyOK describes a response with status code 200, with default header values.

A volume was updated successfully
*/
type VolumeModifyOK struct {
	Payload *models.RegularResponse
}

func (o *VolumeModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}][%d] volumeModifyOK  %+v", 200, o.Payload)
}
func (o *VolumeModifyOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *VolumeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeModifyDefault creates a VolumeModifyDefault with default headers values
func NewVolumeModifyDefault(code int) *VolumeModifyDefault {
	return &VolumeModifyDefault{
		_statusCode: code,
	}
}

/*
	VolumeModifyDefault describes a response with status code -1, with default header values.

Error updating a volume
*/
type VolumeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume modify default response
func (o *VolumeModifyDefault) Code() int {
	return o._statusCode
}

func (o *VolumeModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}][%d] VolumeModify default  %+v", o._statusCode, o.Payload)
}
func (o *VolumeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
