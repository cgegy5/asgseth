//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Information about a gimbal manager. This message should be requested by a ground station using MAV_CMD_REQUEST_MESSAGE. It mirrors some fields of the GIMBAL_DEVICE_INFORMATION message, but not all. If the additional information is desired, also GIMBAL_DEVICE_INFORMATION should be requested.
type MessageStorm32GimbalManagerInformation struct {
	// Gimbal ID (component ID or 1-6 for non-MAVLink gimbal) that this gimbal manager is responsible for.
	GimbalId uint8
	// Gimbal device capability flags. Same flags as reported by GIMBAL_DEVICE_INFORMATION. The flag is only 16 bit wide, but stored in 32 bit, for backwards compatibility (high word is zero).
	DeviceCapFlags GIMBAL_DEVICE_CAP_FLAGS `mavenum:"uint32"`
	// Gimbal manager capability flags.
	ManagerCapFlags MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS `mavenum:"uint32"`
	// Hardware minimum roll angle (positive: roll to the right). NaN if unknown.
	RollMin float32
	// Hardware maximum roll angle (positive: roll to the right). NaN if unknown.
	RollMax float32
	// Hardware minimum pitch/tilt angle (positive: tilt up). NaN if unknown.
	PitchMin float32
	// Hardware maximum pitch/tilt angle (positive: tilt up). NaN if unknown.
	PitchMax float32
	// Hardware minimum yaw/pan angle (positive: pan to the right, relative to the vehicle/gimbal base). NaN if unknown.
	YawMin float32
	// Hardware maximum yaw/pan angle (positive: pan to the right, relative to the vehicle/gimbal base). NaN if unknown.
	YawMax float32
}

// GetID implements the message.Message interface.
func (*MessageStorm32GimbalManagerInformation) GetID() uint32 {
	return 60010
}
