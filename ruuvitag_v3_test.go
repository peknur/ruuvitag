package ruuvitag

import (
	"encoding/hex"
	"testing"
)

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md
func TestDataFormat3MinimumValues(t *testing.T) {
	data, err := hex.DecodeString("99040300FF6300008001800180010000")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMeasurement("", data)
	if err != nil {
		t.Fatal(err)
	}
	if m.Pressure() != 50000 {
		t.Errorf("Pressure = %d; want 50000", m.Pressure())
	}
	if m.Temperature() != -127.99 {
		t.Errorf("Temperature = %.2f; want -127.99", m.Temperature())
	}
	if m.Humidity() != 0.0 {
		t.Errorf("Humidity = %.2f; want 0.0", m.Humidity())
	}
	if m.BatteryVoltage() != 0.000 {
		t.Errorf("BatteryVoltage = %.3f; want 2899", m.BatteryVoltage())
	}
	if m.AccelerationX() != -32.767 {
		t.Errorf("AccelerationX = %.3f; want -32.767", m.AccelerationX())
	}

	if m.AccelerationY() != -32.767 {
		t.Errorf("AccelerationY = %.3f; want -32.767", m.AccelerationY())
	}

	if m.AccelerationZ() != -32.767 {
		t.Errorf("AccelerationZ = %.3f; want -32.767", m.AccelerationZ())
	}
}
func TestDataFormat3MaximumValues(t *testing.T) {
	data, err := hex.DecodeString("990403FF7F63FFFF7FFF7FFF7FFFFFFF")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMeasurement("", data)
	if err != nil {
		t.Fatal(err)
	}
	if m.Pressure() != 115535 {
		t.Errorf("Pressure = %d; want 115535", m.Pressure())
	}
	if m.Temperature() != 127.99 {
		t.Errorf("Temperature = %.2f; want 127.99", m.Temperature())
	}
	if m.Humidity() != 127.5 {
		t.Errorf("Humidity = %.1f; want 127.5", m.Humidity())
	}
	if m.BatteryVoltage() != 65.535 {
		t.Errorf("BatteryVoltage = %.3f; want 65.535", m.BatteryVoltage())
	}
	if m.AccelerationX() != 32.767 {
		t.Errorf("AccelerationX = %.3f; want 32.767", m.AccelerationX())
	}

	if m.AccelerationY() != 32.767 {
		t.Errorf("AccelerationY = %.3f; want 32.767", m.AccelerationY())
	}

	if m.AccelerationZ() != 32.767 {
		t.Errorf("AccelerationZ = %.3f; want 32.767", m.AccelerationZ())
	}
}
