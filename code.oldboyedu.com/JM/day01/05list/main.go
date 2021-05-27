package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func display(head *Node) {
	for head != nil {
		var tmp *Node = head
		fmt.Println(tmp.val)
		head = head.next
	}
}

func main() {
	//链表尾插法
	var head = new(Node)
	head.val = 0

	var tmp *Node
	tmp = head

	for i := 1; i < 10; i++ {
		//node := Node{val: i}  //Node{}方式赋值的初始化不是指针
		node := new(Node)
		node.val = i
		(*tmp).next = node
		tmp = node
	}

	display(head)

	//链表头插法
	fmt.Println("---------头插法-------------")

	var head2 = new(Node)
	head2.val = 0
	var tmp2 *Node
	tmp2 = head2

	for i := 1; i < 10; i++ {
		var node2 = Node{val: i}
		node2.next = tmp2
		tmp2 = &node2
	}
	head2 = tmp2
	display(head2)

}
