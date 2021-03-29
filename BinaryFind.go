package main

import "fmt"

func main() {

	arr := [...]int{1, 8, 10, 89, 1000, 1234}

	binaryFind(&arr, 0, len(arr)-1, 10)
}

func binaryFind(arr *[6]int, leftIndex, rightIndex, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		binaryFind(arr, leftIndex, middle-1, findVal)

	} else if (*arr)[middle] < findVal {
		binaryFind(arr, middle+1, rightIndex, findVal)

	} else {
		fmt.Println("Find index = ", middle)
	}
}
