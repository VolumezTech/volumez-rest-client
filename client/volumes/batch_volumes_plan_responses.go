// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// BatchVolumesPlanReader is a Reader for the BatchVolumesPlan structure.
type BatchVolumesPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchVolumesPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchVolumesPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBatchVolumesPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBatchVolumesPlanOK creates a BatchVolumesPlanOK with default headers values
func NewBatchVolumesPlanOK() *BatchVolumesPlanOK {
	return &BatchVolumesPlanOK{}
}

/*
	BatchVolumesPlanOK describes a response with status code 200, with default header values.

volume plan output
*/
type BatchVolumesPlanOK struct {
	Payload *models.VolumePlanOutput
}

func (o *BatchVolumesPlanOK) Error() string {
	return fmt.Sprintf("[POST /volumes/plan][%d] batchVolumesPlanOK  %+v", 200, o.Payload)
}
func (o *BatchVolumesPlanOK) GetPayload() *models.VolumePlanOutput {
	return o.Payload
}

func (o *BatchVolumesPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumePlanOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchVolumesPlanDefault creates a BatchVolumesPlanDefault with default headers values
func NewBatchVolumesPlanDefault(code int) *BatchVolumesPlanDefault {
	return &BatchVolumesPlanDefault{
		_statusCode: code,
	}
}

/*
	BatchVolumesPlanDefault describes a response with status code -1, with default header values.

Error creating new snapshot
*/
type BatchVolumesPlanDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the batch volumes plan default response
func (o *BatchVolumesPlanDefault) Code() int {
	return o._statusCode
}

func (o *BatchVolumesPlanDefault) Error() string {
	return fmt.Sprintf("[POST /volumes/plan][%d] BatchVolumesPlan default  %+v", o._statusCode, o.Payload)
}
func (o *BatchVolumesPlanDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchVolumesPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
BatchVolumesPlanBody batch volumes plan body
swagger:model BatchVolumesPlanBody
*/
type BatchVolumesPlanBody struct {

	// capacity groups to create the volume from (optional)
	CapacityGroup string `json:"CapacityGroup,omitempty"`

	// policy parameter for all volumes that dont have policy parameter defined in their input
	Policy string `json:"DefaultPolicy,omitempty"`

	// volumes
	// Min Items: 1
	Volumes []*VolumePlanInput `json:"Volumes"`

	// zone parameter for all volumes that dont have zone parameter
	Zone string `json:"DefaultZone,omitempty"`
}

// Validate validates this batch volumes plan body
func (o *BatchVolumesPlanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BatchVolumesPlanBody) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(o.Volumes) { // not required
		return nil
	}

	iVolumesSize := int64(len(o.Volumes))

	if err := validate.MinItems("body"+"."+"Volumes", "body", iVolumesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(o.Volumes); i++ {
		if swag.IsZero(o.Volumes[i]) { // not required
			continue
		}

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this batch volumes plan body based on the context it is used
func (o *BatchVolumesPlanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BatchVolumesPlanBody) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Volumes); i++ {

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BatchVolumesPlanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BatchVolumesPlanBody) UnmarshalBinary(b []byte) error {
	var res BatchVolumesPlanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumePlanInput volume plan input
swagger:model VolumePlanInput
*/
type VolumePlanInput struct {

	// policy name for the planned volume
	Policy string `json:"Policy,omitempty"`

	// volume size
	//
	// volume size in GiB
	// Example: 10
	// Required: true
	Size int64 `json:"Size"`

	// zone
	// Example: us-east-1a
	Zone string `json:"Zone,omitempty"`
}

// Validate validates this volume plan input
func (o *VolumePlanInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumePlanInput) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", int64(o.Size)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume plan input based on context it is used
func (o *VolumePlanInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumePlanInput) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumePlanInput) UnmarshalBinary(b []byte) error {
	var res VolumePlanInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
