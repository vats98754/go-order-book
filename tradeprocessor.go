package main

import (
	"fmt"
	"sync"
)

type TradeProcessor struct {
	OrderBook       *OrderBook
	OrderQueue      chan *Order
	stop            chan struct{}
	waitGroup       *sync.WaitGroup
}

func NewTradeProcessor() *TradeProcessor {
	return &TradeProcessor{
		OrderBook: &OrderBook{
			BuyOrders:  []*PriceLevel{},
			SellOrders: []*PriceLevel{},
		},
		OrderQueue: make(chan *Order),
		stop:       make(chan struct{}),
		waitGroup:  &sync.WaitGroup{},
	}
}

func (tp *TradeProcessor) AddOrder(order *Order) {
	tp.OrderQueue <- order
}

func (tp *TradeProcessor) ProcessTrades() {
	tp.waitGroup.Add(1)

	go func() {
		defer tp.waitGroup.Done()

		for {
			select {
			case order, ok := <-tp.OrderQueue:
				if !ok {
					return
				}

				tp.OrderBook.AddOrder(order)

				if order.Type != Limit {
					tp.matchMarketOrder(order)
				} else {
					tp.matchLimitOrders()
				}

			case <-tp.stop:
				return
			}
		}
	}()
}

func (tp *TradeProcessor) Stop() {
	close(tp.stop)
	tp.waitGroup.Wait()
	close(tp.OrderQueue)
}

func (tp *TradeProcessor) matchLimitOrders() {
    for _, sellPriceLevel := range tp.OrderBook.SellOrders {
        for _, buyPriceLevel := range tp.OrderBook.BuyOrders {
            if sellPriceLevel.Price <= buyPriceLevel.Price {
                sellOrder := sellPriceLevel.Orders.Dequeue()
                buyOrder := buyPriceLevel.Orders.Dequeue()

                // TODO: match the orders based on their volume, etc.
                fmt.Printf("Matched sell order %v with buy order %v\n", sellOrder, buyOrder)
                return
            }
        }
    }
}

func (tp *TradeProcessor) matchMarketOrder(order *Order) {
    var priceLevels []*PriceLevel
    if order.Type == Buy {
        priceLevels = tp.OrderBook.SellOrders
    } else {
        priceLevels = tp.OrderBook.BuyOrders
    }

    for _, priceLevel := range priceLevels {
        if len(priceLevel.Orders) > 0 {
            matchedOrder := priceLevel.Orders.Dequeue()

            // TODO: match the order based on volume, etc.
            fmt.Printf("Matched market order %v with limit order %v\n", order, matchedOrder)
            return
        }
    }
}
