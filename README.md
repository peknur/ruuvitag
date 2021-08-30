# RuuviTag
RuuviTag BT Scanner collects data from all nearby RuuviTag devices.   

Supports RuuviTag v3 (RAWv1) and v5 (RAWv2) formats. 

Only tested in Linux environments

## Usage
___copied from example app___  

Open BT scanner
```go
outputBufferSize := 10
scanner, err := ruuvitag.OpenScanner(outputBufferSize)
if err != nil {
    logger.Fatal(err)
}
```
Start scanner returns measurement output channel
```go
output := scanner.Start()
```

Stopping scanner will stop BT scanning and closes output channel
```go
scanner.Stop()
```

Read incoming RuuviTag measurements from channel
```go
for {
    data, ok := <-output
    if ok == false {
        logger.Println("scanner closed channel")
        break
    }
    logger.Printf("%s[v%d] %.2f / %.2f %%", data.DeviceID(), data.Format(), data.Temperature(), data.Humidity())
}
```
## Examples

See simple example from `example` folder or visit Ruuvibeacon project https://github.com/peknur/ruuvibeacon to see how I use this with Rasberry PI to collect data.
