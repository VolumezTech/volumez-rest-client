// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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
		return nil, runtime.NewAPIError("[POST /signup] UserCreate", response, response.Code())
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

// IsSuccess returns true when this user create o k response has a 2xx status code
func (o *UserCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user create o k response has a 3xx status code
func (o *UserCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user create o k response has a 4xx status code
func (o *UserCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user create o k response has a 5xx status code
func (o *UserCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user create o k response a status code equal to that given
func (o *UserCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user create o k response
func (o *UserCreateOK) Code() int {
	return 200
}

func (o *UserCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /signup][%d] userCreateOK %s", 200, payload)
}

func (o *UserCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /signup][%d] userCreateOK %s", 200, payload)
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

// IsSuccess returns true when this user create internal server error response has a 2xx status code
func (o *UserCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user create internal server error response has a 3xx status code
func (o *UserCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user create internal server error response has a 4xx status code
func (o *UserCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user create internal server error response has a 5xx status code
func (o *UserCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user create internal server error response a status code equal to that given
func (o *UserCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user create internal server error response
func (o *UserCreateInternalServerError) Code() int {
	return 500
}

func (o *UserCreateInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /signup][%d] userCreateInternalServerError %s", 500, payload)
}

func (o *UserCreateInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /signup][%d] userCreateInternalServerError %s", 500, payload)
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
