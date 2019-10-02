package hal

type AnalogInputPin interface {
	Pin
	Read() (float64, error)
	Calibrate([]Measurement) error
	Measure() (float64, error)
}

type AnalogInputDriver interface {
	Driver
	AnalogInputPins() []AnalogInputPin
	AnalogInputPin(int) (AnalogInputPin, error)
}
