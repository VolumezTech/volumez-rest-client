# HealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthCheckResponse

`func NewHealthCheckResponse(statusCode int32, ) *HealthCheckResponse`

NewHealthCheckResponse instantiates a new HealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResponseWithDefaults

`func NewHealthCheckResponseWithDefaults() *HealthCheckResponse`

NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *HealthCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthCheckResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *HealthCheckResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HealthCheckResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HealthCheckResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *HealthCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthCheckResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


