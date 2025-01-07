# TenantCreateAdminUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**Cloudprovider** | Pointer to **string** |  | [optional] 
**Markettoken** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantCreateAdminUserRequest

`func NewTenantCreateAdminUserRequest(email string, password string, name string, ) *TenantCreateAdminUserRequest`

NewTenantCreateAdminUserRequest instantiates a new TenantCreateAdminUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateAdminUserRequestWithDefaults

`func NewTenantCreateAdminUserRequestWithDefaults() *TenantCreateAdminUserRequest`

NewTenantCreateAdminUserRequestWithDefaults instantiates a new TenantCreateAdminUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TenantCreateAdminUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TenantCreateAdminUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TenantCreateAdminUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *TenantCreateAdminUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TenantCreateAdminUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TenantCreateAdminUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *TenantCreateAdminUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreateAdminUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreateAdminUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCloudprovider

`func (o *TenantCreateAdminUserRequest) GetCloudprovider() string`

GetCloudprovider returns the Cloudprovider field if non-nil, zero value otherwise.

### GetCloudproviderOk

`func (o *TenantCreateAdminUserRequest) GetCloudproviderOk() (*string, bool)`

GetCloudproviderOk returns a tuple with the Cloudprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudprovider

`func (o *TenantCreateAdminUserRequest) SetCloudprovider(v string)`

SetCloudprovider sets Cloudprovider field to given value.

### HasCloudprovider

`func (o *TenantCreateAdminUserRequest) HasCloudprovider() bool`

HasCloudprovider returns a boolean if a field has been set.

### GetMarkettoken

`func (o *TenantCreateAdminUserRequest) GetMarkettoken() string`

GetMarkettoken returns the Markettoken field if non-nil, zero value otherwise.

### GetMarkettokenOk

`func (o *TenantCreateAdminUserRequest) GetMarkettokenOk() (*string, bool)`

GetMarkettokenOk returns a tuple with the Markettoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkettoken

`func (o *TenantCreateAdminUserRequest) SetMarkettoken(v string)`

SetMarkettoken sets Markettoken field to given value.

### HasMarkettoken

`func (o *TenantCreateAdminUserRequest) HasMarkettoken() bool`

HasMarkettoken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


