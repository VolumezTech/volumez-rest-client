// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// SnapshotsListReader is a Reader for the SnapshotsList structure.
type SnapshotsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotsListOK creates a SnapshotsListOK with default headers values
func NewSnapshotsListOK() *SnapshotsListOK {
	return &SnapshotsListOK{}
}

/*
SnapshotsListOK describes a response with status code 200, with default header values.

List of snapshots
*/
type SnapshotsListOK struct {
	Payload []*models.Snapshot
}

// IsSuccess returns true when this snapshots list o k response has a 2xx status code
func (o *SnapshotsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshots list o k response has a 3xx status code
func (o *SnapshotsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshots list o k response has a 4xx status code
func (o *SnapshotsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshots list o k response has a 5xx status code
func (o *SnapshotsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshots list o k response a status code equal to that given
func (o *SnapshotsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapshots list o k response
func (o *SnapshotsListOK) Code() int {
	return 200
}

func (o *SnapshotsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots][%d] snapshotsListOK %s", 200, payload)
}

func (o *SnapshotsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots][%d] snapshotsListOK %s", 200, payload)
}

func (o *SnapshotsListOK) GetPayload() []*models.Snapshot {
	return o.Payload
}

func (o *SnapshotsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotsListDefault creates a SnapshotsListDefault with default headers values
func NewSnapshotsListDefault(code int) *SnapshotsListDefault {
	return &SnapshotsListDefault{
		_statusCode: code,
	}
}

/*
SnapshotsListDefault describes a response with status code -1, with default header values.

Error getting a list of snapshots
*/
type SnapshotsListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapshots list default response has a 2xx status code
func (o *SnapshotsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshots list default response has a 3xx status code
func (o *SnapshotsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshots list default response has a 4xx status code
func (o *SnapshotsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshots list default response has a 5xx status code
func (o *SnapshotsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshots list default response a status code equal to that given
func (o *SnapshotsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapshots list default response
func (o *SnapshotsListDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotsListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots][%d] SnapshotsList default %s", o._statusCode, payload)
}

func (o *SnapshotsListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /volumes/{volume}/snapshots][%d] SnapshotsList default %s", o._statusCode, payload)
}

func (o *SnapshotsListDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
