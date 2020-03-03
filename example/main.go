package main

import (
	"log"
	"os"

	"github.com/peknur/ruuvitag"
)

func main() {
	var logger = log.New(os.Stdout, "", log.LstdFlags)
	output := make(chan ruuvitag.Measurement, 10)
	err := ruuvitag.Scan(output)
	if err != nil {
		close(output)
		logger.Fatal(err)
	}
	for {
		data, ok := <-output
		if ok == false {
			logger.Println("scanner closed channel")
			break
		}
		logger.Printf("%s[v%d] %.2f / %.2f %%", data.DeviceID(), data.Format(), data.Temperature(), data.Humidity())
	}
}
