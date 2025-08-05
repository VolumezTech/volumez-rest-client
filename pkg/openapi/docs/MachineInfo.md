# MachineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instanceid** | Pointer to **string** |  | [optional] 
**Machinename** | Pointer to **string** |  | [optional] 
**Accountid** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Resourcenamespace** | Pointer to **string** |  | [optional] 
**PhysicalproximityGroup** | Pointer to **string** |  | [optional] 
**Resiliencydomain** | Pointer to **string** |  | [optional] 
**Faultdomain** | Pointer to **string** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**Ipaddresses** | Pointer to **[]string** |  | [optional] 
**Publicdns** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Additional** | Pointer to [**[]AdditionalParam**](AdditionalParam.md) |  | [optional] 

## Methods

### NewMachineInfo

`func NewMachineInfo() *MachineInfo`

NewMachineInfo instantiates a new MachineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineInfoWithDefaults

`func NewMachineInfoWithDefaults() *MachineInfo`

NewMachineInfoWithDefaults instantiates a new MachineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceid

`func (o *MachineInfo) GetInstanceid() string`

GetInstanceid returns the Instanceid field if non-nil, zero value otherwise.

### GetInstanceidOk

`func (o *MachineInfo) GetInstanceidOk() (*string, bool)`

GetInstanceidOk returns a tuple with the Instanceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceid

`func (o *MachineInfo) SetInstanceid(v string)`

SetInstanceid sets Instanceid field to given value.

### HasInstanceid

`func (o *MachineInfo) HasInstanceid() bool`

HasInstanceid returns a boolean if a field has been set.

### GetMachinename

`func (o *MachineInfo) GetMachinename() string`

GetMachinename returns the Machinename field if non-nil, zero value otherwise.

### GetMachinenameOk

`func (o *MachineInfo) GetMachinenameOk() (*string, bool)`

GetMachinenameOk returns a tuple with the Machinename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinename

`func (o *MachineInfo) SetMachinename(v string)`

SetMachinename sets Machinename field to given value.

### HasMachinename

`func (o *MachineInfo) HasMachinename() bool`

HasMachinename returns a boolean if a field has been set.

### GetAccountid

`func (o *MachineInfo) GetAccountid() string`

GetAccountid returns the Accountid field if non-nil, zero value otherwise.

### GetAccountidOk

`func (o *MachineInfo) GetAccountidOk() (*string, bool)`

GetAccountidOk returns a tuple with the Accountid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountid

`func (o *MachineInfo) SetAccountid(v string)`

SetAccountid sets Accountid field to given value.

### HasAccountid

`func (o *MachineInfo) HasAccountid() bool`

HasAccountid returns a boolean if a field has been set.

### GetRegion

`func (o *MachineInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MachineInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MachineInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MachineInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZone

`func (o *MachineInfo) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MachineInfo) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MachineInfo) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *MachineInfo) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResourcenamespace

`func (o *MachineInfo) GetResourcenamespace() string`

GetResourcenamespace returns the Resourcenamespace field if non-nil, zero value otherwise.

### GetResourcenamespaceOk

`func (o *MachineInfo) GetResourcenamespaceOk() (*string, bool)`

GetResourcenamespaceOk returns a tuple with the Resourcenamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcenamespace

`func (o *MachineInfo) SetResourcenamespace(v string)`

SetResourcenamespace sets Resourcenamespace field to given value.

### HasResourcenamespace

`func (o *MachineInfo) HasResourcenamespace() bool`

HasResourcenamespace returns a boolean if a field has been set.

### GetPhysicalproximityGroup

`func (o *MachineInfo) GetPhysicalproximityGroup() string`

GetPhysicalproximityGroup returns the PhysicalproximityGroup field if non-nil, zero value otherwise.

### GetPhysicalproximityGroupOk

`func (o *MachineInfo) GetPhysicalproximityGroupOk() (*string, bool)`

GetPhysicalproximityGroupOk returns a tuple with the PhysicalproximityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalproximityGroup

`func (o *MachineInfo) SetPhysicalproximityGroup(v string)`

SetPhysicalproximityGroup sets PhysicalproximityGroup field to given value.

### HasPhysicalproximityGroup

`func (o *MachineInfo) HasPhysicalproximityGroup() bool`

HasPhysicalproximityGroup returns a boolean if a field has been set.

### GetResiliencydomain

`func (o *MachineInfo) GetResiliencydomain() string`

GetResiliencydomain returns the Resiliencydomain field if non-nil, zero value otherwise.

### GetResiliencydomainOk

`func (o *MachineInfo) GetResiliencydomainOk() (*string, bool)`

GetResiliencydomainOk returns a tuple with the Resiliencydomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencydomain

`func (o *MachineInfo) SetResiliencydomain(v string)`

SetResiliencydomain sets Resiliencydomain field to given value.

### HasResiliencydomain

`func (o *MachineInfo) HasResiliencydomain() bool`

HasResiliencydomain returns a boolean if a field has been set.

### GetFaultdomain

`func (o *MachineInfo) GetFaultdomain() string`

GetFaultdomain returns the Faultdomain field if non-nil, zero value otherwise.

### GetFaultdomainOk

`func (o *MachineInfo) GetFaultdomainOk() (*string, bool)`

GetFaultdomainOk returns a tuple with the Faultdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultdomain

`func (o *MachineInfo) SetFaultdomain(v string)`

SetFaultdomain sets Faultdomain field to given value.

### HasFaultdomain

`func (o *MachineInfo) HasFaultdomain() bool`

HasFaultdomain returns a boolean if a field has been set.

### GetArchitecture

`func (o *MachineInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *MachineInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *MachineInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *MachineInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetIpaddresses

`func (o *MachineInfo) GetIpaddresses() []string`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *MachineInfo) GetIpaddressesOk() (*[]string, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *MachineInfo) SetIpaddresses(v []string)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *MachineInfo) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetPublicdns

`func (o *MachineInfo) GetPublicdns() string`

GetPublicdns returns the Publicdns field if non-nil, zero value otherwise.

### GetPublicdnsOk

`func (o *MachineInfo) GetPublicdnsOk() (*string, bool)`

GetPublicdnsOk returns a tuple with the Publicdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicdns

`func (o *MachineInfo) SetPublicdns(v string)`

SetPublicdns sets Publicdns field to given value.

### HasPublicdns

`func (o *MachineInfo) HasPublicdns() bool`

HasPublicdns returns a boolean if a field has been set.

### GetCluster

`func (o *MachineInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MachineInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MachineInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MachineInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetAdditional

`func (o *MachineInfo) GetAdditional() []AdditionalParam`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *MachineInfo) GetAdditionalOk() (*[]AdditionalParam, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *MachineInfo) SetAdditional(v []AdditionalParam)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *MachineInfo) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


