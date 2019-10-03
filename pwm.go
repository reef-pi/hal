package hal

type PWMChannel interface {
	DigitalOutputPin
	Set(value float64) error
}

type PWMDriver interface {
	DigitalOutputDriver
	PWMChannels() []PWMChannel
	PWMChannel(int) (PWMChannel, error)
}
