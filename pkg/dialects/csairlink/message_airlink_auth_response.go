//autogenerated:yes
//nolint:revive,misspell,govet,lll
package csairlink

// Response to the authorization request
type MessageAirlinkAuthResponse struct {
	// Response type
	RespType AIRLINK_AUTH_RESPONSE_TYPE `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageAirlinkAuthResponse) GetID() uint32 {
	return 52001
}
