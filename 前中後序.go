package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序 先root 左子樹 右子樹
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序 先左子樹 root 右子樹
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//後序 先左子樹 右子樹 root
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "送",
	}

	left1 := &Hero{
		No:   2,
		Name: "無",
	}

	node10 := &Hero{
		No:   10,
		Name: "tom",
	}

	node12 := &Hero{
		No:   12,
		Name: "jack",
	}

	right1 := &Hero{
		No:   3,
		Name: "盧",
	}
	left1.Left = node10
	left1.Right = node12

	right2 := &Hero{
		No:   4,
		Name: "林",
	}

	root.Left = left1
	root.Right = right1
	right1.Right = right2

	PostOrder(root)
}
