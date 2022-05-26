//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// Send a key-value pair as integer. The use of this message is discouraged for normal packets, but a quite efficient way for testing new messages and getting experimental debug output.
type MessageNamedValueInt struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Name of the debug variable
	Name string `mavlen:"10"`
	// Signed integer value
	Value int32
}

// GetID implements the msg.Message interface.
func (*MessageNamedValueInt) GetID() uint32 {
	return 252
}