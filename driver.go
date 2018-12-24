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

func (c *Capabilities) HasCapabilities(check Capabilities) bool {
	match := true
	if check.Input {
		if !c.Input {
			match = false
		}
	}
	if check.Output {
		if !c.Output {
			match = false
		}
	}
	if check.PWM {
		if !c.PWM {
			match = false
		}
	}
	if check.Temperature {
		if !c.Temperature {
			match = false
		}
	}
	if check.PH {
		if !c.PH {
			match = false
		}
	}
	return match
}

type Driver interface {
	io.Closer
	Metadata() Metadata
}
