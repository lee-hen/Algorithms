package main

import (
	"fmt"
	"log"
	"strconv"
)

type Node struct {
	Value int
	Next *Node
}

type Queue struct {
	n int
	First, Last *Node
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		log.Fatal("Queue underflow")
	}
	return q.First.Value
}

func (q *Queue) Enqueue(value int) {
	if q.IsEmpty() {
		q.First = &Node {
			Value: value,
		}
		q.Last = q.First
	} else {
		oldLast := q.Last
		oldLast.Next = &Node {
			Value: value,
		}
		q.Last = oldLast.Next
	}
	q.n++
}

func (q *Queue) Dequeue() int{
	if q.IsEmpty() {
		log.Fatal("Queue underflow")
	}
	value := q.First.Value
	q.First = q.First.Next
	//if q.IsEmpty() {
	//	q.Last = nil
	//}
	return value
}

// 1.3.37 Josephus problem. In the Josephus problem from antiquity,
// N people are in dire straits and agree to the following strategy to reduce the population.
// They arrange themselves in a circle (at positions numbered from 0 to N–1) and proceed around the circle,
// eliminating every Mth person until only one person is left. Legend has it that
// Josephus figured out where to sit to avoid being eliminated.
// Write a Queue client Josephus that takes M and N from the command line and prints out the order in which people are eliminated
// (and thus would show Josephus where to sit in the circle). % java Josephus 2 7
// 1 3 5 0 4 2 6
func main() {
	var n, m int
	_, err := fmt.Scan(&m)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}

	q := &Queue{}

	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}

	for !q.IsEmpty() {
		for j := 0; j < m-1; j++ {
			q.Enqueue(q.Dequeue())
		}
		fmt.Print(strconv.Itoa(q.Dequeue()) + " ")
	}
}
