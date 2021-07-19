package main

import "fmt"

/*
	- Session: 34
	- Date: 19.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	// 01. Khai báo
	var (
		myBool1 bool = true
		myBool2      = false // Inferred type
		myBool3      = 3 <= 5
		myBool4      = 10 != 10
	)

	// 02. Tính toán
	var result1 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var result2 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true

	// 03. Print
	fmt.Println(myBool1, myBool2, myBool3, myBool4, result1, result2)
}
