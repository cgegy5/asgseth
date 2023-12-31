//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strings"
)

// Emergency status encoding
type UAVIONIX_ADSB_EMERGENCY_STATUS uint32

const (
	UAVIONIX_ADSB_OUT_NO_EMERGENCY                    UAVIONIX_ADSB_EMERGENCY_STATUS = 0
	UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = 1
	UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY             UAVIONIX_ADSB_EMERGENCY_STATUS = 2
	UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY          UAVIONIX_ADSB_EMERGENCY_STATUS = 3
	UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY               UAVIONIX_ADSB_EMERGENCY_STATUS = 4
	UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY UAVIONIX_ADSB_EMERGENCY_STATUS = 5
	UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY       UAVIONIX_ADSB_EMERGENCY_STATUS = 6
	UAVIONIX_ADSB_OUT_RESERVED                        UAVIONIX_ADSB_EMERGENCY_STATUS = 7
)

var labels_UAVIONIX_ADSB_EMERGENCY_STATUS = map[UAVIONIX_ADSB_EMERGENCY_STATUS]string{
	UAVIONIX_ADSB_OUT_NO_EMERGENCY:                    "UAVIONIX_ADSB_OUT_NO_EMERGENCY",
	UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY:               "UAVIONIX_ADSB_OUT_GENERAL_EMERGENCY",
	UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY:             "UAVIONIX_ADSB_OUT_LIFEGUARD_EMERGENCY",
	UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY:          "UAVIONIX_ADSB_OUT_MINIMUM_FUEL_EMERGENCY",
	UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY:               "UAVIONIX_ADSB_OUT_NO_COMM_EMERGENCY",
	UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY: "UAVIONIX_ADSB_OUT_UNLAWFUL_INTERFERANCE_EMERGENCY",
	UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY:       "UAVIONIX_ADSB_OUT_DOWNED_AIRCRAFT_EMERGENCY",
	UAVIONIX_ADSB_OUT_RESERVED:                        "UAVIONIX_ADSB_OUT_RESERVED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_EMERGENCY_STATUS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_UAVIONIX_ADSB_EMERGENCY_STATUS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_EMERGENCY_STATUS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UAVIONIX_ADSB_EMERGENCY_STATUS
	for _, label := range labels {
		found := false
		for value, l := range labels_UAVIONIX_ADSB_EMERGENCY_STATUS {
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
func (e UAVIONIX_ADSB_EMERGENCY_STATUS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
