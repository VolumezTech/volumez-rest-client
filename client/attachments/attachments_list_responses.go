// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// AttachmentsListReader is a Reader for the AttachmentsList structure.
type AttachmentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachmentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachmentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttachmentsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachmentsListOK creates a AttachmentsListOK with default headers values
func NewAttachmentsListOK() *AttachmentsListOK {
	return &AttachmentsListOK{}
}

/*
AttachmentsListOK describes a response with status code 200, with default header values.

List of attachments
*/
type AttachmentsListOK struct {
	Payload []*models.Attachment
}

// IsSuccess returns true when this attachments list o k response has a 2xx status code
func (o *AttachmentsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attachments list o k response has a 3xx status code
func (o *AttachmentsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attachments list o k response has a 4xx status code
func (o *AttachmentsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attachments list o k response has a 5xx status code
func (o *AttachmentsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attachments list o k response a status code equal to that given
func (o *AttachmentsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attachments list o k response
func (o *AttachmentsListOK) Code() int {
	return 200
}

func (o *AttachmentsListOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments][%d] attachmentsListOK  %+v", 200, o.Payload)
}

func (o *AttachmentsListOK) String() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments][%d] attachmentsListOK  %+v", 200, o.Payload)
}

func (o *AttachmentsListOK) GetPayload() []*models.Attachment {
	return o.Payload
}

func (o *AttachmentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachmentsListDefault creates a AttachmentsListDefault with default headers values
func NewAttachmentsListDefault(code int) *AttachmentsListDefault {
	return &AttachmentsListDefault{
		_statusCode: code,
	}
}

/*
AttachmentsListDefault describes a response with status code -1, with default header values.

Error getting a list of attachments
*/
type AttachmentsListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this attachments list default response has a 2xx status code
func (o *AttachmentsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attachments list default response has a 3xx status code
func (o *AttachmentsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attachments list default response has a 4xx status code
func (o *AttachmentsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attachments list default response has a 5xx status code
func (o *AttachmentsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attachments list default response a status code equal to that given
func (o *AttachmentsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attachments list default response
func (o *AttachmentsListDefault) Code() int {
	return o._statusCode
}

func (o *AttachmentsListDefault) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments][%d] AttachmentsList default  %+v", o._statusCode, o.Payload)
}

func (o *AttachmentsListDefault) String() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments][%d] AttachmentsList default  %+v", o._statusCode, o.Payload)
}

func (o *AttachmentsListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttachmentsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
