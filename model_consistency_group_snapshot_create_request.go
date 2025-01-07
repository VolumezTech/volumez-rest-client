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

// checks if the ConsistencyGroupSnapshotCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsistencyGroupSnapshotCreateRequest{}

// ConsistencyGroupSnapshotCreateRequest struct for ConsistencyGroupSnapshotCreateRequest
type ConsistencyGroupSnapshotCreateRequest struct {
	Name        *string  `json:"name,omitempty"`
	Consistency *string  `json:"consistency,omitempty"`
	GroupName   *string  `json:"group_name,omitempty"`
	Volumes     []string `json:"volumes,omitempty"`
}

// NewConsistencyGroupSnapshotCreateRequest instantiates a new ConsistencyGroupSnapshotCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistencyGroupSnapshotCreateRequest() *ConsistencyGroupSnapshotCreateRequest {
	this := ConsistencyGroupSnapshotCreateRequest{}
	return &this
}

// NewConsistencyGroupSnapshotCreateRequestWithDefaults instantiates a new ConsistencyGroupSnapshotCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistencyGroupSnapshotCreateRequestWithDefaults() *ConsistencyGroupSnapshotCreateRequest {
	this := ConsistencyGroupSnapshotCreateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConsistencyGroupSnapshotCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConsistencyGroupSnapshotCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetConsistency returns the Consistency field value if set, zero value otherwise.
func (o *ConsistencyGroupSnapshotCreateRequest) GetConsistency() string {
	if o == nil || IsNil(o.Consistency) {
		var ret string
		return ret
	}
	return *o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) GetConsistencyOk() (*string, bool) {
	if o == nil || IsNil(o.Consistency) {
		return nil, false
	}
	return o.Consistency, true
}

// HasConsistency returns a boolean if a field has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) HasConsistency() bool {
	if o != nil && !IsNil(o.Consistency) {
		return true
	}

	return false
}

// SetConsistency gets a reference to the given string and assigns it to the Consistency field.
func (o *ConsistencyGroupSnapshotCreateRequest) SetConsistency(v string) {
	o.Consistency = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ConsistencyGroupSnapshotCreateRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ConsistencyGroupSnapshotCreateRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ConsistencyGroupSnapshotCreateRequest) GetVolumes() []string {
	if o == nil || IsNil(o.Volumes) {
		var ret []string
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) GetVolumesOk() ([]string, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ConsistencyGroupSnapshotCreateRequest) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []string and assigns it to the Volumes field.
func (o *ConsistencyGroupSnapshotCreateRequest) SetVolumes(v []string) {
	o.Volumes = v
}

func (o ConsistencyGroupSnapshotCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsistencyGroupSnapshotCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Consistency) {
		toSerialize["consistency"] = o.Consistency
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableConsistencyGroupSnapshotCreateRequest struct {
	value *ConsistencyGroupSnapshotCreateRequest
	isSet bool
}

func (v NullableConsistencyGroupSnapshotCreateRequest) Get() *ConsistencyGroupSnapshotCreateRequest {
	return v.value
}

func (v *NullableConsistencyGroupSnapshotCreateRequest) Set(val *ConsistencyGroupSnapshotCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistencyGroupSnapshotCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistencyGroupSnapshotCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistencyGroupSnapshotCreateRequest(val *ConsistencyGroupSnapshotCreateRequest) *NullableConsistencyGroupSnapshotCreateRequest {
	return &NullableConsistencyGroupSnapshotCreateRequest{value: val, isSet: true}
}

func (v NullableConsistencyGroupSnapshotCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistencyGroupSnapshotCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
