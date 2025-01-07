# RepairCmds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmds** | **[]string** |  | 
**Checksum** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewRepairCmds

`func NewRepairCmds(cmds []string, checksum string, ) *RepairCmds`

NewRepairCmds instantiates a new RepairCmds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepairCmdsWithDefaults

`func NewRepairCmdsWithDefaults() *RepairCmds`

NewRepairCmdsWithDefaults instantiates a new RepairCmds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmds

`func (o *RepairCmds) GetCmds() []string`

GetCmds returns the Cmds field if non-nil, zero value otherwise.

### GetCmdsOk

`func (o *RepairCmds) GetCmdsOk() (*[]string, bool)`

GetCmdsOk returns a tuple with the Cmds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmds

`func (o *RepairCmds) SetCmds(v []string)`

SetCmds sets Cmds field to given value.


### GetChecksum

`func (o *RepairCmds) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *RepairCmds) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *RepairCmds) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetMessage

`func (o *RepairCmds) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RepairCmds) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RepairCmds) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RepairCmds) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


