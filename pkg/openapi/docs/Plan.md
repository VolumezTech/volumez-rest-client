# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumegroup** | [**VolumeGroup**](VolumeGroup.md) |  | 
**Replicationvolumegroup** | Pointer to [**VolumeGroup**](VolumeGroup.md) |  | [optional] 

## Methods

### NewPlan

`func NewPlan(volumegroup VolumeGroup, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumegroup

`func (o *Plan) GetVolumegroup() VolumeGroup`

GetVolumegroup returns the Volumegroup field if non-nil, zero value otherwise.

### GetVolumegroupOk

`func (o *Plan) GetVolumegroupOk() (*VolumeGroup, bool)`

GetVolumegroupOk returns a tuple with the Volumegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumegroup

`func (o *Plan) SetVolumegroup(v VolumeGroup)`

SetVolumegroup sets Volumegroup field to given value.


### GetReplicationvolumegroup

`func (o *Plan) GetReplicationvolumegroup() VolumeGroup`

GetReplicationvolumegroup returns the Replicationvolumegroup field if non-nil, zero value otherwise.

### GetReplicationvolumegroupOk

`func (o *Plan) GetReplicationvolumegroupOk() (*VolumeGroup, bool)`

GetReplicationvolumegroupOk returns a tuple with the Replicationvolumegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationvolumegroup

`func (o *Plan) SetReplicationvolumegroup(v VolumeGroup)`

SetReplicationvolumegroup sets Replicationvolumegroup field to given value.

### HasReplicationvolumegroup

`func (o *Plan) HasReplicationvolumegroup() bool`

HasReplicationvolumegroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


