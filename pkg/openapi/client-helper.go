/*
Volumez API

Volumez orchestrator API

API version: 1.0.0 - f3a04f74
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

func NewGenericOpenAPIError(body []byte, error string, model interface{}) *GenericOpenAPIError {
	return &GenericOpenAPIError{body: body, error: error, model: model}
}

func NewGenericOpenAPIErrorFromError(err error) *GenericOpenAPIError {
	if err == nil {
		return nil
	}
	if err, ok := err.(*GenericOpenAPIError); ok {
		return err
	}
	return &GenericOpenAPIError{error: err.Error()}
}
