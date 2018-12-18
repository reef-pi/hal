# reef-pi - Hardware Abstraction Layer

The `hal` package gives base types and interfaces for various
reef-pi hardware modules. This is designed to be primarily interfaces only
void of business logic, zero dependency package that any custom hardware can
use to make a reef-pi compatible device

## Usage

```go
import(
  "github.com/reef-pi/hal"
)

func main() {
  var d hal.Driver = CustomIODriver()
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

  var pwm PWMDriver = CustomPWMDriver()
  ch, _ := pwm.PWMChannel("foo")
  ch.Set(10.23)
  for _, ch := range pwm.PWMChannels() {
     fmt.Println(ch.Name())
  }
}


```

