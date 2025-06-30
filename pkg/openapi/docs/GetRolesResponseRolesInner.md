# GetRolesResponseRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProvider**](CloudProvider.md) |  | 
**AwsRole** | Pointer to [**AWSRoleItem**](AWSRoleItem.md) |  | [optional] 
**AzureRole** | Pointer to [**AzureRoleItem**](AzureRoleItem.md) |  | [optional] 

## Methods

### NewGetRolesResponseRolesInner

`func NewGetRolesResponseRolesInner(cloudProvider CloudProvider, ) *GetRolesResponseRolesInner`

NewGetRolesResponseRolesInner instantiates a new GetRolesResponseRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRolesResponseRolesInnerWithDefaults

`func NewGetRolesResponseRolesInnerWithDefaults() *GetRolesResponseRolesInner`

NewGetRolesResponseRolesInnerWithDefaults instantiates a new GetRolesResponseRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *GetRolesResponseRolesInner) GetCloudProvider() CloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *GetRolesResponseRolesInner) GetCloudProviderOk() (*CloudProvider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *GetRolesResponseRolesInner) SetCloudProvider(v CloudProvider)`

SetCloudProvider sets CloudProvider field to given value.


### GetAwsRole

`func (o *GetRolesResponseRolesInner) GetAwsRole() AWSRoleItem`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *GetRolesResponseRolesInner) GetAwsRoleOk() (*AWSRoleItem, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *GetRolesResponseRolesInner) SetAwsRole(v AWSRoleItem)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *GetRolesResponseRolesInner) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetAzureRole

`func (o *GetRolesResponseRolesInner) GetAzureRole() AzureRoleItem`

GetAzureRole returns the AzureRole field if non-nil, zero value otherwise.

### GetAzureRoleOk

`func (o *GetRolesResponseRolesInner) GetAzureRoleOk() (*AzureRoleItem, bool)`

GetAzureRoleOk returns a tuple with the AzureRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRole

`func (o *GetRolesResponseRolesInner) SetAzureRole(v AzureRoleItem)`

SetAzureRole sets AzureRole field to given value.

### HasAzureRole

`func (o *GetRolesResponseRolesInner) HasAzureRole() bool`

HasAzureRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


