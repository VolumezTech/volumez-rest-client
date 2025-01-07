# MediaAssign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityGroup** | Pointer to **string** |  | [optional] 
**Reprofile** | Pointer to **bool** | Profile media even if it was already profiled | [optional] 
**Blocksize** | Pointer to **int32** | LBA size for media format | [optional] 

## Methods

### NewMediaAssign

`func NewMediaAssign() *MediaAssign`

NewMediaAssign instantiates a new MediaAssign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAssignWithDefaults

`func NewMediaAssignWithDefaults() *MediaAssign`

NewMediaAssignWithDefaults instantiates a new MediaAssign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityGroup

`func (o *MediaAssign) GetCapacityGroup() string`

GetCapacityGroup returns the CapacityGroup field if non-nil, zero value otherwise.

### GetCapacityGroupOk

`func (o *MediaAssign) GetCapacityGroupOk() (*string, bool)`

GetCapacityGroupOk returns a tuple with the CapacityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGroup

`func (o *MediaAssign) SetCapacityGroup(v string)`

SetCapacityGroup sets CapacityGroup field to given value.

### HasCapacityGroup

`func (o *MediaAssign) HasCapacityGroup() bool`

HasCapacityGroup returns a boolean if a field has been set.

### GetReprofile

`func (o *MediaAssign) GetReprofile() bool`

GetReprofile returns the Reprofile field if non-nil, zero value otherwise.

### GetReprofileOk

`func (o *MediaAssign) GetReprofileOk() (*bool, bool)`

GetReprofileOk returns a tuple with the Reprofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReprofile

`func (o *MediaAssign) SetReprofile(v bool)`

SetReprofile sets Reprofile field to given value.

### HasReprofile

`func (o *MediaAssign) HasReprofile() bool`

HasReprofile returns a boolean if a field has been set.

### GetBlocksize

`func (o *MediaAssign) GetBlocksize() int32`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *MediaAssign) GetBlocksizeOk() (*int32, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *MediaAssign) SetBlocksize(v int32)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *MediaAssign) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


