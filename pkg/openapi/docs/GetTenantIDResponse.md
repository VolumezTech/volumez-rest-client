# GetTenantIDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 
**TenantID** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTenantIDResponse

`func NewGetTenantIDResponse(statusCode int32, ) *GetTenantIDResponse`

NewGetTenantIDResponse instantiates a new GetTenantIDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantIDResponseWithDefaults

`func NewGetTenantIDResponseWithDefaults() *GetTenantIDResponse`

NewGetTenantIDResponseWithDefaults instantiates a new GetTenantIDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetTenantIDResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetTenantIDResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetTenantIDResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *GetTenantIDResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetTenantIDResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetTenantIDResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetTenantIDResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTenantID

`func (o *GetTenantIDResponse) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *GetTenantIDResponse) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *GetTenantIDResponse) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *GetTenantIDResponse) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


