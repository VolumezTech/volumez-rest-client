// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Policy policy
//
// swagger:model Policy
type Policy struct {

	// Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	BandwidthRead int64 `json:"bandwidthread"`

	// Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	BandwidthWrite int64 `json:"bandwidthwrite"`

	// Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.
	// Required: true
	// Enum: ["capacity","balanced","performance"]
	CapacityOptimization string `json:"capacityoptimization"`

	// Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%.
	// Example: 20
	// Maximum: 500
	// Minimum: 0
	CapacityReservation int64 `json:"capacityreservation"`

	// Enter the percentage of the volume’s capacity that is expected to be “cold” (i.e. expected to only have infrequent reads). Default is 0%. Values that are greater than 0 give Volumez the option to use more economic media with more relaxed read performance requirements. Valid values: Integers in the range of 0..100.
	// Read Only: true
	// Maximum: 100
	// Minimum: 0
	ColdData int64 `json:"colddata"`

	// created by user email
	CreatedByUserEmail string `json:"createdbyuseremail"`

	// created by user name
	CreatedByUserName string `json:"createdbyuserName"`

	// created time
	CreatedTime DateTime `json:"createdtime"`

	// Setting this value to “Yes” directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to “No” directs Volumez to provision volumes according to the desired performance, however in a case of failure – performance may be impacted. The default value is “No”.
	FailurePerformance bool `json:"failureperformance"`

	// Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Example: 1000
	// Minimum: 0
	IOPSRead int64 `json:"iopsread"`

	// Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Example: 1000
	// Minimum: 0
	IOPSWrite int64 `json:"iopswrite"`

	// Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	LatencyRead int64 `json:"latencyread"`

	//  If not all the reads are hot (i.e., Percentage of Cold Reads is >0) – Enter the more relaxed constraints for read latencies of cold data.  Valid values: non-negative integer number, that is larger than “Read Latency”.
	// Read Only: true
	// Minimum: 0
	LatencyReadCold int64 `json:"latencyreadcold"`

	// Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	// Minimum: 0
	LatencyWrite int64 `json:"latencywrite"`

	// Setting this value to “Yes” directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to “Yes” if you have network constraints (bandwidth or cost) across your zones; otherwise set it to “No”.
	LocalZoneRead bool `json:"localzoneread"`

	// Specifies the maximum bandwidth that Volumez is allowed to consume for replication of this volume (MB/s). 0 means no bandwidth limitation.
	ReplicationBandwidth int64 `json:"replicationbandwidth"`

	// Enter how many seconds are allowed for the replica to stay behind the primary storage. 0 means synchronous replication. Valid values are 0..3600, default value is 0. Max value: 3600. (1 hour).
	// Maximum: 3600
	// Minimum: 0
	ReplicationRPO int64 `json:"replicationrpo"`

	//  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2.
	// Example: 2
	// Maximum: 2
	// Minimum: 0
	ResiliencyMedia int64 `json:"resiliencymedia"`

	// Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1.
	// Example: 1
	// Maximum: 2
	// Minimum: 0
	ResiliencyNode int64 `json:"resiliencynode"`

	// Enter how many regions (e.g. AWS regions zones, DataCenters across continents) failures the system is required to sustain, and still serve volumes of this policy. Note: regions are assumed to reside across WAN distance, with some bandwidth limitations. Valid values are 0 and 1, default value is 0.
	// Example: 1
	// Maximum: 1
	// Minimum: 0
	ResiliencyRegion int64 `json:"resiliencyregion"`

	// Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1.
	// Example: 1
	// Minimum: 0
	ResiliencyZone int64 `json:"resiliencyzone"`

	// snapshot day
	SnapshotDay int64 `json:"snapshotday"`

	// snapshot frequency
	SnapshotFrequency string `json:"snapshotfrequency"`

	// snapshot hour
	SnapshotHour int64 `json:"snapshothour"`

	// snapshot keep
	SnapshotKeep int64 `json:"snapshotkeep"`

	// snapshot minute
	SnapshotMinute int64 `json:"snapshotminute"`

	// update by user email
	UpdateByUserEmail string `json:"updatebyUseremail,omitempty"`

	// update by user name
	UpdateByUserName string `json:"updatebyusername,omitempty"`

	// update time
	UpdateTime DateTime `json:"updatetime,omitempty"`

	// Enter “YES” to encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO.
	Encryption bool `json:"encryption"`

	// Enter “YES” to direct Volumez to activate the “Device Mapper integrity” protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO.
	Integrity bool `json:"integrity"`

	// A name for the policy. The name can be any non-empty string that does not contain a white space.
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// Enter “YES” to direct Volumez to only use media with disk encryption capabilities. Note that specifying “NO” can still use such media, however it is not a must to use it. Default value: NO.
	Sed bool `json:"sed"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidthRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthWrite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityOptimization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityReservation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColdData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIOPSRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIOPSWrite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyReadCold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyWrite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationRPO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) validateBandwidthRead(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthRead) { // not required
		return nil
	}

	if err := validate.MinimumInt("bandwidthread", "body", m.BandwidthRead, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateBandwidthWrite(formats strfmt.Registry) error {
	if swag.IsZero(m.BandwidthWrite) { // not required
		return nil
	}

	if err := validate.MinimumInt("bandwidthwrite", "body", m.BandwidthWrite, 0, false); err != nil {
		return err
	}

	return nil
}

var policyTypeCapacityOptimizationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["capacity","balanced","performance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyTypeCapacityOptimizationPropEnum = append(policyTypeCapacityOptimizationPropEnum, v)
	}
}

const (

	// PolicyCapacityOptimizationCapacity captures enum value "capacity"
	PolicyCapacityOptimizationCapacity string = "capacity"

	// PolicyCapacityOptimizationBalanced captures enum value "balanced"
	PolicyCapacityOptimizationBalanced string = "balanced"

	// PolicyCapacityOptimizationPerformance captures enum value "performance"
	PolicyCapacityOptimizationPerformance string = "performance"
)

// prop value enum
func (m *Policy) validateCapacityOptimizationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyTypeCapacityOptimizationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Policy) validateCapacityOptimization(formats strfmt.Registry) error {

	if err := validate.RequiredString("capacityoptimization", "body", m.CapacityOptimization); err != nil {
		return err
	}

	// value enum
	if err := m.validateCapacityOptimizationEnum("capacityoptimization", "body", m.CapacityOptimization); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateCapacityReservation(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityReservation) { // not required
		return nil
	}

	if err := validate.MinimumInt("capacityreservation", "body", m.CapacityReservation, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("capacityreservation", "body", m.CapacityReservation, 500, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateColdData(formats strfmt.Registry) error {
	if swag.IsZero(m.ColdData) { // not required
		return nil
	}

	if err := validate.MinimumInt("colddata", "body", m.ColdData, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("colddata", "body", m.ColdData, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateCreatedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := m.CreatedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("createdtime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("createdtime")
		}
		return err
	}

	return nil
}

func (m *Policy) validateIOPSRead(formats strfmt.Registry) error {
	if swag.IsZero(m.IOPSRead) { // not required
		return nil
	}

	if err := validate.MinimumInt("iopsread", "body", m.IOPSRead, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateIOPSWrite(formats strfmt.Registry) error {
	if swag.IsZero(m.IOPSWrite) { // not required
		return nil
	}

	if err := validate.MinimumInt("iopswrite", "body", m.IOPSWrite, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateLatencyRead(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyRead) { // not required
		return nil
	}

	if err := validate.MinimumInt("latencyread", "body", m.LatencyRead, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateLatencyReadCold(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyReadCold) { // not required
		return nil
	}

	if err := validate.MinimumInt("latencyreadcold", "body", m.LatencyReadCold, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateLatencyWrite(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyWrite) { // not required
		return nil
	}

	if err := validate.MinimumInt("latencywrite", "body", m.LatencyWrite, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateReplicationRPO(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationRPO) { // not required
		return nil
	}

	if err := validate.MinimumInt("replicationrpo", "body", m.ReplicationRPO, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("replicationrpo", "body", m.ReplicationRPO, 3600, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateResiliencyMedia(formats strfmt.Registry) error {
	if swag.IsZero(m.ResiliencyMedia) { // not required
		return nil
	}

	if err := validate.MinimumInt("resiliencymedia", "body", m.ResiliencyMedia, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("resiliencymedia", "body", m.ResiliencyMedia, 2, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateResiliencyNode(formats strfmt.Registry) error {
	if swag.IsZero(m.ResiliencyNode) { // not required
		return nil
	}

	if err := validate.MinimumInt("resiliencynode", "body", m.ResiliencyNode, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("resiliencynode", "body", m.ResiliencyNode, 2, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateResiliencyRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.ResiliencyRegion) { // not required
		return nil
	}

	if err := validate.MinimumInt("resiliencyregion", "body", m.ResiliencyRegion, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("resiliencyregion", "body", m.ResiliencyRegion, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateResiliencyZone(formats strfmt.Registry) error {
	if swag.IsZero(m.ResiliencyZone) { // not required
		return nil
	}

	if err := validate.MinimumInt("resiliencyzone", "body", m.ResiliencyZone, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := m.UpdateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("updatetime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("updatetime")
		}
		return err
	}

	return nil
}

func (m *Policy) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy based on the context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColdData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatencyReadCold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) contextValidateColdData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "colddata", "body", int64(m.ColdData)); err != nil {
		return err
	}

	return nil
}

func (m *Policy) contextValidateLatencyReadCold(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "latencyreadcold", "body", int64(m.LatencyReadCold)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
