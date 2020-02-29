package main

import (
	"log"
	"os"
	"time"

	"github.com/peknur/ruuvitag"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

func main() {

	output := make(chan ruuvitag.Measurement, 10)
	output, err := ruuvitag.Scan()
	if err != nil {
		logger.Fatal(err)
	}
	for {
		data, ok := <-output
		if ok == false {
			logger.Println("scanner closed channel")
			break
		}
		logger.Printf("%s %d.%dc / %d %%", data.DeviceID, data.Temperature, data.TemperatureFraction, data.Humidity/2)
		time.Sleep(1 * time.Second)
	}
}
