//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// A fence point. Used to set a point when from GCS -> MAV. Also used to return a point from MAV -> GCS.
type MessageFencePoint struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Point index (first point is 1, 0 is for return point).
	Idx uint8
	// Total number of points (for sanity checking).
	Count uint8
	// Latitude of point.
	Lat float32
	// Longitude of point.
	Lng float32
}

// GetID implements the message.Message interface.
func (*MessageFencePoint) GetID() uint32 {
	return 160
}
