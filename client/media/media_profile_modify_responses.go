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

// MediaProfileModifyReader is a Reader for the MediaProfileModify structure.
type MediaProfileModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaProfileModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaProfileModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaProfileModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaProfileModifyOK creates a MediaProfileModifyOK with default headers values
func NewMediaProfileModifyOK() *MediaProfileModifyOK {
	return &MediaProfileModifyOK{}
}

/*
MediaProfileModifyOK describes a response with status code 200, with default header values.

A media profile was updated successfully
*/
type MediaProfileModifyOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this media profile modify o k response has a 2xx status code
func (o *MediaProfileModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this media profile modify o k response has a 3xx status code
func (o *MediaProfileModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media profile modify o k response has a 4xx status code
func (o *MediaProfileModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this media profile modify o k response has a 5xx status code
func (o *MediaProfileModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this media profile modify o k response a status code equal to that given
func (o *MediaProfileModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the media profile modify o k response
func (o *MediaProfileModifyOK) Code() int {
	return 200
}

func (o *MediaProfileModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /media/{media}/profile][%d] mediaProfileModifyOK  %+v", 200, o.Payload)
}

func (o *MediaProfileModifyOK) String() string {
	return fmt.Sprintf("[PATCH /media/{media}/profile][%d] mediaProfileModifyOK  %+v", 200, o.Payload)
}

func (o *MediaProfileModifyOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaProfileModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaProfileModifyDefault creates a MediaProfileModifyDefault with default headers values
func NewMediaProfileModifyDefault(code int) *MediaProfileModifyDefault {
	return &MediaProfileModifyDefault{
		_statusCode: code,
	}
}

/*
MediaProfileModifyDefault describes a response with status code -1, with default header values.

Error updating a media profile
*/
type MediaProfileModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this media profile modify default response has a 2xx status code
func (o *MediaProfileModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this media profile modify default response has a 3xx status code
func (o *MediaProfileModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this media profile modify default response has a 4xx status code
func (o *MediaProfileModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this media profile modify default response has a 5xx status code
func (o *MediaProfileModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this media profile modify default response a status code equal to that given
func (o *MediaProfileModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the media profile modify default response
func (o *MediaProfileModifyDefault) Code() int {
	return o._statusCode
}

func (o *MediaProfileModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /media/{media}/profile][%d] MediaProfileModify default  %+v", o._statusCode, o.Payload)
}

func (o *MediaProfileModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /media/{media}/profile][%d] MediaProfileModify default  %+v", o._statusCode, o.Payload)
}

func (o *MediaProfileModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaProfileModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
