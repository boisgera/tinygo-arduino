package main

import (
	"machine"
	"time"
)

var adc = machine.ADC{Pin: machine.ADC0}

func setup() {
	machine.InitADC()
}

func main() {
	setup()
	for {
		println(adc.Get() >> 6)
		time.Sleep(1 * time.Second)
	}
}
