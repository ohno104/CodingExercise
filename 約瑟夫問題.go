package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("wrong num")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}

		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {

	if first.Next == nil {
		fmt.Println("empty")
	}

	curBoy := first
	for {
		fmt.Printf("no: %d\n", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func JosephuGame(first *Boy, startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("empty")
		return
	}

	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//讓first移動到startNo(刪除以first為準)
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	//開始count, 然後刪除first指向的小孩
	for {
		//開始數countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		//刪除first目前指向的點
		fmt.Println("out no:", first.No)
		first = first.Next
		tail.Next = first

		if tail == first {
			break
		}
	}
	fmt.Println("last out no:", first.No)

}

func main() {
	first := AddBoy(500)
	//ShowBoy(first)
	JosephuGame(first, 20, 31)
}
