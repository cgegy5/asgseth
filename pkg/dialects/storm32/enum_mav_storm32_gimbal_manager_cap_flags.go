//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"fmt"
	"strings"
)

// Gimbal manager capability flags.
type MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS uint32

const (
	// The gimbal manager supports several profiles.
	MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS_HAS_PROFILES MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = 1
)

var labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS = map[MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS]string{
	MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS_HAS_PROFILES: "MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS_HAS_PROFILES",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS {
			if l == label {
				mask |= value
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_MANAGER_CAP_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
