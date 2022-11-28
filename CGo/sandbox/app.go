package main

// #cgo CFLAGS: -I/home/boisgera/VOYAGER/ENS/IDS/tinygo-arduino/CGo/sandbox
// #include "add.h"
import "C"

func main() {
	k := C.add(2, 3)
	println(k)
}
