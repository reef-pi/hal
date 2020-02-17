// Package hal defines the interfaces for implementing a driver in reef-pi
package hal

import "io"

// Capability represents the capabilities of a driver
type Capability int

const (
	None Capability = iota
	DigitalInput
	DigitalOutput
	PWM
	AnalogInput
)

var (
	_capNames = []string{
		"none",
		"digital-input",
		"digital-output",
		"pwm",
		"analog-input",
	}
)

func (c Capability) String() string {
	return _capNames[c]
}

// Metadata represents basic information about a driver
// for the API response.
type Metadata struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Capabilities []Capability `json:"capabilities"`
}

// ConfigParameterType indicates the type of a configuration parameter
type ConfigParameterType int

const (
	String ConfigParameterType = iota + 1
	Number
	Boolean
	IpAddress
	Url
)

// ConfigParameter represent a configuration parameter required by a driver
type ConfigParameter struct {
	Name    string
	Value   interface{}
	Type    ConfigParameterType
	Order   int
	Default interface{}
}

func (m Metadata) HasCapability(cap Capability) bool {
	for _, c := range m.Capabilities {
		if c == cap {
			return true
		}
	}
	return false
}

// Pin represents a single-bit digital input or output
type Pin interface {
	io.Closer
	Name() string
	Number() int
}

type Driver interface {
	io.Closer
	Metadata() Metadata
	Pins(Capability) ([]Pin, error)
}

// DriverFactory is responsible for creating drivers and providing information about a driver.
type DriverFactory interface {

	//GetParameters returns the parameters that the driver expects.
	GetParameters() []ConfigParameter

	//ValidateParameters validates the parameters for a driver.
	//The boolean result is true if the parameters are valid (and the array of failure messages should be nil or empty).
	//The array of failures should contain all of the validation errors. It should not short circuit after the first failure.
	ValidateParameters(parameters map[string]interface{}) (bool, []string)

	//Metadata returns the Metadata the driver can provide.
	Metadata() Metadata

	//CreateDriver validates the parameters and returns the driver if validation succeeds.
	NewDriver(parameters map[string]interface{}, hardwareResources interface{}) (Driver, error)
}
