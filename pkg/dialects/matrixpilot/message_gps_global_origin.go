//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

import (
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
)

// Publishes the GPS coordinates of the vehicle local origin (0,0,0) position. Emitted whenever a new GPS-Local position mapping is requested or set - e.g. following SET_GPS_GLOBAL_ORIGIN message.
type MessageGpsGlobalOrigin = common.MessageGpsGlobalOrigin