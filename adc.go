package hal

type ADCChannel interface {
	Name() string
	Read() (float64, error)
}

type ADCDriver interface {
	Driver
	ADCChannels() []ADCChannel
	ADCChannel(name string) (ADCChannel, error)
}
