# CreateInfraPlanRequestPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IopsWrite** | Pointer to **int32** | Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**IopsRead** | Pointer to **int32** | Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**BandwidthWrite** | Pointer to **int32** | Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**BandwidthRead** | Pointer to **int32** | Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**LatencyWrite** | Pointer to **int32** | Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**LatencyRead** | Pointer to **int32** | Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes. | [optional] 
**LatencyReadCold** | Pointer to **int32** |  If not all the reads are hot (i.e., Percentage of Cold Reads is &gt;0) Enter the more relaxed constraints for read latencies of cold data.  Valid values: non-negative integer number, that is larger than “Read Latency”. | [optional] 
**ColdData** | Pointer to **int32** | Enter the percentage of the volumes capacity that is expected to be cold (i.e. expected to only have infrequent reads). Default is 0%. Values that are greater than 0 give Volumez the option to use more economic media with more relaxed read performance requirements. Valid values: Integers in the range of 0..100. | [optional] 
**LocalZoneRead** | Pointer to **bool** | Setting this value to Yes directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to Yes if you have network constraints (bandwidth or cost) across your zones; otherwise set it to No | [optional] 
**FailurePerformance** | Pointer to **bool** | Setting this value to Yes directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to No directs Volumez to provision volumes according to the desired performance, however in a case of failure performance may be impacted. The default value is No. | [optional] 
**CapacityOptimization** | Pointer to **string** | Choosing Capacity directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing Balanced directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes. | [optional] 
**CapacityReservation** | Pointer to **int32** | Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%. | [optional] 
**ResiliencyMedia** | Pointer to **int32** |  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2. | [optional] 
**ResiliencyNode** | Pointer to **int32** | Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1. | [optional] 
**ResiliencyZone** | Pointer to **int32** | Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1. | [optional] 
**ResiliencyRegion** | Pointer to **int32** | Enter how many regions (e.g. AWS regions zones, DataCenters across continents) failures the system is required to sustain, and still serve volumes of this policy. Note: regions are assumed to reside across WAN distance, with some bandwidth limitations. Valid values are 0 and 1, default value is 0. | [optional] 
**ReplicationRpo** | Pointer to **int32** | Enter how many seconds are allowed for the replica to stay behind the primary storage. 0 means synchronous replication. Valid values are 0..3600, default value is 0. Max value: 3600. (1 hour). | [optional] 
**ReplicationBandwidth** | Pointer to **int32** | Specifies the maximum bandwidth that Volumez is allowed to consume for replication of this volume (MB/s). 0 means no bandwidth limitation. | [optional] 
**Encryption** | Pointer to **bool** | Enter YESto encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO. | [optional] 
**Sed** | Pointer to **bool** | Enter YES to direct Volumez to only use media with disk encryption capabilities. Note that specifying NO can still use such media, however it is not a must to use it. Default value: NO. | [optional] 
**Integrity** | Pointer to **bool** | Enter YES to direct Volumez to activate the Device Mapper integrity protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO. | [optional] 
**SnapshotKeep** | Pointer to **int32** |  | [optional] 
**SnapshotFrequency** | Pointer to **string** |  | [optional] 
**SnapshotDay** | Pointer to **int32** |  | [optional] 
**SnapshotHour** | Pointer to **int32** |  | [optional] 
**SnapshotMinute** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateInfraPlanRequestPolicySettings

`func NewCreateInfraPlanRequestPolicySettings() *CreateInfraPlanRequestPolicySettings`

NewCreateInfraPlanRequestPolicySettings instantiates a new CreateInfraPlanRequestPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraPlanRequestPolicySettingsWithDefaults

`func NewCreateInfraPlanRequestPolicySettingsWithDefaults() *CreateInfraPlanRequestPolicySettings`

NewCreateInfraPlanRequestPolicySettingsWithDefaults instantiates a new CreateInfraPlanRequestPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIopsWrite

`func (o *CreateInfraPlanRequestPolicySettings) GetIopsWrite() int32`

GetIopsWrite returns the IopsWrite field if non-nil, zero value otherwise.

### GetIopsWriteOk

`func (o *CreateInfraPlanRequestPolicySettings) GetIopsWriteOk() (*int32, bool)`

GetIopsWriteOk returns a tuple with the IopsWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrite

`func (o *CreateInfraPlanRequestPolicySettings) SetIopsWrite(v int32)`

SetIopsWrite sets IopsWrite field to given value.

### HasIopsWrite

`func (o *CreateInfraPlanRequestPolicySettings) HasIopsWrite() bool`

HasIopsWrite returns a boolean if a field has been set.

### GetIopsRead

`func (o *CreateInfraPlanRequestPolicySettings) GetIopsRead() int32`

GetIopsRead returns the IopsRead field if non-nil, zero value otherwise.

### GetIopsReadOk

`func (o *CreateInfraPlanRequestPolicySettings) GetIopsReadOk() (*int32, bool)`

GetIopsReadOk returns a tuple with the IopsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRead

`func (o *CreateInfraPlanRequestPolicySettings) SetIopsRead(v int32)`

SetIopsRead sets IopsRead field to given value.

### HasIopsRead

`func (o *CreateInfraPlanRequestPolicySettings) HasIopsRead() bool`

HasIopsRead returns a boolean if a field has been set.

### GetBandwidthWrite

`func (o *CreateInfraPlanRequestPolicySettings) GetBandwidthWrite() int32`

GetBandwidthWrite returns the BandwidthWrite field if non-nil, zero value otherwise.

### GetBandwidthWriteOk

`func (o *CreateInfraPlanRequestPolicySettings) GetBandwidthWriteOk() (*int32, bool)`

GetBandwidthWriteOk returns a tuple with the BandwidthWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthWrite

`func (o *CreateInfraPlanRequestPolicySettings) SetBandwidthWrite(v int32)`

SetBandwidthWrite sets BandwidthWrite field to given value.

### HasBandwidthWrite

`func (o *CreateInfraPlanRequestPolicySettings) HasBandwidthWrite() bool`

HasBandwidthWrite returns a boolean if a field has been set.

### GetBandwidthRead

`func (o *CreateInfraPlanRequestPolicySettings) GetBandwidthRead() int32`

GetBandwidthRead returns the BandwidthRead field if non-nil, zero value otherwise.

### GetBandwidthReadOk

`func (o *CreateInfraPlanRequestPolicySettings) GetBandwidthReadOk() (*int32, bool)`

GetBandwidthReadOk returns a tuple with the BandwidthRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthRead

`func (o *CreateInfraPlanRequestPolicySettings) SetBandwidthRead(v int32)`

SetBandwidthRead sets BandwidthRead field to given value.

### HasBandwidthRead

`func (o *CreateInfraPlanRequestPolicySettings) HasBandwidthRead() bool`

HasBandwidthRead returns a boolean if a field has been set.

### GetLatencyWrite

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyWrite() int32`

GetLatencyWrite returns the LatencyWrite field if non-nil, zero value otherwise.

### GetLatencyWriteOk

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyWriteOk() (*int32, bool)`

GetLatencyWriteOk returns a tuple with the LatencyWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyWrite

`func (o *CreateInfraPlanRequestPolicySettings) SetLatencyWrite(v int32)`

SetLatencyWrite sets LatencyWrite field to given value.

### HasLatencyWrite

`func (o *CreateInfraPlanRequestPolicySettings) HasLatencyWrite() bool`

HasLatencyWrite returns a boolean if a field has been set.

### GetLatencyRead

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyRead() int32`

GetLatencyRead returns the LatencyRead field if non-nil, zero value otherwise.

### GetLatencyReadOk

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyReadOk() (*int32, bool)`

GetLatencyReadOk returns a tuple with the LatencyRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyRead

`func (o *CreateInfraPlanRequestPolicySettings) SetLatencyRead(v int32)`

SetLatencyRead sets LatencyRead field to given value.

### HasLatencyRead

`func (o *CreateInfraPlanRequestPolicySettings) HasLatencyRead() bool`

HasLatencyRead returns a boolean if a field has been set.

### GetLatencyReadCold

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyReadCold() int32`

GetLatencyReadCold returns the LatencyReadCold field if non-nil, zero value otherwise.

### GetLatencyReadColdOk

`func (o *CreateInfraPlanRequestPolicySettings) GetLatencyReadColdOk() (*int32, bool)`

GetLatencyReadColdOk returns a tuple with the LatencyReadCold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyReadCold

`func (o *CreateInfraPlanRequestPolicySettings) SetLatencyReadCold(v int32)`

SetLatencyReadCold sets LatencyReadCold field to given value.

### HasLatencyReadCold

`func (o *CreateInfraPlanRequestPolicySettings) HasLatencyReadCold() bool`

HasLatencyReadCold returns a boolean if a field has been set.

### GetColdData

`func (o *CreateInfraPlanRequestPolicySettings) GetColdData() int32`

GetColdData returns the ColdData field if non-nil, zero value otherwise.

### GetColdDataOk

`func (o *CreateInfraPlanRequestPolicySettings) GetColdDataOk() (*int32, bool)`

GetColdDataOk returns a tuple with the ColdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdData

`func (o *CreateInfraPlanRequestPolicySettings) SetColdData(v int32)`

SetColdData sets ColdData field to given value.

### HasColdData

`func (o *CreateInfraPlanRequestPolicySettings) HasColdData() bool`

HasColdData returns a boolean if a field has been set.

### GetLocalZoneRead

`func (o *CreateInfraPlanRequestPolicySettings) GetLocalZoneRead() bool`

GetLocalZoneRead returns the LocalZoneRead field if non-nil, zero value otherwise.

### GetLocalZoneReadOk

`func (o *CreateInfraPlanRequestPolicySettings) GetLocalZoneReadOk() (*bool, bool)`

GetLocalZoneReadOk returns a tuple with the LocalZoneRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalZoneRead

`func (o *CreateInfraPlanRequestPolicySettings) SetLocalZoneRead(v bool)`

SetLocalZoneRead sets LocalZoneRead field to given value.

### HasLocalZoneRead

`func (o *CreateInfraPlanRequestPolicySettings) HasLocalZoneRead() bool`

HasLocalZoneRead returns a boolean if a field has been set.

### GetFailurePerformance

`func (o *CreateInfraPlanRequestPolicySettings) GetFailurePerformance() bool`

GetFailurePerformance returns the FailurePerformance field if non-nil, zero value otherwise.

### GetFailurePerformanceOk

`func (o *CreateInfraPlanRequestPolicySettings) GetFailurePerformanceOk() (*bool, bool)`

GetFailurePerformanceOk returns a tuple with the FailurePerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePerformance

`func (o *CreateInfraPlanRequestPolicySettings) SetFailurePerformance(v bool)`

SetFailurePerformance sets FailurePerformance field to given value.

### HasFailurePerformance

`func (o *CreateInfraPlanRequestPolicySettings) HasFailurePerformance() bool`

HasFailurePerformance returns a boolean if a field has been set.

### GetCapacityOptimization

`func (o *CreateInfraPlanRequestPolicySettings) GetCapacityOptimization() string`

GetCapacityOptimization returns the CapacityOptimization field if non-nil, zero value otherwise.

### GetCapacityOptimizationOk

`func (o *CreateInfraPlanRequestPolicySettings) GetCapacityOptimizationOk() (*string, bool)`

GetCapacityOptimizationOk returns a tuple with the CapacityOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityOptimization

`func (o *CreateInfraPlanRequestPolicySettings) SetCapacityOptimization(v string)`

SetCapacityOptimization sets CapacityOptimization field to given value.

### HasCapacityOptimization

`func (o *CreateInfraPlanRequestPolicySettings) HasCapacityOptimization() bool`

HasCapacityOptimization returns a boolean if a field has been set.

### GetCapacityReservation

`func (o *CreateInfraPlanRequestPolicySettings) GetCapacityReservation() int32`

GetCapacityReservation returns the CapacityReservation field if non-nil, zero value otherwise.

### GetCapacityReservationOk

`func (o *CreateInfraPlanRequestPolicySettings) GetCapacityReservationOk() (*int32, bool)`

GetCapacityReservationOk returns a tuple with the CapacityReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityReservation

`func (o *CreateInfraPlanRequestPolicySettings) SetCapacityReservation(v int32)`

SetCapacityReservation sets CapacityReservation field to given value.

### HasCapacityReservation

`func (o *CreateInfraPlanRequestPolicySettings) HasCapacityReservation() bool`

HasCapacityReservation returns a boolean if a field has been set.

### GetResiliencyMedia

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyMedia() int32`

GetResiliencyMedia returns the ResiliencyMedia field if non-nil, zero value otherwise.

### GetResiliencyMediaOk

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyMediaOk() (*int32, bool)`

GetResiliencyMediaOk returns a tuple with the ResiliencyMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyMedia

`func (o *CreateInfraPlanRequestPolicySettings) SetResiliencyMedia(v int32)`

SetResiliencyMedia sets ResiliencyMedia field to given value.

### HasResiliencyMedia

`func (o *CreateInfraPlanRequestPolicySettings) HasResiliencyMedia() bool`

HasResiliencyMedia returns a boolean if a field has been set.

### GetResiliencyNode

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyNode() int32`

GetResiliencyNode returns the ResiliencyNode field if non-nil, zero value otherwise.

### GetResiliencyNodeOk

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyNodeOk() (*int32, bool)`

GetResiliencyNodeOk returns a tuple with the ResiliencyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyNode

`func (o *CreateInfraPlanRequestPolicySettings) SetResiliencyNode(v int32)`

SetResiliencyNode sets ResiliencyNode field to given value.

### HasResiliencyNode

`func (o *CreateInfraPlanRequestPolicySettings) HasResiliencyNode() bool`

HasResiliencyNode returns a boolean if a field has been set.

### GetResiliencyZone

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyZone() int32`

GetResiliencyZone returns the ResiliencyZone field if non-nil, zero value otherwise.

### GetResiliencyZoneOk

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyZoneOk() (*int32, bool)`

GetResiliencyZoneOk returns a tuple with the ResiliencyZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyZone

`func (o *CreateInfraPlanRequestPolicySettings) SetResiliencyZone(v int32)`

SetResiliencyZone sets ResiliencyZone field to given value.

### HasResiliencyZone

`func (o *CreateInfraPlanRequestPolicySettings) HasResiliencyZone() bool`

HasResiliencyZone returns a boolean if a field has been set.

### GetResiliencyRegion

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyRegion() int32`

GetResiliencyRegion returns the ResiliencyRegion field if non-nil, zero value otherwise.

### GetResiliencyRegionOk

`func (o *CreateInfraPlanRequestPolicySettings) GetResiliencyRegionOk() (*int32, bool)`

GetResiliencyRegionOk returns a tuple with the ResiliencyRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyRegion

`func (o *CreateInfraPlanRequestPolicySettings) SetResiliencyRegion(v int32)`

SetResiliencyRegion sets ResiliencyRegion field to given value.

### HasResiliencyRegion

`func (o *CreateInfraPlanRequestPolicySettings) HasResiliencyRegion() bool`

HasResiliencyRegion returns a boolean if a field has been set.

### GetReplicationRpo

`func (o *CreateInfraPlanRequestPolicySettings) GetReplicationRpo() int32`

GetReplicationRpo returns the ReplicationRpo field if non-nil, zero value otherwise.

### GetReplicationRpoOk

`func (o *CreateInfraPlanRequestPolicySettings) GetReplicationRpoOk() (*int32, bool)`

GetReplicationRpoOk returns a tuple with the ReplicationRpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRpo

`func (o *CreateInfraPlanRequestPolicySettings) SetReplicationRpo(v int32)`

SetReplicationRpo sets ReplicationRpo field to given value.

### HasReplicationRpo

`func (o *CreateInfraPlanRequestPolicySettings) HasReplicationRpo() bool`

HasReplicationRpo returns a boolean if a field has been set.

### GetReplicationBandwidth

`func (o *CreateInfraPlanRequestPolicySettings) GetReplicationBandwidth() int32`

GetReplicationBandwidth returns the ReplicationBandwidth field if non-nil, zero value otherwise.

### GetReplicationBandwidthOk

`func (o *CreateInfraPlanRequestPolicySettings) GetReplicationBandwidthOk() (*int32, bool)`

GetReplicationBandwidthOk returns a tuple with the ReplicationBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBandwidth

`func (o *CreateInfraPlanRequestPolicySettings) SetReplicationBandwidth(v int32)`

SetReplicationBandwidth sets ReplicationBandwidth field to given value.

### HasReplicationBandwidth

`func (o *CreateInfraPlanRequestPolicySettings) HasReplicationBandwidth() bool`

HasReplicationBandwidth returns a boolean if a field has been set.

### GetEncryption

`func (o *CreateInfraPlanRequestPolicySettings) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CreateInfraPlanRequestPolicySettings) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CreateInfraPlanRequestPolicySettings) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CreateInfraPlanRequestPolicySettings) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetSed

`func (o *CreateInfraPlanRequestPolicySettings) GetSed() bool`

GetSed returns the Sed field if non-nil, zero value otherwise.

### GetSedOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSedOk() (*bool, bool)`

GetSedOk returns a tuple with the Sed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSed

`func (o *CreateInfraPlanRequestPolicySettings) SetSed(v bool)`

SetSed sets Sed field to given value.

### HasSed

`func (o *CreateInfraPlanRequestPolicySettings) HasSed() bool`

HasSed returns a boolean if a field has been set.

### GetIntegrity

`func (o *CreateInfraPlanRequestPolicySettings) GetIntegrity() bool`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *CreateInfraPlanRequestPolicySettings) GetIntegrityOk() (*bool, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *CreateInfraPlanRequestPolicySettings) SetIntegrity(v bool)`

SetIntegrity sets Integrity field to given value.

### HasIntegrity

`func (o *CreateInfraPlanRequestPolicySettings) HasIntegrity() bool`

HasIntegrity returns a boolean if a field has been set.

### GetSnapshotKeep

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotKeep() int32`

GetSnapshotKeep returns the SnapshotKeep field if non-nil, zero value otherwise.

### GetSnapshotKeepOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotKeepOk() (*int32, bool)`

GetSnapshotKeepOk returns a tuple with the SnapshotKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotKeep

`func (o *CreateInfraPlanRequestPolicySettings) SetSnapshotKeep(v int32)`

SetSnapshotKeep sets SnapshotKeep field to given value.

### HasSnapshotKeep

`func (o *CreateInfraPlanRequestPolicySettings) HasSnapshotKeep() bool`

HasSnapshotKeep returns a boolean if a field has been set.

### GetSnapshotFrequency

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotFrequency() string`

GetSnapshotFrequency returns the SnapshotFrequency field if non-nil, zero value otherwise.

### GetSnapshotFrequencyOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotFrequencyOk() (*string, bool)`

GetSnapshotFrequencyOk returns a tuple with the SnapshotFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFrequency

`func (o *CreateInfraPlanRequestPolicySettings) SetSnapshotFrequency(v string)`

SetSnapshotFrequency sets SnapshotFrequency field to given value.

### HasSnapshotFrequency

`func (o *CreateInfraPlanRequestPolicySettings) HasSnapshotFrequency() bool`

HasSnapshotFrequency returns a boolean if a field has been set.

### GetSnapshotDay

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotDay() int32`

GetSnapshotDay returns the SnapshotDay field if non-nil, zero value otherwise.

### GetSnapshotDayOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotDayOk() (*int32, bool)`

GetSnapshotDayOk returns a tuple with the SnapshotDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDay

`func (o *CreateInfraPlanRequestPolicySettings) SetSnapshotDay(v int32)`

SetSnapshotDay sets SnapshotDay field to given value.

### HasSnapshotDay

`func (o *CreateInfraPlanRequestPolicySettings) HasSnapshotDay() bool`

HasSnapshotDay returns a boolean if a field has been set.

### GetSnapshotHour

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotHour() int32`

GetSnapshotHour returns the SnapshotHour field if non-nil, zero value otherwise.

### GetSnapshotHourOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotHourOk() (*int32, bool)`

GetSnapshotHourOk returns a tuple with the SnapshotHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotHour

`func (o *CreateInfraPlanRequestPolicySettings) SetSnapshotHour(v int32)`

SetSnapshotHour sets SnapshotHour field to given value.

### HasSnapshotHour

`func (o *CreateInfraPlanRequestPolicySettings) HasSnapshotHour() bool`

HasSnapshotHour returns a boolean if a field has been set.

### GetSnapshotMinute

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotMinute() int32`

GetSnapshotMinute returns the SnapshotMinute field if non-nil, zero value otherwise.

### GetSnapshotMinuteOk

`func (o *CreateInfraPlanRequestPolicySettings) GetSnapshotMinuteOk() (*int32, bool)`

GetSnapshotMinuteOk returns a tuple with the SnapshotMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotMinute

`func (o *CreateInfraPlanRequestPolicySettings) SetSnapshotMinute(v int32)`

SetSnapshotMinute sets SnapshotMinute field to given value.

### HasSnapshotMinute

`func (o *CreateInfraPlanRequestPolicySettings) HasSnapshotMinute() bool`

HasSnapshotMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


