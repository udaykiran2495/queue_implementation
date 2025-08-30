/* Channel-based queue implementation in Go using goroutines and channels */

package main

type ChannelQueue struct {
	// TODO: Add buffered channel field
	// TODO: Add capacity field
	channel  chan interface{}
	capacity int
}

func NewChannelQueue(capacity int) *ChannelQueue {
	// TODO: Create and return new ChannelQueue instance
	return &ChannelQueue{
		channel:  make(chan interface{}, capacity),
		capacity: capacity,
	}
}

func (cq *ChannelQueue) Enqueue(value interface{}) bool {
	// TODO: Use select with default case for non-blocking send
	// TODO: Return true if successful, false if full
	select {
	case cq.channel <- value:
		return true
	default:
		return false
	}
}

func (cq *ChannelQueue) Dequeue() (interface{}, bool) {
	// TODO: Use select with default case for non-blocking receive
	// TODO: Return value and true if successful, nil and false if empty
	select {
	case value := <-cq.channel:
		return value, true
	default:
		return nil, false
	}
}

func (cq *ChannelQueue) Size() int {
	// TODO: Return len(channel)
	return len(cq.channel)

}

func (cq *ChannelQueue) Close() {
	// TODO: Close the channel
	close(cq.channel)
}

// func main() {
// 	// TODO: Basic demo - create queue, enqueue/dequeue items

// 	// TODO: Concurrent demo with producer/consumer goroutines
// 	// TODO: Use WaitGroup for synchronization
// }
