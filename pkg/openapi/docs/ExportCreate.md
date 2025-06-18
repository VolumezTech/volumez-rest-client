# ExportCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetName** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Volumeid** | Pointer to **string** |  | [optional] 
**Snapshotid** | Pointer to **string** |  | [optional] 
**AllowedHosts** | Pointer to **[]string** |  | [optional] 
**Nodes** | Pointer to [**[]ExportCreateNodesInner**](ExportCreateNodesInner.md) |  | [optional] 

## Methods

### NewExportCreate

`func NewExportCreate() *ExportCreate`

NewExportCreate instantiates a new ExportCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportCreateWithDefaults

`func NewExportCreateWithDefaults() *ExportCreate`

NewExportCreateWithDefaults instantiates a new ExportCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetName

`func (o *ExportCreate) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ExportCreate) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ExportCreate) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ExportCreate) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetProtocol

`func (o *ExportCreate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ExportCreate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ExportCreate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ExportCreate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAuthentication

`func (o *ExportCreate) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ExportCreate) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ExportCreate) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ExportCreate) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetVolumeid

`func (o *ExportCreate) GetVolumeid() string`

GetVolumeid returns the Volumeid field if non-nil, zero value otherwise.

### GetVolumeidOk

`func (o *ExportCreate) GetVolumeidOk() (*string, bool)`

GetVolumeidOk returns a tuple with the Volumeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeid

`func (o *ExportCreate) SetVolumeid(v string)`

SetVolumeid sets Volumeid field to given value.

### HasVolumeid

`func (o *ExportCreate) HasVolumeid() bool`

HasVolumeid returns a boolean if a field has been set.

### GetSnapshotid

`func (o *ExportCreate) GetSnapshotid() string`

GetSnapshotid returns the Snapshotid field if non-nil, zero value otherwise.

### GetSnapshotidOk

`func (o *ExportCreate) GetSnapshotidOk() (*string, bool)`

GetSnapshotidOk returns a tuple with the Snapshotid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotid

`func (o *ExportCreate) SetSnapshotid(v string)`

SetSnapshotid sets Snapshotid field to given value.

### HasSnapshotid

`func (o *ExportCreate) HasSnapshotid() bool`

HasSnapshotid returns a boolean if a field has been set.

### GetAllowedHosts

`func (o *ExportCreate) GetAllowedHosts() []string`

GetAllowedHosts returns the AllowedHosts field if non-nil, zero value otherwise.

### GetAllowedHostsOk

`func (o *ExportCreate) GetAllowedHostsOk() (*[]string, bool)`

GetAllowedHostsOk returns a tuple with the AllowedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHosts

`func (o *ExportCreate) SetAllowedHosts(v []string)`

SetAllowedHosts sets AllowedHosts field to given value.

### HasAllowedHosts

`func (o *ExportCreate) HasAllowedHosts() bool`

HasAllowedHosts returns a boolean if a field has been set.

### GetNodes

`func (o *ExportCreate) GetNodes() []ExportCreateNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ExportCreate) GetNodesOk() (*[]ExportCreateNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ExportCreate) SetNodes(v []ExportCreateNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ExportCreate) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


