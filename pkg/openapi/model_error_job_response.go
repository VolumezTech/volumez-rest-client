/*
Volumez API

Volumez orchestrator API

API version: 1.0.0 - f3a04f74
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ErrorJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorJobResponse{}

// ErrorJobResponse struct for ErrorJobResponse
type ErrorJobResponse struct {
	Message   *string `json:"Message,omitempty"`
	ErrorCode *int32  `json:"ErrorCode,omitempty"`
	ObjectID  *string `json:"ObjectID,omitempty"`
	JobID     *string `json:"JobID,omitempty"`
}

// NewErrorJobResponse instantiates a new ErrorJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorJobResponse() *ErrorJobResponse {
	this := ErrorJobResponse{}
	return &this
}

// NewErrorJobResponseWithDefaults instantiates a new ErrorJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorJobResponseWithDefaults() *ErrorJobResponse {
	this := ErrorJobResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorJobResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorJobResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorJobResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorJobResponse) SetMessage(v string) {
	o.Message = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ErrorJobResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorJobResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorJobResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *ErrorJobResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetObjectID returns the ObjectID field value if set, zero value otherwise.
func (o *ErrorJobResponse) GetObjectID() string {
	if o == nil || IsNil(o.ObjectID) {
		var ret string
		return ret
	}
	return *o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorJobResponse) GetObjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectID) {
		return nil, false
	}
	return o.ObjectID, true
}

// HasObjectID returns a boolean if a field has been set.
func (o *ErrorJobResponse) HasObjectID() bool {
	if o != nil && !IsNil(o.ObjectID) {
		return true
	}

	return false
}

// SetObjectID gets a reference to the given string and assigns it to the ObjectID field.
func (o *ErrorJobResponse) SetObjectID(v string) {
	o.ObjectID = &v
}

// GetJobID returns the JobID field value if set, zero value otherwise.
func (o *ErrorJobResponse) GetJobID() string {
	if o == nil || IsNil(o.JobID) {
		var ret string
		return ret
	}
	return *o.JobID
}

// GetJobIDOk returns a tuple with the JobID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorJobResponse) GetJobIDOk() (*string, bool) {
	if o == nil || IsNil(o.JobID) {
		return nil, false
	}
	return o.JobID, true
}

// HasJobID returns a boolean if a field has been set.
func (o *ErrorJobResponse) HasJobID() bool {
	if o != nil && !IsNil(o.JobID) {
		return true
	}

	return false
}

// SetJobID gets a reference to the given string and assigns it to the JobID field.
func (o *ErrorJobResponse) SetJobID(v string) {
	o.JobID = &v
}

func (o ErrorJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if !IsNil(o.ObjectID) {
		toSerialize["ObjectID"] = o.ObjectID
	}
	if !IsNil(o.JobID) {
		toSerialize["JobID"] = o.JobID
	}
	return toSerialize, nil
}

type NullableErrorJobResponse struct {
	value *ErrorJobResponse
	isSet bool
}

func (v NullableErrorJobResponse) Get() *ErrorJobResponse {
	return v.value
}

func (v *NullableErrorJobResponse) Set(val *ErrorJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorJobResponse(val *ErrorJobResponse) *NullableErrorJobResponse {
	return &NullableErrorJobResponse{value: val, isSet: true}
}

func (v NullableErrorJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
