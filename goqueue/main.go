package main

import (
	"fmt"

	"github.com/adrianbrad/queue"
)

func TestCircularQueue() {

	elems := []int{2, 3, 4}
	circularQueue := queue.NewCircular(elems, 3)

	containsTwo := circularQueue.Contains(2)
	fmt.Println(containsTwo) // true

	size := circularQueue.Size()
	fmt.Println(size) // 3

	empty := circularQueue.IsEmpty()
	fmt.Println(empty) // false

	if err := circularQueue.Offer(1); err != nil {
		// handle err
	}

	elem, err := circularQueue.Get()
	if err != nil {
		// handle err
	}

	fmt.Printf("elem: %d\n", elem) // elem: 1
}

func TestPriorityQueue() {

	elems := []int{2, 3, 4}

	priorityQueue := queue.NewPriority(
		elems,
		func(elem, otherElem int) bool { return elem < otherElem },
	)

	containsTwo := priorityQueue.Contains(2)
	fmt.Println(containsTwo) // true

	size := priorityQueue.Size()
	fmt.Println(size) // 3

	empty := priorityQueue.IsEmpty()
	fmt.Println(empty) // false

	if err := priorityQueue.Offer(1); err != nil {
		// handle err
	}

	elem, err := priorityQueue.Get()
	if err != nil {
		// handle err
	}

	fmt.Printf("elem: %d\n", elem) // elem: 1
}

func main() {

	TestCircularQueue()

	TestPriorityQueue()
}
