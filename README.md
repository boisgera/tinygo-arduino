# TinyGo for Arduino Uno

## Hello world!

```go
package main

func main() {
        println("Hello from Arduino! 👋")
}
```

```python
#!/usr/bin/env python

import serial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    while True:
        bytes = file.readline().strip()
        try:
            print(bytes.decode("utf-8"))
        except UnicodeDecodeError:
            print(bytes)
```

## Time

```go
package main

import "time"

func main() {
    for {
        println("Hello from Arduino! 👋")
		    time.Sleep(500 * time.Millisecond)
	  }
}
```

## LED


```go
package main

import (
        "machine"
        "time"
)

var Output = machine.PinConfig{Mode: machine.PinOutput}

func main() {
        led := machine.LED // i.e. machine.D13 (a Pin)
        led.Configure(Output)
        for {
                led.Low()
                time.Sleep(1000 * time.Millisecond)

                led.High()
                time.Sleep(3000 * time.Millisecond)
        }
}
```