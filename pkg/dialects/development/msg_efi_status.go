//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// EFI status output
type MessageEfiStatus struct {
	// EFI health status
	Health uint8
	// ECU index
	EcuIndex float32
	// RPM
	Rpm float32
	// Fuel consumed
	FuelConsumed float32
	// Fuel flow rate
	FuelFlow float32
	// Engine load
	EngineLoad float32
	// Throttle position
	ThrottlePosition float32
	// Spark dwell time
	SparkDwellTime float32
	// Barometric pressure
	BarometricPressure float32
	// Intake manifold pressure(
	IntakeManifoldPressure float32
	// Intake manifold temperature
	IntakeManifoldTemperature float32
	// Cylinder head temperature
	CylinderHeadTemperature float32
	// Ignition timing (Crank angle degrees)
	IgnitionTiming float32
	// Injection time
	InjectionTime float32
	// Exhaust gas temperature
	ExhaustGasTemperature float32
	// Output throttle
	ThrottleOut float32
	// Pressure/temperature compensation
	PtCompensation float32
	// Supply voltage to EFI sparking system.  Zero in this value means "unknown", so if the supply voltage really is zero volts use 0.0001 instead.
	IgnitionVoltage float32 `mavext:"true"`
	// Fuel pressure. Zero in this value means "unknown", so if the fuel pressure really is zero kPa use 0.0001 instead.
	FuelPressure float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageEfiStatus) GetID() uint32 {
	return 225
}
