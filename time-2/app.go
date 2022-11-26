package main

import (
	"math"
	"time"
)

// Arduino Uno: 16Mhz = 62.5 nanoseconds

func main() {

	t1 := time.Now() // time.Now() takes ~200 microseconds?
	t2 := time.Now()
	t3 := time.Now()

	println("MaxInt:", math.MaxInt)
	// println("32 bit MaxInt:", 1<<(32-1)-1)
	// println("Max elapsed in secs:", float64(math.MaxInt)/1_000_000_000) // 2.14 secs!!!

	println(t1.Nanosecond())
	println(t2.Nanosecond())
	println(t3.Nanosecond())

	// now := time.Now().Nanosecond()
	// for i := 0; i < 100; i++ {
	// 	now = time.Now().Nanosecond()
	// 	println(now)
	// 	//time.Sleep(500 * time.Millisecond)
	// }
	// print(now)

	start := time.Now()

	duration1 := time.Since(start)
	duration2 := time.Since(start)
	duration3 := time.Since(start)

	println("Duration1:", duration1) // 192 ms : time.Now ~= 200 µs
	println("Duration2:", duration2) // 240 ms : time.Since ~= 50 µs. Faster!
	println("Duration3:", duration3) // 285 ms : time.Since ~= 50 µs

	time.Sleep(2 * time.Second)
	duration4 := time.Since(start)
	println("Duration4 (after 2s sleep):", duration4)
	time.Sleep(10 * time.Second)
	duration5 := time.Since(start)

	// Nota: no overflow ! Returns 3108960000, larger than MaxInt!
	// What is the type of Duration? Probably a wrapper around an integer
	// Type, but which one?
	println("Duration5 (after sleep):", duration5)
	println("Duration5 in secs:", duration5.Seconds())

	// Nota: I guess that time.Now() does NOT overflow, but calling
	// Nanosec on the result does. So we may either use durations or
	// call other methods. Nota: the doc says that Nanosecond is an
	// offset wrt the initial second, and that it should be in the
	// [0, 999999999] range (less than 1e9) but this is not what I have
	// seen. So here TinyGo seems to deviate from the Golang spec.

	// Mmmmm I can't demonstrate that issue again ?!?
	// Good in a sense if I was wrong.
	// And indeed that seems to work as expected!
	for i := 0; i < 1000; i++ {
		t := time.Now().Nanosecond()
		f := float64(t) / 1e9
		println("t:", t, f)
		if t >= 1_000_000_000 {
			println("t > 1 sec!!! *******************************")
		}
	}

	println("EOF")

}
