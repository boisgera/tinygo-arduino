#!/usr/bin/env python
import sys

hex_filename = sys.argv[1]

with open(hex_filename, encoding="ascii") as hex_file:
    size = 0
    for line in hex_file.readlines():
        line = line.strip()
        byte_count = line[1:3]
        address = line[3:7]
        record_type = line[7:9]
        data = line[9:-2]
        checksum = line[-2:]

        print(address)
        size += int(byte_count, base=16)

print()
print(size + 1) # not sure about the +1, but matches experiments. Maybe EOF stuff?