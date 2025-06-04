# ProviderPricingInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**Iops** | **int32** |  | 
**Size** | **int32** |  | 
**Throughput** | Pointer to **int32** |  | [optional] 

## Methods

### NewProviderPricingInfoRequest

`func NewProviderPricingInfoRequest(cloudProvider CloudProviderType, iops int32, size int32, ) *ProviderPricingInfoRequest`

NewProviderPricingInfoRequest instantiates a new ProviderPricingInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderPricingInfoRequestWithDefaults

`func NewProviderPricingInfoRequestWithDefaults() *ProviderPricingInfoRequest`

NewProviderPricingInfoRequestWithDefaults instantiates a new ProviderPricingInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ProviderPricingInfoRequest) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ProviderPricingInfoRequest) GetCloudProviderOk() (*CloudProviderType, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ProviderPricingInfoRequest) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.


### GetIops

`func (o *ProviderPricingInfoRequest) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *ProviderPricingInfoRequest) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *ProviderPricingInfoRequest) SetIops(v int32)`

SetIops sets Iops field to given value.


### GetSize

`func (o *ProviderPricingInfoRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProviderPricingInfoRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProviderPricingInfoRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetThroughput

`func (o *ProviderPricingInfoRequest) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *ProviderPricingInfoRequest) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *ProviderPricingInfoRequest) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *ProviderPricingInfoRequest) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


