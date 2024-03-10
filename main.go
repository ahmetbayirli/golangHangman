package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	for {
		board := NewBoard()
		board.ShowBoard()
        // break
	}
	// printHangMan(6)
	// q := english.RandomWord()
	// println(q)
	// isLetter("a")
	// first := getUserData("Enter first number: ")

	//  iFirst, err := strconv.Atoi(first)
	//  if err != nil {
	//     fmt.Printf("please enter number")
	// }

	// second := getUserData("Enter second number: ")

	// iSecond, err := strconv.Atoi(second)
	// if  err != nil {
	//     fmt.Printf("please enter number..")
	// }

	// fmt.Printf("Sum of two numbers: %d" , (iFirst + iSecond))
}

// func getUserData(promtString string) string {
// 	fmt.Print(promtString)
// 	var retVal string
// 	fmt.Scanln(&retVal)
// 	return retVal
// }
