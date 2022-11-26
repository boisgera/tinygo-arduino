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
	start := time.Now()
	// top := pwm.Top()
	// T := 3_000_000_000 // 3 seconds
	for {
		// t := time.Now().Nanosecond() % T
		// x := math.Sin(2.0 * math.Pi * float64(t) / float64(T))
		// x = 0.5 * (1.0 + x)
		// v := uint32(x * float64(top))
		println(time.Since(start).Nanoseconds())
		//pwm.Set(ch, v)
		time.Sleep(time.Millisecond)
	}
}
