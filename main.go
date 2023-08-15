package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create an order book and trade processor
	ob := NewOrderBook()
	tp := NewTradeProcessor(ob)

	// Add initial batch of orders
	batch1 := []*Order{
		{ID: 1, Price: 100.50, Volume: 10, OrderType: Limit, Side: Buy},
		{ID: 2, Price: 101.00, Volume: 5, OrderType: Limit, Side: Buy},
		{ID: 3, Price: 99.00, Volume: 8, OrderType: Limit, Side: Sell},
		{ID: 4, Price: 98.50, Volume: 6, OrderType: Limit, Side: Sell},
	}

	fmt.Println("Adding initial batch of orders...")
	for _, order := range batch1 {
		ob.AddOrder(order)
	}

	// Display initial state
	fmt.Println("Initial State:")
	ob.Print()

	// Process orders
	fmt.Println("\nProcessing Orders...")
	tp.Start()

	// Simulate adding more orders over time to demonstrate concurrent order processing
	go func() {
		time.Sleep(2 * time.Second)

		batch2 := []*Order{
			{ID: 5, Price: 100.75, Volume: 7, OrderType: Limit, Side: Buy},
			{ID: 6, Price: 99.50, Volume: 3, OrderType: Limit, Side: Sell},
		}

		fmt.Println("\nAdding second batch of orders...")
		for _, order := range batch2 {
			ob.AddOrder(order)
			time.Sleep(500 * time.Millisecond) // Space out order additions a bit
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)

		batch3 := []*Order{
			{ID: 7, Price: 100.00, Volume: 2, OrderType: Market, Side: Buy},
			{ID: 8, Price: 0, Volume: 4, OrderType: Market, Side: Sell}, // Price is ignored for market orders
		}

		fmt.Println("\nAdding third batch of orders (market orders)...")
		for _, order := range batch3 {
			ob.AddOrder(order)
			time.Sleep(500 * time.Millisecond) // Space out order additions a bit
		}
	}()

	// Allow demo to run for a while
	time.Sleep(10 * time.Second)

	// Stop the trade processor after the demo
	tp.Stop()

	// Print final state
	fmt.Println("\nFinal State:")
	ob.Print()
}

func randomOrder() *Order {
	orderTypes := []OrderType{Limit, Market}
	sides := []Side{Buy, Sell}

	return &Order{
		ID:        uint64(rand.Intn(10000)),
		Price:     95 + rand.Float64()*10,
		Volume:    rand.Intn(10) + 1,
		OrderType: orderTypes[rand.Intn(len(orderTypes))],
		Side:      sides[rand.Intn(len(sides))],
	}
}
