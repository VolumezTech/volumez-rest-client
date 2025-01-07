# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Zone** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Ipstart** | **string** |  | 
**Ipend** | **string** |  | 

## Methods

### NewNetwork

`func NewNetwork(name string, type_ string, ipstart string, ipend string, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.


### GetZone

`func (o *Network) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Network) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Network) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Network) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *Network) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Network) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Network) SetType(v string)`

SetType sets Type field to given value.


### GetIpstart

`func (o *Network) GetIpstart() string`

GetIpstart returns the Ipstart field if non-nil, zero value otherwise.

### GetIpstartOk

`func (o *Network) GetIpstartOk() (*string, bool)`

GetIpstartOk returns a tuple with the Ipstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpstart

`func (o *Network) SetIpstart(v string)`

SetIpstart sets Ipstart field to given value.


### GetIpend

`func (o *Network) GetIpend() string`

GetIpend returns the Ipend field if non-nil, zero value otherwise.

### GetIpendOk

`func (o *Network) GetIpendOk() (*string, bool)`

GetIpendOk returns a tuple with the Ipend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpend

`func (o *Network) SetIpend(v string)`

SetIpend sets Ipend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


