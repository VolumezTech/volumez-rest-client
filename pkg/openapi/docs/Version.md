# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] 
**Major** | Pointer to **int32** |  | [optional] 
**Minor** | Pointer to **int32** |  | [optional] 
**Patch** | Pointer to **int32** |  | [optional] 
**Prerelease** | Pointer to **string** |  | [optional] 

## Methods

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *Version) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Version) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Version) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *Version) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMajor

`func (o *Version) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *Version) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *Version) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *Version) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *Version) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *Version) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *Version) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *Version) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetPatch

`func (o *Version) GetPatch() int32`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *Version) GetPatchOk() (*int32, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *Version) SetPatch(v int32)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *Version) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPrerelease

`func (o *Version) GetPrerelease() string`

GetPrerelease returns the Prerelease field if non-nil, zero value otherwise.

### GetPrereleaseOk

`func (o *Version) GetPrereleaseOk() (*string, bool)`

GetPrereleaseOk returns a tuple with the Prerelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerelease

`func (o *Version) SetPrerelease(v string)`

SetPrerelease sets Prerelease field to given value.

### HasPrerelease

`func (o *Version) HasPrerelease() bool`

HasPrerelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


