# CreateInfraPlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstancesCount** | **int32** |  | 
**InstanceType** | **string** |  | 
**ProviderInstanceType** | Pointer to **string** |  | [optional] 
**TotalPrice** | **float64** |  | 
**VolumezPrice** | **float64** |  | 
**MediaIOPSRead** | **int32** |  | 
**MediaIOPSWrite** | **int32** |  | 
**MediaBandwidthWrite** | **int32** |  | 
**MediaBandwidthRead** | **int32** |  | 
**RaidResiliency** | **string** |  | 
**Zones** | **[]string** |  | 
**SubnetIds** | **[]string** |  | 
**OsType** | [**OSType**](OSType.md) |  | 
**Encryption** | **bool** |  | 
**MediaSize** | **int32** |  | 
**RaidColumns** | **int32** |  | 
**Redundancy** | **int32** |  | 

## Methods

### NewCreateInfraPlanResponse

`func NewCreateInfraPlanResponse(instancesCount int32, instanceType string, totalPrice float64, volumezPrice float64, mediaIOPSRead int32, mediaIOPSWrite int32, mediaBandwidthWrite int32, mediaBandwidthRead int32, raidResiliency string, zones []string, subnetIds []string, osType OSType, encryption bool, mediaSize int32, raidColumns int32, redundancy int32, ) *CreateInfraPlanResponse`

NewCreateInfraPlanResponse instantiates a new CreateInfraPlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraPlanResponseWithDefaults

`func NewCreateInfraPlanResponseWithDefaults() *CreateInfraPlanResponse`

NewCreateInfraPlanResponseWithDefaults instantiates a new CreateInfraPlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstancesCount

`func (o *CreateInfraPlanResponse) GetInstancesCount() int32`

GetInstancesCount returns the InstancesCount field if non-nil, zero value otherwise.

### GetInstancesCountOk

`func (o *CreateInfraPlanResponse) GetInstancesCountOk() (*int32, bool)`

GetInstancesCountOk returns a tuple with the InstancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCount

`func (o *CreateInfraPlanResponse) SetInstancesCount(v int32)`

SetInstancesCount sets InstancesCount field to given value.


### GetInstanceType

`func (o *CreateInfraPlanResponse) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CreateInfraPlanResponse) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CreateInfraPlanResponse) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetProviderInstanceType

`func (o *CreateInfraPlanResponse) GetProviderInstanceType() string`

GetProviderInstanceType returns the ProviderInstanceType field if non-nil, zero value otherwise.

### GetProviderInstanceTypeOk

`func (o *CreateInfraPlanResponse) GetProviderInstanceTypeOk() (*string, bool)`

GetProviderInstanceTypeOk returns a tuple with the ProviderInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInstanceType

`func (o *CreateInfraPlanResponse) SetProviderInstanceType(v string)`

SetProviderInstanceType sets ProviderInstanceType field to given value.

### HasProviderInstanceType

`func (o *CreateInfraPlanResponse) HasProviderInstanceType() bool`

HasProviderInstanceType returns a boolean if a field has been set.

### GetTotalPrice

`func (o *CreateInfraPlanResponse) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *CreateInfraPlanResponse) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *CreateInfraPlanResponse) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.


### GetVolumezPrice

`func (o *CreateInfraPlanResponse) GetVolumezPrice() float64`

GetVolumezPrice returns the VolumezPrice field if non-nil, zero value otherwise.

### GetVolumezPriceOk

`func (o *CreateInfraPlanResponse) GetVolumezPriceOk() (*float64, bool)`

GetVolumezPriceOk returns a tuple with the VolumezPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumezPrice

`func (o *CreateInfraPlanResponse) SetVolumezPrice(v float64)`

SetVolumezPrice sets VolumezPrice field to given value.


### GetMediaIOPSRead

`func (o *CreateInfraPlanResponse) GetMediaIOPSRead() int32`

GetMediaIOPSRead returns the MediaIOPSRead field if non-nil, zero value otherwise.

### GetMediaIOPSReadOk

`func (o *CreateInfraPlanResponse) GetMediaIOPSReadOk() (*int32, bool)`

GetMediaIOPSReadOk returns a tuple with the MediaIOPSRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaIOPSRead

`func (o *CreateInfraPlanResponse) SetMediaIOPSRead(v int32)`

SetMediaIOPSRead sets MediaIOPSRead field to given value.


### GetMediaIOPSWrite

`func (o *CreateInfraPlanResponse) GetMediaIOPSWrite() int32`

GetMediaIOPSWrite returns the MediaIOPSWrite field if non-nil, zero value otherwise.

### GetMediaIOPSWriteOk

`func (o *CreateInfraPlanResponse) GetMediaIOPSWriteOk() (*int32, bool)`

GetMediaIOPSWriteOk returns a tuple with the MediaIOPSWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaIOPSWrite

`func (o *CreateInfraPlanResponse) SetMediaIOPSWrite(v int32)`

SetMediaIOPSWrite sets MediaIOPSWrite field to given value.


### GetMediaBandwidthWrite

`func (o *CreateInfraPlanResponse) GetMediaBandwidthWrite() int32`

GetMediaBandwidthWrite returns the MediaBandwidthWrite field if non-nil, zero value otherwise.

### GetMediaBandwidthWriteOk

`func (o *CreateInfraPlanResponse) GetMediaBandwidthWriteOk() (*int32, bool)`

GetMediaBandwidthWriteOk returns a tuple with the MediaBandwidthWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBandwidthWrite

`func (o *CreateInfraPlanResponse) SetMediaBandwidthWrite(v int32)`

SetMediaBandwidthWrite sets MediaBandwidthWrite field to given value.


### GetMediaBandwidthRead

`func (o *CreateInfraPlanResponse) GetMediaBandwidthRead() int32`

GetMediaBandwidthRead returns the MediaBandwidthRead field if non-nil, zero value otherwise.

### GetMediaBandwidthReadOk

`func (o *CreateInfraPlanResponse) GetMediaBandwidthReadOk() (*int32, bool)`

GetMediaBandwidthReadOk returns a tuple with the MediaBandwidthRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaBandwidthRead

`func (o *CreateInfraPlanResponse) SetMediaBandwidthRead(v int32)`

SetMediaBandwidthRead sets MediaBandwidthRead field to given value.


### GetRaidResiliency

`func (o *CreateInfraPlanResponse) GetRaidResiliency() string`

GetRaidResiliency returns the RaidResiliency field if non-nil, zero value otherwise.

### GetRaidResiliencyOk

`func (o *CreateInfraPlanResponse) GetRaidResiliencyOk() (*string, bool)`

GetRaidResiliencyOk returns a tuple with the RaidResiliency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidResiliency

`func (o *CreateInfraPlanResponse) SetRaidResiliency(v string)`

SetRaidResiliency sets RaidResiliency field to given value.


### GetZones

`func (o *CreateInfraPlanResponse) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CreateInfraPlanResponse) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CreateInfraPlanResponse) SetZones(v []string)`

SetZones sets Zones field to given value.


### GetSubnetIds

`func (o *CreateInfraPlanResponse) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *CreateInfraPlanResponse) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *CreateInfraPlanResponse) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.


### GetOsType

`func (o *CreateInfraPlanResponse) GetOsType() OSType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateInfraPlanResponse) GetOsTypeOk() (*OSType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateInfraPlanResponse) SetOsType(v OSType)`

SetOsType sets OsType field to given value.


### GetEncryption

`func (o *CreateInfraPlanResponse) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CreateInfraPlanResponse) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CreateInfraPlanResponse) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.


### GetMediaSize

`func (o *CreateInfraPlanResponse) GetMediaSize() int32`

GetMediaSize returns the MediaSize field if non-nil, zero value otherwise.

### GetMediaSizeOk

`func (o *CreateInfraPlanResponse) GetMediaSizeOk() (*int32, bool)`

GetMediaSizeOk returns a tuple with the MediaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSize

`func (o *CreateInfraPlanResponse) SetMediaSize(v int32)`

SetMediaSize sets MediaSize field to given value.


### GetRaidColumns

`func (o *CreateInfraPlanResponse) GetRaidColumns() int32`

GetRaidColumns returns the RaidColumns field if non-nil, zero value otherwise.

### GetRaidColumnsOk

`func (o *CreateInfraPlanResponse) GetRaidColumnsOk() (*int32, bool)`

GetRaidColumnsOk returns a tuple with the RaidColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidColumns

`func (o *CreateInfraPlanResponse) SetRaidColumns(v int32)`

SetRaidColumns sets RaidColumns field to given value.


### GetRedundancy

`func (o *CreateInfraPlanResponse) GetRedundancy() int32`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateInfraPlanResponse) GetRedundancyOk() (*int32, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateInfraPlanResponse) SetRedundancy(v int32)`

SetRedundancy sets Redundancy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


