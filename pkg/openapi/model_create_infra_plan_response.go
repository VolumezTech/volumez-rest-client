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

// checks if the CreateInfraPlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInfraPlanResponse{}

// CreateInfraPlanResponse struct for CreateInfraPlanResponse
type CreateInfraPlanResponse struct {
	InstancesCount       int32    `json:"instancesCount"`
	InstanceType         string   `json:"instanceType"`
	ProviderInstanceType *string  `json:"providerInstanceType,omitempty"`
	TotalPrice           float64  `json:"totalPrice"`
	VolumezPrice         float64  `json:"volumezPrice"`
	MediaIOPSRead        int32    `json:"mediaIOPSRead"`
	MediaIOPSWrite       int32    `json:"mediaIOPSWrite"`
	MediaBandwidthWrite  int32    `json:"mediaBandwidthWrite"`
	MediaBandwidthRead   int32    `json:"mediaBandwidthRead"`
	RaidResiliency       string   `json:"raidResiliency"`
	Zones                []string `json:"zones"`
	SubnetIds            []string `json:"subnetIds"`
	OsType               OSType   `json:"osType"`
	Encryption           bool     `json:"encryption"`
	MediaSize            int32    `json:"mediaSize"`
	RaidColumns          int32    `json:"raidColumns"`
	Redundancy           int32    `json:"redundancy"`
}

type _CreateInfraPlanResponse CreateInfraPlanResponse

// NewCreateInfraPlanResponse instantiates a new CreateInfraPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInfraPlanResponse(instancesCount int32, instanceType string, totalPrice float64, volumezPrice float64, mediaIOPSRead int32, mediaIOPSWrite int32, mediaBandwidthWrite int32, mediaBandwidthRead int32, raidResiliency string, zones []string, subnetIds []string, osType OSType, encryption bool, mediaSize int32, raidColumns int32, redundancy int32) *CreateInfraPlanResponse {
	this := CreateInfraPlanResponse{}
	this.InstancesCount = instancesCount
	this.InstanceType = instanceType
	this.TotalPrice = totalPrice
	this.VolumezPrice = volumezPrice
	this.MediaIOPSRead = mediaIOPSRead
	this.MediaIOPSWrite = mediaIOPSWrite
	this.MediaBandwidthWrite = mediaBandwidthWrite
	this.MediaBandwidthRead = mediaBandwidthRead
	this.RaidResiliency = raidResiliency
	this.Zones = zones
	this.SubnetIds = subnetIds
	this.OsType = osType
	this.Encryption = encryption
	this.MediaSize = mediaSize
	this.RaidColumns = raidColumns
	this.Redundancy = redundancy
	return &this
}

// NewCreateInfraPlanResponseWithDefaults instantiates a new CreateInfraPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInfraPlanResponseWithDefaults() *CreateInfraPlanResponse {
	this := CreateInfraPlanResponse{}
	return &this
}

// GetInstancesCount returns the InstancesCount field value
func (o *CreateInfraPlanResponse) GetInstancesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InstancesCount
}

// GetInstancesCountOk returns a tuple with the InstancesCount field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetInstancesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstancesCount, true
}

// SetInstancesCount sets field value
func (o *CreateInfraPlanResponse) SetInstancesCount(v int32) {
	o.InstancesCount = v
}

// GetInstanceType returns the InstanceType field value
func (o *CreateInfraPlanResponse) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *CreateInfraPlanResponse) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetProviderInstanceType returns the ProviderInstanceType field value if set, zero value otherwise.
func (o *CreateInfraPlanResponse) GetProviderInstanceType() string {
	if o == nil || IsNil(o.ProviderInstanceType) {
		var ret string
		return ret
	}
	return *o.ProviderInstanceType
}

// GetProviderInstanceTypeOk returns a tuple with the ProviderInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetProviderInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderInstanceType) {
		return nil, false
	}
	return o.ProviderInstanceType, true
}

// HasProviderInstanceType returns a boolean if a field has been set.
func (o *CreateInfraPlanResponse) HasProviderInstanceType() bool {
	if o != nil && !IsNil(o.ProviderInstanceType) {
		return true
	}

	return false
}

// SetProviderInstanceType gets a reference to the given string and assigns it to the ProviderInstanceType field.
func (o *CreateInfraPlanResponse) SetProviderInstanceType(v string) {
	o.ProviderInstanceType = &v
}

// GetTotalPrice returns the TotalPrice field value
func (o *CreateInfraPlanResponse) GetTotalPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetTotalPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *CreateInfraPlanResponse) SetTotalPrice(v float64) {
	o.TotalPrice = v
}

// GetVolumezPrice returns the VolumezPrice field value
func (o *CreateInfraPlanResponse) GetVolumezPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.VolumezPrice
}

// GetVolumezPriceOk returns a tuple with the VolumezPrice field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetVolumezPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumezPrice, true
}

// SetVolumezPrice sets field value
func (o *CreateInfraPlanResponse) SetVolumezPrice(v float64) {
	o.VolumezPrice = v
}

// GetMediaIOPSRead returns the MediaIOPSRead field value
func (o *CreateInfraPlanResponse) GetMediaIOPSRead() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaIOPSRead
}

// GetMediaIOPSReadOk returns a tuple with the MediaIOPSRead field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetMediaIOPSReadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaIOPSRead, true
}

// SetMediaIOPSRead sets field value
func (o *CreateInfraPlanResponse) SetMediaIOPSRead(v int32) {
	o.MediaIOPSRead = v
}

// GetMediaIOPSWrite returns the MediaIOPSWrite field value
func (o *CreateInfraPlanResponse) GetMediaIOPSWrite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaIOPSWrite
}

// GetMediaIOPSWriteOk returns a tuple with the MediaIOPSWrite field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetMediaIOPSWriteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaIOPSWrite, true
}

// SetMediaIOPSWrite sets field value
func (o *CreateInfraPlanResponse) SetMediaIOPSWrite(v int32) {
	o.MediaIOPSWrite = v
}

// GetMediaBandwidthWrite returns the MediaBandwidthWrite field value
func (o *CreateInfraPlanResponse) GetMediaBandwidthWrite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaBandwidthWrite
}

// GetMediaBandwidthWriteOk returns a tuple with the MediaBandwidthWrite field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetMediaBandwidthWriteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaBandwidthWrite, true
}

// SetMediaBandwidthWrite sets field value
func (o *CreateInfraPlanResponse) SetMediaBandwidthWrite(v int32) {
	o.MediaBandwidthWrite = v
}

// GetMediaBandwidthRead returns the MediaBandwidthRead field value
func (o *CreateInfraPlanResponse) GetMediaBandwidthRead() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaBandwidthRead
}

// GetMediaBandwidthReadOk returns a tuple with the MediaBandwidthRead field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetMediaBandwidthReadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaBandwidthRead, true
}

// SetMediaBandwidthRead sets field value
func (o *CreateInfraPlanResponse) SetMediaBandwidthRead(v int32) {
	o.MediaBandwidthRead = v
}

// GetRaidResiliency returns the RaidResiliency field value
func (o *CreateInfraPlanResponse) GetRaidResiliency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RaidResiliency
}

// GetRaidResiliencyOk returns a tuple with the RaidResiliency field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetRaidResiliencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RaidResiliency, true
}

// SetRaidResiliency sets field value
func (o *CreateInfraPlanResponse) SetRaidResiliency(v string) {
	o.RaidResiliency = v
}

// GetZones returns the Zones field value
func (o *CreateInfraPlanResponse) GetZones() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetZonesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zones, true
}

// SetZones sets field value
func (o *CreateInfraPlanResponse) SetZones(v []string) {
	o.Zones = v
}

// GetSubnetIds returns the SubnetIds field value
func (o *CreateInfraPlanResponse) GetSubnetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubnetIds
}

// GetSubnetIdsOk returns a tuple with the SubnetIds field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetSubnetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetIds, true
}

// SetSubnetIds sets field value
func (o *CreateInfraPlanResponse) SetSubnetIds(v []string) {
	o.SubnetIds = v
}

// GetOsType returns the OsType field value
func (o *CreateInfraPlanResponse) GetOsType() OSType {
	if o == nil {
		var ret OSType
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetOsTypeOk() (*OSType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *CreateInfraPlanResponse) SetOsType(v OSType) {
	o.OsType = v
}

// GetEncryption returns the Encryption field value
func (o *CreateInfraPlanResponse) GetEncryption() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetEncryptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encryption, true
}

// SetEncryption sets field value
func (o *CreateInfraPlanResponse) SetEncryption(v bool) {
	o.Encryption = v
}

// GetMediaSize returns the MediaSize field value
func (o *CreateInfraPlanResponse) GetMediaSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MediaSize
}

// GetMediaSizeOk returns a tuple with the MediaSize field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetMediaSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaSize, true
}

// SetMediaSize sets field value
func (o *CreateInfraPlanResponse) SetMediaSize(v int32) {
	o.MediaSize = v
}

// GetRaidColumns returns the RaidColumns field value
func (o *CreateInfraPlanResponse) GetRaidColumns() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RaidColumns
}

// GetRaidColumnsOk returns a tuple with the RaidColumns field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetRaidColumnsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RaidColumns, true
}

// SetRaidColumns sets field value
func (o *CreateInfraPlanResponse) SetRaidColumns(v int32) {
	o.RaidColumns = v
}

// GetRedundancy returns the Redundancy field value
func (o *CreateInfraPlanResponse) GetRedundancy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value
// and a boolean to check if the value has been set.
func (o *CreateInfraPlanResponse) GetRedundancyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redundancy, true
}

// SetRedundancy sets field value
func (o *CreateInfraPlanResponse) SetRedundancy(v int32) {
	o.Redundancy = v
}

func (o CreateInfraPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInfraPlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instancesCount"] = o.InstancesCount
	toSerialize["instanceType"] = o.InstanceType
	if !IsNil(o.ProviderInstanceType) {
		toSerialize["providerInstanceType"] = o.ProviderInstanceType
	}
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["volumezPrice"] = o.VolumezPrice
	toSerialize["mediaIOPSRead"] = o.MediaIOPSRead
	toSerialize["mediaIOPSWrite"] = o.MediaIOPSWrite
	toSerialize["mediaBandwidthWrite"] = o.MediaBandwidthWrite
	toSerialize["mediaBandwidthRead"] = o.MediaBandwidthRead
	toSerialize["raidResiliency"] = o.RaidResiliency
	toSerialize["zones"] = o.Zones
	toSerialize["subnetIds"] = o.SubnetIds
	toSerialize["osType"] = o.OsType
	toSerialize["encryption"] = o.Encryption
	toSerialize["mediaSize"] = o.MediaSize
	toSerialize["raidColumns"] = o.RaidColumns
	toSerialize["redundancy"] = o.Redundancy
	return toSerialize, nil
}

func (o *CreateInfraPlanResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instancesCount",
		"instanceType",
		"totalPrice",
		"volumezPrice",
		"mediaIOPSRead",
		"mediaIOPSWrite",
		"mediaBandwidthWrite",
		"mediaBandwidthRead",
		"raidResiliency",
		"zones",
		"subnetIds",
		"osType",
		"encryption",
		"mediaSize",
		"raidColumns",
		"redundancy",
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

	varCreateInfraPlanResponse := _CreateInfraPlanResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInfraPlanResponse)

	if err != nil {
		return err
	}

	*o = CreateInfraPlanResponse(varCreateInfraPlanResponse)

	return err
}

type NullableCreateInfraPlanResponse struct {
	value *CreateInfraPlanResponse
	isSet bool
}

func (v NullableCreateInfraPlanResponse) Get() *CreateInfraPlanResponse {
	return v.value
}

func (v *NullableCreateInfraPlanResponse) Set(val *CreateInfraPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInfraPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInfraPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInfraPlanResponse(val *CreateInfraPlanResponse) *NullableCreateInfraPlanResponse {
	return &NullableCreateInfraPlanResponse{value: val, isSet: true}
}

func (v NullableCreateInfraPlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInfraPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
