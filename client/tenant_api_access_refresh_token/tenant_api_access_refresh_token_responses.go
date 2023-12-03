// Code generated by go-swagger; DO NOT EDIT.

package tenant_api_access_refresh_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// TenantAPIAccessRefreshTokenReader is a Reader for the TenantAPIAccessRefreshToken structure.
type TenantAPIAccessRefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantAPIAccessRefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantAPIAccessRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTenantAPIAccessRefreshTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tenant/apiaccess/credentials/refresh] TenantAPIAccessRefreshToken", response, response.Code())
	}
}

// NewTenantAPIAccessRefreshTokenOK creates a TenantAPIAccessRefreshTokenOK with default headers values
func NewTenantAPIAccessRefreshTokenOK() *TenantAPIAccessRefreshTokenOK {
	return &TenantAPIAccessRefreshTokenOK{}
}

/*
TenantAPIAccessRefreshTokenOK describes a response with status code 200, with default header values.

200 response
*/
type TenantAPIAccessRefreshTokenOK struct {
	AccessControlAllowOrigin string

	Payload *models.RefreshTokenResponse
}

// IsSuccess returns true when this tenant Api access refresh token o k response has a 2xx status code
func (o *TenantAPIAccessRefreshTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenant Api access refresh token o k response has a 3xx status code
func (o *TenantAPIAccessRefreshTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenant Api access refresh token o k response has a 4xx status code
func (o *TenantAPIAccessRefreshTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenant Api access refresh token o k response has a 5xx status code
func (o *TenantAPIAccessRefreshTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenant Api access refresh token o k response a status code equal to that given
func (o *TenantAPIAccessRefreshTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenant Api access refresh token o k response
func (o *TenantAPIAccessRefreshTokenOK) Code() int {
	return 200
}

func (o *TenantAPIAccessRefreshTokenOK) Error() string {
	return fmt.Sprintf("[GET /tenant/apiaccess/credentials/refresh][%d] tenantApiAccessRefreshTokenOK  %+v", 200, o.Payload)
}

func (o *TenantAPIAccessRefreshTokenOK) String() string {
	return fmt.Sprintf("[GET /tenant/apiaccess/credentials/refresh][%d] tenantApiAccessRefreshTokenOK  %+v", 200, o.Payload)
}

func (o *TenantAPIAccessRefreshTokenOK) GetPayload() *models.RefreshTokenResponse {
	return o.Payload
}

func (o *TenantAPIAccessRefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Access-Control-Allow-Origin
	hdrAccessControlAllowOrigin := response.GetHeader("Access-Control-Allow-Origin")

	if hdrAccessControlAllowOrigin != "" {
		o.AccessControlAllowOrigin = hdrAccessControlAllowOrigin
	}

	o.Payload = new(models.RefreshTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantAPIAccessRefreshTokenInternalServerError creates a TenantAPIAccessRefreshTokenInternalServerError with default headers values
func NewTenantAPIAccessRefreshTokenInternalServerError() *TenantAPIAccessRefreshTokenInternalServerError {
	return &TenantAPIAccessRefreshTokenInternalServerError{}
}

/*
TenantAPIAccessRefreshTokenInternalServerError describes a response with status code 500, with default header values.

Error refreshing user Tenant's API Acces Token
*/
type TenantAPIAccessRefreshTokenInternalServerError struct {
	AccessControlAllowOrigin string

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenant Api access refresh token internal server error response has a 2xx status code
func (o *TenantAPIAccessRefreshTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenant Api access refresh token internal server error response has a 3xx status code
func (o *TenantAPIAccessRefreshTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenant Api access refresh token internal server error response has a 4xx status code
func (o *TenantAPIAccessRefreshTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenant Api access refresh token internal server error response has a 5xx status code
func (o *TenantAPIAccessRefreshTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tenant Api access refresh token internal server error response a status code equal to that given
func (o *TenantAPIAccessRefreshTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the tenant Api access refresh token internal server error response
func (o *TenantAPIAccessRefreshTokenInternalServerError) Code() int {
	return 500
}

func (o *TenantAPIAccessRefreshTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenant/apiaccess/credentials/refresh][%d] tenantApiAccessRefreshTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantAPIAccessRefreshTokenInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenant/apiaccess/credentials/refresh][%d] tenantApiAccessRefreshTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantAPIAccessRefreshTokenInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantAPIAccessRefreshTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Access-Control-Allow-Origin
	hdrAccessControlAllowOrigin := response.GetHeader("Access-Control-Allow-Origin")

	if hdrAccessControlAllowOrigin != "" {
		o.AccessControlAllowOrigin = hdrAccessControlAllowOrigin
	}

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
