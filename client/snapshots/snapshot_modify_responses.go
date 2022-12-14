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

// SnapshotModifyReader is a Reader for the SnapshotModify structure.
type SnapshotModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotModifyOK creates a SnapshotModifyOK with default headers values
func NewSnapshotModifyOK() *SnapshotModifyOK {
	return &SnapshotModifyOK{}
}

/* SnapshotModifyOK describes a response with status code 200, with default header values.

A snapshot was updated successfully
*/
type SnapshotModifyOK struct {
	Payload *models.RegularResponse
}

func (o *SnapshotModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}/snapshots/{snapshot}][%d] snapshotModifyOK  %+v", 200, o.Payload)
}
func (o *SnapshotModifyOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *SnapshotModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotModifyDefault creates a SnapshotModifyDefault with default headers values
func NewSnapshotModifyDefault(code int) *SnapshotModifyDefault {
	return &SnapshotModifyDefault{
		_statusCode: code,
	}
}

/* SnapshotModifyDefault describes a response with status code -1, with default header values.

Error updating a snapshot
*/
type SnapshotModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot modify default response
func (o *SnapshotModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{volume}/snapshots/{snapshot}][%d] SnapshotModify default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
