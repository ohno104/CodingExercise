package main

import "fmt"

const (
	NA = iota
	BLACK
	BLUE
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	var chessMap [11][11]int

	chessMap[1][2] = BLACK
	chessMap[2][3] = BLUE

	// for _, v := range chessMap {
	// 	for _, v2 := range v {
	// 		fmt.Printf("%d\t", v2)
	// 	}
	// 	fmt.Println()
	// }

	//轉成稀疏數組
	var sparseArr []ValNode

	//標準稀疏數組含有一個表示紀錄二維數組的規模
	valNode := ValNode{
		row: 11,
		col: 11,
		val: NA,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != NA {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	var chessMap2 [11][11]int

	for i, node := range sparseArr {
		if i != 0 {
			chessMap2[node.row][node.col] = node.val
		}
	}

	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
