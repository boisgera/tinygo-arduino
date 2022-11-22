package main

import (
	"machine"
	"time"
)

var pwm machine.PWM
var pwmPin = machine.D5
var period uint64
var ch uint8

func setup() {
	pwm.Configure(machine.PWMConfig{})
	period = pwm.Period() // 16_000 ns
	var err error
	ch, err = pwm.Channel(pwmPin)
	if err != nil {
		panic(err)
	}
}

func main() {
	setup()
	top := pwm.Top()
	x := top
	for {
		println(x)
		pwm.Set(ch, x)
		x = x - top/100
		if x == 0 {
			x = top
		}
		time.Sleep(10 * time.Duration(period))
	}
}
