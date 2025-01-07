# GetVmRegionZonesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **map[string][]string** |  | [optional] 
**Azure** | Pointer to [**map[string]AzureVMRegion**](AzureVMRegion.md) |  | [optional] 

## Methods

### NewGetVmRegionZonesResponse

`func NewGetVmRegionZonesResponse() *GetVmRegionZonesResponse`

NewGetVmRegionZonesResponse instantiates a new GetVmRegionZonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVmRegionZonesResponseWithDefaults

`func NewGetVmRegionZonesResponseWithDefaults() *GetVmRegionZonesResponse`

NewGetVmRegionZonesResponseWithDefaults instantiates a new GetVmRegionZonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *GetVmRegionZonesResponse) GetAws() map[string][]string`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *GetVmRegionZonesResponse) GetAwsOk() (*map[string][]string, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *GetVmRegionZonesResponse) SetAws(v map[string][]string)`

SetAws sets Aws field to given value.

### HasAws

`func (o *GetVmRegionZonesResponse) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *GetVmRegionZonesResponse) GetAzure() map[string]AzureVMRegion`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *GetVmRegionZonesResponse) GetAzureOk() (*map[string]AzureVMRegion, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *GetVmRegionZonesResponse) SetAzure(v map[string]AzureVMRegion)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *GetVmRegionZonesResponse) HasAzure() bool`

HasAzure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


