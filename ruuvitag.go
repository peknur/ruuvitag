package ruuvitag

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/peknur/gatt"
)

var manufacturerDataID uint16 = 0x0499

// Measurement represents RuuviTag sensor readings
type Measurement interface {
	DeviceID() string
	Format() uint8
	Humidity() float32
	Temperature() float32
	Pressure() uint32
	AccelerationX() float32
	AccelerationY() float32
	AccelerationZ() float32
	BatteryVoltage() float32
	TXPower() int8
	MovementCounter() uint8
	Sequence() uint16
	Timestamp() time.Time
}

// NewMeasurement creates Measurement from ble manufacturer data
func NewMeasurement(ID string, data []byte) (Measurement, error) {
	// switch data format
	switch data[2] {
	case 3:
		return NewDataFormat3(ID, data)
	case 5:
		return NewDataFormat5(ID, data)
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

func isRuuviDevice(data []byte) bool {
	if len(data) < 2 {
		return false
	}
	return binary.LittleEndian.Uint16(data[0:2]) == manufacturerDataID
}

// Scanner handles ble scanning
type Scanner interface {
	Start() chan Measurement
	Stop()
}

// scanner implements Scanner
type scanner struct {
	device  gatt.Device
	output  chan Measurement
	running bool
	started time.Time
	state   gatt.State
	mu      sync.Mutex
}

var instance *scanner

// Stop Scanner
func (s *scanner) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running == true {
		s.device.StopScanning()
		s.running = false
		close(s.output)
	}
}

// Start Scanner
func (s *scanner) Start() chan Measurement {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running == true {
		return s.output
	}
	s.device.Handle(gatt.PeripheralDiscovered(func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		if isRuuviDevice(a.ManufacturerData) {
			data, err := NewMeasurement(p.ID(), a.ManufacturerData)
			if err == nil {
				s.output <- data
			}
		}
	}))
	s.device.Init(func(d gatt.Device, state gatt.State) {
		switch state {
		case gatt.StatePoweredOn:
			d.Scan([]gatt.UUID{}, true)
		default:
			s.Stop()
		}
		s.state = state
	})
	s.started = time.Now()
	s.running = true
	return s.output
}

var once sync.Once

// NewScanner takes output buffer size as parameter
// and creates new Scanner singleton
func NewScanner(bufferSize int) (Scanner, error) {
	var err error
	once.Do(func() {
		var device gatt.Device
		device, err = gatt.NewDevice([]gatt.Option{
			gatt.LnxMaxConnections(1),
			gatt.LnxDeviceID(-1, true),
		}...)
		if err == nil {
			instance = &scanner{device: device, output: make(chan Measurement, bufferSize)}
		}
	})
	return instance, err
}

func init() {
	// Discard log message from gatt module
	log.SetOutput(ioutil.Discard)
}
