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

// MediaDrainReader is a Reader for the MediaDrain structure.
type MediaDrainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaDrainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaDrainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaDrainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaDrainOK creates a MediaDrainOK with default headers values
func NewMediaDrainOK() *MediaDrainOK {
	return &MediaDrainOK{}
}

/*
	MediaDrainOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type MediaDrainOK struct {
	Payload *models.RegularResponse
}

func (o *MediaDrainOK) Error() string {
	return fmt.Sprintf("[POST /media/{media}/drain][%d] mediaDrainOK  %+v", 200, o.Payload)
}
func (o *MediaDrainOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaDrainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaDrainDefault creates a MediaDrainDefault with default headers values
func NewMediaDrainDefault(code int) *MediaDrainDefault {
	return &MediaDrainDefault{
		_statusCode: code,
	}
}

/*
	MediaDrainDefault describes a response with status code -1, with default header values.

Error draining media
*/
type MediaDrainDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the media drain default response
func (o *MediaDrainDefault) Code() int {
	return o._statusCode
}

func (o *MediaDrainDefault) Error() string {
	return fmt.Sprintf("[POST /media/{media}/drain][%d] MediaDrain default  %+v", o._statusCode, o.Payload)
}
func (o *MediaDrainDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaDrainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
