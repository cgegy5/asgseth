//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Backwards compatible version of SERIAL_UDB_EXTRA F17 format
type MessageSerialUdbExtraF17 struct {
	// SUE Feed Forward Gain
	SueFeedForward float32
	// SUE Max Turn Rate when Navigating
	SueTurnRateNav float32
	// SUE Max Turn Rate in Fly By Wire Mode
	SueTurnRateFbw float32
}

// GetID implements the message.Message interface.
func (*MessageSerialUdbExtraF17) GetID() uint32 {
	return 183
}
