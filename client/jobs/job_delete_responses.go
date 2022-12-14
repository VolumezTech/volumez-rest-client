// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// JobDeleteReader is a Reader for the JobDelete structure.
type JobDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewJobDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewJobDeleteOK creates a JobDeleteOK with default headers values
func NewJobDeleteOK() *JobDeleteOK {
	return &JobDeleteOK{}
}

/* JobDeleteOK describes a response with status code 200, with default header values.

A job was deleted successfully
*/
type JobDeleteOK struct {
	Payload *models.RegularResponse
}

func (o *JobDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /jobs/{job}][%d] jobDeleteOK  %+v", 200, o.Payload)
}
func (o *JobDeleteOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *JobDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJobDeleteDefault creates a JobDeleteDefault with default headers values
func NewJobDeleteDefault(code int) *JobDeleteDefault {
	return &JobDeleteDefault{
		_statusCode: code,
	}
}

/* JobDeleteDefault describes a response with status code -1, with default header values.

Error deleting a job
*/
type JobDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the job delete default response
func (o *JobDeleteDefault) Code() int {
	return o._statusCode
}

func (o *JobDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /jobs/{job}][%d] JobDelete default  %+v", o._statusCode, o.Payload)
}
func (o *JobDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *JobDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
