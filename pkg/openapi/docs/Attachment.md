# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumename** | Pointer to **string** |  | [optional] 
**Volumeid** | Pointer to **string** |  | [optional] [readonly] 
**Snapshotname** | Pointer to **string** |  | [optional] 
**Snapshotid** | Pointer to **string** |  | [optional] [readonly] 
**Node** | **string** |  | 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**AllocatedResources** | Pointer to **int32** |  | [optional] 

## Methods

### NewAttachment

`func NewAttachment(node string, ) *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumename

`func (o *Attachment) GetVolumename() string`

GetVolumename returns the Volumename field if non-nil, zero value otherwise.

### GetVolumenameOk

`func (o *Attachment) GetVolumenameOk() (*string, bool)`

GetVolumenameOk returns a tuple with the Volumename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumename

`func (o *Attachment) SetVolumename(v string)`

SetVolumename sets Volumename field to given value.

### HasVolumename

`func (o *Attachment) HasVolumename() bool`

HasVolumename returns a boolean if a field has been set.

### GetVolumeid

`func (o *Attachment) GetVolumeid() string`

GetVolumeid returns the Volumeid field if non-nil, zero value otherwise.

### GetVolumeidOk

`func (o *Attachment) GetVolumeidOk() (*string, bool)`

GetVolumeidOk returns a tuple with the Volumeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeid

`func (o *Attachment) SetVolumeid(v string)`

SetVolumeid sets Volumeid field to given value.

### HasVolumeid

`func (o *Attachment) HasVolumeid() bool`

HasVolumeid returns a boolean if a field has been set.

### GetSnapshotname

`func (o *Attachment) GetSnapshotname() string`

GetSnapshotname returns the Snapshotname field if non-nil, zero value otherwise.

### GetSnapshotnameOk

`func (o *Attachment) GetSnapshotnameOk() (*string, bool)`

GetSnapshotnameOk returns a tuple with the Snapshotname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotname

`func (o *Attachment) SetSnapshotname(v string)`

SetSnapshotname sets Snapshotname field to given value.

### HasSnapshotname

`func (o *Attachment) HasSnapshotname() bool`

HasSnapshotname returns a boolean if a field has been set.

### GetSnapshotid

`func (o *Attachment) GetSnapshotid() string`

GetSnapshotid returns the Snapshotid field if non-nil, zero value otherwise.

### GetSnapshotidOk

`func (o *Attachment) GetSnapshotidOk() (*string, bool)`

GetSnapshotidOk returns a tuple with the Snapshotid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotid

`func (o *Attachment) SetSnapshotid(v string)`

SetSnapshotid sets Snapshotid field to given value.

### HasSnapshotid

`func (o *Attachment) HasSnapshotid() bool`

HasSnapshotid returns a boolean if a field has been set.

### GetNode

`func (o *Attachment) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Attachment) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Attachment) SetNode(v string)`

SetNode sets Node field to given value.


### GetState

`func (o *Attachment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Attachment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Attachment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Attachment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Attachment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attachment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attachment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attachment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Attachment) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Attachment) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Attachment) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Attachment) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetMountpoint

`func (o *Attachment) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *Attachment) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *Attachment) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *Attachment) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetReadonly

`func (o *Attachment) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *Attachment) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *Attachment) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *Attachment) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetAllocatedResources

`func (o *Attachment) GetAllocatedResources() int32`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *Attachment) GetAllocatedResourcesOk() (*int32, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *Attachment) SetAllocatedResources(v int32)`

SetAllocatedResources sets AllocatedResources field to given value.

### HasAllocatedResources

`func (o *Attachment) HasAllocatedResources() bool`

HasAllocatedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


