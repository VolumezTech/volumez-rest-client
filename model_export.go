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

// checks if the Export type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Export{}

// Export struct for Export
type Export struct {
	Id           *string       `json:"id,omitempty"`
	Params       *ExportCreate `json:"params,omitempty"`
	Volumename   *string       `json:"volumename,omitempty"`
	Snapshotname *string       `json:"snapshotname,omitempty"`
	State        *string       `json:"state,omitempty"`
	Status       *string       `json:"status,omitempty"`
	Progress     *int32        `json:"progress,omitempty"`
	Xqn          *string       `json:"xqn,omitempty"`
	Wwn          *string       `json:"wwn,omitempty"`
}

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport() *Export {
	this := Export{}
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Export) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Export) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Export) SetId(v string) {
	o.Id = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Export) GetParams() ExportCreate {
	if o == nil || IsNil(o.Params) {
		var ret ExportCreate
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetParamsOk() (*ExportCreate, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Export) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given ExportCreate and assigns it to the Params field.
func (o *Export) SetParams(v ExportCreate) {
	o.Params = &v
}

// GetVolumename returns the Volumename field value if set, zero value otherwise.
func (o *Export) GetVolumename() string {
	if o == nil || IsNil(o.Volumename) {
		var ret string
		return ret
	}
	return *o.Volumename
}

// GetVolumenameOk returns a tuple with the Volumename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetVolumenameOk() (*string, bool) {
	if o == nil || IsNil(o.Volumename) {
		return nil, false
	}
	return o.Volumename, true
}

// HasVolumename returns a boolean if a field has been set.
func (o *Export) HasVolumename() bool {
	if o != nil && !IsNil(o.Volumename) {
		return true
	}

	return false
}

// SetVolumename gets a reference to the given string and assigns it to the Volumename field.
func (o *Export) SetVolumename(v string) {
	o.Volumename = &v
}

// GetSnapshotname returns the Snapshotname field value if set, zero value otherwise.
func (o *Export) GetSnapshotname() string {
	if o == nil || IsNil(o.Snapshotname) {
		var ret string
		return ret
	}
	return *o.Snapshotname
}

// GetSnapshotnameOk returns a tuple with the Snapshotname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetSnapshotnameOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshotname) {
		return nil, false
	}
	return o.Snapshotname, true
}

// HasSnapshotname returns a boolean if a field has been set.
func (o *Export) HasSnapshotname() bool {
	if o != nil && !IsNil(o.Snapshotname) {
		return true
	}

	return false
}

// SetSnapshotname gets a reference to the given string and assigns it to the Snapshotname field.
func (o *Export) SetSnapshotname(v string) {
	o.Snapshotname = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Export) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Export) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Export) SetState(v string) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Export) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Export) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Export) SetStatus(v string) {
	o.Status = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Export) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Export) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *Export) SetProgress(v int32) {
	o.Progress = &v
}

// GetXqn returns the Xqn field value if set, zero value otherwise.
func (o *Export) GetXqn() string {
	if o == nil || IsNil(o.Xqn) {
		var ret string
		return ret
	}
	return *o.Xqn
}

// GetXqnOk returns a tuple with the Xqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetXqnOk() (*string, bool) {
	if o == nil || IsNil(o.Xqn) {
		return nil, false
	}
	return o.Xqn, true
}

// HasXqn returns a boolean if a field has been set.
func (o *Export) HasXqn() bool {
	if o != nil && !IsNil(o.Xqn) {
		return true
	}

	return false
}

// SetXqn gets a reference to the given string and assigns it to the Xqn field.
func (o *Export) SetXqn(v string) {
	o.Xqn = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *Export) GetWwn() string {
	if o == nil || IsNil(o.Wwn) {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetWwnOk() (*string, bool) {
	if o == nil || IsNil(o.Wwn) {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *Export) HasWwn() bool {
	if o != nil && !IsNil(o.Wwn) {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *Export) SetWwn(v string) {
	o.Wwn = &v
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Export) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Volumename) {
		toSerialize["volumename"] = o.Volumename
	}
	if !IsNil(o.Snapshotname) {
		toSerialize["snapshotname"] = o.Snapshotname
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Xqn) {
		toSerialize["xqn"] = o.Xqn
	}
	if !IsNil(o.Wwn) {
		toSerialize["wwn"] = o.Wwn
	}
	return toSerialize, nil
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
