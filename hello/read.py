#!/bin/env python

# Third-Party Libraries
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
