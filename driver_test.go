package hal

import (
	"testing"
)

func TestDriver(t *testing.T) {
	var d Driver = NewNoopDriver()
	if d.Metadata().Name == "" {
		t.Error("Noop driver name should not be empty")
	}
	if err := d.Close(); err != nil {
		t.Error()
	}

	var input DigitalInputDriver = NewNoopDriver()
	if len(input.DigitalInputPins()) != 0 {
		t.Error("Wrong input pins:", len(input.DigitalInputPins()))
	}
	pin, err := input.DigitalInputPin(4)
	if err != nil {
		t.Error(err)
	}
	if _, err := pin.Read(); err != nil {
		t.Error(err)
	}
	var output DigitalOutputDriver = NewNoopDriver()
	if len(output.DigitalOutputPins()) != 0 {
		t.Error("Wrong output pins:", len(output.DigitalOutputPins()))
	}
	opin, perr := output.DigitalOutputPin(4)
	if perr != nil {
		t.Error(perr)
	}
	if err := opin.Write(false); err != nil {
		t.Error(err)
	}

	var pwm PWMDriver = NewNoopDriver()
	if len(pwm.PWMChannels()) != 0 {
		t.Error("Wrong number of pwm channels: ", len(pwm.PWMChannels()))
	}

	ppin, nerr := pwm.PWMChannel(1)
	if nerr != nil {
		t.Error(nerr)
	}
	if err := ppin.Set(10.23); err != nil {
		t.Error(err)
	}

}

func TestAnalog(t *testing.T) {
	var input AnalogInputDriver = NewNoopDriver()
	if len(input.AnalogInputPins()) != 0 {
		t.Error("Wrong input pins:", len(input.AnalogInputPins()))
	}
	ipin, perr := input.AnalogInputPin(4)
	if perr != nil {
		t.Error(perr)
	}
	if _, err := ipin.Read(); err != nil {
		t.Error(err)
	}
}
