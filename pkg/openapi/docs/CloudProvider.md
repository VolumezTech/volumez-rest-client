# CloudProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CloudProviderType**](CloudProviderType.md) |  | 
**AccountID** | **string** |  | 

## Methods

### NewCloudProvider

`func NewCloudProvider(type_ CloudProviderType, accountID string, ) *CloudProvider`

NewCloudProvider instantiates a new CloudProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderWithDefaults

`func NewCloudProviderWithDefaults() *CloudProvider`

NewCloudProviderWithDefaults instantiates a new CloudProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudProvider) GetType() CloudProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudProvider) GetTypeOk() (*CloudProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudProvider) SetType(v CloudProviderType)`

SetType sets Type field to given value.


### GetAccountID

`func (o *CloudProvider) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *CloudProvider) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *CloudProvider) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


