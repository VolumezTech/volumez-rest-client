# MediaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iopsread** | Pointer to **int32** |  | [optional] 
**Iopswrite** | Pointer to **int32** |  | [optional] 
**Bandwidthread** | Pointer to **int32** |  | [optional] 
**Bandwidthwrite** | Pointer to **int32** |  | [optional] 
**Latencyread** | Pointer to **int32** |  | [optional] 
**Latencywrite** | Pointer to **int32** |  | [optional] 
**Freeiopsread** | Pointer to **int32** |  | [optional] 
**Freeiopswrite** | Pointer to **int32** |  | [optional] 
**Freebandwidthread** | Pointer to **int32** |  | [optional] 
**Freebandwidthwrite** | Pointer to **int32** |  | [optional] 

## Methods

### NewMediaProfile

`func NewMediaProfile() *MediaProfile`

NewMediaProfile instantiates a new MediaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaProfileWithDefaults

`func NewMediaProfileWithDefaults() *MediaProfile`

NewMediaProfileWithDefaults instantiates a new MediaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIopsread

`func (o *MediaProfile) GetIopsread() int32`

GetIopsread returns the Iopsread field if non-nil, zero value otherwise.

### GetIopsreadOk

`func (o *MediaProfile) GetIopsreadOk() (*int32, bool)`

GetIopsreadOk returns a tuple with the Iopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsread

`func (o *MediaProfile) SetIopsread(v int32)`

SetIopsread sets Iopsread field to given value.

### HasIopsread

`func (o *MediaProfile) HasIopsread() bool`

HasIopsread returns a boolean if a field has been set.

### GetIopswrite

`func (o *MediaProfile) GetIopswrite() int32`

GetIopswrite returns the Iopswrite field if non-nil, zero value otherwise.

### GetIopswriteOk

`func (o *MediaProfile) GetIopswriteOk() (*int32, bool)`

GetIopswriteOk returns a tuple with the Iopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopswrite

`func (o *MediaProfile) SetIopswrite(v int32)`

SetIopswrite sets Iopswrite field to given value.

### HasIopswrite

`func (o *MediaProfile) HasIopswrite() bool`

HasIopswrite returns a boolean if a field has been set.

### GetBandwidthread

`func (o *MediaProfile) GetBandwidthread() int32`

GetBandwidthread returns the Bandwidthread field if non-nil, zero value otherwise.

### GetBandwidthreadOk

`func (o *MediaProfile) GetBandwidthreadOk() (*int32, bool)`

GetBandwidthreadOk returns a tuple with the Bandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthread

`func (o *MediaProfile) SetBandwidthread(v int32)`

SetBandwidthread sets Bandwidthread field to given value.

### HasBandwidthread

`func (o *MediaProfile) HasBandwidthread() bool`

HasBandwidthread returns a boolean if a field has been set.

### GetBandwidthwrite

`func (o *MediaProfile) GetBandwidthwrite() int32`

GetBandwidthwrite returns the Bandwidthwrite field if non-nil, zero value otherwise.

### GetBandwidthwriteOk

`func (o *MediaProfile) GetBandwidthwriteOk() (*int32, bool)`

GetBandwidthwriteOk returns a tuple with the Bandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthwrite

`func (o *MediaProfile) SetBandwidthwrite(v int32)`

SetBandwidthwrite sets Bandwidthwrite field to given value.

### HasBandwidthwrite

`func (o *MediaProfile) HasBandwidthwrite() bool`

HasBandwidthwrite returns a boolean if a field has been set.

### GetLatencyread

`func (o *MediaProfile) GetLatencyread() int32`

GetLatencyread returns the Latencyread field if non-nil, zero value otherwise.

### GetLatencyreadOk

`func (o *MediaProfile) GetLatencyreadOk() (*int32, bool)`

GetLatencyreadOk returns a tuple with the Latencyread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyread

`func (o *MediaProfile) SetLatencyread(v int32)`

SetLatencyread sets Latencyread field to given value.

### HasLatencyread

`func (o *MediaProfile) HasLatencyread() bool`

HasLatencyread returns a boolean if a field has been set.

### GetLatencywrite

`func (o *MediaProfile) GetLatencywrite() int32`

GetLatencywrite returns the Latencywrite field if non-nil, zero value otherwise.

### GetLatencywriteOk

`func (o *MediaProfile) GetLatencywriteOk() (*int32, bool)`

GetLatencywriteOk returns a tuple with the Latencywrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencywrite

`func (o *MediaProfile) SetLatencywrite(v int32)`

SetLatencywrite sets Latencywrite field to given value.

### HasLatencywrite

`func (o *MediaProfile) HasLatencywrite() bool`

HasLatencywrite returns a boolean if a field has been set.

### GetFreeiopsread

`func (o *MediaProfile) GetFreeiopsread() int32`

GetFreeiopsread returns the Freeiopsread field if non-nil, zero value otherwise.

### GetFreeiopsreadOk

`func (o *MediaProfile) GetFreeiopsreadOk() (*int32, bool)`

GetFreeiopsreadOk returns a tuple with the Freeiopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeiopsread

`func (o *MediaProfile) SetFreeiopsread(v int32)`

SetFreeiopsread sets Freeiopsread field to given value.

### HasFreeiopsread

`func (o *MediaProfile) HasFreeiopsread() bool`

HasFreeiopsread returns a boolean if a field has been set.

### GetFreeiopswrite

`func (o *MediaProfile) GetFreeiopswrite() int32`

GetFreeiopswrite returns the Freeiopswrite field if non-nil, zero value otherwise.

### GetFreeiopswriteOk

`func (o *MediaProfile) GetFreeiopswriteOk() (*int32, bool)`

GetFreeiopswriteOk returns a tuple with the Freeiopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeiopswrite

`func (o *MediaProfile) SetFreeiopswrite(v int32)`

SetFreeiopswrite sets Freeiopswrite field to given value.

### HasFreeiopswrite

`func (o *MediaProfile) HasFreeiopswrite() bool`

HasFreeiopswrite returns a boolean if a field has been set.

### GetFreebandwidthread

`func (o *MediaProfile) GetFreebandwidthread() int32`

GetFreebandwidthread returns the Freebandwidthread field if non-nil, zero value otherwise.

### GetFreebandwidthreadOk

`func (o *MediaProfile) GetFreebandwidthreadOk() (*int32, bool)`

GetFreebandwidthreadOk returns a tuple with the Freebandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreebandwidthread

`func (o *MediaProfile) SetFreebandwidthread(v int32)`

SetFreebandwidthread sets Freebandwidthread field to given value.

### HasFreebandwidthread

`func (o *MediaProfile) HasFreebandwidthread() bool`

HasFreebandwidthread returns a boolean if a field has been set.

### GetFreebandwidthwrite

`func (o *MediaProfile) GetFreebandwidthwrite() int32`

GetFreebandwidthwrite returns the Freebandwidthwrite field if non-nil, zero value otherwise.

### GetFreebandwidthwriteOk

`func (o *MediaProfile) GetFreebandwidthwriteOk() (*int32, bool)`

GetFreebandwidthwriteOk returns a tuple with the Freebandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreebandwidthwrite

`func (o *MediaProfile) SetFreebandwidthwrite(v int32)`

SetFreebandwidthwrite sets Freebandwidthwrite field to given value.

### HasFreebandwidthwrite

`func (o *MediaProfile) HasFreebandwidthwrite() bool`

HasFreebandwidthwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


