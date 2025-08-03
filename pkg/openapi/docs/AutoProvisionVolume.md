# AutoProvisionVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**Volume**](Volume.md) |  | 
**InfraPlan** | Pointer to [**AutoProvisionInfraPlan**](AutoProvisionInfraPlan.md) |  | [optional] 
**CloudProvider** | **string** |  | 
**AccountId** | **string** |  | 
**SshKeyName** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**AvailabilityZones** | **[]string** |  | 
**Subnets** | **[]string** |  | 
**OsType** | **string** |  | 

## Methods

### NewAutoProvisionVolume

`func NewAutoProvisionVolume(volume Volume, cloudProvider string, accountId string, region string, availabilityZones []string, subnets []string, osType string, ) *AutoProvisionVolume`

NewAutoProvisionVolume instantiates a new AutoProvisionVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoProvisionVolumeWithDefaults

`func NewAutoProvisionVolumeWithDefaults() *AutoProvisionVolume`

NewAutoProvisionVolumeWithDefaults instantiates a new AutoProvisionVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *AutoProvisionVolume) GetVolume() Volume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *AutoProvisionVolume) GetVolumeOk() (*Volume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *AutoProvisionVolume) SetVolume(v Volume)`

SetVolume sets Volume field to given value.


### GetInfraPlan

`func (o *AutoProvisionVolume) GetInfraPlan() AutoProvisionInfraPlan`

GetInfraPlan returns the InfraPlan field if non-nil, zero value otherwise.

### GetInfraPlanOk

`func (o *AutoProvisionVolume) GetInfraPlanOk() (*AutoProvisionInfraPlan, bool)`

GetInfraPlanOk returns a tuple with the InfraPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraPlan

`func (o *AutoProvisionVolume) SetInfraPlan(v AutoProvisionInfraPlan)`

SetInfraPlan sets InfraPlan field to given value.

### HasInfraPlan

`func (o *AutoProvisionVolume) HasInfraPlan() bool`

HasInfraPlan returns a boolean if a field has been set.

### GetCloudProvider

`func (o *AutoProvisionVolume) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AutoProvisionVolume) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AutoProvisionVolume) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetAccountId

`func (o *AutoProvisionVolume) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoProvisionVolume) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoProvisionVolume) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSshKeyName

`func (o *AutoProvisionVolume) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *AutoProvisionVolume) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *AutoProvisionVolume) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.

### HasSshKeyName

`func (o *AutoProvisionVolume) HasSshKeyName() bool`

HasSshKeyName returns a boolean if a field has been set.

### GetImageId

`func (o *AutoProvisionVolume) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AutoProvisionVolume) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AutoProvisionVolume) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AutoProvisionVolume) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetRegion

`func (o *AutoProvisionVolume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoProvisionVolume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoProvisionVolume) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAvailabilityZones

`func (o *AutoProvisionVolume) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AutoProvisionVolume) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AutoProvisionVolume) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSubnets

`func (o *AutoProvisionVolume) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AutoProvisionVolume) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AutoProvisionVolume) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.


### GetOsType

`func (o *AutoProvisionVolume) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AutoProvisionVolume) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AutoProvisionVolume) SetOsType(v string)`

SetOsType sets OsType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


