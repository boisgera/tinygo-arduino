package main

import (
	"machine"
	"time"
)

const (
	Input     = machine.PinConfig{Mode: machine.PinInput}
	Output    = machine.PinConfig{Mode: machine.PinOutput}
	inputPin  = machine.D2
	outputPin = machine.D4
)

func setup() {
	inputPin.Configure(Input)
	outputPin.Configure(Output)
}

func main() {
	setup()
	for {
		pressed := !inputPin.Get()
		if pressed {
			outputPin.High()
		} else {
			outputPin.Low()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
