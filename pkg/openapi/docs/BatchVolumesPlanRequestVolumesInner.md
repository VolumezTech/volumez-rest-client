# BatchVolumesPlanRequestVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | policy name for the planned volume | [optional] 
**Size** | **int32** | volume size in GiB | 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchVolumesPlanRequestVolumesInner

`func NewBatchVolumesPlanRequestVolumesInner(size int32, ) *BatchVolumesPlanRequestVolumesInner`

NewBatchVolumesPlanRequestVolumesInner instantiates a new BatchVolumesPlanRequestVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchVolumesPlanRequestVolumesInnerWithDefaults

`func NewBatchVolumesPlanRequestVolumesInnerWithDefaults() *BatchVolumesPlanRequestVolumesInner`

NewBatchVolumesPlanRequestVolumesInnerWithDefaults instantiates a new BatchVolumesPlanRequestVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *BatchVolumesPlanRequestVolumesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BatchVolumesPlanRequestVolumesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BatchVolumesPlanRequestVolumesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *BatchVolumesPlanRequestVolumesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSize

`func (o *BatchVolumesPlanRequestVolumesInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BatchVolumesPlanRequestVolumesInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BatchVolumesPlanRequestVolumesInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetZone

`func (o *BatchVolumesPlanRequestVolumesInner) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *BatchVolumesPlanRequestVolumesInner) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *BatchVolumesPlanRequestVolumesInner) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *BatchVolumesPlanRequestVolumesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


