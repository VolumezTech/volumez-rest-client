# CreateAzureRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryId** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to [**AzureSecret**](AzureSecret.md) |  | [optional] 

## Methods

### NewCreateAzureRoleRequest

`func NewCreateAzureRoleRequest() *CreateAzureRoleRequest`

NewCreateAzureRoleRequest instantiates a new CreateAzureRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureRoleRequestWithDefaults

`func NewCreateAzureRoleRequestWithDefaults() *CreateAzureRoleRequest`

NewCreateAzureRoleRequestWithDefaults instantiates a new CreateAzureRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryId

`func (o *CreateAzureRoleRequest) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *CreateAzureRoleRequest) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *CreateAzureRoleRequest) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *CreateAzureRoleRequest) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetApplicationId

`func (o *CreateAzureRoleRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateAzureRoleRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateAzureRoleRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CreateAzureRoleRequest) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreateAzureRoleRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateAzureRoleRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateAzureRoleRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateAzureRoleRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSecret

`func (o *CreateAzureRoleRequest) GetSecret() AzureSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateAzureRoleRequest) GetSecretOk() (*AzureSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateAzureRoleRequest) SetSecret(v AzureSecret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateAzureRoleRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


