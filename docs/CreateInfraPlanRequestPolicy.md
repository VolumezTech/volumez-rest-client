# CreateInfraPlanRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Settings** | Pointer to [**NullableCreateInfraPlanRequestPolicySettings**](CreateInfraPlanRequestPolicySettings.md) |  | [optional] 

## Methods

### NewCreateInfraPlanRequestPolicy

`func NewCreateInfraPlanRequestPolicy() *CreateInfraPlanRequestPolicy`

NewCreateInfraPlanRequestPolicy instantiates a new CreateInfraPlanRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraPlanRequestPolicyWithDefaults

`func NewCreateInfraPlanRequestPolicyWithDefaults() *CreateInfraPlanRequestPolicy`

NewCreateInfraPlanRequestPolicyWithDefaults instantiates a new CreateInfraPlanRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInfraPlanRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInfraPlanRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInfraPlanRequestPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInfraPlanRequestPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateInfraPlanRequestPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateInfraPlanRequestPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSettings

`func (o *CreateInfraPlanRequestPolicy) GetSettings() CreateInfraPlanRequestPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateInfraPlanRequestPolicy) GetSettingsOk() (*CreateInfraPlanRequestPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateInfraPlanRequestPolicy) SetSettings(v CreateInfraPlanRequestPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateInfraPlanRequestPolicy) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *CreateInfraPlanRequestPolicy) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *CreateInfraPlanRequestPolicy) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


