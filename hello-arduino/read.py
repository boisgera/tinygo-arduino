#!/usr/bin/env python

import serial  # pyserial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    bytes = file.readline().strip()
    print(bytes.decode("utf-8"))
