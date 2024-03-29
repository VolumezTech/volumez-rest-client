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

// MediaDeleteReader is a Reader for the MediaDelete structure.
type MediaDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaDeleteOK creates a MediaDeleteOK with default headers values
func NewMediaDeleteOK() *MediaDeleteOK {
	return &MediaDeleteOK{}
}

/*
MediaDeleteOK describes a response with status code 200, with default header values.

A node delete job was created successfully
*/
type MediaDeleteOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this media delete o k response has a 2xx status code
func (o *MediaDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this media delete o k response has a 3xx status code
func (o *MediaDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media delete o k response has a 4xx status code
func (o *MediaDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this media delete o k response has a 5xx status code
func (o *MediaDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this media delete o k response a status code equal to that given
func (o *MediaDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the media delete o k response
func (o *MediaDeleteOK) Code() int {
	return 200
}

func (o *MediaDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /media/{media}][%d] mediaDeleteOK  %+v", 200, o.Payload)
}

func (o *MediaDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /media/{media}][%d] mediaDeleteOK  %+v", 200, o.Payload)
}

func (o *MediaDeleteOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *MediaDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaDeleteDefault creates a MediaDeleteDefault with default headers values
func NewMediaDeleteDefault(code int) *MediaDeleteDefault {
	return &MediaDeleteDefault{
		_statusCode: code,
	}
}

/*
MediaDeleteDefault describes a response with status code -1, with default header values.

Error deleting a node
*/
type MediaDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this media delete default response has a 2xx status code
func (o *MediaDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this media delete default response has a 3xx status code
func (o *MediaDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this media delete default response has a 4xx status code
func (o *MediaDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this media delete default response has a 5xx status code
func (o *MediaDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this media delete default response a status code equal to that given
func (o *MediaDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the media delete default response
func (o *MediaDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MediaDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /media/{media}][%d] MediaDelete default  %+v", o._statusCode, o.Payload)
}

func (o *MediaDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /media/{media}][%d] MediaDelete default  %+v", o._statusCode, o.Payload)
}

func (o *MediaDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediaDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
