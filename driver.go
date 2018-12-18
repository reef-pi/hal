package hal

import "io"

// Metadata represents basic information about a driver
// for the API response.
type Metadata struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Capabilities Capabilities `json:"capabilities"`
}

// Capabilities defines which
type Capabilities struct {
	Input       bool `json:"input"`
	Output      bool `json:"output"`
	PWM         bool `json:"pwm"`
	Temperature bool `json:"temperature"`
	PH          bool `json:"ph"`
}

type Driver interface {
	io.Closer
	Metadata() Metadata
}
