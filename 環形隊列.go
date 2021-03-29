package main

import (
	"errors"
	"fmt"
)

//環形列表實際使用的空間會少一格
type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int
	tail    int
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("empty")
	}
	val = this.array[this.head]
	this.head = (this.tail + 1) % this.maxSize
	return
}

func (this *CircleQueue) ListQueue() {
	//取出當前隊列有多少元素
	size := this.Size()
	if size == 0 {
		fmt.Println("empty")
	}

	tempHead := this.head
	for i := 0; i < size; i++ {
		tempHead = (tempHead + 1) % this.maxSize
	}

}

//判斷環形隊列為滿
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判斷環形隊列為滿
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出環形隊列有多少個元素
func (this *CircleQueue) Size() int {
	//!!!
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {

}
