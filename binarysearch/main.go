package main

import "fmt"

var count int

//node struct holds data(key) left child and right child
type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	//move to right child
	if k > n.key {
		// if there is no right node create one
		if n.right == nil {
			n.right = &Node{key: k}
		} else { //when the node already exist's with some value call the insert method on the right child
			n.right.Insert(k)

		}
	} else if k < n.key { //move to left child
		// if there is no left node create one
		if n.left == nil {
			n.left = &Node{key: k}
		} else { //when the node already exist's with some value call the insert method on the left child
			n.left.Insert(k)

		}
	}
}

//Search method will take a key that needs to be search and returns true if found

func (n *Node) Search(k int) bool {
	count++
	//end of the tree
	if n == nil {
		return false
	}
	//move right
	if k > n.key {
		return n.right.Search(k)
		//move left
	} else if k < n.key {
		return n.left.Search(k)
	}
	return true
}

func main() {
	n := &Node{key: 10}
	n.Insert(20)
	n.Insert(30)
	n.Insert(4)
	n.Insert(25)
	res := n.Search(30)
	fmt.Println(res, count)
}
