# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instanceid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Os** | **string** |  | 
**Osversion** | Pointer to **string** |  | [optional] 
**Kversion** | Pointer to **string** |  | [optional] 
**Controladdress** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to **string** |  | [optional] 
**AccountID** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**ResourceNamespace** | Pointer to **string** | global namespace for resources in account empty if not aviliable/supported on cloud provider/node | [optional] 
**PhysicalProximityGroup** | Pointer to **string** | identifier of the physical location of the node empty if not aviliable/supported on on cloud provider/node | [optional] 
**ResiliencyDomain** | Pointer to **string** | virtual domain for the node fault domains if aviliable/supported on on cloud provider/node | [optional] 
**FaultDomain** | Pointer to **string** | identifier for node in FaultDomain | [optional] 
**Offlinetime** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Progress** | Pointer to **int32** |  | [optional] [readonly] 
**Connectorversion** | Pointer to **string** |  | [optional] [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Cloudprovider** | Pointer to **string** |  | [optional] 
**Nodecluster** | Pointer to **string** |  | [optional] 
**AutoprovisionInfraUUID** | Pointer to **string** |  | [optional] [readonly] 
**Instancetype** | Pointer to **string** |  | [optional] 

## Methods

### NewNode

`func NewNode(name string, os string, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceid

`func (o *Node) GetInstanceid() string`

GetInstanceid returns the Instanceid field if non-nil, zero value otherwise.

### GetInstanceidOk

`func (o *Node) GetInstanceidOk() (*string, bool)`

GetInstanceidOk returns a tuple with the Instanceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceid

`func (o *Node) SetInstanceid(v string)`

SetInstanceid sets Instanceid field to given value.

### HasInstanceid

`func (o *Node) HasInstanceid() bool`

HasInstanceid returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.


### GetOs

`func (o *Node) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Node) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Node) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsversion

`func (o *Node) GetOsversion() string`

GetOsversion returns the Osversion field if non-nil, zero value otherwise.

### GetOsversionOk

`func (o *Node) GetOsversionOk() (*string, bool)`

GetOsversionOk returns a tuple with the Osversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsversion

`func (o *Node) SetOsversion(v string)`

SetOsversion sets Osversion field to given value.

### HasOsversion

`func (o *Node) HasOsversion() bool`

HasOsversion returns a boolean if a field has been set.

### GetKversion

`func (o *Node) GetKversion() string`

GetKversion returns the Kversion field if non-nil, zero value otherwise.

### GetKversionOk

`func (o *Node) GetKversionOk() (*string, bool)`

GetKversionOk returns a tuple with the Kversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKversion

`func (o *Node) SetKversion(v string)`

SetKversion sets Kversion field to given value.

### HasKversion

`func (o *Node) HasKversion() bool`

HasKversion returns a boolean if a field has been set.

### GetControladdress

`func (o *Node) GetControladdress() string`

GetControladdress returns the Controladdress field if non-nil, zero value otherwise.

### GetControladdressOk

`func (o *Node) GetControladdressOk() (*string, bool)`

GetControladdressOk returns a tuple with the Controladdress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControladdress

`func (o *Node) SetControladdress(v string)`

SetControladdress sets Controladdress field to given value.

### HasControladdress

`func (o *Node) HasControladdress() bool`

HasControladdress returns a boolean if a field has been set.

### GetCredential

`func (o *Node) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *Node) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *Node) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *Node) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetAccountID

`func (o *Node) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Node) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Node) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Node) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetRegion

`func (o *Node) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Node) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Node) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Node) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZone

`func (o *Node) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Node) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Node) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Node) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResourceNamespace

`func (o *Node) GetResourceNamespace() string`

GetResourceNamespace returns the ResourceNamespace field if non-nil, zero value otherwise.

### GetResourceNamespaceOk

`func (o *Node) GetResourceNamespaceOk() (*string, bool)`

GetResourceNamespaceOk returns a tuple with the ResourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNamespace

`func (o *Node) SetResourceNamespace(v string)`

SetResourceNamespace sets ResourceNamespace field to given value.

### HasResourceNamespace

`func (o *Node) HasResourceNamespace() bool`

HasResourceNamespace returns a boolean if a field has been set.

### GetPhysicalProximityGroup

`func (o *Node) GetPhysicalProximityGroup() string`

GetPhysicalProximityGroup returns the PhysicalProximityGroup field if non-nil, zero value otherwise.

### GetPhysicalProximityGroupOk

`func (o *Node) GetPhysicalProximityGroupOk() (*string, bool)`

GetPhysicalProximityGroupOk returns a tuple with the PhysicalProximityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalProximityGroup

`func (o *Node) SetPhysicalProximityGroup(v string)`

SetPhysicalProximityGroup sets PhysicalProximityGroup field to given value.

### HasPhysicalProximityGroup

`func (o *Node) HasPhysicalProximityGroup() bool`

HasPhysicalProximityGroup returns a boolean if a field has been set.

### GetResiliencyDomain

`func (o *Node) GetResiliencyDomain() string`

GetResiliencyDomain returns the ResiliencyDomain field if non-nil, zero value otherwise.

### GetResiliencyDomainOk

`func (o *Node) GetResiliencyDomainOk() (*string, bool)`

GetResiliencyDomainOk returns a tuple with the ResiliencyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDomain

`func (o *Node) SetResiliencyDomain(v string)`

SetResiliencyDomain sets ResiliencyDomain field to given value.

### HasResiliencyDomain

`func (o *Node) HasResiliencyDomain() bool`

HasResiliencyDomain returns a boolean if a field has been set.

### GetFaultDomain

`func (o *Node) GetFaultDomain() string`

GetFaultDomain returns the FaultDomain field if non-nil, zero value otherwise.

### GetFaultDomainOk

`func (o *Node) GetFaultDomainOk() (*string, bool)`

GetFaultDomainOk returns a tuple with the FaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomain

`func (o *Node) SetFaultDomain(v string)`

SetFaultDomain sets FaultDomain field to given value.

### HasFaultDomain

`func (o *Node) HasFaultDomain() bool`

HasFaultDomain returns a boolean if a field has been set.

### GetOfflinetime

`func (o *Node) GetOfflinetime() string`

GetOfflinetime returns the Offlinetime field if non-nil, zero value otherwise.

### GetOfflinetimeOk

`func (o *Node) GetOfflinetimeOk() (*string, bool)`

GetOfflinetimeOk returns a tuple with the Offlinetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflinetime

`func (o *Node) SetOfflinetime(v string)`

SetOfflinetime sets Offlinetime field to given value.

### HasOfflinetime

`func (o *Node) HasOfflinetime() bool`

HasOfflinetime returns a boolean if a field has been set.

### GetState

`func (o *Node) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Node) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Node) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Node) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *Node) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Node) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Node) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Node) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *Node) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Node) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Node) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Node) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetConnectorversion

`func (o *Node) GetConnectorversion() string`

GetConnectorversion returns the Connectorversion field if non-nil, zero value otherwise.

### GetConnectorversionOk

`func (o *Node) GetConnectorversionOk() (*string, bool)`

GetConnectorversionOk returns a tuple with the Connectorversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorversion

`func (o *Node) SetConnectorversion(v string)`

SetConnectorversion sets Connectorversion field to given value.

### HasConnectorversion

`func (o *Node) HasConnectorversion() bool`

HasConnectorversion returns a boolean if a field has been set.

### GetLabel

`func (o *Node) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Node) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Node) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Node) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *Node) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Node) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Node) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Node) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCloudprovider

`func (o *Node) GetCloudprovider() string`

GetCloudprovider returns the Cloudprovider field if non-nil, zero value otherwise.

### GetCloudproviderOk

`func (o *Node) GetCloudproviderOk() (*string, bool)`

GetCloudproviderOk returns a tuple with the Cloudprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudprovider

`func (o *Node) SetCloudprovider(v string)`

SetCloudprovider sets Cloudprovider field to given value.

### HasCloudprovider

`func (o *Node) HasCloudprovider() bool`

HasCloudprovider returns a boolean if a field has been set.

### GetNodecluster

`func (o *Node) GetNodecluster() string`

GetNodecluster returns the Nodecluster field if non-nil, zero value otherwise.

### GetNodeclusterOk

`func (o *Node) GetNodeclusterOk() (*string, bool)`

GetNodeclusterOk returns a tuple with the Nodecluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodecluster

`func (o *Node) SetNodecluster(v string)`

SetNodecluster sets Nodecluster field to given value.

### HasNodecluster

`func (o *Node) HasNodecluster() bool`

HasNodecluster returns a boolean if a field has been set.

### GetAutoprovisionInfraUUID

`func (o *Node) GetAutoprovisionInfraUUID() string`

GetAutoprovisionInfraUUID returns the AutoprovisionInfraUUID field if non-nil, zero value otherwise.

### GetAutoprovisionInfraUUIDOk

`func (o *Node) GetAutoprovisionInfraUUIDOk() (*string, bool)`

GetAutoprovisionInfraUUIDOk returns a tuple with the AutoprovisionInfraUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoprovisionInfraUUID

`func (o *Node) SetAutoprovisionInfraUUID(v string)`

SetAutoprovisionInfraUUID sets AutoprovisionInfraUUID field to given value.

### HasAutoprovisionInfraUUID

`func (o *Node) HasAutoprovisionInfraUUID() bool`

HasAutoprovisionInfraUUID returns a boolean if a field has been set.

### GetInstancetype

`func (o *Node) GetInstancetype() string`

GetInstancetype returns the Instancetype field if non-nil, zero value otherwise.

### GetInstancetypeOk

`func (o *Node) GetInstancetypeOk() (*string, bool)`

GetInstancetypeOk returns a tuple with the Instancetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancetype

`func (o *Node) SetInstancetype(v string)`

SetInstancetype sets Instancetype field to given value.

### HasInstancetype

`func (o *Node) HasInstancetype() bool`

HasInstancetype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


