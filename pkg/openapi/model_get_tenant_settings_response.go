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

// checks if the GetTenantSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTenantSettingsResponse{}

// GetTenantSettingsResponse struct for GetTenantSettingsResponse
type GetTenantSettingsResponse struct {
	Pricing     NullableGetTenantSettingsResponsePricing     `json:"pricing,omitempty"`
	Declarative NullableGetTenantSettingsResponseDeclarative `json:"declarative,omitempty"`
	Throttle    NullableGetTenantSettingsResponseThrottle    `json:"throttle,omitempty"`
}

// NewGetTenantSettingsResponse instantiates a new GetTenantSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTenantSettingsResponse() *GetTenantSettingsResponse {
	this := GetTenantSettingsResponse{}
	return &this
}

// NewGetTenantSettingsResponseWithDefaults instantiates a new GetTenantSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTenantSettingsResponseWithDefaults() *GetTenantSettingsResponse {
	this := GetTenantSettingsResponse{}
	return &this
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantSettingsResponse) GetPricing() GetTenantSettingsResponsePricing {
	if o == nil || IsNil(o.Pricing.Get()) {
		var ret GetTenantSettingsResponsePricing
		return ret
	}
	return *o.Pricing.Get()
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantSettingsResponse) GetPricingOk() (*GetTenantSettingsResponsePricing, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pricing.Get(), o.Pricing.IsSet()
}

// HasPricing returns a boolean if a field has been set.
func (o *GetTenantSettingsResponse) HasPricing() bool {
	if o != nil && o.Pricing.IsSet() {
		return true
	}

	return false
}

// SetPricing gets a reference to the given NullableGetTenantSettingsResponsePricing and assigns it to the Pricing field.
func (o *GetTenantSettingsResponse) SetPricing(v GetTenantSettingsResponsePricing) {
	o.Pricing.Set(&v)
}

// SetPricingNil sets the value for Pricing to be an explicit nil
func (o *GetTenantSettingsResponse) SetPricingNil() {
	o.Pricing.Set(nil)
}

// UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
func (o *GetTenantSettingsResponse) UnsetPricing() {
	o.Pricing.Unset()
}

// GetDeclarative returns the Declarative field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantSettingsResponse) GetDeclarative() GetTenantSettingsResponseDeclarative {
	if o == nil || IsNil(o.Declarative.Get()) {
		var ret GetTenantSettingsResponseDeclarative
		return ret
	}
	return *o.Declarative.Get()
}

// GetDeclarativeOk returns a tuple with the Declarative field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantSettingsResponse) GetDeclarativeOk() (*GetTenantSettingsResponseDeclarative, bool) {
	if o == nil {
		return nil, false
	}
	return o.Declarative.Get(), o.Declarative.IsSet()
}

// HasDeclarative returns a boolean if a field has been set.
func (o *GetTenantSettingsResponse) HasDeclarative() bool {
	if o != nil && o.Declarative.IsSet() {
		return true
	}

	return false
}

// SetDeclarative gets a reference to the given NullableGetTenantSettingsResponseDeclarative and assigns it to the Declarative field.
func (o *GetTenantSettingsResponse) SetDeclarative(v GetTenantSettingsResponseDeclarative) {
	o.Declarative.Set(&v)
}

// SetDeclarativeNil sets the value for Declarative to be an explicit nil
func (o *GetTenantSettingsResponse) SetDeclarativeNil() {
	o.Declarative.Set(nil)
}

// UnsetDeclarative ensures that no value is present for Declarative, not even an explicit nil
func (o *GetTenantSettingsResponse) UnsetDeclarative() {
	o.Declarative.Unset()
}

// GetThrottle returns the Throttle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantSettingsResponse) GetThrottle() GetTenantSettingsResponseThrottle {
	if o == nil || IsNil(o.Throttle.Get()) {
		var ret GetTenantSettingsResponseThrottle
		return ret
	}
	return *o.Throttle.Get()
}

// GetThrottleOk returns a tuple with the Throttle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantSettingsResponse) GetThrottleOk() (*GetTenantSettingsResponseThrottle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Throttle.Get(), o.Throttle.IsSet()
}

// HasThrottle returns a boolean if a field has been set.
func (o *GetTenantSettingsResponse) HasThrottle() bool {
	if o != nil && o.Throttle.IsSet() {
		return true
	}

	return false
}

// SetThrottle gets a reference to the given NullableGetTenantSettingsResponseThrottle and assigns it to the Throttle field.
func (o *GetTenantSettingsResponse) SetThrottle(v GetTenantSettingsResponseThrottle) {
	o.Throttle.Set(&v)
}

// SetThrottleNil sets the value for Throttle to be an explicit nil
func (o *GetTenantSettingsResponse) SetThrottleNil() {
	o.Throttle.Set(nil)
}

// UnsetThrottle ensures that no value is present for Throttle, not even an explicit nil
func (o *GetTenantSettingsResponse) UnsetThrottle() {
	o.Throttle.Unset()
}

func (o GetTenantSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTenantSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Pricing.IsSet() {
		toSerialize["pricing"] = o.Pricing.Get()
	}
	if o.Declarative.IsSet() {
		toSerialize["declarative"] = o.Declarative.Get()
	}
	if o.Throttle.IsSet() {
		toSerialize["throttle"] = o.Throttle.Get()
	}
	return toSerialize, nil
}

type NullableGetTenantSettingsResponse struct {
	value *GetTenantSettingsResponse
	isSet bool
}

func (v NullableGetTenantSettingsResponse) Get() *GetTenantSettingsResponse {
	return v.value
}

func (v *NullableGetTenantSettingsResponse) Set(val *GetTenantSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTenantSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTenantSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTenantSettingsResponse(val *GetTenantSettingsResponse) *NullableGetTenantSettingsResponse {
	return &NullableGetTenantSettingsResponse{value: val, isSet: true}
}

func (v NullableGetTenantSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTenantSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
