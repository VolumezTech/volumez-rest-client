// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// UserCreateReader is a Reader for the UserCreate structure.
type UserCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUserCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserCreateOK creates a UserCreateOK with default headers values
func NewUserCreateOK() *UserCreateOK {
	return &UserCreateOK{}
}

/*
	UserCreateOK describes a response with status code 200, with default header values.

Successfully Signed up
*/
type UserCreateOK struct {
	Payload *models.SignUpResponse
}

func (o *UserCreateOK) Error() string {
	return fmt.Sprintf("[POST /signup][%d] userCreateOK  %+v", 200, o.Payload)
}
func (o *UserCreateOK) GetPayload() *models.SignUpResponse {
	return o.Payload
}

func (o *UserCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignUpResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCreateInternalServerError creates a UserCreateInternalServerError with default headers values
func NewUserCreateInternalServerError() *UserCreateInternalServerError {
	return &UserCreateInternalServerError{}
}

/*
	UserCreateInternalServerError describes a response with status code 500, with default header values.

Error signing up
*/
type UserCreateInternalServerError struct {
	Payload *models.SignUpResponse
}

func (o *UserCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /signup][%d] userCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *UserCreateInternalServerError) GetPayload() *models.SignUpResponse {
	return o.Payload
}

func (o *UserCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignUpResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
