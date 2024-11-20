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

// AttachmentModifyReader is a Reader for the AttachmentModify structure.
type AttachmentModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachmentModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachmentModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttachmentModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachmentModifyOK creates a AttachmentModifyOK with default headers values
func NewAttachmentModifyOK() *AttachmentModifyOK {
	return &AttachmentModifyOK{}
}

/*
	AttachmentModifyOK describes a response with status code 200, with default header values.

An attachment was updated successfully
*/
type AttachmentModifyOK struct {
	Payload *models.RegularResponse
}

func (o *AttachmentModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}/snapshots/{snapshot}/attachments/{node}][%d] attachmentModifyOK  %+v", 200, o.Payload)
}
func (o *AttachmentModifyOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *AttachmentModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachmentModifyDefault creates a AttachmentModifyDefault with default headers values
func NewAttachmentModifyDefault(code int) *AttachmentModifyDefault {
	return &AttachmentModifyDefault{
		_statusCode: code,
	}
}

/*
	AttachmentModifyDefault describes a response with status code -1, with default header values.

Error updating an attachment
*/
type AttachmentModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the attachment modify default response
func (o *AttachmentModifyDefault) Code() int {
	return o._statusCode
}

func (o *AttachmentModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}/snapshots/{snapshot}/attachments/{node}][%d] AttachmentModify default  %+v", o._statusCode, o.Payload)
}
func (o *AttachmentModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttachmentModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
