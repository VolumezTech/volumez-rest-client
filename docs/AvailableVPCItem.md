# AvailableVPCItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**SubnetId** | **string** |  | 
**SubnetName** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewAvailableVPCItem

`func NewAvailableVPCItem(id string, name string, subnetId string, ) *AvailableVPCItem`

NewAvailableVPCItem instantiates a new AvailableVPCItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableVPCItemWithDefaults

`func NewAvailableVPCItemWithDefaults() *AvailableVPCItem`

NewAvailableVPCItemWithDefaults instantiates a new AvailableVPCItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailableVPCItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailableVPCItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailableVPCItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AvailableVPCItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableVPCItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableVPCItem) SetName(v string)`

SetName sets Name field to given value.


### GetSubnetId

`func (o *AvailableVPCItem) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AvailableVPCItem) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AvailableVPCItem) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetName

`func (o *AvailableVPCItem) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *AvailableVPCItem) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *AvailableVPCItem) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *AvailableVPCItem) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### GetIsDefault

`func (o *AvailableVPCItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AvailableVPCItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AvailableVPCItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AvailableVPCItem) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AvailableVPCItem) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AvailableVPCItem) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AvailableVPCItem) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AvailableVPCItem) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


