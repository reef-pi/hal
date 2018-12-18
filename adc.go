package hal

type ADCChannel interface {
	Name() string
	Set(value float64) error
}
type ADCDriver interface {
	Driver
	ADChannels() []PWMChannel
	ADCChannel(name string) (ADCChannel, error)
}
