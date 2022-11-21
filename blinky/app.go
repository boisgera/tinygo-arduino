package main

import (
	"machine"
	"time"
)

var Output = machine.PinConfig{Mode: machine.PinOutput}

func main() {
	led := machine.LED // i.e. machine.D13 (a Pin)
	led.Configure(Output)
	for {
		led.Low()
		time.Sleep(1000 * time.Millisecond)

		led.High()
		time.Sleep(3000 * time.Millisecond)
	}
}
