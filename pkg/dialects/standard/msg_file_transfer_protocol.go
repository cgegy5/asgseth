//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// File transfer protocol message: https://mavlink.io/en/services/ftp.html.
type MessageFileTransferProtocol struct {
	// Network ID (0 for broadcast)
	TargetNetwork uint8
	// System ID (0 for broadcast)
	TargetSystem uint8
	// Component ID (0 for broadcast)
	TargetComponent uint8
	// Variable length payload. The length is defined by the remaining message length when subtracting the header and other fields. The content/format of this block is defined in https://mavlink.io/en/services/ftp.html.
	Payload [251]uint8
}

// GetID implements the msg.Message interface.
func (*MessageFileTransferProtocol) GetID() uint32 {
	return 110
}
