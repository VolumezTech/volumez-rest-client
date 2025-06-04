# UpdateAzureRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**ActiveDirectoryId** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to [**AzureSecret**](AzureSecret.md) |  | [optional] 

## Methods

### NewUpdateAzureRoleRequest

`func NewUpdateAzureRoleRequest() *UpdateAzureRoleRequest`

NewUpdateAzureRoleRequest instantiates a new UpdateAzureRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAzureRoleRequestWithDefaults

`func NewUpdateAzureRoleRequestWithDefaults() *UpdateAzureRoleRequest`

NewUpdateAzureRoleRequestWithDefaults instantiates a new UpdateAzureRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *UpdateAzureRoleRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *UpdateAzureRoleRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *UpdateAzureRoleRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *UpdateAzureRoleRequest) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetActiveDirectoryId

`func (o *UpdateAzureRoleRequest) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *UpdateAzureRoleRequest) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *UpdateAzureRoleRequest) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *UpdateAzureRoleRequest) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateAzureRoleRequest) GetSecret() AzureSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateAzureRoleRequest) GetSecretOk() (*AzureSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateAzureRoleRequest) SetSecret(v AzureSecret)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateAzureRoleRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


