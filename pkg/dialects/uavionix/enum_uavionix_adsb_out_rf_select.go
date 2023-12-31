//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strings"
)

// Transceiver RF control flags for ADS-B transponder dynamic reports
type UAVIONIX_ADSB_OUT_RF_SELECT uint32

const (
	UAVIONIX_ADSB_OUT_RF_SELECT_STANDBY    UAVIONIX_ADSB_OUT_RF_SELECT = 0
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 1
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED UAVIONIX_ADSB_OUT_RF_SELECT = 2
)

var labels_UAVIONIX_ADSB_OUT_RF_SELECT = map[UAVIONIX_ADSB_OUT_RF_SELECT]string{
	UAVIONIX_ADSB_OUT_RF_SELECT_STANDBY:    "UAVIONIX_ADSB_OUT_RF_SELECT_STANDBY",
	UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED: "UAVIONIX_ADSB_OUT_RF_SELECT_RX_ENABLED",
	UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED: "UAVIONIX_ADSB_OUT_RF_SELECT_TX_ENABLED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_RF_SELECT) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_UAVIONIX_ADSB_OUT_RF_SELECT {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_RF_SELECT) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UAVIONIX_ADSB_OUT_RF_SELECT
	for _, label := range labels {
		found := false
		for value, l := range labels_UAVIONIX_ADSB_OUT_RF_SELECT {
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
func (e UAVIONIX_ADSB_OUT_RF_SELECT) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
