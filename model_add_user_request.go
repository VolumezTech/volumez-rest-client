/*
GitHub

Volumez orchestrator API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package volumezopenapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AddUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserRequest{}

// AddUserRequest struct for AddUserRequest
type AddUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	InvitedUserToken string `json:"invitedUserToken"`
}

type _AddUserRequest AddUserRequest

// NewAddUserRequest instantiates a new AddUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserRequest(email string, password string, name string, invitedUserToken string) *AddUserRequest {
	this := AddUserRequest{}
	this.Email = email
	this.Password = password
	this.Name = name
	this.InvitedUserToken = invitedUserToken
	return &this
}

// NewAddUserRequestWithDefaults instantiates a new AddUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserRequestWithDefaults() *AddUserRequest {
	this := AddUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *AddUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *AddUserRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AddUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AddUserRequest) SetPassword(v string) {
	o.Password = v
}

// GetName returns the Name field value
func (o *AddUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddUserRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddUserRequest) SetName(v string) {
	o.Name = v
}

// GetInvitedUserToken returns the InvitedUserToken field value
func (o *AddUserRequest) GetInvitedUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvitedUserToken
}

// GetInvitedUserTokenOk returns a tuple with the InvitedUserToken field value
// and a boolean to check if the value has been set.
func (o *AddUserRequest) GetInvitedUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitedUserToken, true
}

// SetInvitedUserToken sets field value
func (o *AddUserRequest) SetInvitedUserToken(v string) {
	o.InvitedUserToken = v
}

func (o AddUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	toSerialize["name"] = o.Name
	toSerialize["invitedUserToken"] = o.InvitedUserToken
	return toSerialize, nil
}

func (o *AddUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
		"name",
		"invitedUserToken",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddUserRequest := _AddUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddUserRequest)

	if err != nil {
		return err
	}

	*o = AddUserRequest(varAddUserRequest)

	return err
}

type NullableAddUserRequest struct {
	value *AddUserRequest
	isSet bool
}

func (v NullableAddUserRequest) Get() *AddUserRequest {
	return v.value
}

func (v *NullableAddUserRequest) Set(val *AddUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserRequest(val *AddUserRequest) *NullableAddUserRequest {
	return &NullableAddUserRequest{value: val, isSet: true}
}

func (v NullableAddUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
