package main

import "fmt"

const HashTableSize int = 7

//HashTable struct that holds just an array
type HashTable struct {
	array [HashTableSize]*Bucket
}

//Bucket struct is a linkedList that holds the head node and size of ll
type Bucket struct {
	head *BucketNode
}

//BucketNode is a just a node of ll that holds the data and the address of next node
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert Method will take a key and add that key into the array
func (h *HashTable) Insert(key string) {
	index := HashFunc(key)
	h.array[index].insert(key)
}

// insert method will take key and create a new node and add it to bucket
/*
* headNode
* newNode ---> headNode
* headNode = newNode
 */
func (b *Bucket) insert(k string) {
	newNode := &BucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

//Search Method will take a key and return true if present
func (h *HashTable) Serach(key string) bool {
	index := HashFunc(key)
	res := h.array[index].search(key)
	return res
}

func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

//Delete method will take key and remove the ky node
func (h *HashTable) Delete(key string) {
	index := HashFunc(key)
	h.array[index].delete(key)
}

func (b *Bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

//hash func will take the word(key) and return hashed value
func HashFunc(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % HashTableSize

}

//Init method will create a hashtable and will initalize bucket in each slot of array
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result

}

func main() {
	testHashTable := Init()
	testHashTable.Insert("danish")
	testHashTable.Insert("sh")
	testHashTable.Insert("hen")
	testHashTable.Insert("aka")
	fmt.Println(testHashTable.Serach("danish"))
	testHashTable.Delete("danish")
	fmt.Println(testHashTable.Serach("danish"))

}
