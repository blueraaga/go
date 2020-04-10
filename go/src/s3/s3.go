package main

import ("fmt")

// passing and returning function arguments
func add1(x float64, y float64) float64 {
	return x+y
}

// condensing variable decleration
func add2(x, y float64) float64 {
	return x+y
}

// returning nultiple values
func multiple(a, b string) (string, string) {
	return a,b
}


func main() {
	var num1, num2 float64 = 5.6, 9.5	// condensed type specification with assignment
	
	num3, num4 := 5.5, 6.6	// numbers with type inference

	fmt.Println(add1(num1, num2))
	
	fmt.Println(add2(num1, num2))
	fmt.Println(add2(num3, num4))
	
	word1, word2 := "Hey", "There"	// string with type inference

	fmt.Println(multiple(word1, word2))

	var a, b int = 42, 24

	fmt.Println(add1(float64(a), float64(b)))	// typecasting
}