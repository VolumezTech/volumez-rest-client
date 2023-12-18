// Code generated by go-swagger; DO NOT EDIT.

package amp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*
GetSubscriptionOK describes a response with status code 200, with default header values.

confirmation of subscription registered
*/
type GetSubscriptionOK struct {
	Payload *models.RegularResponse
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /azuremarketplace/subscription][%d] getSubscriptionOK  %+v", 200, o.Payload)
}
func (o *GetSubscriptionOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionNotFound creates a GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

/*
GetSubscriptionNotFound describes a response with status code 404, with default header values.

Subscrition not available
*/
type GetSubscriptionNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /azuremarketplace/subscription][%d] getSubscriptionNotFound  %+v", 404, o.Payload)
}
func (o *GetSubscriptionNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionDefault creates a GetSubscriptionDefault with default headers values
func NewGetSubscriptionDefault(code int) *GetSubscriptionDefault {
	return &GetSubscriptionDefault{
		_statusCode: code,
	}
}

/*
GetSubscriptionDefault describes a response with status code -1, with default header values.

Error Getting volume plan
*/
type GetSubscriptionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get subscription default response
func (o *GetSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /azuremarketplace/subscription][%d] GetSubscription default  %+v", o._statusCode, o.Payload)
}
func (o *GetSubscriptionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
