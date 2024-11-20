// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// MediaUnassignReader is a Reader for the MediaUnassign structure.
type MediaUnassignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaUnassignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaUnassignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaUnassignDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaUnassignOK creates a MediaUnassignOK with default headers values
func NewMediaUnassignOK() *MediaUnassignOK {
	return &MediaUnassignOK{}
}

/*
	MediaUnassignOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type MediaUnassignOK struct {
	Payload *models.RegularResponse
}

func (o *MediaUnassignOK) Error() string {
	return fmt.Sprintf("[PATCH /media/{media}/unassign][%d] mediaUnassignOK  %+v", 200, o.Payload)
}
func (o *MediaUnassignOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaUnassignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaUnassignDefault creates a MediaUnassignDefault with default headers values
func NewMediaUnassignDefault(code int) *MediaUnassignDefault {
	return &MediaUnassignDefault{
		_statusCode: code,
	}
}

/*
	MediaUnassignDefault describes a response with status code -1, with default header values.

Error completing the operation
*/
type MediaUnassignDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the media unassign default response
func (o *MediaUnassignDefault) Code() int {
	return o._statusCode
}

func (o *MediaUnassignDefault) Error() string {
	return fmt.Sprintf("[PATCH /media/{media}/unassign][%d] MediaUnassign default  %+v", o._statusCode, o.Payload)
}
func (o *MediaUnassignDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaUnassignDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
