package main

import (
	"fmt"
	"math"
)

/*
	- Session: 32
	- Date: 19.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	// 01. Khai báo biến
	var number1, number2 = 4, 5

	// 02.Tính toán toán học #1
	var result1 = (number1 + number2) * (number1 + number2) / 2 // Arithmetic operations

	// 03. Tính toán toán học #2
	number1++     // Tăng number1 lên 1
	number2 += 10 // Tưng number2 lên 10
	var result2 = number1 ^ number2 // Bitwise XOR


	// 04. Tính toán toán học #3
	var radius = 3.5
	var result3 = math.Pi * radius * radius // Operations on floating-point type

	fmt.Printf("number1 = %d, number2 = %d", number1,number2)
	fmt.Println("\n================================================")
	fmt.Printf("KẾT QUẢ PHÉP TOÁN\n" +
		"- result1 : %v\n- result2 : %v\n- result3 : %v\n", result1, result2, result3)
}
