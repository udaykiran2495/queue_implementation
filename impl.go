/* lets implement a queue data structure in Go for handling 100 integer values */

package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	items   []interface{}
	current int
	qstart  int
	size    int
	capacity int
	mutex   sync.Mutex
}

func (q *Queue) Enqueue(value interface{}) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.size >= q.capacity {
		fmt.Println("Queue is full, cannot enqueue:", value)
		return false
	}

	q.items[q.current] = value
	q.current++
	q.size++

	if q.current == q.capacity {
		q.current = 0
	}

	fmt.Println("Enqueued:", value)
	return true
}

func (q *Queue) Dequeue() (interface{}, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.size == 0 {
		fmt.Println("Queue is empty, cannot dequeue")
		return nil, false
	}

	value := q.items[q.qstart]
	q.qstart++
	q.size--

	if q.qstart == q.capacity {
		q.qstart = 0
	}

	return value, true
}

func main() {

	Q := Queue{
		items:    make([]interface{}, 100),
		current:  0,
		qstart:   0,
		size:     0,
		capacity: 100,
	}

	for i := 1; i <= 100; i++ {
		Q.Enqueue(i)
	}

	item, ok := Q.Dequeue()
	if ok {
		fmt.Println("Dequeued:", item)
	}

	// Thread-safe concurrent demonstration
	concurrentQueue := Queue{
		items:    make([]interface{}, 100),
		current:  0,
		qstart:   0,
		size:     0,
		capacity: 100,
	}

	var wg sync.WaitGroup

	// Producer goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(producerID int) {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				value := fmt.Sprintf("P%d-Item%d", producerID, j)
				concurrentQueue.Enqueue(value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Consumer goroutines
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for j := 1; j <= 15; j++ {
				if item, ok := concurrentQueue.Dequeue(); ok {
					fmt.Printf("Consumer %d dequeued: %v\n", consumerID, item)
				}
				time.Sleep(150 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")

	// Channel-based queue demonstration
	fmt.Println("\n=== Channel Queue Demo ===")
	channelQueue := NewChannelQueue(10)

	// Add some items
	for i := 1; i <= 5; i++ {
		if channelQueue.Enqueue(i) {
			fmt.Printf("Successfully enqueued: %d\n", i)
		}
	}

	// Remove items
	for i := 1; i <= 3; i++ {
		if item, ok := channelQueue.Dequeue(); ok {
			fmt.Println("Channel dequeued:", item)
		}
	}

	fmt.Printf("Channel queue size: %d\n", channelQueue.Size())
	channelQueue.Close()
}

func NewChannelQueue(capacity int) *ChannelQueue {
	return &ChannelQueue{
		channel:  make(chan interface{}, capacity),
		capacity: capacity,
	}
}

type ChannelQueue struct {
	channel  chan interface{}
	capacity int
}

func (cq *ChannelQueue) Enqueue(value interface{}) bool {
	select {
	case cq.channel <- value:
		return true
	default:
		return false
	}
}

func (cq *ChannelQueue) Dequeue() (interface{}, bool) {
	select {
	case value := <-cq.channel:
		return value, true
	default:
		return nil, false
	}
}

func (cq *ChannelQueue) Size() int {
	return len(cq.channel)
}

func (cq *ChannelQueue) Close() {
	close(cq.channel)
}
