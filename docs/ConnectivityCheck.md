# ConnectivityCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorsSelector** | Pointer to **string** |  | [optional] 
**TargetsSelector** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectivityCheck

`func NewConnectivityCheck() *ConnectivityCheck`

NewConnectivityCheck instantiates a new ConnectivityCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityCheckWithDefaults

`func NewConnectivityCheckWithDefaults() *ConnectivityCheck`

NewConnectivityCheckWithDefaults instantiates a new ConnectivityCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorsSelector

`func (o *ConnectivityCheck) GetInitiatorsSelector() string`

GetInitiatorsSelector returns the InitiatorsSelector field if non-nil, zero value otherwise.

### GetInitiatorsSelectorOk

`func (o *ConnectivityCheck) GetInitiatorsSelectorOk() (*string, bool)`

GetInitiatorsSelectorOk returns a tuple with the InitiatorsSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorsSelector

`func (o *ConnectivityCheck) SetInitiatorsSelector(v string)`

SetInitiatorsSelector sets InitiatorsSelector field to given value.

### HasInitiatorsSelector

`func (o *ConnectivityCheck) HasInitiatorsSelector() bool`

HasInitiatorsSelector returns a boolean if a field has been set.

### GetTargetsSelector

`func (o *ConnectivityCheck) GetTargetsSelector() string`

GetTargetsSelector returns the TargetsSelector field if non-nil, zero value otherwise.

### GetTargetsSelectorOk

`func (o *ConnectivityCheck) GetTargetsSelectorOk() (*string, bool)`

GetTargetsSelectorOk returns a tuple with the TargetsSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsSelector

`func (o *ConnectivityCheck) SetTargetsSelector(v string)`

SetTargetsSelector sets TargetsSelector field to given value.

### HasTargetsSelector

`func (o *ConnectivityCheck) HasTargetsSelector() bool`

HasTargetsSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


