// Code generated by go-swagger; DO NOT EDIT.

package capacity_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// CapacityGroupGetReader is a Reader for the CapacityGroupGet structure.
type CapacityGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CapacityGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCapacityGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCapacityGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCapacityGroupGetOK creates a CapacityGroupGetOK with default headers values
func NewCapacityGroupGetOK() *CapacityGroupGetOK {
	return &CapacityGroupGetOK{}
}

/*
	CapacityGroupGetOK describes a response with status code 200, with default header values.

capacity groups
*/
type CapacityGroupGetOK struct {
	Payload []*models.CapacityGroup
}

func (o *CapacityGroupGetOK) Error() string {
	return fmt.Sprintf("[GET /capacitygroups][%d] capacityGroupGetOK  %+v", 200, o.Payload)
}
func (o *CapacityGroupGetOK) GetPayload() []*models.CapacityGroup {
	return o.Payload
}

func (o *CapacityGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCapacityGroupGetDefault creates a CapacityGroupGetDefault with default headers values
func NewCapacityGroupGetDefault(code int) *CapacityGroupGetDefault {
	return &CapacityGroupGetDefault{
		_statusCode: code,
	}
}

/*
	CapacityGroupGetDefault describes a response with status code -1, with default header values.

Error getting capacity groups
*/
type CapacityGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the capacity group get default response
func (o *CapacityGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *CapacityGroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /capacitygroups][%d] capacityGroupGet default  %+v", o._statusCode, o.Payload)
}
func (o *CapacityGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CapacityGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
