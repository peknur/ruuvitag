package ruuvitag

import (
	"encoding/hex"
	"testing"
)

func TestIsRuuvitag(t *testing.T) {

}

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md
func TestDataFormat3(t *testing.T) {
	data, err := hex.DecodeString("990403291A1ECE1EFC18F94202CA0B53")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMeasurement("", data)
	if err != nil {
		t.Fatal(err)
	}
	if m.Pressure != 52766 {
		t.Errorf("m.Pressure = %d; want 52766", m.Pressure)
	}
	if m.Temperature != 26 {
		t.Errorf("m.Temperature = %d; want 26", m.Temperature)
	}
	if m.TemperatureFraction != 30 {
		t.Errorf("m.TemperatureFctorure = %d; want 30", m.TemperatureFraction)
	}
	if m.Humidity != 41 {
		t.Errorf("m.Humidity = %d; want 41", m.Humidity)
	}
	if m.BatteryVoltage != 2899 {
		t.Errorf("m.Humidity = %d; want 2899", m.BatteryVoltage)
	}
}

// https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md
func TestDataFormat5(t *testing.T) {
	_, err := hex.DecodeString("99040510fe2834bbd1fbe8fff40010b0b6000266f27ac507dadc99040339135ebb6c03e6ffdffff40bd7")
	if err != nil {
		t.Fatal(err)
	}
}
