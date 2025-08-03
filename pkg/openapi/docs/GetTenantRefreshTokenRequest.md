# GetTenantRefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accesstoken** | **string** |  | 
**Hostname** | **string** |  | 
**Tenantid** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTenantRefreshTokenRequest

`func NewGetTenantRefreshTokenRequest(accesstoken string, hostname string, ) *GetTenantRefreshTokenRequest`

NewGetTenantRefreshTokenRequest instantiates a new GetTenantRefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantRefreshTokenRequestWithDefaults

`func NewGetTenantRefreshTokenRequestWithDefaults() *GetTenantRefreshTokenRequest`

NewGetTenantRefreshTokenRequestWithDefaults instantiates a new GetTenantRefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccesstoken

`func (o *GetTenantRefreshTokenRequest) GetAccesstoken() string`

GetAccesstoken returns the Accesstoken field if non-nil, zero value otherwise.

### GetAccesstokenOk

`func (o *GetTenantRefreshTokenRequest) GetAccesstokenOk() (*string, bool)`

GetAccesstokenOk returns a tuple with the Accesstoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesstoken

`func (o *GetTenantRefreshTokenRequest) SetAccesstoken(v string)`

SetAccesstoken sets Accesstoken field to given value.


### GetHostname

`func (o *GetTenantRefreshTokenRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetTenantRefreshTokenRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetTenantRefreshTokenRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetTenantid

`func (o *GetTenantRefreshTokenRequest) GetTenantid() string`

GetTenantid returns the Tenantid field if non-nil, zero value otherwise.

### GetTenantidOk

`func (o *GetTenantRefreshTokenRequest) GetTenantidOk() (*string, bool)`

GetTenantidOk returns a tuple with the Tenantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantid

`func (o *GetTenantRefreshTokenRequest) SetTenantid(v string)`

SetTenantid sets Tenantid field to given value.

### HasTenantid

`func (o *GetTenantRefreshTokenRequest) HasTenantid() bool`

HasTenantid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


