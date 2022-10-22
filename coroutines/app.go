package main

import (
	"time"
)

func ticker(c <-chan bool) {
	for {
		<-c
		println("tick")
	}
}

func main() {
	println("Arduino Timer! 👋")

	c := make(chan bool)
	go ticker(c)

	for {
		c <- true
		time.Sleep(time.Second)
	}

}
