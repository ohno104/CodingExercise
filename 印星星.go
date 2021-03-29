package main

import (
	"fmt"
)

//   *
//  ***
// *****

//星號數 = 2*level-1
//空格數 = 總level-level

// func main() {
// 	totalLevel := 9

// 	for level := 1; level <= totalLevel; level++ {
// 		for i := 1; i <= totalLevel-level; i++ {
// 			fmt.Print(" ")
// 		}

// 		for i := 1; i <= 2*level-1; i++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

//   *
//  * *
// *   *
//*******

func main() {
	totalLevel := 9

	for level := 1; level <= totalLevel; level++ {
		for i := 1; i <= totalLevel-level; i++ {
			fmt.Print(" ")
		}

		for i := 1; i <= 2*level-1; i++ {
			if i == 1 || i == 2*level-1 || level == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
