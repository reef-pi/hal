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

	var input InputDriver = NewNoopDriver()
	if len(input.InputPins()) != 0 {
		t.Error("Wrong input pins:", len(input.InputPins()))
	}
	pin, err := input.InputPin(4)
	if err != nil {
		t.Error(err)
	}
	if _, err := pin.Read(); err != nil {
		t.Error(err)
	}
	var output OutputDriver = NewNoopDriver()
	if len(output.OutputPins()) != 0 {
		t.Error("Wrong output pins:", len(output.OutputPins()))
	}
	opin, perr := output.OutputPin(4)
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
