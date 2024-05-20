// Code generated by go-swagger; DO NOT EDIT.

package tenant_token

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

// TenantTokenReader is a Reader for the TenantToken structure.
type TenantTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenantTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenantTokenOK creates a TenantTokenOK with default headers values
func NewTenantTokenOK() *TenantTokenOK {
	return &TenantTokenOK{}
}

/*
TenantTokenOK describes a response with status code 200, with default header values.

Tenant's Token retrieval was successful
*/
type TenantTokenOK struct {
	Payload *models.TenantTokenResponse
}

// IsSuccess returns true when this tenant token o k response has a 2xx status code
func (o *TenantTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenant token o k response has a 3xx status code
func (o *TenantTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenant token o k response has a 4xx status code
func (o *TenantTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenant token o k response has a 5xx status code
func (o *TenantTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenant token o k response a status code equal to that given
func (o *TenantTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenant token o k response
func (o *TenantTokenOK) Code() int {
	return 200
}

func (o *TenantTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenant/token][%d] tenantTokenOK %s", 200, payload)
}

func (o *TenantTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenant/token][%d] tenantTokenOK %s", 200, payload)
}

func (o *TenantTokenOK) GetPayload() *models.TenantTokenResponse {
	return o.Payload
}

func (o *TenantTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantTokenDefault creates a TenantTokenDefault with default headers values
func NewTenantTokenDefault(code int) *TenantTokenDefault {
	return &TenantTokenDefault{
		_statusCode: code,
	}
}

/*
TenantTokenDefault describes a response with status code -1, with default header values.

Error getting Tenant's Token
*/
type TenantTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenant token default response has a 2xx status code
func (o *TenantTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenant token default response has a 3xx status code
func (o *TenantTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenant token default response has a 4xx status code
func (o *TenantTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenant token default response has a 5xx status code
func (o *TenantTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenant token default response a status code equal to that given
func (o *TenantTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenant token default response
func (o *TenantTokenDefault) Code() int {
	return o._statusCode
}

func (o *TenantTokenDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenant/token][%d] TenantToken default %s", o._statusCode, payload)
}

func (o *TenantTokenDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenant/token][%d] TenantToken default %s", o._statusCode, payload)
}

func (o *TenantTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
