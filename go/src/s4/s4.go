package main

import "fmt"

func main() {
	x := 15
	a := &x		// pointer to x's memory addrss

	// show value and address
	fmt.Println("<value>:", x, ", <addr>:",a)

	// access value through pointer
	fmt.Println("value of *<addr>", *a)

	// change value through pointer
	*a = 5
	fmt.Println("value of <var>", x)
	fmt.Println("value of *<addr>", *a)

	// double asterisk
	*a = *a**a
	fmt.Println("value of <var>", x)
	

}