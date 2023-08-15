package main

import "fmt"

// OrderType is an enum of the types of orders
type OrderType int

const (
	Limit OrderType = iota
	Market
)

// Side is an enum of the types of sides
type Side int

const (
	Buy Side = iota
	Sell
	Cancel
)

// Order represents a single order in the system.
type Order struct {
	ID        uint64    // Unique identifier for the order
	Price     float64   // Price of the order
	Volume    int       // Volume or quantity of the order
	OrderType OrderType // Type of the order: Market or Limit
	Side      Side      // Side of the order: Buy, Sell or Cancel
	CancelID  uint64    // In case of Cancel action, this field denotes the ID of the order to be canceled.
}

func (o Order) String() string {
	switch o.Side {
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
