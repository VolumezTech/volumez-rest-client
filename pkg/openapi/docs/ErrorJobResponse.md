# ErrorJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ObjectID** | Pointer to **string** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorJobResponse

`func NewErrorJobResponse() *ErrorJobResponse`

NewErrorJobResponse instantiates a new ErrorJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorJobResponseWithDefaults

`func NewErrorJobResponseWithDefaults() *ErrorJobResponse`

NewErrorJobResponseWithDefaults instantiates a new ErrorJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorJobResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorJobResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorJobResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorJobResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *ErrorJobResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorJobResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorJobResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorJobResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetObjectID

`func (o *ErrorJobResponse) GetObjectID() string`

GetObjectID returns the ObjectID field if non-nil, zero value otherwise.

### GetObjectIDOk

`func (o *ErrorJobResponse) GetObjectIDOk() (*string, bool)`

GetObjectIDOk returns a tuple with the ObjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectID

`func (o *ErrorJobResponse) SetObjectID(v string)`

SetObjectID sets ObjectID field to given value.

### HasObjectID

`func (o *ErrorJobResponse) HasObjectID() bool`

HasObjectID returns a boolean if a field has been set.

### GetJobID

`func (o *ErrorJobResponse) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *ErrorJobResponse) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *ErrorJobResponse) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *ErrorJobResponse) HasJobID() bool`

HasJobID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


