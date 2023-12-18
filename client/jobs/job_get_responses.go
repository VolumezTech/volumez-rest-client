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

// JobGetReader is a Reader for the JobGet structure.
type JobGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewJobGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewJobGetOK creates a JobGetOK with default headers values
func NewJobGetOK() *JobGetOK {
	return &JobGetOK{}
}

/*
JobGetOK describes a response with status code 200, with default header values.

Properties of a job
*/
type JobGetOK struct {
	Payload *models.Job
}

func (o *JobGetOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{job}][%d] jobGetOK  %+v", 200, o.Payload)
}
func (o *JobGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *JobGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJobGetDefault creates a JobGetDefault with default headers values
func NewJobGetDefault(code int) *JobGetDefault {
	return &JobGetDefault{
		_statusCode: code,
	}
}

/*
JobGetDefault describes a response with status code -1, with default header values.

Error getting properties of a job
*/
type JobGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the job get default response
func (o *JobGetDefault) Code() int {
	return o._statusCode
}

func (o *JobGetDefault) Error() string {
	return fmt.Sprintf("[GET /jobs/{job}][%d] JobGet default  %+v", o._statusCode, o.Payload)
}
func (o *JobGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *JobGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
