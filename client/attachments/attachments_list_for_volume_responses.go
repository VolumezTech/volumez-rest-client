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

// AttachmentsListForVolumeReader is a Reader for the AttachmentsListForVolume structure.
type AttachmentsListForVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachmentsListForVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachmentsListForVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAttachmentsListForVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachmentsListForVolumeOK creates a AttachmentsListForVolumeOK with default headers values
func NewAttachmentsListForVolumeOK() *AttachmentsListForVolumeOK {
	return &AttachmentsListForVolumeOK{}
}

/*
AttachmentsListForVolumeOK describes a response with status code 200, with default header values.

List of attachments
*/
type AttachmentsListForVolumeOK struct {
	Payload []*models.Attachment
}

// IsSuccess returns true when this attachments list for volume o k response has a 2xx status code
func (o *AttachmentsListForVolumeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attachments list for volume o k response has a 3xx status code
func (o *AttachmentsListForVolumeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attachments list for volume o k response has a 4xx status code
func (o *AttachmentsListForVolumeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attachments list for volume o k response has a 5xx status code
func (o *AttachmentsListForVolumeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attachments list for volume o k response a status code equal to that given
func (o *AttachmentsListForVolumeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attachments list for volume o k response
func (o *AttachmentsListForVolumeOK) Code() int {
	return 200
}

func (o *AttachmentsListForVolumeOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/attachments][%d] attachmentsListForVolumeOK  %+v", 200, o.Payload)
}

func (o *AttachmentsListForVolumeOK) String() string {
	return fmt.Sprintf("[GET /volumes/{volume}/attachments][%d] attachmentsListForVolumeOK  %+v", 200, o.Payload)
}

func (o *AttachmentsListForVolumeOK) GetPayload() []*models.Attachment {
	return o.Payload
}

func (o *AttachmentsListForVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachmentsListForVolumeDefault creates a AttachmentsListForVolumeDefault with default headers values
func NewAttachmentsListForVolumeDefault(code int) *AttachmentsListForVolumeDefault {
	return &AttachmentsListForVolumeDefault{
		_statusCode: code,
	}
}

/*
AttachmentsListForVolumeDefault describes a response with status code -1, with default header values.

Error getting a list of attachments
*/
type AttachmentsListForVolumeDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this attachments list for volume default response has a 2xx status code
func (o *AttachmentsListForVolumeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attachments list for volume default response has a 3xx status code
func (o *AttachmentsListForVolumeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attachments list for volume default response has a 4xx status code
func (o *AttachmentsListForVolumeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attachments list for volume default response has a 5xx status code
func (o *AttachmentsListForVolumeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attachments list for volume default response a status code equal to that given
func (o *AttachmentsListForVolumeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attachments list for volume default response
func (o *AttachmentsListForVolumeDefault) Code() int {
	return o._statusCode
}

func (o *AttachmentsListForVolumeDefault) Error() string {
	return fmt.Sprintf("[GET /volumes/{volume}/attachments][%d] AttachmentsListForVolume default  %+v", o._statusCode, o.Payload)
}

func (o *AttachmentsListForVolumeDefault) String() string {
	return fmt.Sprintf("[GET /volumes/{volume}/attachments][%d] AttachmentsListForVolume default  %+v", o._statusCode, o.Payload)
}

func (o *AttachmentsListForVolumeDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AttachmentsListForVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
