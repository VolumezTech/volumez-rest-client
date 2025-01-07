# GetTenantSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pricing** | Pointer to [**NullableGetTenantSettingsResponsePricing**](GetTenantSettingsResponsePricing.md) |  | [optional] 
**Declarative** | Pointer to [**NullableGetTenantSettingsResponseDeclarative**](GetTenantSettingsResponseDeclarative.md) |  | [optional] 
**Throttle** | Pointer to [**NullableGetTenantSettingsResponseThrottle**](GetTenantSettingsResponseThrottle.md) |  | [optional] 

## Methods

### NewGetTenantSettingsResponse

`func NewGetTenantSettingsResponse() *GetTenantSettingsResponse`

NewGetTenantSettingsResponse instantiates a new GetTenantSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantSettingsResponseWithDefaults

`func NewGetTenantSettingsResponseWithDefaults() *GetTenantSettingsResponse`

NewGetTenantSettingsResponseWithDefaults instantiates a new GetTenantSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricing

`func (o *GetTenantSettingsResponse) GetPricing() GetTenantSettingsResponsePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *GetTenantSettingsResponse) GetPricingOk() (*GetTenantSettingsResponsePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *GetTenantSettingsResponse) SetPricing(v GetTenantSettingsResponsePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *GetTenantSettingsResponse) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *GetTenantSettingsResponse) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *GetTenantSettingsResponse) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetDeclarative

`func (o *GetTenantSettingsResponse) GetDeclarative() GetTenantSettingsResponseDeclarative`

GetDeclarative returns the Declarative field if non-nil, zero value otherwise.

### GetDeclarativeOk

`func (o *GetTenantSettingsResponse) GetDeclarativeOk() (*GetTenantSettingsResponseDeclarative, bool)`

GetDeclarativeOk returns a tuple with the Declarative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarative

`func (o *GetTenantSettingsResponse) SetDeclarative(v GetTenantSettingsResponseDeclarative)`

SetDeclarative sets Declarative field to given value.

### HasDeclarative

`func (o *GetTenantSettingsResponse) HasDeclarative() bool`

HasDeclarative returns a boolean if a field has been set.

### SetDeclarativeNil

`func (o *GetTenantSettingsResponse) SetDeclarativeNil(b bool)`

 SetDeclarativeNil sets the value for Declarative to be an explicit nil

### UnsetDeclarative
`func (o *GetTenantSettingsResponse) UnsetDeclarative()`

UnsetDeclarative ensures that no value is present for Declarative, not even an explicit nil
### GetThrottle

`func (o *GetTenantSettingsResponse) GetThrottle() GetTenantSettingsResponseThrottle`

GetThrottle returns the Throttle field if non-nil, zero value otherwise.

### GetThrottleOk

`func (o *GetTenantSettingsResponse) GetThrottleOk() (*GetTenantSettingsResponseThrottle, bool)`

GetThrottleOk returns a tuple with the Throttle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottle

`func (o *GetTenantSettingsResponse) SetThrottle(v GetTenantSettingsResponseThrottle)`

SetThrottle sets Throttle field to given value.

### HasThrottle

`func (o *GetTenantSettingsResponse) HasThrottle() bool`

HasThrottle returns a boolean if a field has been set.

### SetThrottleNil

`func (o *GetTenantSettingsResponse) SetThrottleNil(b bool)`

 SetThrottleNil sets the value for Throttle to be an explicit nil

### UnsetThrottle
`func (o *GetTenantSettingsResponse) UnsetThrottle()`

UnsetThrottle ensures that no value is present for Throttle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


