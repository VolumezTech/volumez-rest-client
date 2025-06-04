# GetTenantSettingsResponseThrottleAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSnapshotResources** | Pointer to **NullableInt32** |  | [optional] 
**MinReservedForTop** | Pointer to **NullableInt32** |  | [optional] 
**EnableCagThrottling** | Pointer to **NullableBool** |  | [optional] 
**EnableRaidThrottling** | Pointer to **NullableBool** |  | [optional] 
**EnableSliceThrottling** | Pointer to **NullableBool** |  | [optional] 
**EnableMultipathThrottling** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewGetTenantSettingsResponseThrottleAttachment

`func NewGetTenantSettingsResponseThrottleAttachment() *GetTenantSettingsResponseThrottleAttachment`

NewGetTenantSettingsResponseThrottleAttachment instantiates a new GetTenantSettingsResponseThrottleAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantSettingsResponseThrottleAttachmentWithDefaults

`func NewGetTenantSettingsResponseThrottleAttachmentWithDefaults() *GetTenantSettingsResponseThrottleAttachment`

NewGetTenantSettingsResponseThrottleAttachmentWithDefaults instantiates a new GetTenantSettingsResponseThrottleAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSnapshotResources

`func (o *GetTenantSettingsResponseThrottleAttachment) GetDefaultSnapshotResources() int32`

GetDefaultSnapshotResources returns the DefaultSnapshotResources field if non-nil, zero value otherwise.

### GetDefaultSnapshotResourcesOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetDefaultSnapshotResourcesOk() (*int32, bool)`

GetDefaultSnapshotResourcesOk returns a tuple with the DefaultSnapshotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSnapshotResources

`func (o *GetTenantSettingsResponseThrottleAttachment) SetDefaultSnapshotResources(v int32)`

SetDefaultSnapshotResources sets DefaultSnapshotResources field to given value.

### HasDefaultSnapshotResources

`func (o *GetTenantSettingsResponseThrottleAttachment) HasDefaultSnapshotResources() bool`

HasDefaultSnapshotResources returns a boolean if a field has been set.

### SetDefaultSnapshotResourcesNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetDefaultSnapshotResourcesNil(b bool)`

 SetDefaultSnapshotResourcesNil sets the value for DefaultSnapshotResources to be an explicit nil

### UnsetDefaultSnapshotResources
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetDefaultSnapshotResources()`

UnsetDefaultSnapshotResources ensures that no value is present for DefaultSnapshotResources, not even an explicit nil
### GetMinReservedForTop

`func (o *GetTenantSettingsResponseThrottleAttachment) GetMinReservedForTop() int32`

GetMinReservedForTop returns the MinReservedForTop field if non-nil, zero value otherwise.

### GetMinReservedForTopOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetMinReservedForTopOk() (*int32, bool)`

GetMinReservedForTopOk returns a tuple with the MinReservedForTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReservedForTop

`func (o *GetTenantSettingsResponseThrottleAttachment) SetMinReservedForTop(v int32)`

SetMinReservedForTop sets MinReservedForTop field to given value.

### HasMinReservedForTop

`func (o *GetTenantSettingsResponseThrottleAttachment) HasMinReservedForTop() bool`

HasMinReservedForTop returns a boolean if a field has been set.

### SetMinReservedForTopNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetMinReservedForTopNil(b bool)`

 SetMinReservedForTopNil sets the value for MinReservedForTop to be an explicit nil

### UnsetMinReservedForTop
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetMinReservedForTop()`

UnsetMinReservedForTop ensures that no value is present for MinReservedForTop, not even an explicit nil
### GetEnableCagThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableCagThrottling() bool`

GetEnableCagThrottling returns the EnableCagThrottling field if non-nil, zero value otherwise.

### GetEnableCagThrottlingOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableCagThrottlingOk() (*bool, bool)`

GetEnableCagThrottlingOk returns a tuple with the EnableCagThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCagThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableCagThrottling(v bool)`

SetEnableCagThrottling sets EnableCagThrottling field to given value.

### HasEnableCagThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) HasEnableCagThrottling() bool`

HasEnableCagThrottling returns a boolean if a field has been set.

### SetEnableCagThrottlingNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableCagThrottlingNil(b bool)`

 SetEnableCagThrottlingNil sets the value for EnableCagThrottling to be an explicit nil

### UnsetEnableCagThrottling
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetEnableCagThrottling()`

UnsetEnableCagThrottling ensures that no value is present for EnableCagThrottling, not even an explicit nil
### GetEnableRaidThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableRaidThrottling() bool`

GetEnableRaidThrottling returns the EnableRaidThrottling field if non-nil, zero value otherwise.

### GetEnableRaidThrottlingOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableRaidThrottlingOk() (*bool, bool)`

GetEnableRaidThrottlingOk returns a tuple with the EnableRaidThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRaidThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableRaidThrottling(v bool)`

SetEnableRaidThrottling sets EnableRaidThrottling field to given value.

### HasEnableRaidThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) HasEnableRaidThrottling() bool`

HasEnableRaidThrottling returns a boolean if a field has been set.

### SetEnableRaidThrottlingNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableRaidThrottlingNil(b bool)`

 SetEnableRaidThrottlingNil sets the value for EnableRaidThrottling to be an explicit nil

### UnsetEnableRaidThrottling
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetEnableRaidThrottling()`

UnsetEnableRaidThrottling ensures that no value is present for EnableRaidThrottling, not even an explicit nil
### GetEnableSliceThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableSliceThrottling() bool`

GetEnableSliceThrottling returns the EnableSliceThrottling field if non-nil, zero value otherwise.

### GetEnableSliceThrottlingOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableSliceThrottlingOk() (*bool, bool)`

GetEnableSliceThrottlingOk returns a tuple with the EnableSliceThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSliceThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableSliceThrottling(v bool)`

SetEnableSliceThrottling sets EnableSliceThrottling field to given value.

### HasEnableSliceThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) HasEnableSliceThrottling() bool`

HasEnableSliceThrottling returns a boolean if a field has been set.

### SetEnableSliceThrottlingNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableSliceThrottlingNil(b bool)`

 SetEnableSliceThrottlingNil sets the value for EnableSliceThrottling to be an explicit nil

### UnsetEnableSliceThrottling
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetEnableSliceThrottling()`

UnsetEnableSliceThrottling ensures that no value is present for EnableSliceThrottling, not even an explicit nil
### GetEnableMultipathThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableMultipathThrottling() bool`

GetEnableMultipathThrottling returns the EnableMultipathThrottling field if non-nil, zero value otherwise.

### GetEnableMultipathThrottlingOk

`func (o *GetTenantSettingsResponseThrottleAttachment) GetEnableMultipathThrottlingOk() (*bool, bool)`

GetEnableMultipathThrottlingOk returns a tuple with the EnableMultipathThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipathThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableMultipathThrottling(v bool)`

SetEnableMultipathThrottling sets EnableMultipathThrottling field to given value.

### HasEnableMultipathThrottling

`func (o *GetTenantSettingsResponseThrottleAttachment) HasEnableMultipathThrottling() bool`

HasEnableMultipathThrottling returns a boolean if a field has been set.

### SetEnableMultipathThrottlingNil

`func (o *GetTenantSettingsResponseThrottleAttachment) SetEnableMultipathThrottlingNil(b bool)`

 SetEnableMultipathThrottlingNil sets the value for EnableMultipathThrottling to be an explicit nil

### UnsetEnableMultipathThrottling
`func (o *GetTenantSettingsResponseThrottleAttachment) UnsetEnableMultipathThrottling()`

UnsetEnableMultipathThrottling ensures that no value is present for EnableMultipathThrottling, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


