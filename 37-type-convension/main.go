package main

import "fmt"

/*
	- Session: 37
	- Date: 19.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	// 01. Kịch bản 1
	var (
		number1 int64   = 4
		number2 int     = int(number1)
		number3 float64 = 6.5
		number4         = float64(number2) + number3
	)
	fmt.Println(number4)

	// 02. Kịch bản 2
	var (
		myInt   int     = 65
		myUint  uint    = uint(myInt)
		myFloat float64 = float64(myInt)
	)

	fmt.Println(myInt, myUint, myFloat)
}
