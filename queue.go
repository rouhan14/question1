package main

import "fmt"

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]string // Change to store strings
	front int
	back  int
}

func (q *Queue) Enqueue(value string) {
	if q.size >= capacity {
		return
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) Dequeue() string {
	if q.size <= 0 {
		return ""
	}
	q.size--
	value := q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Print() {
	if q.front == q.back {
		fmt.Print("Queue is empty.")
	}
	for i := q.front; i < q.back; i++ {
		fmt.Print(q.data[i], " ")
	}
	fmt.Println()
}
