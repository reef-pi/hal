package hal

import (
	"testing"
)

func TestCalibration(t *testing.T) {
	d := 0.2
	point1 := Measurement{
		Expected: 7.1,
		Observed: 7.1 + d,
	}
	point2 := Measurement{
		Expected: 9.4,
		Observed: 9.4 + d,
	}

	noop, err := CalibratorFactory(nil)
	if err != nil {
		t.Error()
	}

	if noop.Calibrate(12) != 12 {
		t.Error("Expected 12, found:", noop.Calibrate(12))
	}
	empty, err := CalibratorFactory([]Measurement{})
	if err != nil {
		t.Error(err)
	}
	if empty.Calibrate(12) != 12 {
		t.Error("Expected 12, found:", noop.Calibrate(12))
	}
	onePoint, err := CalibratorFactory([]Measurement{point1})
	if err != nil {
		t.Error(err)
	}
	if onePoint.Calibrate(12) != (12 - d) {
		t.Error("Expected 11.9, found:", onePoint.Calibrate(12))
	}
	twoPoint, err := CalibratorFactory([]Measurement{point1, point2})
	if err != nil {
		t.Error(err)
	}
	if twoPoint.Calibrate(12) != 11.8 {
		t.Error("Expected 12.1, found:", twoPoint.Calibrate(12))
	}
	if _, err := CalibratorFactory([]Measurement{point1, point2, point2}); err == nil {
		t.Error("Three point calibration is not implemented")
	}
}
