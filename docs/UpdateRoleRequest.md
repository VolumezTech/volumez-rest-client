# UpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**CloudProviderType**](CloudProviderType.md) |  | [optional] 
**AwsRole** | Pointer to [**UpdateAwsRoleRequest**](UpdateAwsRoleRequest.md) |  | [optional] 
**AzureRole** | Pointer to [**UpdateAzureRoleRequest**](UpdateAzureRoleRequest.md) |  | [optional] 

## Methods

### NewUpdateRoleRequest

`func NewUpdateRoleRequest() *UpdateRoleRequest`

NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleRequestWithDefaults

`func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest`

NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *UpdateRoleRequest) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *UpdateRoleRequest) GetCloudProviderOk() (*CloudProviderType, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *UpdateRoleRequest) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *UpdateRoleRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAwsRole

`func (o *UpdateRoleRequest) GetAwsRole() UpdateAwsRoleRequest`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *UpdateRoleRequest) GetAwsRoleOk() (*UpdateAwsRoleRequest, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *UpdateRoleRequest) SetAwsRole(v UpdateAwsRoleRequest)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *UpdateRoleRequest) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetAzureRole

`func (o *UpdateRoleRequest) GetAzureRole() UpdateAzureRoleRequest`

GetAzureRole returns the AzureRole field if non-nil, zero value otherwise.

### GetAzureRoleOk

`func (o *UpdateRoleRequest) GetAzureRoleOk() (*UpdateAzureRoleRequest, bool)`

GetAzureRoleOk returns a tuple with the AzureRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRole

`func (o *UpdateRoleRequest) SetAzureRole(v UpdateAzureRoleRequest)`

SetAzureRole sets AzureRole field to given value.

### HasAzureRole

`func (o *UpdateRoleRequest) HasAzureRole() bool`

HasAzureRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


