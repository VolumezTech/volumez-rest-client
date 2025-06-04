# VolumeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumegroupname** | Pointer to **string** |  | [optional] [readonly] 
**Encryption** | Pointer to **bool** |  | [optional] 
**Integrity** | Pointer to **bool** |  | [optional] 
**Allocation** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **bool** |  | [optional] 
**Deduplication** | Pointer to **bool** |  | [optional] 
**Writecache** | Pointer to **bool** |  | [optional] 
**Redundancy** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Targetsecret** | Pointer to **string** |  | [optional] 
**Allocatedsize** | Pointer to **int32** |  | [optional] 
**Resiliency** | Pointer to **string** |  | [optional] 
**Raidcolumns** | Pointer to **int32** |  | [optional] 
**Mediasize** | Pointer to **int32** |  | [optional] 
**Mediabandwidthwrite** | Pointer to **int32** |  | [optional] 
**Mediabandwidthread** | Pointer to **int32** |  | [optional] 
**Mediaiopswrite** | Pointer to **int32** |  | [optional] 
**Mediaiopsread** | Pointer to **int32** |  | [optional] 
**Media** | Pointer to **[]string** |  | [optional] 
**Cachesize** | Pointer to **int32** |  | [optional] 
**Cacheresiliency** | Pointer to **string** |  | [optional] 
**Cacheredundancy** | Pointer to **int32** |  | [optional] 
**Cacheraidcolumns** | Pointer to **int32** |  | [optional] 
**Cachemediasize** | Pointer to **int32** |  | [optional] 
**Cachemediabandwidthwrite** | Pointer to **int32** |  | [optional] 
**Cachemediabandwidthread** | Pointer to **int32** |  | [optional] 
**Cachemediaiopswrite** | Pointer to **int32** |  | [optional] 
**Cachemediaiopsread** | Pointer to **int32** |  | [optional] 
**Cachemedia** | Pointer to **[]string** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 

## Methods

### NewVolumeGroup

`func NewVolumeGroup() *VolumeGroup`

NewVolumeGroup instantiates a new VolumeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupWithDefaults

`func NewVolumeGroupWithDefaults() *VolumeGroup`

NewVolumeGroupWithDefaults instantiates a new VolumeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumegroupname

`func (o *VolumeGroup) GetVolumegroupname() string`

GetVolumegroupname returns the Volumegroupname field if non-nil, zero value otherwise.

### GetVolumegroupnameOk

`func (o *VolumeGroup) GetVolumegroupnameOk() (*string, bool)`

GetVolumegroupnameOk returns a tuple with the Volumegroupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumegroupname

`func (o *VolumeGroup) SetVolumegroupname(v string)`

SetVolumegroupname sets Volumegroupname field to given value.

### HasVolumegroupname

`func (o *VolumeGroup) HasVolumegroupname() bool`

HasVolumegroupname returns a boolean if a field has been set.

### GetEncryption

`func (o *VolumeGroup) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VolumeGroup) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VolumeGroup) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VolumeGroup) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetIntegrity

`func (o *VolumeGroup) GetIntegrity() bool`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *VolumeGroup) GetIntegrityOk() (*bool, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *VolumeGroup) SetIntegrity(v bool)`

SetIntegrity sets Integrity field to given value.

### HasIntegrity

`func (o *VolumeGroup) HasIntegrity() bool`

HasIntegrity returns a boolean if a field has been set.

### GetAllocation

`func (o *VolumeGroup) GetAllocation() string`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *VolumeGroup) GetAllocationOk() (*string, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *VolumeGroup) SetAllocation(v string)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *VolumeGroup) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### GetCompression

`func (o *VolumeGroup) GetCompression() bool`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *VolumeGroup) GetCompressionOk() (*bool, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *VolumeGroup) SetCompression(v bool)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *VolumeGroup) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetDeduplication

`func (o *VolumeGroup) GetDeduplication() bool`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *VolumeGroup) GetDeduplicationOk() (*bool, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *VolumeGroup) SetDeduplication(v bool)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *VolumeGroup) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetWritecache

`func (o *VolumeGroup) GetWritecache() bool`

GetWritecache returns the Writecache field if non-nil, zero value otherwise.

### GetWritecacheOk

`func (o *VolumeGroup) GetWritecacheOk() (*bool, bool)`

GetWritecacheOk returns a tuple with the Writecache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritecache

`func (o *VolumeGroup) SetWritecache(v bool)`

SetWritecache sets Writecache field to given value.

### HasWritecache

`func (o *VolumeGroup) HasWritecache() bool`

HasWritecache returns a boolean if a field has been set.

### GetRedundancy

`func (o *VolumeGroup) GetRedundancy() int32`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *VolumeGroup) GetRedundancyOk() (*int32, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *VolumeGroup) SetRedundancy(v int32)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *VolumeGroup) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetSize

`func (o *VolumeGroup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeGroup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeGroup) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeGroup) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTargetsecret

`func (o *VolumeGroup) GetTargetsecret() string`

GetTargetsecret returns the Targetsecret field if non-nil, zero value otherwise.

### GetTargetsecretOk

`func (o *VolumeGroup) GetTargetsecretOk() (*string, bool)`

GetTargetsecretOk returns a tuple with the Targetsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsecret

`func (o *VolumeGroup) SetTargetsecret(v string)`

SetTargetsecret sets Targetsecret field to given value.

### HasTargetsecret

`func (o *VolumeGroup) HasTargetsecret() bool`

HasTargetsecret returns a boolean if a field has been set.

### GetAllocatedsize

`func (o *VolumeGroup) GetAllocatedsize() int32`

GetAllocatedsize returns the Allocatedsize field if non-nil, zero value otherwise.

### GetAllocatedsizeOk

`func (o *VolumeGroup) GetAllocatedsizeOk() (*int32, bool)`

GetAllocatedsizeOk returns a tuple with the Allocatedsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedsize

`func (o *VolumeGroup) SetAllocatedsize(v int32)`

SetAllocatedsize sets Allocatedsize field to given value.

### HasAllocatedsize

`func (o *VolumeGroup) HasAllocatedsize() bool`

HasAllocatedsize returns a boolean if a field has been set.

### GetResiliency

`func (o *VolumeGroup) GetResiliency() string`

GetResiliency returns the Resiliency field if non-nil, zero value otherwise.

### GetResiliencyOk

`func (o *VolumeGroup) GetResiliencyOk() (*string, bool)`

GetResiliencyOk returns a tuple with the Resiliency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliency

`func (o *VolumeGroup) SetResiliency(v string)`

SetResiliency sets Resiliency field to given value.

### HasResiliency

`func (o *VolumeGroup) HasResiliency() bool`

HasResiliency returns a boolean if a field has been set.

### GetRaidcolumns

`func (o *VolumeGroup) GetRaidcolumns() int32`

GetRaidcolumns returns the Raidcolumns field if non-nil, zero value otherwise.

### GetRaidcolumnsOk

`func (o *VolumeGroup) GetRaidcolumnsOk() (*int32, bool)`

GetRaidcolumnsOk returns a tuple with the Raidcolumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidcolumns

`func (o *VolumeGroup) SetRaidcolumns(v int32)`

SetRaidcolumns sets Raidcolumns field to given value.

### HasRaidcolumns

`func (o *VolumeGroup) HasRaidcolumns() bool`

HasRaidcolumns returns a boolean if a field has been set.

### GetMediasize

`func (o *VolumeGroup) GetMediasize() int32`

GetMediasize returns the Mediasize field if non-nil, zero value otherwise.

### GetMediasizeOk

`func (o *VolumeGroup) GetMediasizeOk() (*int32, bool)`

GetMediasizeOk returns a tuple with the Mediasize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediasize

`func (o *VolumeGroup) SetMediasize(v int32)`

SetMediasize sets Mediasize field to given value.

### HasMediasize

`func (o *VolumeGroup) HasMediasize() bool`

HasMediasize returns a boolean if a field has been set.

### GetMediabandwidthwrite

`func (o *VolumeGroup) GetMediabandwidthwrite() int32`

GetMediabandwidthwrite returns the Mediabandwidthwrite field if non-nil, zero value otherwise.

### GetMediabandwidthwriteOk

`func (o *VolumeGroup) GetMediabandwidthwriteOk() (*int32, bool)`

GetMediabandwidthwriteOk returns a tuple with the Mediabandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediabandwidthwrite

`func (o *VolumeGroup) SetMediabandwidthwrite(v int32)`

SetMediabandwidthwrite sets Mediabandwidthwrite field to given value.

### HasMediabandwidthwrite

`func (o *VolumeGroup) HasMediabandwidthwrite() bool`

HasMediabandwidthwrite returns a boolean if a field has been set.

### GetMediabandwidthread

`func (o *VolumeGroup) GetMediabandwidthread() int32`

GetMediabandwidthread returns the Mediabandwidthread field if non-nil, zero value otherwise.

### GetMediabandwidthreadOk

`func (o *VolumeGroup) GetMediabandwidthreadOk() (*int32, bool)`

GetMediabandwidthreadOk returns a tuple with the Mediabandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediabandwidthread

`func (o *VolumeGroup) SetMediabandwidthread(v int32)`

SetMediabandwidthread sets Mediabandwidthread field to given value.

### HasMediabandwidthread

`func (o *VolumeGroup) HasMediabandwidthread() bool`

HasMediabandwidthread returns a boolean if a field has been set.

### GetMediaiopswrite

`func (o *VolumeGroup) GetMediaiopswrite() int32`

GetMediaiopswrite returns the Mediaiopswrite field if non-nil, zero value otherwise.

### GetMediaiopswriteOk

`func (o *VolumeGroup) GetMediaiopswriteOk() (*int32, bool)`

GetMediaiopswriteOk returns a tuple with the Mediaiopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaiopswrite

`func (o *VolumeGroup) SetMediaiopswrite(v int32)`

SetMediaiopswrite sets Mediaiopswrite field to given value.

### HasMediaiopswrite

`func (o *VolumeGroup) HasMediaiopswrite() bool`

HasMediaiopswrite returns a boolean if a field has been set.

### GetMediaiopsread

`func (o *VolumeGroup) GetMediaiopsread() int32`

GetMediaiopsread returns the Mediaiopsread field if non-nil, zero value otherwise.

### GetMediaiopsreadOk

`func (o *VolumeGroup) GetMediaiopsreadOk() (*int32, bool)`

GetMediaiopsreadOk returns a tuple with the Mediaiopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaiopsread

`func (o *VolumeGroup) SetMediaiopsread(v int32)`

SetMediaiopsread sets Mediaiopsread field to given value.

### HasMediaiopsread

`func (o *VolumeGroup) HasMediaiopsread() bool`

HasMediaiopsread returns a boolean if a field has been set.

### GetMedia

`func (o *VolumeGroup) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *VolumeGroup) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *VolumeGroup) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *VolumeGroup) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetCachesize

`func (o *VolumeGroup) GetCachesize() int32`

GetCachesize returns the Cachesize field if non-nil, zero value otherwise.

### GetCachesizeOk

`func (o *VolumeGroup) GetCachesizeOk() (*int32, bool)`

GetCachesizeOk returns a tuple with the Cachesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachesize

`func (o *VolumeGroup) SetCachesize(v int32)`

SetCachesize sets Cachesize field to given value.

### HasCachesize

`func (o *VolumeGroup) HasCachesize() bool`

HasCachesize returns a boolean if a field has been set.

### GetCacheresiliency

`func (o *VolumeGroup) GetCacheresiliency() string`

GetCacheresiliency returns the Cacheresiliency field if non-nil, zero value otherwise.

### GetCacheresiliencyOk

`func (o *VolumeGroup) GetCacheresiliencyOk() (*string, bool)`

GetCacheresiliencyOk returns a tuple with the Cacheresiliency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheresiliency

`func (o *VolumeGroup) SetCacheresiliency(v string)`

SetCacheresiliency sets Cacheresiliency field to given value.

### HasCacheresiliency

`func (o *VolumeGroup) HasCacheresiliency() bool`

HasCacheresiliency returns a boolean if a field has been set.

### GetCacheredundancy

`func (o *VolumeGroup) GetCacheredundancy() int32`

GetCacheredundancy returns the Cacheredundancy field if non-nil, zero value otherwise.

### GetCacheredundancyOk

`func (o *VolumeGroup) GetCacheredundancyOk() (*int32, bool)`

GetCacheredundancyOk returns a tuple with the Cacheredundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheredundancy

`func (o *VolumeGroup) SetCacheredundancy(v int32)`

SetCacheredundancy sets Cacheredundancy field to given value.

### HasCacheredundancy

`func (o *VolumeGroup) HasCacheredundancy() bool`

HasCacheredundancy returns a boolean if a field has been set.

### GetCacheraidcolumns

`func (o *VolumeGroup) GetCacheraidcolumns() int32`

GetCacheraidcolumns returns the Cacheraidcolumns field if non-nil, zero value otherwise.

### GetCacheraidcolumnsOk

`func (o *VolumeGroup) GetCacheraidcolumnsOk() (*int32, bool)`

GetCacheraidcolumnsOk returns a tuple with the Cacheraidcolumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheraidcolumns

`func (o *VolumeGroup) SetCacheraidcolumns(v int32)`

SetCacheraidcolumns sets Cacheraidcolumns field to given value.

### HasCacheraidcolumns

`func (o *VolumeGroup) HasCacheraidcolumns() bool`

HasCacheraidcolumns returns a boolean if a field has been set.

### GetCachemediasize

`func (o *VolumeGroup) GetCachemediasize() int32`

GetCachemediasize returns the Cachemediasize field if non-nil, zero value otherwise.

### GetCachemediasizeOk

`func (o *VolumeGroup) GetCachemediasizeOk() (*int32, bool)`

GetCachemediasizeOk returns a tuple with the Cachemediasize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemediasize

`func (o *VolumeGroup) SetCachemediasize(v int32)`

SetCachemediasize sets Cachemediasize field to given value.

### HasCachemediasize

`func (o *VolumeGroup) HasCachemediasize() bool`

HasCachemediasize returns a boolean if a field has been set.

### GetCachemediabandwidthwrite

`func (o *VolumeGroup) GetCachemediabandwidthwrite() int32`

GetCachemediabandwidthwrite returns the Cachemediabandwidthwrite field if non-nil, zero value otherwise.

### GetCachemediabandwidthwriteOk

`func (o *VolumeGroup) GetCachemediabandwidthwriteOk() (*int32, bool)`

GetCachemediabandwidthwriteOk returns a tuple with the Cachemediabandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemediabandwidthwrite

`func (o *VolumeGroup) SetCachemediabandwidthwrite(v int32)`

SetCachemediabandwidthwrite sets Cachemediabandwidthwrite field to given value.

### HasCachemediabandwidthwrite

`func (o *VolumeGroup) HasCachemediabandwidthwrite() bool`

HasCachemediabandwidthwrite returns a boolean if a field has been set.

### GetCachemediabandwidthread

`func (o *VolumeGroup) GetCachemediabandwidthread() int32`

GetCachemediabandwidthread returns the Cachemediabandwidthread field if non-nil, zero value otherwise.

### GetCachemediabandwidthreadOk

`func (o *VolumeGroup) GetCachemediabandwidthreadOk() (*int32, bool)`

GetCachemediabandwidthreadOk returns a tuple with the Cachemediabandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemediabandwidthread

`func (o *VolumeGroup) SetCachemediabandwidthread(v int32)`

SetCachemediabandwidthread sets Cachemediabandwidthread field to given value.

### HasCachemediabandwidthread

`func (o *VolumeGroup) HasCachemediabandwidthread() bool`

HasCachemediabandwidthread returns a boolean if a field has been set.

### GetCachemediaiopswrite

`func (o *VolumeGroup) GetCachemediaiopswrite() int32`

GetCachemediaiopswrite returns the Cachemediaiopswrite field if non-nil, zero value otherwise.

### GetCachemediaiopswriteOk

`func (o *VolumeGroup) GetCachemediaiopswriteOk() (*int32, bool)`

GetCachemediaiopswriteOk returns a tuple with the Cachemediaiopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemediaiopswrite

`func (o *VolumeGroup) SetCachemediaiopswrite(v int32)`

SetCachemediaiopswrite sets Cachemediaiopswrite field to given value.

### HasCachemediaiopswrite

`func (o *VolumeGroup) HasCachemediaiopswrite() bool`

HasCachemediaiopswrite returns a boolean if a field has been set.

### GetCachemediaiopsread

`func (o *VolumeGroup) GetCachemediaiopsread() int32`

GetCachemediaiopsread returns the Cachemediaiopsread field if non-nil, zero value otherwise.

### GetCachemediaiopsreadOk

`func (o *VolumeGroup) GetCachemediaiopsreadOk() (*int32, bool)`

GetCachemediaiopsreadOk returns a tuple with the Cachemediaiopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemediaiopsread

`func (o *VolumeGroup) SetCachemediaiopsread(v int32)`

SetCachemediaiopsread sets Cachemediaiopsread field to given value.

### HasCachemediaiopsread

`func (o *VolumeGroup) HasCachemediaiopsread() bool`

HasCachemediaiopsread returns a boolean if a field has been set.

### GetCachemedia

`func (o *VolumeGroup) GetCachemedia() []string`

GetCachemedia returns the Cachemedia field if non-nil, zero value otherwise.

### GetCachemediaOk

`func (o *VolumeGroup) GetCachemediaOk() (*[]string, bool)`

GetCachemediaOk returns a tuple with the Cachemedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachemedia

`func (o *VolumeGroup) SetCachemedia(v []string)`

SetCachemedia sets Cachemedia field to given value.

### HasCachemedia

`func (o *VolumeGroup) HasCachemedia() bool`

HasCachemedia returns a boolean if a field has been set.

### GetAttachments

`func (o *VolumeGroup) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeGroup) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeGroup) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *VolumeGroup) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


