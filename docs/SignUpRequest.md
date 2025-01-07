# SignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**CloudProvider** | Pointer to **string** |  | [optional] 
**MarketToken** | Pointer to **string** |  | [optional] 

## Methods

### NewSignUpRequest

`func NewSignUpRequest(email string, password string, name string, ) *SignUpRequest`

NewSignUpRequest instantiates a new SignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpRequestWithDefaults

`func NewSignUpRequestWithDefaults() *SignUpRequest`

NewSignUpRequestWithDefaults instantiates a new SignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SignUpRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignUpRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignUpRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *SignUpRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignUpRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignUpRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCloudProvider

`func (o *SignUpRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SignUpRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SignUpRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *SignUpRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMarketToken

`func (o *SignUpRequest) GetMarketToken() string`

GetMarketToken returns the MarketToken field if non-nil, zero value otherwise.

### GetMarketTokenOk

`func (o *SignUpRequest) GetMarketTokenOk() (*string, bool)`

GetMarketTokenOk returns a tuple with the MarketToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketToken

`func (o *SignUpRequest) SetMarketToken(v string)`

SetMarketToken sets MarketToken field to given value.

### HasMarketToken

`func (o *SignUpRequest) HasMarketToken() bool`

HasMarketToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


