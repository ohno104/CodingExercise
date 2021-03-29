package main

import "fmt"

//猴子吃桃子問題

//有一堆桃子, 猴子每天吃其中的一半再加一顆, 第十天時剩一顆桃子, 問最初有幾顆

//第10天只有一顆
//第9天=> (第10天桃子數量 + 1) * 2
//規律: 第n天=> peach(n) = (peach(n+1) + 1) * 2
func main() {
	fmt.Println(peachNum(3))
	fmt.Println(peachNum(4))
	fmt.Println(peachNum(5))
	fmt.Println(peachNum(6))

	fmt.Println(peachNum(1))
}

func peachNum(day int) int {
	if day == 10 {
		return 1

	} else {
		return (peachNum(day+1) + 1) * 2
	}
}
