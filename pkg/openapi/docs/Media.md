# Media

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mediaid** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Media** | Pointer to **string** |  | [optional] 
**Bus** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Sed** | Pointer to **bool** |  | [optional] 
**Node** | Pointer to **string** |  | [optional] 
**Cloudprovider** | Pointer to **string** |  | [optional] 
**AccountID** | Pointer to **string** | the media node ResourceNamespace | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**ResourceNamespace** | Pointer to **string** | the media node ResourceNamespace | [optional] 
**PhysicalProximityGroup** | Pointer to **string** | the media node PhysicalProximityGroup | [optional] 
**ResiliencyDomain** | Pointer to **string** | the media node ResiliencyDomain | [optional] 
**FaultDomain** | Pointer to **string** | the media node FaultDomain | [optional] 
**Firmware** | Pointer to **string** |  | [optional] 
**Sectorsize** | Pointer to **int32** |  | [optional] 
**Iopsread** | Pointer to **int32** |  | [optional] 
**Iopswrite** | Pointer to **int32** |  | [optional] 
**Bandwidthread** | Pointer to **int32** |  | [optional] 
**Bandwidthwrite** | Pointer to **int32** |  | [optional] 
**BandwidthReserved** | Pointer to **int32** |  | [optional] 
**Latencyread** | Pointer to **int32** |  | [optional] 
**Latencywrite** | Pointer to **int32** |  | [optional] 
**Offlinetime** | Pointer to **string** |  | [optional] 
**Freesize** | Pointer to **int32** |  | [optional] 
**Freeiopsread** | Pointer to **int32** |  | [optional] 
**Freeiopswrite** | Pointer to **int32** |  | [optional] 
**Freebandwidthread** | Pointer to **int32** |  | [optional] 
**Freebandwidthwrite** | Pointer to **int32** |  | [optional] 
**Volumescount** | Pointer to **int32** | count of how many volumes are using the media | [optional] [readonly] 
**Assignment** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Capacitygroup** | Pointer to **string** | the capacity group this media belongs to | [optional] [readonly] 
**Lbaformats** | Pointer to **[]string** | Available LBA formats for the media — ensure the block size specified in the media assignment matches one of these formats | [optional] 

## Methods

### NewMedia

`func NewMedia() *Media`

NewMedia instantiates a new Media object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaWithDefaults

`func NewMediaWithDefaults() *Media`

NewMediaWithDefaults instantiates a new Media object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaid

`func (o *Media) GetMediaid() string`

GetMediaid returns the Mediaid field if non-nil, zero value otherwise.

### GetMediaidOk

`func (o *Media) GetMediaidOk() (*string, bool)`

GetMediaidOk returns a tuple with the Mediaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaid

`func (o *Media) SetMediaid(v string)`

SetMediaid sets Mediaid field to given value.

### HasMediaid

`func (o *Media) HasMediaid() bool`

HasMediaid returns a boolean if a field has been set.

### GetSize

`func (o *Media) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Media) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Media) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Media) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetModel

`func (o *Media) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Media) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Media) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Media) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMedia

`func (o *Media) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Media) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Media) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Media) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetBus

`func (o *Media) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *Media) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *Media) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *Media) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetLocation

`func (o *Media) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Media) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Media) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Media) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSed

`func (o *Media) GetSed() bool`

GetSed returns the Sed field if non-nil, zero value otherwise.

### GetSedOk

`func (o *Media) GetSedOk() (*bool, bool)`

GetSedOk returns a tuple with the Sed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSed

`func (o *Media) SetSed(v bool)`

SetSed sets Sed field to given value.

### HasSed

`func (o *Media) HasSed() bool`

HasSed returns a boolean if a field has been set.

### GetNode

`func (o *Media) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Media) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Media) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *Media) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetCloudprovider

`func (o *Media) GetCloudprovider() string`

GetCloudprovider returns the Cloudprovider field if non-nil, zero value otherwise.

### GetCloudproviderOk

`func (o *Media) GetCloudproviderOk() (*string, bool)`

GetCloudproviderOk returns a tuple with the Cloudprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudprovider

`func (o *Media) SetCloudprovider(v string)`

SetCloudprovider sets Cloudprovider field to given value.

### HasCloudprovider

`func (o *Media) HasCloudprovider() bool`

HasCloudprovider returns a boolean if a field has been set.

### GetAccountID

`func (o *Media) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Media) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Media) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Media) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetRegion

`func (o *Media) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Media) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Media) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Media) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZone

`func (o *Media) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Media) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Media) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Media) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResourceNamespace

`func (o *Media) GetResourceNamespace() string`

GetResourceNamespace returns the ResourceNamespace field if non-nil, zero value otherwise.

### GetResourceNamespaceOk

`func (o *Media) GetResourceNamespaceOk() (*string, bool)`

GetResourceNamespaceOk returns a tuple with the ResourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNamespace

`func (o *Media) SetResourceNamespace(v string)`

SetResourceNamespace sets ResourceNamespace field to given value.

### HasResourceNamespace

`func (o *Media) HasResourceNamespace() bool`

HasResourceNamespace returns a boolean if a field has been set.

### GetPhysicalProximityGroup

`func (o *Media) GetPhysicalProximityGroup() string`

GetPhysicalProximityGroup returns the PhysicalProximityGroup field if non-nil, zero value otherwise.

### GetPhysicalProximityGroupOk

`func (o *Media) GetPhysicalProximityGroupOk() (*string, bool)`

GetPhysicalProximityGroupOk returns a tuple with the PhysicalProximityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalProximityGroup

`func (o *Media) SetPhysicalProximityGroup(v string)`

SetPhysicalProximityGroup sets PhysicalProximityGroup field to given value.

### HasPhysicalProximityGroup

`func (o *Media) HasPhysicalProximityGroup() bool`

HasPhysicalProximityGroup returns a boolean if a field has been set.

### GetResiliencyDomain

`func (o *Media) GetResiliencyDomain() string`

GetResiliencyDomain returns the ResiliencyDomain field if non-nil, zero value otherwise.

### GetResiliencyDomainOk

`func (o *Media) GetResiliencyDomainOk() (*string, bool)`

GetResiliencyDomainOk returns a tuple with the ResiliencyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDomain

`func (o *Media) SetResiliencyDomain(v string)`

SetResiliencyDomain sets ResiliencyDomain field to given value.

### HasResiliencyDomain

`func (o *Media) HasResiliencyDomain() bool`

HasResiliencyDomain returns a boolean if a field has been set.

### GetFaultDomain

`func (o *Media) GetFaultDomain() string`

GetFaultDomain returns the FaultDomain field if non-nil, zero value otherwise.

### GetFaultDomainOk

`func (o *Media) GetFaultDomainOk() (*string, bool)`

GetFaultDomainOk returns a tuple with the FaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomain

`func (o *Media) SetFaultDomain(v string)`

SetFaultDomain sets FaultDomain field to given value.

### HasFaultDomain

`func (o *Media) HasFaultDomain() bool`

HasFaultDomain returns a boolean if a field has been set.

### GetFirmware

`func (o *Media) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *Media) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *Media) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *Media) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetSectorsize

`func (o *Media) GetSectorsize() int32`

GetSectorsize returns the Sectorsize field if non-nil, zero value otherwise.

### GetSectorsizeOk

`func (o *Media) GetSectorsizeOk() (*int32, bool)`

GetSectorsizeOk returns a tuple with the Sectorsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorsize

`func (o *Media) SetSectorsize(v int32)`

SetSectorsize sets Sectorsize field to given value.

### HasSectorsize

`func (o *Media) HasSectorsize() bool`

HasSectorsize returns a boolean if a field has been set.

### GetIopsread

`func (o *Media) GetIopsread() int32`

GetIopsread returns the Iopsread field if non-nil, zero value otherwise.

### GetIopsreadOk

`func (o *Media) GetIopsreadOk() (*int32, bool)`

GetIopsreadOk returns a tuple with the Iopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsread

`func (o *Media) SetIopsread(v int32)`

SetIopsread sets Iopsread field to given value.

### HasIopsread

`func (o *Media) HasIopsread() bool`

HasIopsread returns a boolean if a field has been set.

### GetIopswrite

`func (o *Media) GetIopswrite() int32`

GetIopswrite returns the Iopswrite field if non-nil, zero value otherwise.

### GetIopswriteOk

`func (o *Media) GetIopswriteOk() (*int32, bool)`

GetIopswriteOk returns a tuple with the Iopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopswrite

`func (o *Media) SetIopswrite(v int32)`

SetIopswrite sets Iopswrite field to given value.

### HasIopswrite

`func (o *Media) HasIopswrite() bool`

HasIopswrite returns a boolean if a field has been set.

### GetBandwidthread

`func (o *Media) GetBandwidthread() int32`

GetBandwidthread returns the Bandwidthread field if non-nil, zero value otherwise.

### GetBandwidthreadOk

`func (o *Media) GetBandwidthreadOk() (*int32, bool)`

GetBandwidthreadOk returns a tuple with the Bandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthread

`func (o *Media) SetBandwidthread(v int32)`

SetBandwidthread sets Bandwidthread field to given value.

### HasBandwidthread

`func (o *Media) HasBandwidthread() bool`

HasBandwidthread returns a boolean if a field has been set.

### GetBandwidthwrite

`func (o *Media) GetBandwidthwrite() int32`

GetBandwidthwrite returns the Bandwidthwrite field if non-nil, zero value otherwise.

### GetBandwidthwriteOk

`func (o *Media) GetBandwidthwriteOk() (*int32, bool)`

GetBandwidthwriteOk returns a tuple with the Bandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthwrite

`func (o *Media) SetBandwidthwrite(v int32)`

SetBandwidthwrite sets Bandwidthwrite field to given value.

### HasBandwidthwrite

`func (o *Media) HasBandwidthwrite() bool`

HasBandwidthwrite returns a boolean if a field has been set.

### GetBandwidthReserved

`func (o *Media) GetBandwidthReserved() int32`

GetBandwidthReserved returns the BandwidthReserved field if non-nil, zero value otherwise.

### GetBandwidthReservedOk

`func (o *Media) GetBandwidthReservedOk() (*int32, bool)`

GetBandwidthReservedOk returns a tuple with the BandwidthReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthReserved

`func (o *Media) SetBandwidthReserved(v int32)`

SetBandwidthReserved sets BandwidthReserved field to given value.

### HasBandwidthReserved

`func (o *Media) HasBandwidthReserved() bool`

HasBandwidthReserved returns a boolean if a field has been set.

### GetLatencyread

`func (o *Media) GetLatencyread() int32`

GetLatencyread returns the Latencyread field if non-nil, zero value otherwise.

### GetLatencyreadOk

`func (o *Media) GetLatencyreadOk() (*int32, bool)`

GetLatencyreadOk returns a tuple with the Latencyread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyread

`func (o *Media) SetLatencyread(v int32)`

SetLatencyread sets Latencyread field to given value.

### HasLatencyread

`func (o *Media) HasLatencyread() bool`

HasLatencyread returns a boolean if a field has been set.

### GetLatencywrite

`func (o *Media) GetLatencywrite() int32`

GetLatencywrite returns the Latencywrite field if non-nil, zero value otherwise.

### GetLatencywriteOk

`func (o *Media) GetLatencywriteOk() (*int32, bool)`

GetLatencywriteOk returns a tuple with the Latencywrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencywrite

`func (o *Media) SetLatencywrite(v int32)`

SetLatencywrite sets Latencywrite field to given value.

### HasLatencywrite

`func (o *Media) HasLatencywrite() bool`

HasLatencywrite returns a boolean if a field has been set.

### GetOfflinetime

`func (o *Media) GetOfflinetime() string`

GetOfflinetime returns the Offlinetime field if non-nil, zero value otherwise.

### GetOfflinetimeOk

`func (o *Media) GetOfflinetimeOk() (*string, bool)`

GetOfflinetimeOk returns a tuple with the Offlinetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflinetime

`func (o *Media) SetOfflinetime(v string)`

SetOfflinetime sets Offlinetime field to given value.

### HasOfflinetime

`func (o *Media) HasOfflinetime() bool`

HasOfflinetime returns a boolean if a field has been set.

### GetFreesize

`func (o *Media) GetFreesize() int32`

GetFreesize returns the Freesize field if non-nil, zero value otherwise.

### GetFreesizeOk

`func (o *Media) GetFreesizeOk() (*int32, bool)`

GetFreesizeOk returns a tuple with the Freesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreesize

`func (o *Media) SetFreesize(v int32)`

SetFreesize sets Freesize field to given value.

### HasFreesize

`func (o *Media) HasFreesize() bool`

HasFreesize returns a boolean if a field has been set.

### GetFreeiopsread

`func (o *Media) GetFreeiopsread() int32`

GetFreeiopsread returns the Freeiopsread field if non-nil, zero value otherwise.

### GetFreeiopsreadOk

`func (o *Media) GetFreeiopsreadOk() (*int32, bool)`

GetFreeiopsreadOk returns a tuple with the Freeiopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeiopsread

`func (o *Media) SetFreeiopsread(v int32)`

SetFreeiopsread sets Freeiopsread field to given value.

### HasFreeiopsread

`func (o *Media) HasFreeiopsread() bool`

HasFreeiopsread returns a boolean if a field has been set.

### GetFreeiopswrite

`func (o *Media) GetFreeiopswrite() int32`

GetFreeiopswrite returns the Freeiopswrite field if non-nil, zero value otherwise.

### GetFreeiopswriteOk

`func (o *Media) GetFreeiopswriteOk() (*int32, bool)`

GetFreeiopswriteOk returns a tuple with the Freeiopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeiopswrite

`func (o *Media) SetFreeiopswrite(v int32)`

SetFreeiopswrite sets Freeiopswrite field to given value.

### HasFreeiopswrite

`func (o *Media) HasFreeiopswrite() bool`

HasFreeiopswrite returns a boolean if a field has been set.

### GetFreebandwidthread

`func (o *Media) GetFreebandwidthread() int32`

GetFreebandwidthread returns the Freebandwidthread field if non-nil, zero value otherwise.

### GetFreebandwidthreadOk

`func (o *Media) GetFreebandwidthreadOk() (*int32, bool)`

GetFreebandwidthreadOk returns a tuple with the Freebandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreebandwidthread

`func (o *Media) SetFreebandwidthread(v int32)`

SetFreebandwidthread sets Freebandwidthread field to given value.

### HasFreebandwidthread

`func (o *Media) HasFreebandwidthread() bool`

HasFreebandwidthread returns a boolean if a field has been set.

### GetFreebandwidthwrite

`func (o *Media) GetFreebandwidthwrite() int32`

GetFreebandwidthwrite returns the Freebandwidthwrite field if non-nil, zero value otherwise.

### GetFreebandwidthwriteOk

`func (o *Media) GetFreebandwidthwriteOk() (*int32, bool)`

GetFreebandwidthwriteOk returns a tuple with the Freebandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreebandwidthwrite

`func (o *Media) SetFreebandwidthwrite(v int32)`

SetFreebandwidthwrite sets Freebandwidthwrite field to given value.

### HasFreebandwidthwrite

`func (o *Media) HasFreebandwidthwrite() bool`

HasFreebandwidthwrite returns a boolean if a field has been set.

### GetVolumescount

`func (o *Media) GetVolumescount() int32`

GetVolumescount returns the Volumescount field if non-nil, zero value otherwise.

### GetVolumescountOk

`func (o *Media) GetVolumescountOk() (*int32, bool)`

GetVolumescountOk returns a tuple with the Volumescount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumescount

`func (o *Media) SetVolumescount(v int32)`

SetVolumescount sets Volumescount field to given value.

### HasVolumescount

`func (o *Media) HasVolumescount() bool`

HasVolumescount returns a boolean if a field has been set.

### GetAssignment

`func (o *Media) GetAssignment() string`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *Media) GetAssignmentOk() (*string, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *Media) SetAssignment(v string)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *Media) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetState

`func (o *Media) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Media) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Media) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Media) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Media) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Media) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Media) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Media) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Media) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Media) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Media) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Media) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCapacitygroup

`func (o *Media) GetCapacitygroup() string`

GetCapacitygroup returns the Capacitygroup field if non-nil, zero value otherwise.

### GetCapacitygroupOk

`func (o *Media) GetCapacitygroupOk() (*string, bool)`

GetCapacitygroupOk returns a tuple with the Capacitygroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitygroup

`func (o *Media) SetCapacitygroup(v string)`

SetCapacitygroup sets Capacitygroup field to given value.

### HasCapacitygroup

`func (o *Media) HasCapacitygroup() bool`

HasCapacitygroup returns a boolean if a field has been set.

### GetLbaformats

`func (o *Media) GetLbaformats() []string`

GetLbaformats returns the Lbaformats field if non-nil, zero value otherwise.

### GetLbaformatsOk

`func (o *Media) GetLbaformatsOk() (*[]string, bool)`

GetLbaformatsOk returns a tuple with the Lbaformats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbaformats

`func (o *Media) SetLbaformats(v []string)`

SetLbaformats sets Lbaformats field to given value.

### HasLbaformats

`func (o *Media) HasLbaformats() bool`

HasLbaformats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


