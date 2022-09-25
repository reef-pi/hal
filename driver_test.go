package hal

import (
	"testing"
)

func TestDriver(t *testing.T) {
	var f DriverFactory = NoopFactory()

	params := make(map[string]interface{})
	params["Sample Parameter"] = "Value"

	d, err := f.NewDriver(params, nil)
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
	di, _ := f.NewDriver(params, nil)
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
	dod, _ := f.NewDriver(params, nil)
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

	p, _ := f.NewDriver(params, nil)
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

	params := make(map[string]interface{})
	params["Sample Parameter"] = "Value"

	tmp, _ := f.NewDriver(params, nil)
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
	if _, err := ipin.Value(); err != nil {
		t.Error(err)
	}
}

func TestFactory_Get_Parameters_with_Validation_Failure(t *testing.T) {
	var f DriverFactory = NoopFactory()
	p := f.GetParameters()

	if l := len(p); l != 1 {
		t.Error("NoopFactory should return a single parameter")
	}

	params := make(map[string]interface{})
	params["Bad Parameter"] = 33

	_, err := f.NewDriver(params, nil)

	if err == nil {
		t.Error("Invalid parameters should not create a driver")
	}
}

func TestConvertToInt(t *testing.T) {
	var val interface{}
	val = true
	var x int
	var ok bool

	x, ok = ConvertToInt(val)
	if ok {
		t.Error("bool should not convert to int")
	}

	val = float32(37)
	x, ok = ConvertToInt(val)
	if !ok {
		t.Error("float should convert to int")
	}
	if x != 37 {
		t.Error("Float should convert to an int")
	}

	val = float64(37.4)
	x, ok = ConvertToInt(val)
	if ok {
		t.Error("float with decimal value should not convert to int")
	}

	val = "37"
	x, ok = ConvertToInt(val)
	if !ok {
		t.Error("string version of integers should be implemented")
	}

	val = "34.2"
	x, ok = ConvertToInt(val)
	if ok {
		t.Error("string decimal should not convert to int")
	}
}
