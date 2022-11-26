#!/usr/bin/env python

import serial

BAUD_RATE = 9600
SERIAL_PORT = "/dev/ttyACM0"
EOF = "\x04"

with serial.Serial(SERIAL_PORT, BAUD_RATE) as file:
    while True:
        bytes = file.readline().strip()
        if bytes == b"EOF":
            break
        try:
            print(bytes.decode("utf-8"))
        except UnicodeDecodeError:
            print(bytes)
