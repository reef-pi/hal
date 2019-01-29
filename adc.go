package hal

type ADCChannel interface {
	Name() string
	Read() (float64, error)
	Calibrate([]Measurement) error
	Measure() (float64, error)
}

type ADCDriver interface {
	Driver
	ADCChannels() []ADCChannel
	ADCChannel(int) (ADCChannel, error)
}
