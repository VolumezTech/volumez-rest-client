// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// JobResumeSuspendReader is a Reader for the JobResumeSuspend structure.
type JobResumeSuspendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobResumeSuspendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobResumeSuspendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewJobResumeSuspendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewJobResumeSuspendOK creates a JobResumeSuspendOK with default headers values
func NewJobResumeSuspendOK() *JobResumeSuspendOK {
	return &JobResumeSuspendOK{}
}

/*
JobResumeSuspendOK describes a response with status code 200, with default header values.

Operation completed successfully
*/
type JobResumeSuspendOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this job resume suspend o k response has a 2xx status code
func (o *JobResumeSuspendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job resume suspend o k response has a 3xx status code
func (o *JobResumeSuspendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job resume suspend o k response has a 4xx status code
func (o *JobResumeSuspendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job resume suspend o k response has a 5xx status code
func (o *JobResumeSuspendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job resume suspend o k response a status code equal to that given
func (o *JobResumeSuspendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the job resume suspend o k response
func (o *JobResumeSuspendOK) Code() int {
	return 200
}

func (o *JobResumeSuspendOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /jobs/{job}/resume_suspend/{state}][%d] jobResumeSuspendOK %s", 200, payload)
}

func (o *JobResumeSuspendOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /jobs/{job}/resume_suspend/{state}][%d] jobResumeSuspendOK %s", 200, payload)
}

func (o *JobResumeSuspendOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *JobResumeSuspendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJobResumeSuspendDefault creates a JobResumeSuspendDefault with default headers values
func NewJobResumeSuspendDefault(code int) *JobResumeSuspendDefault {
	return &JobResumeSuspendDefault{
		_statusCode: code,
	}
}

/*
JobResumeSuspendDefault describes a response with status code -1, with default header values.

Error resuming or suspending job
*/
type JobResumeSuspendDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this job resume suspend default response has a 2xx status code
func (o *JobResumeSuspendDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this job resume suspend default response has a 3xx status code
func (o *JobResumeSuspendDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this job resume suspend default response has a 4xx status code
func (o *JobResumeSuspendDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this job resume suspend default response has a 5xx status code
func (o *JobResumeSuspendDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this job resume suspend default response a status code equal to that given
func (o *JobResumeSuspendDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the job resume suspend default response
func (o *JobResumeSuspendDefault) Code() int {
	return o._statusCode
}

func (o *JobResumeSuspendDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /jobs/{job}/resume_suspend/{state}][%d] JobResumeSuspend default %s", o._statusCode, payload)
}

func (o *JobResumeSuspendDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /jobs/{job}/resume_suspend/{state}][%d] JobResumeSuspend default %s", o._statusCode, payload)
}

func (o *JobResumeSuspendDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *JobResumeSuspendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
