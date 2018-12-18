# reef-pi - Hardware Abstraction Layer

The `hal` package provide common types for hardware capabilities in reef-pi.
It hides device specific details from the controller logic. reef-pi modules like ATO,
pH, temperature uses hal to perform hardware based operations.

'hal' is intended to be primarily interfaces only, void of business logic and
any other dependency package. A NoopDriver or [null](https://en.wikipedia.org/wiki/Null_object_pattern)
driver is included to ease testing.

## Usage

```go
import(
  "github.com/reef-pi/hal"
)

func main() {
  var d hal.Driver = hal.NewNoopDriver()
  d.Metadata()
  defer d.Close()

  input, _ := d.(hal.InputDriver)
  pin, _ := input.InputPin("GP4")
  v, _ := pin.Read()
  for _, pin := range input.InputPins() {
    fmt.Println(pin.Name())
  }

  output, _ := d.(hal.OutputDriver)
  pin, _ := output.OutputPin("GP4")
  pin.Write(false)

  var pwm PWMDriver = hal.NewNoopDrive()
  ch, _ := pwm.PWMChannel("foo")
  ch.Set(10.23)
  for _, ch := range pwm.PWMChannels() {
     fmt.Println(ch.Name())
  }
}


```

