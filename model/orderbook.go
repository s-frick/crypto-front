package model

type OrderType int

const (
	ASK OrderType = iota
	BID
)

type Order struct {
	Type OrderType
	Size float64
}

type Limit struct {
	Asks []Order
	Bids []Order

	Price float64
}
