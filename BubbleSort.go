package main

import "fmt"

func main() {

	arr := [...]int{24, 69, 80, 57, 13}

	bubbleSort(&arr)
}

func bubbleSort(arr *[5]int) {
	fmt.Println("排序前: ", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
		fmt.Printf("第%d次排序後: %v\n", i, (*arr))
	}
}
