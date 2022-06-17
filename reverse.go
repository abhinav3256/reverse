package main

import "fmt"

func main() {
	var input = -45

	reverse(input)

}

func reverse(input int) {

	var rem int
	var reverse int

	if input > 0 {

		for a := input; a > 0; a = a / 10 {

			rem = a % 10
			//fmt.Println(rem)
			//a = a / 10
			//fmt.Println(a)
			reverse = reverse*10 + rem

		}
	} else {

		for a := input; a < 0; a = a / 10 {

			rem = a % 10
			//fmt.Println(rem)
			//a = a / 10
			//fmt.Println(a)
			reverse = reverse*10 + rem

		}
	}
	fmt.Println(reverse)
}
