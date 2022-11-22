package main

import (
	"machine"
	"time"
)

var pwm machine.PWM
var pwmPin = machine.D5
var ch uint8

func setup() {
	pwm.Configure(machine.PWMConfig{})
	var err error
	ch, err = pwm.Channel(pwmPin)
	if err != nil {
		panic(err)
	}
}

func main() {
	setup()
	x := pwm.Top()
	for {
		pwm.Set(ch, x)
		x = x * 9 / 10
		if x == 0 {
			x = pwm.Top()
		}
		// PWM period is 16 ms
		time.Sleep(100 * time.Millisecond)
	}
}
