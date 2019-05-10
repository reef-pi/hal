package hal

type TemperatureChannel interface {
	// Read returns a logical temperature in degrees celsius
	ReadTemperature() (float64, error)
}

type TemperatureDriver interface {
	Driver
	TemperatureChannels() []TemperatureChannel
	TemperatureChannel(int) (TemperatureChannel, error)
}
