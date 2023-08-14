package main

import "fmt"

// OrderType is an enumeration of the types of orders
type OrderType int

const (
	// Buy represents a buy order
	Buy OrderType = iota
	// Sell represents a sell order
	Sell
	// Cancel represents a cancel order action, but note that it isn't a standalone order type. Instead, the Cancel action would be associated with an existing Buy or Sell order.
	Cancel
	// Limit represents a limit order
	Limit OrderType = iota + 1
)

// Order represents a single order in the system.
type Order struct {
	ID       uint64    // Unique identifier for the order
	Price    float64   // Price of the order
	Volume   int       // Volume or quantity of the order
	Type     OrderType // Type of the order: Buy, Sell or Cancel
	CancelID uint64    // In case of Cancel action, this field denotes the ID of the order to be canceled.
}

func (o Order) String() string {
	switch o.Type {
	case Buy:
		return fmt.Sprintf("Buy Order - ID: %v, Price: %v, Volume: %v", o.ID, o.Price, o.Volume)
	case Sell:
		return fmt.Sprintf("Sell Order - ID: %v, Price: %v, Volume: %v", o.ID, o.Price, o.Volume)
	case Cancel:
		return fmt.Sprintf("Cancel Order for ID: %v", o.CancelID)
	default:
		return "Unknown Order"
	}
}
