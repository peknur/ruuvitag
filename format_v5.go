package ruuvitag

import (
	"encoding/binary"
	"fmt"
	"time"
)

// dataFormat5 (also known as RAWv2) measurement
// @see https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md
type dataFormat5 struct {
	deviceID       string
	format         uint8
	humidity       uint16
	temperature    int16
	pressure       uint16
	accelerationX  int16
	accelerationY  int16
	accelerationZ  int16
	batteryVoltage uint16
	txPower        int8
	movement       uint8
	sequence       uint16
	mac            string
	timestamp      time.Time
}

func (f *dataFormat5) DeviceID() string {
	return f.deviceID
}

func (f *dataFormat5) Format() uint8 {
	return f.format
}

func (f *dataFormat5) Humidity() float32 {
	//  1/0.0025 = 400
	return float32(f.humidity) / 400
}

func (f *dataFormat5) Temperature() float32 {
	// 1 / 0.005 = 200
	return float32(f.temperature) / 200
}

func (f *dataFormat5) Pressure() uint32 {
	return uint32(f.pressure) + 50000
}

func (f *dataFormat5) AccelerationX() float32 {
	return float32(f.accelerationX) / 1000
}

func (f *dataFormat5) AccelerationY() float32 {
	return float32(f.accelerationY) / 1000
}

func (f *dataFormat5) AccelerationZ() float32 {
	return float32(f.accelerationZ) / 1000
}

func (f *dataFormat5) BatteryVoltage() float32 {
	return float32(f.batteryVoltage) / 1000
}

func (f *dataFormat5) Timestamp() time.Time {
	return f.Timestamp()
}

// NewDataFormat5 https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/broadcast_formats.md
func NewDataFormat5(ID string, data []byte) (Measurement, error) {
	if len(data) != 26 {
		return nil, fmt.Errorf("manufacturer data lenght (%d) mismatch", len(data))
	}

	voltage := make([]byte, 2)
	voltage[0] = data[15]
	voltage[1] = data[16] >> 6

	m := dataFormat5{
		deviceID:       ID,
		format:         data[2],
		temperature:    int16(binary.BigEndian.Uint16(data[3:5])),
		humidity:       binary.BigEndian.Uint16(data[5:7]),
		pressure:       binary.BigEndian.Uint16(data[7:9]),
		accelerationX:  int16(binary.BigEndian.Uint16(data[9:11])),
		accelerationY:  int16(binary.BigEndian.Uint16(data[11:13])),
		accelerationZ:  int16(binary.BigEndian.Uint16(data[13:15])),
		batteryVoltage: binary.BigEndian.Uint16(voltage),
		movement:       uint8(data[17]),
		sequence:       binary.BigEndian.Uint16(data[18:20]),
		mac:            string(data[20:]),
		timestamp:      time.Now(),
	}
	return &m, nil
}
