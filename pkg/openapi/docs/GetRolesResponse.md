# GetRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]GetRolesResponseRolesInner**](GetRolesResponseRolesInner.md) |  | 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **NullableString** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetRolesResponse

`func NewGetRolesResponse(roles []GetRolesResponseRolesInner, ) *GetRolesResponse`

NewGetRolesResponse instantiates a new GetRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRolesResponseWithDefaults

`func NewGetRolesResponseWithDefaults() *GetRolesResponse`

NewGetRolesResponseWithDefaults instantiates a new GetRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GetRolesResponse) GetRoles() []GetRolesResponseRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetRolesResponse) GetRolesOk() (*[]GetRolesResponseRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetRolesResponse) SetRoles(v []GetRolesResponseRolesInner)`

SetRoles sets Roles field to given value.


### GetRoleId

`func (o *GetRolesResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GetRolesResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GetRolesResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *GetRolesResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *GetRolesResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *GetRolesResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *GetRolesResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *GetRolesResponse) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *GetRolesResponse) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *GetRolesResponse) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetPolicies

`func (o *GetRolesResponse) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GetRolesResponse) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GetRolesResponse) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *GetRolesResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *GetRolesResponse) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *GetRolesResponse) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


