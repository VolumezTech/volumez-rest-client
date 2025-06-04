# AzureVMRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionDisplayName** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAzureVMRegion

`func NewAzureVMRegion() *AzureVMRegion`

NewAzureVMRegion instantiates a new AzureVMRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVMRegionWithDefaults

`func NewAzureVMRegionWithDefaults() *AzureVMRegion`

NewAzureVMRegionWithDefaults instantiates a new AzureVMRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionDisplayName

`func (o *AzureVMRegion) GetRegionDisplayName() string`

GetRegionDisplayName returns the RegionDisplayName field if non-nil, zero value otherwise.

### GetRegionDisplayNameOk

`func (o *AzureVMRegion) GetRegionDisplayNameOk() (*string, bool)`

GetRegionDisplayNameOk returns a tuple with the RegionDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionDisplayName

`func (o *AzureVMRegion) SetRegionDisplayName(v string)`

SetRegionDisplayName sets RegionDisplayName field to given value.

### HasRegionDisplayName

`func (o *AzureVMRegion) HasRegionDisplayName() bool`

HasRegionDisplayName returns a boolean if a field has been set.

### GetZones

`func (o *AzureVMRegion) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AzureVMRegion) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AzureVMRegion) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AzureVMRegion) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


