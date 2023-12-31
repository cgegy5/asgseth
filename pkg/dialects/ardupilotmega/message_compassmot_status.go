//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Status of compassmot calibration.
type MessageCompassmotStatus struct {
	// Throttle.
	Throttle uint16
	// Current.
	Current float32
	// Interference.
	Interference uint16
	// Motor Compensation X.
	Compensationx float32 `mavname:"CompensationX"`
	// Motor Compensation Y.
	Compensationy float32 `mavname:"CompensationY"`
	// Motor Compensation Z.
	Compensationz float32 `mavname:"CompensationZ"`
}

// GetID implements the message.Message interface.
func (*MessageCompassmotStatus) GetID() uint32 {
	return 177
}
