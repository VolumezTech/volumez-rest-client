// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// ConsistencyGroupSnapshotCreateReader is a Reader for the ConsistencyGroupSnapshotCreate structure.
type ConsistencyGroupSnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupSnapshotCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotCreateOK creates a ConsistencyGroupSnapshotCreateOK with default headers values
func NewConsistencyGroupSnapshotCreateOK() *ConsistencyGroupSnapshotCreateOK {
	return &ConsistencyGroupSnapshotCreateOK{}
}

/*
ConsistencyGroupSnapshotCreateOK describes a response with status code 200, with default header values.

New snapshot was created successfully
*/
type ConsistencyGroupSnapshotCreateOK struct {
	Payload *models.RegularResponse
}

// IsSuccess returns true when this consistency group snapshot create o k response has a 2xx status code
func (o *ConsistencyGroupSnapshotCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group snapshot create o k response has a 3xx status code
func (o *ConsistencyGroupSnapshotCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group snapshot create o k response has a 4xx status code
func (o *ConsistencyGroupSnapshotCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group snapshot create o k response has a 5xx status code
func (o *ConsistencyGroupSnapshotCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group snapshot create o k response a status code equal to that given
func (o *ConsistencyGroupSnapshotCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group snapshot create o k response
func (o *ConsistencyGroupSnapshotCreateOK) Code() int {
	return 200
}

func (o *ConsistencyGroupSnapshotCreateOK) Error() string {
	return fmt.Sprintf("[POST /volumes/snapshot][%d] consistencyGroupSnapshotCreateOK  %+v", 200, o.Payload)
}

func (o *ConsistencyGroupSnapshotCreateOK) String() string {
	return fmt.Sprintf("[POST /volumes/snapshot][%d] consistencyGroupSnapshotCreateOK  %+v", 200, o.Payload)
}

func (o *ConsistencyGroupSnapshotCreateOK) GetPayload() *models.RegularResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegularResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupSnapshotCreateDefault creates a ConsistencyGroupSnapshotCreateDefault with default headers values
func NewConsistencyGroupSnapshotCreateDefault(code int) *ConsistencyGroupSnapshotCreateDefault {
	return &ConsistencyGroupSnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
ConsistencyGroupSnapshotCreateDefault describes a response with status code -1, with default header values.

Error creating new snapshot
*/
type ConsistencyGroupSnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group snapshot create default response has a 2xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group snapshot create default response has a 3xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group snapshot create default response has a 4xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group snapshot create default response has a 5xx status code
func (o *ConsistencyGroupSnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group snapshot create default response a status code equal to that given
func (o *ConsistencyGroupSnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group snapshot create default response
func (o *ConsistencyGroupSnapshotCreateDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotCreateDefault) Error() string {
	return fmt.Sprintf("[POST /volumes/snapshot][%d] ConsistencyGroupSnapshotCreate default  %+v", o._statusCode, o.Payload)
}

func (o *ConsistencyGroupSnapshotCreateDefault) String() string {
	return fmt.Sprintf("[POST /volumes/snapshot][%d] ConsistencyGroupSnapshotCreate default  %+v", o._statusCode, o.Payload)
}

func (o *ConsistencyGroupSnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ConsistencyGroupSnapshotCreateBody consistency group snapshot create body
swagger:model ConsistencyGroupSnapshotCreateBody
*/
type ConsistencyGroupSnapshotCreateBody struct {

	// consistency
	// Enum: [crash application]
	Consistency *string `json:"consistency,omitempty"`

	// group name
	GroupName string `json:"group_name"`

	// name
	Name *string `json:"name,omitempty"`

	// volumes
	Volumes []string `json:"volumes"`
}

// Validate validates this consistency group snapshot create body
func (o *ConsistencyGroupSnapshotCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConsistency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consistencyGroupSnapshotCreateBodyTypeConsistencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupSnapshotCreateBodyTypeConsistencyPropEnum = append(consistencyGroupSnapshotCreateBodyTypeConsistencyPropEnum, v)
	}
}

const (

	// ConsistencyGroupSnapshotCreateBodyConsistencyCrash captures enum value "crash"
	ConsistencyGroupSnapshotCreateBodyConsistencyCrash string = "crash"

	// ConsistencyGroupSnapshotCreateBodyConsistencyApplication captures enum value "application"
	ConsistencyGroupSnapshotCreateBodyConsistencyApplication string = "application"
)

// prop value enum
func (o *ConsistencyGroupSnapshotCreateBody) validateConsistencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupSnapshotCreateBodyTypeConsistencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ConsistencyGroupSnapshotCreateBody) validateConsistency(formats strfmt.Registry) error {
	if swag.IsZero(o.Consistency) { // not required
		return nil
	}

	// value enum
	if err := o.validateConsistencyEnum("body"+"."+"consistency", "body", *o.Consistency); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group snapshot create body based on context it is used
func (o *ConsistencyGroupSnapshotCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConsistencyGroupSnapshotCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConsistencyGroupSnapshotCreateBody) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
