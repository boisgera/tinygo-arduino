package main

import "time"

func main() {
	for i := 0; i < 3; i++ {
		println("Hello from Arduino! 👋")
		time.Sleep(500 * time.Millisecond)
	}
}
