# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alertid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Severity** | Pointer to **string** |  | [optional] [readonly] 
**Objecttype** | Pointer to **string** |  | [optional] [readonly] 
**Objectid** | Pointer to **string** |  | [optional] [readonly] 
**Creationtime** | Pointer to **string** |  | [optional] [readonly] 
**Lastsendtime** | Pointer to **string** |  | [optional] [readonly] 
**Cleartime** | Pointer to **string** |  | [optional] [readonly] 
**Details** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertid

`func (o *Alert) GetAlertid() string`

GetAlertid returns the Alertid field if non-nil, zero value otherwise.

### GetAlertidOk

`func (o *Alert) GetAlertidOk() (*string, bool)`

GetAlertidOk returns a tuple with the Alertid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertid

`func (o *Alert) SetAlertid(v string)`

SetAlertid sets Alertid field to given value.

### HasAlertid

`func (o *Alert) HasAlertid() bool`

HasAlertid returns a boolean if a field has been set.

### GetType

`func (o *Alert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Alert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *Alert) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Alert) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Alert) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Alert) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSeverity

`func (o *Alert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Alert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetObjecttype

`func (o *Alert) GetObjecttype() string`

GetObjecttype returns the Objecttype field if non-nil, zero value otherwise.

### GetObjecttypeOk

`func (o *Alert) GetObjecttypeOk() (*string, bool)`

GetObjecttypeOk returns a tuple with the Objecttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjecttype

`func (o *Alert) SetObjecttype(v string)`

SetObjecttype sets Objecttype field to given value.

### HasObjecttype

`func (o *Alert) HasObjecttype() bool`

HasObjecttype returns a boolean if a field has been set.

### GetObjectid

`func (o *Alert) GetObjectid() string`

GetObjectid returns the Objectid field if non-nil, zero value otherwise.

### GetObjectidOk

`func (o *Alert) GetObjectidOk() (*string, bool)`

GetObjectidOk returns a tuple with the Objectid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectid

`func (o *Alert) SetObjectid(v string)`

SetObjectid sets Objectid field to given value.

### HasObjectid

`func (o *Alert) HasObjectid() bool`

HasObjectid returns a boolean if a field has been set.

### GetCreationtime

`func (o *Alert) GetCreationtime() string`

GetCreationtime returns the Creationtime field if non-nil, zero value otherwise.

### GetCreationtimeOk

`func (o *Alert) GetCreationtimeOk() (*string, bool)`

GetCreationtimeOk returns a tuple with the Creationtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationtime

`func (o *Alert) SetCreationtime(v string)`

SetCreationtime sets Creationtime field to given value.

### HasCreationtime

`func (o *Alert) HasCreationtime() bool`

HasCreationtime returns a boolean if a field has been set.

### GetLastsendtime

`func (o *Alert) GetLastsendtime() string`

GetLastsendtime returns the Lastsendtime field if non-nil, zero value otherwise.

### GetLastsendtimeOk

`func (o *Alert) GetLastsendtimeOk() (*string, bool)`

GetLastsendtimeOk returns a tuple with the Lastsendtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastsendtime

`func (o *Alert) SetLastsendtime(v string)`

SetLastsendtime sets Lastsendtime field to given value.

### HasLastsendtime

`func (o *Alert) HasLastsendtime() bool`

HasLastsendtime returns a boolean if a field has been set.

### GetCleartime

`func (o *Alert) GetCleartime() string`

GetCleartime returns the Cleartime field if non-nil, zero value otherwise.

### GetCleartimeOk

`func (o *Alert) GetCleartimeOk() (*string, bool)`

GetCleartimeOk returns a tuple with the Cleartime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleartime

`func (o *Alert) SetCleartime(v string)`

SetCleartime sets Cleartime field to given value.

### HasCleartime

`func (o *Alert) HasCleartime() bool`

HasCleartime returns a boolean if a field has been set.

### GetDetails

`func (o *Alert) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Alert) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Alert) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Alert) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


