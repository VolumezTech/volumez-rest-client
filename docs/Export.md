# Export

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**ExportCreate**](ExportCreate.md) |  | [optional] 
**Volumename** | Pointer to **string** |  | [optional] 
**Snapshotname** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Xqn** | Pointer to **string** |  | [optional] [readonly] 
**Wwn** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewExport

`func NewExport() *Export`

NewExport instantiates a new Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportWithDefaults

`func NewExportWithDefaults() *Export`

NewExportWithDefaults instantiates a new Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Export) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Export) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Export) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Export) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParams

`func (o *Export) GetParams() ExportCreate`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Export) GetParamsOk() (*ExportCreate, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Export) SetParams(v ExportCreate)`

SetParams sets Params field to given value.

### HasParams

`func (o *Export) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetVolumename

`func (o *Export) GetVolumename() string`

GetVolumename returns the Volumename field if non-nil, zero value otherwise.

### GetVolumenameOk

`func (o *Export) GetVolumenameOk() (*string, bool)`

GetVolumenameOk returns a tuple with the Volumename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumename

`func (o *Export) SetVolumename(v string)`

SetVolumename sets Volumename field to given value.

### HasVolumename

`func (o *Export) HasVolumename() bool`

HasVolumename returns a boolean if a field has been set.

### GetSnapshotname

`func (o *Export) GetSnapshotname() string`

GetSnapshotname returns the Snapshotname field if non-nil, zero value otherwise.

### GetSnapshotnameOk

`func (o *Export) GetSnapshotnameOk() (*string, bool)`

GetSnapshotnameOk returns a tuple with the Snapshotname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotname

`func (o *Export) SetSnapshotname(v string)`

SetSnapshotname sets Snapshotname field to given value.

### HasSnapshotname

`func (o *Export) HasSnapshotname() bool`

HasSnapshotname returns a boolean if a field has been set.

### GetState

`func (o *Export) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Export) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Export) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Export) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Export) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Export) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Export) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Export) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Export) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Export) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Export) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Export) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetXqn

`func (o *Export) GetXqn() string`

GetXqn returns the Xqn field if non-nil, zero value otherwise.

### GetXqnOk

`func (o *Export) GetXqnOk() (*string, bool)`

GetXqnOk returns a tuple with the Xqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXqn

`func (o *Export) SetXqn(v string)`

SetXqn sets Xqn field to given value.

### HasXqn

`func (o *Export) HasXqn() bool`

HasXqn returns a boolean if a field has been set.

### GetWwn

`func (o *Export) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *Export) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *Export) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *Export) HasWwn() bool`

HasWwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


