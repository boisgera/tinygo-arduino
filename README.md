# TinyGo for Arduino Uno

## Hello world!


`app.go`:
```go
package main

func main() {
    println("Hello from Arduino! 👋")
}
```

```bash
tinygo flash -target=arduino app.go
```

```
avrdude: AVR device initialized and ready to accept instructions

Reading | ################################################## | 100% 0.00s

avrdude: Device signature = 0x1e950f (probably m328p)
avrdude: NOTE: "flash" memory has been specified, an erase cycle will be performed
         To disable this feature, specify the -D option.
avrdude: erasing chip
avrdude: reading input file "/tmp/tinygo287452646/main.hex"
avrdude: writing flash (1174 bytes):

Writing | ################################################## | 100% 0.20s

avrdude: 1174 bytes of flash written
avrdude: verifying flash memory against /tmp/tinygo287452646/main.hex:
avrdude: load data flash data from input file /tmp/tinygo287452646/main.hex:
avrdude: input file /tmp/tinygo287452646/main.hex contains 1174 bytes
avrdude: reading on-chip flash data:

Reading | ################################################## | 100% 0.16s

avrdude: verifying ...
avrdude: 1174 bytes of flash verified

avrdude done.  Thank you.

```

`read.py`
```python
#!/usr/bin/env python

import serial  # pyserial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    bytes = file.readline().strip()
    print(bytes.decode("utf-8"))
```

``` bash
$ ./read.py
Hello from Arduino! 👋
```

## Time

`app.go`:
```go
package main

import "time"

func main() {
    for i := 0; i < 3; i++ {
        println("Hello from Arduino! 👋")
        time.Sleep(500 * time.Millisecond)
    }
}
```

`read.py`:
```python
import serial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    while True:
        bytes = file.readline().strip()
        print(bytes.decode("utf-8"))
```

```bash
tinygo flash -target=arduino app.go
```

```bash
$ ./read.py
Hello from Arduino! 👋
Hello from Arduino! 👋
Hello from Arduino! 👋
```

`app.go`:
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

```bash
$ ./read.py
Hello from Arduino! 👋
Hello from Arduino! 👋
Hello from Arduino! 👋
Hello from Arduino! 👋
Hello from Arduino! 👋
...
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