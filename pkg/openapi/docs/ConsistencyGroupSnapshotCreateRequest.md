# ConsistencyGroupSnapshotCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Consistency** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConsistencyGroupSnapshotCreateRequest

`func NewConsistencyGroupSnapshotCreateRequest() *ConsistencyGroupSnapshotCreateRequest`

NewConsistencyGroupSnapshotCreateRequest instantiates a new ConsistencyGroupSnapshotCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsistencyGroupSnapshotCreateRequestWithDefaults

`func NewConsistencyGroupSnapshotCreateRequestWithDefaults() *ConsistencyGroupSnapshotCreateRequest`

NewConsistencyGroupSnapshotCreateRequestWithDefaults instantiates a new ConsistencyGroupSnapshotCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsistencyGroupSnapshotCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsistencyGroupSnapshotCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsistencyGroupSnapshotCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsistencyGroupSnapshotCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConsistency

`func (o *ConsistencyGroupSnapshotCreateRequest) GetConsistency() string`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *ConsistencyGroupSnapshotCreateRequest) GetConsistencyOk() (*string, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *ConsistencyGroupSnapshotCreateRequest) SetConsistency(v string)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *ConsistencyGroupSnapshotCreateRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.

### GetGroupName

`func (o *ConsistencyGroupSnapshotCreateRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ConsistencyGroupSnapshotCreateRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ConsistencyGroupSnapshotCreateRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ConsistencyGroupSnapshotCreateRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetVolumes

`func (o *ConsistencyGroupSnapshotCreateRequest) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ConsistencyGroupSnapshotCreateRequest) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ConsistencyGroupSnapshotCreateRequest) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ConsistencyGroupSnapshotCreateRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


