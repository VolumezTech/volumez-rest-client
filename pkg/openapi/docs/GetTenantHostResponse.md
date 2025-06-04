# GetTenantHostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 
**TenantHost** | Pointer to **string** |  | [optional] 
**TenantID** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTenantHostResponse

`func NewGetTenantHostResponse(statusCode int32, ) *GetTenantHostResponse`

NewGetTenantHostResponse instantiates a new GetTenantHostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantHostResponseWithDefaults

`func NewGetTenantHostResponseWithDefaults() *GetTenantHostResponse`

NewGetTenantHostResponseWithDefaults instantiates a new GetTenantHostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetTenantHostResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetTenantHostResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetTenantHostResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *GetTenantHostResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetTenantHostResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetTenantHostResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetTenantHostResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTenantHost

`func (o *GetTenantHostResponse) GetTenantHost() string`

GetTenantHost returns the TenantHost field if non-nil, zero value otherwise.

### GetTenantHostOk

`func (o *GetTenantHostResponse) GetTenantHostOk() (*string, bool)`

GetTenantHostOk returns a tuple with the TenantHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantHost

`func (o *GetTenantHostResponse) SetTenantHost(v string)`

SetTenantHost sets TenantHost field to given value.

### HasTenantHost

`func (o *GetTenantHostResponse) HasTenantHost() bool`

HasTenantHost returns a boolean if a field has been set.

### GetTenantID

`func (o *GetTenantHostResponse) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *GetTenantHostResponse) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *GetTenantHostResponse) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *GetTenantHostResponse) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


