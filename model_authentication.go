/*
GitHub

Volumez orchestrator API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package volumezopenapiclient

import (
	"encoding/json"
)

// checks if the Authentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authentication{}

// Authentication struct for Authentication
type Authentication struct {
	Method *string             `json:"method,omitempty"`
	Chap   *AuthenticationChap `json:"chap,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *Authentication) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *Authentication) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *Authentication) SetMethod(v string) {
	o.Method = &v
}

// GetChap returns the Chap field value if set, zero value otherwise.
func (o *Authentication) GetChap() AuthenticationChap {
	if o == nil || IsNil(o.Chap) {
		var ret AuthenticationChap
		return ret
	}
	return *o.Chap
}

// GetChapOk returns a tuple with the Chap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetChapOk() (*AuthenticationChap, bool) {
	if o == nil || IsNil(o.Chap) {
		return nil, false
	}
	return o.Chap, true
}

// HasChap returns a boolean if a field has been set.
func (o *Authentication) HasChap() bool {
	if o != nil && !IsNil(o.Chap) {
		return true
	}

	return false
}

// SetChap gets a reference to the given AuthenticationChap and assigns it to the Chap field.
func (o *Authentication) SetChap(v AuthenticationChap) {
	o.Chap = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Chap) {
		toSerialize["chap"] = o.Chap
	}
	return toSerialize, nil
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
