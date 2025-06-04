# VirtualMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**Media**](Media.md) |  | [optional] 
**Flavor** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualMedia

`func NewVirtualMedia() *VirtualMedia`

NewVirtualMedia instantiates a new VirtualMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMediaWithDefaults

`func NewVirtualMediaWithDefaults() *VirtualMedia`

NewVirtualMediaWithDefaults instantiates a new VirtualMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *VirtualMedia) GetMedia() Media`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *VirtualMedia) GetMediaOk() (*Media, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *VirtualMedia) SetMedia(v Media)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *VirtualMedia) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetFlavor

`func (o *VirtualMedia) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *VirtualMedia) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *VirtualMedia) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *VirtualMedia) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetMembers

`func (o *VirtualMedia) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VirtualMedia) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VirtualMedia) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *VirtualMedia) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


