/*
Volumez API

Volumez orchestrator API

API version: 1.0.0 - f3a04f74
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleRequest{}

// CreateRoleRequest struct for CreateRoleRequest
type CreateRoleRequest struct {
	CloudProvider *CloudProviderType      `json:"cloudProvider,omitempty"`
	AwsRole       *CreateAwsRoleRequest   `json:"awsRole,omitempty"`
	AzureRole     *CreateAzureRoleRequest `json:"azureRole,omitempty"`
	RoleId        *string                 `json:"RoleId,omitempty"`
	RoleName      NullableString          `json:"RoleName,omitempty"`
	Policies      []string                `json:"Policies,omitempty"`
}

// NewCreateRoleRequest instantiates a new CreateRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleRequest() *CreateRoleRequest {
	this := CreateRoleRequest{}
	return &this
}

// NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleRequestWithDefaults() *CreateRoleRequest {
	this := CreateRoleRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetCloudProvider() CloudProviderType {
	if o == nil || IsNil(o.CloudProvider) {
		var ret CloudProviderType
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetCloudProviderOk() (*CloudProviderType, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given CloudProviderType and assigns it to the CloudProvider field.
func (o *CreateRoleRequest) SetCloudProvider(v CloudProviderType) {
	o.CloudProvider = &v
}

// GetAwsRole returns the AwsRole field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetAwsRole() CreateAwsRoleRequest {
	if o == nil || IsNil(o.AwsRole) {
		var ret CreateAwsRoleRequest
		return ret
	}
	return *o.AwsRole
}

// GetAwsRoleOk returns a tuple with the AwsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetAwsRoleOk() (*CreateAwsRoleRequest, bool) {
	if o == nil || IsNil(o.AwsRole) {
		return nil, false
	}
	return o.AwsRole, true
}

// HasAwsRole returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasAwsRole() bool {
	if o != nil && !IsNil(o.AwsRole) {
		return true
	}

	return false
}

// SetAwsRole gets a reference to the given CreateAwsRoleRequest and assigns it to the AwsRole field.
func (o *CreateRoleRequest) SetAwsRole(v CreateAwsRoleRequest) {
	o.AwsRole = &v
}

// GetAzureRole returns the AzureRole field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetAzureRole() CreateAzureRoleRequest {
	if o == nil || IsNil(o.AzureRole) {
		var ret CreateAzureRoleRequest
		return ret
	}
	return *o.AzureRole
}

// GetAzureRoleOk returns a tuple with the AzureRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetAzureRoleOk() (*CreateAzureRoleRequest, bool) {
	if o == nil || IsNil(o.AzureRole) {
		return nil, false
	}
	return o.AzureRole, true
}

// HasAzureRole returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasAzureRole() bool {
	if o != nil && !IsNil(o.AzureRole) {
		return true
	}

	return false
}

// SetAzureRole gets a reference to the given CreateAzureRoleRequest and assigns it to the AzureRole field.
func (o *CreateRoleRequest) SetAzureRole(v CreateAzureRoleRequest) {
	o.AzureRole = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *CreateRoleRequest) SetRoleId(v string) {
	o.RoleId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoleRequest) GetRoleName() string {
	if o == nil || IsNil(o.RoleName.Get()) {
		var ret string
		return ret
	}
	return *o.RoleName.Get()
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleRequest) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleName.Get(), o.RoleName.IsSet()
}

// HasRoleName returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasRoleName() bool {
	if o != nil && o.RoleName.IsSet() {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given NullableString and assigns it to the RoleName field.
func (o *CreateRoleRequest) SetRoleName(v string) {
	o.RoleName.Set(&v)
}

// SetRoleNameNil sets the value for RoleName to be an explicit nil
func (o *CreateRoleRequest) SetRoleNameNil() {
	o.RoleName.Set(nil)
}

// UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
func (o *CreateRoleRequest) UnsetRoleName() {
	o.RoleName.Unset()
}

// GetPolicies returns the Policies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoleRequest) GetPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *CreateRoleRequest) SetPolicies(v []string) {
	o.Policies = v
}

func (o CreateRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.AwsRole) {
		toSerialize["awsRole"] = o.AwsRole
	}
	if !IsNil(o.AzureRole) {
		toSerialize["azureRole"] = o.AzureRole
	}
	if !IsNil(o.RoleId) {
		toSerialize["RoleId"] = o.RoleId
	}
	if o.RoleName.IsSet() {
		toSerialize["RoleName"] = o.RoleName.Get()
	}
	if o.Policies != nil {
		toSerialize["Policies"] = o.Policies
	}
	return toSerialize, nil
}

type NullableCreateRoleRequest struct {
	value *CreateRoleRequest
	isSet bool
}

func (v NullableCreateRoleRequest) Get() *CreateRoleRequest {
	return v.value
}

func (v *NullableCreateRoleRequest) Set(val *CreateRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleRequest(val *CreateRoleRequest) *NullableCreateRoleRequest {
	return &NullableCreateRoleRequest{value: val, isSet: true}
}

func (v NullableCreateRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
