package hal

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type noopDigitalPin struct{ name string }

func NewDigitalNoopPin() *noopDigitalPin      { return &noopDigitalPin{name: "noop-pin"} }
func (p *noopDigitalPin) Name() string        { return p.name }
func (p *noopDigitalPin) Close() error        { return nil }
func (p *noopDigitalPin) Read() (bool, error) { return true, nil }
func (p *noopDigitalPin) Write(_ bool) error  { return nil }
func (p *noopDigitalPin) LastState() bool     { return true }
func (p *noopDigitalPin) Number() int         { return 0 }

type noopChannel struct {
	noopDigitalPin
}

func (ch *noopChannel) Set(_ float64) error { return nil }
func (ch *noopChannel) Name() string        { return "noop pwm channel" }

type noopAnalogPin struct {
	noopDigitalPin
}

func (p *noopAnalogPin) Read() (float64, error)          { return 0, nil }
func (p *noopAnalogPin) Measure() (float64, error)       { return 0, nil }
func (p *noopAnalogPin) Calibrate(_ []Measurement) error { return nil }

type noopDriver struct {
	meta Metadata
}

func (n *noopDriver) Close() error                                   { return nil }
func (n *noopDriver) DigitalInputPins() []DigitalInputPin            { return []DigitalInputPin{} }
func (n *noopDriver) AnalogInputPins() []AnalogInputPin              { return []AnalogInputPin{} }
func (n *noopDriver) AnalogInputPin(_ int) (AnalogInputPin, error)   { return new(noopAnalogPin), nil }
func (n *noopDriver) DigitalOutputPins() []DigitalOutputPin          { return []DigitalOutputPin{} }
func (n *noopDriver) DigitalInputPin(_ int) (DigitalInputPin, error) { return new(noopDigitalPin), nil }
func (n *noopDriver) DigitalOutputPin(_ int) (DigitalOutputPin, error) {
	return new(noopDigitalPin), nil
}
func (n *noopDriver) Metadata() Metadata                   { return n.meta }
func (n *noopDriver) PWMChannels() []PWMChannel            { return []PWMChannel{} }
func (n *noopDriver) PWMChannel(_ int) (PWMChannel, error) { return new(noopChannel), nil }

func (n *noopDriver) Pins(cap Capability) ([]Pin, error) {
	switch cap {
	case DigitalInput, DigitalOutput:
		return []Pin{new(noopDigitalPin)}, nil
	case AnalogInput:
		return []Pin{new(noopAnalogPin)}, nil
	case PWM:
		return []Pin{new(noopChannel)}, nil
	default:
		return nil, fmt.Errorf("Unknown capability: %v", cap)
	}
}

type noopFactory struct {
	meta       Metadata
	parameters []ConfigParameter
}

var factory *noopFactory
var once sync.Once

// NoopFactory provides the factory to get NoopDriver parameters and NoopDrivers
func NoopFactory() DriverFactory {
	once.Do(func() {
		factory = &noopFactory{
			meta: Metadata{
				Name:         "noop-driver",
				Description:  "No operation (stub/null) hal driver for testing",
				Capabilities: []Capability{DigitalInput, DigitalOutput, AnalogInput, PWM},
			},
			parameters: []ConfigParameter{
				{
					Name:    "Sample Parameter",
					Value:   nil,
					Type:    String,
					Order:   0,
					Default: "sample",
				},
			},
		}
	})
	return factory
}

func (n *noopFactory) GetParameters() []ConfigParameter {
	return n.parameters
}

func (n *noopFactory) ValidateParameters(parameters map[string]interface{}) (bool, []string) {

	var failures = make([]string, 0, 0)
	if paramCount := len(parameters); paramCount != 1 {
		failures = append(failures, fmt.Sprint("NoopDriver expects 1 and only 1 parameter, but received", paramCount, "."))
	}

	for k, v := range parameters {
		if k == "Sample Parameter" {
			val, ok := v.(string)
			if !ok {
				failures = append(failures, fmt.Sprint("Sample Parameter is not a string.", v, "was received."))
			}
			if len(val) < 3 {
				failures = append(failures, fmt.Sprint("Sample Parameter must be at least 3 characters long. ", v, "was received."))
			}
		} else {
			failures = append(failures, fmt.Sprint("Unrecognized parameter:", k, ":", v, "."))
		}
	}

	return len(failures) == 0, failures
}

func (n *noopFactory) Metadata() Metadata {
	return n.meta
}

func (n *noopFactory) NewDriver(parameters map[string]interface{}, hardwareResources interface{}) (Driver, error) {
	if valid, failures := n.ValidateParameters(parameters); !valid {
		msg := "Invalid parameters: \n" + strings.Join(failures, "\n")
		return nil, errors.New(msg)
	}
	return &noopDriver{meta: n.meta}, nil
}
