package main

import "fmt"

/*
	- Session: 35
	- Date: 19.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	// 01. Khai báo biến
	var (
		complex1 complex64 = 3.4 + 2.9i
		complex2           = 5 + 7i
		number1            = 4.5
		number2            = 7.1
		number3            = 3 + 5i
		number4            = 2 + 4i
	)

	// 02. Print
	fmt.Println(complex1, complex2)

	var result1 = complex(number1, number2)
	fmt.Println(result1)

	var result2 = number3 + number4
	var result3 = number3 - number4
	var result4 = number3 * number4
	var result5 = number3 / number4

	fmt.Println(result2, result3, result4, result5)
}
