# AddUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**InvitedUserToken** | **string** |  | 

## Methods

### NewAddUserRequest

`func NewAddUserRequest(email string, password string, name string, invitedUserToken string, ) *AddUserRequest`

NewAddUserRequest instantiates a new AddUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserRequestWithDefaults

`func NewAddUserRequestWithDefaults() *AddUserRequest`

NewAddUserRequestWithDefaults instantiates a new AddUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AddUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AddUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *AddUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetInvitedUserToken

`func (o *AddUserRequest) GetInvitedUserToken() string`

GetInvitedUserToken returns the InvitedUserToken field if non-nil, zero value otherwise.

### GetInvitedUserTokenOk

`func (o *AddUserRequest) GetInvitedUserTokenOk() (*string, bool)`

GetInvitedUserTokenOk returns a tuple with the InvitedUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserToken

`func (o *AddUserRequest) SetInvitedUserToken(v string)`

SetInvitedUserToken sets InvitedUserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


