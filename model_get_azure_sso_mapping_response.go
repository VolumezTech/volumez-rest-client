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

// checks if the GetAzureSSOMappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAzureSSOMappingResponse{}

// GetAzureSSOMappingResponse struct for GetAzureSSOMappingResponse
type GetAzureSSOMappingResponse struct {
	SecurityGroupID string `json:"SecurityGroupID"`
}

type _GetAzureSSOMappingResponse GetAzureSSOMappingResponse

// NewGetAzureSSOMappingResponse instantiates a new GetAzureSSOMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAzureSSOMappingResponse(securityGroupID string) *GetAzureSSOMappingResponse {
	this := GetAzureSSOMappingResponse{}
	this.SecurityGroupID = securityGroupID
	return &this
}

// NewGetAzureSSOMappingResponseWithDefaults instantiates a new GetAzureSSOMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAzureSSOMappingResponseWithDefaults() *GetAzureSSOMappingResponse {
	this := GetAzureSSOMappingResponse{}
	return &this
}

// GetSecurityGroupID returns the SecurityGroupID field value
func (o *GetAzureSSOMappingResponse) GetSecurityGroupID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityGroupID
}

// GetSecurityGroupIDOk returns a tuple with the SecurityGroupID field value
// and a boolean to check if the value has been set.
func (o *GetAzureSSOMappingResponse) GetSecurityGroupIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityGroupID, true
}

// SetSecurityGroupID sets field value
func (o *GetAzureSSOMappingResponse) SetSecurityGroupID(v string) {
	o.SecurityGroupID = v
}

func (o GetAzureSSOMappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAzureSSOMappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SecurityGroupID"] = o.SecurityGroupID
	return toSerialize, nil
}

func (o *GetAzureSSOMappingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"SecurityGroupID",
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

	varGetAzureSSOMappingResponse := _GetAzureSSOMappingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAzureSSOMappingResponse)

	if err != nil {
		return err
	}

	*o = GetAzureSSOMappingResponse(varGetAzureSSOMappingResponse)

	return err
}

type NullableGetAzureSSOMappingResponse struct {
	value *GetAzureSSOMappingResponse
	isSet bool
}

func (v NullableGetAzureSSOMappingResponse) Get() *GetAzureSSOMappingResponse {
	return v.value
}

func (v *NullableGetAzureSSOMappingResponse) Set(val *GetAzureSSOMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAzureSSOMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAzureSSOMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAzureSSOMappingResponse(val *GetAzureSSOMappingResponse) *NullableGetAzureSSOMappingResponse {
	return &NullableGetAzureSSOMappingResponse{value: val, isSet: true}
}

func (v NullableGetAzureSSOMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAzureSSOMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
