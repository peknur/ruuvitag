package ruuvitag

import (
	"encoding/binary"
	"fmt"
	"time"
)

// dataFormat3 (also known as RAWv1) measurement
// @see https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md
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

func (f *dataFormat3) Humidity() float32 {
	return float32(f.humidity) / 2
}

func (f *dataFormat3) Temperature() float32 {
	if f.temperature < 0 {
		return float32(f.temperature) - (float32(f.temperatureFraction) / 100)
	}
	return float32(f.temperature) + (float32(f.temperatureFraction) / 100)
}

func (f *dataFormat3) Pressure() uint32 {
	return uint32(f.pressure) + 50000
}

func (f *dataFormat3) AccelerationX() float32 {
	return float32(f.accelerationX) / 1000
}

func (f *dataFormat3) AccelerationY() float32 {
	return float32(f.accelerationY) / 1000
}

func (f *dataFormat3) AccelerationZ() float32 {
	return float32(f.accelerationZ) / 1000
}

func (f *dataFormat3) BatteryVoltage() float32 {
	return float32(f.batteryVoltage) / 1000
}

func (f *dataFormat3) Timestamp() time.Time {
	return f.timestamp
}

func (f *dataFormat3) TXPower() int8 {
	return 0
}

func (f *dataFormat3) MovementCounter() uint8 {
	return 0
}

func (f *dataFormat3) Sequence() uint16 {
	return 0
}

// NewDataFormat3 https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/broadcast_formats.md
func NewDataFormat3(ID string, data []byte) (Measurement, error) {
	if len(data) != 16 {
		return nil, fmt.Errorf("manufacturer data lenght mismatch")
	}
	if data[2] != 3 {
		return nil, fmt.Errorf("data format mismatch (%d)", data[2])
	}
	m := dataFormat3{
		deviceID:            ID,
		format:              data[2],
		humidity:            uint8(data[3]),
		temperature:         msbSignedByteToInt8(data[4]),
		temperatureFraction: uint8(data[5]),
		pressure:            binary.BigEndian.Uint16(data[6:8]),
		accelerationX:       int16(binary.BigEndian.Uint16(data[8:10])),
		accelerationY:       int16(binary.BigEndian.Uint16(data[10:12])),
		accelerationZ:       int16(binary.BigEndian.Uint16(data[12:14])),
		batteryVoltage:      binary.BigEndian.Uint16(data[14:16]),
		timestamp:           time.Now(),
	}
	return &m, nil
}
