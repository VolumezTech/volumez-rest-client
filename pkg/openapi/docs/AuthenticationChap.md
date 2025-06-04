# AuthenticationChap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationChap

`func NewAuthenticationChap() *AuthenticationChap`

NewAuthenticationChap instantiates a new AuthenticationChap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationChapWithDefaults

`func NewAuthenticationChapWithDefaults() *AuthenticationChap`

NewAuthenticationChapWithDefaults instantiates a new AuthenticationChap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AuthenticationChap) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticationChap) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticationChap) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticationChap) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *AuthenticationChap) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthenticationChap) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthenticationChap) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthenticationChap) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


