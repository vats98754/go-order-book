package main

import "fmt"

// Queue is a basic queue based on a slice.
type Queue struct {
	data []*Order
}

// NewQueue creates a new Queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Items returns the orders slice for ranging over.
func (q *Queue) Items() []*Order {
	return q.data
}

// Enqueue adds an Order to the end of the queue.
func (q *Queue) Enqueue(order *Order) {
	q.data = append(q.data, order)
}

// Remove removes the Order from the front of the queue and returns it.
// If the queue is empty, it returns nil.
func (q *Queue) Dequeue() *Order {
	if len(q.data) == 0 {
		return nil
	}
	order := q.data[0]
	q.data = q.data[1:]
	return order
}

// Remove removes an Order based on its ID.
func (q *Queue) Remove(id uint64) *Order {
	for i, order := range q.data {
		if order.ID == id {
			q.data = append(q.data[:i], q.data[i+1:]...)
			return order
		}
	}
	return nil
}

func queueDemo() {
	// Demo for the queue and PriceLevel
	queue := NewQueue()
	order1 := &Order{ID: 1, Price: 100.50, Volume: 10}
	order2 := &Order{ID: 2, Price: 101.00, Volume: 5}

	queue.Enqueue(order1)
	queue.Enqueue(order2)

	removedOrder := queue.Dequeue()

	priceLevel := PriceLevel{
		Price:  100.50,
		Orders: *queue,
	}

	fmt.Println("Removed Order:", removedOrder)
	fmt.Println("Price Level after removal:", priceLevel)
}
