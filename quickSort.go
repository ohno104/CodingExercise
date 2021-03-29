package main

import "fmt"

func main() {

	arr := [...]int{10, 1000, 1234, 89, 1, 8}

	quickSort(&arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSort(arr *[6]int, start, end int) {

	if start < end {
		i, j := start+1, end

		for i <= j {
			if (*arr)[i] <= (*arr)[start] {
				i++

			} else if (*arr)[j] >= (*arr)[start] {
				j--

			} else {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
				i++
				j--
			}
		}

		(*arr)[j], (*arr)[start] = (*arr)[start], (*arr)[j]

		quickSort(arr, start, j-1)
		quickSort(arr, j+1, end)
	}
}
