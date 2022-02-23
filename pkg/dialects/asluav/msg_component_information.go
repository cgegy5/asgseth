//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

//
// Component information message, which may be requested using MAV_CMD_REQUEST_MESSAGE.
//
// This contains MAVLink FTP URIs for the component's general and peripherals metadata files along with their CRCs (which can be used for file caching).
// The files must be hosted on the component, and may be xz compressed.
// The general metadata file can be read to get the locations of additional metadata files (COMP_METADATA_TYPE), which may be hosted either on the vehicle or the Internet.
// For more information see: https://mavlink.io/en/services/component_information.html.
//
// Note: Camera components should use CAMERA_INFORMATION instead, and autopilots may use both this message and AUTOPILOT_VERSION.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// CRC32 of the general metadata file (general_metadata_uri).
	GeneralMetadataFileCrc uint32
	// MAVLink FTP URI for the general metadata file (COMP_METADATA_TYPE_GENERAL), which may be compressed with xz. The file contains general component metadata, and may contain URI links for additional metadata (see COMP_METADATA_TYPE). The information is static from boot, and may be generated at compile time.
	GeneralMetadataUri string `mavlen:"100"`
	// CRC32 of peripherals metadata file (peripherals_metadata_uri).
	PeripheralsMetadataFileCrc uint32
	// (Optional) MAVLink FTP URI for the peripherals metadata file (COMP_METADATA_TYPE_PERIPHERALS), which may be compressed with xz. This contains data about "attached components" such as UAVCAN nodes. The peripherals are in a separate file because the information must be generated dynamically at runtime.
	PeripheralsMetadataUri string `mavlen:"100"`
}

// GetID implements the msg.Message interface.
func (*MessageComponentInformation) GetID() uint32 {
	return 395
}
