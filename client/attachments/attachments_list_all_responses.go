// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

// AttachmentsListAllReader is a Reader for the AttachmentsListAll structure.
type AttachmentsListAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachmentsListAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachmentsListAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttachmentsListAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachmentsListAllOK creates a AttachmentsListAllOK with default headers values
func NewAttachmentsListAllOK() *AttachmentsListAllOK {
	return &AttachmentsListAllOK{}
}

/*
AttachmentsListAllOK describes a response with status code 200, with default header values.

List of all attachments
*/
type AttachmentsListAllOK struct {
	Payload []*models.Attachment
}

// IsSuccess returns true when this attachments list all o k response has a 2xx status code
func (o *AttachmentsListAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attachments list all o k response has a 3xx status code
func (o *AttachmentsListAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attachments list all o k response has a 4xx status code
func (o *AttachmentsListAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attachments list all o k response has a 5xx status code
func (o *AttachmentsListAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attachments list all o k response a status code equal to that given
func (o *AttachmentsListAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attachments list all o k response
func (o *AttachmentsListAllOK) Code() int {
	return 200
}

func (o *AttachmentsListAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /attachments][%d] attachmentsListAllOK %s", 200, payload)
}

func (o *AttachmentsListAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /attachments][%d] attachmentsListAllOK %s", 200, payload)
}

func (o *AttachmentsListAllOK) GetPayload() []*models.Attachment {
	return o.Payload
}

func (o *AttachmentsListAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachmentsListAllDefault creates a AttachmentsListAllDefault with default headers values
func NewAttachmentsListAllDefault(code int) *AttachmentsListAllDefault {
	return &AttachmentsListAllDefault{
		_statusCode: code,
	}
}

/*
AttachmentsListAllDefault describes a response with status code -1, with default header values.

Error getting a list of all attachments
*/
type AttachmentsListAllDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this attachments list all default response has a 2xx status code
func (o *AttachmentsListAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attachments list all default response has a 3xx status code
func (o *AttachmentsListAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attachments list all default response has a 4xx status code
func (o *AttachmentsListAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attachments list all default response has a 5xx status code
func (o *AttachmentsListAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attachments list all default response a status code equal to that given
func (o *AttachmentsListAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attachments list all default response
func (o *AttachmentsListAllDefault) Code() int {
	return o._statusCode
}

func (o *AttachmentsListAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /attachments][%d] AttachmentsListAll default %s", o._statusCode, payload)
}

func (o *AttachmentsListAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /attachments][%d] AttachmentsListAll default %s", o._statusCode, payload)
}

func (o *AttachmentsListAllDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttachmentsListAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
