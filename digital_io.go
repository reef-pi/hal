package hal

import (
	"io"
)

// Pin represents a single-bit digital input or output
type Pin interface {
	io.Closer
	Name() string
}

// InputPin represents an input pin with a single digital input
// value.
type InputPin interface {
	Pin
	// Read returns whether this input is logical high (true) or low (false)
	Read() (bool, error)
}

type InputDriver interface {
	Driver
	InputPins() []InputPin
	InputPin(name string) (InputPin, error)
}

type OutputPin interface {
	Pin
	Write(state bool) error
	LastState() bool
}

type OutputDriver interface {
	Driver
	OutputPins() []OutputPin
	OutputPin(name string) (OutputPin, error)
}
