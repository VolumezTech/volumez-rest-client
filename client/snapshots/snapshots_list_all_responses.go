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

// SnapshotsListAllReader is a Reader for the SnapshotsListAll structure.
type SnapshotsListAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotsListAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotsListAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotsListAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotsListAllOK creates a SnapshotsListAllOK with default headers values
func NewSnapshotsListAllOK() *SnapshotsListAllOK {
	return &SnapshotsListAllOK{}
}

/* SnapshotsListAllOK describes a response with status code 200, with default header values.

List of all snapshots
*/
type SnapshotsListAllOK struct {
	Payload []*models.Snapshot
}

func (o *SnapshotsListAllOK) Error() string {
	return fmt.Sprintf("[GET /snapshots][%d] snapshotsListAllOK  %+v", 200, o.Payload)
}
func (o *SnapshotsListAllOK) GetPayload() []*models.Snapshot {
	return o.Payload
}

func (o *SnapshotsListAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotsListAllDefault creates a SnapshotsListAllDefault with default headers values
func NewSnapshotsListAllDefault(code int) *SnapshotsListAllDefault {
	return &SnapshotsListAllDefault{
		_statusCode: code,
	}
}

/* SnapshotsListAllDefault describes a response with status code -1, with default header values.

Error getting a list of all snapshots
*/
type SnapshotsListAllDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshots list all default response
func (o *SnapshotsListAllDefault) Code() int {
	return o._statusCode
}

func (o *SnapshotsListAllDefault) Error() string {
	return fmt.Sprintf("[GET /snapshots][%d] SnapshotsListAll default  %+v", o._statusCode, o.Payload)
}
func (o *SnapshotsListAllDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotsListAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
