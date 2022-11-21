package main

import (
	"machine"
	"time"
)

var Input = machine.PinConfig{Mode: machine.PinInput}
var Button = machine.D2

var channel = make(chan bool, 1)

func setup() {
	Button.Configure(Input)
	Button.SetInterrupt(machine.PinToggle, func(p machine.Pin) {})
}

func ButtonPressed(c <-chan bool) {
	for {
		<-c
		println("button pressed")
	}
}

func main() {
	setup()
	println("Arduino Button Press! 👋")

	go ButtonPressed(channel)

	for {
		time.Sleep(time.Second)
	}

}
