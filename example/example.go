package main

import (
	"log"
	"os"
	"time"

	"github.com/peknur/ruuvitag"
)

func main() {
	var logger = log.New(os.Stdout, "", log.LstdFlags)
	scanner, err := ruuvitag.OpenScanner(10)
	if err != nil {
		logger.Fatal(err)
	}
	output := scanner.Start()
	go func() {
		time.Sleep(10 * time.Second)
		logger.Printf("stopped scanner after 10 sec.")
		scanner.Stop()
	}()
	for i := 1; i <= 100; i++ {
		data, ok := <-output
		if ok == false {
			logger.Println("scanner closed channel")
			break
		}
		logger.Printf("%s[v%d] %.2f / %.2f %%", data.DeviceID(), data.Format(), data.Temperature(), data.Humidity())
	}
}
