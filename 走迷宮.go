package main

import (
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int
}

const (
	UNKNOW = iota
	WALL
	PASS
	BLOCK
)

func SetWay(myMap *[8][7]int, i, j int) bool {
	if myMap[6][5] == PASS {
		return true
	} else {
		if myMap[i][j] == UNKNOW {
			//先假設可以通 然後進行探測
			myMap[i][j] = PASS

			//探測策略 下右上左
			if SetWay(myMap, i+1, j) {
				return true

			} else if SetWay(myMap, i, j+1) {
				return true

			} else if SetWay(myMap, i-1, j) {
				return true

			} else if SetWay(myMap, i, j-1) {
				return true

			} else {
				myMap[i][j] = BLOCK
				return false
			}

		} else {
			return false
		}
	}
}

func main() {
	var myMap [8][7]int

	for i := 0; i < 7; i++ {
		myMap[0][i] = WALL
		myMap[7][i] = WALL
	}

	for i := 0; i < 7; i++ {
		myMap[i][0] = WALL
		myMap[i][6] = WALL
	}
	myMap[3][1] = WALL
	myMap[3][2] = WALL
	//myMap[2][2] = WALL
	//myMap[1][2] = WALL

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("探測完畢:")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
