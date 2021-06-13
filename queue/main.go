package main

import "fmt"

/*
	FIFo(FIRST In First Out)
*/

// Queue is a struct that holds slice
type Queue struct {
	array []int
}

func (q *Queue) Enqueue(item int) {
	q.array = append(q.array, item)
}

func (q *Queue) Dequeue() {
	fmt.Println("Deququed", q.array[0])
	q.array = q.array[1:]
}

func main() {
	queue := Queue{}
	data := []int{1, 2, 3, 5}
	for _, d := range data {
		queue.Enqueue(d)
		fmt.Println(queue)
	}
	for i := 0; i < 4; i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

}
