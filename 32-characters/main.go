package main
import "fmt"

/*
	- Session: 32
	- Date: 18.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}