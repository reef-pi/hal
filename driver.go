package hal

import "io"

type Capability int

const (
	None Capability = iota
	Input
	Output
	PH
	Temperature
	PWM
)

// Metadata represents basic information about a driver
// for the API response.
type Metadata struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Capabilities []Capability `json:"capabilities"`
}

func (m Metadata) HasCapability(cap Capability) bool {
	for _, c := range m.Capabilities {
		if c == cap {
			return true
		}
	}
	return false
}

type Driver interface {
	io.Closer
	Metadata() Metadata
}
