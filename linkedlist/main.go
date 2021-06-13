package main

import "fmt"

// node struct holds the data and address of next node
type Node struct {
	data int
	next *Node
}

// linkedlist struct will hold a head node and lenght of ll
type linkedlist struct {
	head   *Node
	lenght int
}

// data nodes into linkedList
/*
*1. store the current head-node into a variable(b)
*2. create a new head-node a
*3.  point the address of b in head node a
*4  Increment the size of ll
 */
// add node to linkedlist
func (l *linkedlist) Prepend(n *Node) {
	//1.
	oldNode := l.head
	//2.
	l.head = n
	//3.
	n.next = oldNode
	//4.
	l.lenght++
}

//delete given node from the linked list
/*
*1. store the current head-node into a variable
*2. itrate over the nodes
*3. once find the value skip the node and point the address of that prev. node to next node
 */

func (l *linkedlist) Delete(value int) {
	if l.lenght == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.lenght--
		return
	}
	prevNode := l.head
	for prevNode.next.data != value {
		// node not in ll
		if prevNode.next.next == nil {
			return
		}
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
	l.lenght--
}

//print ll data
func (l linkedlist) PrintData() {
	currentNode := l.head
	for l.lenght != 0 {
		fmt.Printf("%d--", currentNode.data)
		currentNode = currentNode.next
		l.lenght--
	}
	fmt.Printf("/n")

}

func main() {
	l1 := linkedlist{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	node5 := &Node{data: 5}
	node6 := &Node{data: 6}
	l1.Prepend(node1)
	l1.Prepend(node2)
	l1.Prepend(node3)
	l1.Prepend(node4)
	l1.Prepend(node5)
	l1.Prepend(node6)
	l1.PrintData()
	l1.Delete(3)
	l1.PrintData()

}
