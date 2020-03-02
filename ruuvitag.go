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

// Measurement represents RuuviTag sensor readings
type Measurement interface {
	DeviceID() string
	Format() uint8
	Humidity() float64
	Temperature() float64
	Pressure() uint32
	AccelerationX() float64
	AccelerationY() float64
	AccelerationZ() float64
	BatteryVoltage() float64
	Timestamp() time.Time
}

func init() {
	// Discard log message from gatt module
	log.SetOutput(ioutil.Discard)
}

func isRuuviDevice(data []byte) bool {
	return binary.LittleEndian.Uint16(data[0:2]) == manufacturerDataID
}

// NewMeasurement creates Measurement from ble manufacturer data
func NewMeasurement(ID string, data []byte) (Measurement, error) {
	// switch data format
	switch data[2] {
	case 3:
		return newDataFormat3(ID, data)
	}
	return nil, fmt.Errorf("format '%d' if not supported", data[2])
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

// Scan starts scanning with default options
func Scan(output chan Measurement) error {
	device, err := gatt.NewDevice([]gatt.Option{
		gatt.LnxMaxConnections(1),
		gatt.LnxDeviceID(-1, true),
	}...)
	if err != nil {
		return err
	}
	go scanDevice(device, output)
	return nil
}
