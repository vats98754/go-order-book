package main

type PriceLevelHeap struct {
	data     []*PriceLevel
	lessFunc func(a, b *PriceLevel) bool
}

func (h PriceLevelHeap) Len() int           { return len(h.data) }
func (h PriceLevelHeap) Less(i, j int) bool { return h.lessFunc(h.data[i], h.data[j]) }
func (h PriceLevelHeap) Swap(i, j int) {
	if h.Len() > 1 {
		h.data[i], h.data[j] = h.data[j], h.data[i]
	}
}

func (h *PriceLevelHeap) Push(x interface{}) {
	h.data = append(h.data, x.(*PriceLevel))
}

func (h *PriceLevelHeap) Pop() interface{} {
	if h.Len() != 0 {
		old := h.data
		n := len(old)
		x := old[n-1]
		h.data = old[0 : n-1]
		return x
	} else {
		return new(PriceLevel)
	}
}

// LessBuy is a comparator function for the BuyOrder heap.
func LessBuy(a, b *PriceLevel) bool {
	return a.Price > b.Price // For buy orders, we want the higher price to be at the top.
}

// LessSell is a comparator function for the SellOrder heap.
func LessSell(a, b *PriceLevel) bool {
	return a.Price < b.Price // For sell orders, we want the lower price to be at the top.
}
