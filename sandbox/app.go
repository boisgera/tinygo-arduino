package main

import (
	"machine"
	"time"
)

var INPUT = machine.PinConfig{Mode: machine.PinInput}
var OUTPUT = machine.PinConfig{Mode: machine.PinOutput}

var togglePin = machine.D2
var toggleState = false
var toggleWasPressed = false
var lightPin = machine.D4

var adc = machine.ADC{Pin: machine.ADC0}
var pwm machine.PWM
var pwmPin = machine.D5
var ch uint8

func setup() {
	machine.LED.Configure(OUTPUT)
	togglePin.Configure(INPUT)
	lightPin.Configure(OUTPUT)

	machine.InitADC()

	pwm.Configure(machine.PWMConfig{})
	cha, err := pwm.Channel(pwmPin)
	ch = cha
	if err != nil {
		println(err)
	}
}

var ledOn = false

func blink() {
	//println("LED On:", ledOn)
	if ledOn {
		machine.LED.High()
	} else {
		machine.LED.Low()
	}
	ledOn = !ledOn
}

func toggler() {
	pressed := togglePin.Get()
	//println("pressed:", pressed)
	if !pressed && toggleWasPressed {
		toggleState = !toggleState
		//println("toggle state:", toggleState)
	}
	toggleWasPressed = pressed
}

func lighter() {
	if toggleState {
		lightPin.High()
	} else {
		lightPin.Low()
	}
}

func read() uint16 {
	return adc.Get() >> 6
}

func main() {

	setup()
	var uint32value uint32
	for { // slow loop (1 sec)
		for i := 0; i < 10; i++ { // fast loop (100 ms)
			toggler()
			lighter()
			int16value := read()
			println(int16value)
			uint32value = uint32(int16value)
			//println("u32:", uint32value)
			pwm.Set(ch, pwm.Top()*uint32value/1024)
			// println(uint32value)
			time.Sleep(100 * time.Millisecond)
		}
		blink()
	}
}
