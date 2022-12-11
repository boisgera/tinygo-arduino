package main

import (
	"time"
)

// Arduino Uno: 16Mhz = 62.5 nanoseconds

func DisplayTime(t time.Time) {
	seconds := int(t.Second())
	nanoseconds := t.Nanosecond()
	println(seconds, nanoseconds)
}

func main() {
	ticker := time.NewTimer(time.Second)
	for {
		t := <-ticker.C
		println(t.Hour()) //, t.Minute(), t.Second(), t.Nanosecond())
	}
}
