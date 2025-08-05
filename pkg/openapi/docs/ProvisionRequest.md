# ProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**TenantHostToken** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**MachineInfo** | Pointer to [**MachineInfo**](MachineInfo.md) |  | [optional] 
**BootTime** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisionRequest

`func NewProvisionRequest() *ProvisionRequest`

NewProvisionRequest instantiates a new ProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionRequestWithDefaults

`func NewProvisionRequestWithDefaults() *ProvisionRequest`

NewProvisionRequestWithDefaults instantiates a new ProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ProvisionRequest) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisionRequest) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisionRequest) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisionRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOs

`func (o *ProvisionRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ProvisionRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ProvisionRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ProvisionRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetTenantHostToken

`func (o *ProvisionRequest) GetTenantHostToken() string`

GetTenantHostToken returns the TenantHostToken field if non-nil, zero value otherwise.

### GetTenantHostTokenOk

`func (o *ProvisionRequest) GetTenantHostTokenOk() (*string, bool)`

GetTenantHostTokenOk returns a tuple with the TenantHostToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantHostToken

`func (o *ProvisionRequest) SetTenantHostToken(v string)`

SetTenantHostToken sets TenantHostToken field to given value.

### HasTenantHostToken

`func (o *ProvisionRequest) HasTenantHostToken() bool`

HasTenantHostToken returns a boolean if a field has been set.

### GetHostname

`func (o *ProvisionRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ProvisionRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ProvisionRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ProvisionRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ProvisionRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ProvisionRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ProvisionRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ProvisionRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMachineInfo

`func (o *ProvisionRequest) GetMachineInfo() MachineInfo`

GetMachineInfo returns the MachineInfo field if non-nil, zero value otherwise.

### GetMachineInfoOk

`func (o *ProvisionRequest) GetMachineInfoOk() (*MachineInfo, bool)`

GetMachineInfoOk returns a tuple with the MachineInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineInfo

`func (o *ProvisionRequest) SetMachineInfo(v MachineInfo)`

SetMachineInfo sets MachineInfo field to given value.

### HasMachineInfo

`func (o *ProvisionRequest) HasMachineInfo() bool`

HasMachineInfo returns a boolean if a field has been set.

### GetBootTime

`func (o *ProvisionRequest) GetBootTime() string`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *ProvisionRequest) GetBootTimeOk() (*string, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *ProvisionRequest) SetBootTime(v string)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *ProvisionRequest) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


