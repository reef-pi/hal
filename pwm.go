package hal

type PWMChannel interface {
	OutputPin
	Set(value float64) error
}

type PWMDriver interface {
	OutputDriver
	PWMChannels() []PWMChannel
	PWMChannel(name string) (PWMChannel, error)
}
