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
}

// func getUserData(promtString string) string {
// 	fmt.Print(promtString)
// 	var retVal string
// 	fmt.Scanln(&retVal)
// 	return retVal
// }
