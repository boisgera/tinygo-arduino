package main

import (
	"machine"
	"time"
)

var Input = machine.PinConfig{Mode: machine.PinInput}
var Output = machine.PinConfig{Mode: machine.PinOutput}

var ButtonPin = machine.D2
var ButtonWasPressed = false

var lightPin = machine.D4
var LightOn = false

func setup() {
	machine.LED.Configure(Output)
	ButtonPin.Configure(Input)
}

func ButtonHandler() {
	ButtonPressed := !ButtonPin.Get()
	ButtonRelease := !ButtonPressed && ButtonWasPressed
	if ButtonRelease {
		LightOn = !LightOn
	}
	ButtonWasPressed = ButtonPressed
}

func LightHandler() {
	if LightOn {
		lightPin.High()
	} else {
		lightPin.Low()
	}
}

func main() {
	setup()
	for {
		ButtonHandler()
		LightHandler()
		time.Sleep(100 * time.Millisecond)
	}
}
