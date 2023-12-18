// Code generated by go-swagger; DO NOT EDIT.

package signin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// SignInReader is a Reader for the SignIn structure.
type SignInReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignInReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignInOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSignInDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSignInOK creates a SignInOK with default headers values
func NewSignInOK() *SignInOK {
	return &SignInOK{}
}

/*
SignInOK describes a response with status code 200, with default header values.

Sign in was successful
*/
type SignInOK struct {
	Payload *models.SignInResponse
}

// IsSuccess returns true when this sign in o k response has a 2xx status code
func (o *SignInOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sign in o k response has a 3xx status code
func (o *SignInOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sign in o k response has a 4xx status code
func (o *SignInOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sign in o k response has a 5xx status code
func (o *SignInOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sign in o k response a status code equal to that given
func (o *SignInOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sign in o k response
func (o *SignInOK) Code() int {
	return 200
}

func (o *SignInOK) Error() string {
	return fmt.Sprintf("[POST /signin][%d] signInOK  %+v", 200, o.Payload)
}

func (o *SignInOK) String() string {
	return fmt.Sprintf("[POST /signin][%d] signInOK  %+v", 200, o.Payload)
}

func (o *SignInOK) GetPayload() *models.SignInResponse {
	return o.Payload
}

func (o *SignInOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignInResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignInDefault creates a SignInDefault with default headers values
func NewSignInDefault(code int) *SignInDefault {
	return &SignInDefault{
		_statusCode: code,
	}
}

/*
SignInDefault describes a response with status code -1, with default header values.

Error signing in
*/
type SignInDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this sign in default response has a 2xx status code
func (o *SignInDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sign in default response has a 3xx status code
func (o *SignInDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sign in default response has a 4xx status code
func (o *SignInDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sign in default response has a 5xx status code
func (o *SignInDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sign in default response a status code equal to that given
func (o *SignInDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sign in default response
func (o *SignInDefault) Code() int {
	return o._statusCode
}

func (o *SignInDefault) Error() string {
	return fmt.Sprintf("[POST /signin][%d] SignIn default  %+v", o._statusCode, o.Payload)
}

func (o *SignInDefault) String() string {
	return fmt.Sprintf("[POST /signin][%d] SignIn default  %+v", o._statusCode, o.Payload)
}

func (o *SignInDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SignInDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
