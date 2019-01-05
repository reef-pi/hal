package hal

type noopPin struct{ name string }

func NewNoopPin() *noopPin             { return &noopPin{name: "noop-pin"} }
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
			Capabilities: []Capability{Input, Output, PH, Temperature, PWM},
		},
	}
}
func (n *noopDriver) Metadata() Metadata                      { return n.meta }
func (n *noopDriver) Close() error                            { return nil }
func (n *noopDriver) InputPins() []InputPin                   { return []InputPin{} }
func (n *noopDriver) OutputPins() []OutputPin                 { return []OutputPin{} }
func (n *noopDriver) InputPin(_ string) (InputPin, error)     { return new(noopPin), nil }
func (n *noopDriver) OutputPin(_ string) (OutputPin, error)   { return new(noopPin), nil }
func (n *noopDriver) PWMChannels() []PWMChannel               { return []PWMChannel{} }
func (n *noopDriver) PWMChannel(_ string) (PWMChannel, error) { return new(noopChannel), nil }
