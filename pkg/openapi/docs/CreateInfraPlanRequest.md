# CreateInfraPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**Policy** | [**CreateInfraPlanRequestPolicy**](CreateInfraPlanRequestPolicy.md) |  | 
**OsType** | Pointer to [**OSType**](OSType.md) |  | [optional] 
**Size** | **int32** |  | 
**Region** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 
**SubnetIds** | Pointer to **[]string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateInfraPlanRequest

`func NewCreateInfraPlanRequest(cloudProvider CloudProviderType, policy CreateInfraPlanRequestPolicy, size int32, ) *CreateInfraPlanRequest`

NewCreateInfraPlanRequest instantiates a new CreateInfraPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraPlanRequestWithDefaults

`func NewCreateInfraPlanRequestWithDefaults() *CreateInfraPlanRequest`

NewCreateInfraPlanRequestWithDefaults instantiates a new CreateInfraPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateInfraPlanRequest) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateInfraPlanRequest) GetCloudProviderOk() (*CloudProviderType, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateInfraPlanRequest) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.


### GetPolicy

`func (o *CreateInfraPlanRequest) GetPolicy() CreateInfraPlanRequestPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateInfraPlanRequest) GetPolicyOk() (*CreateInfraPlanRequestPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateInfraPlanRequest) SetPolicy(v CreateInfraPlanRequestPolicy)`

SetPolicy sets Policy field to given value.


### GetOsType

`func (o *CreateInfraPlanRequest) GetOsType() OSType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateInfraPlanRequest) GetOsTypeOk() (*OSType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateInfraPlanRequest) SetOsType(v OSType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *CreateInfraPlanRequest) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetSize

`func (o *CreateInfraPlanRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateInfraPlanRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateInfraPlanRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetRegion

`func (o *CreateInfraPlanRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateInfraPlanRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateInfraPlanRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateInfraPlanRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZones

`func (o *CreateInfraPlanRequest) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CreateInfraPlanRequest) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CreateInfraPlanRequest) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CreateInfraPlanRequest) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetSubnetIds

`func (o *CreateInfraPlanRequest) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *CreateInfraPlanRequest) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *CreateInfraPlanRequest) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *CreateInfraPlanRequest) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateInfraPlanRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateInfraPlanRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateInfraPlanRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateInfraPlanRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


