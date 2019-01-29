package hal

import (
	"fmt"
)

type Measurement struct {
	Expected float64 `json:"expected"`
	Observed float64 `json:"observed"`
}

type Calibrator interface {
	Calibrate(float64) float64
}

func CalibratorFactory(points []Measurement) (Calibrator, error) {
	l := len(points)
	switch l {
	case 0:
		return &noopCalibrator{}, nil
	case 1:
		return &onePointCalibration{offset: points[0].Expected - points[0].Observed}, nil
	case 2:
		min := points[0]
		max := points[1]
		if points[0].Expected >= points[1].Expected {
			min = points[1]
			max = points[0]
		}
		return &twoPointCalibration{
			refLow:  min.Expected,
			refHigh: max.Expected,
			rawLow:  min.Observed,
			rawHigh: max.Observed,
		}, nil
	default:
		return nil, fmt.Errorf("Expected calibration type can only be 0, 1 or 2. Found: %d", l)
	}
}

type twoPointCalibration struct {
	rawLow, rawHigh, refLow, refHigh float64
}

func (c *twoPointCalibration) Calibrate(value float64) float64 {
	// CorrectedValue = (((RawValue – RawLow) * ReferenceRange) / RawRange) + ReferenceLow
	return (((value - c.rawLow) * (c.refHigh - c.refLow)) / (c.rawHigh - c.rawLow)) + c.refLow
}

type onePointCalibration struct {
	offset float64
}

func (c *onePointCalibration) Calibrate(value float64) float64 {
	return value + c.offset
}

type noopCalibrator struct{}

func (n *noopCalibrator) Calibrate(v float64) float64 {
	return v
}
