# AzureSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureSecret

`func NewAzureSecret() *AzureSecret`

NewAzureSecret instantiates a new AzureSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSecretWithDefaults

`func NewAzureSecretWithDefaults() *AzureSecret`

NewAzureSecretWithDefaults instantiates a new AzureSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *AzureSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AzureSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AzureSecret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AzureSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


