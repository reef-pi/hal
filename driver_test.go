package hal

import (
	"testing"
)

func TestDriver(t *testing.T) {
	var f DriverFactory = NoopFactory()
	var params []ConfigParameter = []ConfigParameter{{
		Name:  "Sample Parameter",
		Value: "Value",
	}}

	d, err := f.CreateDriver(params, nil)
	if err != nil {
		t.Error("Error creating noop driver")
	}

	if d.Metadata().Name == "" {
		t.Error("Noop driver name should not be empty")
	}
	if err := d.Close(); err != nil {
		t.Error()
	}

	var input DigitalInputDriver
	di, _ := f.CreateDriver(params, nil)
	if di.Metadata().HasCapability(DigitalInput) {
		input = di.(DigitalInputDriver)
	} else {
		t.Error("Unable to convert to DigitalInputDriver")
	}

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

	var output DigitalOutputDriver
	dod, _ := f.CreateDriver(params, nil)
	if dod.Metadata().HasCapability(DigitalOutput) {
		output = di.(DigitalOutputDriver)
	} else {
		t.Error("Unable to convert to DigitalOutputDriver")
	}

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

	p, _ := f.CreateDriver(params, nil)
	var pwm PWMDriver
	if p.Metadata().HasCapability(PWM) {
		pwm = p.(PWMDriver)
	} else {
		t.Error("Unable to convert to PWMDriver")
	}

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

	var f DriverFactory = NoopFactory()
	var params []ConfigParameter = []ConfigParameter{{
		Name:  "Sample Parameter",
		Value: "Value",
	}}

	tmp, _ := f.CreateDriver(params, nil)
	var input AnalogInputDriver
	if tmp.Metadata().HasCapability(AnalogInput) {
		input = tmp.(AnalogInputDriver)
	} else {
		t.Error("Unable to convert to AnalogInputDriver")
	}

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

func TestFactory_Get_Parameters_with_Validation_Failure(t *testing.T) {
	var f DriverFactory = NoopFactory()
	params := f.GetParameters()

	if l := len(params); l != 1 {
		t.Error("NoopFactory should return a single parameter")
	}

	var param = ConfigParameter{
		Name:  "Bad Parameter",
		Value: 33,
	}
	params = []ConfigParameter{param}

	_, err := f.CreateDriver(params, nil)

	if err == nil {
		t.Error("Invalid parameters should not create a driver")
	}
}
