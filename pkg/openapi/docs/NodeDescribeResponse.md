# NodeDescribeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumesMap** | Pointer to [**map[string]Volume**](Volume.md) | Map of volume IDs to volume objects | [optional] 

## Methods

### NewNodeDescribeResponse

`func NewNodeDescribeResponse() *NodeDescribeResponse`

NewNodeDescribeResponse instantiates a new NodeDescribeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDescribeResponseWithDefaults

`func NewNodeDescribeResponseWithDefaults() *NodeDescribeResponse`

NewNodeDescribeResponseWithDefaults instantiates a new NodeDescribeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumesMap

`func (o *NodeDescribeResponse) GetVolumesMap() map[string]Volume`

GetVolumesMap returns the VolumesMap field if non-nil, zero value otherwise.

### GetVolumesMapOk

`func (o *NodeDescribeResponse) GetVolumesMapOk() (*map[string]Volume, bool)`

GetVolumesMapOk returns a tuple with the VolumesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesMap

`func (o *NodeDescribeResponse) SetVolumesMap(v map[string]Volume)`

SetVolumesMap sets VolumesMap field to given value.

### HasVolumesMap

`func (o *NodeDescribeResponse) HasVolumesMap() bool`

HasVolumesMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


