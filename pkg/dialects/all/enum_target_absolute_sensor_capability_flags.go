//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/development"
)

// These flags indicate the sensor reporting capabilities for TARGET_ABSOLUTE.
type TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS

const (
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_POSITION
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_VELOCITY
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_ACCELERATION
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE     TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_ATTITUDE
	TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES        TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS = development.TARGET_ABSOLUTE_SENSOR_CAPABILITY_RATES
)