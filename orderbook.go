package main

import (
	"container/heap"
	"sync"
)

type PriceLevel struct {
	Price  float64
	Orders Queue
}

type OrderBook struct {
	BuyOrders  []*PriceLevel
	SellOrders []*PriceLevel
	lock       sync.Mutex
}

func (ob *OrderBook) AddOrder(order *Order) {
	ob.lock.Lock()
	defer ob.lock.Unlock()

	var priceLevel *PriceLevel
	var orders *[]*PriceLevel
	if order.Type == Buy {
		priceLevel = ob.getBuyPriceLevel(order.Price)
		orders = &ob.BuyOrders
	} else {
		priceLevel = ob.getSellPriceLevel(order.Price)
		orders = &ob.SellOrders
	}

	if priceLevel == nil {
		// This price level does not yet exist, so create it
		priceLevel = &PriceLevel{
			Price:  order.Price,
			Orders: Queue{},
		}
		*orders = append(*orders, priceLevel)
	}

	if order.Type == Cancel {
		priceLevel.Orders.Remove(order.CancelID)
	} else {
		priceLevel.Orders.Enqueue(order)
	}
}

// CancelOrder will remove the order with the given ID from the order book.
func (ob *OrderBook) CancelOrder(orderId uint64) {
	ob.lock.Lock()
	defer ob.lock.Unlock()

	for _, priceLevel := range append(ob.BuyOrders, ob.SellOrders...) {
		for index, order := range priceLevel.Orders.Items() {
			if order.ID == orderId {
				ordersSlice := priceLevel.Orders.Items()
				priceLevel.Orders = Queue{data: append(ordersSlice[:index], ordersSlice[index+1:]...)}
				break
			}
		}
	}
}

func (ob *OrderBook) getBuyPriceLevel(price float64) *PriceLevel {
	heap.Init(&PriceLevelHeap{ob.BuyOrders, LessBuy})
	return heap.Pop(&PriceLevelHeap{ob.BuyOrders, LessBuy}).(*PriceLevel)
}

func (ob *OrderBook) getSellPriceLevel(price float64) *PriceLevel {
	heap.Init(&PriceLevelHeap{ob.SellOrders, LessSell})
	return heap.Pop(&PriceLevelHeap{ob.SellOrders, LessSell}).(*PriceLevel)
}
