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

// MediaProfileReader is a Reader for the MediaProfile structure.
type MediaProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaProfileOK creates a MediaProfileOK with default headers values
func NewMediaProfileOK() *MediaProfileOK {
	return &MediaProfileOK{}
}

/*
MediaProfileOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type MediaProfileOK struct {
	Payload *models.RegularResponse
}

func (o *MediaProfileOK) Error() string {
	return fmt.Sprintf("[GET /media/{media}/profile][%d] mediaProfileOK  %+v", 200, o.Payload)
}
func (o *MediaProfileOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaProfileDefault creates a MediaProfileDefault with default headers values
func NewMediaProfileDefault(code int) *MediaProfileDefault {
	return &MediaProfileDefault{
		_statusCode: code,
	}
}

/*
MediaProfileDefault describes a response with status code -1, with default header values.

Error completing the operation
*/
type MediaProfileDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the media profile default response
func (o *MediaProfileDefault) Code() int {
	return o._statusCode
}

func (o *MediaProfileDefault) Error() string {
	return fmt.Sprintf("[GET /media/{media}/profile][%d] MediaProfile default  %+v", o._statusCode, o.Payload)
}
func (o *MediaProfileDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
