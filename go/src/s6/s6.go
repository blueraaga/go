package main

import "fmt"

type car struct {
	gas_pedal uint16
	break_pedal uint16
	steer_wheel int16
	top_speed float64
}

func main() {
	a_car := car{gas_pedal: 22222,
				 break_pedal: 13212,
				 steer_wheel: 13322,
				 top_speed: 432.10}

	fmt.Println(a_car.gas_pedal)
}