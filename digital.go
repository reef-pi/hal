package hal

// InputPin represents an input pin with a single digital input
// value.
type DigitalInputPin interface {
	Pin
	// Read returns whether this input is logical high (true) or low (false)
	Read() (bool, error)
}

type DigitalInputDriver interface {
	Driver
	DigitalInputPins() []DigitalInputPin
	DigitalInputPin(int) (DigitalInputPin, error)
}

type DigitalOutputPin interface {
	Pin
	Write(state bool) error
	LastState() bool
}

type DigitalOutputDriver interface {
	Driver
	DigitalOutputPins() []DigitalOutputPin
	DigitalOutputPin(int) (DigitalOutputPin, error)
}
