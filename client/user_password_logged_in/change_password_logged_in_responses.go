// Code generated by go-swagger; DO NOT EDIT.

package user_password_logged_in

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// ChangePasswordLoggedInReader is a Reader for the ChangePasswordLoggedIn structure.
type ChangePasswordLoggedInReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePasswordLoggedInReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePasswordLoggedInOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangePasswordLoggedInDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangePasswordLoggedInOK creates a ChangePasswordLoggedInOK with default headers values
func NewChangePasswordLoggedInOK() *ChangePasswordLoggedInOK {
	return &ChangePasswordLoggedInOK{}
}

/* ChangePasswordLoggedInOK describes a response with status code 200, with default header values.

New password changed successfully
*/
type ChangePasswordLoggedInOK struct {
	Payload *models.RegularResponse
}

func (o *ChangePasswordLoggedInOK) Error() string {
	return fmt.Sprintf("[POST /tenant/user/changepasswordloggedin][%d] changePasswordLoggedInOK  %+v", 200, o.Payload)
}
func (o *ChangePasswordLoggedInOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *ChangePasswordLoggedInOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePasswordLoggedInDefault creates a ChangePasswordLoggedInDefault with default headers values
func NewChangePasswordLoggedInDefault(code int) *ChangePasswordLoggedInDefault {
	return &ChangePasswordLoggedInDefault{
		_statusCode: code,
	}
}

/* ChangePasswordLoggedInDefault describes a response with status code -1, with default header values.

Error changing password
*/
type ChangePasswordLoggedInDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the change password logged in default response
func (o *ChangePasswordLoggedInDefault) Code() int {
	return o._statusCode
}

func (o *ChangePasswordLoggedInDefault) Error() string {
	return fmt.Sprintf("[POST /tenant/user/changepasswordloggedin][%d] ChangePasswordLoggedIn default  %+v", o._statusCode, o.Payload)
}
func (o *ChangePasswordLoggedInDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ChangePasswordLoggedInDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
