/*
Volumez API

Volumez orchestrator API

API version: 1.0.0 - f3a04f74
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SignInResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignInResponse{}

// SignInResponse struct for SignInResponse
type SignInResponse struct {
	AccessToken  string `json:"AccessToken"`
	IdToken      string `json:"IdToken"`
	RefreshToken string `json:"RefreshToken"`
	ExpiresIn    int32  `json:"ExpiresIn"`
	TokenType    string `json:"TokenType"`
}

type _SignInResponse SignInResponse

// NewSignInResponse instantiates a new SignInResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignInResponse(accessToken string, idToken string, refreshToken string, expiresIn int32, tokenType string) *SignInResponse {
	this := SignInResponse{}
	this.AccessToken = accessToken
	this.IdToken = idToken
	this.RefreshToken = refreshToken
	this.ExpiresIn = expiresIn
	this.TokenType = tokenType
	return &this
}

// NewSignInResponseWithDefaults instantiates a new SignInResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignInResponseWithDefaults() *SignInResponse {
	this := SignInResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *SignInResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SignInResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SignInResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetIdToken returns the IdToken field value
func (o *SignInResponse) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *SignInResponse) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *SignInResponse) SetIdToken(v string) {
	o.IdToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *SignInResponse) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *SignInResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *SignInResponse) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *SignInResponse) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *SignInResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *SignInResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetTokenType returns the TokenType field value
func (o *SignInResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *SignInResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *SignInResponse) SetTokenType(v string) {
	o.TokenType = v
}

func (o SignInResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignInResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AccessToken"] = o.AccessToken
	toSerialize["IdToken"] = o.IdToken
	toSerialize["RefreshToken"] = o.RefreshToken
	toSerialize["ExpiresIn"] = o.ExpiresIn
	toSerialize["TokenType"] = o.TokenType
	return toSerialize, nil
}

func (o *SignInResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AccessToken",
		"IdToken",
		"RefreshToken",
		"ExpiresIn",
		"TokenType",
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

	varSignInResponse := _SignInResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignInResponse)

	if err != nil {
		return err
	}

	*o = SignInResponse(varSignInResponse)

	return err
}

type NullableSignInResponse struct {
	value *SignInResponse
	isSet bool
}

func (v NullableSignInResponse) Get() *SignInResponse {
	return v.value
}

func (v *NullableSignInResponse) Set(val *SignInResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignInResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignInResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignInResponse(val *SignInResponse) *NullableSignInResponse {
	return &NullableSignInResponse{value: val, isSet: true}
}

func (v NullableSignInResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignInResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
