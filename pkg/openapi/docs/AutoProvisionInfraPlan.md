# AutoProvisionInfraPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **int32** |  | [optional] 

## Methods

### NewAutoProvisionInfraPlan

`func NewAutoProvisionInfraPlan() *AutoProvisionInfraPlan`

NewAutoProvisionInfraPlan instantiates a new AutoProvisionInfraPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoProvisionInfraPlanWithDefaults

`func NewAutoProvisionInfraPlanWithDefaults() *AutoProvisionInfraPlan`

NewAutoProvisionInfraPlanWithDefaults instantiates a new AutoProvisionInfraPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *AutoProvisionInfraPlan) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AutoProvisionInfraPlan) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AutoProvisionInfraPlan) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AutoProvisionInfraPlan) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AutoProvisionInfraPlan) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AutoProvisionInfraPlan) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AutoProvisionInfraPlan) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AutoProvisionInfraPlan) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *AutoProvisionInfraPlan) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AutoProvisionInfraPlan) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AutoProvisionInfraPlan) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *AutoProvisionInfraPlan) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetOsType

`func (o *AutoProvisionInfraPlan) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AutoProvisionInfraPlan) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AutoProvisionInfraPlan) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AutoProvisionInfraPlan) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPrice

`func (o *AutoProvisionInfraPlan) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AutoProvisionInfraPlan) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AutoProvisionInfraPlan) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AutoProvisionInfraPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


