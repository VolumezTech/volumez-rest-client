# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Starttime** | Pointer to **string** |  | [optional] 
**Endtime** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Useremail** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *Job) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Job) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Job) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Job) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetArgs

`func (o *Job) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Job) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Job) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Job) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetState

`func (o *Job) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Job) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Job) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Job) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Job) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Job) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Job) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Job) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStarttime

`func (o *Job) GetStarttime() string`

GetStarttime returns the Starttime field if non-nil, zero value otherwise.

### GetStarttimeOk

`func (o *Job) GetStarttimeOk() (*string, bool)`

GetStarttimeOk returns a tuple with the Starttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttime

`func (o *Job) SetStarttime(v string)`

SetStarttime sets Starttime field to given value.

### HasStarttime

`func (o *Job) HasStarttime() bool`

HasStarttime returns a boolean if a field has been set.

### GetEndtime

`func (o *Job) GetEndtime() string`

GetEndtime returns the Endtime field if non-nil, zero value otherwise.

### GetEndtimeOk

`func (o *Job) GetEndtimeOk() (*string, bool)`

GetEndtimeOk returns a tuple with the Endtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndtime

`func (o *Job) SetEndtime(v string)`

SetEndtime sets Endtime field to given value.

### HasEndtime

`func (o *Job) HasEndtime() bool`

HasEndtime returns a boolean if a field has been set.

### GetUsername

`func (o *Job) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Job) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Job) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Job) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUseremail

`func (o *Job) GetUseremail() string`

GetUseremail returns the Useremail field if non-nil, zero value otherwise.

### GetUseremailOk

`func (o *Job) GetUseremailOk() (*string, bool)`

GetUseremailOk returns a tuple with the Useremail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseremail

`func (o *Job) SetUseremail(v string)`

SetUseremail sets Useremail field to given value.

### HasUseremail

`func (o *Job) HasUseremail() bool`

HasUseremail returns a boolean if a field has been set.

### GetDetails

`func (o *Job) GetDetails() []string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Job) GetDetailsOk() (*[]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Job) SetDetails(v []string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Job) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


