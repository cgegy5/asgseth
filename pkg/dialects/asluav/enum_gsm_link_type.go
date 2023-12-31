//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"fmt"
	"strings"
)

type GSM_LINK_TYPE uint32

const (
	// no service
	GSM_LINK_TYPE_NONE GSM_LINK_TYPE = 0
	// link type unknown
	GSM_LINK_TYPE_UNKNOWN GSM_LINK_TYPE = 1
	// 2G (GSM/GRPS/EDGE) link
	GSM_LINK_TYPE_2G GSM_LINK_TYPE = 2
	// 3G link (WCDMA/HSDPA/HSPA)
	GSM_LINK_TYPE_3G GSM_LINK_TYPE = 3
	// 4G link (LTE)
	GSM_LINK_TYPE_4G GSM_LINK_TYPE = 4
)

var labels_GSM_LINK_TYPE = map[GSM_LINK_TYPE]string{
	GSM_LINK_TYPE_NONE:    "GSM_LINK_TYPE_NONE",
	GSM_LINK_TYPE_UNKNOWN: "GSM_LINK_TYPE_UNKNOWN",
	GSM_LINK_TYPE_2G:      "GSM_LINK_TYPE_2G",
	GSM_LINK_TYPE_3G:      "GSM_LINK_TYPE_3G",
	GSM_LINK_TYPE_4G:      "GSM_LINK_TYPE_4G",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GSM_LINK_TYPE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GSM_LINK_TYPE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GSM_LINK_TYPE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GSM_LINK_TYPE
	for _, label := range labels {
		found := false
		for value, l := range labels_GSM_LINK_TYPE {
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
func (e GSM_LINK_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
