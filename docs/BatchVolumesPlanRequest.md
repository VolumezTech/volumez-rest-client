# BatchVolumesPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]BatchVolumesPlanRequestVolumesInner**](BatchVolumesPlanRequestVolumesInner.md) |  | [optional] 
**CapacityGroup** | Pointer to **string** | capacity groups to create the volume from (optional) | [optional] 
**DefaultZone** | Pointer to **string** | zone parameter for all volumes that dont have zone parameter | [optional] 
**DefaultPolicy** | Pointer to **string** | policy parameter for all volumes that dont have policy parameter defined in their input | [optional] 

## Methods

### NewBatchVolumesPlanRequest

`func NewBatchVolumesPlanRequest() *BatchVolumesPlanRequest`

NewBatchVolumesPlanRequest instantiates a new BatchVolumesPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchVolumesPlanRequestWithDefaults

`func NewBatchVolumesPlanRequestWithDefaults() *BatchVolumesPlanRequest`

NewBatchVolumesPlanRequestWithDefaults instantiates a new BatchVolumesPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *BatchVolumesPlanRequest) GetVolumes() []BatchVolumesPlanRequestVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *BatchVolumesPlanRequest) GetVolumesOk() (*[]BatchVolumesPlanRequestVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *BatchVolumesPlanRequest) SetVolumes(v []BatchVolumesPlanRequestVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *BatchVolumesPlanRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetCapacityGroup

`func (o *BatchVolumesPlanRequest) GetCapacityGroup() string`

GetCapacityGroup returns the CapacityGroup field if non-nil, zero value otherwise.

### GetCapacityGroupOk

`func (o *BatchVolumesPlanRequest) GetCapacityGroupOk() (*string, bool)`

GetCapacityGroupOk returns a tuple with the CapacityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGroup

`func (o *BatchVolumesPlanRequest) SetCapacityGroup(v string)`

SetCapacityGroup sets CapacityGroup field to given value.

### HasCapacityGroup

`func (o *BatchVolumesPlanRequest) HasCapacityGroup() bool`

HasCapacityGroup returns a boolean if a field has been set.

### GetDefaultZone

`func (o *BatchVolumesPlanRequest) GetDefaultZone() string`

GetDefaultZone returns the DefaultZone field if non-nil, zero value otherwise.

### GetDefaultZoneOk

`func (o *BatchVolumesPlanRequest) GetDefaultZoneOk() (*string, bool)`

GetDefaultZoneOk returns a tuple with the DefaultZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZone

`func (o *BatchVolumesPlanRequest) SetDefaultZone(v string)`

SetDefaultZone sets DefaultZone field to given value.

### HasDefaultZone

`func (o *BatchVolumesPlanRequest) HasDefaultZone() bool`

HasDefaultZone returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *BatchVolumesPlanRequest) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *BatchVolumesPlanRequest) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *BatchVolumesPlanRequest) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *BatchVolumesPlanRequest) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


