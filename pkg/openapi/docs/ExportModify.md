# ExportModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedHosts** | Pointer to **[]string** |  | [optional] 
**Nodes** | Pointer to [**[]ExportCreateNodesInner**](ExportCreateNodesInner.md) |  | [optional] 

## Methods

### NewExportModify

`func NewExportModify() *ExportModify`

NewExportModify instantiates a new ExportModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportModifyWithDefaults

`func NewExportModifyWithDefaults() *ExportModify`

NewExportModifyWithDefaults instantiates a new ExportModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedHosts

`func (o *ExportModify) GetAllowedHosts() []string`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *ExportModify) GetAllowedHostsOk() (*[]string, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *ExportModify) SetAllowedHosts(v []string)`

SetAllowedHosts sets AllowedHosts field to given value.

### HasAllowedHosts

`func (o *ExportModify) HasAllowedHosts() bool`

HasAllowedHosts returns a boolean if a field has been set.

### GetNodes

`func (o *ExportModify) GetNodes() []ExportCreateNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ExportModify) GetNodesOk() (*[]ExportCreateNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ExportModify) SetNodes(v []ExportCreateNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ExportModify) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


