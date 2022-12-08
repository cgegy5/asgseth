//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package pythonarraytest

import (
	"errors"
)

// Flags for high level gimbal manager operation The first 16 bits are identical to the GIMBAL_DEVICE_FLAGS.
type GIMBAL_MANAGER_FLAGS uint32

const (
	// Based on GIMBAL_DEVICE_FLAGS_RETRACT.
	GIMBAL_MANAGER_FLAGS_RETRACT GIMBAL_MANAGER_FLAGS = 1
	// Based on GIMBAL_DEVICE_FLAGS_NEUTRAL.
	GIMBAL_MANAGER_FLAGS_NEUTRAL GIMBAL_MANAGER_FLAGS = 2
	// Based on GIMBAL_DEVICE_FLAGS_ROLL_LOCK.
	GIMBAL_MANAGER_FLAGS_ROLL_LOCK GIMBAL_MANAGER_FLAGS = 4
	// Based on GIMBAL_DEVICE_FLAGS_PITCH_LOCK.
	GIMBAL_MANAGER_FLAGS_PITCH_LOCK GIMBAL_MANAGER_FLAGS = 8
	// Based on GIMBAL_DEVICE_FLAGS_YAW_LOCK.
	GIMBAL_MANAGER_FLAGS_YAW_LOCK GIMBAL_MANAGER_FLAGS = 16
	// Based on GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME.
	GIMBAL_MANAGER_FLAGS_YAW_IN_VEHICLE_FRAME GIMBAL_MANAGER_FLAGS = 32
	// Based on GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_FLAGS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_FLAGS = 64
	// Based on GIMBAL_DEVICE_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_FLAGS_ACCEPTS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_FLAGS = 128
	// Based on GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE.
	GIMBAL_MANAGER_FLAGS_RC_EXCLUSIVE GIMBAL_MANAGER_FLAGS = 256
	// Based on GIMBAL_DEVICE_FLAGS_RC_MIXED.
	GIMBAL_MANAGER_FLAGS_RC_MIXED GIMBAL_MANAGER_FLAGS = 512
)

var labels_GIMBAL_MANAGER_FLAGS = map[GIMBAL_MANAGER_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_MANAGER_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_MANAGER_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_MANAGER_FLAGS = map[string]GIMBAL_MANAGER_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_MANAGER_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_MANAGER_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_MANAGER_FLAGS) String() string {
	if l, ok := labels_GIMBAL_MANAGER_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
