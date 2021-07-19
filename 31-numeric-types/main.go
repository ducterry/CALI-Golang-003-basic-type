package main

import (
	"fmt"
	_ "fmt"
)

/*
	- Session: 31
	- Date: 18.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	var (
		myInt8        int8    = 97
		myInt                 = 1200
		myUint        uint    = 500
		myHexNumber           = 0xFF
		myOctalNumber         = 034
		myFloat32     float32 = 4.5
		myFloat               = 9.2
	)

	fmt.Println("Kiem tra kieu so:")
	fmt.Println("=============================")
	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)
}
