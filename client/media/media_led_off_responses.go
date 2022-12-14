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

// MediaLedOffReader is a Reader for the MediaLedOff structure.
type MediaLedOffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaLedOffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaLedOffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaLedOffDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaLedOffOK creates a MediaLedOffOK with default headers values
func NewMediaLedOffOK() *MediaLedOffOK {
	return &MediaLedOffOK{}
}

/* MediaLedOffOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type MediaLedOffOK struct {
	Payload *models.RegularResponse
}

func (o *MediaLedOffOK) Error() string {
	return fmt.Sprintf("[GET /media/{media}/ledoff][%d] mediaLedOffOK  %+v", 200, o.Payload)
}
func (o *MediaLedOffOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaLedOffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaLedOffDefault creates a MediaLedOffDefault with default headers values
func NewMediaLedOffDefault(code int) *MediaLedOffDefault {
	return &MediaLedOffDefault{
		_statusCode: code,
	}
}

/* MediaLedOffDefault describes a response with status code -1, with default header values.

Error completing the operation
*/
type MediaLedOffDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the media led off default response
func (o *MediaLedOffDefault) Code() int {
	return o._statusCode
}

func (o *MediaLedOffDefault) Error() string {
	return fmt.Sprintf("[GET /media/{media}/ledoff][%d] MediaLedOff default  %+v", o._statusCode, o.Payload)
}
func (o *MediaLedOffDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaLedOffDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
