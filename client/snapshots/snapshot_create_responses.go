// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// SnapshotCreateReader is a Reader for the SnapshotCreate structure.
type SnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotCreateOK creates a SnapshotCreateOK with default headers values
func NewSnapshotCreateOK() *SnapshotCreateOK {
	return &SnapshotCreateOK{}
}

/*
SnapshotCreateOK describes a response with status code 200, with default header values.

New snapshot was created successfully
*/
type SnapshotCreateOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this snapshot create o k response has a 2xx status code
func (o *SnapshotCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot create o k response has a 3xx status code
func (o *SnapshotCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot create o k response has a 4xx status code
func (o *SnapshotCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot create o k response has a 5xx status code
func (o *SnapshotCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot create o k response a status code equal to that given
func (o *SnapshotCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshot create o k response
func (o *SnapshotCreateOK) Code() int {
	return 200
}

func (o *SnapshotCreateOK) Error() string {
	return fmt.Sprintf("[POST /volumes/{volume}/snapshots][%d] snapshotCreateOK  %+v", 200, o.Payload)
}

func (o *SnapshotCreateOK) String() string {
	return fmt.Sprintf("[POST /volumes/{volume}/snapshots][%d] snapshotCreateOK  %+v", 200, o.Payload)
}

func (o *SnapshotCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *SnapshotCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotCreateDefault creates a SnapshotCreateDefault with default headers values
func NewSnapshotCreateDefault(code int) *SnapshotCreateDefault {
	return &SnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
SnapshotCreateDefault describes a response with status code -1, with default header values.

Error creating new snapshot
*/
type SnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshot create default response has a 2xx status code
func (o *SnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot create default response has a 3xx status code
func (o *SnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot create default response has a 4xx status code
func (o *SnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot create default response has a 5xx status code
func (o *SnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot create default response a status code equal to that given
func (o *SnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshot create default response
func (o *SnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotCreateDefault) Error() string {
	return fmt.Sprintf("[POST /volumes/{volume}/snapshots][%d] SnapshotCreate default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotCreateDefault) String() string {
	return fmt.Sprintf("[POST /volumes/{volume}/snapshots][%d] SnapshotCreate default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
