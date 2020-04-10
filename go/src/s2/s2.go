package main

import ("fmt"
		"math"
		"math/rand")	// required for randon number

func sr2() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func r100() {
	fmt.Println("A random number between 1 and 100 is", rand.Intn(100))
}


func main() {
	sr2()
	r100()	// Will repeat the same number as seed is not set
}