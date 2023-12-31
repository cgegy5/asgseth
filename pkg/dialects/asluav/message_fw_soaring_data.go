//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Fixed-wing soaring (i.e. thermal seeking) data
type MessageFwSoaringData struct {
	// Timestamp
	Timestamp uint64
	// Timestamp since last mode change
	Timestampmodechanged uint64 `mavname:"timestampModeChanged"`
	// Thermal core updraft strength
	Xw float32 `mavname:"xW"`
	// Thermal radius
	Xr float32 `mavname:"xR"`
	// Thermal center latitude
	Xlat float32 `mavname:"xLat"`
	// Thermal center longitude
	Xlon float32 `mavname:"xLon"`
	// Variance W
	Varw float32 `mavname:"VarW"`
	// Variance R
	Varr float32 `mavname:"VarR"`
	// Variance Lat
	Varlat float32 `mavname:"VarLat"`
	// Variance Lon
	Varlon float32 `mavname:"VarLon"`
	// Suggested loiter radius
	Loiterradius float32 `mavname:"LoiterRadius"`
	// Suggested loiter direction
	Loiterdirection float32 `mavname:"LoiterDirection"`
	// Distance to soar point
	Disttosoarpoint float32 `mavname:"DistToSoarPoint"`
	// Expected sink rate at current airspeed, roll and throttle
	Vsinkexp float32 `mavname:"vSinkExp"`
	// Measurement / updraft speed at current/local airplane position
	Z1Localupdraftspeed float32 `mavname:"z1_LocalUpdraftSpeed"`
	// Measurement / roll angle tracking error
	Z2Deltaroll float32 `mavname:"z2_DeltaRoll"`
	// Expected measurement 1
	Z1Exp float32
	// Expected measurement 2
	Z2Exp float32
	// Thermal drift (from estimator prediction step only)
	Thermalgsnorth float32 `mavname:"ThermalGSNorth"`
	// Thermal drift (from estimator prediction step only)
	Thermalgseast float32 `mavname:"ThermalGSEast"`
	// Total specific energy change (filtered)
	TseDot float32 `mavname:"TSE_dot"`
	// Debug variable 1
	Debugvar1 float32 `mavname:"DebugVar1"`
	// Debug variable 2
	Debugvar2 float32 `mavname:"DebugVar2"`
	// Control Mode [-]
	Controlmode uint8 `mavname:"ControlMode"`
	// Data valid [-]
	Valid uint8
}

// GetID implements the message.Message interface.
func (*MessageFwSoaringData) GetID() uint32 {
	return 8011
}
