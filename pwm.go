package hal

type PWMChannel interface {
	Name() string
	Set(value float64) error
}

type PWMDriver interface {
	Driver
	PWMChannels() []PWMChannel
	PWMChannel(name string) (PWMChannel, error)
}
