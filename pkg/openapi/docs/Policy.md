# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the policy. The name can be any non-empty string that does not contain a white space. | 
**Iopswrite** | Pointer to **int32** | Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Iopsread** | Pointer to **int32** | Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Bandwidthwrite** | Pointer to **int32** | Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Bandwidthread** | Pointer to **int32** | Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Latencywrite** | Pointer to **int32** | Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Latencyread** | Pointer to **int32** | Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**Latencyreadcold** | Pointer to **int32** |  If not all the reads are hot (i.e., Percentage of Cold Reads is &gt;0) – Enter the more relaxed constraints for read latencies of cold data.  Valid values: non-negative integer number, that is larger than “Read Latency”. | [optional] [readonly] 
**Colddata** | Pointer to **int32** | Enter the percentage of the volume’s capacity that is expected to be “cold” (i.e. expected to only have infrequent reads). Default is 0%. Values that are greater than 0 give Volumez the option to use more economic media with more relaxed read performance requirements. Valid values: Integers in the range of 0..100. | [optional] [readonly] 
**Localzoneread** | Pointer to **bool** | Setting this value to “Yes” directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to “Yes” if you have network constraints (bandwidth or cost) across your zones; otherwise set it to “No”. | [optional] 
**Failureperformance** | Pointer to **bool** | Setting this value to “Yes” directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to “No” directs Volumez to provision volumes according to the desired performance, however in a case of failure – performance may be impacted. The default value is “No”. | [optional] 
**Capacityoptimization** | **string** | Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes. | 
**Capacityreservation** | Pointer to **int32** | Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%. | [optional] 
**Resiliencymedia** | Pointer to **int32** |  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2. | [optional] 
**Resiliencynode** | Pointer to **int32** | Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1. | [optional] 
**Resiliencyzone** | Pointer to **int32** | Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1. | [optional] 
**Encryption** | Pointer to **bool** | Enter “YES” to encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO. | [optional] 
**Sed** | Pointer to **bool** | Enter “YES” to direct Volumez to only use media with disk encryption capabilities. Note that specifying “NO” can still use such media, however it is not a must to use it. Default value: NO. | [optional] 
**Integrity** | Pointer to **bool** | Enter “YES” to direct Volumez to activate the “Device Mapper integrity” protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO. | [optional] 
**Snapshotkeep** | Pointer to **int32** |  | [optional] 
**Snapshotfrequency** | Pointer to **string** |  | [optional] 
**Snapshotday** | Pointer to **int32** |  | [optional] 
**Snapshothour** | Pointer to **int32** |  | [optional] 
**Snapshotminute** | Pointer to **int32** |  | [optional] 
**CreatedbyuserName** | Pointer to **string** |  | [optional] 
**Createdbyuseremail** | Pointer to **string** |  | [optional] 
**Createdtime** | Pointer to **string** |  | [optional] 
**Updatebyusername** | Pointer to **string** |  | [optional] 
**UpdatebyUseremail** | Pointer to **string** |  | [optional] 
**Updatetime** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy(name string, capacityoptimization string, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.


### GetIopswrite

`func (o *Policy) GetIopswrite() int32`

GetIopswrite returns the Iopswrite field if non-nil, zero value otherwise.

### GetIopswriteOk

`func (o *Policy) GetIopswriteOk() (*int32, bool)`

GetIopswriteOk returns a tuple with the Iopswrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopswrite

`func (o *Policy) SetIopswrite(v int32)`

SetIopswrite sets Iopswrite field to given value.

### HasIopswrite

`func (o *Policy) HasIopswrite() bool`

HasIopswrite returns a boolean if a field has been set.

### GetIopsread

`func (o *Policy) GetIopsread() int32`

GetIopsread returns the Iopsread field if non-nil, zero value otherwise.

### GetIopsreadOk

`func (o *Policy) GetIopsreadOk() (*int32, bool)`

GetIopsreadOk returns a tuple with the Iopsread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsread

`func (o *Policy) SetIopsread(v int32)`

SetIopsread sets Iopsread field to given value.

### HasIopsread

`func (o *Policy) HasIopsread() bool`

HasIopsread returns a boolean if a field has been set.

### GetBandwidthwrite

`func (o *Policy) GetBandwidthwrite() int32`

GetBandwidthwrite returns the Bandwidthwrite field if non-nil, zero value otherwise.

### GetBandwidthwriteOk

`func (o *Policy) GetBandwidthwriteOk() (*int32, bool)`

GetBandwidthwriteOk returns a tuple with the Bandwidthwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthwrite

`func (o *Policy) SetBandwidthwrite(v int32)`

SetBandwidthwrite sets Bandwidthwrite field to given value.

### HasBandwidthwrite

`func (o *Policy) HasBandwidthwrite() bool`

HasBandwidthwrite returns a boolean if a field has been set.

### GetBandwidthread

`func (o *Policy) GetBandwidthread() int32`

GetBandwidthread returns the Bandwidthread field if non-nil, zero value otherwise.

### GetBandwidthreadOk

`func (o *Policy) GetBandwidthreadOk() (*int32, bool)`

GetBandwidthreadOk returns a tuple with the Bandwidthread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthread

`func (o *Policy) SetBandwidthread(v int32)`

SetBandwidthread sets Bandwidthread field to given value.

### HasBandwidthread

`func (o *Policy) HasBandwidthread() bool`

HasBandwidthread returns a boolean if a field has been set.

### GetLatencywrite

`func (o *Policy) GetLatencywrite() int32`

GetLatencywrite returns the Latencywrite field if non-nil, zero value otherwise.

### GetLatencywriteOk

`func (o *Policy) GetLatencywriteOk() (*int32, bool)`

GetLatencywriteOk returns a tuple with the Latencywrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencywrite

`func (o *Policy) SetLatencywrite(v int32)`

SetLatencywrite sets Latencywrite field to given value.

### HasLatencywrite

`func (o *Policy) HasLatencywrite() bool`

HasLatencywrite returns a boolean if a field has been set.

### GetLatencyread

`func (o *Policy) GetLatencyread() int32`

GetLatencyread returns the Latencyread field if non-nil, zero value otherwise.

### GetLatencyreadOk

`func (o *Policy) GetLatencyreadOk() (*int32, bool)`

GetLatencyreadOk returns a tuple with the Latencyread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyread

`func (o *Policy) SetLatencyread(v int32)`

SetLatencyread sets Latencyread field to given value.

### HasLatencyread

`func (o *Policy) HasLatencyread() bool`

HasLatencyread returns a boolean if a field has been set.

### GetLatencyreadcold

`func (o *Policy) GetLatencyreadcold() int32`

GetLatencyreadcold returns the Latencyreadcold field if non-nil, zero value otherwise.

### GetLatencyreadcoldOk

`func (o *Policy) GetLatencyreadcoldOk() (*int32, bool)`

GetLatencyreadcoldOk returns a tuple with the Latencyreadcold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyreadcold

`func (o *Policy) SetLatencyreadcold(v int32)`

SetLatencyreadcold sets Latencyreadcold field to given value.

### HasLatencyreadcold

`func (o *Policy) HasLatencyreadcold() bool`

HasLatencyreadcold returns a boolean if a field has been set.

### GetColddata

`func (o *Policy) GetColddata() int32`

GetColddata returns the Colddata field if non-nil, zero value otherwise.

### GetColddataOk

`func (o *Policy) GetColddataOk() (*int32, bool)`

GetColddataOk returns a tuple with the Colddata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColddata

`func (o *Policy) SetColddata(v int32)`

SetColddata sets Colddata field to given value.

### HasColddata

`func (o *Policy) HasColddata() bool`

HasColddata returns a boolean if a field has been set.

### GetLocalzoneread

`func (o *Policy) GetLocalzoneread() bool`

GetLocalzoneread returns the Localzoneread field if non-nil, zero value otherwise.

### GetLocalzonereadOk

`func (o *Policy) GetLocalzonereadOk() (*bool, bool)`

GetLocalzonereadOk returns a tuple with the Localzoneread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalzoneread

`func (o *Policy) SetLocalzoneread(v bool)`

SetLocalzoneread sets Localzoneread field to given value.

### HasLocalzoneread

`func (o *Policy) HasLocalzoneread() bool`

HasLocalzoneread returns a boolean if a field has been set.

### GetFailureperformance

`func (o *Policy) GetFailureperformance() bool`

GetFailureperformance returns the Failureperformance field if non-nil, zero value otherwise.

### GetFailureperformanceOk

`func (o *Policy) GetFailureperformanceOk() (*bool, bool)`

GetFailureperformanceOk returns a tuple with the Failureperformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureperformance

`func (o *Policy) SetFailureperformance(v bool)`

SetFailureperformance sets Failureperformance field to given value.

### HasFailureperformance

`func (o *Policy) HasFailureperformance() bool`

HasFailureperformance returns a boolean if a field has been set.

### GetCapacityoptimization

`func (o *Policy) GetCapacityoptimization() string`

GetCapacityoptimization returns the Capacityoptimization field if non-nil, zero value otherwise.

### GetCapacityoptimizationOk

`func (o *Policy) GetCapacityoptimizationOk() (*string, bool)`

GetCapacityoptimizationOk returns a tuple with the Capacityoptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityoptimization

`func (o *Policy) SetCapacityoptimization(v string)`

SetCapacityoptimization sets Capacityoptimization field to given value.


### GetCapacityreservation

`func (o *Policy) GetCapacityreservation() int32`

GetCapacityreservation returns the Capacityreservation field if non-nil, zero value otherwise.

### GetCapacityreservationOk

`func (o *Policy) GetCapacityreservationOk() (*int32, bool)`

GetCapacityreservationOk returns a tuple with the Capacityreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityreservation

`func (o *Policy) SetCapacityreservation(v int32)`

SetCapacityreservation sets Capacityreservation field to given value.

### HasCapacityreservation

`func (o *Policy) HasCapacityreservation() bool`

HasCapacityreservation returns a boolean if a field has been set.

### GetResiliencymedia

`func (o *Policy) GetResiliencymedia() int32`

GetResiliencymedia returns the Resiliencymedia field if non-nil, zero value otherwise.

### GetResiliencymediaOk

`func (o *Policy) GetResiliencymediaOk() (*int32, bool)`

GetResiliencymediaOk returns a tuple with the Resiliencymedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencymedia

`func (o *Policy) SetResiliencymedia(v int32)`

SetResiliencymedia sets Resiliencymedia field to given value.

### HasResiliencymedia

`func (o *Policy) HasResiliencymedia() bool`

HasResiliencymedia returns a boolean if a field has been set.

### GetResiliencynode

`func (o *Policy) GetResiliencynode() int32`

GetResiliencynode returns the Resiliencynode field if non-nil, zero value otherwise.

### GetResiliencynodeOk

`func (o *Policy) GetResiliencynodeOk() (*int32, bool)`

GetResiliencynodeOk returns a tuple with the Resiliencynode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencynode

`func (o *Policy) SetResiliencynode(v int32)`

SetResiliencynode sets Resiliencynode field to given value.

### HasResiliencynode

`func (o *Policy) HasResiliencynode() bool`

HasResiliencynode returns a boolean if a field has been set.

### GetResiliencyzone

`func (o *Policy) GetResiliencyzone() int32`

GetResiliencyzone returns the Resiliencyzone field if non-nil, zero value otherwise.

### GetResiliencyzoneOk

`func (o *Policy) GetResiliencyzoneOk() (*int32, bool)`

GetResiliencyzoneOk returns a tuple with the Resiliencyzone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyzone

`func (o *Policy) SetResiliencyzone(v int32)`

SetResiliencyzone sets Resiliencyzone field to given value.

### HasResiliencyzone

`func (o *Policy) HasResiliencyzone() bool`

HasResiliencyzone returns a boolean if a field has been set.

### GetEncryption

`func (o *Policy) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Policy) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Policy) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Policy) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetSed

`func (o *Policy) GetSed() bool`

GetSed returns the Sed field if non-nil, zero value otherwise.

### GetSedOk

`func (o *Policy) GetSedOk() (*bool, bool)`

GetSedOk returns a tuple with the Sed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSed

`func (o *Policy) SetSed(v bool)`

SetSed sets Sed field to given value.

### HasSed

`func (o *Policy) HasSed() bool`

HasSed returns a boolean if a field has been set.

### GetIntegrity

`func (o *Policy) GetIntegrity() bool`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *Policy) GetIntegrityOk() (*bool, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *Policy) SetIntegrity(v bool)`

SetIntegrity sets Integrity field to given value.

### HasIntegrity

`func (o *Policy) HasIntegrity() bool`

HasIntegrity returns a boolean if a field has been set.

### GetSnapshotkeep

`func (o *Policy) GetSnapshotkeep() int32`

GetSnapshotkeep returns the Snapshotkeep field if non-nil, zero value otherwise.

### GetSnapshotkeepOk

`func (o *Policy) GetSnapshotkeepOk() (*int32, bool)`

GetSnapshotkeepOk returns a tuple with the Snapshotkeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotkeep

`func (o *Policy) SetSnapshotkeep(v int32)`

SetSnapshotkeep sets Snapshotkeep field to given value.

### HasSnapshotkeep

`func (o *Policy) HasSnapshotkeep() bool`

HasSnapshotkeep returns a boolean if a field has been set.

### GetSnapshotfrequency

`func (o *Policy) GetSnapshotfrequency() string`

GetSnapshotfrequency returns the Snapshotfrequency field if non-nil, zero value otherwise.

### GetSnapshotfrequencyOk

`func (o *Policy) GetSnapshotfrequencyOk() (*string, bool)`

GetSnapshotfrequencyOk returns a tuple with the Snapshotfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotfrequency

`func (o *Policy) SetSnapshotfrequency(v string)`

SetSnapshotfrequency sets Snapshotfrequency field to given value.

### HasSnapshotfrequency

`func (o *Policy) HasSnapshotfrequency() bool`

HasSnapshotfrequency returns a boolean if a field has been set.

### GetSnapshotday

`func (o *Policy) GetSnapshotday() int32`

GetSnapshotday returns the Snapshotday field if non-nil, zero value otherwise.

### GetSnapshotdayOk

`func (o *Policy) GetSnapshotdayOk() (*int32, bool)`

GetSnapshotdayOk returns a tuple with the Snapshotday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotday

`func (o *Policy) SetSnapshotday(v int32)`

SetSnapshotday sets Snapshotday field to given value.

### HasSnapshotday

`func (o *Policy) HasSnapshotday() bool`

HasSnapshotday returns a boolean if a field has been set.

### GetSnapshothour

`func (o *Policy) GetSnapshothour() int32`

GetSnapshothour returns the Snapshothour field if non-nil, zero value otherwise.

### GetSnapshothourOk

`func (o *Policy) GetSnapshothourOk() (*int32, bool)`

GetSnapshothourOk returns a tuple with the Snapshothour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshothour

`func (o *Policy) SetSnapshothour(v int32)`

SetSnapshothour sets Snapshothour field to given value.

### HasSnapshothour

`func (o *Policy) HasSnapshothour() bool`

HasSnapshothour returns a boolean if a field has been set.

### GetSnapshotminute

`func (o *Policy) GetSnapshotminute() int32`

GetSnapshotminute returns the Snapshotminute field if non-nil, zero value otherwise.

### GetSnapshotminuteOk

`func (o *Policy) GetSnapshotminuteOk() (*int32, bool)`

GetSnapshotminuteOk returns a tuple with the Snapshotminute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotminute

`func (o *Policy) SetSnapshotminute(v int32)`

SetSnapshotminute sets Snapshotminute field to given value.

### HasSnapshotminute

`func (o *Policy) HasSnapshotminute() bool`

HasSnapshotminute returns a boolean if a field has been set.

### GetCreatedbyuserName

`func (o *Policy) GetCreatedbyuserName() string`

GetCreatedbyuserName returns the CreatedbyuserName field if non-nil, zero value otherwise.

### GetCreatedbyuserNameOk

`func (o *Policy) GetCreatedbyuserNameOk() (*string, bool)`

GetCreatedbyuserNameOk returns a tuple with the CreatedbyuserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedbyuserName

`func (o *Policy) SetCreatedbyuserName(v string)`

SetCreatedbyuserName sets CreatedbyuserName field to given value.

### HasCreatedbyuserName

`func (o *Policy) HasCreatedbyuserName() bool`

HasCreatedbyuserName returns a boolean if a field has been set.

### GetCreatedbyuseremail

`func (o *Policy) GetCreatedbyuseremail() string`

GetCreatedbyuseremail returns the Createdbyuseremail field if non-nil, zero value otherwise.

### GetCreatedbyuseremailOk

`func (o *Policy) GetCreatedbyuseremailOk() (*string, bool)`

GetCreatedbyuseremailOk returns a tuple with the Createdbyuseremail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedbyuseremail

`func (o *Policy) SetCreatedbyuseremail(v string)`

SetCreatedbyuseremail sets Createdbyuseremail field to given value.

### HasCreatedbyuseremail

`func (o *Policy) HasCreatedbyuseremail() bool`

HasCreatedbyuseremail returns a boolean if a field has been set.

### GetCreatedtime

`func (o *Policy) GetCreatedtime() string`

GetCreatedtime returns the Createdtime field if non-nil, zero value otherwise.

### GetCreatedtimeOk

`func (o *Policy) GetCreatedtimeOk() (*string, bool)`

GetCreatedtimeOk returns a tuple with the Createdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedtime

`func (o *Policy) SetCreatedtime(v string)`

SetCreatedtime sets Createdtime field to given value.

### HasCreatedtime

`func (o *Policy) HasCreatedtime() bool`

HasCreatedtime returns a boolean if a field has been set.

### GetUpdatebyusername

`func (o *Policy) GetUpdatebyusername() string`

GetUpdatebyusername returns the Updatebyusername field if non-nil, zero value otherwise.

### GetUpdatebyusernameOk

`func (o *Policy) GetUpdatebyusernameOk() (*string, bool)`

GetUpdatebyusernameOk returns a tuple with the Updatebyusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatebyusername

`func (o *Policy) SetUpdatebyusername(v string)`

SetUpdatebyusername sets Updatebyusername field to given value.

### HasUpdatebyusername

`func (o *Policy) HasUpdatebyusername() bool`

HasUpdatebyusername returns a boolean if a field has been set.

### GetUpdatebyUseremail

`func (o *Policy) GetUpdatebyUseremail() string`

GetUpdatebyUseremail returns the UpdatebyUseremail field if non-nil, zero value otherwise.

### GetUpdatebyUseremailOk

`func (o *Policy) GetUpdatebyUseremailOk() (*string, bool)`

GetUpdatebyUseremailOk returns a tuple with the UpdatebyUseremail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatebyUseremail

`func (o *Policy) SetUpdatebyUseremail(v string)`

SetUpdatebyUseremail sets UpdatebyUseremail field to given value.

### HasUpdatebyUseremail

`func (o *Policy) HasUpdatebyUseremail() bool`

HasUpdatebyUseremail returns a boolean if a field has been set.

### GetUpdatetime

`func (o *Policy) GetUpdatetime() string`

GetUpdatetime returns the Updatetime field if non-nil, zero value otherwise.

### GetUpdatetimeOk

`func (o *Policy) GetUpdatetimeOk() (*string, bool)`

GetUpdatetimeOk returns a tuple with the Updatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatetime

`func (o *Policy) SetUpdatetime(v string)`

SetUpdatetime sets Updatetime field to given value.

### HasUpdatetime

`func (o *Policy) HasUpdatetime() bool`

HasUpdatetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


