/* lets implement a queue data structure in Go for handling 100 integer values */

package main

import "fmt"

type Queue struct {
	items   [100]interface{}
	current int
	qstart  int
}

func (q *Queue) Enqueue(value interface{}) {
	q.items[q.current] = value

	q.current++

	if q.current == 100 {
		q.current = 0
	}

	fmt.Println("Enqueued:", value)

}

func (q *Queue) Dequeue() interface{} {
	value := q.items[q.qstart]

	q.qstart++

	if q.qstart == 100 {
		q.qstart = 0
	}

	return value
}

func main() {

	Q := Queue{}
	Q.current = 0
	Q.qstart = 0

	for i := 1; i <= 100; i++ {
		Q.Enqueue(i)
	}

	item := Q.Dequeue()

	fmt.Println(item) // should print 1

	S := Queue{}
	S.current = 0
	S.qstart = 0

	// queue other data structures

	S.Enqueue("Hello")
	S.Enqueue(3.14)

	for i := 1; i <= 5; i++ {
		S.Enqueue(i)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(S.Dequeue())
	}

}
