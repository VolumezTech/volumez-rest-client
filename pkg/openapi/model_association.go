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

// checks if the Association type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Association{}

// Association struct for Association
type Association struct {
	Name         *string `json:"name,omitempty"`
	Id           *string `json:"id,omitempty"`
	Volumename   *string `json:"volumename,omitempty"`
	Volumeid     *string `json:"volumeid,omitempty"`
	Snapshotname *string `json:"snapshotname,omitempty"`
	Snapshotid   *string `json:"snapshotid,omitempty"`
	State        *string `json:"state,omitempty"`
	Status       *string `json:"status,omitempty"`
}

// NewAssociation instantiates a new Association object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociation() *Association {
	this := Association{}
	return &this
}

// NewAssociationWithDefaults instantiates a new Association object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationWithDefaults() *Association {
	this := Association{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Association) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Association) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Association) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Association) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Association) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Association) SetId(v string) {
	o.Id = &v
}

// GetVolumename returns the Volumename field value if set, zero value otherwise.
func (o *Association) GetVolumename() string {
	if o == nil || IsNil(o.Volumename) {
		var ret string
		return ret
	}
	return *o.Volumename
}

// GetVolumenameOk returns a tuple with the Volumename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetVolumenameOk() (*string, bool) {
	if o == nil || IsNil(o.Volumename) {
		return nil, false
	}
	return o.Volumename, true
}

// HasVolumename returns a boolean if a field has been set.
func (o *Association) HasVolumename() bool {
	if o != nil && !IsNil(o.Volumename) {
		return true
	}

	return false
}

// SetVolumename gets a reference to the given string and assigns it to the Volumename field.
func (o *Association) SetVolumename(v string) {
	o.Volumename = &v
}

// GetVolumeid returns the Volumeid field value if set, zero value otherwise.
func (o *Association) GetVolumeid() string {
	if o == nil || IsNil(o.Volumeid) {
		var ret string
		return ret
	}
	return *o.Volumeid
}

// GetVolumeidOk returns a tuple with the Volumeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetVolumeidOk() (*string, bool) {
	if o == nil || IsNil(o.Volumeid) {
		return nil, false
	}
	return o.Volumeid, true
}

// HasVolumeid returns a boolean if a field has been set.
func (o *Association) HasVolumeid() bool {
	if o != nil && !IsNil(o.Volumeid) {
		return true
	}

	return false
}

// SetVolumeid gets a reference to the given string and assigns it to the Volumeid field.
func (o *Association) SetVolumeid(v string) {
	o.Volumeid = &v
}

// GetSnapshotname returns the Snapshotname field value if set, zero value otherwise.
func (o *Association) GetSnapshotname() string {
	if o == nil || IsNil(o.Snapshotname) {
		var ret string
		return ret
	}
	return *o.Snapshotname
}

// GetSnapshotnameOk returns a tuple with the Snapshotname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetSnapshotnameOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshotname) {
		return nil, false
	}
	return o.Snapshotname, true
}

// HasSnapshotname returns a boolean if a field has been set.
func (o *Association) HasSnapshotname() bool {
	if o != nil && !IsNil(o.Snapshotname) {
		return true
	}

	return false
}

// SetSnapshotname gets a reference to the given string and assigns it to the Snapshotname field.
func (o *Association) SetSnapshotname(v string) {
	o.Snapshotname = &v
}

// GetSnapshotid returns the Snapshotid field value if set, zero value otherwise.
func (o *Association) GetSnapshotid() string {
	if o == nil || IsNil(o.Snapshotid) {
		var ret string
		return ret
	}
	return *o.Snapshotid
}

// GetSnapshotidOk returns a tuple with the Snapshotid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetSnapshotidOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshotid) {
		return nil, false
	}
	return o.Snapshotid, true
}

// HasSnapshotid returns a boolean if a field has been set.
func (o *Association) HasSnapshotid() bool {
	if o != nil && !IsNil(o.Snapshotid) {
		return true
	}

	return false
}

// SetSnapshotid gets a reference to the given string and assigns it to the Snapshotid field.
func (o *Association) SetSnapshotid(v string) {
	o.Snapshotid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Association) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Association) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Association) SetState(v string) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Association) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Association) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Association) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Association) SetStatus(v string) {
	o.Status = &v
}

func (o Association) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Association) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Volumename) {
		toSerialize["volumename"] = o.Volumename
	}
	if !IsNil(o.Volumeid) {
		toSerialize["volumeid"] = o.Volumeid
	}
	if !IsNil(o.Snapshotname) {
		toSerialize["snapshotname"] = o.Snapshotname
	}
	if !IsNil(o.Snapshotid) {
		toSerialize["snapshotid"] = o.Snapshotid
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAssociation struct {
	value *Association
	isSet bool
}

func (v NullableAssociation) Get() *Association {
	return v.value
}

func (v *NullableAssociation) Set(val *Association) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociation(val *Association) *NullableAssociation {
	return &NullableAssociation{value: val, isSet: true}
}

func (v NullableAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
