//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// This message informs about the currently active SCRIPT.
type MessageScriptCurrent struct {
	// Active Sequence
	Seq uint16
}

// GetID implements the message.Message interface.
func (*MessageScriptCurrent) GetID() uint32 {
	return 184
}
