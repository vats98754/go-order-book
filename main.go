package main

import (
	"time"
)

func main() {
	tradeProcessor := NewTradeProcessor()
    go tradeProcessor.ProcessTrades()

    for i := 0; i < 100; i++ {
        go func() {
            for {
                tradeProcessor.AddOrder(&Order{ID: 1, Price: 100.0, Volume: 10, Type: Buy})
                time.Sleep(1 * time.Second)
            }
        }()

        go func() {
            for {
                tradeProcessor.AddOrder(&Order{ID: 2, Price: 110.0, Volume: 5, Type: Sell})
                time.Sleep(1 * time.Second)
            }
        }()
        
        go func() {
            tradeProcessor.OrderBook.CancelOrder(1)  // Cancel order with ID 1
            time.Sleep(1 * time.Second)
        }()
    }

    time.Sleep(10 * time.Second)
    tradeProcessor.Stop()
}
