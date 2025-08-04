# GetVMMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpc** | Pointer to [**AvailableVPCItem**](AvailableVPCItem.md) |  | [optional] 

## Methods

### NewGetVMMetadataResponse

`func NewGetVMMetadataResponse() *GetVMMetadataResponse`

NewGetVMMetadataResponse instantiates a new GetVMMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMMetadataResponseWithDefaults

`func NewGetVMMetadataResponseWithDefaults() *GetVMMetadataResponse`

NewGetVMMetadataResponseWithDefaults instantiates a new GetVMMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpc

`func (o *GetVMMetadataResponse) GetVpc() AvailableVPCItem`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *GetVMMetadataResponse) GetVpcOk() (*AvailableVPCItem, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *GetVMMetadataResponse) SetVpc(v AvailableVPCItem)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *GetVMMetadataResponse) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


