package hal

type AnalogInputPin interface {
	Name() string
	Read() (float64, error)
	Calibrate([]Measurement) error
	Measure() (float64, error)
}

type AnalogInputDriver interface {
	Driver
	AnalogInputPins() []AnalogInputPin
	AnalogInputPin(int) (AnalogInputPin, error)
}
