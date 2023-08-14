package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create an order book and trade processor
	ob := NewOrderBook()
	tp := NewTradeProcessor(ob)

	// Start the TradeProcessor in a goroutine
	go tp.ProcessTrades()

	// Use WaitGroup to wait for all goroutines
	var wg sync.WaitGroup

	// Concurrently place a couple of Buy orders
	wg.Add(2)
	go func() {
		defer wg.Done()
		buyOrder1 := &Order{ID: 1, Price: 100.50, Volume: 10, Type: Buy}
		ob.AddOrder(buyOrder1)
	}()
	go func() {
		defer wg.Done()
		buyOrder2 := &Order{ID: 2, Price: 101.00, Volume: 5, Type: Buy}
		ob.AddOrder(buyOrder2)
	}()
	wg.Wait() // Wait for Buy orders to be processed

	// Print order book after buy orders
	fmt.Println("Order Book after placing Buy Orders:")
	ob.Print()

	// Concurrently place a couple of Sell orders
	wg.Add(2)
	go func() {
		defer wg.Done()
		sellOrder1 := &Order{ID: 3, Price: 101.50, Volume: 7, Type: Sell}
		ob.AddOrder(sellOrder1)
	}()
	go func() {
		defer wg.Done()
		sellOrder2 := &Order{ID: 4, Price: 101.00, Volume: 4, Type: Sell}
		ob.AddOrder(sellOrder2)
	}()
	wg.Wait() // Wait for Sell orders to be processed

	// ... [rest of the previous demo code]

	// Concurrently add more orders to further demonstrate the concurrent nature
	wg.Add(4)
	go func() {
		defer wg.Done()
		order := &Order{ID: 9, Price: 99.50, Volume: 8, Type: Buy}
		ob.AddOrder(order)
	}()
	go func() {
		defer wg.Done()
		order := &Order{ID: 10, Price: 102.00, Volume: 9, Type: Sell}
		ob.AddOrder(order)
	}()
	go func() {
		defer wg.Done()
		order := &Order{ID: 11, Price: 100.75, Volume: 6, Type: Buy}
		ob.AddOrder(order)
	}()
	go func() {
		defer wg.Done()
		order := &Order{ID: 12, Price: 101.25, Volume: 4, Type: Sell}
		ob.AddOrder(order)
	}()
	wg.Wait() // Wait for all concurrent order additions to finish

	// Print the final state of order book
	fmt.Println("\nOrder Book after all concurrent operations:")
	ob.Print()

	// Shutdown the TradeProcessor to stop processing trades
	tp.Stop()
}
