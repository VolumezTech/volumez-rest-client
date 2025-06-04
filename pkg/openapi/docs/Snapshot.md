# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumename** | Pointer to **string** |  | [optional] 
**Volumeid** | Pointer to **string** |  | [optional] [readonly] 
**Volumesize** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Snapshotid** | Pointer to **string** |  | [optional] [readonly] 
**Targetsecret** | Pointer to **string** |  | [optional] 
**Consistency** | **string** |  | 
**Time** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **bool** |  | [optional] 
**Consistencygroup** | Pointer to **bool** |  | [optional] 
**Consistencygroupname** | Pointer to **string** |  | [optional] 
**Used** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Numberofattachments** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSnapshot

`func NewSnapshot(name string, consistency string, ) *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumename

`func (o *Snapshot) GetVolumename() string`

GetVolumename returns the Volumename field if non-nil, zero value otherwise.

### GetVolumenameOk

`func (o *Snapshot) GetVolumenameOk() (*string, bool)`

GetVolumenameOk returns a tuple with the Volumename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumename

`func (o *Snapshot) SetVolumename(v string)`

SetVolumename sets Volumename field to given value.

### HasVolumename

`func (o *Snapshot) HasVolumename() bool`

HasVolumename returns a boolean if a field has been set.

### GetVolumeid

`func (o *Snapshot) GetVolumeid() string`

GetVolumeid returns the Volumeid field if non-nil, zero value otherwise.

### GetVolumeidOk

`func (o *Snapshot) GetVolumeidOk() (*string, bool)`

GetVolumeidOk returns a tuple with the Volumeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeid

`func (o *Snapshot) SetVolumeid(v string)`

SetVolumeid sets Volumeid field to given value.

### HasVolumeid

`func (o *Snapshot) HasVolumeid() bool`

HasVolumeid returns a boolean if a field has been set.

### GetVolumesize

`func (o *Snapshot) GetVolumesize() int32`

GetVolumesize returns the Volumesize field if non-nil, zero value otherwise.

### GetVolumesizeOk

`func (o *Snapshot) GetVolumesizeOk() (*int32, bool)`

GetVolumesizeOk returns a tuple with the Volumesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesize

`func (o *Snapshot) SetVolumesize(v int32)`

SetVolumesize sets Volumesize field to given value.

### HasVolumesize

`func (o *Snapshot) HasVolumesize() bool`

HasVolumesize returns a boolean if a field has been set.

### GetName

`func (o *Snapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snapshot) SetName(v string)`

SetName sets Name field to given value.


### GetSnapshotid

`func (o *Snapshot) GetSnapshotid() string`

GetSnapshotid returns the Snapshotid field if non-nil, zero value otherwise.

### GetSnapshotidOk

`func (o *Snapshot) GetSnapshotidOk() (*string, bool)`

GetSnapshotidOk returns a tuple with the Snapshotid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotid

`func (o *Snapshot) SetSnapshotid(v string)`

SetSnapshotid sets Snapshotid field to given value.

### HasSnapshotid

`func (o *Snapshot) HasSnapshotid() bool`

HasSnapshotid returns a boolean if a field has been set.

### GetTargetsecret

`func (o *Snapshot) GetTargetsecret() string`

GetTargetsecret returns the Targetsecret field if non-nil, zero value otherwise.

### GetTargetsecretOk

`func (o *Snapshot) GetTargetsecretOk() (*string, bool)`

GetTargetsecretOk returns a tuple with the Targetsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsecret

`func (o *Snapshot) SetTargetsecret(v string)`

SetTargetsecret sets Targetsecret field to given value.

### HasTargetsecret

`func (o *Snapshot) HasTargetsecret() bool`

HasTargetsecret returns a boolean if a field has been set.

### GetConsistency

`func (o *Snapshot) GetConsistency() string`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *Snapshot) GetConsistencyOk() (*string, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *Snapshot) SetConsistency(v string)`

SetConsistency sets Consistency field to given value.


### GetTime

`func (o *Snapshot) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Snapshot) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Snapshot) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *Snapshot) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPolicy

`func (o *Snapshot) GetPolicy() bool`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Snapshot) GetPolicyOk() (*bool, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Snapshot) SetPolicy(v bool)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Snapshot) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetConsistencygroup

`func (o *Snapshot) GetConsistencygroup() bool`

GetConsistencygroup returns the Consistencygroup field if non-nil, zero value otherwise.

### GetConsistencygroupOk

`func (o *Snapshot) GetConsistencygroupOk() (*bool, bool)`

GetConsistencygroupOk returns a tuple with the Consistencygroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencygroup

`func (o *Snapshot) SetConsistencygroup(v bool)`

SetConsistencygroup sets Consistencygroup field to given value.

### HasConsistencygroup

`func (o *Snapshot) HasConsistencygroup() bool`

HasConsistencygroup returns a boolean if a field has been set.

### GetConsistencygroupname

`func (o *Snapshot) GetConsistencygroupname() string`

GetConsistencygroupname returns the Consistencygroupname field if non-nil, zero value otherwise.

### GetConsistencygroupnameOk

`func (o *Snapshot) GetConsistencygroupnameOk() (*string, bool)`

GetConsistencygroupnameOk returns a tuple with the Consistencygroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencygroupname

`func (o *Snapshot) SetConsistencygroupname(v string)`

SetConsistencygroupname sets Consistencygroupname field to given value.

### HasConsistencygroupname

`func (o *Snapshot) HasConsistencygroupname() bool`

HasConsistencygroupname returns a boolean if a field has been set.

### GetUsed

`func (o *Snapshot) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Snapshot) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Snapshot) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Snapshot) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetState

`func (o *Snapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Snapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Snapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Snapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Snapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Snapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Snapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Snapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Snapshot) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Snapshot) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Snapshot) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Snapshot) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetNumberofattachments

`func (o *Snapshot) GetNumberofattachments() int32`

GetNumberofattachments returns the Numberofattachments field if non-nil, zero value otherwise.

### GetNumberofattachmentsOk

`func (o *Snapshot) GetNumberofattachmentsOk() (*int32, bool)`

GetNumberofattachmentsOk returns a tuple with the Numberofattachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberofattachments

`func (o *Snapshot) SetNumberofattachments(v int32)`

SetNumberofattachments sets Numberofattachments field to given value.

### HasNumberofattachments

`func (o *Snapshot) HasNumberofattachments() bool`

HasNumberofattachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


