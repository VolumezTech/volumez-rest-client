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

// checks if the AzureVMRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureVMRegion{}

// AzureVMRegion struct for AzureVMRegion
type AzureVMRegion struct {
	RegionDisplayName *string  `json:"regionDisplayName,omitempty"`
	Zones             []string `json:"zones,omitempty"`
}

// NewAzureVMRegion instantiates a new AzureVMRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureVMRegion() *AzureVMRegion {
	this := AzureVMRegion{}
	return &this
}

// NewAzureVMRegionWithDefaults instantiates a new AzureVMRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureVMRegionWithDefaults() *AzureVMRegion {
	this := AzureVMRegion{}
	return &this
}

// GetRegionDisplayName returns the RegionDisplayName field value if set, zero value otherwise.
func (o *AzureVMRegion) GetRegionDisplayName() string {
	if o == nil || IsNil(o.RegionDisplayName) {
		var ret string
		return ret
	}
	return *o.RegionDisplayName
}

// GetRegionDisplayNameOk returns a tuple with the RegionDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureVMRegion) GetRegionDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionDisplayName) {
		return nil, false
	}
	return o.RegionDisplayName, true
}

// HasRegionDisplayName returns a boolean if a field has been set.
func (o *AzureVMRegion) HasRegionDisplayName() bool {
	if o != nil && !IsNil(o.RegionDisplayName) {
		return true
	}

	return false
}

// SetRegionDisplayName gets a reference to the given string and assigns it to the RegionDisplayName field.
func (o *AzureVMRegion) SetRegionDisplayName(v string) {
	o.RegionDisplayName = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *AzureVMRegion) GetZones() []string {
	if o == nil || IsNil(o.Zones) {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureVMRegion) GetZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *AzureVMRegion) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *AzureVMRegion) SetZones(v []string) {
	o.Zones = v
}

func (o AzureVMRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureVMRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegionDisplayName) {
		toSerialize["regionDisplayName"] = o.RegionDisplayName
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableAzureVMRegion struct {
	value *AzureVMRegion
	isSet bool
}

func (v NullableAzureVMRegion) Get() *AzureVMRegion {
	return v.value
}

func (v *NullableAzureVMRegion) Set(val *AzureVMRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureVMRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureVMRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureVMRegion(val *AzureVMRegion) *NullableAzureVMRegion {
	return &NullableAzureVMRegion{value: val, isSet: true}
}

func (v NullableAzureVMRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureVMRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
