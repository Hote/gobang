package main

import "fmt"

func main() {
	for n := 0; n <= 5; n++ {
		go StartA() // execute Function StartA
		go StartB() // execute Function StartB
		x, y, z := StartCount(n)
		var result string = fmt.Sprintf("x=%d, y=%d,z=%d", x, y, z)
		fmt.Println(result)
	}
}

func StartA() {
	fmt.Println("I am A")
}

func StartB() {
	fmt.Println("I am B")
}

func StartCount(x int) (int, int, int) {
	return x, x | 2, x * 3
}
