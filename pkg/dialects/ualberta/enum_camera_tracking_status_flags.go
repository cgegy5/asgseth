//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Camera tracking status flags
type CAMERA_TRACKING_STATUS_FLAGS uint32

const (
	// Camera is not tracking
	CAMERA_TRACKING_STATUS_FLAGS_IDLE CAMERA_TRACKING_STATUS_FLAGS = 0
	// Camera is tracking
	CAMERA_TRACKING_STATUS_FLAGS_ACTIVE CAMERA_TRACKING_STATUS_FLAGS = 1
	// Camera tracking in error state
	CAMERA_TRACKING_STATUS_FLAGS_ERROR CAMERA_TRACKING_STATUS_FLAGS = 2
)

var labels_CAMERA_TRACKING_STATUS_FLAGS = map[CAMERA_TRACKING_STATUS_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_TRACKING_STATUS_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_TRACKING_STATUS_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_TRACKING_STATUS_FLAGS = map[string]CAMERA_TRACKING_STATUS_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_TRACKING_STATUS_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_TRACKING_STATUS_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_TRACKING_STATUS_FLAGS) String() string {
	if l, ok := labels_CAMERA_TRACKING_STATUS_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}