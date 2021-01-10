# reef-pi - Hardware Abstraction Layer

[![Build Status](https://github.com/reef-pi/hal/workflows/go/badge.svg?branch=master)
[![Coverage Status](https://codecov.io/gh/reef-pi/hal/branch/master/graph/badge.svg)](https://codecov.io/gh/reef-pi/hal)
[![Go Report Card](https://goreportcard.com/badge/reef-pi/hal)](https://goreportcard.com/report/reef-pi/hal)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/reef-pi/hal/blob/master/LICENSE.txt)
[![GoDoc](https://godoc.org/github.com/reef-pi/hal?status.svg)](https://godoc.org/github.com/reef-pi/hal)

## Introduction


The `hal` package provide common types for hardware capabilities in reef-pi.
It hides device specific details from the controller logic. reef-pi modules like ATO,
and pH, uses hal interfaces to perform hardware based operations. The HAL interfaces cover
metdata and methods for digital or analog I/O. Check out [reef-pi/drivers](https://github.com/reef-pi/drivers)
for examples.

'hal' is mostly made of interfaces and void of business logic and
or dependency package. A No-Op or [null](https://en.wikipedia.org/wiki/Null_object_pattern)
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

