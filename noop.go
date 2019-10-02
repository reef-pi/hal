package hal

import (
	"fmt"
)

type noopPin struct{ name string }

func NewDigitalNoopPin() *noopPin      { return &noopPin{name: "noop-pin"} }
func (p *noopPin) Name() string        { return p.name }
func (p *noopPin) Close() error        { return nil }
func (p *noopPin) Read() (bool, error) { return true, nil }
func (p *noopPin) Write(_ bool) error  { return nil }
func (p *noopPin) LastState() bool     { return true }

type noopChannel struct {
	noopPin
}

func (ch *noopChannel) Set(_ float64) error { return nil }
func (ch *noopChannel) Name() string        { return "noop pwm channel" }

type noopDriver struct {
	meta Metadata
}

func NewNoopDriver() *noopDriver {
	return &noopDriver{
		meta: Metadata{
			Name:         "noop-driver",
			Description:  "No operation (stub/null) hal driver for testing",
			Capabilities: []Capability{DigitalInput, DigitalOutput, AnalogInput, PWM},
		},
	}
}
func (n *noopDriver) Metadata() Metadata                               { return n.meta }
func (n *noopDriver) Close() error                                     { return nil }
func (n *noopDriver) DigitalInputPins() []DigitalInputPin              { return []DigitalInputPin{} }
func (n *noopDriver) DigitalOutputPins() []DigitalOutputPin            { return []DigitalOutputPin{} }
func (n *noopDriver) DigitalInputPin(_ int) (DigitalInputPin, error)   { return new(noopPin), nil }
func (n *noopDriver) DigitalOutputPin(_ int) (DigitalOutputPin, error) { return new(noopPin), nil }
func (n *noopDriver) PWMChannels() []PWMChannel                        { return []PWMChannel{} }
func (n *noopDriver) PWMChannel(_ int) (PWMChannel, error)             { return new(noopChannel), nil }

func (n *noopDriver) Pins(cap Capability) ([]Pin, error) {
	switch cap {
	case DigitalInput:
	case DigitalOutput:
		return []Pin{new(noopPin)}, nil
	case PWM:
		return []Pin{new(noopChannel)}, nil
	default:
		return nil, fmt.Errorf("Unknown capability: %v", cap)
	}
	return nil, fmt.Errorf("Unknown capability: %v", cap)
}
