//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Message encoding a command with parameters as scaled integers and additional metadata. Scaling depends on the actual command value.
type MessageCommandIntStamped struct {
	// UTC time, seconds elapsed since 01.01.1970
	UtcTime uint32
	// Microseconds elapsed since vehicle boot
	VehicleTimestamp uint64
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// The coordinate system of the COMMAND, as defined by MAV_FRAME enum
	Frame MAV_FRAME `mavenum:"uint8"`
	// The scheduled action for the mission item, as defined by MAV_CMD enum
	Command MAV_CMD `mavenum:"uint16"`
	// false:0, true:1
	Current uint8
	// autocontinue to next wp
	Autocontinue uint8
	// PARAM1, see MAV_CMD enum
	Param1 float32
	// PARAM2, see MAV_CMD enum
	Param2 float32
	// PARAM3, see MAV_CMD enum
	Param3 float32
	// PARAM4, see MAV_CMD enum
	Param4 float32
	// PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
	X int32
	// PARAM6 / local: y position in meters * 1e4, global: longitude in degrees * 10^7
	Y int32
	// PARAM7 / z position: global: altitude in meters (MSL, WGS84, AGL or relative to home - depending on frame).
	Z float32
}

// GetID implements the message.Message interface.
func (*MessageCommandIntStamped) GetID() uint32 {
	return 223
}
