// Code generated by go-swagger; DO NOT EDIT.

package tenant_id_from_tenant_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// TenantIDgetReader is a Reader for the TenantIDget structure.
type TenantIDgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantIDgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantIDgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTenantIDgetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenantIDgetOK creates a TenantIDgetOK with default headers values
func NewTenantIDgetOK() *TenantIDgetOK {
	return &TenantIDgetOK{}
}

/* TenantIDgetOK describes a response with status code 200, with default header values.

Got Tenant's ID successfully
*/
type TenantIDgetOK struct {
	Payload *models.GetTenantIDResponse
}

func (o *TenantIDgetOK) Error() string {
	return fmt.Sprintf("[GET /tenant/tenantid][%d] tenantIDgetOK  %+v", 200, o.Payload)
}
func (o *TenantIDgetOK) GetPayload() *models.GetTenantIDResponse {
	return o.Payload
}

func (o *TenantIDgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTenantIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantIDgetInternalServerError creates a TenantIDgetInternalServerError with default headers values
func NewTenantIDgetInternalServerError() *TenantIDgetInternalServerError {
	return &TenantIDgetInternalServerError{}
}

/* TenantIDgetInternalServerError describes a response with status code 500, with default header values.

Error getting Tenant's ID
*/
type TenantIDgetInternalServerError struct {
	Payload *models.GetTenantIDResponse
}

func (o *TenantIDgetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenant/tenantid][%d] tenantIDgetInternalServerError  %+v", 500, o.Payload)
}
func (o *TenantIDgetInternalServerError) GetPayload() *models.GetTenantIDResponse {
	return o.Payload
}

func (o *TenantIDgetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTenantIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
