# ProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Tenanthosttoken** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Cloudprovider** | Pointer to **string** |  | [optional] 
**Machineinfo** | Pointer to [**MachineInfo**](MachineInfo.md) |  | [optional] 
**Boottime** | Pointer to **string** |  | [optional] 

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

### GetTenanthosttoken

`func (o *ProvisionRequest) GetTenanthosttoken() string`

GetTenanthosttoken returns the Tenanthosttoken field if non-nil, zero value otherwise.

### GetTenanthosttokenOk

`func (o *ProvisionRequest) GetTenanthosttokenOk() (*string, bool)`

GetTenanthosttokenOk returns a tuple with the Tenanthosttoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenanthosttoken

`func (o *ProvisionRequest) SetTenanthosttoken(v string)`

SetTenanthosttoken sets Tenanthosttoken field to given value.

### HasTenanthosttoken

`func (o *ProvisionRequest) HasTenanthosttoken() bool`

HasTenanthosttoken returns a boolean if a field has been set.

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

### GetCloudprovider

`func (o *ProvisionRequest) GetCloudprovider() string`

GetCloudprovider returns the Cloudprovider field if non-nil, zero value otherwise.

### GetCloudproviderOk

`func (o *ProvisionRequest) GetCloudproviderOk() (*string, bool)`

GetCloudproviderOk returns a tuple with the Cloudprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudprovider

`func (o *ProvisionRequest) SetCloudprovider(v string)`

SetCloudprovider sets Cloudprovider field to given value.

### HasCloudprovider

`func (o *ProvisionRequest) HasCloudprovider() bool`

HasCloudprovider returns a boolean if a field has been set.

### GetMachineinfo

`func (o *ProvisionRequest) GetMachineinfo() MachineInfo`

GetMachineinfo returns the Machineinfo field if non-nil, zero value otherwise.

### GetMachineinfoOk

`func (o *ProvisionRequest) GetMachineinfoOk() (*MachineInfo, bool)`

GetMachineinfoOk returns a tuple with the Machineinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineinfo

`func (o *ProvisionRequest) SetMachineinfo(v MachineInfo)`

SetMachineinfo sets Machineinfo field to given value.

### HasMachineinfo

`func (o *ProvisionRequest) HasMachineinfo() bool`

HasMachineinfo returns a boolean if a field has been set.

### GetBoottime

`func (o *ProvisionRequest) GetBoottime() string`

GetBoottime returns the Boottime field if non-nil, zero value otherwise.

### GetBoottimeOk

`func (o *ProvisionRequest) GetBoottimeOk() (*string, bool)`

GetBoottimeOk returns a tuple with the Boottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoottime

`func (o *ProvisionRequest) SetBoottime(v string)`

SetBoottime sets Boottime field to given value.

### HasBoottime

`func (o *ProvisionRequest) HasBoottime() bool`

HasBoottime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


