package ruuvitag

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/paypal/gatt"
)

var manufacturerDataID uint16 = 0x0499

func init() {
	// Discard log message from gatt module
	log.SetOutput(ioutil.Discard)
}

func isRuuviDevice(data []byte) bool {
	return binary.LittleEndian.Uint16(data[0:2]) == manufacturerDataID
}

// Measurement represents RuuviTag sensor data
type Measurement struct {
	DeviceID            string
	Format              uint8
	Humidity            uint8
	Temperature         int8
	TemperatureFraction uint8
	Pressure            uint16
	AccelerationX       int16
	AccelerationY       int16
	AccelerationZ       int16
	BatteryVoltage      uint16
	Timestamp           time.Time
}

// Data format 3 https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/broadcast_formats.md
func dataFormat3(ID string, data []byte) Measurement {
	return Measurement{
		DeviceID:            ID,
		Format:              data[2],
		Humidity:            uint8(data[3]),
		Temperature:         int8(data[4]),
		TemperatureFraction: uint8(data[5]),
		Pressure:            uint16(binary.BigEndian.Uint16(data[6:8])),
		AccelerationX:       int16(binary.BigEndian.Uint16(data[8:10])),
		AccelerationY:       int16(binary.BigEndian.Uint16(data[10:12])),
		AccelerationZ:       int16(binary.BigEndian.Uint16(data[12:14])),
		BatteryVoltage:      uint16(binary.BigEndian.Uint16(data[14:16])),
		Timestamp:           time.Now(),
	}
}

// NewMeasurement creates Measurement from ble manufacturer data
func NewMeasurement(ID string, data []byte) (Measurement, error) {
	switch data[2] {
	case 3:
		return dataFormat3(ID, data), nil
	}
	return Measurement{}, fmt.Errorf("format '%d' if not supported", data[2])
}

func scanDevice(device gatt.Device, output chan Measurement) {
	device.Handle(gatt.PeripheralDiscovered(func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		if isRuuviDevice(a.ManufacturerData) {
			data, err := NewMeasurement(p.ID(), a.ManufacturerData)
			if err == nil && output != nil {
				output <- data
			}
		}
	}))
	device.Init(func(d gatt.Device, s gatt.State) {
		switch s {
		case gatt.StatePoweredOn:
			d.Scan([]gatt.UUID{}, true)
		default:
			d.StopScanning()
			close(output)
		}
	})
}

// Scan starts scanning with default values
func Scan() (chan Measurement, error) {
	device, err := gatt.NewDevice([]gatt.Option{
		gatt.LnxMaxConnections(1),
		gatt.LnxDeviceID(-1, true),
	}...)
	output := make(chan Measurement, 10)
	if err != nil {
		close(output)
		return output, err
	}
	go scanDevice(device, output)
	return output, nil
}
