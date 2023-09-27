package main

import (
	"machine"
	"time"
)

func main() {
	go led1()
	led2()
}

func led1() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("+")
		led.Low()
		time.Sleep(time.Millisecond * 5000)

		println("-")
		led.High()
		time.Sleep(time.Millisecond * 5000)
	}
}

func led2() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("  +")
		led.Low()
		time.Sleep(time.Millisecond * 500)

		println("  -")
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
