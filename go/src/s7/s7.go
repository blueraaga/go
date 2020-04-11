package main

import "fmt"

const u16bit_max float64 = 65535
const kmh_mih float64 = 1.60934

type car struct {
	gas_pedal uint16
	break_pedal uint16
	steer_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bit_max)
}

func (c car) mih() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bit_max/kmh_mih)
}

func main() {
	a_car := car{gas_pedal: 65000,
				 break_pedal: 13212,
				 steer_wheel: 13322,
				 top_speed_kmh: 432.10}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mih())
}