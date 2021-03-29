package main

import "fmt"

func main() {
	fmt.Println(fabn(3))
	fmt.Println(fabn(4))
	fmt.Println(fabn(5))
	fmt.Println(fabn(6))

	//fmt.Println(fx(5))
}

func fabn(n int) int {
	if n == 1 || n == 2 {
		return 1

	} else {
		return fabn(n-1) + fabn(n-2)
	}
}

// func fx(n int) int {
// 	if n == 1 {
// 		return 3

// 	} else {
// 		return 2*fx(n-1) + 1
// 	}
// }
