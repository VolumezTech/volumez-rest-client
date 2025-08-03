# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**CloudProviderType**](CloudProviderType.md) |  | [optional] 
**AwsRole** | Pointer to [**CreateAwsRoleRequest**](CreateAwsRoleRequest.md) |  | [optional] 
**AzureRole** | Pointer to [**CreateAzureRoleRequest**](CreateAzureRoleRequest.md) |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **NullableString** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest() *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateRoleRequest) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateRoleRequest) GetCloudProviderOk() (*CloudProviderType, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateRoleRequest) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateRoleRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAwsRole

`func (o *CreateRoleRequest) GetAwsRole() CreateAwsRoleRequest`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *CreateRoleRequest) GetAwsRoleOk() (*CreateAwsRoleRequest, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *CreateRoleRequest) SetAwsRole(v CreateAwsRoleRequest)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *CreateRoleRequest) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetAzureRole

`func (o *CreateRoleRequest) GetAzureRole() CreateAzureRoleRequest`

GetAzureRole returns the AzureRole field if non-nil, zero value otherwise.

### GetAzureRoleOk

`func (o *CreateRoleRequest) GetAzureRoleOk() (*CreateAzureRoleRequest, bool)`

GetAzureRoleOk returns a tuple with the AzureRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRole

`func (o *CreateRoleRequest) SetAzureRole(v CreateAzureRoleRequest)`

SetAzureRole sets AzureRole field to given value.

### HasAzureRole

`func (o *CreateRoleRequest) HasAzureRole() bool`

HasAzureRole returns a boolean if a field has been set.

### GetRoleId

`func (o *CreateRoleRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateRoleRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateRoleRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CreateRoleRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *CreateRoleRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *CreateRoleRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *CreateRoleRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *CreateRoleRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *CreateRoleRequest) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *CreateRoleRequest) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetPolicies

`func (o *CreateRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CreateRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CreateRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CreateRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CreateRoleRequest) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CreateRoleRequest) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


