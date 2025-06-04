# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Chap** | Pointer to [**AuthenticationChap**](AuthenticationChap.md) |  | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *Authentication) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Authentication) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Authentication) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Authentication) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetChap

`func (o *Authentication) GetChap() AuthenticationChap`

GetChap returns the Chap field if non-nil, zero value otherwise.

### GetChapOk

`func (o *Authentication) GetChapOk() (*AuthenticationChap, bool)`

GetChapOk returns a tuple with the Chap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChap

`func (o *Authentication) SetChap(v AuthenticationChap)`

SetChap sets Chap field to given value.

### HasChap

`func (o *Authentication) HasChap() bool`

HasChap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


