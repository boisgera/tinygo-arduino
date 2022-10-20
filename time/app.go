package main

import "time"

func main() {
	for {
		println("Hello from Arduino! 👋")
		time.Sleep(500 * time.Millisecond)
	}
}
