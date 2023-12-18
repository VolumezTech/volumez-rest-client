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

// AttachmentGetReader is a Reader for the AttachmentGet structure.
type AttachmentGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachmentGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachmentGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttachmentGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachmentGetOK creates a AttachmentGetOK with default headers values
func NewAttachmentGetOK() *AttachmentGetOK {
	return &AttachmentGetOK{}
}

/*
AttachmentGetOK describes a response with status code 200, with default header values.

Properties of an attachment
*/
type AttachmentGetOK struct {
	Payload *models.Attachment
}

func (o *AttachmentGetOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments/{node}][%d] attachmentGetOK  %+v", 200, o.Payload)
}
func (o *AttachmentGetOK) GetPayload() *models.Attachment {
	return o.Payload
}

func (o *AttachmentGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachmentGetDefault creates a AttachmentGetDefault with default headers values
func NewAttachmentGetDefault(code int) *AttachmentGetDefault {
	return &AttachmentGetDefault{
		_statusCode: code,
	}
}

/*
AttachmentGetDefault describes a response with status code -1, with default header values.

Error getting properties of an attachment
*/
type AttachmentGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the attachment get default response
func (o *AttachmentGetDefault) Code() int {
	return o._statusCode
}

func (o *AttachmentGetDefault) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots/{snapshot}/attachments/{node}][%d] AttachmentGet default  %+v", o._statusCode, o.Payload)
}
func (o *AttachmentGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttachmentGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
