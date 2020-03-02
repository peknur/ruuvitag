package ruuvitag

import (
	"encoding/binary"
	"fmt"
	"time"
)

// DataFormat3 (also known as RAWv1) measurement
type dataFormat3 struct {
	deviceID            string
	format              uint8
	humidity            uint8
	temperature         int8
	temperatureFraction uint8
	pressure            uint16
	accelerationX       int16
	accelerationY       int16
	accelerationZ       int16
	batteryVoltage      uint16
	timestamp           time.Time
}

func (f *dataFormat3) DeviceID() string {
	return f.deviceID
}

func (f *dataFormat3) Format() uint8 {
	return f.format
}
func (f *dataFormat3) Humidity() float64 {
	return float64(f.humidity) / 2
}

func (f *dataFormat3) Temperature() float64 {
	return float64(f.temperature) + float64(f.temperatureFraction)/100
}

func (f *dataFormat3) Pressure() uint32 {
	return uint32(f.pressure) + 50000
}
func (f *dataFormat3) AccelerationX() float64 {
	return float64(f.accelerationX / 1000)
}
func (f *dataFormat3) AccelerationY() float64 {
	return float64(f.accelerationY) / 1000
}
func (f *dataFormat3) AccelerationZ() float64 {
	return float64(f.accelerationZ) / 1000
}
func (f *dataFormat3) BatteryVoltage() float64 {
	return float64(f.batteryVoltage)
}
func (f *dataFormat3) Timestamp() time.Time {
	return f.Timestamp()
}

// Data format 3 https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/broadcast_formats.md
func newDataFormat3(ID string, data []byte) (Measurement, error) {
	if len(data) != 16 {
		return nil, fmt.Errorf("manufacturer data lenght mismatch")
	}

	m := dataFormat3{
		deviceID:            ID,
		format:              data[2],
		humidity:            uint8(data[3]),
		temperature:         int8(data[4]),
		temperatureFraction: uint8(data[5]),
		pressure:            uint16(binary.BigEndian.Uint16(data[6:8])),
		accelerationX:       int16(binary.BigEndian.Uint16(data[8:10])),
		accelerationY:       int16(binary.BigEndian.Uint16(data[10:12])),
		accelerationZ:       int16(binary.BigEndian.Uint16(data[12:14])),
		batteryVoltage:      uint16(binary.BigEndian.Uint16(data[14:16])),
		timestamp:           time.Now(),
	}
	return &m, nil
}
