//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Acknowldge sucess or failure of a flexifunction command
type MessageFlexifunctionCommandAck struct {
	// Command acknowledged
	CommandType uint16
	// result of acknowledge
	Result uint16
}

// GetID implements the msg.Message interface.
func (*MessageFlexifunctionCommandAck) GetID() uint32 {
	return 158
}