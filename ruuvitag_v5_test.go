package ruuvitag

import (
	"encoding/hex"
	"testing"
)

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md
func TestDataFormat5MaximumValues(t *testing.T) {
	data, err := hex.DecodeString("9904057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMeasurement("", data)
	if err != nil {
		t.Fatal(err)
	}
	if m.Pressure() != 115534 {
		t.Errorf("Pressure = %d; want 115534", m.Pressure())
	}
	if m.Temperature() != 163.835 {
		t.Errorf("Temperature = %f; want 163.835", m.Temperature())
	}
	if m.Humidity() != 163.8350 {
		t.Errorf("Humidity = %.4f; want 163.8350", m.Humidity())
	}
	if m.BatteryVoltage() != 3.646 {
		t.Errorf("BatteryVoltage = %.3f; want 3.646", m.BatteryVoltage())
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

	if m.TXPower() != 20 {
		t.Errorf("TXPower = %d; want 20", m.TXPower())
	}

	if m.MovementCounter() != 254 {
		t.Errorf("MovementCounter = %d; want 254", m.MovementCounter())
	}

	if m.Sequence() != 65534 {
		t.Errorf("MovementCounter = %d; want 65534", m.Sequence())
	}
}

func TestDataFormat5MinimumValues(t *testing.T) {
	data, err := hex.DecodeString("9904058001000000008001800180010000000000CBB8334C884F")
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
	if m.Temperature() != -163.835 {
		t.Errorf("Temperature = %.3f; want -163.835", m.Temperature())
	}
	if m.Humidity() != 0.000 {
		t.Errorf("Humidity = %.3f; want 0.000", m.Humidity())
	}
	if m.BatteryVoltage() != 1.600 {
		t.Errorf("BatteryVoltage = %.3f; want -3.646", m.BatteryVoltage())
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

	if m.TXPower() != -40 {
		t.Errorf("TXPower = %d; want -40", m.TXPower())
	}

	if m.MovementCounter() != 0 {
		t.Errorf("MovementCounter = %d; want 0", m.MovementCounter())
	}

	if m.Sequence() != 0 {
		t.Errorf("MovementCounter = %d; want 0", m.Sequence())
	}
}
