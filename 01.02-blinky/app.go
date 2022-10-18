package main

import (
	"machine"
	"time"
)

// Source: https://tinygo.org/docs/tutorials/blinky/

func main() {
	led := machine.LED // i.e. machine.D13 (Output Pin)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
