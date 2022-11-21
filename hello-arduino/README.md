
`app.go`:

```go
package main

func main() {
	println("Hello from Arduino! 👋")
}
```


```bash
$ tinygo flash -target=arduino app.go

avrdude: AVR device initialized and ready to accept instructions

Reading | ################################################## | 100% 0.00s

avrdude: Device signature = 0x1e950f (probably m328p)
avrdude: NOTE: "flash" memory has been specified, an erase cycle will be performed
         To disable this feature, specify the -D option.
avrdude: erasing chip
avrdude: reading input file "/tmp/tinygo3179804462/main.hex"
avrdude: writing flash (1174 bytes):

Writing | ################################################## | 100% 0.20s

avrdude: 1174 bytes of flash written
avrdude: verifying flash memory against /tmp/tinygo3179804462/main.hex:
avrdude: load data flash data from input file /tmp/tinygo3179804462/main.hex:
avrdude: input file /tmp/tinygo3179804462/main.hex contains 1174 bytes
avrdude: reading on-chip flash data:

Reading | ################################################## | 100% 0.16s

avrdude: verifying ...
avrdude: 1174 bytes of flash verified

avrdude done.  Thank you.
```

Our TinyGo program is now running on the Arduino Uno ... but it prints nothing
on our computer! It sends the message on the Arduino serial port ; we need to
read it. Let's do this with Python ; we may install the module pyserial

```bash
$ pip install pyserial
```

and then implement a `read.py` program:

```python
import serial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    while True:
        bytes = file.readline()
        try:
            print(bytes.decode("utf-8"))
        except UnicodeDecodeError:
            print(bytes)
```

Now, 

```bash
$ python read.py
Hello from Arduino! 👋
```