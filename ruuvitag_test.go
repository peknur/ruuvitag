package ruuvitag

import (
	"encoding/hex"
	"testing"
)

func TestIsRuuvitag(t *testing.T) {

}

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md
func TestDataFormat3(t *testing.T) {
	// min 0x0300FF6300008001800180010000
	// max 0x03FF7F63FFFF7FFF7FFF7FFFFFFF
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

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md
func TestDataFormat5(t *testing.T) {
	_, err := hex.DecodeString("99040510fe2834bbd1fbe8fff40010b0b6000266f27ac507dadc99040339135ebb6c03e6ffdffff40bd7")
	if err != nil {
		t.Fatal(err)
	}
}
