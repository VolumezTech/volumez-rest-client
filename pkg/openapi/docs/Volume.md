# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Volumeid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**Contentvolume** | Pointer to **string** |  | [optional] 
**Contentsnapshot** | Pointer to **string** |  | [optional] 
**Size** | **int32** |  | 
**Maxsize** | Pointer to **int32** |  | [optional] [readonly] 
**Policy** | **string** |  | 
**Consistencygroup** | Pointer to **string** |  | [optional] 
**Node** | Pointer to **string** |  | [optional] [readonly] 
**Zone** | Pointer to **string** |  | [optional] 
**Zonereplica** | Pointer to **string** |  | [optional] 
**Volumegroupname** | Pointer to **string** |  | [optional] 
**Volumegroupid** | Pointer to **string** |  | [optional] [readonly] 
**Replicationnode** | Pointer to **string** |  | [optional] 
**Replicationvolumegroupname** | Pointer to **string** |  | [optional] 
**Replicationvolumegroupid** | Pointer to **string** |  | [optional] [readonly] 
**Volumerecoveryjob** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Capacitygroup** | Pointer to **string** |  | [optional] 
**Throttlingscheme** | Pointer to **string** |  | [optional] 
**Allowdatamovement** | Pointer to **bool** |  | [optional] [default to false]
**Flavor** | Pointer to **string** |  | [optional] 

## Methods

### NewVolume

`func NewVolume(name string, type_ string, size int32, policy string, ) *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.


### GetVolumeid

`func (o *Volume) GetVolumeid() string`

GetVolumeid returns the Volumeid field if non-nil, zero value otherwise.

### GetVolumeidOk

`func (o *Volume) GetVolumeidOk() (*string, bool)`

GetVolumeidOk returns a tuple with the Volumeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeid

`func (o *Volume) SetVolumeid(v string)`

SetVolumeid sets Volumeid field to given value.

### HasVolumeid

`func (o *Volume) HasVolumeid() bool`

HasVolumeid returns a boolean if a field has been set.

### GetType

`func (o *Volume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Volume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Volume) SetType(v string)`

SetType sets Type field to given value.


### GetContentvolume

`func (o *Volume) GetContentvolume() string`

GetContentvolume returns the Contentvolume field if non-nil, zero value otherwise.

### GetContentvolumeOk

`func (o *Volume) GetContentvolumeOk() (*string, bool)`

GetContentvolumeOk returns a tuple with the Contentvolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentvolume

`func (o *Volume) SetContentvolume(v string)`

SetContentvolume sets Contentvolume field to given value.

### HasContentvolume

`func (o *Volume) HasContentvolume() bool`

HasContentvolume returns a boolean if a field has been set.

### GetContentsnapshot

`func (o *Volume) GetContentsnapshot() string`

GetContentsnapshot returns the Contentsnapshot field if non-nil, zero value otherwise.

### GetContentsnapshotOk

`func (o *Volume) GetContentsnapshotOk() (*string, bool)`

GetContentsnapshotOk returns a tuple with the Contentsnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsnapshot

`func (o *Volume) SetContentsnapshot(v string)`

SetContentsnapshot sets Contentsnapshot field to given value.

### HasContentsnapshot

`func (o *Volume) HasContentsnapshot() bool`

HasContentsnapshot returns a boolean if a field has been set.

### GetSize

`func (o *Volume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMaxsize

`func (o *Volume) GetMaxsize() int32`

GetMaxsize returns the Maxsize field if non-nil, zero value otherwise.

### GetMaxsizeOk

`func (o *Volume) GetMaxsizeOk() (*int32, bool)`

GetMaxsizeOk returns a tuple with the Maxsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsize

`func (o *Volume) SetMaxsize(v int32)`

SetMaxsize sets Maxsize field to given value.

### HasMaxsize

`func (o *Volume) HasMaxsize() bool`

HasMaxsize returns a boolean if a field has been set.

### GetPolicy

`func (o *Volume) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Volume) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Volume) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetConsistencygroup

`func (o *Volume) GetConsistencygroup() string`

GetConsistencygroup returns the Consistencygroup field if non-nil, zero value otherwise.

### GetConsistencygroupOk

`func (o *Volume) GetConsistencygroupOk() (*string, bool)`

GetConsistencygroupOk returns a tuple with the Consistencygroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencygroup

`func (o *Volume) SetConsistencygroup(v string)`

SetConsistencygroup sets Consistencygroup field to given value.

### HasConsistencygroup

`func (o *Volume) HasConsistencygroup() bool`

HasConsistencygroup returns a boolean if a field has been set.

### GetNode

`func (o *Volume) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Volume) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Volume) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *Volume) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetZone

`func (o *Volume) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Volume) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Volume) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Volume) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonereplica

`func (o *Volume) GetZonereplica() string`

GetZonereplica returns the Zonereplica field if non-nil, zero value otherwise.

### GetZonereplicaOk

`func (o *Volume) GetZonereplicaOk() (*string, bool)`

GetZonereplicaOk returns a tuple with the Zonereplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonereplica

`func (o *Volume) SetZonereplica(v string)`

SetZonereplica sets Zonereplica field to given value.

### HasZonereplica

`func (o *Volume) HasZonereplica() bool`

HasZonereplica returns a boolean if a field has been set.

### GetVolumegroupname

`func (o *Volume) GetVolumegroupname() string`

GetVolumegroupname returns the Volumegroupname field if non-nil, zero value otherwise.

### GetVolumegroupnameOk

`func (o *Volume) GetVolumegroupnameOk() (*string, bool)`

GetVolumegroupnameOk returns a tuple with the Volumegroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumegroupname

`func (o *Volume) SetVolumegroupname(v string)`

SetVolumegroupname sets Volumegroupname field to given value.

### HasVolumegroupname

`func (o *Volume) HasVolumegroupname() bool`

HasVolumegroupname returns a boolean if a field has been set.

### GetVolumegroupid

`func (o *Volume) GetVolumegroupid() string`

GetVolumegroupid returns the Volumegroupid field if non-nil, zero value otherwise.

### GetVolumegroupidOk

`func (o *Volume) GetVolumegroupidOk() (*string, bool)`

GetVolumegroupidOk returns a tuple with the Volumegroupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumegroupid

`func (o *Volume) SetVolumegroupid(v string)`

SetVolumegroupid sets Volumegroupid field to given value.

### HasVolumegroupid

`func (o *Volume) HasVolumegroupid() bool`

HasVolumegroupid returns a boolean if a field has been set.

### GetReplicationnode

`func (o *Volume) GetReplicationnode() string`

GetReplicationnode returns the Replicationnode field if non-nil, zero value otherwise.

### GetReplicationnodeOk

`func (o *Volume) GetReplicationnodeOk() (*string, bool)`

GetReplicationnodeOk returns a tuple with the Replicationnode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationnode

`func (o *Volume) SetReplicationnode(v string)`

SetReplicationnode sets Replicationnode field to given value.

### HasReplicationnode

`func (o *Volume) HasReplicationnode() bool`

HasReplicationnode returns a boolean if a field has been set.

### GetReplicationvolumegroupname

`func (o *Volume) GetReplicationvolumegroupname() string`

GetReplicationvolumegroupname returns the Replicationvolumegroupname field if non-nil, zero value otherwise.

### GetReplicationvolumegroupnameOk

`func (o *Volume) GetReplicationvolumegroupnameOk() (*string, bool)`

GetReplicationvolumegroupnameOk returns a tuple with the Replicationvolumegroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationvolumegroupname

`func (o *Volume) SetReplicationvolumegroupname(v string)`

SetReplicationvolumegroupname sets Replicationvolumegroupname field to given value.

### HasReplicationvolumegroupname

`func (o *Volume) HasReplicationvolumegroupname() bool`

HasReplicationvolumegroupname returns a boolean if a field has been set.

### GetReplicationvolumegroupid

`func (o *Volume) GetReplicationvolumegroupid() string`

GetReplicationvolumegroupid returns the Replicationvolumegroupid field if non-nil, zero value otherwise.

### GetReplicationvolumegroupidOk

`func (o *Volume) GetReplicationvolumegroupidOk() (*string, bool)`

GetReplicationvolumegroupidOk returns a tuple with the Replicationvolumegroupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationvolumegroupid

`func (o *Volume) SetReplicationvolumegroupid(v string)`

SetReplicationvolumegroupid sets Replicationvolumegroupid field to given value.

### HasReplicationvolumegroupid

`func (o *Volume) HasReplicationvolumegroupid() bool`

HasReplicationvolumegroupid returns a boolean if a field has been set.

### GetVolumerecoveryjob

`func (o *Volume) GetVolumerecoveryjob() string`

GetVolumerecoveryjob returns the Volumerecoveryjob field if non-nil, zero value otherwise.

### GetVolumerecoveryjobOk

`func (o *Volume) GetVolumerecoveryjobOk() (*string, bool)`

GetVolumerecoveryjobOk returns a tuple with the Volumerecoveryjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumerecoveryjob

`func (o *Volume) SetVolumerecoveryjob(v string)`

SetVolumerecoveryjob sets Volumerecoveryjob field to given value.

### HasVolumerecoveryjob

`func (o *Volume) HasVolumerecoveryjob() bool`

HasVolumerecoveryjob returns a boolean if a field has been set.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Volume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Volume) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Volume) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Volume) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Volume) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCapacitygroup

`func (o *Volume) GetCapacitygroup() string`

GetCapacitygroup returns the Capacitygroup field if non-nil, zero value otherwise.

### GetCapacitygroupOk

`func (o *Volume) GetCapacitygroupOk() (*string, bool)`

GetCapacitygroupOk returns a tuple with the Capacitygroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitygroup

`func (o *Volume) SetCapacitygroup(v string)`

SetCapacitygroup sets Capacitygroup field to given value.

### HasCapacitygroup

`func (o *Volume) HasCapacitygroup() bool`

HasCapacitygroup returns a boolean if a field has been set.

### GetThrottlingscheme

`func (o *Volume) GetThrottlingscheme() string`

GetThrottlingscheme returns the Throttlingscheme field if non-nil, zero value otherwise.

### GetThrottlingschemeOk

`func (o *Volume) GetThrottlingschemeOk() (*string, bool)`

GetThrottlingschemeOk returns a tuple with the Throttlingscheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingscheme

`func (o *Volume) SetThrottlingscheme(v string)`

SetThrottlingscheme sets Throttlingscheme field to given value.

### HasThrottlingscheme

`func (o *Volume) HasThrottlingscheme() bool`

HasThrottlingscheme returns a boolean if a field has been set.

### GetAllowdatamovement

`func (o *Volume) GetAllowdatamovement() bool`

GetAllowdatamovement returns the Allowdatamovement field if non-nil, zero value otherwise.

### GetAllowdatamovementOk

`func (o *Volume) GetAllowdatamovementOk() (*bool, bool)`

GetAllowdatamovementOk returns a tuple with the Allowdatamovement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowdatamovement

`func (o *Volume) SetAllowdatamovement(v bool)`

SetAllowdatamovement sets Allowdatamovement field to given value.

### HasAllowdatamovement

`func (o *Volume) HasAllowdatamovement() bool`

HasAllowdatamovement returns a boolean if a field has been set.

### GetFlavor

`func (o *Volume) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Volume) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Volume) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *Volume) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


