//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Vehicle status report that is sent out while figure eight execution is in progress (see MAV_CMD_DO_FIGURE_EIGHT).
// This may typically send at low rates: of the order of 2Hz.
type MessageFigureEightExecutionStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Major axis radius of the figure eight. Positive: orbit the north circle clockwise. Negative: orbit the north circle counter-clockwise.
	MajorRadius float32
	// Minor axis radius of the figure eight. Defines the radius of two circles that make up the figure.
	MinorRadius float32
	// Orientation of the figure eight major axis with respect to true north in [-pi,pi).
	Orientation float32
	// The coordinate system of the fields: x, y, z.
	Frame MAV_FRAME `mavenum:"uint8"`
	// X coordinate of center point. Coordinate system depends on frame field.
	X int32
	// Y coordinate of center point. Coordinate system depends on frame field.
	Y int32
	// Altitude of center point. Coordinate system depends on frame field.
	Z float32
}

// GetID implements the message.Message interface.
func (*MessageFigureEightExecutionStatus) GetID() uint32 {
	return 361
}
